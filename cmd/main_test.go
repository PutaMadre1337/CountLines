package main_test

import (
	main "CountLines/cmd"
	"testing"
)

func Test_expandPath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "home directory",
			path:    "/home/username",
			want:    "/home/username",
			wantErr: false,
		},
		{
			name:    "current directory",
			path:    ".",
			want:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd",
			wantErr: false,
		},
		{
			name:    "absolute path",
			path:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/",
			want:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines",
			wantErr: false,
		},
		{
			name:    "relative path",
			path:    "testfiles",
			want:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles",
			wantErr: false,
		},
		{
			name:    "non exist path",
			path:    "internal",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := main.ExpandPath(tt.path)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("\nReturned error = %v", gotErr)
				return
			}
			if got != tt.want {
				t.Errorf("\nReturned path: %v\nWant: %v", got, tt.want)
			}
		})
	}
}