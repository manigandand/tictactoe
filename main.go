package main

import (
	"fmt"
	"strings"
)

// | a1   | a2   | a3   |
// | b1   | b2   | b3   |
// | c1   | c2   | c3   |
type tictactoe struct {
	board [3][3]string
}

func (t *tictactoe) setMove(position boardPosition, user string) {
	// parse the input
	x, y := position.parse()
	t.board[x][y] = user
}

// checks for winning move vertically and horizontally
func (t *tictactoe) checkWinner() (bool, string) {
	for i := 0; i < 3; i++ {
		if t.board[i][0] == t.board[i][1] && t.board[i][1] == t.board[i][2] {
			return true, t.board[i][0]
		}
		if t.board[0][i] == t.board[1][i] && t.board[1][i] == t.board[2][i] {
			return true, t.board[0][i]
		}
	}
	if t.board[0][0] == t.board[1][1] && t.board[1][1] == t.board[2][2] {
		return true, t.board[0][0]
	}
	if t.board[2][0] == t.board[1][1] && t.board[1][1] == t.board[0][2] {
		return true, t.board[2][0]
	}

	return false, ""
}

func (t *tictactoe) printGrid() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j < 3 {
				fmt.Printf("|")
			}
			fmt.Printf("%05s", t.board[i][j])
		}
		fmt.Println()
	}
}

// will store all the mode moved postions as occupied,
// so others can avoid inputing the same
type moveStore map[boardPosition]bool

// {user1: {a1},{a1},{b1}}
type usersMove map[string][]boardPosition

func (m usersMove) userMovesStr(user string) (res string) {
	moves := m[user]

	for i, mv := range moves {
		res += mv.str()
		if i != len(moves)-1 {
			res += ":"
		}
	}
	return
}

func main() {
	fmt.Println("****************************************")
	fmt.Println("************ TicTacToe Game ************")
	fmt.Println("****************************************")
	moveStore := make(moveStore)
	usersMove := usersMove{
		"user1": []boardPosition{},
		"user2": []boardPosition{},
	}

	// postion: user1
	// upon move check if it filed(made a move)
	grid := &tictactoe{
		board: [3][3]string{},
	}

	moveCount := 1
	user := ""
	for {
		// fmt.Println(strings.Repeat("-", 50))
		grid.printGrid()
		fmt.Println()
		// ask
		user = whoNext(user)
		fmt.Println(user, " please make a move")
		// read the input
		var input boardPosition
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("invalid input, please continue again. Err: ", err.Error())
			continue
		}
		input = boardPosition(strings.TrimSpace(input.str()))

		// validate the input
		if isValid := validPositions[input]; !isValid {
			fmt.Println(input, " is invalid, please continue again!")
			user = whoNext(user)
			continue
		}
		if isVisited := moveStore[input]; isVisited {
			fmt.Println(input, " is already marked, please try other choice!")
			user = whoNext(user)
			continue
		}

		// fmt.Println("user input: ", input)
		// mark visited
		moveStore[input] = true
		usersMove[user] = append(usersMove[user], input)

		grid.setMove(input, user)

		// fmt.Println("MoveCount: ", moveCount)
		// 1,3,5  2,4,6,
		// upon the 3rd move of any user find wether they made the winning pattern
		if moveCount == 5 || moveCount == 6 {
			pattern := usersMove.userMovesStr(user)
			// fmt.Println("user pattern: ", pattern)
			win := winningMoves[pattern]
			if win {
				fmt.Println(user, " Won")
				grid.printGrid()
				fmt.Println()
				return
			}
		}
		if moveCount > 6 {
			if win, u := grid.checkWinner(); win {
				fmt.Println(u, " Won")
				grid.printGrid()
				fmt.Println()
				return
			}
		}

		if moveCount > 9 {
			fmt.Println("Game Drawn")
			grid.printGrid()
			fmt.Println()
			return
		}
		moveCount++
	}
}

func whoNext(previousMove string) string {
	switch previousMove {
	case "user1":
		return "user2"
	case "user2":
		return "user1"
	case "":
		return "user1"
	}
	return ""
}
