package session

import (
	"gitlab.xchch.top/golang-group/go-101/training/week4-6/web_v9/v1"
)

type Manager struct {
	SessCtxKey string
	s          v1.Store
	p          v1.Propagator
}

func (m *Manager) InitSession(ctx *v1.Context, id string) (v1.Session, error) {
	sess, err := m.s.Generate(ctx.Req.Context(), id)
	if err != nil {
		return nil, err
	}
	if err = m.p.Inject(id, ctx.Resp); err != nil {
		return nil, err
	}
	return sess, nil
}

func (m *Manager) GetSession(ctx *v1.Context) (v1.Session, error) {
	// 先从UserValue中取
	if ctx.UserValues == nil {
		ctx.UserValues = make(map[string]any, 1)
	}
	sess, ok := ctx.UserValues[m.SessCtxKey]
	if ok {
		return sess.(v1.Session), nil
	}

	id, err := m.p.Extract(ctx.Req)
	if err != nil {
		return nil, err
	}
	sess, err = m.s.Get(ctx.Req.Context(), id)
	if err != nil {
		return nil, err
	}
	// 存储到UserValues
	ctx.UserValues[m.SessCtxKey] = sess
	return sess.(v1.Session), nil
}

// RefreshSession 刷新 Session
func (m *Manager) RefreshSession(ctx *v1.Context) (v1.Session, error) {
	sess, err := m.GetSession(ctx)
	if err != nil {
		return nil, err
	}
	// 刷新存储的过期时间
	err = m.s.Refresh(ctx.Req.Context(), sess.ID())
	if err != nil {
		return nil, err
	}
	// 重新注入 HTTP 里面
	if err = m.p.Inject(sess.ID(), ctx.Resp); err != nil {
		return nil, err
	}
	return sess, nil
}

func (m *Manager) RemoveSession(ctx *v1.Context) error {
	sess, err := m.GetSession(ctx)
	if err != nil {
		return err
	}
	err = m.s.Remove(ctx.Req.Context(), sess.ID())
	if err != nil {
		return err
	}
	return m.p.Remove(ctx.Resp)
}
