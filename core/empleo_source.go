package core

type EmpleoSource interface {
	Init() error
	Fetch() ([]Empleo, bool)
}
