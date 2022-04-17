-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 17, 2022 at 02:30 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gosample`
--

-- --------------------------------------------------------

--
-- Table structure for table `audit_logs`
--

CREATE TABLE `audit_logs` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `userid` int(11) DEFAULT NULL,
  `logtype` varchar(255) DEFAULT NULL,
  `module` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `audit_logs`
--

INSERT INTO `audit_logs` (`id`, `name`, `userid`, `logtype`, `module`, `created_at`, `updated_at`, `deleted_at`) VALUES
(21, 'john location ', 1, 'Portal', 'Calendar', '2022-04-17 12:23:47', '2022-04-17 12:23:47', NULL),
(22, 'joe location', 2, 'Portal', 'Calendar', '2022-04-17 12:23:47', '2022-04-17 12:23:47', NULL),
(23, 'jane location', 3, 'Portal', 'Calendar', '2022-04-17 12:23:48', '2022-04-17 12:23:48', NULL),
(24, 'john location ', 1, 'Portal', 'Calendar', '2022-04-17 12:25:49', '2022-04-17 12:25:49', NULL),
(25, 'joe location', 2, 'Portal', 'Calendar', '2022-04-17 12:25:49', '2022-04-17 12:25:49', NULL),
(26, 'jane location', 3, 'Portal', 'Calendar', '2022-04-17 12:25:49', '2022-04-17 12:25:49', NULL),
(27, 'john location ', 1, 'Portal', 'Calendar', '2022-04-17 12:29:22', '2022-04-17 12:29:22', NULL),
(28, 'joe location', 2, 'Portal', 'Calendar', '2022-04-17 12:29:22', '2022-04-17 12:29:22', NULL),
(29, 'jane location', 3, 'Portal', 'Calendar', '2022-04-17 12:29:22', '2022-04-17 12:29:22', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `encounters`
--

CREATE TABLE `encounters` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `userid` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `encounters`
--

INSERT INTO `encounters` (`id`, `name`, `userid`, `created_at`, `updated_at`, `deleted_at`) VALUES
(17, 'john location ', 1, '2022-04-17 12:23:47', '2022-04-17 12:23:47', NULL),
(18, 'joe location', 2, '2022-04-17 12:23:47', '2022-04-17 12:23:47', NULL),
(19, 'jane location', 3, '2022-04-17 12:23:47', '2022-04-17 12:23:47', NULL),
(20, 'john location ', 1, '2022-04-17 12:25:48', '2022-04-17 12:25:48', NULL),
(21, 'joe location', 2, '2022-04-17 12:25:49', '2022-04-17 12:25:49', NULL),
(22, 'jane location', 3, '2022-04-17 12:25:49', '2022-04-17 12:25:49', NULL),
(23, 'john location ', 1, '2022-04-17 12:29:21', '2022-04-17 12:29:21', NULL),
(24, 'joe location', 2, '2022-04-17 12:29:22', '2022-04-17 12:29:22', NULL),
(25, 'jane location', 3, '2022-04-17 12:29:22', '2022-04-17 12:29:22', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`) VALUES
(1, 'test1');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `audit_logs`
--
ALTER TABLE `audit_logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `encounters`
--
ALTER TABLE `encounters`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `audit_logs`
--
ALTER TABLE `audit_logs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;

--
-- AUTO_INCREMENT for table `encounters`
--
ALTER TABLE `encounters`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
