package algo

import "testing"

func Test_bestTimeBuySellStock(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestTimeBuySellStock(tt.args.prices); got != tt.want {
				t.Errorf("bestTimeBuySellStock() = %v, want %v", got, tt.want)
			}
		})
	}
}
