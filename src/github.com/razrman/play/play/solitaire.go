// This program solves the (English) peg
// solitaire board game.
// http://en.wikipedia.org/wiki/Peg_solitaire

package main

import "fmt"

const N = 11 + 1 // length of a row (+1 for \n)

// The board must be surrounded by 2 illegal
// fields in each direction so that move()
// doesn't need to check the board boundaries.
// Periods represent illegal fields,
// â are pegs, and â are holes.

var board = []rune(
`...........
...........
....âââ....
....âââ....
..âââââââ..
..âââââââ..
..âââââââ..
....âââ....
....âââ....
...........
...........
`
)

// center is the position of the center hole if
// there is a single one; otherwise it is -1.
var center int

func init() {
	n := 0
	for pos, field := range board {
		if field == 'â'�' {
			center = pos
			n++
		}
	}
	if n != 1 {
		center = -1 // no single hole
	}
}

var moves int // number of times move is called

// move tests if there is a peg at position pos that
// can jump over another peg in direction dir. If the
// move is valid, it is executed and move returns true.
// Otherwise, move returns false.
func move(pos, dir int) bool {
	moves++
	if board[pos] == 'â' && board[pos+dir] == 'â' && board[pos+2*dir] == 'â' {
		board[pos] = 'â'
		board[pos+dir] = 'â'
		board[pos+2*dir] = 'â'
		return true
	}
	return false
}

// unmove reverts a previously executed valid move.
func unmove(pos, dir int) {
	board[pos] = 'â'
	board[pos+dir] = 'â'
	board[pos+2*dir] = 'â'
}

// solve tries to find a sequence of moves such that
// there is only one peg left at the end; if center is
// >= 0, that last peg must be in the center position.
// If a solution is found, solve prints the board after
// each move in a backward fashion (i.e., the last
// board position is printed first, all the way back to
// the starting board position).
func solve() bool {
	var last, n int
	for pos, field := range board {
		// try each board position
		if field == 'â' {
			// found a peg
			for _, dir := range [...]int{-1, -N, +1, +N} {
				// try each direction
				if move(pos, dir) {
					// a valid move was found and executed,
					// see if this new board has a solution
					if solve() {
						unmove(pos, dir)
						fmt.Println(string(board))
						return true
					}
					unmove(pos, dir)
				}
			}
			last = pos
			n++
		}
	}
	// tried each possible move
	if n == 1 && (center < 0 || last == center) {
		// there's only one peg left
		fmt.Println(string(board))
		return true
	}
	// no solution found for this board
	return false
}

func main() {
	if !solve() {
		fmt.Println("no solution found")
	}
	fmt.Println(moves, "moves tried")
}
