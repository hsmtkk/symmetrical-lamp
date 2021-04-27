package proverb

import (
	"fmt"
	"strings"
)

func Proverb(words []string) string {
	phrases := []string{}
	for i := 0; i < len(words)-1; i++ {
		phrases = append(phrases, fmt.Sprintf("For want of a %s the %s was lost.", words[i], words[i+1]))
	}
	phrases = append(phrases, "And all for the want of a nail.")
	return strings.Join(phrases, "\n")
}
