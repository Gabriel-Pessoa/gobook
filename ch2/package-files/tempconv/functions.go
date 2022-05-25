package tempconv

// CtoF converte uma temperatura em Celsius para Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC converte uma temperatura em Fahrenheit para Celsius
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
