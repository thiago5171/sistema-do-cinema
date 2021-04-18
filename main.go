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
			tarde: "0",
			noite: "1",
		},
	})
	customer = append(customer, estruturas.Movie{
		Name: "Harry Potter", Sessao: &estruturas.Sessao{
			RoomL: "0",
			RoomD: "1",
			tarde: "1",
			noite: "0",
		},
	})


	customer = append(customer, estruturas.Movie{
		Name: " jogos mortais", Sessao: &estruturas.Sessao{
			RoomL: "0",
			RoomD: "1",// 1 siginifica  que a sala foi selecionada
			tarde: "0",
			noite: "1",
		},
	})
	

	customer = append(customer, estruturas.Movie{
		Name: " jogos mortais", Sessao: &estruturas.Sessao{
			RoomL: "1",
			RoomD: "0",
			tarde: "1",
			noite: "0",
		},
	})

	
	customer[0].SeeMovie()
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")
	customer[1].SeeMovie()
	//customer[2]=append(customer,*estruturas.NewMovie("anjos da lei","1","0"))

	
	for i:=0;i<=3;i++ {
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")
		customer[i].SeeMovie()	
	}
}
