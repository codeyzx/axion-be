-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 19, 2023 at 08:58 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `axion`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `delete_product` (IN `p_product_id` VARCHAR(255), IN `p_user_id` INT)   BEGIN DELETE FROM products WHERE products.id = (SELECT auctions.product_id FROM auctions WHERE auctions.product_id = p_product_id AND auctions.user_id = p_user_id); END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `update_product` (IN `p_product_id` VARCHAR(255), IN `p_user_id` INT, IN `p_name` VARCHAR(255), IN `p_description` TEXT, IN `p_price` DECIMAL(10,2), IN `p_image` VARCHAR(255))   BEGIN UPDATE products SET name = p_name, description = p_description, price = p_price, image = p_image WHERE products.id = (SELECT auctions.product_id FROM auctions WHERE auctions.product_id = p_product_id AND auctions.user_id = p_user_id); END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `auctions`
--

CREATE TABLE `auctions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `last_price` bigint(20) DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `bidders_count` bigint(20) DEFAULT NULL,
  `product_id` varchar(191) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `end_at` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `auction_histories`
--

CREATE TABLE `auction_histories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `auction_id` bigint(20) UNSIGNED DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Triggers `auction_histories`
--
DELIMITER $$
CREATE TRIGGER `trigger_auction` BEFORE INSERT ON `auction_histories` FOR EACH ROW UPDATE `auctions` SET last_price = new.price, user_id = new.user_id, bidders_count = bidders_count + 1	WHERE id = new.auction_id AND new.price > last_price AND status = 'open'
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `histories`
--

CREATE TABLE `histories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `log` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` varchar(191) NOT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Triggers `products`
--
DELIMITER $$
CREATE TRIGGER `trigger_products_delete_history` AFTER DELETE ON `products` FOR EACH ROW INSERT INTO histories VALUES(NULL, CONCAT('Product with id ', OLD.id, ' (', OLD.name, ')', ' has been removed'),NOW(), NOW(), NULL)
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `trigger_products_insert_history` AFTER INSERT ON `products` FOR EACH ROW INSERT INTO histories VALUES(NULL, CONCAT('Product with id ', NEW.id, ' (', NEW.name, ')', ' has been added'),NOW(), NOW(), NULL)
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `trigger_products_update_history` AFTER UPDATE ON `products` FOR EACH ROW INSERT INTO histories VALUES(NULL, CONCAT('Product with id ', OLD.id, ' (', OLD.name, ')', ' has been updated'),NOW(), NOW(), NULL)
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `role` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auctions`
--
ALTER TABLE `auctions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `product_id` (`product_id`),
  ADD KEY `fk_users_auctions` (`user_id`);

--
-- Indexes for table `auction_histories`
--
ALTER TABLE `auction_histories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_auction_histories_user` (`user_id`),
  ADD KEY `fk_auctions_auction_history` (`auction_id`);

--
-- Indexes for table `histories`
--
ALTER TABLE `histories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auctions`
--
ALTER TABLE `auctions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `auction_histories`
--
ALTER TABLE `auction_histories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `histories`
--
ALTER TABLE `histories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `auctions`
--
ALTER TABLE `auctions`
  ADD CONSTRAINT `fk_products_auctions` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `fk_users_auctions` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `auction_histories`
--
ALTER TABLE `auction_histories`
  ADD CONSTRAINT `fk_auction_histories_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `fk_auctions_auction_history` FOREIGN KEY (`auction_id`) REFERENCES `auctions` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
