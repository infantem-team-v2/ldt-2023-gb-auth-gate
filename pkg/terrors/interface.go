package terrors

type IError interface {
	LoggerMessage() string
	Wrap(err IError) IError
}
