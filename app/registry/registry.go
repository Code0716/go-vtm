// Package registry registry package
package registry

import (
	"github.com/Code0716/go-vtm/app/interfaces/database"
	"github.com/Code0716/go-vtm/app/usecase/interactors"
	"github.com/Code0716/go-vtm/app/usecase/repositories"
)

// Registry returns initialized repositories and interactores.
type Registry struct {
	db database.SQLHandlerInterface
}

// New initializes registry with gorm-database.
func New(db database.SQLHandlerInterface) *Registry {
	return &Registry{
		db: db,
	}
}

/*
	以下に具体的な依存性を解決する初期化処理を書く
*/

// UserRepository returns users database.
func (r *Registry) UserRepository() repositories.UserRepository {
	return database.NewUser(r.db)
}

// UserInteractor returns User interactor.
func (r *Registry) UserInteractor() *interactors.UserInteractor {
	return interactors.NewUser(r.UserRepository())
}
