package runtime

type Service interface {
	Start() error
	Stop() error
}
