package v3

import (
	"context"
	"fmt"
	"gitlab.xchch.top/zhangsai/go-101/training/week07/orm_v3/v3/internal/errs"
	"gitlab.xchch.top/zhangsai/go-101/training/week07/orm_v3/v3/internal/model"
	"gitlab.xchch.top/zhangsai/go-101/training/week07/orm_v3/v3/internal/valuer"
	"strings"
)

type Selector[T any] struct {
	sb    strings.Builder
	tbl   string
	where []Predicate

	args  []any
	model *model.Model
	db    *DB
}

func (s *Selector[T]) Build() (*Query, error) {
	t := new(T)
	var err error
	s.model, err = s.db.r.Get(t)
	if err != nil {
		return nil, err
	}

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

func NewSelector[T any](db *DB) *Selector[T] {
	return &Selector[T]{
		db: db,
	}
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	q, err := s.Build()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.db.QueryContext(ctx, q.SQL, q.Args...)
	if err != nil {
		return nil, err
	}

	// 没有找到数据
	if !rows.Next() {
		return nil, ErrNoRows
	}

	t := new(T)
	val := valuer.NewReflectValue(t, s.model)
	return t, val.SetColumns(rows)
}

func (s *Selector[T]) GetMulti(ctx context.Context) ([]*T, error) {
	panic("implement me")
}
