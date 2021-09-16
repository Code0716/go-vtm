package registry

import (
	"github.com/Code0716/go-vtm/app/interactor"
	"github.com/Code0716/go-vtm/app/interfaces/repository"
)

// Getter gets registered instances.
type Getter interface {
	RepositoryGetter
	InteractorGetter
}

// RepositoryGetter gets registered repository instances.
type RepositoryGetter interface {
	AdminRepository() repository.AdminInterface
	MembersRepository() repository.MembersInterface
}

// InteractorGetter gets registered interactor instances.
type InteractorGetter interface {
	AdminInteractor() *interactor.AdminInteractor
	MembersInteractor() *interactor.MembersInteractor
}
