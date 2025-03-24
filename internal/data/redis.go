package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"xiuyiPro/internal/biz"
)

func rankingKey(id int64) string {
	return fmt.Sprintf("ranking_roomid:%d", id)
}

// UpsertRanking Increment the score of a member if it exists, or insert it with a score of 1 if it does not exist
func (r *IdiomRepo) UpsertRanking(ctx context.Context, roomId int64, viewer *biz.ViewerRanking, timeout int64) error {
	data, _ := json.Marshal(viewer)
	if err := r.data.rdb.ZIncrBy(ctx, rankingKey(roomId), 1, string(data)).Err(); err != nil {
		return err
	}
	// 刷新超时时间
	return r.data.rdb.Expire(ctx, rankingKey(roomId), time.Duration(timeout)).Err()
}

// GetTopRanking Retrieve the top `limit` members based on their scores in descending order
func (r *IdiomRepo) GetTopRanking(ctx context.Context, roomId int64, limit int64) ([]*biz.ViewerRanking, error) {
	res, err := r.data.rdb.ZRevRangeWithScores(ctx, rankingKey(roomId), 0, limit-1).Result()
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

func (r *IdiomRepo) SetRedisKey(ctx context.Context, key, value string, timeout int64) error {
	return r.data.rdb.Set(ctx, key, value, time.Duration(timeout)).Err()
}

func (r *IdiomRepo) GetRedisKey(ctx context.Context, key string) (string, error) {
	return r.data.rdb.Get(ctx, key).Result()
}
