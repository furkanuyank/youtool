package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/kkdai/youtube/v2/downloader"
)

func DownloadNotification(title string) {
	fmt.Printf("%v has been downloaded to %s\n", title, OUTDIR)
}

func SetOutdirInApp(newValue string) {
	OUTDIR = newValue
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func GetEnvLines(path string) []string {
	var readFile *os.File
	var err error
	readFile, err = os.Open(path)
	if err != nil {
		CreateEnvFile(path)
		readFile, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	return fileLines
}

func GetEnvValue(path string, key string) string {
	lines := GetEnvLines(path)
	var line = ""
	for _, v := range lines {
		if strings.Contains(v, key) {
			line = v
			break
		}
	}
	if line == "" {
		log.Fatal("env key not found")
	}
	values := strings.Split(line, "=")
	return values[1]
}

func SetEnvValue(path string, key string, value string) {
	lines := GetEnvLines(path)
	found := false
	for i, v := range lines {
		if strings.Contains(v, key) {
			lines[i] = key + "=" + value
			found = true
			break
		}
	}
	if !found {
		log.Fatal("cannot found")
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range lines {
		_, err := f.WriteString(v + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer f.Close()
}

func CreateEnvFile(path string) {
	file, err := os.Create(CONFIGENVPATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range DEFAULTVALUES {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getVideo(URL string) *youtube.Video {
	downloader := new(downloader.Downloader)
	downloader.ChunkSize = youtube.Size1Mb
	downloader.OutputDir = OUTDIR
	video, err := downloader.GetVideo(URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return video
}
func getVideoAndDownloader(URL string) (*youtube.Video, *downloader.Downloader) {
	downloader := new(downloader.Downloader)
	downloader.ChunkSize = youtube.Size1Mb
	downloader.OutputDir = OUTDIR
	video, err := downloader.GetVideo(URL)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	return video, downloader
}
func getPlaylist(URL string) *youtube.Playlist {
	downloader := new(downloader.Downloader)
	downloader.ChunkSize = youtube.Size1Mb
	downloader.OutputDir = OUTDIR
	playlist, err := downloader.GetPlaylist(URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return playlist
}
func getPlaylistAndDownloader(URL string) (*youtube.Playlist, *downloader.Downloader) {
	downloader := new(downloader.Downloader)
	downloader.ChunkSize = youtube.Size1Mb
	downloader.OutputDir = OUTDIR
	playlist, err := downloader.GetPlaylist(URL)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	return playlist, downloader
}
func getVideoIDFromPlaylist(URL string, index string) string {
	i, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal("convert error: ", err)
	}
	downloader := new(downloader.Downloader)
	downloader.ChunkSize = youtube.Size1Mb
	downloader.OutputDir = OUTDIR
	playlist, err := downloader.GetPlaylist(URL)
	if err != nil {
		log.Fatal("playlist error: ", err)
	}
	if i > len(playlist.Videos) || i <= 0 {
		log.Fatal("could not find a video with this number in the playlist")
	}
	id := playlist.Videos[i-1].ID
	return id
}

func IsVideo(URL string) bool {
	downloader := new(downloader.Downloader)
	downloader.ChunkSize = youtube.Size1Mb
	downloader.OutputDir = OUTDIR
	_, err := downloader.GetVideo(URL)
	if err != nil {
		return false
	}
	return true
}

func setTitle(name string, extension string, defaultName string, defaultExtension string) string {
	var videoName = defaultName
	var videoExt = defaultExtension
	if name != "" {
		videoName = name
	}
	if extension != "" {
		videoExt = extension
	}
	return videoName + "." + videoExt
}
