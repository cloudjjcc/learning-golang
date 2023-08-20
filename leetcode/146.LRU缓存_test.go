package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		cap int
		op  [][3]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestNewLRUCache",
			args: args{
				cap: 2,
				op: [][3]int{
					{0, 1, 1},
					{0, 2, 2},
					{1, 1, 1},
					{0, 3, 3},
					{1, 2, -1},
					{0, 4, 4},
					{1, 1, -1},
					{1, 3, 3},
					{1, 4, 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lruCache := NewLRUCache(tt.args.cap)
			for _, v := range tt.args.op {
				if v[0] == 0 {
					lruCache.Put(v[1], v[2])
					t.Log(lruCache.Print())
				} else {
					assert.Equal(t, v[2], lruCache.Get(v[1]))
				}
			}
		})
	}
}
