package huawei

import "testing"

func Test_sortGame(t *testing.T) {
	type args struct {
		records []*record
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "",
			args: args{records: []*record{
				{1, 0, 100},
				{2, 0, 200},
				{3, 0, 300},
				{4, 0, 200},
				{5, 0, 200},
			}},
			want:  0,
			want1: 150,
		},
		{
			name: "",
			args: args{records: []*record{
				{1, 0, 223},
				{2, 0, 323},
				{3, 2, 1203},
			}},
			want:  0,
			want1: 105,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := bossIncome(tt.args.records)
			if got != tt.want {
				t.Errorf("sortGame() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("sortGame() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
