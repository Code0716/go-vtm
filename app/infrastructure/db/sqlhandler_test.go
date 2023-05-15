package db_test

import (
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/util"
)

func TestSQLHandler_CreateAdmin(t *testing.T) {
	t.Parallel()

	adminUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		MailAddress: "test@test.com",
		Password:    "password",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	adminUser2 := &domain.AdminUser{
		Id:          1,
		AdminId:     "be458a2c-b6b7-472b-823b-0a3755a6004b",
		Name:        "hogehoge",
		MailAddress: "test@test.com",
		Password:    "password",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	testDB, close, err := getTestDB(t, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	tests := []struct {
		name    string
		want    []*domain.AdminUser
		args    domain.AdminUser
		wantErr bool
	}{
		{
			"success",
			[]*domain.AdminUser{
				adminUser,
			},
			*adminUser,
			false,
		},
		{
			"failed same id",
			[]*domain.AdminUser{
				adminUser2,
			},
			*adminUser,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := db.SQLHandler(*testDB)
			err := d.Create(tt.args).Conn.Error
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.CreateAdmin error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLHandler_CreateMember(t *testing.T) {
	t.Parallel()

	Member := &domain.Member{
		Id:          1,
		MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		PhoneNumber: "090000000000",
		HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	Member2 := &domain.Member{
		Id:          2,
		MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		PhoneNumber: "090000000000",
		HourlyPrice: util.LiteralToPtrGenerics[int64](1200),
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	testDB, close, err := getTestDB(t, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	tests := []struct {
		name    string
		want    []*domain.Member
		args    domain.Member
		wantErr bool
	}{
		{
			"success",
			[]*domain.Member{
				Member,
			},
			*Member,
			false,
		},
		{
			"failed same member",
			[]*domain.Member{
				Member2,
			},
			*Member,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := db.SQLHandler(*testDB)
			err := d.Create(tt.args).Conn.Error
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.CreateMember error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLHandler_AdminMemberGetAll(t *testing.T) {
	t.Parallel()

	member1 := &domain.Member{
		Id:          1,
		MemberId:    "873a2824-8006-4e67-aed7-ec427df5fce8",
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
		MemberId:    "uuuu",
		Name:        "uuuu",
		PhoneNumber: "09000000000",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member4 := &domain.Member{
		Id:          4,
		MemberId:    "rrrr",
		Name:        "rrrr",
		PhoneNumber: "09000000000",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member5 := &domain.Member{
		Id:          5,
		MemberId:    "kkkk",
		Name:        "kkkk",
		PhoneNumber: "09000000000",
		Status:      "active",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	member6 := &domain.Member{
		Id:          6,
		MemberId:    "iiii",
		Name:        "iiii",
		PhoneNumber: "09000000000",
		Status:      "other",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}
	seeds := []any{
		member1,
		member2,
		member3,
		member4,
		member5,
		member6,
	}
	testDB, close, err := getTestDB(t, seeds)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	tests := []struct {
		name    string
		args    domain.Pager
		want    []*domain.Member
		count   int64
		wantErr bool
	}{
		{
			"success",
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
			6,
			false,
		},
		{
			"MmberStatus int",
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
			d := db.SQLHandler(*testDB)

			got, gotCount, err := d.AdminMemberGetAll(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.AdminMemberGetAll error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLHandler.AdminMemberGetAll got = %v, want %v", got, tt.want)
			}
			if gotCount != tt.count {
				t.Errorf("SQLHandler.AdminMemberGetAll gotCount = %v, want %v", gotCount, tt.count)
			}
		})
	}

}

func TestSQLHandler_First(t *testing.T) {
	t.Parallel()

	adminUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		MailAddress: "hoge@test.com",
		Password:    "hoge",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	fakeUser := &domain.AdminUser{
		Id:          1,
		AdminId:     "fuga", // 本来はuuid
		Name:        "fugafuga",
		MailAddress: "fuga@test.com",
		Password:    "fuga",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	seeds := []any{
		adminUser,
	}

	testDB, close, err := getTestDB(t, seeds)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	type args struct {
		mail string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.AdminUser
		wantErr bool
	}{
		{
			"Success",
			args{
				mail: adminUser.MailAddress,
			},
			adminUser,
			false,
		},
		{
			"faild",
			args{
				mail: fakeUser.MailAddress,
			},
			&domain.AdminUser{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := db.SQLHandler(*testDB)
			var admin *domain.AdminUser

			err := d.First(&admin, domain.AdminUser{MailAddress: tt.args.mail}).Conn.Error

			if (err != nil) != tt.wantErr {
				t.Fatal(err)
				t.Errorf("Admin.GetAdminByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(admin, tt.want) {
				t.Errorf("Admin.GetAdminByEmail() got = %v, want %v", admin, tt.want)
			}
		})
	}
}

func TestSQLHandler_Delete(t *testing.T) {
	t.Parallel()

	adminUser1 := &domain.AdminUser{
		Id:          1,
		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:        "hogehoge",
		MailAddress: "hoge@test.com",
		Password:    "hoge",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	adminUser2 := &domain.AdminUser{
		Id:          2,
		AdminId:     "fugafuga",
		Name:        "fugafuga",
		MailAddress: "fuga@test.com",
		Password:    "fuga",
		Permission:  "admin",
		Status:      "init",
		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
	}

	seeds := []any{
		adminUser1,
		adminUser2,
	}

	testDB, close, err := getTestDB(t, seeds)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	type args struct {
		value *domain.AdminUser
		where domain.AdminUser
	}

	var adminUser domain.AdminUser

	tests := []struct {
		name    string
		want    domain.AdminUser
		args    args
		wantErr bool
	}{
		{
			"Success",
			*adminUser1,
			args{value: &adminUser, where: domain.AdminUser{AdminId: adminUser1.AdminId}},
			false,
		},
		{
			"faild",
			*adminUser2,
			args{value: &adminUser, where: domain.AdminUser{AdminId: "testest"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := db.SQLHandler(*testDB)
			err := d.Delete(tt.args.value, tt.args.where).Conn.Error
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
