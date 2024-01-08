-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 08, 2024 at 12:09 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_2205573_bintang_uas`
--

-- --------------------------------------------------------

--
-- Table structure for table `inventory_bintang`
--

CREATE TABLE `inventory_bintang` (
  `id` int(11) NOT NULL,
  `nama_barang` varchar(255) NOT NULL,
  `jumlah` int(11) NOT NULL,
  `harga_satuan` int(11) NOT NULL,
  `lokasi` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `inventory_bintang`
--

INSERT INTO `inventory_bintang` (`id`, `nama_barang`, `jumlah`, `harga_satuan`, `lokasi`, `deskripsi`) VALUES
(1, 'Buku', 10, 50000, 'Rak Buku', 'Buku pelajaran, novel, dan komik'),
(2, 'Pensil', 100, 2000, 'Laci Meja', 'Pensil kayu dan pensil mekanik'),
(3, 'Pulpen', 50, 3000, 'Dapur', 'Pulpen tinta dan pulpen gel'),
(4, 'Penghapus', 200, 1000, 'Laci Meja', 'Penghapus pensil dan penghapus papan tulis'),
(5, 'Rautan', 50, 2000, 'Laci Meja', 'Rautan pensil dan rautan papan tulis');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `inventory_bintang`
--
ALTER TABLE `inventory_bintang`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `inventory_bintang`
--
ALTER TABLE `inventory_bintang`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
