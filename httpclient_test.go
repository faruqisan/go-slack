package slack

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func mockHandler(httpStatus int, responseBody []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		w.Write(responseBody)
	}
}

func TestEngine_do(t *testing.T) {

	type fields struct {
		opt    Option
		client *http.Client
	}
	type args struct {
		method string
		param  []byte
		path   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test success",
			fields: fields{
				opt: Option{
					WebHookURLs: []string{"http://foo.bar"},
				},
				client: &http.Client{},
			},
			args: args{
				method: http.MethodGet,
			},
			want: []byte("ok"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{
				opt:    tt.fields.opt,
				client: tt.fields.client,
			}

			handler := mockHandler(http.StatusOK, []byte("ok"))
			server := httptest.NewServer(http.HandlerFunc(handler))
			defer server.Close()

			got, err := e.do(tt.args.method, server.URL, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Engine.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Engine.do() = %v, want %v", got, tt.want)
			}
		})
	}
}
