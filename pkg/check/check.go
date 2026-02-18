package check

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "linter",
	Doc:  "notifies of usage uppercase",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		if !isLog(call.Fun) {
			return true
		}

		if len(call.Args) == 0 {
			return true
		}

		message := call.Args[0]

		lit, ok := message.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			return true
		}

		checkSensitiveData(pass, lit)

		messageText := lit.Value

		if messageText == "" {
			return true
		}

		if len(messageText) >= 2 {
			messageText = messageText[1 : len(messageText)-1]
		}

		checkCase(pass, lit, messageText)
		checkLanguageAndSymbols(pass, lit, messageText)

		return true
	}

	for _, file := range pass.Files {
		ast.Inspect(file, inspect)

	}

	return nil, nil
}

func isLog(expr ast.Expr) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	if id.Name == "log" || id.Name == "slog" {
		return true
	}

	return false
}

func checkCase(pass *analysis.Pass, lit *ast.BasicLit, messageText string) {
	if len(messageText) > 0 {
		first := []rune(messageText)[0]
		if unicode.IsUpper(first) {
			pass.Reportf(lit.Pos(), "the message must start with a lowercase letter")
		}
	}
}

func checkLanguageAndSymbols(pass *analysis.Pass, lit *ast.BasicLit, massageText string) {
	for _, char := range massageText {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || unicode.IsDigit(char) || char == ' ' {
			continue
		} else {
			pass.Reportf(lit.Pos(), "the message must be english letters without special symbols only")
			break
		}
	}
}

func checkSensitiveData(pass *analysis.Pass, lit *ast.BasicLit) {
	var sensitiveWords = map[string]bool{
		"password":     true,
		"pass":         true,
		"userPassword": true,
		"userPass":     true,
		"secret":       true,
		"auth":         true,
		"token":        true,
		"api_key":      true,
		"apikey":       true,
		"private_key":  true,
		"privateKey":   true,
	}

	message := strings.ToLower(strings.Trim(lit.Value, `"`))

	for word := range sensitiveWords {
		if strings.Contains(message, word+":") || strings.Contains(message, word+"=") {
			pass.Reportf(lit.Pos(), "the message contains sensitive data")
			return
		}
	}
}
