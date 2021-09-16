package interactor_test

import (
	"context"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interactor"
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
			ia := interactor.NewAdmin(adminRepo)

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
		fakeIsAdminExist func(ctx context.Context, name, mail string) (bool, error)
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
				fakeIsAdminExist: func(ctx context.Context, name, mail string) (bool, error) {
					return false, nil
				},
			},
			args{name: "name", mail: "mail"},
			false,
			false,
		},
		{
			"failed",
			fakes{
				fakeIsAdminExist: func(ctx context.Context, name, mail string) (bool, error) {
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
			ia := interactor.NewAdmin(adminRepo)

			got, err := ia.IsAdminExist(ctx, tt.args.name, tt.args.mail)
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
