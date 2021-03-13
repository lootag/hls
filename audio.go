package hls

type audio struct {
	converter AudioConverter
	path      string
}

func (this *audio) constructor(converter AudioConverter, path string) *audio {
	this.converter = converter
	this.path = path
	return this
}

func NewAudio(converter AudioConverter, path string) *audio {
	audio := &audio{}
	return audio.constructor(converter, path)
}

func (this *audio) Path() string {
	return this.path
}

func (this audio) ToHLS() {
	cmd := this.converter.MakeCommand(this)
	cmd.Execute()
}
