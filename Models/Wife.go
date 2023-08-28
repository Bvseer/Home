package Models

import "fmt"

const COST_OF_PRODUCTS = 10
const HAPPINESS_POINTS_FOR_BUYING_FURCOAT = 60
const COST_OF_FURCOAT = 350

type Wife struct {
	Human Human
}

func (w Wife) BuyProducts(home *Home) {
	home.Money -= COST_OF_PRODUCTS
}

func (w *Wife) BuyFurcoat(home *Home) {
	if home.Money < COST_OF_FURCOAT+100 {
		return
	}
	w.Human.Happiness += HAPPINESS_POINTS_FOR_BUYING_FURCOAT
	home.Money -= COST_OF_FURCOAT
	fmt.Println("FURCOAT HAS BEEN BOUGHT")
	fmt.Println("DIRT", home.Dirt)
}

func (w Wife) CleanHome(home *Home, points uint16) {
	home.Dirt -= points
}
