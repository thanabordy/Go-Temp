package main

func main() {
	// creds, err := credentials.NewClientTLSFromFile("storage/cert/certificate.crt", "*")
	// if err != nil {
	// 	log.Fatalf("failed to load credentials: %v", err)
	// }
	// opts := []grpc.DialOption{
	// 	// credentials.
	// 	grpc.WithTransportCredentials(creds),
	// }
	// conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// uc := pb.NewUserClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := uc.GetUser(ctx, &pb.UserID{Id: 11})
	// if err != nil {
	// 	log.Fatalf("Could not user: %v", err)
	// }
	// log.Printf("User: %v", r)
}
