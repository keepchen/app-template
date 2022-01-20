package cmd

import (
	"github.com/spf13/cobra"
)

var cfgPath string

var RootCMD = &cobra.Command{
	Use: "app",
}
