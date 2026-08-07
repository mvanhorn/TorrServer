package torrstor

import (
	"math"
	"testing"
)

func TestGetPieceIndex(t *testing.T) {
	tests := []struct {
		name        string
		offset      int64
		fileOffset  int64
		pieceLength int64
		pieceCount  int
		want        int
	}{
		{
			name:        "interior offset with file offset",
			offset:      512,
			fileOffset:  1024,
			pieceLength: 1024,
			pieceCount:  4,
			want:        1,
		},
		{
			name:        "ordinary interior offset",
			offset:      2500,
			pieceLength: 1024,
			pieceCount:  4,
			want:        2,
		},
		{
			name:        "negative absolute position",
			offset:      -2048,
			fileOffset:  1024,
			pieceLength: 1024,
			pieceCount:  4,
			want:        0,
		},
		{
			name:        "exact torrent end",
			offset:      4096,
			pieceLength: 1024,
			pieceCount:  4,
			want:        3,
		},
		{
			name:        "beyond torrent end",
			offset:      8192,
			pieceLength: 1024,
			pieceCount:  4,
			want:        3,
		},
		{
			name:        "positive offset overflow",
			offset:      math.MaxInt64,
			fileOffset:  1,
			pieceLength: 1024,
			pieceCount:  4,
			want:        3,
		},
		{
			name:        "negative offset overflow",
			offset:      math.MinInt64,
			fileOffset:  -1,
			pieceLength: 1024,
			pieceCount:  4,
			want:        0,
		},
		{
			name:        "one piece start",
			pieceLength: 1024,
			pieceCount:  1,
			want:        0,
		},
		{
			name:        "one piece end",
			offset:      1024,
			pieceLength: 1024,
			pieceCount:  1,
			want:        0,
		},
		{
			name:        "uninitialized piece length",
			offset:      1024,
			pieceLength: 0,
			pieceCount:  4,
			want:        0,
		},
		{
			name:        "uninitialized piece count",
			offset:      1024,
			pieceLength: 1024,
			pieceCount:  0,
			want:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := &Cache{
				pieceLength: tt.pieceLength,
				pieceCount:  tt.pieceCount,
			}

			if got := getPieceIndex(tt.offset, tt.fileOffset, cache); got != tt.want {
				t.Fatalf("getPieceIndex(%d, %d) = %d, want %d", tt.offset, tt.fileOffset, got, tt.want)
			}
		})
	}
}

func TestGetPieceIndexWithNilCache(t *testing.T) {
	if got := getPieceIndex(1024, 0, nil); got != 0 {
		t.Fatalf("getPieceIndex with nil cache = %d, want 0", got)
	}
}
