package controller

type Reply interface {
	ToBytes() []byte
}
