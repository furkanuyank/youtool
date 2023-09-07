package pkg

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

func DownloadPlaylist(args []string, sim bool, name string, extension string) {
	if len(args) == 1 {
		if name != "" {
			log.Fatal("cannot used name flag if you download more than 1 video")
		}
		downloadPlaylistAll(args[0], sim, extension)

	} else if len(args) == 2 {
		downloadPlaylistSelected(args[0], args[1:], sim, name, extension)
	} else {
		if name != "" {
			log.Fatal("cannot used name flag if you download more than 1 video")
		}
		downloadPlaylistSelected(args[0], args[1:], sim, name, extension)
	}
}

func DownloadPlaylistAudio(args []string, sim bool, name string, extension string) {
	if len(args) == 1 {
		if name != "" {
			log.Fatal("cannot used name flag if you download more than 1 video")
		}
		downloadPlaylistAudioAll(args[0], sim, extension)

	} else if len(args) == 2 {
		downloadPlaylistAudioSelected(args[0], args[1:], sim, name, extension)
	} else {
		if name != "" {
			log.Fatal("cannot used name flag if you download more than 1 video")
		}
		downloadPlaylistAudioSelected(args[0], args[1:], sim, name, extension)
	}
}

func DownloadPlaylistMuted(args []string, sim bool, name string, extension string) {
	if len(args) == 1 {
		if name != "" {
			log.Fatal("cannot used name flag if you download more than 1 video")
		}
		downloadPlaylistMutedAll(args[0], sim, extension)

	} else if len(args) == 2 {
		downloadPlaylistMutedSelected(args[0], args[1:], sim, name, extension)
	} else {
		if name != "" {
			log.Fatal("cannot used name flag if you download more than 1 video")
		}
		downloadPlaylistMutedSelected(args[0], args[1:], sim, name, extension)
	}
}

func downloadPlaylistAll(URL string, sim bool, extension string) {
	playlist := getPlaylist(URL)

	if sim {
		var wg sync.WaitGroup
		wg.Add(len(playlist.Videos))
		for _, v := range playlist.Videos {
			videoId := v.ID
			go func(id string) {
				defer wg.Done()
				DownloadVideo(videoId, "", extension, false)
			}(videoId)
		}
		wg.Wait()
	} else {
		for _, v := range playlist.Videos {
			videoId := v.ID
			DownloadVideo(videoId, "", extension, true)
		}
	}
}

func downloadPlaylistAudioAll(URL string, sim bool, extension string) {
	playlist := getPlaylist(URL)

	if sim {
		var wg sync.WaitGroup
		wg.Add(len(playlist.Videos))
		for _, v := range playlist.Videos {
			videoId := v.ID
			go func(id string) {
				defer wg.Done()
				DownloadVideoAudio(videoId, "", extension, false)
			}(videoId)
		}
		wg.Wait()
	} else {
		for _, v := range playlist.Videos {
			videoId := v.ID
			DownloadVideoAudio(videoId, "", extension, true)
		}
	}
}

func downloadPlaylistMutedAll(URL string, sim bool, extension string) {
	playlist := getPlaylist(URL)

	if sim {
		var wg sync.WaitGroup
		wg.Add(len(playlist.Videos))
		for _, v := range playlist.Videos {
			videoId := v.ID
			go func(id string) {
				defer wg.Done()
				DownloadVideoMuted(videoId, "", extension, false)
			}(videoId)
		}
		wg.Wait()
	} else {
		for _, v := range playlist.Videos {
			videoId := v.ID
			DownloadVideoMuted(videoId, "", extension, true)
		}
	}
}

func downloadPlaylistSelected(URL string, numbers []string, sim bool, name string, extension string) {
	playlist := getPlaylist(URL)
	var validNumbers []int
	for _, v := range numbers {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v is not a number\n", v)
		} else if n > len(playlist.Videos) || n <= 0 {
			fmt.Printf("incorrect number\n")
		} else {
			validNumbers = append(validNumbers, n)
		}
	}
	if sim {
		var wg sync.WaitGroup
		wg.Add(len(validNumbers))
		for _, v := range validNumbers {
			videoId := playlist.Videos[v-1].ID
			go func() {
				DownloadVideo(videoId, "", extension, false)
				defer wg.Done()
			}()
		}
		wg.Wait()
	} else {
		for _, v := range validNumbers {
			videoId := playlist.Videos[v-1].ID
			DownloadVideo(videoId, "", extension, true)
		}
	}
}

func downloadPlaylistAudioSelected(URL string, numbers []string, sim bool, name string, extension string) {
	playlist := getPlaylist(URL)
	var validNumbers []int
	for _, v := range numbers {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v is not a number\n", v)
		} else if n > len(playlist.Videos) || n <= 0 {
			fmt.Printf("incorrect number\n")
		} else {
			validNumbers = append(validNumbers, n)
		}
	}
	if sim {
		var wg sync.WaitGroup
		wg.Add(len(validNumbers))
		for _, v := range validNumbers {
			videoId := playlist.Videos[v-1].ID
			go func() {
				DownloadVideo(videoId, "", extension, false)
				defer wg.Done()
			}()
		}
		wg.Wait()
	} else {
		for _, v := range validNumbers {
			videoId := playlist.Videos[v-1].ID
			DownloadVideo(videoId, "", extension, true)
		}
	}
}

func downloadPlaylistMutedSelected(URL string, numbers []string, sim bool, name string, extension string) {
	playlist := getPlaylist(URL)
	var validNumbers []int
	for _, v := range numbers {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v is not a number\n", v)
		} else if n > len(playlist.Videos) || n <= 0 {
			fmt.Printf("incorrect number\n")
		} else {
			validNumbers = append(validNumbers, n)
		}
	}
	if sim {
		var wg sync.WaitGroup
		wg.Add(len(validNumbers))
		for _, v := range validNumbers {
			videoId := playlist.Videos[v-1].ID
			go func() {
				DownloadVideo(videoId, "", extension, false)
				defer wg.Done()
			}()
		}
		wg.Wait()
	} else {
		for _, v := range validNumbers {
			videoId := playlist.Videos[v-1].ID
			DownloadVideo(videoId, "", extension, true)
		}
	}
}
