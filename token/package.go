// this file was generated by gomacro command: import "github.com/cosmos72/gomacro/token"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package token

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/token"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/token"] = imports.Package{
		Binds: map[string]r.Value{
			"IsKeyword":      r.ValueOf(IsKeyword),
			"IsLiteral":      r.ValueOf(IsLiteral),
			"IsMacroKeyword": r.ValueOf(IsMacroKeyword),
			"IsOperator":     r.ValueOf(IsOperator),
			"Lookup":         r.ValueOf(Lookup),
			"MACRO":          r.ValueOf(MACRO),
			"QUASIQUOTE":     r.ValueOf(QUASIQUOTE),
			"QUOTE":          r.ValueOf(QUOTE),
			"SPLICE":         r.ValueOf(SPLICE),
			"String":         r.ValueOf(String),
			"UNQUOTE":        r.ValueOf(UNQUOTE),
			"UNQUOTE_SPLICE": r.ValueOf(UNQUOTE_SPLICE),
		},
		Types:   map[string]r.Type{},
		Proxies: map[string]r.Type{}}
}
