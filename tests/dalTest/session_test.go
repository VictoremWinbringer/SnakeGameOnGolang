package tests

import (
	"testing"

	dal "../../src/Server/serverModule/dal"
)

func TestMul(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{"ad", args{2, 3}, 6}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dal.Mul(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
