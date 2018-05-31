// package test

// import (
// 	"testing"

// 	dal "../src/Server/serverModule/dal"
// )

// func TestMul(t *testing.T) {
// 	type args struct {
// 		x int
// 		y int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{{"ad", args{2, 3}, 6}}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := dal.Mul(tt.args.x, tt.args.y); got != tt.want {
// 				t.Errorf("Mul() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_session_HandleCommand(t *testing.T) {
// 	type args struct {
// 		command int
// 	}
// 	tests := []struct {
// 		name string
// 		this *session
// 		args args
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.this.HandleCommand(tt.args.command)
// 		})
// 	}
// }