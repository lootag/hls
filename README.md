# HLS
This Golang library is a wrapper around ffmpeg, and provides an intuitive API to
convert media to HLS. It requires go modules.

# Example
```
package main

import (
    "github.com/lootag/hls"
)

func main() {
	encKeyPath := "enc.keyinfo" //your keyinfo file(if want encryption)
	audioPath := "path/to/media"
	folder := "path/to/folder" //The folder where the manifest and the fragments are
        //going to end up
	converter := hls.NewBaseConverter(hls.NewEncryptionDecorator(nil, encKeyPath), folder)
        //If you don't want the files to be encrypted just use
        //hls.NewBaseConverter(nil, folder)
	audio := hls.NewAudio(converter, audioPath)
	audio.ToHLS()
}
```
