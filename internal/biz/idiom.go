package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// Idiom is an Idiom model.
type Idiom struct {
	ID         int64
	Iid        string
	Answer     string
	Image      string
	Difficulty int32
	Creator    string
	State      int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ViewerRanking struct {
	UID   string
	Face  string
	Name  string
	Index int64
	Score int64
}

// IdiomRepo is an Idiom repo.
type IdiomRepo interface {
	Save(context.Context, *Idiom) (*Idiom, error)
	Update(context.Context, *Idiom) (*Idiom, error)
	FindByID(context.Context, int64) (*Idiom, error)
	ListAll(context.Context) ([]*Idiom, error)
	SaveBatch(context.Context, []*Idiom) ([]*Idiom, error)
	ListByPage(context.Context, int32, int32, *Idiom) ([]*Idiom, int32, error)
	GetTopRanking(context.Context, int64, int64) ([]*ViewerRanking, error)
	UpsertRanking(context.Context, int64, *ViewerRanking, int64) error
	SetRedisKey(context.Context, string, string, int64) error
	GetRedisKey(context.Context, string) (string, error)
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

func (uc *IdiomUsecase) GetTopRanking(ctx context.Context, roomid, limit int64) ([]*ViewerRanking, error) {
	return uc.repo.GetTopRanking(ctx, roomid, limit)

}

func (uc *IdiomUsecase) UpsertRanking(ctx context.Context, roomId int64, viewer *ViewerRanking, timeout int64) error {
	return uc.repo.UpsertRanking(ctx, roomId, viewer, timeout)
}

func (uc *IdiomUsecase) SetGameAnswer(ctx context.Context, gameId int64, answer string, timeout int64) error {
	return uc.repo.SetRedisKey(ctx, getGameAnswerKey(gameId), answer, timeout)
}

func (uc *IdiomUsecase) GetGameAnswer(ctx context.Context, gameId int64) (string, error) {
	return uc.repo.GetRedisKey(ctx, getGameAnswerKey(gameId))
}

func getGameAnswerKey(gameId int64) string {
	return fmt.Sprintf("game_answer_%d", gameId)
}
