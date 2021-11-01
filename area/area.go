// Package area para testar import de pacote externo
package area

import "math"

// Pi e uma proporcao numerica definida pela relacao entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ eh responsavel por calcular a area de uma circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect eh responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Nao eh visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

func Alegria() string {
	return "Alegria blah saco mano eita agora sim"
}
