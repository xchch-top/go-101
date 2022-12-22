package c28_v2

// 接口越小, 抽象程度越高, 对应的事物集合越大

// 会飞的
type Flyable interface {
	Fly()
}

// 会游泳的
type Swimable interface {
	Swim()
}

// 会飞且会游泳的
type FlySwimable interface {
	Flyable
	Swimable
}
