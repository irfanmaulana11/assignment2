package main

import (
	"assignment2/service"
	"fmt"
	"sync"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	names := []string{"Irfan", "Aulia", "Giva", "Fahmi", "Yusuf", "Taslim", "Burok"}

	var wg1 sync.WaitGroup
	wg1.Add(len(names))

	for _, n := range names {
		go RegisterUser(n, userSvc, &wg1)
	}
	wg1.Wait()

	resGet := userSvc.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	var wg sync.WaitGroup
	wg.Add(len(resGet))
	for _, v := range resGet {
		go cetakNama(&wg, v.Name)
	}
	wg.Wait()
}

func cetakNama(wg *sync.WaitGroup, nama string) {
	fmt.Println(nama)
	wg.Done()
}

func RegisterUser(name string, svc service.UserInterface, wg *sync.WaitGroup) {
	res := svc.Register(&service.User{Name: name})
	fmt.Println(res)
	wg.Done()
}
