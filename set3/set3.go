package main

import (
	"fmt"
	"strings"
	"sampledata"
)

// Relationship status
func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {
	// Check if fromMember follows toMember
	following := socialGraph[fromMember]["following"]
	follows := false
	for _, followee := range strings.Split(following, ",") {
		if followee == toMember {
			follows = true
			break
		}
	}

	// Check if toMember follows fromMember
	followedBy := socialGraph[toMember]["following"]
	followedByFrom := false
	for _, follower := range strings.Split(followedBy, ",") {
		if follower == fromMember {
			followedByFrom = true
			break
		}
	}

	// Determine relationship status
	if follows && followedByFrom {
		return "friends"
	} else if follows {
		return "follower"
	} else if followedByFrom {
		return "followed by"
	} else {
		return "no relationship"
	}
}

// Tic tac toe
func ticTacToe(board [][]string) string {
	// Check rows for a winner
	for _, row := range board {
		if row[0] != "" && row[0] == row[1] && row[1] == row[2] {
			return row[0]
		}
	}

	// Check columns for a winner
	for i := 0; i < len(board); i++ {
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}

	// Check diagonals for a winner
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	// No winner
	return "NO WINNER"
}

// ETA
func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	// Try the direct route (firstStop to secondStop)
	if time, exists := routeMap[firstStop+","+secondStop]; exists {
		return time["travel_time_mins"]
	}

	// Try the reverse route (secondStop to firstStop)
	if time, exists := routeMap[secondStop+","+firstStop]; exists {
		return time["travel_time_mins"]
	}

	// If no route found, return 0
	return 0
}

func main() {
	// Test the relationshipStatus function
	fmt.Println(relationshipStatus("@joaquin", "@chums", sampledata.SocialGraph))

	// Test the ticTacToe function
	fmt.Println(ticTacToe(sampledata.Board))

	// Test the eta function
	fmt.Println(eta("upd", "admu", sampledata.RouteMap))
}
