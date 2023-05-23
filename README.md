# Yalua - A Yalo library for Lua

Lua does not have regular expressions. These need to be implemented as a library to use the host's capabilities. [goluago][glg] already implements a lot of the go standard library. It is missing a few methods. This is a demo for adding a `replaceAll` method from the go standard library to regex-replace occurrences of a regular expression in a string.

**Why does Lua not have regular expressions? Isn't that kind of basic?**

The entire Lua language is about [63KB][size] of code. A regular expression implementation would probably be bigger than the entire language!

## Usage

``` lua
local re = require("yalo/regexp")

local r = re.compile("a.")

local got = r.replaceAll("-ab-axb-", -1)

print(got)
-- prints "--1--1b-"
```


The code here is extracted and modified from the Shopify/goluago [regexp][goluago] library.

[size]: https://www.lua.org/notes/ltn002.html
[glg]: https://github.com/Shopify/goluago
[goluago]: https://github.com/Shopify/goluago/blob/main/pkg/regexp/regexp.go
