package example

var a = b + c // inicializada em terceiro lugar
var b = f()   // inicializada em segundo lugar
var c = 1     // inicializada em primeiro lugar

func f() int {
	return c + 1
}
