package dal

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/g8rswimmer/go-data-access-example/pkg/model"
)

func TestUser_Create(t *testing.T) {
	type fields struct {
		DB           *sql.DB
		GenerateUUID GenerateUUID
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserEntity
		wantErr bool
	}{
		{
			name: "Create",
			fields: fields{
				DB: setupDB([]string{UserTable}),
				GenerateUUID: func() string {
					return "1234"
				},
			},
			args: args{
				user: &model.User{
					FirstName: "test",
					LastName:  "one",
				},
			},
			want: &model.UserEntity{
				Entity: model.Entity{
					ID: "1234",
				},
				User: model.User{
					FirstName: "test",
					LastName:  "one",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				DB:           tt.fields.DB,
				GenerateUUID: tt.fields.GenerateUUID,
			}
			defer u.DB.Close()

			got, err := u.Create(context.Background(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.CreatedAt = time.Time{}
			got.UpdatedAt = time.Time{}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_FetchByID(t *testing.T) {
	type fields struct {
		DB           *sql.DB
		GenerateUUID GenerateUUID
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserEntity
		wantErr bool
	}{
		{
			name: "Fetch by id",
			fields: fields{
				DB: setupDB([]string{
					UserTable,
					`INSERT INTO user (id, first_name, last_name) VALUES ('123456789012345678901234567890123456', 'test', 'one')`,
				}),
			},
			args: args{
				id: "123456789012345678901234567890123456",
			},
			want: &model.UserEntity{
				Entity: model.Entity{
					ID: "123456789012345678901234567890123456",
				},
				User: model.User{
					FirstName: "test",
					LastName:  "one",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				DB:           tt.fields.DB,
				GenerateUUID: tt.fields.GenerateUUID,
			}
			defer u.DB.Close()

			got, err := u.FetchByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.FetchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.CreatedAt = time.Time{}
			got.UpdatedAt = time.Time{}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.FetchByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_FetchAll(t *testing.T) {
	type fields struct {
		DB           *sql.DB
		GenerateUUID GenerateUUID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.UserEntity
		wantErr bool
	}{
		{
			name: "Fetch all",
			fields: fields{
				DB: setupDB([]string{
					UserTable,
					`INSERT INTO user (id, first_name, last_name) VALUES ('123456789012345678901234567890123456', 'test', 'one')`,
					`INSERT INTO user (id, first_name, last_name) VALUES ('123456789012345678901234567890123457', 'test', 'two')`,
				}),
			},
			want: []*model.UserEntity{
				{
					Entity: model.Entity{
						ID: "123456789012345678901234567890123456",
					},
					User: model.User{
						FirstName: "test",
						LastName:  "one",
					},
				},
				{
					Entity: model.Entity{
						ID: "123456789012345678901234567890123457",
					},
					User: model.User{
						FirstName: "test",
						LastName:  "two",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				DB:           tt.fields.DB,
				GenerateUUID: tt.fields.GenerateUUID,
			}
			defer u.DB.Close()

			got, err := u.FetchAll(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("User.FetchAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range got {
				got[i].CreatedAt = time.Time{}
				got[i].UpdatedAt = time.Time{}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.FetchAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	type fields struct {
		DB           *sql.DB
		GenerateUUID GenerateUUID
	}
	type args struct {
		id   string
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserEntity
		wantErr bool
	}{
		{
			name: "Update",
			fields: fields{
				DB: setupDB([]string{
					UserTable,
					`INSERT INTO user (id, first_name, last_name) VALUES ('123456789012345678901234567890123456', 'test', 'one')`,
				}),
			},
			args: args{
				id: "123456789012345678901234567890123456",
				user: &model.User{
					FirstName: "testy",
					LastName:  "two",
				},
			},
			want: &model.UserEntity{
				Entity: model.Entity{
					ID: "123456789012345678901234567890123456",
				},
				User: model.User{
					FirstName: "testy",
					LastName:  "two",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				DB:           tt.fields.DB,
				GenerateUUID: tt.fields.GenerateUUID,
			}
			defer u.DB.Close()

			got, err := u.Update(context.Background(), tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got.CreatedAt = time.Time{}
			got.UpdatedAt = time.Time{}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Update() = %v, want %v", got, tt.want)
			}

			got2, err := u.FetchByID(context.Background(), tt.args.id)
			if err != nil {
				t.Errorf("User.Update() = %v", err)
				return
			}

			got2.CreatedAt = time.Time{}
			got2.UpdatedAt = time.Time{}

			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("User.Update() = %v, want %v", got2, tt.want)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	type fields struct {
		DB           *sql.DB
		GenerateUUID GenerateUUID
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete",
			fields: fields{
				DB: setupDB([]string{
					UserTable,
					`INSERT INTO user (id, first_name, last_name) VALUES ('123456789012345678901234567890123456', 'test', 'one')`,
				}),
			},
			args: args{
				id: "123456789012345678901234567890123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				DB:           tt.fields.DB,
				GenerateUUID: tt.fields.GenerateUUID,
			}
			defer u.DB.Close()

			if err := u.Delete(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("User.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
