CREATE TABLE users(
                      id        INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT unique ,
                      name varchar(255) not null,
                      username varchar(255) not null unique,
                      password_hash varchar(255) not null
)