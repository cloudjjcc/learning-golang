package bisect

import "testing"

func TestBisect(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name      string
		args      args
		wantLeft  int
		wantRight int
	}{
		{
			name: "TestBisect",
			args: args{
				nums:   []int{1, 3, 5, 5, 6, 8, 10},
				target: 5,
			},
			wantLeft:  2,
			wantRight: 3,
		},
		{
			name: "TestBisect",
			args: args{
				nums:   []int{1, 3, 5, 5, 6, 8, 10},
				target: 4,
			},
			wantLeft:  2,
			wantRight: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft := LeftBound(tt.args.nums, tt.args.target)
			gotRight := RightBound(tt.args.nums, tt.args.target)
			if gotLeft != tt.wantLeft {
				t.Errorf("TestBisect() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if gotRight != tt.wantRight {
				t.Errorf("TestBisect() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}
