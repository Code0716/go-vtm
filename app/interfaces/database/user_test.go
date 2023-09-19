package database_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/interfaces/database"
	"github.com/Code0716/go-vtm/app/util"
)

// TODO:もう少しうまくできないか。
// errorのtestがリテラルってのもだっさい。
func TestUsersInterface_CreateUser(t *testing.T) {
	t.Parallel()
	type fakes struct {
		fakeCreate func(c context.Context, user any) db.SQLHandler
		fakeFirst  func(c context.Context, user any) db.SQLHandler
	}

	user1 := domain.User{
		UserID:           "873a2824-8006-4e67-aed7-ec427df5fce8",
		Name:             "hoge",
		MailAddress:      util.LiteralToPtrGenerics("test@test.com"),
		PhoneNumber:      util.LiteralToPtrGenerics("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics(1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	user2 := domain.User{
		UserID:           "dcacc0ed-9dc7-49e2-84ac-31f6fabaf952",
		Name:             "fuga",
		MailAddress:      util.LiteralToPtrGenerics("test@test2.com"),
		PhoneNumber:      util.LiteralToPtrGenerics("09000000000"),
		Status:           domain.UserStatusActive,
		Role:             domain.UserRoleCommon,
		EmploymentStatus: domain.EmploymentStatusHourly,
		UnitPrice:        util.LiteralToPtrGenerics(1200),
		DepartmentID:     nil,
		CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
		UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
	}

	testDB, close, err := getTestDB(t, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer close()

	dataBase := db.SQLHandler(*testDB)

	fakeCreateErr := util.LiteralToPtrGenerics("Error 1062 (23000): Duplicate entry '873a2824-8006-4e67-aed7-ec427df5fce8' for key 'user_id'")
	fakeFirstErr := util.LiteralToPtrGenerics("record not found")

	tests := []struct {
		name    string
		fakes   fakes
		args    domain.User
		want    *domain.User
		wantErr *string
	}{
		{
			"success",
			fakes{
				fakeCreate: func(c context.Context, user any) db.SQLHandler {
					return dataBase.Create(c, user1)
				},
				fakeFirst: func(c context.Context, user any) db.SQLHandler {
					return dataBase.First(c, &user)

				},
			},
			user1,
			&domain.User{
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
				CreatedAt:        util.TimeFromStr("2023-09-14 15:08:54"),
				UpdatedAt:        util.TimeFromStr("2023-10-19 15:09:32"),
				DeletedAt:        nil,
			},
			nil,
		},
		{
			"error fakeCreate - Duplicate entry",
			fakes{
				fakeCreate: func(c context.Context, user any) db.SQLHandler {
					return dataBase.Create(c, user1)

				},
				fakeFirst: func(c context.Context, user any) db.SQLHandler {
					return dataBase.First(c, &user1)

				},
			},
			user1,
			nil,
			fakeCreateErr,
		},
		{
			"error fakeFirst - recode not found",
			fakes{
				fakeCreate: func(c context.Context, user any) db.SQLHandler {
					return dataBase.Create(c, user2)

				},
				fakeFirst: func(c context.Context, user any) db.SQLHandler {
					return dataBase.First(c, &domain.User{ID: util.LiteralToPtrGenerics("100")})

				},
			},
			user2,
			nil,
			fakeFirstErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := mockUsersRepo{}
			userRepo.FakeCreate = tt.fakes.fakeCreate
			userRepo.FakeFirst = tt.fakes.fakeFirst

			r := database.NewUser(&userRepo)
			got, err := r.CreateUser(testCtx, tt.args)

			if err != nil && err.Error() != *tt.wantErr {
				t.Errorf(" func TestUsersInterface CreateUser() = %v, wantErr %v", err, *tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("func TestUsersInterface CreateUser() = got = %v, want %v", got, tt.want)
			}
		})
	}
}
