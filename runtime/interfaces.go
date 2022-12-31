package runtime

type Setupper interface {
	Setup() error
}

type Service interface {
	Start() error
	Stop() error
}

type Logger interface {
	Info(log string)
	Error(log string, err error)
}
