package v4

type Middleware func(next HandleFunc) HandleFunc
