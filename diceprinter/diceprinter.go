package diceprintger

import (
	"fmt"

	"github.com/sbucek/dice"
)

func PrintRoll(sides int, comment string) {
	fmt.Printf("%s: %dn", comment, dice.Roll(sides))
}
