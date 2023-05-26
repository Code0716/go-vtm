package registry

import (
	"github.com/Code0716/go-vtm/app/usecase/interactors"
	"github.com/Code0716/go-vtm/app/usecase/repositories"
)

// Getter gets registered instances.
type Getter interface {
	RepositoryGetter
	InteractorGetter
}

// RepositoryGetter gets registered database instances.
type RepositoryGetter interface {
	UserRepository() repositories.UserRepository
}

// InteractorGetter gets registered interactor instances.
type InteractorGetter interface {
	UserInteractor() *interactors.UserInteractor
}
