package main

import (
	"assignment2/service"
	"fmt"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	names := []string{"Irfan", "Aulia", "Giva", "Fahmi", "Yusuf", "Taslim", "Burok"}
	for _, n := range names {
		res := userSvc.Register(&service.User{Name: n})
		fmt.Println(res)
	}

	resGet := userSvc.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	for _, v := range resGet {
		cetakNama(v.Name)
	}
}

func cetakNama(nama string) {
	fmt.Println(nama)
}
