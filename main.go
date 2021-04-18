package main

import(
	"fmt"
	estruturas "sistema-do-cinema/estruturadedados"
)

var customer []estruturas.Movie

func main() {
	customer = append(customer, estruturas.Movie{
		Name: "Harry Potter", Sessao: &estruturas.Sessao{
			Room_l: "1",// 1 siginifica  que a sala foi selecionada
			Room_d: "0",// 0 significa que a sala nao foi selecionada
		},
	})


	customer = append(customer, estruturas.Movie{
		Name: " jogos mortais", Sessao: &estruturas.Sessao{
			Room_l: "0",
			Room_d: "1",
		},
	})

	
	customer[0].SeeMovie()
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")
	customer[1].SeeMovie()
	customer[2]= estruturas.NewMovie("anjos da lei","1","0")


	for(int i:=0;i<= len(customer);i++){  
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=--=-")
	customer[i].SeeMovie()
	}
}