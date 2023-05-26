package database_test

// import (
// 	"errors"
// 	"testing"

// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/util"
// )

// func TestUsersRepository_AdminUserGetAll(t *testing.T) {
// 	t.Parallel()
// 	type fakes struct {
// 		fakeAdminUserGetAll func(params domain.Pager) ([]*domain.User, int64, error)
// 	}

// 	user1 := &domain.User{
// 		Id:          1,
// 		UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 		Name:        "hoge",
// 		PhoneNumber: "09000000000",
// 		Status:      "active",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}
// 	user2 := &domain.User{
// 		Id:          2,
// 		UserId:    "fuga",
// 		Name:        "fuga",
// 		PhoneNumber: "09000000000",
// 		Status:      "active",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}
// 	user3 := &domain.User{
// 		Id:          3,
// 		UserId:    "1111",
// 		Name:        "1111",
// 		PhoneNumber: "09000000000",
// 		Status:      "init",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}
// 	user4 := &domain.User{
// 		Id:          4,
// 		UserId:    "4444",
// 		Name:        "4444",
// 		PhoneNumber: "09000000000",
// 		Status:      "init",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}
// 	user5 := &domain.User{
// 		Id:          5,
// 		UserId:    "5555",
// 		Name:        "5555",
// 		PhoneNumber: "09000000000",
// 		Status:      "active",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}
// 	user6 := &domain.User{
// 		Id:          6,
// 		UserId:    "6666",
// 		Name:        "6666",
// 		PhoneNumber: "09000000000",
// 		Status:      "other",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	tests := []struct {
// 		name    string
// 		fakes   fakes
// 		args    domain.Pager
// 		want    []*domain.User
// 		count   int64
// 		wantErr bool
// 	}{
// 		{
// 			"success",
// 			fakes{
// 				fakeAdminUserGetAll: func(params domain.Pager) ([]*domain.User, int64, error) {
// 					if params.Limit != 50 || params.Offset != 0 {
// 						t.Fatal("params not match")
// 					}
// 					return []*domain.User{
// 							user1,
// 							user2,
// 							user3,
// 							user4,
// 							user5,
// 							user6,
// 						},
// 						6, nil
// 				},
// 			},
// 			domain.Pager{
// 				Limit:  50,
// 				Offset: 0,
// 				Status: "",
// 			},
// 			[]*domain.User{
// 				user1,
// 				user2,
// 				user3,
// 				user4,
// 				user5,
// 				user6,
// 			},
// 			6,
// 			false,
// 		},
// 		{
// 			"offset 3",
// 			fakes{
// 				fakeAdminUserGetAll: func(params domain.Pager) ([]*domain.User, int64, error) {
// 					if params.Limit != 50 || params.Offset != 3 {
// 						t.Fatal("params not match")
// 					}
// 					return []*domain.User{
// 							user4,
// 							user5,
// 							user6,
// 						},
// 						3, nil
// 				},
// 			},
// 			domain.Pager{
// 				Limit:  50,
// 				Offset: 3,
// 				Status: "",
// 			},
// 			[]*domain.User{
// 				user4,
// 				user5,
// 				user6,
// 			},
// 			3,
// 			false,
// 		},
// 		{
// 			"MmberStatus init",
// 			fakes{
// 				fakeAdminUserGetAll: func(params domain.Pager) ([]*domain.User, int64, error) {
// 					if params.Limit != 50 || params.Offset != 0 || params.Status != "init" {
// 						t.Fatal("params not match")
// 					}
// 					return []*domain.User{
// 							user3,
// 							user4,
// 						},
// 						2, nil
// 				},
// 			},
// 			domain.Pager{
// 				Limit:  50,
// 				Offset: 0,
// 				Status: "init",
// 			},
// 			[]*domain.User{
// 				user3,
// 				user4,
// 			},
// 			2,
// 			false,
// 		},
// 		{
// 			"MmberStatus other",
// 			fakes{
// 				fakeAdminUserGetAll: func(params domain.Pager) ([]*domain.User, int64, error) {
// 					if params.Limit != 100 || params.Offset != 0 || params.Status != "other" {
// 						t.Fatal("params not match")
// 					}
// 					return []*domain.User{
// 							user6,
// 						},
// 						1, nil
// 				},
// 			},
// 			domain.Pager{
// 				Limit:  100,
// 				Offset: 0,
// 				Status: "other",
// 			},
// 			[]*domain.User{
// 				user6,
// 			},
// 			1,
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// userRepo := mockUsersRepo{}
// 			// userRepo.FakeAdminUserGetAll = tt.fakes.fakeAdminUserGetAll
// 			// r := database.NewUsers(userRepo)
// 			// got, gotCount, err := r.AdminUserGetAll(testCtx, tt.args)
// 			// if (err != nil) != tt.wantErr {
// 			// 	t.Errorf("Users.GetAll() error = %v, wantErr %v", err, tt.wantErr)
// 			// 	return
// 			// }
// 			// if !reflect.DeepEqual(got, tt.want) {
// 			// 	t.Errorf("Users.GetAll() got = %v, want %v", got, tt.want)
// 			// }
// 			// if gotCount != tt.count {
// 			// 	t.Errorf("Users.GetAll() gotCount = %v, want %v", gotCount, tt.count)
// 			// }
// 		})
// 	}
// }

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
