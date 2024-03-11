package main

import (
	"github.com/acework2u/air-iot-app-api-service/config"
	"testing"
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

	startGinServer(config.Config{Origin: "*"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := SumCal(tt.args.x, tt.args.y); got != tt.want {
			// 	t.Errorf("SamCal = %v,%v", got, tt.want)
			// }
		})
	}

}
