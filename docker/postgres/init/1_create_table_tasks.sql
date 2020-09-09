create table tasks (
    id bigint primary key,
    name varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);
