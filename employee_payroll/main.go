// Ish haqini hisoblash uchun CalculateSalary() va DisplayDetails() kabi methodlardan foydalangan holda
// Employee deb nomlangan interfeysni aniqlang va
// to'liq kunlik, yarim kunlik va shartnoma xodimlari kabi har xil turdagi xodimlarning tafsilotlarini ko'rsating.
// Har bir xodim turi uchun interfeysni amalga oshiring.
// Shartnoma xodimlari oyiga 160 soatni ishashini farase qiling
package main

import "fmt"

func main() {
	Axrorbek := FullDayEmployee{
		Fullname:         "Axrorbek Olimjonov",
		SalaryPerHour:    4,
		MonthlyWorkHours: 1.5 * 3 * 22,
	}

	Dostonbek := HalfDayEmployee{
		Fullname:         "Dostonbek Soliyev",
		SalaryPerHour:    4,
		MonthlyWorkHours: 2 * 22,
	}

	Nodirjon := FullDayEmployee{
		Fullname:         "Nodirbek Numonov",
		SalaryPerHour:    4,
		MonthlyWorkHours: 8 * 22,
	}

	employees := []Employee{
		Axrorbek, Dostonbek, Nodirjon,
	}

	for _, employee := range employees {
		fmt.Println(DisplayDetails(employee))
	}

}

type Employee interface {
	GetSalary() float32
	GetDetails() string
}

func CalculateSalary(e Employee) float32 {
	return e.GetSalary()
}

func DisplayDetails(e Employee) string {
	return e.GetDetails()
}

type FullDayEmployee struct {
	Fullname         string
	SalaryPerHour    float32
	MonthlyWorkHours float32
}

type HalfDayEmployee struct {
	Fullname         string
	SalaryPerHour    float32
	MonthlyWorkHours float32
}

type ContractedEmployee struct {
	Fullname         string
	SalaryPerHour    float32
	MonthlyWorkHours float32
}

func (f FullDayEmployee) GetSalary() float32 {
	return f.MonthlyWorkHours * f.SalaryPerHour
}

func (h HalfDayEmployee) GetSalary() float32 {
	return h.MonthlyWorkHours * h.SalaryPerHour
}

func (c ContractedEmployee) GetSalary() float32 {
	return c.MonthlyWorkHours * c.SalaryPerHour
}

func (f FullDayEmployee) GetDetails() string {
	return fmt.Sprintf("Employee %s works %f hours monthly and gets $%f a month", f.Fullname, f.MonthlyWorkHours, CalculateSalary(f))
}

func (h HalfDayEmployee) GetDetails() string {
	return fmt.Sprintf("Employee %s works %f hours monthly and gets $%f a month", h.Fullname, h.MonthlyWorkHours, CalculateSalary(h))
}

func (c ContractedEmployee) GetDetails() string {
	return fmt.Sprintf("Employee %s works %f hours monthly and gets $%f a month", c.Fullname, c.MonthlyWorkHours, CalculateSalary(c))
}
