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
