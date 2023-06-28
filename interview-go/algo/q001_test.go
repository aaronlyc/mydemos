package algo

import (
	"testing"
)

func Test_fib1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "10",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib1(tt.args.n); got != tt.want {
				t.Errorf("fib1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{
				n: 100,
			},
			want: 3736710778780434371,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("fib2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10",
			args: args{
				n: 10,
			},
			want: 55,
		},
		{
			name: "100",
			args: args{
				n: 100,
			},
			want: 3736710778780434371,
		},
		{
			name: "1000",
			args: args{
				n: 1000,
			},
			want: 817770325994397771,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib3(tt.args.n); got != tt.want {
				t.Errorf("fib3() = %v, want %v", got, tt.want)
			}
		})
	}
}
