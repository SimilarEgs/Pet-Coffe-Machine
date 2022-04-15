package main

import "fmt"

type Machine struct {
	water  int
	milk   int
	beans  int
	matcha int
	money  int
	cups   int
}

func main() {

	m := Machine{
		money: 550,
		water: 400,
		milk:  540,
		beans: 120,
		cups:  9,
	}

	var action string
	var coffe_type, quant int

	for action != "exit" {
		fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
		fmt.Scanln(&action)

		switch action {
		case "buy":
			fmt.Println("\nWhat do you want to buy? \n1 - espresso, 2 - latte, 3 - cappuccino, 4 - matcha:")
			fmt.Scanln(&coffe_type)
			fmt.Println("How many cups would u buy?")
			fmt.Scanln(&quant)
			m.buy(coffe_type, quant)
		case "fill":
			m.fill()
		case "take":
			if m.money >= 1 {
				fmt.Printf("\nI gave you $%d\n", m.money)
				m.money -= m.money
			} else {
				fmt.Println("\nInvalid: out of money")
			}
		case "remaining":
			m.info()
		default:
			fmt.Println("\nNot an option, try again.")
		}

	}

}

func (m Machine) info() {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d ml of water\n", m.water)
	fmt.Printf("%d ml of milk\n", m.milk)
	fmt.Printf("%d g of coffee beans\n", m.beans)
	fmt.Printf("%d g of matcha\n", m.matcha)
	fmt.Printf("%d disposable cups\n", m.cups)
	fmt.Printf("$%d of money\n", m.money)
}

func (m *Machine) buy(sell, quant int) {

	espresso_water, espresso_beans, espresso_money := 250, 16, 4
	espresso_water, espresso_beans, espresso_money = espresso_water*quant, espresso_beans*quant, espresso_money*quant

	latte_water, latte_milk, latte_beans, latte_money := 350, 75, 20, 7
	latte_water, latte_milk, latte_beans, latte_money = latte_water*quant, latte_milk*quant, latte_beans*quant, latte_money*quant

	cappucchino_water, cappucchino_milk, cappucchino_beans, cappucchino_money := 200, 100, 12, 6
	cappucchino_water, cappucchino_milk, cappucchino_beans, cappucchino_money = cappucchino_water*quant, cappucchino_milk*quant, cappucchino_beans*quant, cappucchino_money*quant

	matcha_water, matcha_milk, matcha_beans, matcha_powder, matcha_money := 50, 150, 5, 2, 12
	matcha_water, matcha_milk, matcha_beans, matcha_powder, matcha_money = matcha_water*quant, matcha_milk*quant, matcha_beans*quant, matcha_powder*quant, matcha_money*quant

	switch sell {
	case 1:
		if m.cups >= quant && m.money >= espresso_money && m.water >= espresso_water && m.beans >= espresso_beans {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.beans, m.cups = m.money+espresso_money, m.water-espresso_water, m.beans-espresso_beans, m.cups-quant
		} else if quant == 1 && m.money >= 4 && m.water >= 250 && m.beans >= 16 {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.beans, m.cups = m.money+4, m.water-250, m.beans-16, m.cups-1
		} else if m.money < espresso_money || m.water < espresso_water || m.beans < espresso_beans || m.cups < quant {
			switch {
			case m.water < espresso_water:
				fmt.Println("Sorry, not enough water")
			case m.beans < espresso_beans:
				fmt.Println("Sorry, not enough coffee beans!")
			case m.cups < quant:
				fmt.Println("Sorry, not eough cups")
			case m.money < espresso_money:
				fmt.Println("Sorry, not enough money")
			}
		} else {
			switch {
			case m.water < 250:
				fmt.Println("Sorry, not enough water")
			case m.beans < 16:
				fmt.Println("Sorry, not enough coffee beans")
			case m.cups < 1:
				fmt.Println("Sorry, not eough cups")
			case m.money < 4:
				fmt.Println("Not enough money")
			}
		}
	case 2:
		if m.cups >= quant && m.money >= latte_money && m.water >= latte_water && m.beans >= latte_beans && m.milk >= latte_milk {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.beans, m.milk, m.cups = m.money+latte_money, m.water-latte_water, m.beans-latte_beans, m.milk-latte_milk, m.cups-quant
		} else if quant == 1 && m.money >= 7 && m.water >= 350 && m.milk >= 75 && m.beans >= 20 {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.milk, m.beans, m.cups = m.money+7, m.water-350, m.milk-75, m.beans-20, m.cups-1
		} else if m.money < latte_money || m.water < latte_water || m.beans < latte_beans || m.milk < latte_milk || m.cups < quant {
			switch {
			case m.water < latte_water:
				fmt.Println("Sorry, not enough water")
			case m.beans < latte_beans:
				fmt.Println("Sorry, not enough coffee beans")
			case m.milk < latte_milk:
				fmt.Println("Sorry, not enough milk")
			case m.cups < quant:
				fmt.Println("Sorry, not eough cups")
			case m.money < latte_money:
				fmt.Println("Sorry, not enough money")
			}
		} else {
			switch {
			case m.water < 350:
				fmt.Println("Sorry, not enough water")
			case m.beans < 20:
				fmt.Println("Sorry, not enough coffee beans")
			case m.milk < 75:
				fmt.Println("Sorry, not enough milk")
			case m.milk < 75:
				fmt.Println("Sorry, not enough milk")
			case m.cups < 1:
				fmt.Println("Sorry, not eough cups")
			case m.money < 7:
				fmt.Println("Not enough money")
			}
		}
	case 3:
		if m.cups >= quant && m.money >= cappucchino_money && m.water >= cappucchino_water && m.beans >= cappucchino_beans && m.milk >= cappucchino_milk {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.beans, m.milk, m.cups = m.money+cappucchino_money, m.water-cappucchino_water, m.beans-cappucchino_beans, m.milk-cappucchino_milk, m.cups-quant
		} else if quant == 1 && m.money >= 6 && m.water >= 200 && m.milk >= 100 && m.beans >= 12 {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.milk, m.beans, m.cups = m.money+6, m.water-200, m.milk-100, m.beans-12, m.cups-1
		} else if m.money < cappucchino_money || m.water < cappucchino_water || m.beans < cappucchino_beans || m.milk < cappucchino_milk || m.cups < quant {
			switch {
			case m.water < cappucchino_water:
				fmt.Println("Sorry, not enough water")
			case m.beans < cappucchino_beans:
				fmt.Println("Sorry, not enough coffee beans!")
			case m.milk < cappucchino_milk:
				fmt.Println("Sorry, not enough milk")
			case m.cups < quant:
				fmt.Println("Sorry, not eough cups")
			case m.money < cappucchino_money:
				fmt.Println("Sorry, not enough money")
			}
		} else {
			switch {
			case m.water < 200:
				fmt.Println("Sorry, not enough water")
			case m.beans < 12:
				fmt.Println("Sorry, not enough coffee beans")
			case m.milk < 100:
				fmt.Println("Sorry, not enough milk")
			case m.cups < 1:
				fmt.Println("Sorry, not eough cups")
			case m.money < 6:
				fmt.Println("Not enough money")
			}
		}
	case 4:
		if m.cups >= quant && m.money >= matcha_money && m.water >= matcha_water && m.beans >= matcha_beans && m.matcha >= matcha_powder && m.milk >= matcha_milk {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.beans, m.matcha, m.milk, m.cups = m.money+matcha_money, m.water-matcha_water, m.beans-matcha_beans, m.matcha-matcha_powder, m.milk-matcha_milk, m.cups-quant
		} else if quant == 1 && m.money >= 12 && m.water >= 50 && m.milk >= 150 && m.beans >= 5 && m.matcha >= 2 {
			fmt.Println("\nI have enough resources, making you a coffee!")
			m.money, m.water, m.milk, m.beans, m.matcha, m.cups = m.money+12, m.water-50, m.milk-150, m.beans-5, m.matcha-2, m.cups-1
		} else if m.money < matcha_money || m.water < matcha_water || m.beans < matcha_beans || m.matcha < matcha_powder || m.milk < matcha_milk || m.cups < quant {
			switch {
			case m.water < cappucchino_water:
				fmt.Println("Sorry, not enough water")
			case m.beans < cappucchino_beans:
				fmt.Println("Sorry, not enough coffee beans")
			case m.matcha < matcha_powder:
				fmt.Println("Sorry, not enough matcha powder")
			case m.milk < cappucchino_milk:
				fmt.Println("Sorry, not enough milk")
			case m.cups < quant:
				fmt.Println("Sorry, not eough cups")
			case m.money < cappucchino_money:
				fmt.Println("Sorry, not enough money")
			}
		} else {
			switch {
			case m.water < 200:
				fmt.Println("Sorry, not enough water")
			case m.beans < 12:
				fmt.Println("Sorry, not enough coffee beans")
			case m.matcha < 2:
				fmt.Println("Sorry, not enough matcha powder")
			case m.milk < 100:
				fmt.Println("Sorry, not enough milk")
			case m.cups < 1:
				fmt.Println("Sorry, not eough cups")
			case m.money < 6:
				fmt.Println("Not enough money")
			}
		}
	}
}

func (m *Machine) fill() {
	var (
		add_water  int
		add_milk   int
		add_beans  int
		add_matcha int
		add_cups   int
	)
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scanln(&add_water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scanln(&add_milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scanln(&add_beans)
	fmt.Println("Write how many grams of matcha powder you want to add:")
	fmt.Scanln(&add_matcha)
	fmt.Println("Write how many disposable cups of coffee you want to add:")
	fmt.Scanln(&add_cups)
	m.water += add_water
	m.milk += add_milk
	m.beans += add_beans
	m.matcha += add_matcha
	m.cups += add_cups
}
