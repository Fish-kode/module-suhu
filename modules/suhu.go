package suhu

// Fahrenheit
func Fahrenheit(c float32) float32 {

	f := (9/5)*c + 32

	return f
}

// Kelvin
// K = c + 273.15
func Kelvin(c float32) float32 {
	k := c + 273.15

	return k
}

// Rankine
// R = (c 	+ 173.15) * 9.5

func Rankie(c float32) float32 {
	rnk := (c + 273.15) * 9.5

	return rnk
}
