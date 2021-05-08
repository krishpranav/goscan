package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/krishpranav/goscan/core/cli"
	"github.com/krishpranav/goscan/core/utils"
)

var (
	author  string
	version string
)

func showBanner() {
	name := fmt.Sprintf("goscan (v.%s)", version)
	banner := `GOSCAN`

	// Shell width
	all_lines := strings.Split(banner, "\n")
	w := len(all_lines[1])

	// Print Centered
	fmt.Println(banner)
	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))
	fmt.Println()
}

func initCore() {
	utils.CheckSudo()
	showBanner()
	utils.InitConfig()
}

// main function
func main() {
	// Setup core
	initCore()

	// Start CLI
	p := prompt.New(
		cli.Executor,
		cli.Completer,
		prompt.OptionTitle("goscan: Interactive Network Scanner"),
		prompt.OptionPrefix("[goscan] > "),
		prompt.OptionInputTextColor(prompt.White),
	)
	p.Run()
}
