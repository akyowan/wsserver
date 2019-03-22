package wsserver

import (
	"io"
	"sync"
)

type Object interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type Session struct {
	io   sync.Mutex
	conn io.ReadWriteCloser

	params  sync.Map
	sendBuf chan Object
}

func (*Session) receive() error {
	return nil
}

func (*Session) writer() {
}

func (s *Session) Set(key, value interface{}) {
	s.params.Store(key, value)
}

func (s *Session) Get(key interface{}) (val interface{}, exists bool) {
	return s.params.Load(key)
}

func (*Session) Send(obj Object) error {
	return nil
}

func (*Session) Close() error {
	return nil
}
