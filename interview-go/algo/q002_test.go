package algo

import (
	"testing"
)

func Test_coinChange1(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins:  []int{1, 5, 10},
				amount: 13,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange2(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins:  []int{1, 5, 10},
				amount: 13,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				coins:  []int{1, 2, 5, 10},
				amount: 163,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange3(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				coins:  []int{1, 2, 5, 10},
				amount: 163,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange3(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange3() = %v, want %v", got, tt.want)
			}
		})
	}
}
