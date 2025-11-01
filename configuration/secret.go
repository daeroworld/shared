package configuration

type Variable struct {
	Database       *Database
	VoiceWriterApi *Api
	VoiceReaderApi *Api
	TextWriterApi  *Api
	Frontend       *Api
}

func NewVariable() *Variable {
	return &Variable{

		VoiceWriterApi: &Api{
			Port: 25011,
		},
		VoiceReaderApi: &Api{
			Port: 25012,
		},
		TextWriterApi: &Api{
			Port: 25013,
		},
	}
}
