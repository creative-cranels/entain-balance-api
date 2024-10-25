-- +goose Up
CREATE TYPE public.TransactionType AS ENUM ('ADDITION', 'SUBTRACTION');

create table public.transactions
(
	id serial not null
		constraint transactions_pk
			primary key,
  user_id integer references public.users (id),
	amount numeric not null,
  transaction_type  public.TransactionType not null,

	created_at timestamp default now() not null,
	updated_at timestamp default now() not null,
	deleted_at timestamp
);

-- +goose Down
DROP TABLE public.transactions;
DROP TYPE public.TransactionType;