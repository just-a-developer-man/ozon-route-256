package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func BenchmarkSolution(t *testing.B) {
	out, _ := os.Open(os.DevNull)

	type args struct {
		in  io.ReadCloser
		out io.Writer
	}
	type test struct {
		name string
		args args
	}
	tests := make([]test, 0, 30)

	for i := 1; i <= 30; i++ {
		if i == 15 {
			continue
		}
		file, err := os.Open(fmt.Sprintf("/home/computer/Downloads/242/%d", i))
		if err != nil {
			panic(err)
		}
		tests = append(tests, test{
			name: fmt.Sprintf("%d", i),
			args: args{
				in:  file,
				out: out,
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			Solution(tt.args.in, tt.args.out)
			tt.args.in.Close()
		})
	}
}
