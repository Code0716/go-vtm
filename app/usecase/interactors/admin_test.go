package interactors_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/interactors"
	"github.com/Code0716/go-vtm/app/util"
)

func TestAdmin_RegistAdmin(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type fakes struct {
		fakeRegistAdmin func(ctx context.Context, params domain.AdminUser) error
	}

	type args struct {
		params domain.RegistAdminJSONRequestBody
	}
	tests := []struct {
		name    string
		fakes   fakes
		args    args
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeRegistAdmin: func(ctx context.Context, params domain.AdminUser) error {
					return nil
				},
			},
			args{
				params: domain.RegistAdminJSONRequestBody{
					Name:        "hogehoge",
					MailAddress: "test@test.com",
					Password:    "password",
				},
			},
			false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeRegistAdmin = tt.fakes.fakeRegistAdmin
			ia := interactors.NewAdmin(adminRepo)

			if err := ia.RegistAdmin(ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Admin.RegistAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdmin_IsAdminExist(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type fakes struct {
		fakeIsAdminExist func(ctx context.Context, mail string) (bool, error)
	}

	type args struct {
		name string
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
			"Success",
			fakes{
				fakeIsAdminExist: func(ctx context.Context, mail string) (bool, error) {
					return false, nil
				},
			},
			args{mail: "mail"},
			false,
			false,
		},
		{
			"failed",
			fakes{
				fakeIsAdminExist: func(ctx context.Context, mail string) (bool, error) {
					return true, nil
				},
			},
			args{name: "name", mail: "mail"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeIsAdminExist = tt.fakes.fakeIsAdminExist
			ia := interactors.NewAdmin(adminRepo)

			got, err := ia.IsAdminExist(ctx, tt.args.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("Admin.IsAdminExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Admin.IsAdminExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminInteractor_GetAdminByUUID(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	adminUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "hogehoge",
		Name:        "hogehoge",
		Password:    "password",
		MailAddress: "test@test.com",
		Status:      "active",
		Permission:  "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	type fakes struct {
		fakeGetAdminByUUID func(ctx context.Context, uuid string) (*domain.AdminUser, error)
	}

	type args struct {
		uuid string
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
				fakeGetAdminByUUID: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
					return adminUser, nil
				},
			},
			args{uuid: "hogehoge"},
			adminUser,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeGetAdminByUUID = tt.fakes.fakeGetAdminByUUID
			ia := interactors.NewAdmin(adminRepo)

			got, err := ia.GetAdminByUUID(ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminInteractor.GetAdminByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminInteractor.GetAdminByUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdminInteractor_GetAdminList(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	adminUser1 := &domain.AdminUser{
		Id:          1,
		AdminId:     "hogehoge",
		Name:        "hogehoge",
		Password:    "password",
		MailAddress: "test@test.com",
		Status:      "active",
		Permission:  "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser2 := &domain.AdminUser{
		Id:          2,
		AdminId:     "fugafuga",
		Name:        "fuga",
		Password:    "password",
		MailAddress: "test@test.com",
		Status:      "active",
		Permission:  "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser3 := &domain.AdminUser{
		Id:          3,
		AdminId:     "hoge2",
		Name:        "hoge2",
		Password:    "password",
		MailAddress: "test@test.com",
		Status:      "active",
		Permission:  "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser4 := &domain.AdminUser{
		Id:          4,
		AdminId:     "fuga2",
		Name:        "fuga2",
		Password:    "password",
		MailAddress: "test@test.com",
		Status:      "active",
		Permission:  "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUsers := []*domain.AdminUser{
		adminUser1,
		adminUser2,
		adminUser3,
		adminUser4,
	}

	type fakes struct {
		fakeGetAllAdminUser func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error)
	}

	type args struct {
		params domain.Pager
	}
	tests := []struct {
		name    string
		args    args
		fakes   fakes
		want    []*domain.AdminUser
		count   int64
		wantErr bool
	}{
		{
			"Success",
			args{params: domain.Pager{}},
			fakes{
				fakeGetAllAdminUser: func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
					return adminUsers, 4, nil
				},
			},
			adminUsers,
			4,
			false,
		},
		{
			"Success Limit and offset",
			args{params: domain.Pager{Limit: 1, Offset: 2}},
			fakes{
				fakeGetAllAdminUser: func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
					return []*domain.AdminUser{
						adminUser3,
						adminUser4,
					}, 2, nil
				},
			},
			[]*domain.AdminUser{
				adminUser3,
				adminUser4,
			},
			2,
			false,
		},
		{
			"faild",
			args{params: domain.Pager{}},
			fakes{
				fakeGetAllAdminUser: func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
					return nil, 0, errors.New("error message")
				},
			},
			nil,
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}
			adminRepo.FakeGetAllAdminUser = tt.fakes.fakeGetAllAdminUser
			ia := interactors.NewAdmin(adminRepo)

			got, count, err := ia.GetAdminList(ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminInteractor.GetAdminList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminInteractor.GetAdminList() got = %v, want %v", got, tt.want)
			}
			if count != tt.count || int64(len(got)) != count {
				t.Errorf("AdminInteractor.GetAdminList() count = %v, want %v", count, tt.count)
			}
		},
		)
	}
}

func TestAdminInteractor_DeleteAdmin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	adminUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "be458a2c-b6b7-472b-823b-0a3755a6004b",
		Name:        "hogehoge",
		Password:    "password",
		MailAddress: "test@test.com",
		Status:      "active",
		Permission:  "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	type fakes struct {
		fakeDeleteAdmin func(ctx context.Context, uuid string) (*domain.AdminUser, error)
	}

	type args struct {
		uuid string
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
				fakeDeleteAdmin: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
					au := &domain.AdminUser{
						Id:          1,
						AdminId:     "be458a2c-b6b7-472b-823b-0a3755a6004b",
						Name:        "hogehoge",
						Password:    "password",
						MailAddress: "test@test.com",
						Status:      "active",
						Permission:  "admin",
						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
					}
					return au, nil
				},
			},
			args{
				uuid: "be458a2c-b6b7-472b-823b-0a3755a6004b",
			},
			adminUser,
			false,
		},
		{
			"faild",
			fakes{
				fakeDeleteAdmin: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
					return nil, domain.NewError(domain.ErrorTypeContentNotFound)
				},
			},
			args{
				uuid: "hogehoge",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adminRepo := mockAdminRepo{}

			adminRepo.FakeDeleteAdmin = tt.fakes.fakeDeleteAdmin
			ia := interactors.NewAdmin(adminRepo)

			got, err := ia.DeleteAdmin(ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminInteractor.DeleteAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdminInteractor.DeleteAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
