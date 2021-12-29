package main

import (
	"flag"
	"fmt"
	"os"

	Q "github.com/ileossa/cli/quitoque"
)

func main() {
	var url string

	fQuitoque := flag.NewFlagSet("quitoque", flag.PanicOnError)
	fQuitoque.StringVar(&url, "url", "none", "url de la recette")

	flag.Parse()

	switch os.Args[1] {
	case "quitoque":
		fQuitoque.Parse(os.Args[2:])
		Q.GetHTML("https://www.quitoque.fr/recettes/papillotte-de-poisson-blanc-aux-tomates-sechees/2021-01")
		break
	default:
		fmt.Println("pls use -h for help")
		os.Exit(1)
	}

}
