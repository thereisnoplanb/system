package system

import (
	"testing"
)

func TestGetFunctionName(t *testing.T) {
	type args struct {
		function []interface{}
	}
	var alaMaKota bool
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "anonymous",
			args:    args{},
			want:    "func1",
			wantErr: false,
		},
		{
			name: "TestGetFunctionName",
			args: args{
				function: []interface{}{TestGetFunctionName},
			},
			want:    "TestGetFunctionName",
			wantErr: false,
		},
		{
			name: "alaMaKota",
			args: args{
				function: []interface{}{alaMaKota},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFunctionName(tt.args.function...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFunctionName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
