package ast

import (
	"go/ast"
	"strings"
)

type FileVisitor struct {
	ans map[string]string
}

func (f *FileVisitor) Visit(node ast.Node) (w ast.Visitor) {
	fn, ok := node.(*ast.File)
	if !ok {
		return f
	}
	if fn.Doc == nil || len(fn.Doc.List) == 0 {
		return f
	}

	for _, doc := range fn.Doc.List {
		if !strings.HasPrefix(doc.Text, "// @") {
			continue
		}
		text := strings.TrimPrefix(doc.Text, "// @")
		if text == "" {
			continue
		}

		segs := strings.SplitN(text, " ", 2)
		if len(segs) == 0 {
			continue
		}
		key := segs[0]
		val := ""
		if len(segs) > 1 {
			val = segs[1]
		}
		f.ans[key] = val
	}

	return f
}
