package interactor

import (
	"context"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/repository"
	uuid "github.com/satori/go.uuid"
)

// MembersInteractor is member interactor.
type MembersInteractor struct {
	MembersRepository repository.MembersInterface
}

// NewMembers initializes item interactor.
func NewMembers(
	membersRepo repository.MembersInterface,
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
	params.MemberId = uuid.NewV4().String()

	params.Status = domain.StatusCodeInit.GetMembeStatus()
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
