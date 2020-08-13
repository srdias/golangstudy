CREATE SCHEMA golang;

ALTER SCHEMA golang OWNER TO POSTGRES;

create table golang.pessoas(
	i_pessoas integer,
	nome character(64),
	primary key(i_pessoas)
);

commit;
