create table if not exists signals (
    id serial primary key,
    signal varchar(255) not null,
    inserted_at timestamp not null default current_timestamp
);