package v0

type Assignable interface {
	assign()
}

type Assignment struct {
	column string
	val    any
}

func (Assignment) assign() {}

func Assign(column string, val any) Assignment {
	return Assignment{
		column: column,
		val:    val,
	}
}
