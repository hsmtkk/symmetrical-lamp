package proverb_test

import (
	"testing"

	"github.com/hsmtkk/symmetrical-lamp/proverb"
	"github.com/stretchr/testify/assert"
)

func TestProverb(t *testing.T) {
	words := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
	want := `For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.`
	got := proverb.Proverb(words)
	assert.Equal(t, want, got)
}
