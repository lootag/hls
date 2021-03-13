package hls

type AudioConverter interface {
	MakeCommand(audio audio) *command
}
