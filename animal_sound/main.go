package main

import (
	"fmt"
)

// Hayvon tomonidan chiqarilgan tovushni qaytaruvchi MakeSound() methodo yordamida
// Animal deb nomlangan interfeys yarating.
// It, mushuk va qush kabi turli hayvonlar uchun MakeSound() methodini qo'llang.
// Turli hayvonlarning tovush chiqarishi uchun interfeysdan foydalaning.
func main() {
	cat := Cat{
		name:  "cat",
		sound: "Mew mew",
	}

	dog := Dog{
		name:  "dog",
		sound: "Gof gof",
	}

	bird := Bird{
		name:  "bird",
		sound: "chrrr chrrr",
	}

	animals := []Animal{
		cat, dog, bird,
	}

	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}
}

type Cat struct {
	name  string
	sound string
}

type Dog struct {
	name  string
	sound string
}

type Bird struct {
	name  string
	sound string
}

func (c Cat) MakeSound() string {
	return "a " + c.name + " cat makes sound like " + c.sound
}

func (d Dog) MakeSound() string {
	return "a " + d.name + " dog makes sound like " + d.sound
}

func (b Bird) MakeSound() string {
	return "a " + b.name + " bird makes sound like " + b.sound
}

type Animal interface {
	MakeSound() string
}

func MakeSound(a Animal) string {
	return a.MakeSound()
}
