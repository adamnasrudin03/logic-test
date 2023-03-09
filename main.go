package main

import (
	"flag"
	"log"

	"github.com/adamnasrudin03/logic-test/game"
)

func main() {
	var typeProgram = flag.String("type", "dice-game", " The type of program to be run ")
	var players = flag.Int("players", 3, " The number of players in the dice game ")
	var dices = flag.Int("dices", 4, " The number of dice in the dice game ")

	flag.Parse()

	switch *typeProgram {
	case "dice-game":
		log.Println("Dice game it's running.")
		game.DiceGame(*players, *dices)
	default:
		log.Println("Please enter the appropriate program type!")

	}

}
