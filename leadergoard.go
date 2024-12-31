package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type SimpleLeaderboard struct {
	Name string
}

// Add Adds player with its score to the leaderboard
func (sl SimpleLeaderboard) Add(player string, score float64) error {
	ctx := context.Background()
	rdb := redisClient()
	add := rdb.ZAdd(
		ctx,
		sl.Name,
		redis.Z{
			Score:  score,
			Member: player,
		},
	)
	return add.Err()
}

// TopN return a slice with top n players with their scores
func (sl SimpleLeaderboard) TopN(n int64) ([]redis.Z, error) {
	ctx := context.Background()
	rdb := redisClient()
	topN, err := rdb.ZRevRangeWithScores(ctx, sl.Name, 0, n).Result()
	return topN, err
}

// BottomN return a slice with bottom n players with their scores
func (sl SimpleLeaderboard) BottomN(n int64) ([]redis.Z, error) {
	ctx := context.Background()
	rdb := redisClient()
	count, err := rdb.ZCard(ctx, sl.Name).Result()
	if err != nil {
		return []redis.Z{}, err
	}
	bottomN, err := rdb.ZRevRangeWithScores(ctx, sl.Name, count-n, count).Result()
	if err != nil {
		return []redis.Z{}, err
	}
	return bottomN, nil
}

// Rank return the position in the rank of the player.
func (sl SimpleLeaderboard) Rank(player string) (int64, error) {
	ctx := context.Background()
	rdb := redisClient()
	rank, err := rdb.ZRevRank(ctx, sl.Name, player).Result()
	if err != nil {
		return 0, err
	}
	return rank, nil
}

// Score return the score of the player
func (sl SimpleLeaderboard) Score(player string) (float64, error) {
	ctx := context.Background()
	rdb := redisClient()
	score, err := rdb.ZScore(ctx, sl.Name, player).Result()
	if err != nil {
		return 0, err
	}
	return score, nil
}

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
