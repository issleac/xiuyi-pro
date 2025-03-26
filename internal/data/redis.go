package data

import (
	"context"
	"encoding/json"
	"time"
	"xiuyiPro/internal/biz"
)

// UpsertRanking Increment the score of a member if it exists, or insert it with a score of 1 if it does not exist
func (r *IdiomRepo) UpsertRanking(ctx context.Context, key string, viewer *biz.ViewerRanking, timeout int64) error {
	data, _ := json.Marshal(viewer)
	if err := r.data.rdb.ZIncrBy(ctx, key, 1, string(data)).Err(); err != nil {
		return err
	}
	// 刷新超时时间
	return r.data.rdb.Expire(ctx, key, time.Duration(timeout)).Err()
}

// GetTopRanking Retrieve the top `limit` members based on their scores in descending order
func (r *IdiomRepo) GetTopRanking(ctx context.Context, key string, limit int64) ([]*biz.ViewerRanking, error) {
	res, err := r.data.rdb.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}
	rankings := make([]*biz.ViewerRanking, len(res))
	r.log.WithContext(ctx).Infof("GetTopRanking res(%+v)", res)
	for i, z := range res {
		uidStr, _ := z.Member.(string)
		data := new(biz.ViewerRanking)
		_ = json.Unmarshal([]byte(uidStr), data)
		data.Index, data.Score = int64(i+1), int64(z.Score)
		rankings[i] = data
	}
	return rankings, nil
}

func (r *IdiomRepo) SetRedisKey(ctx context.Context, key string, value interface{}, timeout int64) error {
	return r.data.rdb.Set(ctx, key, value, time.Duration(timeout)).Err()
}

func (r *IdiomRepo) GetRedisKey(ctx context.Context, key string) (interface{}, error) {
	return r.data.rdb.Get(ctx, key).Result()
}
