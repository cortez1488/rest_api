CREATE TABLE users(
                      id        INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT unique ,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
)

CREATE TABLE todo_lists
(
    id        INTEGER NOT NULL
        PRIMARY KEY AUTOINCREMENT unique ,
    title varchar(255) not null,
    description varchar(255),
)

CREATE TABLE users_lists
(
    id        INTEGER NOT NULL
        PRIMARY KEY AUTOINCREMENT unique ,
    user_id int references users (id) on delete cascade not null,
    list_id int references  todo_lists(id) on delete cascade not null,
)

CREATE TABLE todo_items
(
    id        INTEGER NOT NULL
        PRIMARY KEY AUTOINCREMENT unique ,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
)

CREATE TABLE lists_items
(
    id        INTEGER NOT NULL
        PRIMARY KEY AUTOINCREMENT unique ,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references  todo_lists(id) on delete cascade not null,
)