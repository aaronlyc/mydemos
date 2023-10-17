package algo

import "testing"

func Test_cuttingBamboo(t *testing.T) {
	type args struct {
		bamboo_len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				bamboo_len: 3,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				bamboo_len: 4,
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				bamboo_len: 12,
			},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingBamboo(tt.args.bamboo_len); got != tt.want {
				t.Errorf("cuttingBamboo() = %v, want %v", got, tt.want)
			}
		})
	}
}
