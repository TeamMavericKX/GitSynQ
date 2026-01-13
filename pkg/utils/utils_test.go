package utils

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{500, "500 B"},
		{1024, "1.0 KB"},
		{1024 * 1024, "1.0 MB"},
		{1024 * 1024 * 1024, "1.0 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatBytes(tt.bytes); got != tt.want {
				t.Errorf("FormatBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
