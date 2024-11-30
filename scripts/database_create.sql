CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price INT NOT NULL
);

INSERT INTO products (name, price) VALUES
('Product 1', 100),
('Product 2', 200),
('Product 3', 300),
('Product 4', 400),
('Product 5', 500);
