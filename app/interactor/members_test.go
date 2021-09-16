package interactor_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interactor"
	"github.com/Code0716/go-vtm/app/util"
)

func TestMembers_GetAll(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	successData := []*domain.Member{
		{
			Id:        1,
			MemberId:  "hoge", // 本来はuuid
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
			im := interactor.NewMembers(memberRepo)
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
