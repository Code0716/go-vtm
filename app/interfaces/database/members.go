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
func (r *MembersRepository) AdminRegistMember(ctx context.Context, member domain.Member) error {
	err := r.SQLHandler.Create(&member)
	if err != nil {
		return err
	}
	return nil
}

// IsMemberExist check member name
func (r *MembersRepository) IsMemberExist(ctx context.Context, name, phone string) (bool, error) {
	var member domain.Member
	bool, err := r.SQLHandler.IsExist(member.TableName(), "name = ? OR phone_number = ?", name, phone)
	if err != nil {
		return false, err
	}
	return bool, nil
}

// AdminMemberGetAll return members found by params
func (r *MembersRepository) AdminMemberGetAll(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error) {

	members, count, err := r.SQLHandler.AdminMemberGetAll(params)

	if err != nil {
		return nil, 0, err
	}

	return members, count, nil
}

// GetMemberByUUID  get member by uuid
func (r *MembersRepository) GetMemberByUUID(ctx context.Context, uuid string) (*domain.Member, error) {
	var member domain.Member
	err := r.SQLHandler.First(&member, domain.Member{MemberId: uuid})
	if err != nil {
		return nil, err
	}
	return &member, nil
}
