package main

import (
	"github.com/Shopify/go-lua"
	"regexp"
)

// Open exposes the regexp functions to Lua code in the `goluago/regexp`
// namespace.
func Open(l *lua.State) {
	reOpen := func(l *lua.State) int {
		lua.NewLibrary(l, regexpLibrary)
		return 1
	}
	lua.Require(l, "yalo/regexp", reOpen, false)
	l.Pop(1)
}

var regexpLibrary = []lua.RegistryFunction{
	{"compile", compile},
}

func compile(l *lua.State) int {
	expr := lua.CheckString(l, 1)
	re, err := regexp.Compile(expr)
	if err != nil {
		lua.Errorf(l, err.Error())
		panic("unreachable")
	}

	l.NewTable()
	for name, goFn := range regexpFunc {
		// -1: tbl
		l.PushGoFunction(goFn(re))
		// -1: fn, -2:tbl
		l.SetField(-2, name)
	}

	return 1
}

var regexpFunc = map[string]func(*regexp.Regexp) lua.Function{
	"replaceAll": reReplaceAll,
}

// replaceAll
func reReplaceAll(re *regexp.Regexp) lua.Function {
	return func(l *lua.State) int {
		s := lua.CheckString(l, 1)
		r := lua.CheckString(l, 2)
		str := re.ReplaceAllString(s, r)
		l.PushString(str)
		return 1
	}
}

func main() {
	l := lua.NewState()
	Open(l)

	lua.OpenLibraries(l)
	if err := lua.DoFile(l, "hello.lua"); err != nil {
		panic(err)
	}
}
