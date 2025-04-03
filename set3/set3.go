// Placeholder
package main

func main() {
	// Placeholder
}

// Relationship status
//
// Let's pretend that you are building a new app with social media functionality.
// Users can have relationships with other users.
//
// The two guidelines for describing relationships are:
// 1. Any user can follow any other user.
// 2. If two users follow each other, they are considered friends.
//
// This function describes the relationship that two users have with each other.
//
// Please see the sample data for examples of `socialGraph`.
//
// Params:
// - fromMember, the subject member
// - toMember, the object member
// - socialGraph, the relationship data
//
// Returns:
// - "follower" if fromMember follows toMember; "followed by" if fromMember is followed by toMember; "friends" if fromMember and toMember follow each other; "no relationship otherwise."
func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]interface{}) string {
	// Replace this with your code
	return ""
}

// Tic tac toe
//
// Tic Tac Toe is a common paper-and-pencil game.
// Players must attempt to draw a line of their symbol across a grid.
// The player that does this first is considered the winner.
//
// This function evaluates a Tic Tac Toe game board and returns the winner.
//
// Please see the sample data for examples of `board`.
//
// Params:
// - board, the representation of the Tic Tac Toe board as a square slice of slices of strings. The size of the slice will range between 3x3 to 6x6. The board will never have more than 1 winner. There will only ever be 2 unique symbols at the same time.
//
// Returns:
// - the symbol of the winner, or "NO WINNER" if there is no winner.
func ticTacToe(board [][]string) string {
	// Replace this with your code
	return ""
}

// ETA
//
// A shuttle van service is tasked to travel one way along a predefined circular route.
// The route is divided into several legs between stops.
// The route is fully connected to itself.
//
// This function returns how long it will take the shuttle to arrive at a stop after leaving anothe rstop.
//
// Please see the sample data for examples of `routeMap`.
//
// Params:
// - firstStop, the stop that the shuttle will leave
// - secondStop, the stop that the shuttle will arrive at
// - routeMap, the data describing the routes
//
// Returns:
// - the time that it will take the shuttle to travel from firstStop to secondStop
func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	// Replace this with your code
	return 0
}