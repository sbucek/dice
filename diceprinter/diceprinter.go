package diceprinter

import (
	"fmt"

	"github.com/sbucek/dice"
)

func PrintRoll(sides int, comment string) {
	fmt.Printf("%s: %d", comment, dice.Roll(sides))

}
