package v1

// 这种叫做标记接口
// expr这个方法并不具有实际含义和实际用处
type Expression interface {
	expr()
}

type Predicate struct {
	left  Expression
	op    op
	right Expression
}

func (Predicate) expr() {}

type op string

const (
	opEQ = "="
	opLT = "<"
	opGT = ">"

	opNot = "NOT"
	opAnd = "AND"
	opOR  = "OR"
)

func (o op) String() string {
	return string(o)
}

// Not(C("id").EQ(12)) ==> not (id=?), 12
func Not(p Predicate) Predicate {
	return Predicate{
		op:    opNot,
		right: p,
	}
}

// C("id").EQ(12).And(C("name").EQ("Tom")) ==> not (id=)
func (p1 Predicate) And(p2 Predicate) Predicate {
	return Predicate{
		left:  p1,
		op:    opAnd,
		right: p2,
	}
}

func (p1 Predicate) Or(p2 Predicate) Predicate {
	return Predicate{
		left:  p1,
		op:    opOR,
		right: p2,
	}
}
