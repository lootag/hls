package hls

type ConverterDecorator interface {
	Decorate(command *command) *command
}
