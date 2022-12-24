package c26_v1

type E interface {
	M1()
	M2()
}

// type I interface {
// 	M1()
// 	M2()
// 	M3()
// }

// 接口类型的类型嵌入
// 下面的接口声明与第8行等价
type I interface {
	E
	M3()
}

// go官方包中的接口类型嵌入/src/io/io.go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadCloser interface {
	Reader
	Closer
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// go1.17之后没有下面代码没有编译错误
type Interface1 interface {
	M1()
}

type Interface2 interface {
	M1()
	M2()
}

type Interface3 interface {
	Interface1
	Interface2 // Error: duplicate method M1
}

type Interface4 interface {
	Interface2
	M2() // Error: duplicate method M2
}
