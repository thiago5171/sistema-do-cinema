create DATABASE base_de_dados;
use base_de_dados;

create TABLE filme(
id int not null auto_increment,
nome Varchar(50),
data_lancamento date,
primary key(id)
);


 CREATE table sessao(
id int not null auto_increment,
idioma enum('D','L'),
horario datetime,
id_filme int not null,
primary key(id),
FOREIGN KEY (id_filme) REFERENCES filme(id)
 );
 
create table sala(
id int not null auto_increment,
numero_sala int,
primary key(id),
id_sessao int not null,
FOREIGN KEY (id_sessao) REFERENCES sessao(id)
);

create table ingresso(
id  int not null auto_increment primary key,
preco decimal(5,2),
id_sessao int not null,
FOREIGN KEY (id_sessao) REFERENCES sessao(id)
);

create table cliente(
id int not null auto_increment,
nome varchar(50),
primary key(id)
);

create table compra(
id int not null auto_increment primary key,
id_ingresso int not null,
id_cliente int not null
);

alter table compra add constraint compra_ingresso foreign key (id_ingresso) references ingresso(id);
alter table compra add constraint compra_cliente foreign key (id_cliente) references cliente(id);