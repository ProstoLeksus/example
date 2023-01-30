package main

import "fmt"

type Colour struct {
	Name string
	Code string
}
type Car struct {
	Colours []Colour
}

// func (car *Car) GetColour() string {
// 	return car.Colour.Name
// }

// func (car *Car) AddCar(colour string) Colour {
// 	car.Colour = colour
// 	return car.Colour
// }

// func (car *Car) DeleteColour() {
// 	car.Colour = Colour{}
// }

func main() {
	// var green Colour
	green := Colour{Name: "green", Code: "#000000"}
	blue := Colour{Name: "blue", Code: "#000333"}
	var colours []Colour
	colours = append(colours, green, blue)

	var car1 Car
	car1.Colours = colours
	fmt.Println(car1.Colours)
	// car1.DeleteColour()
	// fmt.Println(car1.Colour)
	// colour1 := colour1.DeleteColour()
	// var car1 Car
	// colour1 := colour1.AddColour("Blue")
	// var car1 Car
	// colour1 := colour1.GetColour()
	// fmt.Println(colour1)
}

/*
1 - Создать СТРУКТУРУ Map
2 - Добавить в нее поля для хранения ключей и значений
3 - Создать метод структуры Add - метод который будет добавлять структуру новые значения
4 - Создать метод структуры Get - метод который будет возвращать знначение по ключу
5 - Создать метод структуры Delete - метод который будет удалять из структуры по ключу
*/
//  6* - реализовать  сохранение данных любого типа/  почитать про приведение типов (UpCast DownCast)

// func main() {
// 	var student1 Student
// 	student1.Name = "Jony"
// 	student1.GroupNumber = 4

// 	var student2 Student
// 	student2.Name = "Vasya"
// 	student2.GroupNumber = 5

// 	var student3 Student
// 	student3.Name = "Nikita"
// 	student3.GroupNumber = 5

// 	name := student3.GetName()
// 	fmt.Println(name)

// 	id := uuid.New()
// 	idString := id.String()
// 	m := make(map[string]Student)
// 	m[idString] = student1
// 	m[uuid.New().String()] = student2
// 	m[uuid.New().String()] = student3

// 	for key, value := range m {
// 		if value.Name == "Nikita" {
// 			value.GroupNumber = 3
// 			fmt.Println(key, value)
// 		}
// 	}
// }
