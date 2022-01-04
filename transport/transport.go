package transport

type Transport interface {
	drawBoard()
	clearBoard()
	next() bool
	Run()
}
