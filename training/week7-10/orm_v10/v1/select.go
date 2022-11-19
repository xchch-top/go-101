package v1

import (
	"context"
	"fmt"
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v10/v1/internal/errs"
	"strings"
)

type Selector[T any] struct {
	builder
	core
	sess Session

	tbl     string
	where   []Predicate
	columns []Selectable
}

type Selectable interface {
	selectable()
}

// s.Select("id", "age")
func (s *Selector[T]) Select(cols ...Selectable) *Selector[T] {
	s.columns = cols
	return s
}

func (s *Selector[T]) Build() (*Query, error) {
	t := new(T)
	var err error
	s.model, err = s.r.Get(t)
	if err != nil {
		return nil, err
	}

	s.sb.WriteString("SELECT ")
	s.ColumnsFragment()
	s.sb.WriteString(" FROM ")

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
		s.sb.WriteString("`" + s.model.TableName + "`")
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
	err := s.buildExpression(pred)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Selector[T]) buildExpression(expression Expression) error {
	switch expr := expression.(type) {
	case nil:
		return nil
	case Column:
		s.sb.WriteByte('`')
		fd, ok := s.model.FieldMap[expr.name]
		if !ok {
			return errs.NewErrUnknownField(expr.name)
		}
		s.sb.WriteString(fd.ColName)
		s.sb.WriteByte('`')
	case Value:
		s.sb.WriteByte('?')
		s.args = append(s.args, expr.val)
	case Predicate:
		_, ok := expr.left.(Predicate)
		if ok {
			s.sb.WriteByte('(')
		}
		err := s.buildExpression(expr.left)
		if err != nil {
			return err
		}
		if ok {
			s.sb.WriteByte(')')
		}

		s.sb.WriteString(" " + expr.op.String() + " ")

		_, ok = expr.right.(Predicate)
		if ok {
			s.sb.WriteByte('(')
		}
		err = s.buildExpression(expr.right)
		if err != nil {
			return err
		}
		if ok {
			s.sb.WriteByte(')')
		}
	default:
		return errs.ErrUnsupportedExpressionType
	}
	return nil
}

func NewSelector[T any](sess Session) *Selector[T] {
	return &Selector[T]{
		sess: sess,
		core: sess.getCore(),
	}
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	res := get[T](ctx, s.core, s.sess, &QueryContext{
		Type:    "SELECT",
		Builder: s,
	})
	if res.Result != nil {
		return res.Result.(*T), res.Err
	}
	return nil, res.Err

}

func (s *Selector[T]) GetMulti(ctx context.Context) ([]*T, error) {
	panic("implement me")
}

func (s *Selector[T]) ColumnsFragment() {
	if len(s.columns) == 0 {
		s.sb.WriteByte('*')
	} else {
		for i, c := range s.columns {
			if i > 0 {
				s.sb.WriteByte(',')
			}
			switch col := c.(type) {
			case Column:
				fd, ok := s.model.FieldMap[col.name]
				if !ok {
					fmt.Println(errs.NewErrUnknownField(col.name).Error())
				}
				s.sb.WriteString("`" + fd.ColName + "`")
			case Aggregate:
				fd, ok := s.model.FieldMap[col.arg]
				if !ok {
					fmt.Println(errs.NewErrUnknownField(col.arg).Error())
				}
				s.sb.WriteString(col.fn + "(`" + fd.ColName + "`)")
			case RawExpr:
				s.sb.WriteString(col.raw)
				if len(col.args) > 0 {
					s.args = append(s.args, col.args...)
				}
			}
		}
	}
}
