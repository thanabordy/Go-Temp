package console

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func keygen() *cobra.Command {
	var b2args, b3args bool
	cmd := &cobra.Command{
		Use: "keygen <key> [bits=1024] [years=5]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a key argument")
			}
			switch key := args[0]; key {
			case "rsa", "RSA", "ecdsa", "ECDSA":
			default:
				return fmt.Errorf("Unsupported signature \"%s\"", key)
			}
			if len(args) == 2 {
				if _, err := strconv.Atoi(args[1]); err != nil {
					return err
				}
				b2args = true
			}
			if len(args) == 3 {
				if _, err := strconv.Atoi(args[2]); err != nil {
					return err
				}
				b3args = true
			}
			if len(args) > 3 {
				return errors.New("Need 3 argument only")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			bits := 1024
			years := 5
			if b2args {
				bits, _ = strconv.Atoi(args[1])
			}
			if b3args {
				years, _ = strconv.Atoi(args[2])
			}
			pk, err := privateKey(args[0], bits)
			if err != nil {
				log.Fatal(err)
			}
			ran, err := rand.Int(rand.Reader, big.NewInt(2^63))
			if err != nil {
				log.Fatal(err)
			}
			template := x509.Certificate{
				SerialNumber: ran,
				DNSNames:     []string{"*"},
				NotBefore:    time.Now(),
				NotAfter:     time.Now().AddDate(years, 0, 0),

				KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
				ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
				BasicConstraintsValid: true,
			}
			pub, err := publicKey(pk)
			if err != nil {
				log.Fatalf("Failed to create public key: %s", err)
			}
			derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, pub, pk)
			if err != nil {
				log.Fatalf("Failed to create certificate: %s", err)
			}
			out := &bytes.Buffer{}
			pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
			err = ioutil.WriteFile("storage/cert/certificate.crt", out.Bytes(), 0600)
			if err != nil {
				log.Fatalf("Failed to create certificate file: %s", err)
			}
			fmt.Println("Created certificate")
			out.Reset()
			pem.Encode(out, pemBlockForKey(pk))
			err = ioutil.WriteFile("storage/cert/privatekey.key", out.Bytes(), 0600)
			if err != nil {
				log.Fatalf("Failed to create privatekey file: %s", err)
			}
			fmt.Println("Created private key")
		},
	}
	return cmd
}

func privateKey(key string, bits int) (crypto.PrivateKey, error) {
	switch key {
	case "rsa", "RSA":
		return rsa.GenerateKey(rand.Reader, bits)
	case "ecdsa", "ECDSA":
		return ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		return nil, fmt.Errorf("Unsupported signature \"%s\"", key)
	}
}

func publicKey(priv crypto.PrivateKey) (crypto.PublicKey, error) {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey, nil
	case *ecdsa.PrivateKey:
		return &k.PublicKey, nil
	default:
		return nil, errors.New("Can't get public key")
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}
