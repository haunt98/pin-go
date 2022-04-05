package pin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinToString(t *testing.T) {
	tests := []struct {
		name   string
		number int64
		length int
		want   string
	}{
		{
			name:   "0001",
			number: 1,
			length: 4,
			want:   "0001",
		},
		{
			name:   "000069",
			number: 69,
			length: 6,
			want:   "000069",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := pinToString(tc.number, tc.length)
			assert.Equal(t, tc.want, got)
		})
	}
}
