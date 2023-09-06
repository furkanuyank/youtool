package pkg

import (
	"strings"

	"github.com/fatih/color"
	"github.com/kkdai/youtube/v2"
	"github.com/rodaine/table"
)

var (
	headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt = color.New(color.FgYellow).SprintfFunc()
)

func ListAudioOnlyFormat(args []string) {
	tbl := table.New("Itag", "Type", "Quality", "FPS", "Bitrate", "Codec")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var video *youtube.Video
	if len(args) == 1 {
		video = getVideo(args[0])
	} else {
		videoId := getVideoIDFromPlaylist(args[0], args[1])
		video = getVideo(videoId)
	}
	printAudioOnlyFormat(video, tbl)
	tbl.Print()
}

func ListMutedOnlyFormat(args []string) {
	tbl := table.New("Itag", "Type", "Quality", "FPS", "Bitrate", "Codec")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var video *youtube.Video
	if len(args) == 1 {
		video = getVideo(args[0])
	} else {
		videoId := getVideoIDFromPlaylist(args[0], args[1])
		video = getVideo(videoId)
	}
	printMutedOnlyFormat(video, tbl)
	tbl.Print()
}

func ListVideoOnlyFormat(args []string) {
	tbl := table.New("Itag", "Type", "Quality", "FPS", "Bitrate", "Codec")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	var video *youtube.Video
	if len(args) == 1 {
		video = getVideo(args[0])
	} else {
		videoId := getVideoIDFromPlaylist(args[0], args[1])
		video = getVideo(videoId)
	}
	printVideoOnlyFormat(video, tbl)
	tbl.Print()
}

func ListAllFormats(args []string) {
	tbl := table.New("Itag", "Type", "Quality", "FPS", "Bitrate", "Codec")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var video *youtube.Video
	if len(args) == 1 {
		video = getVideo(args[0])
	} else {
		videoId := getVideoIDFromPlaylist(args[0], args[1])
		video = getVideo(videoId)
	}
	printAllFormats(video, tbl)
	tbl.Print()
}

func printAllFormats(video *youtube.Video, tbl table.Table) {
	printVideoOnlyFormat(video, tbl)
	printAudioOnlyFormat(video, tbl)
	printMutedOnlyFormat(video, tbl)
}
func printVideoOnlyFormat(video *youtube.Video, tbl table.Table) {
	for _, v := range video.Formats.WithAudioChannels().Type("video") {
		l := strings.Index(v.MimeType, ";")
		tbl.AddRow(v.ItagNo, "video", v.QualityLabel, v.FPS, v.Bitrate, v.MimeType[l+10:len(v.MimeType)-1])
	}
}
func printAudioOnlyFormat(video *youtube.Video, tbl table.Table) {
	for _, v := range video.Formats.WithAudioChannels().Type("audio") {
		l := strings.Index(v.MimeType, ";")
		tbl.AddRow(v.ItagNo, "audio", v.QualityLabel, v.FPS, v.Bitrate, v.MimeType[l+10:len(v.MimeType)-1])
	}
}
func printMutedOnlyFormat(video *youtube.Video, tbl table.Table) {
	for _, v := range video.Formats.AudioChannels(0).Type("video") {
		l := strings.Index(v.MimeType, ";")
		tbl.AddRow(v.ItagNo, "muted", v.QualityLabel, v.FPS, v.Bitrate, v.MimeType[l+10:len(v.MimeType)-1])
	}
}
