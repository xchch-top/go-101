package ast

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFileVisitor(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go",
		`
// annotation go through the source code and extra the annotation
// @author Deng Ming
// @date 2022/04/02
package annotation
`, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	fv := &FileVisitor{
		ans: map[string]string{},
	}
	ast.Walk(fv, f)

	res := map[string]string{
		"date":   "2022/04/02",
		"author": "Deng Ming",
	}
	assert.Equal(t, res, fv.ans)
}
