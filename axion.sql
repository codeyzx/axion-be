-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 25, 2023 at 04:18 AM
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
CREATE DEFINER=`root`@`localhost` PROCEDURE `get_product` (IN `user_id` INT)   BEGIN SELECT products.* FROM products INNER JOIN auctions ON products.id = auctions.product_id WHERE auctions.user_id = user_id; END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `update_product` (IN `p_product_id` VARCHAR(255), IN `p_user_id` INT, IN `p_name` VARCHAR(255), IN `p_description` TEXT, IN `p_price` DECIMAL(10,2), IN `p_image` VARCHAR(255))   BEGIN UPDATE products SET name = p_name, description = p_description, price = p_price, image = p_image WHERE products.id = (SELECT auctions.product_id FROM auctions WHERE auctions.product_id = p_product_id AND auctions.user_id = p_user_id); END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `auctions`
--

CREATE TABLE `auctions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `last_price` bigint(20) DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `bidders_count` bigint(20) DEFAULT NULL,
  `product_id` varchar(191) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `bidder_id` bigint(20) UNSIGNED DEFAULT NULL,
  `end_at` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `auctions`
--

INSERT INTO `auctions` (`id`, `name`, `last_price`, `status`, `bidders_count`, `product_id`, `user_id`, `bidder_id`, `end_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'NVIDIA GeForce RTX 3080', 14000000, 'Open', 0, 'c79d816e-76b6-49ca-8d98-a9e5f22ff7ee', 4, 2, '2023-05-15T18:00:00Z', '2023-03-01 15:41:06.000', '2023-03-24 11:47:31.460', NULL),
(2, 'Sony Alpha 7 IV', 3526000, 'Open', 0, 'bf54cb7c-f652-47ee-90ad-eb3857734c88', 5, 4, '2023-04-28T10:00:00Z', '2023-03-01 00:00:00.000', '2023-03-24 15:12:50.684', NULL),
(3, 'Fitbit Charge 5', 150000, 'Open', 0, '93fa0e6b-2052-4078-b941-f420a83d6df8', 2, 5, '2023-05-05T16:30:00Z', '2023-03-02 00:00:00.000', '2023-03-24 15:11:45.689', NULL),
(4, 'ASUS ROG Swift PG279QM', 8000000, 'Open', 0, 'be96009c-e308-47de-ab29-1e2b6f86a611', 4, 2, '2023-04-20T22:00:00Z', '2023-03-02 00:00:00.000', '2023-03-24 15:11:06.397', NULL),
(5, 'Dyson V15 Detect', 952300, 'Open', 0, '8c0b9cfd-2fd1-420a-adee-d6b1de0a8974', 5, 4, '2023-04-25T08:00:00Z', '2023-03-02 00:00:00.000', '2023-03-24 15:12:04.723', NULL),
(6, 'Sony WH-1000XM4', 652000, 'Open', 0, 'b5d8e7e7-94fd-4cdf-ac93-9b8391441e13', 2, 5, '2023-05-08T14:00:00Z', '2023-03-02 00:00:00.000', '2023-03-24 15:11:57.494', NULL),
(7, 'Microsoft Surface Laptop 4', 10000000, 'Open', 0, '14d91620-cfe1-42ee-ac59-e99bde138b54', 4, 2, '2023-04-27T19:00:00Z', '2023-03-08 00:00:00.000', '2023-03-24 15:12:12.009', NULL),
(8, 'LG C1', 6503500, 'Open', 0, '257b80e0-5f25-416f-9b9b-bc3498b0f7bb', 5, 4, '2023-05-03T11:30:00Z', '2023-03-06 00:00:00.000', '2023-03-24 15:12:23.518', NULL),
(9, 'LG OLED77C1PUB', 6512000, 'Open', 0, 'd895b3db-676c-4a27-918f-95403b7b0c34', 4, 2, '2023-05-20T22:00:00Z', '2023-03-07 00:00:00.000', '2023-03-24 15:13:14.306', NULL),
(10, 'ASUS ROG Phone 5', 2350000, 'Open', 0, '0a0937c1-7ee7-4588-8c4c-67277c4865ba', 2, 5, '2023-03-11T16:06:22Z', '2023-03-07 00:00:00.000', '2023-03-24 15:10:52.132', NULL),
(11, 'MacBook Pro M1 Pro', 9850000, 'Open', 0, 'c8bc0848-33db-481a-bf18-6fa7626165fe', 2, 4, '2023-03-12T16:06:22Z', '2023-03-11 00:00:00.000', '2023-03-24 15:12:37.814', NULL),
(12, 'LG OLED77C1PUB', 5000, 'Open', 0, '88589c95-a045-43bb-9d89-1a3552d462ba', 2, 4, '2023-03-13T16:06:22Z', '2023-03-11 00:00:00.000', '2023-03-24 14:11:45.842', NULL),
(13, 'HP Spectre x360 14', 6545000, 'Open', 0, 'a4c3c088-9077-4668-91b3-f3efa9026cf0', 4, 5, '2023-03-14T16:06:22Z', '2023-03-11 00:00:00.000', '2023-03-24 15:12:43.799', NULL),
(14, 'ASUS ROG Swift PG279QZ', 7452000, 'Open', 0, 'c16b1886-9cc9-4462-b246-3dcdccc14570', 5, 2, '2023-03-15T16:06:22Z', '2023-03-11 00:00:00.000', '2023-03-24 15:11:20.522', NULL),
(15, 'Alienware Aurora R13', 18000000, 'Open', 0, '0fe653ce-6455-49ff-9f2d-e9e125a7e07a', 4, 5, '2023-03-16T16:06:22Z', '2023-03-11 00:00:00.000', '2023-03-24 14:12:50.323', NULL),
(16, 'Sony Alpha 7 IV', 3526000, 'Open', 0, 'a1d00476-a8d6-468f-80a8-6630fb56b5df', 2, 4, '2023-04-28T10:00:00Z', '2023-03-17 00:00:00.000', '2023-03-24 15:13:00.640', NULL),
(17, 'Logitech G Pro X', 15000000, 'Open', 1, '81fc7a46-83a7-468c-b90b-f5102acb1092', 4, 2, '2023-04-30T12:00:00Z', '2023-03-18 00:00:00.000', '2023-03-24 14:16:04.590', NULL),
(18, 'LG OLED65C1PUB', 8562000, 'Open', 0, '47f228bc-1b1c-4e1a-a4e5-ddaf5855b1f5', 2, 4, '2023-05-10T20:00:00Z', '2023-03-18 00:00:00.000', '2023-03-24 15:13:07.297', NULL),
(19, 'Bose SoundLink Revolve+', 750000, 'Open', 6, 'fff8c36b-751e-4b1c-864c-43997728a5f1', 5, 5, '2023-05-01T14:00:00Z', '2023-03-18 00:00:00.000', '2023-03-24 15:11:32.766', NULL),
(20, 'Fitbit Charge 5', 105000, 'Open', 0, '2d63c459-5a77-4473-9fbd-0a281f494db6', 2, 5, '2023-07-20T15:00:00Z', '2023-03-21 00:00:00.000', '2023-03-24 14:18:28.007', NULL),
(21, 'Nest Cam Outdoor', 2050000, 'Open', 0, 'c7f29bfd-1bc3-476b-9ac5-4e9f6c07795b', 4, 2, '2023-07-10T20:00:00Z', '2023-03-22 00:00:00.000', '2023-03-24 14:18:37.055', NULL),
(22, 'Xbox Series X', 552000, 'Open', 0, '28ad385c-ae43-4737-a23c-fe5405327bc6', 5, 4, '2023-06-30T17:30:00Z', '2023-03-23 00:00:00.000', '2023-03-24 14:18:48.366', NULL),
(23, 'DJI Mavic 3', 1050000, 'Open', 0, '8860626f-0684-40c7-920a-a3800b437eaf', 2, 5, '2023-06-15T09:00:00Z', '2023-03-23 00:00:00.000', '2023-03-24 14:19:03.896', NULL),
(24, 'Bose QuietComfort 45', 3012500, 'Open', 0, '16b46919-4cbd-4938-89ad-016bc3125692', 4, 2, '2023-06-01T12:00:00Z', '2023-03-23 00:00:00.000', '2023-03-24 14:19:15.230', NULL),
(25, 'Sennheiser HD 800 S', 564000, 'Closed', 3, 'e0ebc6be-6607-4395-be4f-2f178da7aae3', 4, 5, '2023-05-07T16:30:00Z', '2023-03-24 16:16:15.356', '2023-03-24 16:24:22.738', NULL),
(26, 'Anker PowerCore+ 26800mAh', 625000, 'Closed', 0, '8c32989c-3f86-4f82-a7cb-aafc46477647', 2, 4, '2023-05-09T12:00:00Z', '2023-03-24 16:16:29.740', '2023-03-24 16:20:04.383', NULL),
(27, 'PlayStation 5', 5000000, 'Closed', 0, '58812ffd-619c-4dc7-aed7-f5d92cff5f96', 2, 5, '2023-05-12T14:00:00Z', '2023-03-24 16:16:39.159', '2023-03-24 16:19:36.112', NULL),
(28, 'NordicTrack Commercial 2950', 3250000, 'Closed', 0, '4fee7d54-8c84-46a4-9550-b3faeeb48634', 2, 4, '2023-05-08T08:00:00Z', '2023-03-24 16:16:50.082', '2023-03-24 16:18:49.849', NULL);

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
-- Dumping data for table `auction_histories`
--

INSERT INTO `auction_histories` (`id`, `auction_id`, `user_id`, `price`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 17, 2, 15000000, '2023-03-24 21:32:50.091', '2023-03-24 21:32:50.091', NULL),
(2, 19, 2, 461000, '2023-03-24 21:35:09.711', '2023-03-24 21:35:09.711', NULL),
(3, 19, 2, 471000, '2023-03-24 21:35:11.604', '2023-03-24 21:35:11.604', NULL),
(4, 19, 1, 480000, '2023-03-24 21:40:19.254', '2023-03-24 21:40:19.254', NULL),
(5, 19, 3, 485555, '2023-03-24 21:40:34.817', '2023-03-24 21:40:34.817', NULL),
(6, 19, 4, 512500, '2023-03-24 21:40:44.453', '2023-03-24 21:40:44.453', NULL),
(7, 19, 5, 750000, '2023-03-24 21:45:58.600', '2023-03-24 21:45:58.600', NULL);

--
-- Triggers `auction_histories`
--
DELIMITER $$
CREATE TRIGGER `trigger_auction` BEFORE INSERT ON `auction_histories` FOR EACH ROW UPDATE `auctions` SET last_price = new.price, bidder_id = new.user_id, bidders_count = bidders_count + 1 WHERE id = new.auction_id AND new.price > last_price AND status = 'open'
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

--
-- Dumping data for table `histories`
--

INSERT INTO `histories` (`id`, `log`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Product with id c79d816e-76b6-49ca-8d98-a9e5f22ff7ee (NVIDIA GeForce RTX 3080 Graphics Card) has been added', '2023-03-24 11:47:31.000', '2023-03-24 11:47:31.000', NULL),
(2, 'Product with id bf54cb7c-f652-47ee-90ad-eb3857734c88 (Sony Alpha 7 IV Mirrorless Camera) has been added', '2023-03-24 11:49:08.000', '2023-03-24 11:49:08.000', NULL),
(3, 'Product with id 93fa0e6b-2052-4078-b941-f420a83d6df8 (Fitbit Charge 5 Fitness Tracker) has been added', '2023-03-24 11:49:15.000', '2023-03-24 11:49:15.000', NULL),
(4, 'Product with id be96009c-e308-47de-ab29-1e2b6f86a611 (ASUS ROG Swift PG279QM Gaming Monitor) has been added', '2023-03-24 11:49:21.000', '2023-03-24 11:49:21.000', NULL),
(5, 'Product with id 8c0b9cfd-2fd1-420a-adee-d6b1de0a8974 (Dyson V15 Detect Cordless Vacuum Cleaner) has been added', '2023-03-24 11:49:27.000', '2023-03-24 11:49:27.000', NULL),
(6, 'Product with id b5d8e7e7-94fd-4cdf-ac93-9b8391441e13 (Sony WH-1000XM4 Noise-Canceling Headphones) has been added', '2023-03-24 11:49:32.000', '2023-03-24 11:49:32.000', NULL),
(7, 'Product with id 14d91620-cfe1-42ee-ac59-e99bde138b54 (Microsoft Surface Laptop 4) has been added', '2023-03-24 11:49:38.000', '2023-03-24 11:49:38.000', NULL),
(8, 'Product with id 257b80e0-5f25-416f-9b9b-bc3498b0f7bb (LG C1 OLED TV) has been added', '2023-03-24 11:49:43.000', '2023-03-24 11:49:43.000', NULL),
(9, 'Product with id d895b3db-676c-4a27-918f-95403b7b0c34 (LG OLED77C1PUB 77-inch OLED TV) has been added', '2023-03-24 14:06:22.000', '2023-03-24 14:06:22.000', NULL),
(10, 'Product with id 0a0937c1-7ee7-4588-8c4c-67277c4865ba (ASUS ROG Phone 5) has been added', '2023-03-24 14:06:34.000', '2023-03-24 14:06:34.000', NULL),
(11, 'Product with id c8bc0848-33db-481a-bf18-6fa7626165fe (MacBook Pro with M1 Pro chip) has been added', '2023-03-24 14:06:38.000', '2023-03-24 14:06:38.000', NULL),
(12, 'Product with id 88589c95-a045-43bb-9d89-1a3552d462ba (LG OLED77C1PUB 77-Inch 4K Smart OLED TV) has been added', '2023-03-24 14:11:45.000', '2023-03-24 14:11:45.000', NULL),
(13, 'Product with id a4c3c088-9077-4668-91b3-f3efa9026cf0 (HP Spectre x360 14 Laptop) has been added', '2023-03-24 14:11:56.000', '2023-03-24 14:11:56.000', NULL),
(14, 'Product with id c16b1886-9cc9-4462-b246-3dcdccc14570 (ASUS ROG Swift PG279QZ Gaming Monitor) has been added', '2023-03-24 14:12:01.000', '2023-03-24 14:12:01.000', NULL),
(15, 'Product with id 0fe653ce-6455-49ff-9f2d-e9e125a7e07a (Alienware Aurora R13 Gaming Desktop) has been added', '2023-03-24 14:12:50.000', '2023-03-24 14:12:50.000', NULL),
(16, 'Product with id a1d00476-a8d6-468f-80a8-6630fb56b5df (Sony Alpha 7 IV Mirrorless Camera) has been added', '2023-03-24 14:13:11.000', '2023-03-24 14:13:11.000', NULL),
(17, 'Product with id 81fc7a46-83a7-468c-b90b-f5102acb1092 (Logitech G Pro X Mechanical Gaming Keyboard) has been added', '2023-03-24 14:16:04.000', '2023-03-24 14:16:04.000', NULL),
(18, 'Product with id 47f228bc-1b1c-4e1a-a4e5-ddaf5855b1f5 (LG OLED65C1PUB 65-inch OLED TV) has been added', '2023-03-24 14:16:13.000', '2023-03-24 14:16:13.000', NULL),
(19, 'Product with id fff8c36b-751e-4b1c-864c-43997728a5f1 (Bose SoundLink Revolve+ Bluetooth Speaker) has been added', '2023-03-24 14:16:20.000', '2023-03-24 14:16:20.000', NULL),
(20, 'Product with id 2d63c459-5a77-4473-9fbd-0a281f494db6 (Fitbit Charge 6 Fitness & Healthy Tracker) has been added', '2023-03-24 14:18:28.000', '2023-03-24 14:18:28.000', NULL),
(21, 'Product with id c7f29bfd-1bc3-476b-9ac5-4e9f6c07795b (Nest Cam Outdoor Security Camera) has been added', '2023-03-24 14:18:37.000', '2023-03-24 14:18:37.000', NULL),
(22, 'Product with id 28ad385c-ae43-4737-a23c-fe5405327bc6 (Xbox Series X Gaming Console) has been added', '2023-03-24 14:18:48.000', '2023-03-24 14:18:48.000', NULL),
(23, 'Product with id 8860626f-0684-40c7-920a-a3800b437eaf (DJI Mavic 3 Drone) has been added', '2023-03-24 14:19:03.000', '2023-03-24 14:19:03.000', NULL),
(24, 'Product with id 16b46919-4cbd-4938-89ad-016bc3125692 (Bose QuietComfort 45 Noise Cancelling Headphones) has been added', '2023-03-24 14:19:15.000', '2023-03-24 14:19:15.000', NULL),
(25, 'Product with id 0a0937c1-7ee7-4588-8c4c-67277c4865ba (ASUS ROG Phone 5) has been updated', '2023-03-24 14:52:31.000', '2023-03-24 14:52:31.000', NULL),
(26, 'Product with id 0fe653ce-6455-49ff-9f2d-e9e125a7e07a (Alienware Aurora R13 Gaming Desktop) has been updated', '2023-03-24 14:53:26.000', '2023-03-24 14:53:26.000', NULL),
(27, 'Product with id 14d91620-cfe1-42ee-ac59-e99bde138b54 (Microsoft Surface Laptop 4) has been updated', '2023-03-24 14:53:59.000', '2023-03-24 14:53:59.000', NULL),
(28, 'Product with id 16b46919-4cbd-4938-89ad-016bc3125692 (Bose QuietComfort 45 Noise Cancelling Headphones) has been updated', '2023-03-24 14:54:13.000', '2023-03-24 14:54:13.000', NULL),
(29, 'Product with id 257b80e0-5f25-416f-9b9b-bc3498b0f7bb (LG C1 OLED TV) has been updated', '2023-03-24 14:56:00.000', '2023-03-24 14:56:00.000', NULL),
(30, 'Product with id 14d91620-cfe1-42ee-ac59-e99bde138b54 (Microsoft Surface Laptop 4) has been updated', '2023-03-24 14:56:21.000', '2023-03-24 14:56:21.000', NULL),
(31, 'Product with id 14d91620-cfe1-42ee-ac59-e99bde138b54 (Microsoft Surface Laptop 4) has been updated', '2023-03-24 14:56:26.000', '2023-03-24 14:56:26.000', NULL),
(32, 'Product with id 28ad385c-ae43-4737-a23c-fe5405327bc6 (Xbox Series X Gaming Console) has been updated', '2023-03-24 14:57:46.000', '2023-03-24 14:57:46.000', NULL),
(33, 'Product with id 2d63c459-5a77-4473-9fbd-0a281f494db6 (Fitbit Charge 6 Fitness & Healthy Tracker) has been updated', '2023-03-24 14:58:12.000', '2023-03-24 14:58:12.000', NULL),
(34, 'Product with id 47f228bc-1b1c-4e1a-a4e5-ddaf5855b1f5 (LG OLED65C1PUB 65-inch OLED TV) has been updated', '2023-03-24 14:58:41.000', '2023-03-24 14:58:41.000', NULL),
(35, 'Product with id 47f228bc-1b1c-4e1a-a4e5-ddaf5855b1f5 (LG OLED65C1PUB 65-inch OLED TV) has been updated', '2023-03-24 14:58:56.000', '2023-03-24 14:58:56.000', NULL),
(36, 'Product with id be96009c-e308-47de-ab29-1e2b6f86a611 (ASUS ROG Swift PG279QM Gaming Monitor) has been updated', '2023-03-24 14:59:22.000', '2023-03-24 14:59:22.000', NULL),
(37, 'Product with id c16b1886-9cc9-4462-b246-3dcdccc14570 (ASUS ROG Swift PG279QZ Gaming Monitor) has been updated', '2023-03-24 14:59:45.000', '2023-03-24 14:59:45.000', NULL),
(38, 'Product with id fff8c36b-751e-4b1c-864c-43997728a5f1 (Bose SoundLink Revolve+ Bluetooth Speaker) has been updated', '2023-03-24 15:00:15.000', '2023-03-24 15:00:15.000', NULL),
(39, 'Product with id 8860626f-0684-40c7-920a-a3800b437eaf (DJI Mavic 3 Drone) has been updated', '2023-03-24 15:00:42.000', '2023-03-24 15:00:42.000', NULL),
(40, 'Product with id 8c0b9cfd-2fd1-420a-adee-d6b1de0a8974 (Dyson V15 Detect Cordless Vacuum Cleaner) has been updated', '2023-03-24 15:01:16.000', '2023-03-24 15:01:16.000', NULL),
(41, 'Product with id 81fc7a46-83a7-468c-b90b-f5102acb1092 (Logitech G Pro X Mechanical Gaming Keyboard) has been updated', '2023-03-24 15:02:04.000', '2023-03-24 15:02:04.000', NULL),
(42, 'Product with id 93fa0e6b-2052-4078-b941-f420a83d6df8 (Fitbit Charge 5 Fitness Tracker) has been updated', '2023-03-24 15:02:25.000', '2023-03-24 15:02:25.000', NULL),
(43, 'Product with id 93fa0e6b-2052-4078-b941-f420a83d6df8 (Fitbit Charge 5 Fitness Tracker) has been updated', '2023-03-24 15:02:29.000', '2023-03-24 15:02:29.000', NULL),
(44, 'Product with id b5d8e7e7-94fd-4cdf-ac93-9b8391441e13 (Sony WH-1000XM4 Noise-Canceling Headphones) has been updated', '2023-03-24 15:02:55.000', '2023-03-24 15:02:55.000', NULL),
(45, 'Product with id bf54cb7c-f652-47ee-90ad-eb3857734c88 (Sony Alpha 7 IV Mirrorless Camera) has been updated', '2023-03-24 15:03:18.000', '2023-03-24 15:03:18.000', NULL),
(46, 'Product with id a1d00476-a8d6-468f-80a8-6630fb56b5df (Sony Alpha 7 IV Mirrorless Camera) has been updated', '2023-03-24 15:03:43.000', '2023-03-24 15:03:43.000', NULL),
(47, 'Product with id a1d00476-a8d6-468f-80a8-6630fb56b5df (Sony Alpha 5 Mirrorless Camera) has been updated', '2023-03-24 15:03:55.000', '2023-03-24 15:03:55.000', NULL),
(48, 'Product with id c7f29bfd-1bc3-476b-9ac5-4e9f6c07795b (Nest Cam Outdoor Security Camera) has been updated', '2023-03-24 15:04:19.000', '2023-03-24 15:04:19.000', NULL),
(49, 'Product with id c79d816e-76b6-49ca-8d98-a9e5f22ff7ee (NVIDIA GeForce RTX 3080 Graphics Card) has been updated', '2023-03-24 15:04:53.000', '2023-03-24 15:04:53.000', NULL),
(50, 'Product with id c8bc0848-33db-481a-bf18-6fa7626165fe (MacBook Pro with M1 Pro chip) has been updated', '2023-03-24 15:05:20.000', '2023-03-24 15:05:20.000', NULL),
(51, 'Product with id d895b3db-676c-4a27-918f-95403b7b0c34 (LG OLED77C1PUB 77-inch OLED TV) has been updated', '2023-03-24 15:05:48.000', '2023-03-24 15:05:48.000', NULL),
(52, 'Product with id 88589c95-a045-43bb-9d89-1a3552d462ba (LG OLED77C1PUB 77-Inch 4K Smart OLED TV) has been updated', '2023-03-24 15:06:56.000', '2023-03-24 15:06:56.000', NULL),
(53, 'Product with id a4c3c088-9077-4668-91b3-f3efa9026cf0 (HP Spectre x360 14 Laptop) has been updated', '2023-03-24 15:07:22.000', '2023-03-24 15:07:22.000', NULL),
(54, 'Product with id a1d00476-a8d6-468f-80a8-6630fb56b5df (Sony Alpha 5 Mirrorless Camera) has been updated', '2023-03-24 15:07:39.000', '2023-03-24 15:07:39.000', NULL),
(55, 'Product with id e0ebc6be-6607-4395-be4f-2f178da7aae3 (Sennheiser HD 800 S Headphones) has been added', '2023-03-24 16:16:15.000', '2023-03-24 16:16:15.000', NULL),
(56, 'Product with id 8c32989c-3f86-4f82-a7cb-aafc46477647 (Anker PowerCore+ 26800mAh Portable Charger) has been added', '2023-03-24 16:16:29.000', '2023-03-24 16:16:29.000', NULL),
(57, 'Product with id 58812ffd-619c-4dc7-aed7-f5d92cff5f96 (PlayStation 5 Console) has been added', '2023-03-24 16:16:39.000', '2023-03-24 16:16:39.000', NULL),
(58, 'Product with id 4fee7d54-8c84-46a4-9550-b3faeeb48634 (PlayStation 5 Console) has been added', '2023-03-24 16:16:50.000', '2023-03-24 16:16:50.000', NULL),
(59, 'Product with id 4fee7d54-8c84-46a4-9550-b3faeeb48634 (PlayStation 5 Console) has been updated', '2023-03-24 16:18:43.000', '2023-03-24 16:18:43.000', NULL),
(60, 'Product with id 58812ffd-619c-4dc7-aed7-f5d92cff5f96 (PlayStation 5 Console) has been updated', '2023-03-24 16:19:37.000', '2023-03-24 16:19:37.000', NULL),
(61, 'Product with id 8c32989c-3f86-4f82-a7cb-aafc46477647 (Anker PowerCore+ 26800mAh Portable Charger) has been updated', '2023-03-24 16:20:00.000', '2023-03-24 16:20:00.000', NULL),
(62, 'Product with id e0ebc6be-6607-4395-be4f-2f178da7aae3 (Sennheiser HD 800 S Headphones) has been updated', '2023-03-24 16:20:33.000', '2023-03-24 16:20:33.000', NULL);

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
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `price`, `image`, `created_at`, `updated_at`, `deleted_at`) VALUES
('0a0937c1-7ee7-4588-8c4c-67277c4865ba', 'ASUS ROG Phone 5', 'Experience the ultimate in mobile gaming with the new ASUS ROG Phone 5. With a 144Hz AMOLED display and a powerful Snapdragon 888 processor, you can play your favorite games at the highest settings with no lag.', 2350000, './public/covers/1679644351-9194e173-1905-4b8f-8c44-2d99ab4b2a3b.jpg', '2023-03-24 14:06:34.038', '2023-03-24 14:52:31.750', NULL),
('0fe653ce-6455-49ff-9f2d-e9e125a7e07a', 'Alienware Aurora R13 Gaming Desktop', 'Experience the ultimate in gaming performance with the new Alienware Aurora R13 gaming desktop. With a powerful NVIDIA GeForce RTX 3090 graphics card and an Intel Core i9 processor, you can play your favorite games at the highest settings with no lag.', 18000000, './public/covers/1679644406-20220422_170407-01-scaled.jpeg', '2023-03-24 14:12:50.322', '2023-03-24 14:53:26.556', NULL),
('14d91620-cfe1-42ee-ac59-e99bde138b54', 'Microsoft Surface Laptop 4', 'Stay organized and productive with the new Microsoft Surface Laptop 4. With a powerful AMD Ryzen processor and up to 19 hours of battery life, you can work and play wherever you go.', 10000000, './public/covers/1679644439-b21b03e7-ea23-4339-99b8-185e3d3e916e.png', '2023-03-24 11:49:38.321', '2023-03-24 14:56:26.476', NULL),
('16b46919-4cbd-4938-89ad-016bc3125692', 'Bose QuietComfort 45 Noise Cancelling Headphones', 'Experience immersive sound with the new Bose QuietComfort 45 noise cancelling headphones. With advanced noise cancellation technology, you can enjoy your music in peace and quiet.', 3012500, './public/covers/1679644453-118f8dc9-6549-4a2f-9ce1-ff1ef77845d7.png', '2023-03-24 14:19:15.229', '2023-03-24 14:54:13.888', NULL),
('257b80e0-5f25-416f-9b9b-bc3498b0f7bb', 'LG C1 OLED TV', 'Transform your living room into a home theater with the new LG C1 OLED TV. With stunning picture quality and advanced AI processing, you can watch your favorite movies and shows like never before.', 6503500, './public/covers/1679644560-lg-c1-oled-tv-cnet-review-2021-hero.webp', '2023-03-24 11:49:43.178', '2023-03-24 14:56:00.090', NULL),
('28ad385c-ae43-4737-a23c-fe5405327bc6', 'Xbox Series X Gaming Console', 'Get your game on with the new Xbox Series X. With lightning-fast load times and stunning 4K graphics, you can experience gaming like never before.', 5552000, './public/covers/1679644666-ef9e8f9f-4141-409c-9bab-659102bcc6b2.jpg', '2023-03-24 14:18:48.365', '2023-03-24 14:57:46.422', NULL),
('2d63c459-5a77-4473-9fbd-0a281f494db6', 'Fitbit Charge 6 Fitness & Healthy Tracker', 'Get fit and stay healthy with the new Fitbit Charge 6. With advanced health tracking features and up to 14 days of battery life, you can achieve your fitness goals with ease.', 1050000, './public/covers/1679644692-61xh+Cewq5L.jpg', '2023-03-24 14:18:28.005', '2023-03-24 14:58:12.015', NULL),
('47f228bc-1b1c-4e1a-a4e5-ddaf5855b1f5', 'LG OLED65C1PUB 65-inch OLED TV', 'Experience the ultimate viewing experience with the new LG OLED65C1PUB 65-inch OLED TV. With stunning 4K resolution and HDR technology, you can enjoy your favorite movies and shows with lifelike colors and deep blacks.', 8562000, './public/covers/1679644721-LG-C1-OLED-Contrast-Black-Level.jpg', '2023-03-24 14:16:13.950', '2023-03-24 14:58:56.949', NULL),
('4fee7d54-8c84-46a4-9550-b3faeeb48634', 'NordicTrack Commercial 2950', 'Get the ultimate workout with the new NordicTrack Commercial 2950 treadmill. With a 22-inch touchscreen display and iFit membership, you can enjoy a personalized workout experience from the comfort of your home.', 3250000, './public/covers/1679649523-download (8).jfif', '2023-03-24 16:16:50.081', '2023-03-24 16:18:43.697', NULL),
('58812ffd-619c-4dc7-aed7-f5d92cff5f96', 'PlayStation 5 Console', 'Experience immersive gaming with the new PlayStation 5 console. With lightning-fast load times and stunning graphics, you can play the latest games like never before.', 5000000, './public/covers/1679649577-PlayStation-5-Full-Package-Rental-Service-in-Ho-Chi-Minh-City-by-Thue-Nhanh-beead06d-62ec-4af7-a60f-886c188badfc.webp', '2023-03-24 16:16:39.158', '2023-03-24 16:19:37.213', NULL),
('81fc7a46-83a7-468c-b90b-f5102acb1092', 'Logitech G Pro X Mechanical Gaming Keyboard', 'Upgrade your gaming setup with the new Logitech G Pro X Mechanical Gaming Keyboard. With swappable switches and advanced customization options, you can optimize your gaming experience for any game or playstyle.', 1500000, './public/covers/1679644924-logitech-g-pro-x-keyboard.jpg', '2023-03-24 14:16:04.588', '2023-03-24 15:02:04.032', NULL),
('88589c95-a045-43bb-9d89-1a3552d462ba', 'LG OLED77C1PTB 4K Smart OLED TV', 'Get the ultimate home entertainment experience with the new LG OLED77C1PUB 77-Inch 4K Smart OLED TV. With stunning picture quality and immersive sound, you can enjoy your favorite movies and shows like never before.', 654100, './public/covers/1679645216-b36d45e2-84a2-45a4-8290-3c8b4c40d190.jpg', '2023-03-24 14:11:45.840', '2023-03-24 15:06:56.437', NULL),
('8860626f-0684-40c7-920a-a3800b437eaf', 'DJI Mavic 3 Drone', 'Take your photography to the next level with the new DJI Mavic 3 drone. With a Hasselblad camera and 4K video recording, you can capture stunning aerial shots with ease.', 1050000, './public/covers/1679644842-618a060c4a887.jpg', '2023-03-24 14:19:03.894', '2023-03-24 15:00:42.219', NULL),
('8c0b9cfd-2fd1-420a-adee-d6b1de0a8974', 'Dyson V15 Detect Cordless Vacuum Cleaner', 'Keep your home clean and tidy with the new Dyson V15 Detect cordless vacuum cleaner. With powerful suction and advanced laser technology, you can remove even the smallest dust particles and allergens.', 952300, './public/covers/1679644876-a86094ba-adaf-4159-b757-96326dd339d2.jpg', '2023-03-24 11:49:27.154', '2023-03-24 15:01:16.293', NULL),
('8c32989c-3f86-4f82-a7cb-aafc46477647', 'Anker PowerCore+ 26800mAh Portable Charger', 'Keep your devices charged on the go with the new Anker PowerCore+ 26800mAh portable charger. With USB-C and USB-A ports, you can charge multiple devices simultaneously.', 625000, './public/covers/1679649600-61vNDyrUlxL._AC_UF894,1000_QL80_.jpg', '2023-03-24 16:16:29.738', '2023-03-24 16:20:00.295', NULL),
('93fa0e6b-2052-4078-b941-f420a83d6df8', 'Fitbit Charge 5 Fitness Tracker', 'Get in shape and stay motivated with the new Fitbit Charge 5 fitness tracker. With advanced health monitoring features and up to 7 days of battery life, you can track your progress and reach your goals.', 150000, './public/covers/1679644949-download.jfif', '2023-03-24 11:49:15.327', '2023-03-24 15:02:29.716', NULL),
('a1d00476-a8d6-468f-80a8-6630fb56b5df', 'Sony Alpha 5 Mirrorless Camera', 'Capture stunning photos and videos with the new Sony Alpha 5 mirrorless camera. With a 33-megapixel full-frame sensor and 4K video recording, you can take your creativity to the next level.', 651023, './public/covers/1679645035-images.jfif', '2023-03-24 14:13:11.480', '2023-03-24 15:07:39.552', NULL),
('a4c3c088-9077-4668-91b3-f3efa9026cf0', 'HP Spectre x360 14 Laptop', 'Experience powerful performance and sleek style with the new HP Spectre x360 14 laptop. With a 11th Gen Intel Core i7 processor and a stunning 14-inch OLED display, you can work and play with ease.', 6545000, './public/covers/1679645242-download (7).jfif', '2023-03-24 14:11:56.419', '2023-03-24 15:07:22.767', NULL),
('b5d8e7e7-94fd-4cdf-ac93-9b8391441e13', 'Sony WH-1000XM4 Noise-Canceling Headphones', 'Enjoy your favorite music on-the-go with the new Sony WH-1000XM4 noise-canceling headphones. With up to 30 hours of battery life and advanced noise cancellation technology, you can listen to your music without any distractions.', 652000, './public/covers/1679644975-download (1).jfif', '2023-03-24 11:49:32.867', '2023-03-24 15:02:55.605', NULL),
('be96009c-e308-47de-ab29-1e2b6f86a611', 'ASUS ROG Swift PG279QM Gaming Monitor', 'Experience smooth and immersive gameplay with the new ASUS ROG Swift PG279QM gaming monitor. With a fast 165Hz refresh rate and G-SYNC technology, you can play your favorite games without any lag or tearing.', 8000000, './public/covers/1679644762-77521944-4a98-41ec-ab0e-d59e9778c3cf.jpg', '2023-03-24 11:49:21.468', '2023-03-24 14:59:22.923', NULL),
('bf54cb7c-f652-47ee-90ad-eb3857734c88', 'Sony Alpha 7 IV Mirrorless Camera', 'Capture stunning photos and videos with the new Sony Alpha 7 IV mirrorless camera. With a 33-megapixel full-frame sensor and 4K video recording, you can take your creativity to the next level.', 3526000, './public/covers/1679644998-download (2).jfif', '2023-03-24 11:49:08.459', '2023-03-24 15:03:18.599', NULL),
('c16b1886-9cc9-4462-b246-3dcdccc14570', 'ASUS ROG Swift PG279QZ Gaming Monitor', 'Upgrade your gaming experience with the powerful ASUS ROG Swift PG279QZ gaming monitor. With a 27-inch WQHD display and a 165Hz refresh rate, you can play your favorite games at the highest settings with no lag.', 7452000, './public/covers/1679644785-h525.png', '2023-03-24 14:12:01.002', '2023-03-24 14:59:45.037', NULL),
('c79d816e-76b6-49ca-8d98-a9e5f22ff7ee', 'NVIDIA GeForce RTX 3080 Graphics Card', 'Upgrade your gaming experience with the powerful NVIDIA GeForce RTX 3080 graphics card. With cutting-edge ray tracing technology and advanced AI rendering, you can play the latest games at ultra-high settings.', 1400000, './public/covers/1679645093-download (4).jfif', '2023-03-24 11:47:31.458', '2023-03-24 15:04:53.455', NULL),
('c7f29bfd-1bc3-476b-9ac5-4e9f6c07795b', 'Nest Cam Outdoor Security Camera', 'Keep your home safe and secure with the new Nest Cam Outdoor. With 24/7 live streaming and advanced motion detection, you can monitor your property from anywhere.', 2050000, './public/covers/1679645059-download (3).jfif', '2023-03-24 14:18:37.053', '2023-03-24 15:04:19.707', NULL),
('c8bc0848-33db-481a-bf18-6fa7626165fe', 'MacBook Pro with M1 Pro chip', 'Stay connected and productive on the go with the new MacBook Pro with M1 Pro chip. With up to 20 hours of battery life and lightning-fast performance, you can work and play wherever you go.', 9850000, './public/covers/1679645120-download (5).jfif', '2023-03-24 14:06:38.834', '2023-03-24 15:05:20.644', NULL),
('d895b3db-676c-4a27-918f-95403b7b0c34', 'LG OLED77C1PUB 77-inch OLED TV', 'Upgrade your home entertainment system with the new LG OLED77C1PUB 77-inch OLED TV. With perfect blacks and infinite contrast, you can enjoy your favorite movies and shows in stunning detail.', 6512000, './public/covers/1679645148-download (6).jfif', '2023-03-24 14:06:22.604', '2023-03-24 15:05:48.286', NULL),
('e0ebc6be-6607-4395-be4f-2f178da7aae3', 'Sennheiser HD 800 S Headphones', 'Experience high-fidelity sound with the new Sennheiser HD 800 S headphones. With an open-back design and innovative transducer technology, you can hear every detail of your music with clarity and precision.', 564000, './public/covers/1679649633-design-medium.jpg', '2023-03-24 16:16:15.353', '2023-03-24 16:20:33.155', NULL),
('fff8c36b-751e-4b1c-864c-43997728a5f1', 'Bose SoundLink Revolve+ Bluetooth Speaker', 'Get immersive sound and powerful bass with the new Bose SoundLink Revolve+ Bluetooth speaker. With a 360-degree design, you can enjoy your music from any angle.', 451000, './public/covers/1679644815-9a283992-8324-49bc-bf90-31c60f8329d7.jpg', '2023-03-24 14:16:20.459', '2023-03-24 15:00:15.441', NULL);

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
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `address`, `phone`, `role`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Ahmad Joni', 'admin@gmail.com', '$2a$14$RbnY1kuLqATjLcLWXgmyJ.Jtg8M8BuB8YvjBPoAIqHeqweNjM/YVq', 'Jl. Mahar Martanegara No.48, Utama, Kec. Cimahi Sel., Kota Cimahi, Jawa Barat 40521', '+6285624918686', 'Admin', '2023-03-23 20:47:01.736', '2023-03-23 20:47:01.736', NULL),
(2, 'Ruby Sakura', 'ruby@gmail.com', '$2a$14$k0f248GjjYHPw7U815Q7jeD2/b7n1jTb9dzBFDIR0mzYMYPc8NQl6', 'Jl. Raya Wangun No.21, RT.01/RW.06, Sindangsari, Kec. Bogor Tim., Kota Bogor, Jawa Barat 16146', '+6285642918600', 'Users', '2023-03-23 22:32:34.736', '2023-03-23 22:32:34.736', NULL),
(3, 'Theron J Ryan', 'theron@gmail.com', '$2a$14$ZHfJXDlbsV.YowDP0eEjs..Qg2kBnuRq9U1ENhIZeV9WyofDopmPW', '3507 Henery Street', '317-218-0076', 'Operator', '2023-03-23 22:34:30.369', '2023-03-23 22:34:30.369', NULL),
(4, 'Kristie A Lockhart', 'kristie@gmail.com', '$2a$14$rKysaPOgyxb3pfMswzZ59etbjSJMkTX8dWtlVdSYU5U.tfp2a6AHy', '4351 Grant Street', '903-876-3401', 'Users', '2023-03-23 22:35:03.120', '2023-03-23 22:35:03.120', NULL),
(5, 'Shari R Bischoff', 'shari@gmail.com', '$2a$14$gO/NFvGRwpBSv8fEoawOx.CXdTSSlf3SwAZJ4DdXPofBxbbCxhaay', '248 Turkey Pen Lane', '334-548-5988', 'Users', '2023-03-23 22:35:51.362', '2023-03-23 22:35:51.362', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auctions`
--
ALTER TABLE `auctions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `product_id` (`product_id`),
  ADD KEY `fk_auctions_bidder` (`bidder_id`),
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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT for table `auction_histories`
--
ALTER TABLE `auction_histories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `histories`
--
ALTER TABLE `histories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=63;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `auctions`
--
ALTER TABLE `auctions`
  ADD CONSTRAINT `fk_auctions_bidder` FOREIGN KEY (`bidder_id`) REFERENCES `users` (`id`),
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
