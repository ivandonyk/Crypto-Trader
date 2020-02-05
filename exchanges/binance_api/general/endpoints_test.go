package general

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func testServer() *httptest.Server {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		switch r.URL.Path {
		case "/api/v3/time":
			f, err := ioutil.ReadFile("testdata/serverTime.json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(f)
		case "/api/v3/ping":
			f, err := ioutil.ReadFile("testdata/ping.json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			w.Write(f)
		default:
			http.Error(w, errors.New("unknown path").Error(), http.StatusInternalServerError)

		}
	}))

	return ts

}
func TestGeneral_CheckServiceTime(t *testing.T) {
	tests := []struct {
		name    string
		gen     *General
		want    *checkServerTimeResponse
		wantErr bool
	}{
		{
			name:    "server check test #1",
			gen:     &General{baseURL: testServer().URL},
			want:    &checkServerTimeResponse{ServerTime: 10000},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.gen.CheckServiceTime()
			if (err != nil) != tt.wantErr {
				t.Fatalf("unknown error occurred when trying check server time %v", err)
			}

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("CheckServerTime() failed, expected=%+v got=%+v", result, tt.want)
			}

		})
	}
}

func TestGeneral_GetPing(t *testing.T) {
	tests := []struct {
		name    string
		gen     *General
		want    *pingResponse
		wantErr bool
	}{
		{
			name:    "server check test #1",
			gen:     &General{baseURL: testServer().URL},
			want:    &pingResponse{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.gen.GetPing()
			if (err != nil) != tt.wantErr {
				t.Fatalf("unknown error occurred when trying check server time %v", err)
			}

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("CheckServerTime() failed, expected=%+v got=%+v", result, tt.want)
			}

		})
	}

}
