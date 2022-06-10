package figuras

import "math"

type Ciruclo struct {
	Radio float64
}

func (cir *Ciruclo) calcularArea() float64 {
	return math.Pi * (cir.Radio * cir.Radio)
}

func (cir *Ciruclo) calcularPerimetro() float64 {
	return 2 * math.Pi * cir.Radio
}
