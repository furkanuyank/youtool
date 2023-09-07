package cmd

import (
	pkg "github/furkanuyank/youtool/pkg"

	"github.com/spf13/cobra"
)

var listAudio bool
var listMuted bool
var listVideo bool

func init() {
	rootCmd.AddCommand(cmdListFormat)
	applyListFormatFlags(cmdListFormat)
}

func applyListFormatFlags(command *cobra.Command) {
	command.Flags().BoolVarP(&listAudio, "audio", "a", false, "list audio only formats")
	command.Flags().BoolVarP(&listMuted, "muted", "m", false, "list no sound video formats")
	command.Flags().BoolVarP(&listVideo, "video", "v", false, "list video formats")
	command.MarkFlagsMutuallyExclusive("audio", "muted", "video")
}

var cmdListFormat = &cobra.Command{
	Use:   "list-format [URL]",
	Short: "list download formats of video",
	Long:  "List download formats of video.\nIt can be filtered by using flags",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		if listAudio {
			pkg.ListAudioOnlyFormat(args)
		} else if listMuted {
			pkg.ListMutedOnlyFormat(args)
		} else if listVideo {
			pkg.ListVideoOnlyFormat(args)
		} else {
			pkg.ListAllFormats(args)
		}
	},
}
