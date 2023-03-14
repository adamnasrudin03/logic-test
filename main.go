package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"github.com/adamnasrudin03/logic-test/game"
	"github.com/adamnasrudin03/logic-test/geometri"
	"github.com/adamnasrudin03/logic-test/other"
)

func main() {
	var typeProgram = flag.String("type", "dice-game", " The type of program to be run ")
	var players = flag.Int("players", 3, " The number of players in the dice game ")
	var dices = flag.Int("dices", 4, " The number of dice in the dice game ")
	var total = flag.Int("total", 3, " The number of arrays and their contents")
	var totalTree = flag.Int("totalTree", 5, " The number of total tree")

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
	case "trees-and-the-sun":
		log.Println("Logic count trees and the sun it's running.")
		var trees = []int{}
		var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// make up the whole height of the tree
		for i := 0; i < *totalTree; i++ {
			randomIndex := rand.Intn(len(numbers))
			pick := numbers[randomIndex]
			trees = append(trees, pick)
		}

		fmt.Println("Trees: ", trees)
		countTree := other.TreesAndTheSun(trees, *totalTree)
		fmt.Println(countTree)
	case "min-swap-array":
		log.Println("Logic count minimum swap array it's running.")
		var arr = []int{4, 3, 1, 2}
		count := other.MinSwapArray(arr)
		fmt.Println(count)
	default:
		log.Println("Please enter the appropriate program type!")

	}

}
