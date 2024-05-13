package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Troxsoft/SkikDB/pkg"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Expectative argument: 1 but found: nothing")
	} else {
		if len(os.Args)-1 == 2 && os.Args[1] == "new" {
			if strings.HasSuffix(os.Args[2], ".skl") {
				f, _ := strings.CutSuffix(os.Args[2], ".skl")
				_, err := pkg.CreateDB(f)
				if err != nil {
					fmt.Printf("%v\n", err)
				} else {
					fmt.Printf("DB:%s created sucessfull :)\n", f)
				}
			} else {
				_, err := pkg.CreateDB(os.Args[2])
				if err != nil {
					fmt.Printf("%v\n", err)
				} else {
					fmt.Printf("DB:%s created sucessfull :)\n", os.Args[2])
				}
			}

		} else if len(os.Args)-1 == 0 {
			fmt.Printf(
				`	SkikDB version: %v   A Database that supports JSON/key value ✔
		SkikLang version %v

`, pkg.VERSION, pkg.LANG_VERSION)
		} else {
			fmt.Println("Invalid arguments :/ 😒")
		}
	}
	// lang := pkg.NewSkikLang(
	// 	`set *`)
	// lang.PreInit()
	// err := lang.Tokenize()
	// lang.Tokens = pkg.RemoveGarbageTokens(lang.Tokens)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // var g any
	// // json.Unmarshal([]byte("[1,2,3,4,\"hola\"]"), &g)
	// // fmt.Printf("%+v\n", g)
	// b, _ := json.Marshal(lang.Tokens)
	// fmt.Printf("%v\n", string(b))
}
