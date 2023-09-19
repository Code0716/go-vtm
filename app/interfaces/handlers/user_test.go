package handlers_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/handlers"
	"github.com/Code0716/go-vtm/app/util"
	"github.com/Code0716/go-vtm/graph/model"
)

func Test_handler_CreateUser(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	successResponse := model.User{
		ID:               "1",
		UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:             "hoge",
		MailAddress:      util.LiteralToPtrGenerics("test@test.com"),
		PhoneNumber:      util.LiteralToPtrGenerics("09000000000"),
		Status:           model.UserStatusActive,
		Role:             model.UserRoleCommon,
		EmploymentStatus: model.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics(1200),
		DepartmentID:     nil,
		CreatedAt:        "2023-09-14 00:08:54",
		UpdatedAt:        "2023-10-19 00:09:32",
		DeletedAt:        nil,
	}

	successResponse2 := model.User{
		ID:               "1",
		UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:             "hoge",
		MailAddress:      util.LiteralToPtrGenerics("test@test.com"),
		PhoneNumber:      util.LiteralToPtrGenerics("09000000000"),
		Status:           model.UserStatusActive,
		Role:             model.UserRoleCommon,
		EmploymentStatus: model.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics(1200),
		DepartmentID:     nil,
		CreatedAt:        "2023-09-14 00:08:54",
		UpdatedAt:        "2023-10-19 00:09:32",
		DeletedAt:        util.LiteralToPtrGenerics("2023-10-19 00:09:32"),
	}

	type args struct {
		user model.CreateUserInput
	}

	type fakes struct {
		fakeCreateUser func(ctx context.Context, user domain.User) (*domain.User, error)
	}

	tests := []struct {
		name    string
		args    args
		fakes   fakes
		wantRes *model.User
		wantErr bool
	}{
		{
			"success",
			args{model.CreateUserInput{
				Name:         "hoge",
				MailAddress:  util.LiteralToPtrGenerics("test@test.com"),
				PhoneNumber:  util.LiteralToPtrGenerics("09000000000"),
				UnitPrice:    util.LiteralToPtrGenerics(1200),
				DepartmentID: nil,
			}},
			fakes{
				fakeCreateUser: func(ctx context.Context, user domain.User) (*domain.User, error) {
					return &domain.User{
						ID:               util.LiteralToPtrGenerics("1"),
						UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
						Name:             "hoge",
						MailAddress:      util.LiteralToPtrGenerics("test@test.com"),
						PhoneNumber:      util.LiteralToPtrGenerics("09000000000"),
						Status:           domain.UserStatusActive,
						Role:             domain.UserRoleCommon,
						EmploymentStatus: domain.EmploymentStatusHourly,
						UnitPrice:        util.LiteralToPtrGenerics(1200),
						DepartmentID:     nil,
						CreatedAt:        util.TimeFromStr("2023-09-14 00:08:54"),
						UpdatedAt:        util.TimeFromStr("2023-10-19 00:09:32"),
					}, nil
				},
			},
			&successResponse,
			false,
		},
		{
			"success - deleted user",
			args{model.CreateUserInput{
				Name:         "hoge",
				MailAddress:  util.LiteralToPtrGenerics("test@test.com"),
				PhoneNumber:  util.LiteralToPtrGenerics("09000000000"),
				UnitPrice:    util.LiteralToPtrGenerics(1200),
				DepartmentID: nil,
			}},
			fakes{
				fakeCreateUser: func(ctx context.Context, user domain.User) (*domain.User, error) {
					return &domain.User{
						ID:               util.LiteralToPtrGenerics("1"),
						UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
						Name:             "hoge",
						MailAddress:      util.LiteralToPtrGenerics("test@test.com"),
						PhoneNumber:      util.LiteralToPtrGenerics("09000000000"),
						Status:           domain.UserStatusActive,
						Role:             domain.UserRoleCommon,
						EmploymentStatus: domain.EmploymentStatusHourly,
						UnitPrice:        util.LiteralToPtrGenerics(1200),
						DepartmentID:     nil,
						CreatedAt:        util.TimeFromStr("2023-09-14 00:08:54"),
						UpdatedAt:        util.TimeFromStr("2023-10-19 00:09:32"),
						DeletedAt:        util.LiteralToPtrGenerics[time.Time](util.TimeFromStr("2023-10-19 00:09:32")),
					}, nil
				},
			},
			&successResponse2,
			false,
		},

		{
			"failed - internal server error",
			args{model.CreateUserInput{
				Name:         "hoge",
				MailAddress:  util.LiteralToPtrGenerics("test@test.com"),
				PhoneNumber:  util.LiteralToPtrGenerics("09000000000"),
				UnitPrice:    util.LiteralToPtrGenerics(1200),
				DepartmentID: nil,
			}},
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
			reg := &registryMock{}
			reg.mockUserRepo.FakeCreateUser = tt.fakes.fakeCreateUser
			h := handlers.New(reg)

			got, err := h.CreateUser(ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("interactors CreateUser() error = %v,\n wantErr %v\n", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantRes) {
				t.Errorf("interactors CreateUser() got = %v,\n want %v\n", got, tt.wantRes)
			}

		})
	}
}

// func Test_usersHandler_GetUser(t *testing.T) {
// 	t.Parallel()

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
// 		wantRes wantRes
// 		wantErr bool
// 	}{
// 		{
// 			"success",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return &domain.User{
// 							Id:          1,
// 							UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 							Name:        "hogehoge",
// 							PhoneNumber: "09000000000",
// 							Status:      "active",
// 							CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 							UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 						},
// 						nil
// 				},
// 			},
// 			args{uuid: "ab6ddfb6-ccec-45c2-9269-976c401612da"},
// 			wantRes{
// 				code: http.StatusOK,
// 				body: domain.User{
// 					Id:          1,
// 					UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 					Name:        "hogehoge",
// 					PhoneNumber: "09000000000",
// 					Status:      "active",
// 					CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 					UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 				}},
// 			false,
// 		},
// 		{
// 			"feild invalid uuid",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return &domain.User{
// 							Id:          1,
// 							UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 							Name:        "hogehoge",
// 							PhoneNumber: "09000000000",
// 							Status:      "active",
// 							CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 							UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 						},
// 						nil
// 				},
// 			},
// 			args{uuid: "hogehoge"},
// 			wantRes{
// 				code: http.StatusBadRequest,
// 				body: domain.ErrorResponse{
// 					Error: domain.Error{
// 						Type:       domain.ErrorTypeUUIDValidationFailed,
// 						Status:     http.StatusBadRequest,
// 						Message:    "invalid uuid",
// 						InnerError: nil,
// 					},
// 				},
// 			},
// 			false,
// 		},
// 		{"feild none uuid",
// 			fakes{
// 				fakeGetUserByUUID: func(ctx context.Context, uuid string) (*domain.User, error) {
// 					return &domain.User{
// 							Id:          1,
// 							UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 							Name:        "hogehoge",
// 							PhoneNumber: "09000000000",
// 							Status:      "active",
// 							CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 							UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 						},
// 						nil
// 				},
// 			},
// 			args{},
// 			wantRes{
// 				code: http.StatusOK,
// 				body: 0,
// 			},
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reg := &registryMock{}
// 			reg.mockUserRepo.FakeGetUserByUUID = tt.fakes.fakeGetUserByUUID
// 			h := handlers.New(reg)

// 			path := fmt.Sprintf("https://test.com/api/v1/users/:uuid")
// 			req := httptest.NewRequest(http.MethodGet, path, nil)
// 			c, res := newTestEchoContext(t, req)
// 			c.SetPath(path)
// 			c.SetParamNames("uuid")
// 			c.SetParamValues(tt.args.uuid)
// 			si := api.ServerInterfaceWrapper{Handler: h}

// 			if err := si.GetUser(c); (err != nil) != tt.wantErr {
// 				t.Errorf("usersHandler.GetUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			if res.Code != tt.wantRes.code {
// 				t.Errorf("usersHandler.GetUser()status code = %v, want code %v", res.Code, tt.wantRes.code)
// 				return
// 			}
// 			if len(res.Body.Bytes()) != 0 {
// 				testJSON(t, res.Body.Bytes(), tt.wantRes.body)
// 			}
// 		})
// 	}
// }
