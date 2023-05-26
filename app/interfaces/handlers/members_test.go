package handlers_test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"reflect"
// 	"testing"

// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/gen/api"
// 	"github.com/Code0716/go-vtm/app/interfaces/handlers"
// 	"github.com/Code0716/go-vtm/app/util"
// )

// func Test_usersHandler_GetUserList(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		Limit  *int64
// 		Offset *int64
// 		Status *string
// 	}
// 	type fakes struct {
// 		fakeGetAll func(ctx context.Context, params domain.Pager) ([]*domain.User, int64, error)
// 	}
// 	// TODO:テストケース増やす
// 	tests := []struct {
// 		name    string
// 		args    args
// 		fakes   fakes
// 		wantRes wantRes
// 		wantErr bool
// 	}{
// 		{
// 			"success",
// 			args{},
// 			fakes{
// 				fakeGetAll: func(ctx context.Context, params domain.Pager) ([]*domain.User, int64, error) {
// 					return []*domain.User{
// 							{
// 								Id:          1,
// 								UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 								Name:        "hogehoge",
// 								PhoneNumber: "09000000000",
// 								Status:      "active",
// 								CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 								UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 							},
// 						},
// 						1,
// 						nil
// 				},
// 			},
// 			wantRes{
// 				code: http.StatusOK,
// 				body: domain.UsersResponse{
// 					Users: []*domain.User{
// 						{
// 							Id:          1,
// 							UserId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
// 							Name:        "hogehoge",
// 							PhoneNumber: "09000000000",
// 							Status:      "active",
// 							CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 							UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 						},
// 					},
// 					Total: 1,
// 				},
// 			},
// 			false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reg := &registryMock{}
// 			reg.mockUserRepo.FakeGetAll = tt.fakes.fakeGetAll
// 			h := handlers.New(reg)

// 			path := fmt.Sprintf("https://test.com/api/v1/admin/users")
// 			req := httptest.NewRequest(http.MethodGet, path, nil)

// 			u := req.URL
// 			q := u.Query()
// 			if tt.args.Offset != nil {
// 				q.Add("offset", fmt.Sprint(*tt.args.Offset))
// 			}
// 			if tt.args.Limit != nil {
// 				q.Add("limit", fmt.Sprint(*tt.args.Limit))
// 			}
// 			u.RawQuery = q.Encode()

// 			c, res := newTestEchoContext(t, req)
// 			si := api.ServerInterfaceWrapper{Handler: h}
// 			if err := si.AdminGetUserList(c); (err != nil) != tt.wantErr {
// 				t.Errorf("userHandler.GetUserList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if res.Code != tt.wantRes.code {
// 				t.Errorf("userHandler.GetUserList()  status code = %v, want code %v", res.Code, tt.wantRes.code)
// 				return
// 			}

// 			// TODO:helperに汎用的に使える関数を用意したい。
// 			var gotBody domain.UsersResponse
// 			var wantBody domain.UsersResponse

// 			err := json.Unmarshal(res.Body.Bytes(), &gotBody)
// 			if err != nil {
// 				t.Errorf("userHandler.GetUserList()  response Body = %v, want code %v", res.Body, tt.wantRes.body)
// 			}

// 			tmpWant, err := json.Marshal(tt.wantRes.body)
// 			if err != nil {
// 				t.Error("userHandler.GetUserList()  Marshal error")
// 			}

// 			err = json.Unmarshal(tmpWant, &wantBody)
// 			if err != nil {
// 				t.Errorf("userHandler.GetUserList()  response Body = %v, want code %v", gotBody, wantBody)
// 			}

// 			for index, user := range gotBody.Users {
// 				if !reflect.DeepEqual(user, wantBody.Users[index]) {
// 					t.Errorf("userHandler.GetUserList()  User = %v, want code %v", user, wantBody.Users[index])
// 					return
// 				}
// 			}

// 			if gotBody.Total != wantBody.Total {
// 				t.Errorf("userHandler.GetUserList()  Total = %v, want code %v", wantBody.Total, wantBody.Total)
// 				return
// 			}

// 		})
// 	}
// }

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
