package Models

const SALARY = 150
const HAPPINESS_POINTS_FOR_GAMING = 20

type Husband struct {
	Human Human
}

func (h *Husband) Play() {
	h.Human.Happiness += HAPPINESS_POINTS_FOR_GAMING
}

func (h *Husband) Work(home *Home) {
	home.Money += SALARY
}
