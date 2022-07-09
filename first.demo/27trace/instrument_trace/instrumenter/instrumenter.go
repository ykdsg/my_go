package instrumenter

type Instrumenter interface {
	Instrument(string2 string) ([]byte, error)
}
