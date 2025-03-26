package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"xiuyiPro/internal/biz"
	"xiuyiPro/internal/data/ent"
	"xiuyiPro/internal/data/ent/idiom"
)

type IdiomRepo struct {
	data *Data
	log  *log.Helper
}

// NewIdiomRepo .
func NewIdiomRepo(data *Data, logger log.Logger) biz.IdiomRepo {
	return &IdiomRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *IdiomRepo) Save(ctx context.Context, i *biz.Idiom) (*biz.Idiom, error) {
	_, err := r.data.db.Idiom.
		Create().
		SetIid(i.Iid).
		SetAnswer(i.Answer).
		SetImage(i.Image).
		SetDifficulty(i.Difficulty).
		SetCreator(i.Creator).
		SetState(i.State).
		Save(ctx)
	return i, err
}

func (r *IdiomRepo) Update(ctx context.Context, g *biz.Idiom) (*biz.Idiom, error) {
	return nil, nil
}

func (r *IdiomRepo) FindByID(ctx context.Context, id int64) (*biz.Idiom, error) {
	p, err := r.data.db.Idiom.Get(ctx, id)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return &biz.Idiom{
		ID:         p.ID,
		Iid:        p.Iid,
		Answer:     p.Answer,
		Image:      p.Image,
		Difficulty: p.Difficulty,
		Creator:    p.Creator,
		State:      p.State,
		CreatedAt:  p.Ctime,
		UpdatedAt:  p.Mtime,
	}, nil
}

func (r *IdiomRepo) FindByGTEID(ctx context.Context, id int64) (*biz.Idiom, error) {
	p, err := r.data.db.Idiom.Query().Where(idiom.IDGTE(id)).First(ctx) // ID >= id
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return &biz.Idiom{
		ID:         p.ID,
		Iid:        p.Iid,
		Answer:     p.Answer,
		Image:      p.Image,
		Difficulty: p.Difficulty,
		Creator:    p.Creator,
		State:      p.State,
		CreatedAt:  p.Ctime,
		UpdatedAt:  p.Mtime,
	}, nil

}

func (r *IdiomRepo) ListAll(ctx context.Context) ([]*biz.Idiom, error) {
	ps, err := r.data.db.Idiom.Query().All(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	ret := make([]*biz.Idiom, 0)
	for _, p := range ps {
		ret = append(ret, &biz.Idiom{
			ID:         p.ID,
			Iid:        p.Iid,
			Answer:     p.Answer,
			Image:      p.Image,
			Difficulty: p.Difficulty,
			Creator:    p.Creator,
			State:      p.State,
			CreatedAt:  p.Ctime,
			UpdatedAt:  p.Mtime,
		})
	}
	return ret, nil
}

func (r *IdiomRepo) ListByPage(ctx context.Context, limit, offset int32, t *biz.Idiom) ([]*biz.Idiom, int32, error) {
	return nil, 0, nil
}

func (r *IdiomRepo) SaveBatch(ctx context.Context, Idioms []*biz.Idiom) ([]*biz.Idiom, error) {
	if len(Idioms) == 0 {
		return nil, nil
	}
	bulk := make([]*ent.IdiomCreate, len(Idioms))
	for i, t := range Idioms {
		bulk[i] = r.data.db.Idiom.Create().
			SetIid(t.Iid).
			SetAnswer(t.Answer).
			SetImage(t.Image).
			SetDifficulty(t.Difficulty).
			SetCreator(t.Creator).
			SetState(t.State)
	}
	createdIdioms, err := r.data.db.Idiom.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Error(err)
		return nil, err
	}
	result := make([]*biz.Idiom, len(createdIdioms))
	for i, t := range createdIdioms {
		result[i] = &biz.Idiom{
			ID:         t.ID,
			Iid:        t.Iid,
			Answer:     t.Answer,
			Image:      t.Image,
			Creator:    t.Creator,
			Difficulty: t.Difficulty,
			State:      t.State,
			CreatedAt:  t.Ctime,
			UpdatedAt:  t.Mtime,
		}
	}
	return result, nil
}
