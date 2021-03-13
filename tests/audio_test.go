package tests

import (
	"github.com/lootag/hls"
	"io/ioutil"
	"os"
	"testing"
)

func TestAudioIsConvertedCorrectly(t *testing.T) {
	//Arrange
	expectedNumberOfFiles := 4
	encKeyPath := "enc.keyinfo"
	audioPath := "mp3_sound.mp3"
	folder := "seq"
	converter := hls.NewBaseConverter(hls.NewEncryptionDecorator(nil, encKeyPath), folder)
	audio := hls.NewAudio(converter, audioPath)
	os.RemoveAll(folder)

	//Act
	audio.ToHLS()

	//Assert
	fileInfos, err := ioutil.ReadDir(folder)
	if err != nil {
		t.Errorf("Couldn't open " + folder + " directory")
	}
	if len(fileInfos) != expectedNumberOfFiles {
		t.Errorf("Expected " + string(expectedNumberOfFiles) + " but got " + string(len(fileInfos)))
	}

}
