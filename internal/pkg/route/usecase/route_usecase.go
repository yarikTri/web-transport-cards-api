package usecase

import (
	"github.com/yarikTri/web-transport-cards/internal/models"
	"github.com/yarikTri/web-transport-cards/internal/pkg/route"
)

// Usecase implements route.Usecase
type Usecase struct {
	repo route.Repository
}

func NewUsecase(rr route.Repository) *Usecase {
	return &Usecase{
		repo: rr,
	}
}

func (u *Usecase) GetByID(routeID uint32) (*models.Route, error) {
	return u.repo.GetByID(routeID)
}

func (u *Usecase) List() ([]models.Route, error) {
	return u.repo.List()
}

func (u *Usecase) Create(route models.Route) (*models.Route, error) {
	return u.repo.Create(route)
}

func (u *Usecase) DeleteByID(routeID uint32) error {
	return u.repo.DeleteByID(routeID)
}
