package drawer

import (
	"errors"
	"fmt"
)

var stages = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

func DrawHangman(stage int) error {
	if stage < 0 || stage >= len(stages) {
		return errors.New("invalid stage number")
	}

	fmt.Println(stages[stage])
	return nil
}
