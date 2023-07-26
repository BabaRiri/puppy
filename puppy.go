package puppy

import (
	"fmt"

	"github.com/BabaRiri/dog"
)

func Bark() string {
	return "🐶Woof!"
}

func Barks() string {
	return "🐶Woof! 🐶Woof! 🐶Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func version11() {
	fmt.Println("This is version 1.1")
}

func version12() {
	fmt.Println("This is version 1.2")
}
