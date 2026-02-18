package plugin

import (
	"linter/pkg/check"

	"github.com/golangci/golangci-lint/pkg/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func Linter() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"linter",
		"notifies of usage uppercase",
		[]*analysis.Analyzer{check.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
