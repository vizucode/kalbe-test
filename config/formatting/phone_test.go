package formatting

import (
	"fmt"
	"testing"
)

func TestPhoneFormatting(t *testing.T) {
	type args struct {
		rawPhoneNumber string
	}
	tests := []struct {
		name     string
		args     args
		wantResp string
	}{
		{
			name: "test-1",
			args: args{
				rawPhoneNumber: "087708466279",
			},
			wantResp: "6287708466279",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp := PhoneFormatting(tt.args.rawPhoneNumber)
			fmt.Println(gotResp)
		})
	}
}
