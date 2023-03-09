package main

import (
	"flag"
	"log"

	"github.com/adamnasrudin03/logic-test/game"
	"github.com/adamnasrudin03/logic-test/geometri"
)

func main() {
	var typeProgram = flag.String("type", "dice-game", " The type of program to be run ")
	var players = flag.Int("players", 3, " The number of players in the dice game ")
	var dices = flag.Int("dices", 4, " The number of dice in the dice game ")
	var total = flag.Int("total", 3, " The number of arrays and their contents")

	flag.Parse()

	switch *typeProgram {
	case "dice-game":
		log.Println("Dice game it's running.")
		game.DiceGame(*players, *dices)
	case "calculate-square-diagonal":
		if *total <= 1 {
			log.Println("total must be more than 1")
			break
		}

		log.Println("Logic calculate square diagonal it's running.")
		geometri.CalculateSquareDiagonal(*total)
	default:
		log.Println("Please enter the appropriate program type!")

	}

}
