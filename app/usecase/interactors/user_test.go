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

// func TestUsersInteractor_GetUserByUUID(t *testing.T) {
// 	t.Parallel()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	expectUser := domain.User{
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		DeletedAt:   nil,
// 		HourlyPrice: util.LiteralToPtrGenerics[int64](1000),
// 		Id:          1,
// 		UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 		Name:        "hoge",
// 		PhoneNumber: "09000000000",
// 		Status:      "active",
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	type fakes struct {
// 		fakeGetUserByUUID func(ctx context.Context, uuid string) (*domain.User, error)
// 	}

// 	type args struct {
// 		uuid string
// 	}
// 	tests := []struct {
// 		name    string
// 		fakes   fakes
// 		args    args
// 		want    *domain.User
// 		wantErr bool
// 	}{
// 		{
// 			"success",
// 			fakes{fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 				return &domain.User{
// 					CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 					DeletedAt:   nil,
// 					HourlyPrice: util.LiteralToPtrGenerics[int64](1000),
// 					Id:          1,
// 					UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 					Name:        "hoge",
// 					PhoneNumber: "09000000000",
// 					Status:      "active",
// 					UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 				}, nil
// 			},
// 			},
// 			args{uuid: "873a2824-8006-4e67-aed7-ec427df5fce8"},
// 			&expectUser,
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			userRepo := mockUserRepo{}
// 			userRepo.FakeGetUserByUUID = tt.fakes.fakeGetUserByUUID
// 			im := interactors.NewUsers(userRepo)

// 			got, err := im.GetUserByUUID(ctx, tt.args.uuid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UsersInteractor.GetUserByUUID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UsersInteractor.GetUserByUUID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUsersInteractor_UpdateUser(t *testing.T) {
// 	t.Parallel()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	type fakes struct {
// 		fakeGetUserByUUID func(ctx context.Context, uuid string) (*domain.User, error)
// 		fakeUpdateUser    func(ctx context.Context, oldUser domain.User) (*domain.User, error)
// 	}

// 	type args struct {
// 		params domain.UpdateUserJSONBody
// 		uuid   string
// 	}

// 	tests := []struct {
// 		name    string
// 		fakes   fakes
// 		args    args
// 		want    *domain.User
// 		wantErr bool
// 	}{
// 		{
// 			"success change all",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return &domain.User{
// 						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 						DeletedAt:   nil,
// 						HourlyPrice: util.LiteralToPtrGenerics[int64](1000),
// 						Id:          1,
// 						UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 						Name:        "hoge",
// 						PhoneNumber: "09000000000",
// 						Status:      "active",
// 						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 					}, nil
// 				},
// 				fakeUpdateUser: func(ctx context.Context, oldUser domain.User) (*domain.User, error) {
// 					oldUser.UpdatedAt = util.TimeFromStr("2021-10-19 15:09:32")
// 					return &oldUser, nil
// 				},
// 			},
// 			args{
// 				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
// 				params: domain.UpdateUserJSONBody{
// 					HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
// 					Name:        "fuga",
// 					PhoneNumber: "08000000000",
// 					Status:      "other",
// 				},
// 			},
// 			&domain.User{
// 				CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 				DeletedAt:   nil,
// 				HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
// 				Id:          1,
// 				UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 				Name:        "fuga",
// 				PhoneNumber: "08000000000",
// 				Status:      "other",
// 				UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 			},
// 			false,
// 		},
// 		{
// 			"success change PhoneNumber and HourlyPrice",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return &domain.User{
// 						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 						DeletedAt:   nil,
// 						HourlyPrice: util.LiteralToPtrGenerics[int64](1000),
// 						Id:          1,
// 						UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 						Name:        "hoge",
// 						PhoneNumber: "09000000000",
// 						Status:      "active",
// 						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 					}, nil
// 				},
// 				fakeUpdateUser: func(ctx context.Context, oldUser domain.User) (*domain.User, error) {
// 					oldUser.UpdatedAt = util.TimeFromStr("2021-10-19 15:09:32")
// 					return &oldUser, nil
// 				},
// 			},
// 			args{
// 				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
// 				params: domain.UpdateUserJSONBody{
// 					HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
// 					PhoneNumber: "08000000000",
// 				},
// 			},
// 			&domain.User{
// 				CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 				DeletedAt:   nil,
// 				HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
// 				Id:          1,
// 				UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 				Name:        "hoge",
// 				PhoneNumber: "08000000000",
// 				Status:      "active",
// 				UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 			},
// 			false,
// 		},
// 		{
// 			"faild not found user by uuid",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return nil, domain.NewError(domain.ErrorTypeContentNotFound)
// 				},
// 				fakeUpdateUser: func(ctx context.Context, oldUser domain.User) (*domain.User, error) {
// 					oldUser.UpdatedAt = util.TimeFromStr("2021-10-19 15:09:32")
// 					return &oldUser, nil
// 				},
// 			},
// 			args{
// 				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
// 				params: domain.UpdateUserJSONBody{
// 					HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
// 					PhoneNumber: "08000000000",
// 				},
// 			},
// 			nil,
// 			true,
// 		},
// 		{
// 			"faild internal server error",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return &domain.User{
// 						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 						DeletedAt:   nil,
// 						HourlyPrice: util.LiteralToPtrGenerics[int64](1000),
// 						Id:          1,
// 						UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 						Name:        "hoge",
// 						PhoneNumber: "09000000000",
// 						Status:      "active",
// 						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 					}, nil
// 				},
// 				fakeUpdateUser: func(ctx context.Context, oldUser domain.User) (*domain.User, error) {

// 					return nil, domain.NewError(domain.ErrorTypeInternalError)
// 				},
// 			},
// 			args{
// 				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
// 				params: domain.UpdateUserJSONBody{
// 					HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
// 					PhoneNumber: "08000000000",
// 				},
// 			},
// 			nil,
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			userRepo := mockUserRepo{}
// 			userRepo.FakeGetUserByUUID = tt.fakes.fakeGetUserByUUID
// 			userRepo.FakeUpdateUser = tt.fakes.fakeUpdateUser
// 			im := interactors.NewUsers(userRepo)

// 			got, err := im.UpdateUser(ctx, tt.args.params, tt.args.uuid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UsersInteractor.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UsersInteractor.UpdateUser() = %v, want %v", got, tt.want)
// 			}

// 		})
// 	}
// }
