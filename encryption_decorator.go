package hls

type EncryptionDecorator struct {
	decorator  ConverterDecorator
	encKeyPath string
}

func (this *EncryptionDecorator) constructor(decorator ConverterDecorator, encKeyPath string) *EncryptionDecorator {
	this.decorator = decorator
	this.encKeyPath = encKeyPath
	return this
}

func NewEncryptionDecorator(decorator ConverterDecorator, encKeyPath string) *EncryptionDecorator {
	encryptionDecorator := EncryptionDecorator{}
	return encryptionDecorator.constructor(decorator, encKeyPath)
}

func (this *EncryptionDecorator) Decorate(cmd *command) *command {
	if this.decorator != nil {
		return this.decorator.Decorate(cmd)
	}
	decoratedArguments := append(cmd.Arguments(), "-hls_key_info_file", this.encKeyPath)
	return NewCommand(cmd.App(), decoratedArguments)
}
