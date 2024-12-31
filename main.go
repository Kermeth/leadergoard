package main

import (
	"fmt"
	"math/rand"
)

func main() {

	leaderboard := &SimpleLeaderboard{
		Name: "Season 0",
	}

	for i := 0; i < 2000; i++ {
		score := rand.Float64() * 10000
		name := fmt.Sprintf("user%d", i)
		err := leaderboard.Add(name, score)
		if err != nil {
			return
		}
	}

	fmt.Println("Top 10")
	top10, err := leaderboard.TopN(10)
	if err != nil {
		return
	}
	for _, v := range top10 {
		fmt.Println(v)
	}

	fmt.Println("Bottom 10")
	bottom10, err := leaderboard.BottomN(10)
	if err != nil {
		return
	}
	for _, v := range bottom10 {
		fmt.Println(v)
	}
}
