package session

import (
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v9/v1"
)

type Manager struct {
	SessCtxKey string
	s          web.Store
	p          web.Propagator
}

func (m *Manager) InitSession(ctx *web.Context, id string) (web.Session, error) {
	sess, err := m.s.Generate(ctx.Req.Context(), id)
	if err != nil {
		return nil, err
	}
	if err = m.p.Inject(id, ctx.Resp); err != nil {
		return nil, err
	}
	return sess, nil
}

func (m *Manager) GetSession(ctx *web.Context) (web.Session, error) {
	// 先从UserValue中取
	if ctx.UserValues == nil {
		ctx.UserValues = make(map[string]any, 1)
	}
	sess, ok := ctx.UserValues[m.SessCtxKey]
	if ok {
		return sess.(web.Session), nil
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
	return sess.(web.Session), nil
}

// RefreshSession 刷新 Session
func (m *Manager) RefreshSession(ctx *web.Context) (web.Session, error) {
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

func (m *Manager) RemoveSession(ctx *web.Context) error {
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
