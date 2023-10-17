package question

import "testing"

func Test_replaceBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "normal",
			args: args{
				s: "ada r dd ",
			},
			want:  "ada%20r%20dd%20",
			want1: true,
		},
		{
			name: "string out of bounds",
			args: args{
				s: "ada r dd dddfc sdsd dgfg aad",
			},
			want:  "",
			want1: false,
		},
		{
			name: "string has number",
			args: args{
				s: "ada 124 sd",
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := replaceBlank(tt.args.s)
			if got != tt.want {
				t.Errorf("replaceBlank() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("replaceBlank() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
