package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			want:    "abc123",
			wantErr: nil,
		},
		{
			name:    "no authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(err, tt.wantErr) {
				t.Fatalf("expected: %v, got: %v, expected error: %v, got error: %v", tt.want, got, tt.wantErr, err)
			}
		})
	}
}
