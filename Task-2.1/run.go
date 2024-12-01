package products

import (
	"fmt"
	"log"
)

type User struct {
	name string
	age  int
}
type Employee struct {
	User
	Title string
}

func (u *User) Get1() (string, int) {
	return u.GetName(), u.GetAge()
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) GetName() string {
	return u.name
}

func Run() {
	/*var num1 int = 10
	// Создает копию значения num1.
	var num2 int = num1
	// Изменяет только num2.
	num2 = 20

	fmt.Println(num1)
	fmt.Println(num2)

	// Типы данных по ссылке.
	var slice1 = []int{1, 2, 3}
	// Создает ссылку на slice1.
	var slice2 = slice1
	// Изменяет исходный срез slice1.
	slice2[0] = 4

	fmt.Println(slice1)
	fmt.Println(slice2)

	user1 := &User{name: "Ivan", age: 25}
	user2 := new(User)
	user2.name = "Petr"
	user2.age = 30
	fmt.Println(user1)
	fmt.Println(user2)

	fmt.Println(user2.Get1())

	product := new(Product)
	product.Name = "Misha"
	product.price = 30
	fmt.Println(product)

	emp := Employee{User: User{name: "Иван", age: 30}, Title: "Программист"}
	// Вывод: Иван 30 Программист.
	fmt.Println(emp.name, emp.age, emp.Title)

	// Пример использования.
	contractor := Contractor{
		User: User{
			name: "Ivan",
			age:  30,
		},
		ID: "C123",
	}

	employee := Employee1{
		User: User{
			name: "Stepan",
			age:  25,
		},
		ID: "4561",
	}

	// Проверка авторизации.
	if contractor.Authorizable() {
		fmt.Println("Правильный возраст")
	} else {
		fmt.Println("Ошибка возраста")
	}

	if employee.Authorizable() {
		fmt.Println("Правильный возраст")
	} else {
		fmt.Println("Ошибка возраста")
	}*/

	order := 0
	courier := CourierDelivery{}
	post := PostDelivery{}
	pickup := PickupDelivery{}

	fmt.Printf("Стоимость доставки курьером: %d\n", CalculateOrderShippingCost(courier, order))
	fmt.Printf("Стоимость доставки почтой: %d\n", CalculateOrderShippingCost(post, order))
	fmt.Printf("Стоимость самовывоза: %d\n", CalculateOrderShippingCost(pickup, order))

	sword := &Character{}
	sword.SetAttackStrategy(&SwordAttack{})

	bow := &Character{}
	bow.SetAttackStrategy(&BowAttack{})

	magic := &Character{}
	magic.SetAttackStrategy(&MagicAttack{})

	fmt.Printf("Персонаж атакует, урон мечом нанес: %d\n", sword.AttackProcessing())
	fmt.Printf("Персонаж атакует, урон луком нанесено: %d\n", bow.AttackProcessing())
	fmt.Printf("Персонаж атакует, урон магией нанесено: %d\n", magic.AttackProcessing())

	log1 := GetInstance()
	log2 := GetInstance()

	// Проверяем, что оба экземпляра одинаковые.
	if log1 == log2 {
		log.Println("один и тот же экземпляр.")
	} else {
		log.Println("разные")
	}
	log1.Log("fvddf")

	// Получение единственного экземпляра конфигурации.
	cfg := GetInstance()
	log.Println(cfg)
}
