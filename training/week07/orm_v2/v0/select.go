package v0

import (
	"context"
	"log"
	"reflect"
	"strings"
)

type Selector[T any] struct {
	sb    strings.Builder
	tbl   string
	where []Predicate

	args []any
}

func (s *Selector[T]) Build() (*Query, error) {
	s.sb.WriteString("SELECT * FROM ")

	// 拼接表名部分
	s.tableNameFragment()
	// 拼接where部分
	s.whereFragment()

	s.sb.WriteByte(';')
	return &Query{
		SQL:  s.sb.String(),
		Args: s.args,
	}, nil
}

func (s *Selector[T]) From(tbl string) *Selector[T] {
	s.tbl = tbl
	return s
}

func (s *Selector[T]) Where(ps ...Predicate) *Selector[T] {
	s.where = ps
	return s
}

// tableNameFragment 拼接表名部分的sql片段
func (s *Selector[T]) tableNameFragment() {
	if s.tbl == "" {
		var t T
		// 用结构体名字转下划线
		typ := reflect.TypeOf(t)
		goName := typ.Name()
		s.sb.WriteString("`" + goName + "`")
	} else {
		segs := strings.SplitN(s.tbl, ".", 2)
		if len(segs) == 1 {
			s.sb.WriteString("`" + segs[0] + "`")
		} else {
			s.sb.WriteString("`" + segs[0] + "`.`" + segs[1] + "`")
		}
	}
}

// whereFragment 拼接where部分的sql片段
func (s *Selector[T]) whereFragment() {
	if len(s.where) == 0 {
		return
	}

	s.sb.WriteString(" WHERE ")
	pred := s.where[0]
	for i := 1; i < len(s.where); i++ {
		pred = pred.And(s.where[i])
	}
	s.buildExpression(pred)
}

func (s *Selector[T]) buildExpression(expression Expression) {
	switch expr := expression.(type) {
	case nil:
		return
	case Column:
		s.sb.WriteByte('`')
		s.sb.WriteString(expr.name)
		s.sb.WriteByte('`')
	case Value:
		s.sb.WriteByte('?')
		s.args = append(s.args, expr.val)
	case Predicate:
		_, ok := expr.left.(Predicate)
		if ok {
			s.sb.WriteByte('(')
		}
		s.buildExpression(expr.left)
		if ok {
			s.sb.WriteByte(')')
		}

		s.sb.WriteString(" " + expr.op.String() + " ")

		_, ok = expr.right.(Predicate)
		if ok {
			s.sb.WriteByte('(')
		}
		s.buildExpression(expr.right)
		if ok {
			s.sb.WriteByte(')')
		}
	default:
		log.Println("不支持的表达式: ", expression)
	}
}

func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Selector[T]) GetMulti(ctx context.Context) ([]*T, error) {
	//TODO implement me
	panic("implement me")
}
