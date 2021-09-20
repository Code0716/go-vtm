package repositories

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
)

// MembersRepository  is data access methods to RegistMember.
type MembersRepository interface {
	AdminMemberGetAll(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error)
	AdminRegistMember(ctx context.Context, member domain.Member) error
	IsMemberExist(ctx context.Context, name, phone string) (bool, error)
}
