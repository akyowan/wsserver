package wsserver

import (
	"github.com/mailru/easygo/netpoll"
	"sync"
	"time"
)

type handleMessageFunc func(*Session, Object)
type handleSessionFunc func(*Session)
type handleErrorFunc func(*Session, error)

type Config struct {
	WriteWait time.Duration
	PongWait  time.Duration
}

type Hub struct {
	config            *Config
	sessions          sync.Map
	poller            *netpoll.Poller
	messageHandler    handleMessageFunc
	connectHandler    handleSessionFunc
	disconnectHandler handleSessionFunc
	pongHandler       handleSessionFunc
	closeHandler      handleSessionFunc
	errorHandler      handleErrorFunc
}
