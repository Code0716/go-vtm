package repository_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/repository"
	"github.com/Code0716/go-vtm/app/util"
)

func TestMembersRepository_AdminMemberGetAll(t *testing.T) {
	t.Parallel()
	type fakes struct {
		fakeAdminMemberGetAll func(params domain.Pager) ([]*domain.Member, int64, error)
	}

	member1 := &domain.Member{
		Id:          1,
		MemberId:    "hoge", // 本来はuuid
		Name:        "hoge",
		PhoneNumber: "09000000000",
		Status:      "active",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member2 := &domain.Member{
		Id:          2,
		MemberId:    "fuga",
		Name:        "fuga",
		PhoneNumber: "09000000000",
		Status:      "active",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member3 := &domain.Member{
		Id:          3,
		MemberId:    "1111",
		Name:        "1111",
		PhoneNumber: "09000000000",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member4 := &domain.Member{
		Id:          4,
		MemberId:    "4444",
		Name:        "4444",
		PhoneNumber: "09000000000",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member5 := &domain.Member{
		Id:          5,
		MemberId:    "5555",
		Name:        "5555",
		PhoneNumber: "09000000000",
		Status:      "active",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member6 := &domain.Member{
		Id:          6,
		MemberId:    "6666",
		Name:        "6666",
		PhoneNumber: "09000000000",
		Status:      "other",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    domain.Pager
		want    []*domain.Member
		count   int64
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeAdminMemberGetAll: func(params domain.Pager) ([]*domain.Member, int64, error) {
					if params.Limit != 50 || params.Offset != 0 {
						t.Fatal("params not match")
					}
					return []*domain.Member{
							member1,
							member2,
							member3,
							member4,
							member5,
							member6,
						},
						6, nil
				},
			},
			domain.Pager{
				Limit:  50,
				Offset: 0,
				Status: "",
			},
			[]*domain.Member{
				member1,
				member2,
				member3,
				member4,
				member5,
				member6,
			},
			6,
			false,
		},
		{
			"offset 3",
			fakes{
				fakeAdminMemberGetAll: func(params domain.Pager) ([]*domain.Member, int64, error) {
					if params.Limit != 50 || params.Offset != 3 {
						t.Fatal("params not match")
					}
					return []*domain.Member{
							member4,
							member5,
							member6,
						},
						3, nil
				},
			},
			domain.Pager{
				Limit:  50,
				Offset: 3,
				Status: "",
			},
			[]*domain.Member{
				member4,
				member5,
				member6,
			},
			3,
			false,
		},
		{
			"MmberStatus init",
			fakes{
				fakeAdminMemberGetAll: func(params domain.Pager) ([]*domain.Member, int64, error) {
					if params.Limit != 50 || params.Offset != 0 || params.Status != "init" {
						t.Fatal("params not match")
					}
					return []*domain.Member{
							member3,
							member4,
						},
						2, nil
				},
			},
			domain.Pager{
				Limit:  50,
				Offset: 0,
				Status: "init",
			},
			[]*domain.Member{
				member3,
				member4,
			},
			2,
			false,
		},
		{
			"MmberStatus other",
			fakes{
				fakeAdminMemberGetAll: func(params domain.Pager) ([]*domain.Member, int64, error) {
					if params.Limit != 100 || params.Offset != 0 || params.Status != "other" {
						t.Fatal("params not match")
					}
					return []*domain.Member{
							member6,
						},
						1, nil
				},
			},
			domain.Pager{
				Limit:  100,
				Offset: 0,
				Status: "other",
			},
			[]*domain.Member{
				member6,
			},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memberRepo := mockMembersRepo{}
			memberRepo.FakeAdminMemberGetAll = tt.fakes.fakeAdminMemberGetAll
			r := repository.NewMembers(memberRepo)
			got, gotCount, err := r.AdminMemberGetAll(testCtx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Members.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Members.GetAll() got = %v, want %v", got, tt.want)
			}
			if gotCount != tt.count {
				t.Errorf("Members.GetAll() gotCount = %v, want %v", gotCount, tt.count)
			}
		})
	}
}

func TestMembersRepository_AdminRegistMember(t *testing.T) {
	t.Parallel()

	type fakes struct {
		fakeCreateMember func(m interface{}) error
	}

	member := &domain.Member{
		Id:          1,
		MemberId:    "hoge", // 本来はuuid
		Name:        "hogehoge",
		PhoneNumber: "09000000000",
		Password:    util.StrPtr("hoge"),
		Status:      "init",
		HourlyPrice: util.Int64Ptr(1500),
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	type args struct {
		newMember domain.Member
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    args
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeCreateMember: func(m interface{}) error {
					return nil
				},
			},
			args{newMember: *member},
			false,
		},
		{
			"fail",
			fakes{
				fakeCreateMember: func(m interface{}) error {
					return domain.WrapInternalError(errors.New("create faild"))
				},
			},
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memberRepo := mockMembersRepo{}
			memberRepo.FakeCreateMember = tt.fakes.fakeCreateMember
			r := repository.NewMembers(memberRepo)

			if err := r.AdminRegistMember(testCtx, tt.args.newMember); (err != nil) != tt.wantErr {
				t.Errorf("MembersRepository.AdminRegistMember() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMembersRepository_IsMemberExist(t *testing.T) {
	type fakes struct {
		fakeIsMemberExist func(tableName string, query interface{}, args ...interface{}) (bool, error)
	}

	type args struct {
		name  string
		phone string
	}

	tests := []struct {
		name    string
		fakes   fakes
		args    args
		want    bool
		wantErr bool
	}{
		{
			"success",
			fakes{
				fakeIsMemberExist: func(tableName string, query interface{}, args ...interface{}) (bool, error) {
					if args[0] == "" || args[1] == "" {
						return false, errors.New("faild")
					}
					return false, nil
				},
			},
			args{name: "hogehoge", phone: "09000000000"},
			false,
			false,
		},
		{
			"faild",
			fakes{
				fakeIsMemberExist: func(tableName string, query interface{}, args ...interface{}) (bool, error) {
					if args[0] == "" || args[1] == "" {
						return false, errors.New("faild")
					}
					return true, nil
				},
			},
			args{name: "hogehoge", phone: "09000000000"},
			true,
			false,
		},
		{
			"validate error",
			fakes{fakeIsMemberExist: func(tableName string, query interface{}, args ...interface{}) (bool, error) {
				if args[0] == "" || args[1] == "" {
					return false, errors.New("faild")
				}
				return false, nil
			},
			},
			args{},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memberRepo := mockMembersRepo{}
			memberRepo.FakeIsMemberExist = tt.fakes.fakeIsMemberExist
			r := repository.NewMembers(memberRepo)
			got, err := r.IsMemberExist(testCtx, tt.args.name, tt.args.phone)

			if (err != nil) != tt.wantErr {
				t.Errorf("AdminRepository.IsAdminExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AdminRepository.IsAdminExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
