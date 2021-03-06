package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("%v isMatch() = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
