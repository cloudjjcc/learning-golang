package nowcoder

import "testing"

func TestFirstAppearingOnce(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "test-01",
			args: args{str: "google"},
			want: 'l',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstAppearingOnce(tt.args.str); got != tt.want {
				t.Errorf("FirstAppearingOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}
