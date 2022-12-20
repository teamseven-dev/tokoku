USE tokoku;

CREATE TABLE staffs (
	id_staff INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT NULL,
	password VARCHAR(50) DEFAULT NULL,
	created_date TIMESTAMP DEFAULT now(),
	updated_date TIMESTAMP DEFAULT now()
);

CREATE TABLE customers (
	id_customer INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT NULL,
	id_staff INT NOT NULL,
	CONSTRAINT fk_customers_staff FOREIGN KEY (id_staff) REFERENCES staffs(id_staff)
);

CREATE TABLE products (
	id_product INT AUTO_INCREMENT PRIMARY KEY,
	product_name VARCHAR(50) DEFAULT NULL,
	qty INT DEFAULT NULL,
	id_staff INT NOT NULL,
	created_date TIMESTAMP DEFAULT now(),
	updated_date TIMESTAMP DEFAULT now(),
	CONSTRAINT fk_products_staff FOREIGN KEY (id_staff) REFERENCES staffs(id_staff)
);

CREATE TABLE transactions (
	id_transaction INT AUTO_INCREMENT PRIMARY KEY,
	id_staff INT NOT NULL,
	id_customer INT NOT NULL,
	created_date TIMESTAMP DEFAULT now(),
	CONSTRAINT fk_transactions_staff FOREIGN KEY (id_staff) REFERENCES staffs(id_staff),
	CONSTRAINT fk_transactions_customer FOREIGN KEY (id_customer) REFERENCES customers(id_customer)
);

CREATE TABLE items (
	id_transaction INT NOT NULL,
	id_product INT NOT NULL,
	qty INT DEFAULT NULL,
	PRIMARY KEY(id_transaction, id_product),
	CONSTRAINT fk_items_transaction FOREIGN KEY (id_transaction) REFERENCES transactions(id_transaction),
	CONSTRAINT fk_items_product FOREIGN KEY (id_product) REFERENCES products(id_product)
);

