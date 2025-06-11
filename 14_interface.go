package main

import "fmt"

// * Interfaces
type InfoI interface {
	getName() string
	totalPrice() float32
}

type TimeI interface {
	makingTime() int
}

// *Classes
type Food struct {
	foodName      string
	foodPrice     float32
	serviceCharge float32
	coockingTime  int
}

type Juice struct {
	name             string
	juicePrice       float32
	serviceCharge    float32
	numOfFruits      int
	fruitSqueezeTime int
}

type Car struct {
	brandName   string
	isAutomatic bool
	price       float32
	vat         float32
}

// * Methods
func (f Food) getName() string {
	return f.foodName
}
func (j Juice) getName() string {
	return j.name
}
func (c Car) getName() string {
	return c.brandName
}

// --
func (f Food) totalPrice() float32 {
	return f.foodPrice + f.serviceCharge
}
func (j Juice) totalPrice() float32 {
	return j.juicePrice + j.serviceCharge
}
func (c Car) totalPrice() float32 {
	return c.price + c.vat
}

// --
func (f Food) makingTime() int {
	return f.coockingTime
}
func (j Juice) makingTime() int {
	return j.fruitSqueezeTime * j.numOfFruits
}

// * Fucntion
func itemNameAndPriceIncVat(i InfoI) {
	fmt.Println(i.getName(), "- Cost:", i.totalPrice())
}

func itemsExpense(items []InfoI) ([]string, float32) {
	var names []string = make([]string, 0, 10)
	var totalCost float32 = 0
	for _, value := range items {
		names = append(names, value.getName())
		totalCost += value.totalPrice()
	}
	return names, totalCost
}

func restaurentWaitingTimeFor(item TimeI, nameInfo InfoI) {
	fmt.Println(item.makingTime(), "minute, for:", nameInfo.getName())
}

func main() {
	food1 := Food{
		foodName:      "Burger",
		foodPrice:     250,
		serviceCharge: 30,
		coockingTime:  20,
	}
	juice1 := Juice{
		name:             "OrangeMintLime",
		juicePrice:       60,
		serviceCharge:    20,
		numOfFruits:      3,
		fruitSqueezeTime: 5,
	}
	car1 := Car{
		brandName:   "Tesla",
		isAutomatic: true,
		price:       10000,
		vat:         2000,
	}
	//
	itemNameAndPriceIncVat(food1)
	itemNameAndPriceIncVat(juice1)
	itemNameAndPriceIncVat(car1)

	//
	name, totalExpese := itemsExpense([]InfoI{food1, juice1, car1})
	fmt.Println(name)
	fmt.Println(totalExpese)
	//

	restaurentWaitingTimeFor(food1, food1)
	restaurentWaitingTimeFor(juice1, juice1)
	// restaurentWaitingTimeFor(car1, car1)

	//* Type Assertions
	var data interface{} = "Mahadi Hassan"
	value := data.(string)
	fmt.Println(value)
	value1, ok1 := data.(string)
	fmt.Println(ok1, value1)
	value2, ok2 := data.(int)
	fmt.Println(ok2, value2)

	/*
		Burger - Cost: 280
		OrangeMintLime - Cost: 80
		Tesla - Cost: 12000
		[Burger OrangeMintLime Tesla]
		12360
		20 minute, for: Burger
		15 minute, for: OrangeMintLime
		Mahadi Hassan
		true Mahadi Hassan
		false 0
	*/

}
