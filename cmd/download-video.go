package cmd

import (
	pkg "github/furkanuyank/youtube-download-cli/pkg"

	"github.com/spf13/cobra"
)

func init() {
	cmdDownload.AddCommand(cmdDownloadVideo)
	applyDownloadCommonFlags(cmdDownloadVideo)
	applyDownloadVideoFlags(cmdDownloadVideo)
}

var cmdDownloadVideo = &cobra.Command{
	Use:   "video [URL]",
	Short: "download video",
	Long:  "- Download video with video Url or video code.",
	Args:  cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if downloadOutdir != "" {
			pkg.SetOutdirInApp(downloadOutdir)
		}
		if downloadItag != 0 {
			pkg.DownloadVideoWithItag(args[0], downloadItag, name, extension, true)
		} else if downloadQuality != "" {
			pkg.DownloadVideoWithQuality(args[0], downloadQuality, name, extension, true)
		} else if downloadAudio {
			pkg.DownloadVideoAudio(args[0], name, extension, true)
		} else if downloadMuted {
			pkg.DownloadVideoMuted(args[0], name, extension, true)
		} else {
			pkg.DownloadVideo(args[0], name, extension, true)
		}
	},
}
