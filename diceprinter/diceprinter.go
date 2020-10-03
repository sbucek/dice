package diceprinter

import (
	"fmt"

	"github.com/sbucek/dice"
)

func PrintRoll(sides int, comment string) string {
	fmt.Printf("%s: %dn", comment, dice.Roll(sides))
}
