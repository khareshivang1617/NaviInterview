package main

import (
	constants "connectFour/constants"
	game "connectFour/game"
	player "connectFour/player"
	"fmt"
)

var GamesPlayed []*game.Game

func main() {
	players := []*player.Player{}

	var name string

	fmt.Print("Enter name for player 1 (R): ")
	fmt.Scan(&name)
	players = append(players, &player.Player{Id: 1, Name: name, Color: "r"})

	fmt.Print("Enter name for player 2 (Y): ")
	fmt.Scan(&name)
	players = append(players, &player.Player{Id: 2, Name: name, Color: "y"})

	fmt.Println("Press 1 to play, 2 to exit")
	var option uint32
	fmt.Scan(&option)

	for option == 1 {
		var game game.Game

		game.InitializeGame(uint32(constants.ROWS), uint32(constants.COLUMNS), players)

		game.PrintBoard()

		game.StartGame()

		for game.GetStatus() == constants.IN_PROGRESS {
			err := game.PlayTurn()
			if err != nil {
				fmt.Println(err)
				game.EndGame()
			}
			game.PrintBoard()
		}

		GamesPlayed = append(GamesPlayed, &game)

		fmt.Println("\n----------------------------------------------------------------\n")
		fmt.Println("Press 1 to play, 2 to exit")
		fmt.Scan(&option)
	}

}
