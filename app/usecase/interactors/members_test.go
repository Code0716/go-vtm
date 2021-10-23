package interactors_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/interactors"
	"github.com/Code0716/go-vtm/app/util"
)

func TestMembers_GetAll(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	successData := []*domain.Member{
		{
			Id:        1,
			MemberId:  "873a2824-8006-4e67-aed7-ec427df5fce8",
			Name:      "hoge",
			Status:    "active",
			CreatedAt: util.TimeFromStr("2021-09-14 15:08:54"),
			UpdatedAt: util.TimeFromStr("2021-10-19 15:09:32"),
		},
		{
			Id:        2,
			MemberId:  "fuga",
			Name:      "fuga",
			Status:    "active",
			CreatedAt: util.TimeFromStr("2021-09-14 15:08:54"),
			UpdatedAt: util.TimeFromStr("2021-10-19 15:09:32"),
		},
	}

	type fakes struct {
		fakeGetAll func(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error)
	}

	type args struct {
		params domain.Pager
	}
	tests := []struct {
		name    string
		args    args
		fakes   fakes
		want    []*domain.Member
		total   int64
		wantErr bool
	}{
		{
			"success",
			args{params: domain.Pager{}},
			fakes{
				fakeGetAll: func(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error) {
					return successData, 0, nil
				},
			},
			successData,
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memberRepo := mockMemberRepo{}
			memberRepo.FakeGetAll = tt.fakes.fakeGetAll
			im := interactors.NewMembers(memberRepo)
			got, count, err := im.MemberGetAll(ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Members.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Members.GetAll() got = %v, want %v", got, tt.want)
			}
			if count != tt.total {
				t.Errorf("Members.GetAll() got1 = %v, want %v", count, tt.total)
			}
		})
	}
}

func TestMembersInteractor_GetMemberByUUID(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	expectMember := domain.Member{
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		DeletedAt:   nil,
		HourlyPrice: util.Int64Ptr(1000),
		Id:          1,
		MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hoge",
		PhoneNumber: "09000000000",
		Status:      "active",
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	type fakes struct {
		fakeGetMemberByUUID func(ctx context.Context, uuid string) (*domain.Member, error)
	}

	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fakes   fakes
		args    args
		want    *domain.Member
		wantErr bool
	}{
		{
			"success",
			fakes{fakeGetMemberByUUID: func(ctx context.Context, uuid string) (*domain.Member, error) {
				return &domain.Member{
					CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
					DeletedAt:   nil,
					HourlyPrice: util.Int64Ptr(1000),
					Id:          1,
					MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
					Name:        "hoge",
					PhoneNumber: "09000000000",
					Status:      "active",
					UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
				}, nil
			},
			},
			args{uuid: "873a2824-8006-4e67-aed7-ec427df5fce8"},
			&expectMember,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memberRepo := mockMemberRepo{}
			memberRepo.FakeGetMemberByUUID = tt.fakes.fakeGetMemberByUUID
			im := interactors.NewMembers(memberRepo)

			got, err := im.GetMemberByUUID(ctx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersInteractor.GetMemberByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MembersInteractor.GetMemberByUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembersInteractor_UpdateMember(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type fakes struct {
		fakeGetMemberByUUID func(ctx context.Context, uuid string) (*domain.Member, error)
		fakeUpdateMember    func(ctx context.Context, oldMember domain.Member) (*domain.Member, error)
	}

	type args struct {
		params domain.UpdateMemberJSONBody
		uuid   string
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    args
		want    *domain.Member
		wantErr bool
	}{
		{
			"success change all",
			fakes{
				fakeGetMemberByUUID: func(ctx context.Context, uuid string) (*domain.Member, error) {
					return &domain.Member{
						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
						DeletedAt:   nil,
						HourlyPrice: util.Int64Ptr(1000),
						Id:          1,
						MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
						Name:        "hoge",
						PhoneNumber: "09000000000",
						Status:      "active",
						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
					}, nil
				},
				fakeUpdateMember: func(ctx context.Context, oldMember domain.Member) (*domain.Member, error) {
					oldMember.UpdatedAt = util.TimeFromStr("2021-10-19 15:09:32")
					return &oldMember, nil
				},
			},
			args{
				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
				params: domain.UpdateMemberJSONBody{
					HourlyPrice: util.Int64Ptr(1200),
					Name:        "fuga",
					PhoneNumber: "08000000000",
					Status:      "other",
				},
			},
			&domain.Member{
				CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
				DeletedAt:   nil,
				HourlyPrice: util.Int64Ptr(1200),
				Id:          1,
				MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
				Name:        "fuga",
				PhoneNumber: "08000000000",
				Status:      "other",
				UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
			},
			false,
		},
		{
			"success change PhoneNumber and HourlyPrice",
			fakes{
				fakeGetMemberByUUID: func(ctx context.Context, uuid string) (*domain.Member, error) {
					return &domain.Member{
						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
						DeletedAt:   nil,
						HourlyPrice: util.Int64Ptr(1000),
						Id:          1,
						MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
						Name:        "hoge",
						PhoneNumber: "09000000000",
						Status:      "active",
						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
					}, nil
				},
				fakeUpdateMember: func(ctx context.Context, oldMember domain.Member) (*domain.Member, error) {
					oldMember.UpdatedAt = util.TimeFromStr("2021-10-19 15:09:32")
					return &oldMember, nil
				},
			},
			args{
				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
				params: domain.UpdateMemberJSONBody{
					HourlyPrice: util.Int64Ptr(1200),
					PhoneNumber: "08000000000",
				},
			},
			&domain.Member{
				CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
				DeletedAt:   nil,
				HourlyPrice: util.Int64Ptr(1200),
				Id:          1,
				MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
				Name:        "hoge",
				PhoneNumber: "08000000000",
				Status:      "active",
				UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
			},
			false,
		},
		{
			"faild not found member by uuid",
			fakes{
				fakeGetMemberByUUID: func(ctx context.Context, uuid string) (*domain.Member, error) {
					return nil, domain.NewError(domain.ErrorTypeContentNotFound)
				},
				fakeUpdateMember: func(ctx context.Context, oldMember domain.Member) (*domain.Member, error) {
					oldMember.UpdatedAt = util.TimeFromStr("2021-10-19 15:09:32")
					return &oldMember, nil
				},
			},
			args{
				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
				params: domain.UpdateMemberJSONBody{
					HourlyPrice: util.Int64Ptr(1200),
					PhoneNumber: "08000000000",
				},
			},
			nil,
			true,
		},
		{
			"faild internal server error",
			fakes{
				fakeGetMemberByUUID: func(ctx context.Context, uuid string) (*domain.Member, error) {
					return &domain.Member{
						CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
						DeletedAt:   nil,
						HourlyPrice: util.Int64Ptr(1000),
						Id:          1,
						MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
						Name:        "hoge",
						PhoneNumber: "09000000000",
						Status:      "active",
						UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
					}, nil
				},
				fakeUpdateMember: func(ctx context.Context, oldMember domain.Member) (*domain.Member, error) {

					return nil, domain.NewError(domain.ErrorTypeInternalError)
				},
			},
			args{
				uuid: "873a2824-8006-4e67-aed7-ec427df5fce8",
				params: domain.UpdateMemberJSONBody{
					HourlyPrice: util.Int64Ptr(1200),
					PhoneNumber: "08000000000",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memberRepo := mockMemberRepo{}
			memberRepo.FakeGetMemberByUUID = tt.fakes.fakeGetMemberByUUID
			memberRepo.FakeUpdateMember = tt.fakes.fakeUpdateMember
			im := interactors.NewMembers(memberRepo)

			got, err := im.UpdateMember(ctx, tt.args.params, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MembersInteractor.UpdateMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MembersInteractor.UpdateMember() = %v, want %v", got, tt.want)
			}

		})
	}
}
