package battleships

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CnC struct {
	game *DefGame
}

func (c CnC) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		response := c.commandLoop(scanner.Text())
		fmt.Println(response)
	}
}

const SpaceSeparator string = " "

func (c *CnC) commandLoop(command string) string {
	commArr := strings.Split(command, SpaceSeparator)
	if len(commArr) == 0 {
		return fmt.Sprintf("No imput")
	}

	command1 := commArr[0]
	switch command1 {
	case "start":
		c.startNewGame()
		return fmt.Sprintf("Starting game")
	case "shoot":
		return fmt.Sprintf("Attacking")
	default:
		return fmt.Sprintf("Unknown command")
	}
}

func (c *CnC) startNewGame() {
	if c.game == nil {
		c.game = MakeNewGame()
	}

	c.game.m = MakeNewMap()
}
