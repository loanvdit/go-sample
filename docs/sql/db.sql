-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th8 19, 2021 lúc 06:49 AM
-- Phiên bản máy phục vụ: 10.4.17-MariaDB
-- Phiên bản PHP: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `fuji_challenge`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `fuji_bank`
--

CREATE TABLE `fuji_bank` (
  `id` char(36) NOT NULL,
  `code` char(30) NOT NULL,
  `name` varchar(300) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `fuji_bank`
--

INSERT INTO `fuji_bank` (`id`, `code`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
('05e3bcda-fd76-11eb-9a03-0242ac130003', 'Agribank', 'Ngân hàng Nông nghiệp và Phát triển Nông thôn VN', '2021-08-15 03:08:16', '2021-08-15 03:08:16', NULL),
('3082496c-fd75-11eb-9a03-0242ac130003', 'ACB', 'Ngân hàng Á Châu', '2021-08-15 03:03:12', '2021-08-15 03:03:12', NULL),
('49ca8f6a-fd75-11eb-9a03-0242ac130003', '	VIB', 'NH TMCP Quốc tế Việt Nam (VIB)', '2021-08-15 03:03:12', '2021-08-15 03:03:12', NULL),
('7ccc4e3a-fd75-11eb-9a03-0242ac130003', 'TP Bank', 'Ngân hàng Tiên Phong', '2021-08-15 03:04:53', '2021-08-15 03:04:53', NULL),
('86e204a0-fd75-11eb-9a03-0242ac130003', 'Techcombank', 'Kỹ Thương Việt Nam', '2021-08-15 03:04:53', '2021-08-15 03:04:53', NULL),
('960235ae-fd75-11eb-9a03-0242ac130003', 'VPBank', 'Việt Nam Thịnh Vượng', '2021-08-15 03:05:42', '2021-08-15 03:05:42', NULL),
('a91296de-fd75-11eb-9a03-0242ac130003', 'LienVietPostBank', 'Bưu điện Liên Việt', '2021-08-15 03:05:42', '2021-08-15 03:05:42', NULL),
('bcc675b0-fd75-11eb-9a03-0242ac130003', 'MB', 'Quân đội', '2021-08-15 03:06:33', '2021-08-15 03:06:33', NULL),
('c7c41be8-fd75-11eb-9a03-0242ac130003', 'VCB', 'Ngoại thương Việt Nam', '2021-08-15 03:06:33', '2021-08-15 03:06:33', NULL),
('d72d6ecc-fd75-11eb-9a03-0242ac130003', 'Vietinbank', 'Công Thương Việt Nam', '2021-08-15 03:07:14', '2021-08-15 03:07:14', NULL),
('dada93e2-fd75-11eb-9a03-0242ac130003', 'BIDV', 'Đầu tư và Phát triển Việt Nam', '2021-08-15 03:07:14', '2021-08-15 03:07:14', NULL),
('f3597064-fd75-11eb-9a03-0242ac130003', 'ABBank', 'Ngân hàng An Bình', '2021-08-15 03:07:52', '2021-08-15 03:07:52', NULL),
('f717a81a-fd75-11eb-9a03-0242ac130003', 'MSB', 'Hàng Hải Việt Nam', '2021-08-15 03:07:52', '2021-08-15 03:07:52', NULL);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `fuji_bank_branch`
--

CREATE TABLE `fuji_bank_branch` (
  `id` char(36) NOT NULL,
  `bank_id` char(36) NOT NULL,
  `name` varchar(300) NOT NULL,
  `address` varchar(300) NOT NULL,
  `phone` char(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `fuji_bank_branch`
--

INSERT INTO `fuji_bank_branch` (`id`, `bank_id`, `name`, `address`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES
('7ec217c8-fd76-11eb-9a03-0242ac130003', 'c7c41be8-fd75-11eb-9a03-0242ac130003', 'PGD TRUNG HÒA', 'Số 39 Vạn Phúc, phường Kim Mã, quận Ba Đình, thành phố Hà Nội', '024.37264333', '2021-08-15 03:12:21', '2021-08-15 03:12:21', NULL),
('943393f2-fd76-11eb-9a03-0242ac130003', 'c7c41be8-fd75-11eb-9a03-0242ac130003', 'PGD DUY TÂN', 'Số 2 phố Duy Tân,phường Dịch Vọng Hậu, Cầu Giấy', '024.35134146', '2021-08-15 03:12:21', '2021-08-15 03:12:21', NULL),
('ab905788-fd76-11eb-9a03-0242ac130003', 'c7c41be8-fd75-11eb-9a03-0242ac130003', 'PGD Khúc Thừa Dụ', 'Số 18, phố Khúc Thừa Dụ, phường Dịch Vọng, quận Cầu Giấy, tp. Hà Nội', '', '2021-08-15 03:13:32', '2021-08-15 03:13:32', NULL),
('aedd9c02-fd76-11eb-9a03-0242ac130003', 'c7c41be8-fd75-11eb-9a03-0242ac130003', 'CHI NHÁNH THĂNG LONG', 'Tòa nhà PVOIL Phú Thọ, số 148 Hoàng Quốc Việt, phường Nghĩa Tân, Quận Cầu Giấy, thành phố Hà Nội', '', '2021-08-15 03:13:32', '2021-08-15 03:13:32', NULL),
('c909a5c6-fd76-11eb-9a03-0242ac130003', 'c7c41be8-fd75-11eb-9a03-0242ac130003', 'PGD NGUYỄN CHÁNH', 'Số 16C, khu nhà ở thấp tầng tại Ô đất A10 (Biệt thự A10 phố Nguyễn Chánh), KĐT Nam Trung Yên, quận Cầu Giấy, Hà Nội', '024.38251139', '2021-08-15 03:14:29', '2021-08-15 03:14:29', NULL),
('cda60052-fd76-11eb-9a03-0242ac130003', 'c7c41be8-fd75-11eb-9a03-0242ac130003', 'PGD NGUYỄN VĂN HUYÊN', 'Tầng 1, Tầng 2 nhà số 35 đường Nguyễn Văn Huyên, tổ 22, phường Quan Hoa, quận Cầu Giấy, Tp Hà Nội', '024.22484111', '2021-08-15 03:14:29', '2021-08-15 03:14:29', NULL);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `fuji_partner`
--

CREATE TABLE `fuji_partner` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `bank_account` char(30) NOT NULL,
  `address` varchar(300) NOT NULL,
  `email` char(100) NOT NULL,
  `phone` char(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `fuji_transaction`
--

CREATE TABLE `fuji_transaction` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `partner_id` char(36) DEFAULT NULL,
  `bank_id` char(36) NOT NULL,
  `bank_branch_id` char(36) NOT NULL,
  `amount` bigint(30) NOT NULL,
  `fee` bigint(30) NOT NULL,
  `message` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '0-Create; 1-Pending; 2-Success',
  `mode` tinyint(1) DEFAULT NULL COMMENT '1-Send money; 2-Receive money',
  `currency` tinyint(1) NOT NULL COMMENT '0-VND; 1-BATH; 2-JPY',
  `sended_at` timestamp NULL DEFAULT NULL,
  `received_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `fuji_user`
--

CREATE TABLE `fuji_user` (
  `id` char(36) NOT NULL,
  `first_name` varchar(30) NOT NULL,
  `last_name` varchar(30) NOT NULL,
  `phone` char(20) NOT NULL,
  `address` varchar(300) NOT NULL,
  `avatar` varchar(300) NOT NULL,
  `email` char(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `verify_key` varchar(5) DEFAULT NULL,
  `verify_key_created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `is_active` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Đang đổ dữ liệu cho bảng `fuji_user`
--

INSERT INTO `fuji_user` (`id`, `first_name`, `last_name`, `phone`, `address`, `avatar`, `email`, `password`, `verify_key`, `verify_key_created_at`, `is_active`, `created_at`, `updated_at`, `deleted_at`) VALUES
('5fcf5b48-7f44-4508-9495-70ab5fd7ddba', 'Loan', 'Vu Dinh', '03562191943', 'Ha Noi, Viet Nam', '', 'loanvdit@gmail.com', '123456', '28802', '2021-08-19 03:25:41', 1, '2021-08-15 01:33:52', '2021-08-19 03:25:41', NULL);

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `fuji_user_device`
--

CREATE TABLE `fuji_user_device` (
  `id` bigint(20) NOT NULL,
  `user_id` char(36) NOT NULL,
  `device_os` char(100) NOT NULL,
  `device_mac` char(24) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `fuji_bank`
--
ALTER TABLE `fuji_bank`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Chỉ mục cho bảng `fuji_bank_branch`
--
ALTER TABLE `fuji_bank_branch`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `id_2` (`id`);

--
-- Chỉ mục cho bảng `fuji_partner`
--
ALTER TABLE `fuji_partner`
  ADD PRIMARY KEY (`id`);

--
-- Chỉ mục cho bảng `fuji_transaction`
--
ALTER TABLE `fuji_transaction`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD KEY `id_2` (`id`);

--
-- Chỉ mục cho bảng `fuji_user`
--
ALTER TABLE `fuji_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email_2` (`email`),
  ADD UNIQUE KEY `phone` (`phone`),
  ADD KEY `id` (`id`),
  ADD KEY `email` (`email`),
  ADD KEY `phone_2` (`phone`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
