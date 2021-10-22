-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 21, 2021 at 03:03 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.3.31

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ecommerce`
--

-- --------------------------------------------------------

--
-- Table structure for table `tbl_orders`
--

CREATE TABLE `tbl_orders` (
  `id` int(11) UNSIGNED NOT NULL,
  `product_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `created_date` varchar(155) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_orders`
--

INSERT INTO `tbl_orders` (`id`, `product_id`, `user_id`, `quantity`, `price`, `created_date`) VALUES
(7, 2, 8, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(8, 3, 8, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(10, 2, 3, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(11, 3, 3, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(12, 2, 1, 100, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(13, 2, 2, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(15, 3, 2, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(16, 1, 6, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(17, 2, 6, 1, NULL, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(20, 2, 1, 105, 8999, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(29, 2, 1, 105, 8999, '2021-10-18 20:39:00.4536171 +0700 +07 m=+1.550416701'),
(30, 2, 1, 3, 8999, '2021-10-18 21:11:10.2185393 +0700 +07 m=+40.680273001'),
(32, 2, 1, 3, 8999, '2021-10-18 21:21:05.9685964 +0700 +07 m=+2.803518001'),
(33, 2, 1, 0, 0, '2021-10-18 21:26:44.5511819 +0700 +07 m=+341.386103501');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_products`
--

CREATE TABLE `tbl_products` (
  `id` int(10) UNSIGNED NOT NULL,
  `product_name` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `price` float(10,2) DEFAULT NULL,
  `image_path` varchar(255) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `category` varchar(100) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `modified_date` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_products`
--

INSERT INTO `tbl_products` (`id`, `product_name`, `description`, `price`, `image_path`, `quantity`, `category`, `created_date`, `modified_date`, `isActive`) VALUES
(1, 'Mousepad Artisan Ninja FX Hayate Otsu', 'Filling Material: Foam\r\nFrame Material: Solid Wood\r\ndiy- Basic assembly to be done with simple tools by the customer, comes with instructions.\r\nThis sofa set is a clear all-rounder', 10000.00, 'https://images.tokopedia.net/img/cache/900/VqbcmM/2021/10/19/b1b357c3-995a-4ed2-989d-598a55c84342.jpg', 1000, NULL, '2021-01-04 12:19:16', NULL, 1),
(2, 'Premium Garskin G Pro Wireless Skin Mouse', '== Premium Garskin G Pro Wireless Skin Mouse ==', 8999.00, 'https://images.tokopedia.net/img/cache/900/VqbcmM/2021/5/25/3f2cf9c0-77dc-49a0-8396-021366c5a76a.jpg', 997, NULL, '2021-01-04 12:32:11', NULL, 1),
(3, 'Moondrop Aria 2021 High Performance LCP Diaphragm In Ear Earphone', 'Moondrop has designed the latest Aria from the ground up. They have changed the simple bullet-styled earpieces to beautiful metallic ergonomic ear cavities for the new Aria. They have a premium matte-finish to them. Moondrop has made them using an industrial-grade CNC machining process providing an ultimate look and beautiful artwork style design to them.', 11999.00, 'https://images.tokopedia.net/img/cache/900/VqbcmM/2021/6/4/52e37f69-62ee-40ef-9de6-e4047efec7bd.jpg', 1000, NULL, '2021-01-05 11:33:14', NULL, 1);

-- --------------------------------------------------------

--
-- Table structure for table `tbl_user`
--

CREATE TABLE `tbl_user` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(455) DEFAULT NULL,
  `created_date` varchar(155) DEFAULT NULL,
  `modified_date` varchar(155) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_user`
--

INSERT INTO `tbl_user` (`id`, `username`, `email`, `password`, `created_date`, `modified_date`) VALUES
(1, 'tono', 'tono@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-02 07:45:27', NULL),
(2, 'toni', 'toni@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-04 10:25:42 AM', NULL),
(3, 'tini', 'tini@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-05 10:58:46 AM', NULL),
(4, 'tina', 'tina@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-09 10:2:35 AM', NULL),
(5, 'mega', 'mega@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-09 3:19:7 PM', NULL),
(6, 'setyo', 'setyo@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-11 12:15:31 PM', NULL),
(7, 'stalin', 'stalin@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-13 11:3:28 AM', NULL),
(8, 'setya', 'setya@gmail.com', '$2a$10$fRIW2Dw0GFaBWT0e2A0WteKwAyb2uONdwdatlSD0IO.mZbBlWvix2', '2021-01-18 11:30:46 AM', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `tbl_wishlist`
--

CREATE TABLE `tbl_wishlist` (
  `id` int(10) UNSIGNED NOT NULL,
  `product_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `created_date` varchar(155) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tbl_wishlist`
--

INSERT INTO `tbl_wishlist` (`id`, `product_id`, `user_id`, `quantity`, `created_date`) VALUES
(20, 1, 8, 1, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(21, 2, 8, 1, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(50, 3, 3, 1, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(62, 1, 1, 101, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(63, 1, 2, 1, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(64, 2, 2, 1, '2021-10-18 20:25:45.787326 +0700 +07 m=+10.939956501'),
(69, 3, 1, 1, '2021-10-18 21:03:24.608803 +0700 +07 m=+52.171879501');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tbl_orders`
--
ALTER TABLE `tbl_orders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_products`
--
ALTER TABLE `tbl_products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_user`
--
ALTER TABLE `tbl_user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tbl_wishlist`
--
ALTER TABLE `tbl_wishlist`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tbl_orders`
--
ALTER TABLE `tbl_orders`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT for table `tbl_products`
--
ALTER TABLE `tbl_products`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `tbl_user`
--
ALTER TABLE `tbl_user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `tbl_wishlist`
--
ALTER TABLE `tbl_wishlist`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=70;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
