package database

// import (
// 	"context"

// 	"github.com/Code0716/go-vtm/app/domain"
// )

// // UsersRepository is user database.
// type UsersRepository struct {
// 	SQLHandler SQLHandlerInterface
// }

// // NewUsers initializes users database.
// func NewUsers(sqlHandler SQLHandlerInterface) *UsersRepository {
// 	return &UsersRepository{
// 		sqlHandler,
// 	}
// }

// // AdminRegistUser regist user to users db
// func (r *UsersRepository) AdminRegistUser(_ context.Context, user domain.User) error {
// 	err := r.SQLHandler.Create(&user).Conn.Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // UpdateUser update user
// func (r *UsersRepository) UpdateUser(_ context.Context, user domain.User) (*domain.User, error) {
// 	err := r.SQLHandler.Save(&user).Conn.Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// // IsUserExist check user name
// func (r *UsersRepository) IsUserExist(_ context.Context, name, phone string) (bool, error) {
// 	var user domain.User
// 	isExist, err := r.SQLHandler.IsExist(
// 		user.TableName(),
// 		"name = ? OR phone_number = ?",
// 		name,
// 		phone,
// 	)
// 	if err != nil {
// 		return isExist, err
// 	}
// 	return isExist, nil
// }

// // AdminUserGetAll return users found by params
// func (r *UsersRepository) AdminUserGetAll(_ context.Context, params domain.Pager) ([]*domain.User, int64, error) {

// 	users, count, err := r.SQLHandler.AdminUserGetAll(params)

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return users, count, nil
// }

// // GetUserByUUID  get user by uuid
// func (r *UsersRepository) GetUserByUUID(_ context.Context, uuid string) (*domain.User, error) {
// 	var user domain.User
// 	err := r.SQLHandler.First(&user, domain.User{UserId: uuid}).Conn.Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
