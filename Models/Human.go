package Models

type Human struct {
	Name      string
	Satiety   uint16
	Happiness uint16
}

func (h *Human) Eat(points uint16) {
	h.Satiety += points
}
