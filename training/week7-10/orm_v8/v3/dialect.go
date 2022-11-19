package v3

import (
	"gitlab.xchch.top/zhangsai/go-101/training/week7-10/orm_v8/v3/internal/errs"
)

// Dialect 方言, 构造个性部分
type Dialect interface {
	// 引号
	quoter() byte

	//
	buildDuplicateKey(b *builder, odk *Upsert) error
}

// standardSQL 标准的SQL实现
type standardSQL struct {
}

type mysqlDialect struct {
	standardSQL
}

func (dial *mysqlDialect) quoter() byte {
	return '`'
}

func (dial *mysqlDialect) buildDuplicateKey(b *builder, odk *Upsert) error {
	b.sb.WriteString(" ON DUPLICATE KEY UPDATE ")
	for idx, assign := range odk.assigns {
		if idx > 0 {
			b.sb.WriteByte(',')
		}
		switch expr := assign.(type) {
		case Assignment:
			fd, ok := b.model.FieldMap[expr.column]
			if !ok {
				return errs.NewErrUnknownField(expr.column)
			}
			b.quote(fd.ColName)
			b.sb.WriteString("=?")

			b.args = append(b.args, expr.val)
		case Column:
			fd, ok := b.model.FieldMap[expr.name]
			if !ok {
				return errs.NewErrUnknownField(expr.name)
			}
			b.quote(fd.ColName)
			b.sb.WriteString("=VALUES(")
			b.quote(fd.ColName)
			b.sb.WriteByte(')')
		}
	}
	return nil
}

type sqliteDialect struct {
	standardSQL
}

func (dial *sqliteDialect) buildDuplicateKey(b *builder, odk *Upsert) error {
	b.sb.WriteString(" ON CONFLICT")
	if len(odk.conflictColumns) > 0 {
		b.sb.WriteByte('(')
		for i, col := range odk.conflictColumns {
			if i > 0 {
				b.sb.WriteByte(',')
			}
			fd, ok := b.model.FieldMap[col]
			if !ok {
				return errs.NewErrUnknownField(col)
			}
			b.quote(fd.ColName)
		}
		b.sb.WriteByte(')')
	}

	b.sb.WriteString(" DO UPDATE SET ")
	for idx, assign := range odk.assigns {
		if idx > 0 {
			b.sb.WriteByte(',')
		}
		switch expr := assign.(type) {
		case Assignment:
			fd, ok := b.model.FieldMap[expr.column]
			if !ok {
				return errs.NewErrUnknownField(expr.column)
			}
			b.quote(fd.ColName)
			b.sb.WriteString("=?")

			b.args = append(b.args, expr.val)
		case Column:
			fd, ok := b.model.FieldMap[expr.name]
			if !ok {
				return errs.NewErrUnknownField(expr.name)
			}
			b.quote(fd.ColName)
			b.sb.WriteString("=excluded.")
			b.quote(fd.ColName)
		}
	}
	return nil
}
