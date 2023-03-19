-- 1a
INSERT INTO operator
(id, name, create_at, update_at)
VALUES(1, 'aul', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'lia', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'ret', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'no', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'ekawati', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--1b
INSERT INTO product_type
(id, name, create_at, update_at)
VALUES(1, 'jeans', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'jacket', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'sweater', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--1c
INSERT INTO product
(id, product_type_id, operator_id, code, create_at, update_at, status)
VALUES(1, 1, 3, 'jeans New', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
(2, 1, 3, 'shirt new', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);

--1d
INSERT INTO product
(id, product_type_id, operator_id, code, create_at, update_at, status)
VALUES(3, 2, 1, 'preloved jeans', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
(4, 2, 1, 'new jacket', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
(5, 2, 1, 'preloved sweater', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);

--1e
INSERT INTO product
(id, product_type_id, operator_id, code, create_at, update_at, status)
VALUES(6, 3, 4, 'new t shirt', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
(7, 3, 4, 'preloved jacket', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0),
(8, 3, 4, 'new sweater', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0);

--1f
INSERT INTO product_description
(id, product_id, description, create_at, update_at)
VALUES(1, 0, 'jeans New', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 1, 'shirt new', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 2, 'preloved jeans', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 2, 'new jacket', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 2, 'preloved sweater', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 3, 'new t shirt', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, 3, 'preloved jacket', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 3, 'new sweater', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--1g
INSERT INTO payment_method
(id, name, create_at, update_at)
VALUES(1, 'shopeepay', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'transfer bank', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'dana', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--1h
INSERT INTO `user`
(id, name, alamat_id, status_user, tanggal_lahir, gender, create_at, update_at)
VALUES(1, 'mark', 1, 0, '2000-01-02', 'perempuan', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'chenle', 2, 0, '2003-01-02', 'perempuan', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'jisung', 3, 0, '2004-02-02', 'perempuan', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'jeno', 4, 0, '2004-02-02', 'perempuan', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'jaemin', 5, 0, '2004-02-02', 'perempuan', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--1i
INSERT INTO `transaction`
(id, user_id, payment_method_id, quantity, price, status, create_at, update_at)
VALUES(1, 1, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 1, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 1, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 2, 2, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 2, 2, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 2, 2, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, 3, 3, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 3, 3, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(9, 3, 3, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(10, 4, 3, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(11, 4, 3, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(12, 4, 3, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(13, 5, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(14, 5, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(15, 5, 1, 0, 0, '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--1j

--2a
SELECT id, name, alamat_id, status_user, tanggal_lahir, gender, create_at, update_at
FROM `user`where gender = 'M';

--2b
SELECT * FROM product where id = 3;

--2c
SELECT *FROM `user` WHERE create_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';

--2d
SELECT count(*) FROM `user`WHERE gender = 'F';

--2e
SELECT * FROM `user` ORDER BY name asc ;

--3a
UPDATE product
SET product_type_id=0, operator_id=0, code='product dummy', create_at=CURRENT_TIMESTAMP, update_at=CURRENT_TIMESTAMP, status=0
WHERE id=1;

--3b
UPDATE transaction_details
SET quantity=3
WHERE id=1;

--4a
DELETE FROM product
WHERE id=1;

--4b
DELETE FROM product
WHERE product_type_id=1;

--UNION JOIN SUB QUERY FUNCTION
--1
SELECT * FROM transaction WHERE user_id = 1 OR user_id = 2;

--2
SELECT SUM(price) AS total_price FROM transaction WHERE user_id = 1;

--3
select count(*) from transaction t
join transaction_details td on t.id = td.transaction_id
join product p on p.id = td.product_id
where p.product_type_id = 2;

--4
SELECT product.*, product_type.name AS product_type_name FROM product INNER JOIN product_type ON product.product_type_id = product_type.id;

--5
SELECT transaction.*, quantity  AS quantity, price AS price 
FROM transaction
INNER JOIN product ON transaction.payment_method_id = product.id
INNER JOIN user ON transaction.user_id = user.id;

--6
delimiter $$
 create trigger delete_transaction_details
 before delete on transaction for each row
 begin
 declare trans_id int;
 set trans_id = old.id;
 delete from transaction_details where transaction_id = trans_id;
 end $$
 
 delete from transaction where id = 4;

--7
select * from product where id not in(
 select product_id from transaction_details
 );
 
select * from product where id in (
select p.id from product p
left outer join transaction_details td on p.id = td.product_id
where td.transaction_id is null
);

--8
SELECT *
FROM product
WHERE id NOT IN (
  SELECT product_id
  FROM transaction_details
);