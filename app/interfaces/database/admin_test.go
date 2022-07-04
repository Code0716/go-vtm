package database_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/database"
	"github.com/Code0716/go-vtm/app/util"
)

func TestAdmin_GetAdminByEmail(t *testing.T) {
	t.Parallel()

	type fakes struct {
		fakeFirst func(value interface{}, where ...interface{}) error
	}

	adminUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		MailAddress: "hoge@test.com",
		Password:    "hoge",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	fakeUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		MailAddress: "hoge@test.com",
		Password:    "hoge",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	type args struct {
		mail string
	}
	tests := []struct {
		name    string
		fakes   fakes
		args    args
		want    *domain.AdminUser
		wantErr bool
	}{
		{
			"Success",
			fakes{
				fakeFirst: func(value interface{}, where ...interface{}) error {
					value = &domain.AdminUser{
						Id:          1,
						AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
						Name:        "hogehoge",
						MailAddress: "hoge@test.com",
						Password:    "hoge",
						Permission:  "admin",
						Status:      "init",
						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
					}

					return nil
				},
			},
			args{
				mail: adminUser.MailAddress,
			},
			adminUser,
			false,
		},
		{
			"faild",
			fakes{
				fakeFirst: func(value interface{}, where ...interface{}) error {
					return domain.WrapInternalError(errors.New("not found"))
				},
			},
			args{
				mail: fakeUser.MailAddress,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeFirst = tt.fakes.fakeFirst
			r := database.NewAdmin(adminRepo)

			_, err := r.GetAdminByEmail(testCtx, tt.args.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("Admin.GetAdminByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Admin.GetAdminByEmail() got = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestAdmin_RegistAdmin(t *testing.T) {
	t.Parallel()
	type fakes struct {
		fakeCreateAdmin func(newAdmin interface{}) error
	}

	adminUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		MailAddress: "test@test.com",
		Password:    "password",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	tests := []struct {
		name    string
		fakes   fakes
		want    error
		args    domain.AdminUser
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeCreateAdmin: func(newAdmin interface{}) error {
					return nil
				},
			},
			nil,
			*adminUser,
			false,
		},
		{
			"failed ",
			fakes{
				fakeCreateAdmin: func(newAdmin interface{}) error {
					return domain.WrapInternalError(errors.New("create faild"))
				},
			},
			domain.WrapInternalError(errors.New("create faild")),
			*adminUser,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeCreateAdmin = tt.fakes.fakeCreateAdmin
			r := database.NewAdmin(adminRepo)
			err := r.RegistAdmin(testCtx, tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("Admin.RegistAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.TypeOf(err) != reflect.TypeOf(tt.want) {
				t.Errorf("Admin.RegistAdmin() error = %v, want %v", reflect.TypeOf(err), reflect.TypeOf(tt.want))
			}

		})
	}
}

func TestAdminRepository_IsAdminExist(t *testing.T) {
	type fakes struct {
		fakeIsAdminExist func(tableName string, query interface{}, args ...interface{}) (bool, error)
	}

	type args struct {
		mail string
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    args
		want    bool
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeIsAdminExist: func(tableName string, query interface{}, args ...interface{}) (bool, error) {

					if args[0] == "" {
						return false, errors.New("faild")
					}
					return false, nil
				},
			},
			args{mail: "test@test.com"},
			false,
			false,
		},
		{
			"faild",
			fakes{
				fakeIsAdminExist: func(tableName string, query interface{}, args ...interface{}) (bool, error) {

					if args[0] == "" {
						return false, errors.New("faild")
					}
					return true, nil
				},
			},
			args{mail: "test@test.com"},
			true,
			false,
		},
		{
			"validate error",
			fakes{
				fakeIsAdminExist: func(tableName string, query interface{}, args ...interface{}) (bool, error) {
					if args[0] == "" {
						return false, errors.New("faild")
					}
					return false, nil
				},
			},
			args{},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeIsAdminExist = tt.fakes.fakeIsAdminExist
			r := database.NewAdmin(adminRepo)
			got, err := r.IsAdminExist(testCtx, tt.args.mail)

			if (err != nil) != tt.wantErr {
				t.Errorf("AdminRepository.IsAdminExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AdminRepository.IsAdminExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
