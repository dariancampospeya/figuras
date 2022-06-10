package figuras

import "fmt"

type geometrica interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func ImprimirGeometrica(geo geometrica) {
	fmt.Printf("El area es: %f\n", geo.calcularArea())
	fmt.Printf("El perímetro es: %f\n", geo.calcularPerimetro())
}
