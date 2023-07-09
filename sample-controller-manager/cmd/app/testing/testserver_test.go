package testing

import (
	"context"
	"reflect"
	"testing"
)

func TestStartTestServerOrDie(t *testing.T) {
	type args struct {
		ctx   context.Context
		flags []string
	}
	tests := []struct {
		name string
		args args
		want *TestServer
	}{
		{
			name: "test no flags",
			args: args{
				ctx: context.TODO(),
				flags: []string{
					"--kubeconfig=/Users/aaron/.kube/config",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartTestServerOrDie(tt.args.ctx, tt.args.flags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartTestServerOrDie() = %v, want %v", got, tt.want)
			}
		})
	}
}
