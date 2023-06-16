package database_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/interfaces/database"
	"github.com/Code0716/go-vtm/app/util"
)

func TestUsersInterFace_CreateUser(t *testing.T) {
	t.Parallel()
	type fakes struct {
		fakeCreate func(c context.Context, user any) db.SQLHandler
		fakeFirst  func(c context.Context, user any) db.SQLHandler
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

	testDB, close, err := getTestDB(t, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	dataBase := db.SQLHandler(*testDB)

	tests := []struct {
		name    string
		fakes   fakes
		args    domain.User
		want    *domain.User
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeCreate: func(c context.Context, user any) db.SQLHandler {
					return dataBase.Create(c, user1)
				},
				fakeFirst: func(c context.Context, user any) db.SQLHandler {
					return dataBase.First(c, &user)

				},
			},
			user1,
			&domain.User{
				ID:               util.LiteralToPtrGenerics[string]("1"),
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
				DeletedAt:        nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := mockUsersRepo{}
			userRepo.FakeCreate = tt.fakes.fakeCreate
			userRepo.FakeFirst = tt.fakes.fakeFirst
			r := database.NewUser(&userRepo)
			got, err := r.CreateUser(testCtx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Users.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Users.GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestUsersRepository_AdminRegistUser(t *testing.T) {
// 	t.Parallel()

// 	type fakes struct {
// 		fakeCreateUser func(m any) error
// 	}

// 	user := &domain.User{
// 		Id:          1,
// 		UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 		Name:        "hogehoge",
// 		PhoneNumber: "09000000000",
// 		Password:    util.LiteralToPtrGenerics("hoge"),
// 		Status:      "init",
// 		HourlyPrice: util.LiteralToPtrGenerics[int64](1500),
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	type args struct {
// 		newUser domain.User
// 	}

// 	tests := []struct {
// 		name    string
// 		fakes   fakes
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			"success",
// 			fakes{
// 				fakeCreateUser: func(m any) error {
// 					return nil
// 				},
// 			},
// 			args{newUser: *user},
// 			false,
// 		},
// 		{
// 			"fail",
// 			fakes{
// 				fakeCreateUser: func(m any) error {
// 					return domain.WrapInternalError(errors.New("create faild"))
// 				},
// 			},
// 			args{},
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// userRepo := mockUsersRepo{}
// 			// userRepo.FakeCreateUser = tt.fakes.fakeCreateUser
// 			// r := database.NewUsers(userRepo)

// 			// if err := r.AdminRegistUser(testCtx, tt.args.newUser); (err != nil) != tt.wantErr {
// 			// 	t.Errorf("UsersRepository.AdminRegistUser() error = %v, wantErr %v", err, tt.wantErr)
// 			// }
// 		})
// 	}
// }

// func TestUsersRepository_IsUserExist(t *testing.T) {
// 	type fakes struct {
// 		fakeIsUserExist func(tableName string, query any, args ...any) (bool, error)
// 	}

// 	type args struct {
// 		name  string
// 		phone string
// 	}

// 	tests := []struct {
// 		name    string
// 		fakes   fakes
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			"success",
// 			fakes{
// 				fakeIsUserExist: func(tableName string, query any, args ...any) (bool, error) {
// 					if args[0] == "" || args[1] == "" {
// 						return false, errors.New("faild")
// 					}
// 					return false, nil
// 				},
// 			},
// 			args{name: "hogehoge", phone: "09000000000"},
// 			false,
// 			false,
// 		},
// 		{
// 			"faild",
// 			fakes{
// 				fakeIsUserExist: func(tableName string, query any, args ...any) (bool, error) {
// 					if args[0] == "" || args[1] == "" {
// 						return false, errors.New("faild")
// 					}
// 					return true, nil
// 				},
// 			},
// 			args{name: "hogehoge", phone: "09000000000"},
// 			true,
// 			false,
// 		},
// 		{
// 			"validate error",
// 			fakes{fakeIsUserExist: func(tableName string, query any, args ...any) (bool, error) {
// 				if args[0] == "" || args[1] == "" {
// 					return false, errors.New("faild")
// 				}
// 				return false, nil
// 			},
// 			},
// 			args{},
// 			false,
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// userRepo := mockUsersRepo{}
// 			// userRepo.FakeIsUserExist = tt.fakes.fakeIsUserExist
// 			// r := database.NewUsers(userRepo)
// 			// got, err := r.IsUserExist(testCtx, tt.args.name, tt.args.phone)

// 			// if (err != nil) != tt.wantErr {
// 			// 	t.Errorf("AdminRepository.IsAdminExist() error = %v, wantErr %v", err, tt.wantErr)
// 			// 	return
// 			// }
// 			// if got != tt.want {
// 			// 	t.Errorf("AdminRepository.IsAdminExist() = %v, want %v", got, tt.want)
// 			// }
// 		})
// 	}
// }
