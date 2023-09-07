package pkg

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func DownloadVideo(URL string, name string, extension string, notification bool) {
	video, downloader := getVideoAndDownloader(URL)
	title := setTitle(name, extension, video.Title, "mp4")
	if notification {
		defer DownloadNotification(title)
	}
	err := downloader.Download(context.Background(), video, &video.Formats.WithAudioChannels().Type("video")[0], title)
	if err != nil {
		log.Fatal("Download error: ", err)
	}

}

func DownloadVideoWithQuality(URL string, q string, name string, extension string, notification bool) {
	video, downloader := getVideoAndDownloader(URL)
	format, err := getVideoFormatWithQuality(*video, q)
	if err != nil {
		log.Fatal("Format error: ", err)
	}
	title := setTitle(name, extension, video.Title, "mp4")
	if notification {
		defer DownloadNotification(title)
	}
	err = downloader.Download(context.Background(), video, format, title)
	if err != nil {
		log.Fatal("Download error: ", err)
	}
}

func DownloadVideoWithItag(URL string, i int, name string, extension string, notification bool) {
	video, downloader := getVideoAndDownloader(URL)
	format, err := getFormatWithItag(*video, i)
	if err != nil {
		log.Fatal("Format error: ", err)
	}
	l := strings.Index(format.MimeType, ";")
	title := setTitle(name, extension, video.Title, format.MimeType[6:l])
	if notification {
		defer DownloadNotification(title)
	}
	err = downloader.Download(context.Background(), video, format, title)
	if err != nil {
		log.Fatal("Download error: ", err)
	}
}

func DownloadVideoAudio(URL string, name string, extension string, notification bool) {
	video, downloader := getVideoAndDownloader(URL)
	format, err := getAudioFormat(*video)
	if err != nil {
		log.Fatalf("Format error: ", err)
	}
	title := setTitle(name, extension, video.Title, "mp3")
	if notification {
		defer DownloadNotification(title)
	}
	err = downloader.Download(context.Background(), video, format, title)
	if err != nil {
		log.Fatal("Download error: ", err)
	}
}

func DownloadVideoMuted(URL string, name string, extension string, notification bool) {
	video, downloader := getVideoAndDownloader(URL)
	format, err := getMutedVideoFormat(*video)
	if err != nil {
		log.Fatalf("Format error: ", err)
	}
	title := setTitle(name, extension, video.Title, "mp4")
	if notification {
		defer DownloadNotification(title)
	}
	err = downloader.Download(context.Background(), video, format, title)
	if err != nil {
		log.Fatal("Download error: ", err)
	}
}

func getVideoFormatWithQuality(video youtube.Video, q string) (*youtube.Format, error) {
	for _, v := range video.Formats.WithAudioChannels().Type("video") {
		if q == v.QualityLabel || q+"p" == v.QualityLabel {
			return &v, nil
		}
	}
	return nil, errors.New("no format")
}
func getFormatWithItag(video youtube.Video, i int) (*youtube.Format, error) {
	format := video.Formats.FindByItag(i)
	if format == nil {
		return nil, errors.New("no format")
	}
	return format, nil
}
func getAudioFormat(video youtube.Video) (*youtube.Format, error) {
	formats := video.Formats.WithAudioChannels().Type("audio")
	if formats == nil {
		return nil, errors.New("no format for audio")
	} else {
		return &formats[0], nil
	}
}
func getMutedVideoFormat(video youtube.Video) (*youtube.Format, error) {
	formats := video.Formats.AudioChannels(0).Type("video")
	if formats == nil {
		return nil, errors.New("no format for audio")
	} else {
		return &formats[0], nil
	}
}
