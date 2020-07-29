package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/g8rswimmer/go-data-access-example/pkg/model"
	"github.com/gorilla/mux"
)

func TestHandler_Create(t *testing.T) {
	type fields struct {
		UserDAO DAO
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		status int
		body   interface{}
	}{
		{
			name: "created",
			fields: fields{
				UserDAO: &mockUserDAO{
					user: &model.UserEntity{
						Entity: model.Entity{
							ID:        "1234",
							CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
						},
						User: model.User{
							FirstName: "test",
							LastName:  "testison",
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					u := model.User{
						FirstName: "test",
						LastName:  "testison",
					}
					enc, _ := json.Marshal(u)
					return httptest.NewRequest(http.MethodPost, "http://www.google.com", bytes.NewReader(enc))
				}(),
			},
			status: http.StatusCreated,
			body: model.UserEntity{
				Entity: model.Entity{
					ID:        "1234",
					CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
				},
				User: model.User{
					FirstName: "test",
					LastName:  "testison",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				UserDAO: tt.fields.UserDAO,
			}
			writer := httptest.NewRecorder()
			handler := h.create()
			handler.ServeHTTP(writer, tt.args.req)

			if writer.Result().StatusCode != tt.status {
				t.Errorf("Handler.Create() = %v, want %v", writer.Result().StatusCode, tt.status)
				return
			}

			var bodyMap map[string]interface{}
			if err := json.NewDecoder(writer.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.Create() = json body decode error %v", err)
				return
			}

			var wantBodyMap map[string]interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.Create() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}

func TestHandler_FetchByID(t *testing.T) {
	type fields struct {
		UserDAO DAO
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		status int
		body   interface{}
	}{
		{
			name: "fetched",
			fields: fields{
				UserDAO: &mockUserDAO{
					user: &model.UserEntity{
						Entity: model.Entity{
							ID:        "1234",
							CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
						},
						User: model.User{
							FirstName: "test",
							LastName:  "testison",
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					return httptest.NewRequest(http.MethodGet, "http://www.google.com/1234", strings.NewReader(""))
				}(),
			},
			status: http.StatusOK,
			body: model.UserEntity{
				Entity: model.Entity{
					ID:        "1234",
					CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
				},
				User: model.User{
					FirstName: "test",
					LastName:  "testison",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				UserDAO: tt.fields.UserDAO,
			}
			writer := httptest.NewRecorder()
			handler := h.fetchByID()
			handler.ServeHTTP(writer, tt.args.req)

			if writer.Result().StatusCode != tt.status {
				t.Errorf("Handler.FetchByID() = %v, want %v", writer.Result().StatusCode, tt.status)
				return
			}

			var bodyMap map[string]interface{}
			if err := json.NewDecoder(writer.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.FetchByID() = json body decode error %v", err)
				return
			}

			var wantBodyMap map[string]interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.FetchByID() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}

func TestHandler_List(t *testing.T) {
	type fields struct {
		UserDAO DAO
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		status int
		body   interface{}
	}{
		{
			name: "fetched",
			fields: fields{
				UserDAO: &mockUserDAO{
					users: []*model.UserEntity{
						{
							Entity: model.Entity{
								ID:        "1234",
								CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
							},
							User: model.User{
								FirstName: "test",
								LastName:  "testison",
							},
						},
						{
							Entity: model.Entity{
								ID:        "9876",
								CreatedAt: time.Date(2020, time.July, 30, 0, 0, 0, 0, time.UTC),
							},
							User: model.User{
								FirstName: "test-2",
								LastName:  "testison-2",
							},
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					return httptest.NewRequest(http.MethodGet, "http://www.google.com/", strings.NewReader(""))
				}(),
			},
			status: http.StatusOK,
			body: []*model.UserEntity{
				{
					Entity: model.Entity{
						ID:        "1234",
						CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
					},
					User: model.User{
						FirstName: "test",
						LastName:  "testison",
					},
				},
				{
					Entity: model.Entity{
						ID:        "9876",
						CreatedAt: time.Date(2020, time.July, 30, 0, 0, 0, 0, time.UTC),
					},
					User: model.User{
						FirstName: "test-2",
						LastName:  "testison-2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				UserDAO: tt.fields.UserDAO,
			}
			writer := httptest.NewRecorder()
			handler := h.list()
			handler.ServeHTTP(writer, tt.args.req)

			if writer.Result().StatusCode != tt.status {
				t.Errorf("Handler.List() = %v, want %v", writer.Result().StatusCode, tt.status)
				return
			}

			var bodyMap []interface{}
			if err := json.NewDecoder(writer.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.List() = json body decode error %v", err)
				return
			}

			var wantBodyMap []interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.List() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}

func TestHandler_Update(t *testing.T) {
	type fields struct {
		UserDAO DAO
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		status int
		body   interface{}
	}{
		{
			name: "updated",
			fields: fields{
				UserDAO: &mockUserDAO{
					user: &model.UserEntity{
						Entity: model.Entity{
							ID:        "1234",
							CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
						},
						User: model.User{
							FirstName: "test",
							LastName:  "testison",
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					u := model.User{
						FirstName: "test",
						LastName:  "testison",
					}
					enc, _ := json.Marshal(u)
					return httptest.NewRequest(http.MethodPatch, "http://www.google.com/1234", bytes.NewReader(enc))
				}(),
			},
			status: http.StatusOK,
			body: model.UserEntity{
				Entity: model.Entity{
					ID:        "1234",
					CreatedAt: time.Date(2020, time.July, 23, 0, 0, 0, 0, time.UTC),
				},
				User: model.User{
					FirstName: "test",
					LastName:  "testison",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				UserDAO: tt.fields.UserDAO,
			}
			writer := httptest.NewRecorder()
			handler := h.update()
			handler.ServeHTTP(writer, tt.args.req)

			if writer.Result().StatusCode != tt.status {
				t.Errorf("Handler.Update() = %v, want %v", writer.Result().StatusCode, tt.status)
				return
			}

			var bodyMap map[string]interface{}
			if err := json.NewDecoder(writer.Body).Decode(&bodyMap); err != nil {
				t.Errorf("Handler.Update() = json body decode error %v", err)
				return
			}

			var wantBodyMap map[string]interface{}
			if enc, err := json.Marshal(tt.body); err == nil {
				_ = json.Unmarshal(enc, &wantBodyMap)
			}

			if !reflect.DeepEqual(bodyMap, wantBodyMap) {
				t.Errorf("Handler.Update() = %v, want %v", bodyMap, wantBodyMap)
			}
		})
	}
}

func TestHandler_Delete(t *testing.T) {
	type fields struct {
		UserDAO DAO
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		status int
	}{
		{
			name: "deleted",
			fields: fields{
				UserDAO: &mockUserDAO{},
			},
			args: args{
				req: func() *http.Request {
					return httptest.NewRequest(http.MethodDelete, "http://www.google.com/1234", strings.NewReader(""))
				}(),
			},
			status: http.StatusNoContent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				UserDAO: tt.fields.UserDAO,
			}
			writer := httptest.NewRecorder()
			handler := h.delete()
			handler.ServeHTTP(writer, tt.args.req)

			if writer.Result().StatusCode != tt.status {
				t.Errorf("Handler.Delete() = %v, want %v", writer.Result().StatusCode, tt.status)
				return
			}
		})
	}
}

func TestHandler_Add(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "create",
			args: args{
				req: httptest.NewRequest(http.MethodPost, "http://localhost:8080/user", nil),
			},
			want: true,
		},
		{
			name: "fetch",
			args: args{
				req: httptest.NewRequest(http.MethodGet, "http://localhost:8080/users/1234", nil),
			},
			want: true,
		},
		{
			name: "list",
			args: args{
				req: httptest.NewRequest(http.MethodGet, "http://localhost:8080/users", nil),
			},
			want: true,
		},
		{
			name: "update",
			args: args{
				req: httptest.NewRequest(http.MethodPatch, "http://localhost:8080/users/1234", nil),
			},
			want: true,
		},
		{
			name: "delete",
			args: args{
				req: httptest.NewRequest(http.MethodDelete, "http://localhost:8080/users/1234", nil),
			},
			want: true,
		},
		{
			name: "nope",
			args: args{
				req: httptest.NewRequest(http.MethodGet, "http://localhost:8080/user", nil),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{}
			r := mux.NewRouter()

			h.Add(r)

			var match mux.RouteMatch
			if ok := r.Match(tt.args.req, &match); ok != tt.want {
				t.Errorf("Handler.Add() %v", tt.want)
			}
		})
	}
}
