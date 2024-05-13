package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cli struct {
	Dbd *DB
}

func NewCLI(dbName string) (*Cli, error) {
	f, err := NewDB(dbName)
	if err != nil {
		return nil, err
	}
	return &Cli{
		Dbd: f,
	}, nil
}
func (cli *Cli) Run() {
	execute := NewExecuter(cli.Dbd)
	for {

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "@close" {
			break
		}
		d := execute.Execute(text)
		fmt.Println(d)
	}
}
