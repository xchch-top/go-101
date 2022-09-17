package v2

type Middleware func(next HandleFunc) HandleFunc
