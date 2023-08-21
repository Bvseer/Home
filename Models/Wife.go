package Models

const COST_OF_PRODUCTS = 10
const HAPPINESS_POINTS_FOR_BUYING_FURCOAT = 60

type Wife struct {
	Human Human
}

func (w Wife) eat(wife *Wife, points uint16) {
	wife.Human.Satiety += points
}

func (w Wife) buyProducts(wife *Wife) {
	wife.Human.Satiety -= COST_OF_PRODUCTS
}

func (w Wife) buyFurcoat(wife *Wife) {
	wife.Human.Satiety += HAPPINESS_POINTS_FOR_BUYING_FURCOAT
}

func (w Wife) cleanHome(home Home, points int16) {
	home.Dirt -= points
}
