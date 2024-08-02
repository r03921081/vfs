package util

import (
	"r03921081/vfs/constant"
	"regexp"
	"strings"
	"testing"
)

func TestIsValidInput(t *testing.T) {
	type args struct {
		input string
		r     *regexp.Regexp
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid command",
			args: args{
				input: "list-folders user1 --sort-name asc",
				r:     ValidCommand,
			},
			want: true,
		},
		{
			name: "Valid command length",
			args: args{
				input: strings.Repeat("a", constant.MaxLengthCommand),
				r:     ValidCommand,
			},
			want: true,
		},
		{
			name: "Invalid command",
			args: args{
				input: "# list-folders user1 --sort-name asc",
				r:     ValidCommand,
			},
			want: false,
		},
		{
			name: "Command too long",
			args: args{
				input: strings.Repeat("a", constant.MaxLengthCommand+1),
				r:     ValidCommand,
			},
			want: false,
		},
		{
			name: "Command too short",
			args: args{
				input: "",
				r:     ValidCommand,
			},
			want: false,
		},
		{
			name: "Valid name",
			args: args{
				input: "user1",
				r:     ValidName,
			},
			want: true,
		},
		{
			name: "Valid name length",
			args: args{
				input: strings.Repeat("a", constant.MaxLengthName),
				r:     ValidName,
			},
			want: true,
		},
		{
			name: "Invalid name",
			args: args{
				input: "user 1",
				r:     ValidName,
			},
			want: false,
		},
		{
			name: "Name too long",
			args: args{
				input: strings.Repeat("a", constant.MaxLengthName+1),
				r:     ValidName,
			},
			want: false,
		},
		{
			name: "Name too short",
			args: args{
				input: "",
				r:     ValidName,
			},
			want: false,
		},
		{
			name: "Valid description",
			args: args{
				input: strings.Repeat("a", constant.MaxLengthDescription),
				r:     ValidDescription,
			},
			want: true,
		},
		{
			name: "Invalid description",
			args: args{
				input: "~description 1",
				r:     ValidDescription,
			},
			want: false,
		},
		{
			name: "Description too long",
			args: args{
				input: strings.Repeat("a", constant.MaxLengthDescription+1),
				r:     ValidDescription,
			},
			want: false,
		},
		{
			name: "Description too short",
			args: args{
				input: "",
				r:     ValidDescription,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidInput(tt.args.input, tt.args.r); got != tt.want {
				t.Errorf("IsValidInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
