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

// AdminRepository returns Admin database.
func (r *Registry) AdminRepository() repositories.AdminRepository {
	return database.NewAdmin(r.db)
}

// AdminInteractor returns Admin interactor.
func (r *Registry) AdminInteractor() *interactors.AdminInteractor {
	return interactors.NewAdmin(r.AdminRepository())
}

// MembersRepository returns members database.
func (r *Registry) MembersRepository() repositories.MembersRepository {
	return database.NewMembers(r.db)
}

// MembersInteractor returns members interactor.
func (r *Registry) MembersInteractor() *interactors.MembersInteractor {
	return interactors.NewMembers(r.MembersRepository())
}

// AttendanceRepository returns attendance database.
func (r *Registry) AttendanceRepository() repositories.AttendanceRepository {
	return database.NewAttendance(r.db)
}

// AttendanceInteractor returns attendance interactor.
func (r *Registry) AttendanceInteractor() *interactors.AttendanceInteractor {
	return interactors.NewAttendance(r.AttendanceRepository())
}
