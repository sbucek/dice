package diceprintger

import (
	"fmt"

	"github.com/sbucek/dice"
)

func PrintRoll(sides int, string comment) {
	fmt.Printf("%s: %dn", comment, dice.Roll(sides))
}
