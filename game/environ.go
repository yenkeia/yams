package game

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
)

var log = golog.New("server")

type Environ struct {
	Peer cellnet.GenericPeer
}

func NewEnviron() *Environ {
	return nil
}

func (env *Environ) Update() {
	log.Debugln("Update")
}

func (env *Environ) HandleEvent(ev cellnet.Event) {
	log.Debugln("HandleEvent")
}
