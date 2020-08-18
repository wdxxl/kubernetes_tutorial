package main

import "testing"

func Test_printMessage1(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				msg: "Hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printMessage(tt.args.msg); got != tt.want {
				t.Errorf("printMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}