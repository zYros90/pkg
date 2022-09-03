package jwt

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerateJWTHS256(t *testing.T) {
	type args struct {
		secret    string
		claimsMap map[string]string
		expire    time.Duration
	}
	tests := []struct {
		name              string
		args              args
		wantClaims        map[string]string
		wantGenerateErr   bool
		wantValidateErr   bool
		wantClaimMismatch bool
	}{
		{
			"jwt",
			args{
				secret:    "mysupersecretkey",
				claimsMap: map[string]string{"username": "test"},
				expire:    0,
			},
			map[string]string{"username": "test"},
			false,
			false,
			false,
		},
		{
			"jwt",
			args{
				secret:    "mysupersecretkey",
				claimsMap: map[string]string{"username": "test"},
				expire:    0,
			},
			map[string]string{"user_name": "test"},
			false,
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := New()
			token, err := j.GenerateJWTHS256(tt.args.secret, tt.args.claimsMap, tt.args.expire)
			if (err != nil) != tt.wantGenerateErr {
				t.Errorf("GenerateJWTHS256() error = %v, wantErr %v", err, tt.wantGenerateErr)
				return
			}
			claimMap, err := j.ValidateJWTHS256([]byte(tt.args.secret), token)
			if (err != nil) != tt.wantValidateErr {
				t.Errorf("ValidateJWTHS256() error = %v, wantErr %v", err, tt.wantGenerateErr)
				return
			}
			if !reflect.DeepEqual(claimMap, tt.wantClaims) != tt.wantClaimMismatch {
				t.Errorf("ValidateJWTHS256() claims = %v, wantClaims %v", claimMap, tt.wantClaims)
			}
		})
	}
}
