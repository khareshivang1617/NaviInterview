package game

import (
	constants "connectFour/constants"
	player "connectFour/player"
	"errors"
	"fmt"
)

// ravi.chandora@navi.com

type Game struct {
	board         [][]string
	rows          uint32
	columns       uint32
	players       []*player.Player
	currentPlayer uint32
	status        constants.GameStatus
}

func (game *Game) GetStatus() constants.GameStatus {
	return game.status
}

func (game *Game) StartGame() {
	game.status = constants.IN_PROGRESS
	fmt.Println("Game has started!!!")
}

func (game *Game) EndGame() {
	game.status = constants.NOT_STARTED
	fmt.Println("Game Ended!!!")
}

func (game *Game) InitializeGame(r uint32, c uint32, players []*player.Player) {
	//Initialize Board
	game.rows = r
	game.columns = c

	game.board = [][]string{}

	for i := 0; i < int(r); i++ {
		row := []string{}

		for j := 0; j < int(c); j++ {
			row = append(row, "-")
		}
		game.board = append(game.board, row)
	}

	//Initialize Players
	game.players = players
	game.currentPlayer = 0
}

func (game *Game) CheckIfValidMove(columnSelected uint32) bool {
	if columnSelected < game.columns && game.board[0][columnSelected] == "-" {
		return true
	}

	return false
}

func (game *Game) InsertDisc(columnSelected uint32, color string) (uint32, uint32) {
	i := game.rows - 1

	for i > 0 && game.board[i][columnSelected] != "-" {
		i--
	}

	game.board[i][columnSelected] = color
	game.PrintBoard()

	return i, columnSelected
}

func (game *Game) IsWinningMove(i uint32, j uint32) bool {
	color := game.board[i][j]
	count := 1

	//check four in a col
	k := i - 1
	for i >= 1 && k > 0 && game.board[k][j] == color {
		count++
		k--
	}

	k = i + 1
	for k < game.rows && game.board[k][j] == color {
		count++
		k++
	}

	if count == constants.WIN_COUNT {
		return true
	}

	//check four in a row
	count = 1
	l := j - 1
	for j >= 1 && l > 0 && game.board[i][l] == color {
		count++
		l--
	}

	l = j + 1
	for l < game.columns && game.board[i][l] == color {
		count++
		l++
	}

	if count == constants.WIN_COUNT {
		return true
	}

	//check four in a diagonal

	count = 1
	k = i - 1
	l = j - 1

	for j >= 1 && i >= 1 && k > 0 && l > 0 && game.board[k][l] == color {
		count++
		k--
		l--
	}

	k = i + 1
	l = j + 1

	for k < game.rows && l < game.columns && game.board[k][l] == color {
		count++
		k++
		l++
	}

	if count == constants.WIN_COUNT {
		return true
	}

	count = 1
	k = i + 1
	l = j - 1

	for k < game.rows && j >= 1 && l > 0 && game.board[k][l] == color {
		count++
		k++
		l--
	}

	k = i - 1
	l = j + 1

	for i >= 1 && k > 0 && l < game.columns && game.board[k][l] == color {
		count++
		k--
		l++
	}

	if count == constants.WIN_COUNT {
		return true
	}

	return false

}

func (game *Game) PlayTurn() error {
	if game.status == constants.NOT_STARTED {
		return errors.New("Please Start the Game before playing")
	}

	//check if game is still playabale
	if len(game.getPlayableColumns()) == 0 {
		game.EndGame()
		fmt.Println("Game over! DRAW!")
	}

	currentPlayer := game.players[game.currentPlayer]

	fmt.Println("Player ", currentPlayer.Name, "'s turn, please select a column: ")
	var columnSelected uint32
	fmt.Scan(&columnSelected)

	//check if it is valid column
	if !game.CheckIfValidMove(columnSelected) {
		return errors.New("Invalid Move!")
	}

	//Insert Move into board
	i, j := game.InsertDisc(columnSelected, currentPlayer.Color)

	//check if its winning move
	if game.IsWinningMove(i, j) {
		fmt.Println("Player ", currentPlayer.Name, " wins!!!")
		game.EndGame()
	}

	//Shift the turn to next player
	game.currentPlayer = uint32(int(game.currentPlayer+1) % len(game.players))

	return nil
}

func (game *Game) getPlayableColumns() []uint32 {
	playableColumns := []uint32{}

	for i := uint32(0); i < game.columns; i++ {
		if game.board[0][i] == "-" {
			playableColumns = append(playableColumns, i)
		}
	}

	return playableColumns
}

func (game *Game) PrintBoard() {

	for j := 0; j < int(game.columns); j++ {
		fmt.Print(j, " ")
	}

	fmt.Println()

	for j := 0; j < int(game.columns); j++ {
		fmt.Print("--")
	}

	fmt.Println()

	for i := 0; i < int(game.rows); i++ {
		for j := 0; j < int(game.columns); j++ {
			fmt.Print(game.board[i][j], " ")
		}
		fmt.Println()
	}

	for j := 0; j < int(game.columns); j++ {
		fmt.Print("--")
	}

	fmt.Println()
}
