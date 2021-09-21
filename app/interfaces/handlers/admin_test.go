package handlers_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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

			if err := si.Handler.RegistAdmin(c); (err != nil) != tt.wantErr {
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

			reqS := fmt.Sprintf(
				`{"uuid":"%s"}`,
				tt.args.uuid,
			)

			reqBody := bytes.NewBuffer([]byte(reqS))

			path := fmt.Sprintf("https://test.com/admin/%s", tt.args.uuid)
			req := httptest.NewRequest(http.MethodPost, path, reqBody)
			req.Header.Set("Content-Type", "application/json")

			c, res := newTestEchoContext(t, req)

			si := api.ServerInterfaceWrapper{Handler: h}

			if err := si.Handler.GetAdminInfo(c, tt.args.uuid); (err != nil) != tt.wantErr {
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
