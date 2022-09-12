package v0

type Middleware func(next HandleFunc) HandleFunc
