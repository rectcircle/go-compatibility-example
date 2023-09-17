module github.com/rectcircle/go-compatibility-example/02-gotoolchain/02-goline-old-require-new

go 1.21.0

require github.com/rectcircle/go-compatibility-example/02-gotoolchain/01-goline-new v0.0.0

replace github.com/rectcircle/go-compatibility-example/02-gotoolchain/01-goline-new => ../01-goline-new
