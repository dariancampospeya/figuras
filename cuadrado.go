package figuras

type Cuadrado struct {
	Base   float64
	Altura float64
}

func (cua *Cuadrado) calcularArea() float64 {
	return cua.Base * cua.Altura
}

func (cua *Cuadrado) calcularPerimetro() float64 {
	return 2*cua.Base + 2*cua.Altura
}
