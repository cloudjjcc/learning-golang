package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t1 *testing.T) {
	type args struct {
		s         string
		notExistS string
		startWith string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_Trie",
			args: args{
				s:         "apple",
				notExistS: "app",
				startWith: "app",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{}
			t.Insert(tt.args.s)
			assert.True(t1, t.Search(tt.args.s))
			assert.False(t1, t.Search(tt.args.notExistS))
			assert.True(t1, t.StartWith(tt.args.startWith))
		})
	}
}
