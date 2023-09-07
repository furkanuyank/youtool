package cmd

import (
	pkg "github/furkanuyank/youtool/pkg"

	"github.com/spf13/cobra"
)

func init() {
	cmdDownload.AddCommand(cmdDownloadPlaylist)
	applyDownloadCommonFlags(cmdDownloadPlaylist)
	applyDownloadPlaylistFlags(cmdDownloadPlaylist)
}

var cmdDownloadPlaylist = &cobra.Command{
	Use:   "playlist [URL]",
	Short: "download playlist",
	Long: "- Download playlist with playlist Url or code." +
		"\n- Download video with playlist Url or code and number of videos you want in the playlist.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if downloadOutdir != "" {
			pkg.SetOutdirInApp(downloadOutdir)
		}
		if downloadAudio {
			pkg.DownloadPlaylistAudio(args, downloadSimultaneously, name, extension)
		} else if downloadMuted {
			pkg.DownloadPlaylistMuted(args, downloadSimultaneously, name, extension)
		} else {
			pkg.DownloadPlaylist(args, downloadSimultaneously, name, extension)
		}
	},
}
