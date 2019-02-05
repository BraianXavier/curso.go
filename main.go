package goarea

import "math"

// Pi é uma proporção numerca definida pela relação entre
// o perimetro de uma circunferência e seu diametro
const Pi = 3.1415

//Circ é responsabel por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é responsavel por calcular a área de um Triangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _TriagnguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
