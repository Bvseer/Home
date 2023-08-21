package main

import (
	"awesomeProject/Models"
	"fmt"
	"math/rand"
)

// Создать модель жизни небольшой семьи.
//
// Каждый день участники жизни могут делать только одно действие.
// Все вместе они должны прожить год и не умереть.
//
// Муж может:
//   есть,
//   играть в Комп,
//   ходить на работу,
// Жена может:
//   есть,
//   покупать продукты,
//   покупать шубу,
//   убираться в доме,
//
// Все они живут в одном доме, дом характеризуется:
//   кол-во денег в тумбочке (в начале - 100)
//   кол-во еды в холодильнике (в начале - 50)
//   кол-во грязи (в начале - 0)
//
// У людей есть имя, степень сытости (в начале - 30) и степень счастья (в начале - 100).
//
// Любое действие, кроме "есть", приводит к уменьшению степени сытости на 10 пунктов
// Кушают взрослые максимум по 30 единиц еды, степень сытости растет на 1 пункт за 1 пункт еды.
// Степень сытости не должна падать ниже 0, иначе чел умрет от голода.
//
// Деньги в тумбочку добавляет муж, после работы - 150 единиц за раз.
// Еда стоит 10 денег 10 единиц еды. Шуба стоит 350 единиц.
//
// Грязь добавляется каждый день по 5 пунктов, за одну уборку жена может убирать до 100 единиц грязи.
// Если в доме грязи больше 90 - у людей падает степень счастья каждый день на 10 пунктов,
// Степень счастья растет: у мужа от игры на Компьютере (на 20), у жены от покупки шубы (на 60, но шуба дорогая)
// Степень счастья не должна падать ниже 10, иначе чел умрает от депресии.
//
// Подвести итоги жизни за год: сколько было заработано денег, сколько сьедено еды, сколько куплено шуб.

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

	wife := Models.Wife{
		Human: Models.Human{
			Name:      "Naomi",
			Satiety:   30,
			Happiness: 100,
		},
	}

	for i := uint16(0); i < 365; i++ {

		addDirt(&home)

		if home.Dirt >= 90 {
			decreaseHappiness(&husband, &wife)
		}

		husbandAction := rand.Intn(3)
		if husbandAction == 0 {
			min := 20
			max := 30
			foodAmount := rand.Intn(max-min) + min
			fmt.Println(foodAmount)
			husband.Human.Eat(uint16(foodAmount))
		} else if husbandAction == 1 {
			husband.Play()
		} else {
			husband.Work(&home)
		}

		wifeAction := rand.Intn(2)
		if husbandAction == 0 {
			min := 20
			max := 30
			foodAmount := rand.Intn(max-min) + min
			fmt.Println(foodAmount)
			husband.Human.Eat(uint16(foodAmount))
		} else if husbandAction == 1 {
			husband.Play()
		} else {
			husband.Work(&home)
		}
	}
}

func addDirt(home *Models.Home) {
	home.Dirt += 5
}

func decreaseHappiness(husband *Models.Husband, wife *Models.Wife) {
	husband.Human.Happiness -= 10
	wife.Human.Happiness -= 10
}
