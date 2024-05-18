package pkg

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
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
func (cli *Cli) autocompleter(d prompt.Document) []prompt.Suggest {
	var suggestions []prompt.Suggest

	if !strings.Contains(d.Text, " ") {
		suggestions = []prompt.Suggest{
			{Text: "get", Description: "get a key"},
			{Text: "set", Description: "set a key"},
			{Text: "delete", Description: "delete key"},
			{Text: "save", Description: "save database on disk"},
			{Text: "@close", Description: "close cli sesion"},
			{Text: "exists", Description: "return true if key exists  else return false"},
		}
	} else {
		suggestions = []prompt.Suggest{}
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (cli *Cli) Run() {
	done := make(chan bool)
	execute := NewExecuter(cli.Dbd)

	prom := prompt.New(func(text string) {
		text = strings.TrimSpace(text)

		if text == "@close" {
			done <- true
			return
		} else if text == "@clear" || text == "@cls" {
			screen.Clear()
		}
		d := execute.Execute(text)
		var g map[string]any
		json.Unmarshal([]byte(d), &g)
		if g["ok"] != true {
			if e, e2 := g["errorInfo"]; e2 == true {
				color.Red(e.(string))

			} else {
				color.Red("Error :/")
			}
		} else {

			fmt.Printf("%v\n", jsonIdentToStr(g))
		}
	}, cli.autocompleter, prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC, Fn: func(b *prompt.Buffer) {
			done <- true
		}}))
	go func() {
		prom.Run()
	}()

	<-done
	screen.Clear()
}
