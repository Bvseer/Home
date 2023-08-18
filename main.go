package main

import "awesomeProject/Models"

func main() {

	home := Models.Home{
		Money: 100,
		Food:  50,
		Dirt:  0,
	}

	husband := Models.Husband{
		Human: Models.Human{
			Name:      "Alex",
			Satiety:   30,
			Happiness: 100,
		},
	}

	wife := Models.Husband{
		Human: Models.Human{
			Name:      "Naomi",
			Satiety:   30,
			Happiness: 100,
		},
	}
}
