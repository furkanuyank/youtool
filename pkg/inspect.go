package pkg

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kkdai/youtube/v2"
)

func PrintContent(args []string,
	title bool,
	description bool,
	author bool,
	length bool,
	views bool,
	thumbnails bool,
) {
	var video *youtube.Video
	if len(args) == 1 {
		video = getVideo(args[0])
	} else {
		videoId := getVideoIDFromPlaylist(args[0], args[1])
		video = getVideo(videoId)
	}
	printContentWithFilter(video,
		title,
		description,
		author,
		length,
		views,
		thumbnails,
	)
}

func printContentWithFilter(v *youtube.Video, title bool, description bool, author bool, length bool, views bool, thumbnails bool) {
	bold := color.New(color.Bold)

	if title {
		bold.Print("Title: ")
		fmt.Println(v.Title, "\n")
	}
	if description {
		bold.Print("Description: ")
		fmt.Println(v.Description, "\n")
	}
	if author {
		bold.Print("Author: ")
		fmt.Println(v.Author, "\n")
	}
	if length {
		bold.Print("Length: ")
		fmt.Println(v.Duration, "\n")
	}
	if views {
		bold.Print("Views: ")
		fmt.Println(v.Views, "\n")
	}
	if thumbnails {
		bold.Print("Video Pictures: \n")
		printVideoThumbnails(v)
	}
	if !title && !description && !author && !length && !views && !thumbnails {
		printAllContent(v, bold)
	}
}
func printAllContent(v *youtube.Video, bold *color.Color) {
	bold.Print("Title: ")
	fmt.Println(v.Title, "\n")

	bold.Println("Description:")
	fmt.Println(v.Description, "\n")

	bold.Print("Author: ")
	fmt.Println(v.Author, "\n")

	bold.Print("Video Duration: ")
	fmt.Println(v.Duration, "\n")

	bold.Print("Views: ")
	fmt.Println(v.Views, "\n")

	bold.Print("Video Pictures: \n")
	printVideoThumbnails(v)
}
func printVideoThumbnails(v *youtube.Video) {
	thumbnails := v.Thumbnails

	for _, v := range thumbnails {
		fmt.Printf("%vx%v -> %v\n", v.Height, v.Width, v.URL)
	}
}
