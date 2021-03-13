package hls

type baseConverter struct {
	decorator ConverterDecorator
	folder    string
}

func (this *baseConverter) constructor(decorator ConverterDecorator, folder string) *baseConverter {
	this.decorator = decorator
	this.folder = folder
	return this
}

func NewBaseConverter(decorator ConverterDecorator, folder string) *baseConverter {
	baseConverter := baseConverter{}
	return baseConverter.constructor(decorator, folder)
}

func (this baseConverter) MakeCommand(audio audio) *command {
	app := "ffmpeg"
	arguments := []string{"-y", "-i", audio.Path(), "-hls_time", "10", "-hls_playlist_type", "vod"}
	cmd := NewCommand(app, arguments)
	decoratedCommand := cmd
	if this.decorator != nil {
		decoratedCommand = this.decorator.Decorate(cmd)
	}
	return this.finalizeCommand(decoratedCommand)
}
