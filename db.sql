CREATE TABLE recipes (
	id serial PRIMARY KEY,
	name VARCHAR ( 100 ) UNIQUE NOT NULL,
	description VARCHAR ( 300 ),
	instuctions VARCHAR ( 2000 ) NOT NULL,
	created_on TIMESTAMP NOT NULL,
    prep_time INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    foreign key (category_id) references categories(id)
);


CREATE TABLE categories (
    id serial PRIMARY KEY,
    name VARCHAR ( 100 ) UNIQUE NOT NULL
);

CREATE TABLE ingredients (
    id serial PRIMARY KEY,
    name VARCHAR ( 100 ) UNIQUE NOT NULL,
    units VARCHAR ( 100 ) NOT NULL
);

CREATE TABLE recipe_ingredients (
    id serial PRIMARY KEY,
    recipe_id INTEGER NOT NULL,
    ingredient_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    foreign key (recipe_id) references recipes(id),
    foreign key (ingredient_id) references ingredients(id)
);

CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR ( 100 ) UNIQUE NOT NULL,
    password VARCHAR ( 100 ) NOT NULL,
    email VARCHAR ( 100 ) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL
);