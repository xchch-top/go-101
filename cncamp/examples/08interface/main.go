package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

type Plane struct {
	vendor string
	model  string
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

func main() {
	// 定义一个接口类型的切片
	interfaces := []IF{}

	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	// 给切片中添加一个IF类型的对象
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		// 通过接口方法输出对象的getName方法的调用
		fmt.Println(f.getName())
	}

	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	// vendor: testVendor, model: testModel
	fmt.Println(p.getName())
}
