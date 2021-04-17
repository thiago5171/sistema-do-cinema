package cineProject
import(
	"fmt"
	
	)

type Movie struct{
	Name  		string
	Session 	*Session
}

func (x Movie)  SeeMovie(){
	fmt.Print("Filme->",x.Name)
	if(Room_l=="1"){
		fmt.Print("sessão legendada")
	}
	if(Room_d=="1"){
		fmt.Print("sessão dublada")
	}

}

