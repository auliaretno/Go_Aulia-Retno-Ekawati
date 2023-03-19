-- No 1 Create database alta_online_shop.
CREATE SCHEMA `alta_online_shop` ;

-- No 2a
CREATE TABLE `alta_online_shop`.`user` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `alamat_id` INT NULL,
  `status_user` SMALLINT NULL,
  `tanggal_lahir` DATE NULL,
  `gender` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

-- No 2b
CREATE TABLE `alta_online_shop`.`product` (
  `id` INT NOT NULL,
  `product_type_id` INT NULL,
  `operator_id` INT NULL,
  `code` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `status` SMALLINT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `alta_online_shop`.`product_type` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

CREATE TABLE `alta_online_shop`.`operator` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

CREATE TABLE `alta_online_shop`.`product_description` (
  `id` INT NOT NULL,
  `product_id` INT NULL,
  `description` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

CREATE TABLE `alta_online_shop`.`payment_method` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

-- No 2c
CREATE TABLE `alta_online_shop`.`transaction` (
  `id` INT NOT NULL,
  `user_id` INT NULL,
  `payment_method_id` INT NULL,
  `quantity` INT NULL,
  `price` INT NULL,
  `status` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

CREATE TABLE `alta_online_shop`.`transaction_details` (
  `id` INT NOT NULL,
  `product_id` INT NULL,
  `transaction_id` INT NULL,
  `quantity` INT NULL,
  `price` INT NULL,
  `status` VARCHAR(45) NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`));

-- No 3
CREATE TABLE kurir (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- No 4
ALTER TABLE kurir ADD ongkos_dasar DECIMAL(10,2) NOT NULL DEFAULT 0;

-- No 5
ALTER TABLE kurir RENAME TO shipping;

-- No 6
-- No 6 hapus/drop TABLE

-- No 7a
CREATE TABLE payment_method_description (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    payment_method_id INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (payment_method_id) REFERENCES payment_method(id) ON DELETE CASCADE
);

-- No 7b
CREATE TABLE `user` (
  `id` int NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `alamat_id` int DEFAULT NULL,
  `status_user` smallint DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `gender` varchar(45) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (alamat_id) REFERENCES alamat(id) ON DELETE cascade
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- No 7c
CREATE TABLE `alta_online_shop`.`user_payment_method_detail` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NULL,
  `payment_method_id` INT NULL,
  `create_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
  FOREIGN KEY (payment_method_id) REFERENCES payment_method(id) ON DELETE CASCADE
 );