package task

import (
	"testing"
)

func Test_initBaseDir(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "initBaseDir",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initBaseDir()
		})
	}
}
