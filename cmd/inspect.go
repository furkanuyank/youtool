package cmd

import (
	pkg "github/furkanuyank/youtube-download-cli/pkg"

	"github.com/spf13/cobra"
)

var title bool
var description bool
var author bool
var length bool
var views bool
var thumbnails bool

func init() {
	rootCmd.AddCommand(cmdInspect)
	cmdInspect.AddCommand(cmdInspectVideo)
	applyInspectFlags(cmdInspect)
	applyInspectFlags(cmdInspectVideo)
}

func applyInspectFlags(command *cobra.Command) {
	command.Flags().BoolVarP(&title, "title", "t", false, "get title of video")
	command.Flags().BoolVarP(&description, "description", "d", false, "get description of video")
	command.Flags().BoolVarP(&author, "author", "a", false, "get author of video")
	command.Flags().BoolVarP(&length, "length", "l", false, "get duration of video")
	command.Flags().BoolVarP(&views, "views", "w", false, "get views of video")
	command.Flags().BoolVarP(&thumbnails, "photo", "p", false, "get thumbnails of video")
}

var cmdInspect = &cobra.Command{
	Use:   "inspect [URL]",
	Short: "get video contents",
	Long:  "Get video contents.\nIt can be filtered by using flags",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.PrintContent(
			args,
			title,
			description,
			author,
			length,
			views,
			thumbnails,
		)
	},
}

var cmdInspectVideo = &cobra.Command{
	Use:   "video [URL]",
	Short: "get video contents",
	Long:  "Get video contents.\nIt can be filtered by using flags",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.PrintContent(
			args,
			title,
			description,
			author,
			length,
			views,
			thumbnails,
		)
	},
}
