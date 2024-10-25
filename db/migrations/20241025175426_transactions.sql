-- +goose Up
CREATE TYPE main.TransactionType AS ENUM ('ADDITION', 'SUBTRACTION');

create table main.transactions
(
	id serial not null
		constraint transactions_pk
			primary key,
  user_id integer references main.users (id),
	amount numeric not null,
  transaction_type  main.TransactionType not null,

	created_at timestamp default now() not null,
	updated_at timestamp default now() not null,
	deleted_at timestamp
);

-- +goose Down
DROP TABLE main.transactions;
DROP TYPE main.TransactionType;