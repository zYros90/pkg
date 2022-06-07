package hash

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name          string
		args          args
		pswdToCompare string
		wantErr       bool
		wantMismatch  bool
	}{
		{
			"passwd",
			args{password: "hallo123"},
			"hallo123",
			false,
			false,
		},
		{
			"passwd",
			args{password: "pj&UPPA$8F42L^SV"},
			"pj&UPPA$8F42L^SV",
			false,
			false,
		},
		{
			"passwd",
			args{password: "d7tjZurMbSWWZJxS"},
			"differentPass",
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() = %v", got)
			}

			ok := CheckPasswordHash(tt.pswdToCompare, got)
			if !ok && !tt.wantMismatch {
				t.Errorf("HashPassword() = %v, want %v", got, tt.pswdToCompare)
			}
		})
	}
}
