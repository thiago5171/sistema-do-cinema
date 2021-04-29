package main

import(
	"fmt"
	estruturas "sistema-do-cinema/estruturadedados"
)

var customer []estruturas.Movie

func main() {
	customer = append(customer, estruturas.Movie{
		Name: "Harry Potter", Sessao: &estruturas.Sessao{
			RoomL: "1",// 1 siginifica  que a sala foi selecionada
			RoomD: "0",// 0 significa que a sala nao foi selecionada
			Tarde: "1",
			Noite: "0",
		},
	})
	customer = append(customer, estruturas.Movie{
		Name: "Harry Potter", Sessao: &estruturas.Sessao{
			RoomL: "0",
			RoomD: "1",// 1 siginifica  que a sala foi selecionada
			Tarde: "0",
			Noite: "1",
		},
	})


	customer = append(customer, estruturas.Movie{
		Name: " jogos mortais", Sessao: &estruturas.Sessao{
			RoomL: "1",
			RoomD: "0",
			Tarde: "1",
			Noite: "0",
		},
	})

	
	customer = append(customer, estruturas.Movie{
		Name: " jogos mortais", Sessao: &estruturas.Sessao{
			RoomL: "0",
			RoomD: "1",
			Tarde: "0",
			Noite: "1",
		},
	})
	

	
	//customer[2]=append(customer,*estruturas.NewMovie("anjos da lei","1","0"))

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")

	for i:=0;i<=3;i++ {
		
		customer[i].SeeMovie()	
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")

	}
}