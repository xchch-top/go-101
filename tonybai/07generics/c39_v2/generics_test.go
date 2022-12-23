package c39_v2

type Set[T comparable] map[T]struct{}

type sliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

type node[K, V any] struct{}

type Map[K, V any] struct {
	root    *node[K, V]
	compare func(K, K) int
}

type element[T any] struct {
	next *element[T]
	val  T
}

type P[T1, T2 any] struct {
	// 类型参数的顺序要与声明中类型参数列表中的顺序一致
	F *P[T1, T2] // ok
}

// type P2[T1, T2 any] struct {
// 	F *P2[T2, T1] // 不符合技术方案，但Go 1.19编译器并未报错
// }

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

type NumericAbs[T Numeric] interface {
	Abs() T
}

// 泛型方法
func (P[T1, T2]) max() {}
