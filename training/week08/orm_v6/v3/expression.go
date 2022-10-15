package v3

type RawExpr struct {
	raw  string
	args []any
}

func (RawExpr) selectable() {}

func Raw(expr string, args ...any) RawExpr {
	return RawExpr{
		raw:  expr,
		args: args,
	}
}
