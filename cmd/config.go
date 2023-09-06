package cmd

import (
	"fmt"
	pkg "github/furkanuyank/youtube-download-cli/pkg"

	"github.com/spf13/cobra"
)

var getValue bool

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configOutDirCmd)
	configCmd.PersistentFlags().BoolVarP(&getValue, "get", "g", false, "get config value")
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "set config",
	Long:  "Set selected config via commands.",
	Run: func(cmd *cobra.Command, args []string) {
		if getValue {
			lines := pkg.GetEnvLines(pkg.CONFIGENVPATH)
			for _, line := range lines {
				fmt.Println(line)
			}
		} else {
			cmd.Help()
		}
	},
}

var configOutDirCmd = &cobra.Command{
	Use:   "outdir [PATH]",
	Short: "set outdir by default",
	Long:  "Set out direction by default.\nIf nothing is typed, the default is set to ./",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if getValue {
			fmt.Println(pkg.OUTDIRKEY + "=" + pkg.OUTDIR)
		} else {
			if len(args) == 0 {
				pkg.SetEnvValue(pkg.CONFIGENVPATH, pkg.OUTDIRKEY, "./")
				fmt.Println("Out Directory is set './' by default")
			} else {
				pkg.SetEnvValue(pkg.CONFIGENVPATH, pkg.OUTDIRKEY, args[0])
				fmt.Printf("Out Directory is set '%v' by default\n", args[0])
			}
		}

	},
}
