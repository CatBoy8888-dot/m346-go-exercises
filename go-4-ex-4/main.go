package main

import "fmt"

// Neue Typen
type Celsius float64
type Fahrenheit float64

// Methode für Celsius → Fahrenheit
func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32)
}

// Methode für Fahrenheit → Celsius
func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5.0 / 9.0)
}

func main() {
	// Testdaten
	var cozy Celsius = 23.0
	fmt.Println(cozy.ConvertToFahrenheit()) // 73.4°F

	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius()) // -9.277777777777779°C

	// Zusatzfrage: Methoden-Ketten vs. Funktions-Ketten
	fmt.Println(Celsius(23).ConvertToFahrenheit().ConvertToCelsius())       // 23°C
	fmt.Println(convertCelsiusToFahrenheit(23))                             // 73.4
	fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(23))) // 23°C
}

func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5.0 / 9.0
}
