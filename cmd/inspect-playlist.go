package cmd

import (
	pkg "github/furkanuyank/youtool/pkg"

	"github.com/spf13/cobra"
)

var playlistTitle bool
var playlistDescription bool
var playlistAuthor bool
var playlistVideos bool

func init() {
	cmdInspect.AddCommand(cmdInspectPlaylist)
	applyInspectPlayistFlags(cmdInspectPlaylist)
}

func applyInspectPlayistFlags(command *cobra.Command) {
	command.Flags().BoolVarP(&playlistTitle, "title", "t", false, "get title of playlist")
	command.Flags().BoolVarP(&playlistDescription, "description", "d", false, "get description of playlist")
	command.Flags().BoolVarP(&playlistAuthor, "author", "a", false, "get author of playlist")
	command.Flags().BoolVarP(&playlistVideos, "videos", "v", false, "get videos of playlist")
}

var cmdInspectPlaylist = &cobra.Command{
	Use:   "playlist [URL]",
	Short: "get playist contents",
	Long:  "Get playlist contents.\nIt can be filtered by using flags",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.PrintPlaylistContent(args[0],
			playlistTitle,
			playlistDescription,
			playlistAuthor,
			playlistVideos)
	},
}
