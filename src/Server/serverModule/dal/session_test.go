package dal

import "testing"

func Test_session_HandleCommand(t *testing.T) {
	type args struct {
		command int
	}
	tests := []struct {
		name string
		this *session
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.HandleCommand(tt.args.command)
		})
	}
}
