package interactors

import (
	"context"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/repositories"
	"github.com/Code0716/go-vtm/app/util"
)

// MembersInteractor is member interactor.
type MembersInteractor struct {
	MembersRepository repositories.MembersRepository
}

// NewMembers initializes item interactor.
func NewMembers(
	membersRepo repositories.MembersRepository,
) *MembersInteractor {
	return &MembersInteractor{
		MembersRepository: membersRepo,
	}
}

// MemberGetAll returns member list
// im: members interactor
func (im *MembersInteractor) MemberGetAll(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error) {
	memberList, count, err := im.MembersRepository.AdminMemberGetAll(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return memberList, count, nil
}

// RegistMember regist member
// im: members interactor
func (im *MembersInteractor) RegistMember(ctx context.Context, params domain.Member) error {
	currentTime := time.Now()
	params.CreatedAt = currentTime
	params.UpdatedAt = currentTime
	params.MemberId = util.UUIDGenerator()

	params.Status = domain.StatusCodeInit.GetWorkStatus()
	err := im.MembersRepository.AdminRegistMember(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

// IsMemberExist check regist member
// im: members interactor
func (im *MembersInteractor) IsMemberExist(ctx context.Context, name, phone string) (bool, error) {
	isExist, err := im.MembersRepository.IsMemberExist(ctx, name, phone)
	return isExist, err
}

//  GetMemberByUUID get regist member by uuid
// im: members interactor
func (im *MembersInteractor) GetMemberByUUID(ctx context.Context, uuid string) (*domain.Member, error) {
	member, err := im.MembersRepository.GetMemberByUUID(ctx, uuid)
	return member, err
}
