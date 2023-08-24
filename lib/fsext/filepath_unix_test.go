//go:build unix

package fsext_test

import (
	"testing"

	"github.com/gesemaya/k6/lib/fsext"
	"github.com/stretchr/testify/assert"
)

func TestJoinFilePath(t *testing.T) {
	t.Parallel()

	type args struct {
		b string
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "join root and some catalog",
			args: args{
				b: "/",
				p: "test",
			},
			want: "/test",
		},
		{
			name: "join root and some catalog with leading slash",
			args: args{
				b: "/",
				p: "/test",
			},
			want: "/test",
		},
		{
			name: "join root and some catalog with several leading slash",
			args: args{
				b: "/",
				p: "//test",
			},
			want: "/test",
		},
		{
			name: "join catalog and some other catalog",
			args: args{
				b: "/path/to",
				p: "test",
			},
			want: "/path/to/test",
		},
		{
			name: "join catalog and some other catalog with leading slash",
			args: args{
				b: "/path/to",
				p: "/test",
			},
			want: "/path/to/test",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, fsext.JoinFilePath(tt.args.b, tt.args.p))
		})
	}
}