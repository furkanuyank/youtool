
* [Youtool](#youtool)
* [Getting Started](#getting-started)
* [How to use](#how-to-use)


# Youtool
Youtool is a tool to download video or audio in the format and extension you want.

# Getting Started
For Linux
```
wget https://github.com/furkanuyank/youtool/releases/download/v1/youtool_Linux_x86_64.tar.gz
sudo tar -C /usr/bin -xzf youtool_Linux_x86_64.tar.gz && rm youtool_Linux_x86_64.tar.gz
```

# How to use

## Config

<h5>youtool config outdir [path] --> set default out directory (it is your home directory if you didn't change it)</h5>
<br>

## Download
<h5>youtool download [Video Url] --> download video</h5>
<h5>youtool download [Playlist Url] --> download playlist </h5>
<h5>youtool download [Playist Url] [Video number] ... --> download specific videos from playlist</h5>

-a --> download audio <br>
-m --> muted audio <br>
-n --> set name of the video (default is title of the video) <br>
-. --> set extension of video (default is mp4 for video, mp3 for audio) <br>
-o --> set out directory at once for downloaded video <br>

youtool download --help for more
<br><br>

## Inspect

<h5>youtool inspect [Video Url] or [Playist Url] [Video number] --> gives information about video such as title, author, views, thumbnails, etc.</h5>
<h5>youtool inspect [Playist Url] --> gives information about playlist such as title, author, videos contained in, etc.</h5>
