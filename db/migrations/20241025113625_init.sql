-- +goose Up
create table main.users
(
	id serial not null
		constraint users_pk
			primary key,
	balance numeric not null check (balance >= 0),

	created_at timestamp default now() not null,
	updated_at timestamp default now() not null,
	deleted_at timestamp
);

INSERT INTO main.users 
(id, balance, created_at, updated_at, deleted_at) 
VALUES 
(DEFAULT, 1000.20, DEFAULT, DEFAULT, null),
(DEFAULT, 1500, DEFAULT, DEFAULT, null),
(DEFAULT, 3000.520, DEFAULT, DEFAULT, null);

-- +goose Down
DROP TABLE main.users;