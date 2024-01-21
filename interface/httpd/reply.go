package httpd

type Reply interface {
	ToBytes() []byte
}
