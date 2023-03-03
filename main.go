package main

import (
	"fmt"
	"os"
	"sort"
)

type Horse struct {
	Id int
	Pi int
}

func main() {
	var numberOfHorses int

	fmt.Println("Enter number of horses: ")
	fmt.Fscan(os.Stdin, &numberOfHorses)

	if numberOfHorses <= 1 {
		fmt.Println("\tNot enough horses. Add some more.")
		main()
	} else {
		horses := registerHorses(numberOfHorses)
		sortHorsesByPower(horses)
		result := findClosestCompetitors(horses)
		printResult(result)
	}
}

func createHorse(id int, pi int) *Horse {
	horse := Horse{id, pi}

	return &horse
}

func registerHorses(numberOfHorses int) []Horse {
	var horsePower int
	horses := make([]Horse, 0, numberOfHorses)
	for i := 0; i < numberOfHorses; i++ {
		fmt.Printf("Enter Pi for horse %d: ", i+1)
		fmt.Fscan(os.Stdin, &horsePower)
		horses = append(horses, *createHorse(i, horsePower))
	}
	return horses
}

func sortHorsesByPower(horses []Horse) {
	sort.Slice(horses, func(i, j int) bool {
		return horses[i].Pi < horses[j].Pi
	})
}

func findClosestCompetitors(horses []Horse) [2]Horse {
	var interestingHorses [2]Horse
	difference := horses[len(horses)-1].Pi

	for i := 0; i < len(horses); i++ {
		if len(horses) <= i+1 {
			continue
		}
		tmpDifference := horses[i+1].Pi - horses[i].Pi
		if tmpDifference <= difference {
			difference = tmpDifference
			interestingHorses[0] = horses[i]
			interestingHorses[1] = horses[i+1]
		}
	}
	return interestingHorses
}

func printResult(interestingHorses [2]Horse) {
	fmt.Printf(
		"\n\tMost interesting horses: %d - %d. Their power: %d vs %d\n",
		interestingHorses[0].Id+1, interestingHorses[1].Id+1, interestingHorses[0].Pi, interestingHorses[1].Pi)
}
