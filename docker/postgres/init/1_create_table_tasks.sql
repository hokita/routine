create table tasks (
    id serial primary key,
    name varchar(255) not null,
    done boolean not null,
    created_at timestamp not null,
    updated_at timestamp not null
);
