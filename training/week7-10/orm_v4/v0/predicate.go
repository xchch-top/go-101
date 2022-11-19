package v0

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

type Column struct {
	name string
}

func C(name string) Column {
	return Column{name: name}
}

func (Column) expr() {}

// C("id").EQ(12)
func (c Column) EQ(val any) Predicate {
	return Predicate{
		left:  c,
		op:    opEQ,
		right: Value{val: val},
	}
}

func (c Column) GT(val any) Predicate {
	return Predicate{
		left:  c,
		op:    opGT,
		right: Value{val: val},
	}
}

func (c Column) LT(val any) Predicate {
	return Predicate{
		left:  c,
		op:    opLT,
		right: Value{val: val},
	}
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

type Value struct {
	val any
}

func (Value) expr() {}
