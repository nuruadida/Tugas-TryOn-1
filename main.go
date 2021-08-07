package main

import (
	"fmt"
)

type education struct {
	School    string
	ThnMasuk  string
	ThnKeluar string
}
type teacher struct {
	Classroom   string
	TeacherName string
}
type person struct {
	Name      string
	Hobby     string
	Education []education
	Teacher   teacher
}

func main() {


	Person := person{
		Name:      "nurul hidayahtul",
		Hobby:     "bike",
		Education: []education{
			{
				School:    "SDN syalala",
				ThnMasuk:  "2003",
				ThnKeluar: "2009",
			},
			{
				School:    "Daarelqolam Boardingschool",
				ThnMasuk:  "2012",
				ThnKeluar: "2015",
			},
			{
				School:    "SMAN1 tangerang",
				ThnMasuk:  "2012",
				ThnKeluar: "2015",
			},
		},
		Teacher:   teacher{
			Classroom:   "12IPA",
			TeacherName: "Mr. Ali",
		},
	}


	fmt.Println(Person)
	fmt.Println("===========================")
	fmt.Printf("Nama \t: %v\n", Person.Name)
	fmt.Printf("Hobby \t: %v\n", Person.Hobby)
	fmt.Printf("Education\t: %v", Person.Education)
	fmt.Println("")
	fmt.Printf("Teacher \t: %v", Person.Teacher)
	fmt.Println("")
}
