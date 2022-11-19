package v2

type Column struct {
	name string
}

func C(name string) Column {
	return Column{name: name}
}

func (Column) expr() {}

func (Column) selectable() {}

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

type Value struct {
	val any
}

func (Value) expr() {}
