# Yalua - A Yalo library for Lua

## Usage

``` lua
local re = require("yalo/regexp")

local r = re.compile("a.")

local got = r.replaceAll("-ab-axb-", -1)

print(got)
-- prints "--1--1b-"
```


The code here is extracted and modified from the Shopify/goluago [regexp][goluago] library.

[goluago]: https://github.com/Shopify/goluago/blob/main/pkg/regexp/regexp.go
