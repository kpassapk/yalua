local re = require("yalo/regexp")

local r = re.compile("a.")

local got = r.replaceAll("-ab-axb-", -1)

print(got)
