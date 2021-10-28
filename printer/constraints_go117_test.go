//go:build go1.17
// +build go1.17

package printer_test

import (
	"testing"

	"github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/printer"
)

func TestConstraints(t *testing.T) {
	ctx := build.NewContext()
	ctx.ConstraintExpr("linux,386 darwin,!cgo")
	ctx.ConstraintExpr("!noasm")

	AssertPrintsLines(t, ctx, printer.NewGoAsm, []string{
		"// Code generated by avo. DO NOT EDIT.",
		"",
		"//go:build ((linux && 386) || (darwin && !cgo)) && !noasm",
		"// +build linux,386 darwin,!cgo",
		"// +build !noasm",
		"",
	})
}
