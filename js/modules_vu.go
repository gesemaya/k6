package js

import (
	"context"

	"github.com/dop251/goja"
	"github.com/gesemaya/k6/event"
	"github.com/gesemaya/k6/js/common"
	"github.com/gesemaya/k6/js/eventloop"
	"github.com/gesemaya/k6/lib"
)

type events struct {
	global, local *event.System
}

type moduleVUImpl struct {
	ctx       context.Context
	initEnv   *common.InitEnvironment
	state     *lib.State
	runtime   *goja.Runtime
	eventLoop *eventloop.EventLoop
	events    events
}

func (m *moduleVUImpl) Context() context.Context {
	return m.ctx
}

func (m *moduleVUImpl) Events() common.Events {
	return common.Events{Global: m.events.global, Local: m.events.local}
}

func (m *moduleVUImpl) InitEnv() *common.InitEnvironment {
	return m.initEnv
}

func (m *moduleVUImpl) State() *lib.State {
	return m.state
}

func (m *moduleVUImpl) Runtime() *goja.Runtime {
	return m.runtime
}

func (m *moduleVUImpl) RegisterCallback() func(func() error) {
	return m.eventLoop.RegisterCallback()
}
