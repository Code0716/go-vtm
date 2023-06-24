package db_test

import (
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/util"
)

func TestSQLHandler_Create(t *testing.T) {
	t.Parallel()

	testDB, close, err := getTestDB(t, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	user := &domain.User{
		UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:             "hoge",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	tests := []struct {
		name    string
		args    *domain.User
		wantErr bool
	}{
		{
			"success create",
			user,
			false,
		},
		{
			"failed create",
			user,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := db.SQLHandler(*testDB)
			err := d.Create(testCtx, tt.args).Conn.Error
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.Create error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLHandler_where(t *testing.T) {
	t.Parallel()

	user1 := domain.User{
		UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:             "hoge",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	user2 := domain.User{
		UserID:           "dfb759de-1bd9-479e-b99e-9f14e97ebe67",
		Name:             "fuga",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test2.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	user3 := domain.User{
		UserID:           "03da531b-11fd-4430-bf80-c6caf2bca1d8",
		Name:             "hogefuga",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test3.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-10-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	user4 := domain.User{
		UserID:           "6f93ebf2-bfd5-40e8-a907-47d9bd976f06",
		Name:             "hogehoge",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test4.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-10-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	user5 := domain.User{
		UserID:           "10141b9f-3f51-4f8e-972b-91da21303435",
		Name:             "fugafuga",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test5.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-10-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	user6 := domain.User{
		UserID:           "12c90c00-f72c-4ac3-b7e2-2a527ee4459d",
		Name:             "hogehogefugafuga",
		MailAddress:      util.LiteralToPtrGenerics[string]("test@test6.com"),
		PhoneNumber:      util.LiteralToPtrGenerics[string]("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics[int](1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-10-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	seeds := []any{
		user1,
		user2,
		user3,
		user4,
		user5,
		user6,
	}

	testDB, close, err := getTestDB(t, seeds)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	tests := []struct {
		name    string
		where   string
		args    string
		wantID  string
		wantErr bool
	}{
		{
			"success user_id",
			"user_id = ?",
			user1.UserID,
			user1.UserID,
			false,
		},
		{
			"success name",
			"name = ?",
			user2.Name,
			user2.UserID,
			false,
		},
		{
			"failed not found",
			"user_id = ?",
			"hogehoge",
			user2.UserID,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := db.SQLHandler(*testDB)
			var got []domain.User

			err := d.Where(tt.where, tt.args).Find(testCtx, &got).Conn.Error
			if err != nil {
				t.Errorf("SQLHandler.Find error = %v", err)
				return
			}

			if (len(got) == 0) != tt.wantErr {
				t.Error("SQLHandler.Find not found")
				return
			}

			for _, v := range got {
				if v.UserID != tt.wantID {
					t.Errorf("SQLHandler.Where() got = %v, want %v", v, tt.wantID)
				}
			}

		})
	}

}

// func TestSQLHandler_First(t *testing.T) {
// 	t.Parallel()

// 	adminUser := &domain.AdminUser{
// 		Id:          1,
// 		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
// 		Name:        "hogehoge",
// 		MailAddress: "hoge@test.com",
// 		Password:    "hoge",
// 		Permission:  "admin",
// 		Status:      "init",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	fakeUser := &domain.AdminUser{
// 		Id:          1,
// 		AdminId:     "fuga", // 本来はuuid
// 		Name:        "fugafuga",
// 		MailAddress: "fuga@test.com",
// 		Password:    "fuga",
// 		Permission:  "admin",
// 		Status:      "init",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	seeds := []any{
// 		adminUser,
// 	}

// 	testDB, close, err := getTestDB(t, seeds)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer close()

// 	type args struct {
// 		mail string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *domain.AdminUser
// 		wantErr bool
// 	}{
// 		{
// 			"Success",
// 			args{
// 				mail: adminUser.MailAddress,
// 			},
// 			adminUser,
// 			false,
// 		},
// 		{
// 			"faild",
// 			args{
// 				mail: fakeUser.MailAddress,
// 			},
// 			&domain.AdminUser{},
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := db.SQLHandler(*testDB)
// 			var admin *domain.AdminUser

// 			err := d.First(&admin, domain.AdminUser{MailAddress: tt.args.mail}).Conn.Error

// 			if (err != nil) != tt.wantErr {
// 				t.Fatal(err)
// 				t.Errorf("Admin.GetAdminByEmail() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(admin, tt.want) {
// 				t.Errorf("Admin.GetAdminByEmail() got = %v, want %v", admin, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSQLHandler_Delete(t *testing.T) {
// 	t.Parallel()

// 	adminUser1 := &domain.AdminUser{
// 		Id:          1,
// 		AdminId:     "873a2824-8006-4e67-aed7-ec427df5fce8",
// 		Name:        "hogehoge",
// 		MailAddress: "hoge@test.com",
// 		Password:    "hoge",
// 		Permission:  "admin",
// 		Status:      "init",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	adminUser2 := &domain.AdminUser{
// 		Id:          2,
// 		AdminId:     "fugafuga",
// 		Name:        "fugafuga",
// 		MailAddress: "fuga@test.com",
// 		Password:    "fuga",
// 		Permission:  "admin",
// 		Status:      "init",
// 		CreatedAt:   util.TimeFromStr("2021-09-14 15:08:54"),
// 		UpdatedAt:   util.TimeFromStr("2021-10-19 15:09:32"),
// 	}

// 	seeds := []any{
// 		adminUser1,
// 		adminUser2,
// 	}

// 	testDB, close, err := getTestDB(t, seeds)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer close()

// 	type args struct {
// 		value *domain.AdminUser
// 		where domain.AdminUser
// 	}

// 	var adminUser domain.AdminUser

// 	tests := []struct {
// 		name    string
// 		want    domain.AdminUser
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			"Success",
// 			*adminUser1,
// 			args{value: &adminUser, where: domain.AdminUser{AdminId: adminUser1.AdminId}},
// 			false,
// 		},
// 		{
// 			"faild",
// 			*adminUser2,
// 			args{value: &adminUser, where: domain.AdminUser{AdminId: "testest"}},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := db.SQLHandler(*testDB)
// 			err := d.Delete(tt.args.value, tt.args.where).Conn.Error
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("SQLHandler.Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
