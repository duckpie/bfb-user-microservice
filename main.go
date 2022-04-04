package main

import (
	"github.com/spf13/cobra"

	"github.com/wrs-news/bfb-user-microservice/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewRootCmd().Execute())
}
