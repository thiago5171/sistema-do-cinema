package cineProject
import(
	"fmt"
	
	)

type Movie struct{
	Name  		string
	Sessao 	    *Sessao
}

func (x Movie)  SeeMovie(){
	fmt.Println("Filme->",x.Name)  {
		if(Room_l == "1"){
			fmt.Println("sessão legendada")
		}
		if(Room_d =="1"){
			fmt.Println("sessão dublada")
		}

	}
	
}

