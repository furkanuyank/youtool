package pkg

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kkdai/youtube/v2"
)

func PrintPlaylistContent(URL string,
	playlistTitle bool,
	playlistDescription bool,
	playlistAuthor bool,
	playlistVideos bool) {

	p := getPlaylist(URL)
	bold := color.New(color.Bold)

	if playlistTitle {
		bold.Print("Title: ")
		fmt.Println(p.Title, "\n")
	}
	if playlistDescription {
		bold.Print("Description: ")
		fmt.Println(p.Description, "\n")
	}
	if playlistAuthor {
		bold.Print("Author: ")
		fmt.Println(p.Author, "\n")
	}
	if playlistVideos {
		printVideosOfPlaylist(p)
	}
	if !playlistTitle && !playlistDescription && !playlistAuthor && !playlistVideos {

		printPlaylistAllContent(p, bold)
	}
}

func printPlaylistAllContent(p *youtube.Playlist, bold *color.Color) {
	bold.Print("Title: ")
	fmt.Println(p.Title, "\n")

	bold.Print("Description: ")
	fmt.Println(p.Description, "\n")

	bold.Print("Author: ")
	fmt.Println(p.Author, "\n")

	bold.Print("Videos: \n")
	printVideosOfPlaylist(p)

}

func printVideosOfPlaylist(p *youtube.Playlist) {
	for i, v := range p.Videos {
		fmt.Println(i+1, v.Title)
	}
}
