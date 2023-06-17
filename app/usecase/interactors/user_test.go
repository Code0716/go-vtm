package interactors_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/interactors"
	"github.com/Code0716/go-vtm/app/util"
)

func TestUser_CreateUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type fakes struct {
		fakeCreateUser func(ctx context.Context, user domain.User) (*domain.User, error)
	}

	type args struct {
		user domain.User
	}

	user1 := domain.User{
		UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:             "hoge",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	tests := []struct {
		name    string
		args    args
		fakes   fakes
		want    *domain.User
		wantErr bool
	}{
		{
			"success",
			args{user: user1},
			fakes{
				fakeCreateUser: func(ctx context.Context, user domain.User) (*domain.User, error) {
					return &user, nil
				},
			},
			&user1,
			false,
		},
		{
			"failed - internal server error",
			args{user: user1},
			fakes{
				fakeCreateUser: func(ctx context.Context, user domain.User) (*domain.User, error) {
					return nil, domain.NewError(domain.ErrorTypeInternalError)
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := mockUserRepo{}
			userRepo.FakeCreateUser = tt.fakes.fakeCreateUser

			im := interactors.NewUser(userRepo)
			got, err := im.CreateUser(ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("interactors CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interactors CreateUser() got = %v, want %v", got, tt.want)
			}

		})
	}
}
