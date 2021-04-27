package main

import (
	"fmt"

	"github.com/hsmtkk/symmetrical-lamp/proverb"
)

func main() {
	words := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
	fmt.Println(proverb.Proverb(words))
}
