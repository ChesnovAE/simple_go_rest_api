CREATE TABLE users (
    id  serial not null unique,
    name VARCHAR(255) not null,
    username VARCHAR(255) not null unique,
    password_hash VARCHAR(255) not null
);

CREATE TABLE todo_lists (
    id  serial not null unique,
    title VARCHAR(255) not null,
    description VARCHAR(255) not null
);

CREATE TABLE users_list (
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items (
    id  serial not null unique,
    title VARCHAR(255) not null,
    description VARCHAR(255) not null,
    done boolean not null default false
);

CREATE TABLE list_items (
    id serial not null unique,
    list_id int references todo_lists (id) on delete cascade not null,
    item_id int references todo_items (id) on delete cascade not null
);