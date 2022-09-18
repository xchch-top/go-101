package v3

type Middleware func(next HandleFunc) HandleFunc
