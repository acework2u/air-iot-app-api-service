package main

import (
	"testing"

	_ "golang.org/x/lint"
)

func Test_SummaryCal(t *testing.T) {
	type args struct {
		x float32
		y float32
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "with 5 and 6",
			args: args{
				x: 5,
				y: 6,
			},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := SumCal(tt.args.x, tt.args.y); got != tt.want {
			// 	t.Errorf("SamCal = %v,%v", got, tt.want)
			// }
		})
	}

}
