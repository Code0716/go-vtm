package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/gen/api"
	"github.com/Code0716/go-vtm/app/interfaces/handlers"
	"github.com/Code0716/go-vtm/app/util"
)

func Test_adminHandler_RegistAdmin(t *testing.T) {
	t.Parallel()

	type fakes struct {
		fakeIsAdminExist func(ctx context.Context, name, mail string) (bool, error)
		fakeRegistAdmin  func(ctx context.Context, params domain.AdminUser) error
	}
	type args struct {
		domain.RegistAdminJSONRequestBody
	}
	tests := []struct {
		name    string
		fakes   fakes
		args    args
		wantRes wantRes
		wantErr bool
	}{
		{
			"Success",
			fakes{
				fakeIsAdminExist: func(ctx context.Context, name, mail string) (bool, error) {
					return false, nil
				},
				fakeRegistAdmin: func(ctx context.Context, params domain.AdminUser) error {
					return nil
				},
			},
			args{
				domain.RegistAdminJSONRequestBody{
					// admin mail address
					MailAddress: "test@test.co.jp",
					// admin name
					Name: "name",
					// admin password
					Password: "password",
				},
			},
			wantRes{
				code: http.StatusCreated,
				body: domain.CommonSuccessResponse{
					Message: "Success",
				},
			},
			false,
		},
		{
			"no name error",
			fakes{
				fakeIsAdminExist: func(ctx context.Context, name, mail string) (bool, error) {
					return false, nil
				},
				fakeRegistAdmin: func(ctx context.Context, params domain.AdminUser) error {
					return nil
				},
			},
			args{
				domain.RegistAdminJSONRequestBody{
					// admin mail address
					MailAddress: "test@test.co.jp",
					// admin password
					Password: "password",
				},
			},
			wantRes{
				code: http.StatusBadRequest,
				body: domain.ErrorResponse{
					Error: domain.Error{
						Type:    domain.ErrorTypeRegistAdminValidationFailed,
						Status:  http.StatusBadRequest,
						Message: domain.ErrorMessageMap[domain.ErrorTypeRegistAdminValidationFailed],
					},
				},
			},
			false,
		},
		{
			"admin user already exest",
			fakes{
				fakeIsAdminExist: func(ctx context.Context, name, mail string) (bool, error) {
					return true, nil
				},
				fakeRegistAdmin: func(ctx context.Context, params domain.AdminUser) error {
					return nil
				},
			},
			args{
				domain.RegistAdminJSONRequestBody{
					// admin mail address
					MailAddress: "test@test.co.jp",
					Name:        "name",
					// admin password
					Password: "password",
				},
			},
			wantRes{
				code: http.StatusBadRequest,
				body: domain.ErrorResponse{
					Error: domain.Error{
						Type:    domain.ErrorTypeRegistItemAlreadyRegistered,
						Status:  http.StatusBadRequest,
						Message: domain.ErrorMessageMap[domain.ErrorTypeRegistItemAlreadyRegistered],
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := &registryMock{}
			reg.mockAdminRepo.FakeIsAdminExist = tt.fakes.fakeIsAdminExist
			reg.mockAdminRepo.FakeRegistAdmin = tt.fakes.fakeRegistAdmin
			h := handlers.New(reg)

			reqS := fmt.Sprintf(
				`{"name":"%s","mail_address":"%s","password":"%s"}`,
				tt.args.Name,
				tt.args.MailAddress,
				tt.args.Password,
			)

			reqBody := bytes.NewBuffer([]byte(reqS))

			path := fmt.Sprintf("https://test.com/admin/regist")
			req := httptest.NewRequest(http.MethodPost, path, reqBody)
			req.Header.Set("Content-Type", "application/json")

			c, res := newTestEchoContext(t, req)

			si := api.ServerInterfaceWrapper{Handler: h}

			if err := si.RegistAdmin(c); (err != nil) != tt.wantErr {
				t.Errorf("adminHandler.RegistAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if res.Code != tt.wantRes.code {
				t.Errorf("adminHandler.RegistAdmin() http status got = %v want = %v", res.Code, tt.wantRes.code)
			}

			if !testJSON(t, res.Body.Bytes(), tt.wantRes.body) {
				t.Errorf("adminHandler.RegistAdmin() response got = %v, want %v", res.Body, tt.wantRes.body)
			}
		})
	}
}

func Test_adminHandler_GetAdminInfo(t *testing.T) {
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
		wantRes wantRes
		wantErr bool
	}{
		{
			"Success",
			fakes{
				fakeGetAdminByUUID: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
					return &domain.AdminUser{
							Id:          1,
							Name:        "hogehoge",
							AdminId:     "4b4fd313-1a9d-4210-b74d-d214d9fb40ce",
							Password:    "password",
							MailAddress: "test@test.com",
							Status:      "active",
							Authority:   "admin",
							CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
							UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
						},
						nil
				},
			},
			args{
				uuid: "4b4fd313-1a9d-4210-b74d-d214d9fb40ce",
			},
			wantRes{
				code: http.StatusOK,
				body: domain.AdminUserResponse{
					Id:          1,
					Name:        "hogehoge",
					AdminId:     "4b4fd313-1a9d-4210-b74d-d214d9fb40ce",
					MailAddress: "test@test.com",
					Status:      "active",
					Authority:   "admin",
					CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
					UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
				},
			},
			false,
		},
		{
			"faild invalid uuid",
			fakes{
				fakeGetAdminByUUID: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
					return &domain.AdminUser{
							Id:          1,
							Name:        "hogehoge",
							AdminId:     "4b4fd313-1a9d-4210-b74d-d214d9fb40ce",
							Password:    "password",
							MailAddress: "test@test.com",
							Status:      "active",
							Authority:   "admin",
							CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
							UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
						},
						nil
				},
			},
			args{
				uuid: "hogehoge",
			},
			wantRes{
				code: http.StatusBadRequest,
				body: domain.ErrorResponse{
					Error: domain.Error{
						Type:    domain.ErrorTypeUUIDValidationFailed,
						Status:  http.StatusBadRequest,
						Message: domain.ErrorMessageMap[domain.ErrorTypeUUIDValidationFailed],
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := &registryMock{}
			reg.mockAdminRepo.FakeGetAdminByUUID = tt.fakes.fakeGetAdminByUUID
			h := handlers.New(reg)

			path := fmt.Sprintf("https://test.com/admin/:uuid")
			req := httptest.NewRequest(http.MethodGet, path, nil)

			c, res := newTestEchoContext(t, req)
			c.SetPath(path)
			c.SetParamNames("uuid")
			c.SetParamValues(tt.args.uuid)

			si := api.ServerInterfaceWrapper{Handler: h}

			if err := si.GetAdminInfo(c); (err != nil) != tt.wantErr {
				t.Errorf("adminHandler.GetAdminInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if res.Code != tt.wantRes.code {
				t.Errorf("adminHandler.GetAdminInfo() http status got = %v want = %v", res.Code, tt.wantRes.code)
			}

			if !testJSON(t, res.Body.Bytes(), tt.wantRes.body) {
				t.Errorf("adminHandler.GetAdminInfo() response got = %v\n, want %v\n", res.Body, tt.wantRes.body)
			}
		})
	}
}

func Test_adminHandler_GetAdminList(t *testing.T) {
	type fakes struct {
		fakeGetAllAdminUser func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error)
	}

	adminUser1 := &domain.AdminUser{
		Id:          1,
		AdminId:     "hogehoge",
		Name:        "hogehoge",
		MailAddress: "test@test.com",
		Status:      "active",
		Authority:   "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser2 := &domain.AdminUser{
		Id:          2,
		AdminId:     "fugafuga",
		Name:        "fuga",
		MailAddress: "test@test.com",
		Status:      "active",
		Authority:   "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser3 := &domain.AdminUser{
		Id:          3,
		AdminId:     "hoge2",
		Name:        "hoge2",
		MailAddress: "test@test.com",
		Status:      "active",
		Authority:   "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser4 := &domain.AdminUser{
		Id:          4,
		AdminId:     "fuga2",
		Name:        "fuga2",
		MailAddress: "test@test.com",
		Status:      "active",
		Authority:   "admin",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUsers := []*domain.AdminUser{
		adminUser1,
		adminUser2,
		adminUser3,
		adminUser4,
	}

	type args struct {
		Limit  *int64
		Offset *int64
		Status *string
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    args
		wantRes wantRes
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeGetAllAdminUser: func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
					return adminUsers, 4, nil
				},
			},
			args{},
			wantRes{
				code: http.StatusOK,
				body: domain.AdminUsersResponse{
					AdminUsers: adminUsers,
					Total:      4,
				},
			},
			false,
		},
		{
			"success fix limit",
			fakes{
				fakeGetAllAdminUser: func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
					if params.Limit != 50 {
						panic("not fix limit")
					}
					return adminUsers, 4, nil
				},
			},
			args{
				Limit: util.Int64Ptr(1000000),
			},
			wantRes{
				code: http.StatusOK,
				body: domain.AdminUsersResponse{
					AdminUsers: adminUsers,
					Total:      4,
				},
			},
			false,
		},
		{
			"not found",
			fakes{
				fakeGetAllAdminUser: func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
					return nil, 0, domain.NewError(domain.ErrorTypeContentNotFound)
				},
			},
			args{},
			wantRes{
				code: http.StatusBadRequest,
				body: domain.AdminUsersResponse{
					AdminUsers: adminUsers,
					Total:      0,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := &registryMock{}
			reg.mockAdminRepo.FakeGetAllAdminUser = tt.fakes.fakeGetAllAdminUser
			h := handlers.New(reg)

			path := fmt.Sprintf("https://test.com/admin")
			req := httptest.NewRequest(http.MethodGet, path, nil)

			u := req.URL
			q := u.Query()
			if tt.args.Offset != nil {
				q.Add("offset", fmt.Sprint(*tt.args.Offset))
			}
			if tt.args.Limit != nil {
				q.Add("limit", fmt.Sprint(*tt.args.Limit))
			}

			u.RawQuery = q.Encode()

			c, res := newTestEchoContext(t, req)

			si := api.ServerInterfaceWrapper{Handler: h}

			if err := si.GetAdminList(c); (err != nil) != tt.wantErr {
				t.Errorf("adminHandler.GetAdminList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if res.Code != tt.wantRes.code {
				t.Errorf("adminHandler.GetAdminList() http status got = %v want = %v", res.Code, tt.wantRes.code)
			}

			var gotBody domain.AdminUsersResponse
			var wantBody domain.AdminUsersResponse

			err := json.Unmarshal(res.Body.Bytes(), &gotBody)
			if err != nil {
				t.Errorf("adminHandler.GetAdminList()  response Body = %v, want code %v", res.Body, tt.wantRes.body)
			}

			tmpWant, err := json.Marshal(tt.wantRes.body)
			if err != nil {
				t.Error("adminHandler.GetAdminList()  Marshal error")
			}

			err = json.Unmarshal(tmpWant, &wantBody)
			if err != nil {
				t.Errorf("adminHandler.GetAdminList()  response Body = %v, want code %v", gotBody, wantBody)
			}

			for index, admin := range gotBody.AdminUsers {
				if !reflect.DeepEqual(wantBody, gotBody) {
					t.Errorf("adminHandler.GetAdminList()  AdminUsers = %v\n, want %v\n", admin, wantBody.AdminUsers[index])
					return
				}
			}

			if gotBody.Total != wantBody.Total {
				t.Errorf("adminHandler.GetAdminList()  Total = %v, want code %v", wantBody.Total, wantBody.Total)
				return
			}
		})
	}
}

func Test_adminHandler_DeleteAdminInfo(t *testing.T) {
	t.Parallel()

	type fakes struct {
		fakeDeleteAdminUser func(ctx context.Context, uuid string) (*domain.AdminUser, error)
	}

	type args struct {
		uuid string
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    args
		wantRes wantRes
		wantErr bool
	}{
		{
			"Success",
			fakes{fakeDeleteAdminUser: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
				adminUser := &domain.AdminUser{
					Id:          1,
					AdminId:     "be458a2c-b6b7-472b-823b-0a3755a6004b",
					Name:        "hogehoge",
					Password:    "password",
					MailAddress: "test@test.com",
					Status:      "active",
					Authority:   "admin",
					CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
					UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
				}

				return adminUser, nil
			},
			},
			args{uuid: "be458a2c-b6b7-472b-823b-0a3755a6004b"},
			wantRes{
				code: http.StatusOK,
				body: domain.DeleteAdminUserResponse{
					Id:      1,
					Name:    "hogehoge",
					AdminId: "be458a2c-b6b7-472b-823b-0a3755a6004b",
				},
			},
			false,
		},
		{
			"faild uuid validate",
			fakes{fakeDeleteAdminUser: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {
				adminUser := &domain.AdminUser{
					Id:          1,
					AdminId:     "be458a2c-b6b7-472b-823b-0a3755a6004b",
					Name:        "hogehoge",
					Password:    "password",
					MailAddress: "test@test.com",
					Status:      "active",
					Authority:   "admin",
					CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
					UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
				}

				return adminUser, nil
			},
			},
			args{uuid: "hogehoge"},
			wantRes{
				code: http.StatusBadRequest,
				body: domain.ErrorResponse{
					Error: domain.Error{
						Type:    domain.ErrorTypeUUIDValidationFailed,
						Status:  http.StatusBadRequest,
						Message: domain.ErrorMessageMap[domain.ErrorTypeUUIDValidationFailed],
					},
				},
			},
			false,
		},
		{
			"faild not found",
			fakes{fakeDeleteAdminUser: func(ctx context.Context, uuid string) (*domain.AdminUser, error) {

				return nil, domain.NewError(domain.ErrorTypeContentNotFound)
			},
			},
			args{uuid: "be458a2c-b6b7-472b-823b-0a3755a6004b"},
			wantRes{
				code: http.StatusBadRequest,
				body: domain.ErrorResponse{
					Error: domain.Error{
						Type:    domain.ErrorTypeContentNotFound,
						Status:  http.StatusBadRequest,
						Message: domain.ErrorMessageMap[domain.ErrorTypeContentNotFound],
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := &registryMock{}
			reg.mockAdminRepo.FakeDeleteAdminUser = tt.fakes.fakeDeleteAdminUser
			h := handlers.New(reg)

			path := fmt.Sprintf("https://test.com/admin/:uuid")
			req := httptest.NewRequest(http.MethodDelete, path, nil)

			c, res := newTestEchoContext(t, req)
			c.SetPath(path)
			c.SetParamNames("uuid")
			c.SetParamValues(tt.args.uuid)
			si := api.ServerInterfaceWrapper{Handler: h}

			if err := si.DeleteAdminInfo(c); (err != nil) != tt.wantErr {
				t.Errorf("adminHandler.DeleteAdminInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if res.Code != tt.wantRes.code {
				t.Errorf("adminHandler.DeleteAdminInfo() http status got = %v want = %v", res.Code, tt.wantRes.code)
			}
			if !testJSON(t, res.Body.Bytes(), tt.wantRes.body) {
				t.Errorf("adminHandler.DeleteAdminInfo() response got = %v\n, want %v\n", res.Body, tt.wantRes.body)
			}
		})
	}
}
