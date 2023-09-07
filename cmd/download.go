package cmd

import (
	"github/furkanuyank/youtube-download-cli/pkg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdDownload)
	applyDownloadCommonFlags(cmdDownload)
	applyDownloadVideoFlags(cmdDownload)
}

var cmdDownload = &cobra.Command{
	Use:   "download [URL]",
	Short: "download video or playlist",
	Long: "- Download video with video Url or video code." +
		"\n- Download playlist with playlist Url or code." +
		"\n- Download video with playlist Url or code and number of videos you want in the playlist.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if downloadOutdir != "" {
			pkg.SetOutdirInApp(downloadOutdir)
		}
		if len(args) == 1 && pkg.IsVideo(args[0]) {
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
		} else {
			if downloadAudio {
				pkg.DownloadPlaylistAudio(args, downloadSimultaneously, name, extension)
			} else if downloadMuted {
				pkg.DownloadPlaylistMuted(args, downloadSimultaneously, name, extension)
			} else {
				pkg.DownloadPlaylist(args, downloadSimultaneously, name, extension)
			}
		}
	},
}

var downloadOutdir string
var name string
var extension string
var downloadItag int
var downloadQuality string
var downloadAudio bool
var downloadMuted bool
var downloadSimultaneously bool

func applyDownloadCommonFlags(command *cobra.Command) {
	command.Flags().StringVarP(&downloadOutdir, "outdir", "o", "", "set out directory")
	command.Flags().BoolVarP(&downloadAudio, "audio", "a", false, "download audio")
	command.Flags().BoolVarP(&downloadMuted, "muted", "m", false, "download muted video")
	command.Flags().StringVarP(&name, "name", "n", "", "set name of video")
	command.Flags().StringVarP(&extension, "extension", ".", "", "set extension of video")
}

func applyDownloadVideoFlags(command *cobra.Command) {
	command.Flags().IntVarP(&downloadItag, "itag", "i", 0, "download video by itag")
	command.Flags().StringVarP(&downloadQuality, "quality", "q", "", "download video by quality")
	command.MarkFlagsMutuallyExclusive("itag", "quality", "audio", "muted")
}

func applyDownloadPlaylistFlags(command *cobra.Command) {
	command.Flags().BoolVarP(&downloadSimultaneously, "simultaneously", "s", false, "download simultaneously")
	command.MarkFlagsMutuallyExclusive("audio", "muted")
}
