package c24_v1

// type MyInt *int
//
// // receiver参数的基类型本身不能为指针类型
// // Invalid receiver type 'MyInt' ('MyInt' is a pointer type)
// func (r MyInt) String() string {
// 	return fmt.Sprintf("%d", *(*int)(r))
// }
//
// type MyReader io.Reader
//
// // receiver参数的基类型本身不能为接口类型
// // Invalid receiver type 'MyReader' ('MyReader' is an interface type)
// func (r MyReader) Read(p []byte) (int, error) {
// 	return r.Read(p)
// }
