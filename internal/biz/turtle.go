package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Turtle is a Turtle model.
type Turtle struct {
	Qid      string
	Title    string
	Content  string
	Answer   string
	Category int32
	Creator  string
	State    int32
	Ctime    time.Time
	Mtime    time.Time
}

// TurtleRepo is a Turtle repo.
type TurtleRepo interface {
	Save(context.Context, *Turtle) (*Turtle, error)
	Update(context.Context, *Turtle) (*Turtle, error)
	FindByID(context.Context, int64) (*Turtle, error)
	ListAll(context.Context) ([]*Turtle, error)
	SaveBatch(context.Context, []*Turtle) ([]*Turtle, error)
	ListByPage(c context.Context, limit, offset int32, t *Turtle) ([]*Turtle, error)
}

// TurtleUsecase is a Turtle usecase.
type TurtleUsecase struct {
	repo TurtleRepo
	log  *log.Helper
}

// NewTurtleUsecase creates a new Turtle usecase.
func NewTurtleUsecase(repo TurtleRepo, logger log.Logger) *TurtleUsecase {
	return &TurtleUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// CreateTurtle creates a Tutle, and returns the new Turtle.
func (uc *TurtleUsecase) CreateTurtle(ctx context.Context, t *Turtle) (*Turtle, error) {
	return uc.repo.Save(ctx, t)
}

// GetTurtleByPage get a Turtle by page.
func (uc *TurtleUsecase) GetTurtleByPage(ctx context.Context, limit int32, offset int32, t *Turtle) ([]*Turtle, error) {
	return uc.repo.ListByPage(ctx, limit, offset, t)
}

// CreateTurtles creates Tutles, and returns the new Turtles.
func (uc *TurtleUsecase) CreateTurtles(ctx context.Context, t []*Turtle) ([]*Turtle, error) {
	return uc.repo.SaveBatch(ctx, t)
}
