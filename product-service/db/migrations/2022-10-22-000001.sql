CREATE DATABASE product;

CREATE table products (
    id SERIAL,
    name text,
    price numeric,
    CONSTRAINT product_pk PRIMARY KEY(id)
);
