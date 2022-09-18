package v1

type Middleware func(next HandleFunc) HandleFunc
