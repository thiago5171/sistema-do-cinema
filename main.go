package main

import(
	"fmt"
	"cineProject/estruturadedados"
)
var customer []Movie

func main() {

	customer = append(customer,Movie{
		Name : "Harry Potter",	&Session{
			Room_l : "1",// 1 siginifica  que a sala foi selecionada
			Room_d : "0",// 0 significa que a sala nao foi selecionada
		},
	})


	customer = append(customer,Movie{
		Name : "",	&Session{
			Room_l : "0",
			Room_d : "1",
		},
	})

	customer[0].SeeMovie()
	fmt.Print("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")
	customer[1].SeeMovie()



}