package main

import (
	"fmt"
	"os"

	"github.com/kitdevelop-org/vtx/cmd"
	"github.com/kitdevelop-org/vtx/internal/i18n"
)

func main() {
	i18n.Init()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
