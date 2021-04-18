package estruturadedados

import(
	"fmt"
)

type Movie struct{
	Name  		string
	Sessao 	    *Sessao
}

func (x Movie)  SeeMovie(){
	fmt.Println("Filme->", x.Name);

	if(x.Sessao.RoomL == "1"){
		fmt.Println("sessão legendada");

	}
	if(x.Sessao.RoomD =="1"){
		fmt.Println("sessão dublada");
	}

	if (x.Sessao.tarde=="1"){
		fmt.Println("Horario de inicio: 16:00");
		
	}
	if (x.Sessao.noite=="1"){
		fmt.Println("Horario de inicio: 20:00");
		
	}
}
/*
func NewMovie(Name string ,RoomD, RoomL string) *Movie{
	return &Movie{Name,&Sessao{RoomD, RoomL}}
}*/