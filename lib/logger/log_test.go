package logger

import (
	"log/syslog"
	"reflect"
	"testing"
)

func Test_getOutputLogLevel(t *testing.T) {
	tests := []struct {
		name string
		want syslog.Priority
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOutputLogLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOutputLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	type args struct {
		level     Priority
		calldepth int
		v         []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "output_1",
			args: args{
				LogDebug,
				2,
				[]interface{}{
					"test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Output(tt.args.level, tt.args.calldepth, tt.args.v...)
		})
	}
}

func TestOutputf(t *testing.T) {
	type args struct {
		level     Priority
		calldepth int
		format    string
		v         []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "output_1",
			args: args{
				LogDebug,
				2,
				"%s",
				[]interface{}{
					"test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Outputf(tt.args.level, tt.args.calldepth, tt.args.format, tt.args.v...)
		})
	}
}
