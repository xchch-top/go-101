package v4

type Aggregate struct {
	arg string
	fn  string
}

func (Aggregate) selectable() {}

func Avg(col string) Aggregate {
	return Aggregate{
		arg: col,
		fn:  "AVG",
	}
}

func Max(col string) Aggregate {
	return Aggregate{
		arg: col,
		fn:  "MAX",
	}
}
