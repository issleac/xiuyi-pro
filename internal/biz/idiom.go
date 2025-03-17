package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Idiom is an Idiom model.
type Idiom struct {
	ID         int64
	Iid        string
	Name       string
	Image      string
	Difficulty int32
	Creator    string
	State      int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// IdiomRepo is an Idiom repo.
type IdiomRepo interface {
	Save(context.Context, *Idiom) (*Idiom, error)
	Update(context.Context, *Idiom) (*Idiom, error)
	FindByID(context.Context, int64) (*Idiom, error)
	ListAll(context.Context) ([]*Idiom, error)
	SaveBatch(context.Context, []*Idiom) ([]*Idiom, error)
	ListByPage(context.Context, int32, int32, *Idiom) ([]*Idiom, int32, error)
}

// IdiomUsecase is an Idiom usecase.
type IdiomUsecase struct {
	repo IdiomRepo
	log  *log.Helper
}

// NewIdiomUsecase creates a new Idiom usecase.
func NewIdiomUsecase(repo IdiomRepo, logger log.Logger) *IdiomUsecase {
	return &IdiomUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// CreateIdiom creates an Idiom, and returns the new Idiom.
func (uc *IdiomUsecase) CreateIdiom(ctx context.Context, i *Idiom) (*Idiom, error) {
	return uc.repo.Save(ctx, i)
}

// GetIdiomByPage gets Idioms by page.
func (uc *IdiomUsecase) GetIdiomByPage(ctx context.Context, limit int32, offset int32, i *Idiom) ([]*Idiom, int32, error) {
	return uc.repo.ListByPage(ctx, limit, offset, i)
}

// CreateIdioms creates Idioms, and returns the new Idioms.
func (uc *IdiomUsecase) CreateIdioms(ctx context.Context, i []*Idiom) ([]*Idiom, error) {
	return uc.repo.SaveBatch(ctx, i)
}

func (uc *IdiomUsecase) FindByID(ctx context.Context, id int64) (*Idiom, error) {
	return uc.repo.FindByID(ctx, id)
}
