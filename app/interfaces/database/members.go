package database

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
)

// MembersRepository is member database.
type MembersRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewMembers initializes members database.
func NewMembers(sqlHandler SQLHandlerInterface) *MembersRepository {
	return &MembersRepository{
		sqlHandler,
	}
}

// AdminRegistMember regist member to members db
func (r *MembersRepository) AdminRegistMember(_ context.Context, member domain.Member) error {
	err := r.SQLHandler.Create(&member)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMember update member
func (r *MembersRepository) UpdateMember(_ context.Context, member domain.Member) (*domain.Member, error) {
	err := r.SQLHandler.Save(&member).Conn.Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

// IsMemberExist check member name
func (r *MembersRepository) IsMemberExist(_ context.Context, name, phone string) (bool, error) {
	var member domain.Member
	isExist, err := r.SQLHandler.IsExist(
		member.TableName(),
		"name = ? OR phone_number = ?",
		name,
		phone,
	)
	if err != nil {
		return isExist, err
	}
	return isExist, nil
}

// AdminMemberGetAll return members found by params
func (r *MembersRepository) AdminMemberGetAll(_ context.Context, params domain.Pager) ([]*domain.Member, int64, error) {

	members, count, err := r.SQLHandler.AdminMemberGetAll(params)

	if err != nil {
		return nil, 0, err
	}

	return members, count, nil
}

// GetMemberByUUID  get member by uuid
func (r *MembersRepository) GetMemberByUUID(_ context.Context, uuid string) (*domain.Member, error) {
	var member domain.Member
	err := r.SQLHandler.First(&member, domain.Member{MemberId: uuid}).Conn.Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}
