package main

import (
	"fmt"
	"testing"
)

func TestSimpleLeaderboard_Add(t *testing.T) {
	leaderboard := &SimpleLeaderboard{
		Name: "Add-Test-Ranking",
	}
	err := leaderboard.Add("Test", 10)
	if err != nil {
		t.Errorf("Error adding to leaderboard: %v", err)
	}
	err = leaderboard.Add("Test", 45)
	if err != nil {
		t.Errorf("Error adding to leaderboard: %v", err)
	}
	score, err := leaderboard.Score("Test")
	if err != nil {
		t.Errorf("Error getting score: %v", err)
	}
	if score != 45 {
		t.Errorf("Score was incorrect, got: %f, want: %d.", score, 45)
	}
}

func TestSimpleLeaderboard_TopN(t *testing.T) {
	leaderboard := &SimpleLeaderboard{
		Name: "Top10-Test-Ranking",
	}
	for i := 0; i <= 30; i++ {
		player := fmt.Sprintf("player%d", i)
		score := i
		err := leaderboard.Add(player, float64(score))
		if err != nil {
			t.Errorf("Error adding to leaderboard: %v", err)
		}
	}
	top10, err := leaderboard.TopN(10)
	if err != nil {
		t.Errorf("Error getting top 10: %v", err)
	}
	for index, player := range top10 {
		shouldBe := 30 - index // Top10 players should be from 30 to 20
		if player.Member != fmt.Sprintf("player%d", shouldBe) {
			t.Errorf("Player was incorrect, got: %s, want: %d.", player.Member, shouldBe)
		}
	}
}

func TestSimpleLeaderboard_BottomN(t *testing.T) {
	leaderboard := &SimpleLeaderboard{
		Name: "Bottom10-Test-Ranking",
	}
	for i := 0; i <= 30; i++ {
		player := fmt.Sprintf("player%d", i)
		score := i
		err := leaderboard.Add(player, float64(score))
		if err != nil {
			t.Errorf("Error adding to leaderboard: %v", err)
		}
	}
	bottom10, err := leaderboard.BottomN(10)
	if err != nil {
		t.Errorf("Error getting bottom 10: %v", err)
	}
	for index, player := range bottom10 {
		shouldBe := 9 - index // Bottom10 should be players from 0 to 9
		if player.Member != fmt.Sprintf("player%d", shouldBe) {
			t.Errorf("Player was incorrect, got: %s, want: %d.", player.Member, shouldBe)
		}
	}
}

func TestSimpleLeaderboard_Rank(t *testing.T) {
	leaderboard := &SimpleLeaderboard{
		Name: "Ranking-Test-Ranking",
	}
	for i := 0; i <= 30; i++ {
		player := fmt.Sprintf("player%d", i)
		score := i
		err := leaderboard.Add(player, float64(score))
		if err != nil {
			t.Errorf("Error adding to leaderboard: %v", err)
		}
	}
	rank, err := leaderboard.Rank("player26")
	if err != nil {
		t.Errorf("Error getting rank: %v", err)
	}
	if rank != 30-26 {
		t.Errorf("Rank was incorrect, got: %d, want: %d.", rank, 30-26)
	}
}

func TestSimpleLeaderboard_Score(t *testing.T) {
	leaderboard := &SimpleLeaderboard{
		Name: "Score-Test-Ranking",
	}
	for i := 0; i <= 30; i++ {
		player := fmt.Sprintf("player%d", i)
		score := i
		err := leaderboard.Add(player, float64(score))
		if err != nil {
			t.Errorf("Error adding to leaderboard: %v", err)
		}
	}
	score, err := leaderboard.Score("player26")
	if err != nil {
		t.Errorf("Error getting score: %v", err)
	}
	if score != 26 {
		t.Errorf("Score was incorrect, got: %f, want: %d.", score, 26)
	}
}
