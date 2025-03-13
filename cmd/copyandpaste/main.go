package main

import (
	"log"

	"github.com/alingse/copyandpaste"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	setting := copyandpaste.LinterSetting{}

	analyzer, err := copyandpaste.NewAnalyzer(setting)
	if err != nil {
		log.Fatal(err)
	}

	singlechecker.Main(analyzer)
}
