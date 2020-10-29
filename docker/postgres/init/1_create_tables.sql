create table routines (
    id serial primary key,
    date date not null unique,
    created_at timestamp not null,
    updated_at timestamp not null
);

create table tasks (
    id serial primary key,
    name varchar(255) not null,
    done boolean not null,
    routine_id integer not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    foreign key (routine_id) references routines(id)
);
