create table book (
	id serial primary key,
	name varchar(50) not null,
	author varchar(50) not null,
	price int not null
)

insert into book (name, author, price) values ('walden', 'henry david', 100000) returning *;

delete from book where id = 8 returning *

update book set name = 'reksadana lanjutan', author = 'andy', price = 45000 where id = 10 