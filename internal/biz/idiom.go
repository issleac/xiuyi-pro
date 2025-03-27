package biz

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
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

type IdiomAnswer struct {
	IdiomId int64 // Idiom.ID
	Answer  string
}

// IdiomRepo is an Idiom repo.
type IdiomRepo interface {
	Save(context.Context, *Idiom) (*Idiom, error)
	Update(context.Context, *Idiom) (*Idiom, error)
	FindByID(context.Context, int64) (*Idiom, error)
	FindByGTEID(context.Context, int64) (*Idiom, error)
	ListAll(context.Context) ([]*Idiom, error)
	SaveBatch(context.Context, []*Idiom) ([]*Idiom, error)
	ListByPage(context.Context, int32, int32, *Idiom) ([]*Idiom, int32, error)
	GetTopRanking(context.Context, string, int64) ([]*ViewerRanking, error)
	UpsertRanking(context.Context, string, *ViewerRanking, int64) error
	SetRedisKey(context.Context, string, interface{}, int64) error
	GetRedisKey(context.Context, string) (interface{}, error)
	DeleteRedisKey(context.Context, string) error
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

func gameIdiomKey(gameId string) string {
	return fmt.Sprintf("game_idiom_%s", gameId)
}

func roomGameKey(roomId int64) string {
	return fmt.Sprintf("room_game_%d", roomId)
}

func rankingKey(id string) string {
	return fmt.Sprintf("game_ranking_%s", id)
}

func giftViewKey(gameId, uid string) string {
	return fmt.Sprintf("gift_game_%s_uid_%s", gameId, uid)
}

func gameCancelFuncKey(gameId string) string {
	return fmt.Sprintf("game_cancel_func_%s", gameId)
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

func (uc *IdiomUsecase) FindByGTEID(ctx context.Context, id int64) (*Idiom, error) {
	res, err := uc.repo.FindByGTEID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (uc *IdiomUsecase) GetTopRanking(ctx context.Context, gameId string, limit int64) ([]*ViewerRanking, error) {
	res, err := uc.repo.GetTopRanking(ctx, rankingKey(gameId), limit)
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

func (uc *IdiomUsecase) UpsertRanking(ctx context.Context, gameId string, viewer *ViewerRanking, timeout int64) error {
	return uc.repo.UpsertRanking(ctx, rankingKey(gameId), viewer, timeout)
}

func (uc *IdiomUsecase) SetGameIdiom(ctx context.Context, gameId string, idiomId int64, answer string, timeout int64) error {
	return uc.repo.SetRedisKey(ctx, gameIdiomKey(gameId), &IdiomAnswer{
		IdiomId: idiomId,
		Answer:  answer,
	}, timeout)
}

func (uc *IdiomUsecase) GetGameIdiom(ctx context.Context, gameId string) (*IdiomAnswer, error) {
	res, err := uc.repo.GetRedisKey(ctx, gameIdiomKey(gameId))
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	answer, _ := res.(*IdiomAnswer)
	return answer, nil
}

func (uc *IdiomUsecase) SetRoomGame(ctx context.Context, roomId int64, gameId string, timeout int64) (err error) {
	return uc.repo.SetRedisKey(ctx, roomGameKey(roomId), gameId, timeout)
}

func (uc *IdiomUsecase) GetRoomGame(ctx context.Context, roomId int64) (string, error) {
	res, err := uc.repo.GetRedisKey(ctx, roomGameKey(roomId))
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	str, _ := res.(string)
	return str, nil
}

func (uc *IdiomUsecase) DelRoomGame(ctx context.Context, roomId int64) error {
	return uc.repo.DeleteRedisKey(ctx, roomGameKey(roomId))
}

func (uc *IdiomUsecase) SetGameGiftView(ctx context.Context, gameId string, uid string, timeout int64) error {
	return uc.repo.SetRedisKey(ctx, giftViewKey(gameId, uid), time.Now(), timeout)
}

func (uc *IdiomUsecase) GetGameGiftView(ctx context.Context, gameId string, uid string) (exist bool, err error) {
	res, err := uc.repo.GetRedisKey(ctx, giftViewKey(gameId, uid))
	if err != nil {
		return false, err
	}
	return res != nil, nil
}

func (uc *IdiomUsecase) SetGameCancelFunc(ctx context.Context, gameId string, f context.CancelFunc, timeout int64) error {
	return uc.repo.SetRedisKey(ctx, gameCancelFuncKey(gameId), f, timeout)
}

func (uc *IdiomUsecase) GetGameCancelFunc(ctx context.Context, gameId string) (f context.CancelFunc, err error) {
	res, err := uc.repo.GetRedisKey(ctx, gameCancelFuncKey(gameId))
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	f, _ = res.(context.CancelFunc)
	return f, nil
}
