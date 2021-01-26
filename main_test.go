package main

import (
	"os"
	"testing"
)

func Test_checkPort(t *testing.T) {
	tests := []struct {
		name string
		port string
		want bool
	}{
		{
			name: "test port 5000",
			port: "5000",
			want: true,
		},
		{
			name: "test port 100000",
			port: "100000",
			want: false,
		},
		{
			name: "test negative port",
			port: "-2000",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			servePort = tt.port
			if got := checkPort(); got != tt.want {
				t.Errorf("checkPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_args(t *testing.T) {
	t.Run("test args parsing", func(t *testing.T) {
		port := "7000"
		os.Args = append(os.Args, "-p", port)
		processFlags()
		if servePort != port {
			t.Errorf("Flag -p was not properly parsed. Got: %s, want: %s", servePort, port)
		}
	})
}

func Test_envvars(t *testing.T) {
	t.Run("test envvars", func(t *testing.T) {
		port := "9000"
		os.Setenv("HTTP_PORT", port)
		processEnvVars()
		if servePort != port {
			t.Errorf("Envvar HTTP_PORT was not properly readed. Got: %s, want: %s", servePort, "7000")
		}
	})
}
