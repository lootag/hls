package hls

import (
	"os"
	"path"
)

func (this *baseConverter) finalizeCommand(command *command) *command {
	err := os.Mkdir(this.folder, 0700)
	if err != nil {
		panic(this.folder + " already exists!")
	}
	segmentNameFlag := "-hls_segment_filename"
	segmentNameValue := path.Join(this.folder, "sequence%d.ts")
	manifestName := path.Join(this.folder, this.folder+".m3u8")
	finalizedArgs := command.Arguments()
	finalizedArgs = append(finalizedArgs, segmentNameFlag, segmentNameValue, manifestName)
	return NewCommand(command.App(), finalizedArgs)
}
