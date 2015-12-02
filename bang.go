package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
)

func main() {
	homePath, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	bangs := NewBangs(homePath + "/.bang")
	app := cli.NewApp()
	app.Name = "Bang"
	app.Usage = "Text snippets in your terminal"

	app.Action = func(c *cli.Context) {
		bangKey := ""
		bangValue := ""

		if len(c.Args()) == 0 {
			for _, bang := range bangs.ListBangs() {
				fmt.Println(bang)
			}
		}

		if len(c.Args()) == 1 {
			bangKey = c.Args()[0]
			println(bangs.GetBang(bangKey))
		}

		if len(c.Args()) == 2 {
			bangKey = c.Args()[0]
			bangValue = c.Args()[1]
			bangs.SetBang(bangKey, bangValue)
		}
	}
	app.Run(os.Args)
}

type Bangs struct {
	bangFilePath string
	bangList     map[string]string
}

func NewBangs(filePath string) *Bangs {
	b := new(Bangs)
	b.bangList = make(map[string]string)
	b.bangFilePath = filePath
	if _, err := os.Stat(b.bangFilePath); os.IsNotExist(err) {
		error := ioutil.WriteFile(b.bangFilePath, []byte("{}"), 0644)
		if error != nil {
			panic(error)
		}
	}
	encoded, error := ioutil.ReadFile(b.bangFilePath)
	if error != nil {
		panic(error)
	}
	err := json.Unmarshal([]byte(encoded), &b.bangList)
	if err != nil {
		panic(err)
	}
	return b
}

func (b Bangs) ListBangs() []string {
	var allBangs []string
	for k, v := range b.bangList {
		allBangs = append(allBangs, fmt.Sprintf("%s: %s", k, v))
	}
	return allBangs
}

func (b Bangs) GetBang(key string) string {
	return b.bangList[key]
}

func (b Bangs) SetBang(key string, value string) {
	b.bangList[key] = value
	jsonString, error := json.Marshal(b.bangList)
	if error != nil {
		panic(error)
	}
	err := ioutil.WriteFile(b.bangFilePath, jsonString, 0644)
	if err != nil {
		panic(err)
	}
	return
}
