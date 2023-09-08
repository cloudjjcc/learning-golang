package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_replaceWords(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_replaceWords",
			args: args{
				dictionary: []string{"cat", "bat", "rat"},
				sentence:   "the cattle was rattled by the battery",
			},
			want: "the cat was rat by the bat",
		},
		{
			name: "Test_replaceWords",
			args: args{
				dictionary: []string{"a", "aa", "aaa", "aaaa"},
				sentence:   "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa",
			},
			want: "a a a a a a a a bbb baba a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, replaceWords(tt.args.dictionary, tt.args.sentence), "replaceWords(%v, %v)", tt.args.dictionary, tt.args.sentence)
		})
	}
}
