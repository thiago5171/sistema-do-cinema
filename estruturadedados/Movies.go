package estruturadedados

import(
	"fmt"
)

type Movie struct{
	Name  		string
	Sessao 	    *Sessao
}


/*
func NewMovie(Name string ,RoomD, RoomL string) *Movie{
	return &Movie{Name,&Sessao{RoomD, RoomL}}
}*/