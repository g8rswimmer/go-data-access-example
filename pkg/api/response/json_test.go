package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
	type args struct {
		status int
		body   interface{}
	}
	tests := []struct {
		name   string
		args   args
		status int
		body   map[string]interface{}
	}{
		{
			name: "Success",
			args: args{
				status: http.StatusOK,
				body: map[string]interface{}{
					"a": "b",
					"z": true,
				},
			},
			status: http.StatusOK,
			body: map[string]interface{}{
				"a": "b",
				"z": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()

			JSON(writer, tt.args.status, tt.args.body)

			if writer.Result().StatusCode != tt.status {
				t.Errorf("JSON() got %v want %v", writer.Result().StatusCode, tt.status)
				return
			}

			var body map[string]interface{}
			if err := json.NewDecoder(writer.Body).Decode(&body); err != nil {
				t.Errorf("JSON decode err %v", err)
				return
			}

			if reflect.DeepEqual(body, tt.body) == false {
				t.Errorf("JSON() got %v want %v", body, tt.body)
			}
		})
	}
}
