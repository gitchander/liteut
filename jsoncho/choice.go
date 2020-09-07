package jsoncho

type Choice interface {
	GetType() (t string, err error)
	SetType(t string) error

	Value() interface{}
}
