package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"xiuyiPro/internal/biz"
	"xiuyiPro/internal/data/ent"
	"xiuyiPro/internal/data/ent/turtle"
)

type TurtleRepo struct {
	data *Data
	log  *log.Helper
}

// NewTurtleRepo .
func NewTurtleRepo(data *Data, logger log.Logger) biz.TurtleRepo {
	return &TurtleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *TurtleRepo) Save(ctx context.Context, g *biz.Turtle) (*biz.Turtle, error) {
	_, err := r.data.db.Turtle.
		Create().
		SetQid(g.Qid).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetAnswer(g.Answer).
		SetCategory(g.Category).
		SetCreator(g.Creator).
		SetState(g.State).
		Save(ctx)
	return g, err
}

func (r *TurtleRepo) Update(ctx context.Context, g *biz.Turtle) (*biz.Turtle, error) {
	return nil, nil
}

func (r *TurtleRepo) FindByID(ctx context.Context, id int64) (*biz.Turtle, error) {
	p, err := r.data.db.Turtle.Get(ctx, id)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	return &biz.Turtle{
		Qid:      p.Qid,
		Title:    p.Title,
		Content:  p.Content,
		Answer:   p.Answer,
		Category: p.Category,
		Creator:  p.Creator,
		State:    p.State,
		Ctime:    p.Ctime,
		Mtime:    p.Mtime,
	}, nil
}

func (r *TurtleRepo) ListAll(ctx context.Context) ([]*biz.Turtle, error) {
	ps, err := r.data.db.Turtle.Query().All(ctx)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	rv := make([]*biz.Turtle, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Turtle{
			Qid:      p.Qid,
			Title:    p.Title,
			Content:  p.Content,
			Answer:   p.Answer,
			Category: p.Category,
			Creator:  p.Creator,
			State:    p.State,
			Ctime:    p.Ctime,
			Mtime:    p.Mtime,
		})
	}
	return rv, nil
}

func (r *TurtleRepo) ListByPage(ctx context.Context, offset, limit int32, t *biz.Turtle) ([]*biz.Turtle, error) {
	sql := r.data.db.Turtle.Query()
	if t != nil && len(t.Qid) > 0 {
		sql.Where(turtle.QidEQ(t.Qid))
	}
	if t != nil && len(t.Title) > 0 {
		sql.Where(turtle.TitleContains(t.Title))
	}
	if t != nil && t.State > 0 {
		sql.Where(turtle.StateEQ(t.State))
	}
	if t != nil && t.Category > 0 {
		sql.Where(turtle.CategoryEQ(t.Category))
	}
	if limit > 0 {
		sql.Limit(int(limit)).Offset(int(offset))
	}
	ps, err := sql.All(ctx)
	if err != nil {
		r.log.WithContext(ctx).Error(err)
		return nil, err
	}
	rv := make([]*biz.Turtle, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Turtle{
			Qid:      p.Qid,
			Title:    p.Title,
			Content:  p.Content,
			Answer:   p.Answer,
			Category: p.Category,
			Creator:  p.Creator,
			State:    p.State,
			Ctime:    p.Ctime,
			Mtime:    p.Mtime,
		})
	}
	return rv, nil
}

func (r *TurtleRepo) SaveBatch(ctx context.Context, turtles []*biz.Turtle) ([]*biz.Turtle, error) {
	if len(turtles) == 0 {
		return nil, nil
	}
	bulk := make([]*ent.TurtleCreate, len(turtles))
	for i, t := range turtles {
		bulk[i] = r.data.db.Turtle.Create().
			SetQid(t.Qid).
			SetTitle(t.Title).
			SetContent(t.Content).
			SetAnswer(t.Answer).
			SetCategory(t.Category).
			SetCreator(t.Creator).
			SetState(t.State)
	}
	createdTurtles, err := r.data.db.Turtle.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		r.log.WithContext(ctx).Error(err)
		return nil, err
	}
	result := make([]*biz.Turtle, len(createdTurtles))
	for i, t := range createdTurtles {
		result[i] = &biz.Turtle{
			Qid:      t.Qid,
			Title:    t.Title,
			Content:  t.Content,
			Answer:   t.Answer,
			Category: t.Category,
			Creator:  t.Creator,
			State:    t.State,
			Ctime:    t.Ctime,
			Mtime:    t.Mtime,
		}
	}
	return result, nil
}
