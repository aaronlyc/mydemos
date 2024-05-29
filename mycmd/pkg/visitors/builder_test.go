package visitors

import (
	"io"
	"reflect"
	"testing"
)

func TestNewStreamVisitor(t *testing.T) {
	type args struct {
		r      io.Reader
		source string
	}
	tests := []struct {
		name string
		args args
		want *StreamVisitor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStreamVisitor(tt.args.r, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStreamVisitor() = %v, want %v", got, tt.want)
			}
		})
	}
}
