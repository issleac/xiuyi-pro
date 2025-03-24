package data

import (
	"context"
	"fmt"
	"strconv"
	"xiuyiPro/internal/biz"
)

func rankingKey(id int64) string {
	return fmt.Sprintf("ranking_roomid:%d", id)
}

// UpsertRanking Increment the score of a member if it exists, or insert it with a score of 1 if it does not exist
func (r *IdiomRepo) UpsertRanking(ctx context.Context, roomId int64, uid string) error {
	return r.data.rdb.ZIncrBy(ctx, rankingKey(roomId), 1, uid).Err()
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
		uid, _ := strconv.ParseInt(uidStr, 10, 64)
		rankings[i] = &biz.ViewerRanking{
			UID:   uid,
			Index: int64(i + 1),
			Score: int64(z.Score),
		}
	}
	return rankings, nil
}
