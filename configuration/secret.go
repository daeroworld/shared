package configuration

type Variable struct {
	Database       *Database
	VoiceWriterApi *Api
	VoiceReaderApi *Api
	TextWriterApi  *Api
	TextReaderApi  *Api
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
		TextReaderApi: &Api{
			Port: 25014,
		},
		Frontend: &Api{
			Port: 25001,
		},
	}
}
