-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 21, 2019 at 07:18 PM
-- Server version: 10.4.8-MariaDB
-- PHP Version: 7.3.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `desaku`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id` int(11) NOT NULL,
  `uuid_person` char(36) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `user_id_group` int(11) NOT NULL,
  `aktif` enum('Y','T') NOT NULL DEFAULT 'Y'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `uuid_person`, `username`, `password`, `email`, `user_id_group`, `aktif`) VALUES
(1, '000c04ff-4da8-42cd-9985-1f67c1ce0b54', 'faiz', '9e13342cfc1ea881aeb27a51e009ee9162d45dfb', 'faizuloke@gmail.com', 1, 'Y');

-- --------------------------------------------------------

--
-- Table structure for table `keluarga`
--

CREATE TABLE `keluarga` (
  `id` int(11) NOT NULL,
  `dasar_id_person` char(36) NOT NULL,
  `id_status_keluarga` tinyint(4) NOT NULL,
  `kepada_id_person` char(36) NOT NULL,
  `wali` enum('Y','T') NOT NULL DEFAULT 'T'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `keluarga`
--

INSERT INTO `keluarga` (`id`, `dasar_id_person`, `id_status_keluarga`, `kepada_id_person`, `wali`) VALUES
(1, '000c04ff-4da8-42cd-9985-1f67c1ce0b54', 9, '003af9ca-78c0-40d6-8747-f74b958ec887', 'Y'),
(2, '000c04ff-4da8-42cd-9985-1f67c1ce0b54', 12, '00786def-d38d-447c-9fc2-07af809d5dcc', 'T');

-- --------------------------------------------------------

--
-- Table structure for table `pekerjaan`
--

CREATE TABLE `pekerjaan` (
  `id` int(11) NOT NULL,
  `nama` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pekerjaan`
--

INSERT INTO `pekerjaan` (`id`, `nama`) VALUES
(1, 'Pejabat Eksekutif/Pemerintah'),
(2, 'Pejabat Legislatif (DPR/DPD/DPRD)'),
(3, 'Hakim/Jaksa/Pengacara'),
(4, 'TNI'),
(5, 'POLRI'),
(6, 'Pegawai Negeri Sipil (PNS)'),
(7, 'Dosen/Guru'),
(8, 'Dokter'),
(9, 'Perawat/Bidan'),
(10, 'Nelayan/Perikanan'),
(11, 'Petani/Pekebun'),
(12, 'Peternak'),
(13, 'Pedagang'),
(14, 'Pengrajin/Pemahat'),
(15, 'Arsitek'),
(16, 'Konstruksi'),
(17, 'Transportasi'),
(18, 'Montir/Mekanik'),
(19, 'Pertambangan'),
(20, 'Karyawan Pabrik'),
(21, 'Perkapalan'),
(22, 'Perhotelan'),
(23, 'Periklanan/Pemasaran'),
(24, 'Penulis'),
(25, 'Wartawan/Jurnalis'),
(26, 'Seniman'),
(27, 'Fotografer'),
(28, 'Desain Grafis'),
(29, 'Programmer'),
(30, 'Akuntan'),
(31, 'Aktor/Aktris'),
(32, 'Musikus'),
(33, 'Pilot'),
(34, 'Pramugari/Pramugara'),
(35, 'Atlet'),
(36, 'Pekerjaan Tidak Tetap');

-- --------------------------------------------------------

--
-- Table structure for table `pendidikan`
--

CREATE TABLE `pendidikan` (
  `id` int(11) NOT NULL,
  `nama` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pendidikan`
--

INSERT INTO `pendidikan` (`id`, `nama`) VALUES
(1, 'Setingkat SD'),
(2, 'Setingkat SMP'),
(3, 'Setingkat SMA'),
(4, 'D1'),
(5, 'D2'),
(6, 'D3'),
(7, 'S1'),
(8, 'S2'),
(9, 'S3');

-- --------------------------------------------------------

--
-- Table structure for table `penghasilan`
--

CREATE TABLE `penghasilan` (
  `id` int(11) NOT NULL,
  `keterangan` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `penghasilan`
--

INSERT INTO `penghasilan` (`id`, `keterangan`) VALUES
(1, '< 500.000'),
(2, '500.000 - 1.500.000'),
(3, '1.500.000 - 2.500.000'),
(4, '2.500.000 - 3.500.000'),
(5, '3.500.000 - 5.000.000'),
(6, '5.000.000 - 10.000.000'),
(7, '> 10.000.000');

-- --------------------------------------------------------

--
-- Table structure for table `person`
--

CREATE TABLE `person` (
  `uuid` char(36) NOT NULL,
  `nokk` char(16) NOT NULL,
  `nik` char(16) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `jk` enum('L','P') NOT NULL DEFAULT 'L',
  `tempat_lahir` varchar(25) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `anak_ke` tinyint(4) NOT NULL,
  `jum_saudara` tinyint(4) NOT NULL,
  `phone1` varchar(20) DEFAULT NULL,
  `phone2` varchar(20) DEFAULT NULL,
  `id_pendidikan` tinyint(4) DEFAULT NULL,
  `id_pekerjaan` tinyint(4) DEFAULT NULL,
  `id_penghasilan` tinyint(4) DEFAULT NULL,
  `pembuat` char(36) DEFAULT NULL,
  `wafat` enum('Y','T') NOT NULL DEFAULT 'T'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `person`
--

INSERT INTO `person` (`uuid`, `nokk`, `nik`, `nama`, `jk`, `tempat_lahir`, `tanggal_lahir`, `anak_ke`, `jum_saudara`, `phone1`, `phone2`, `id_pendidikan`, `id_pekerjaan`, `id_penghasilan`, `pembuat`, `wafat`) VALUES
('000c04ff-4da8-42cd-9985-1f67c1ce0b54', '3529073010210001', '3529073010210002', 'Faizul Amali', 'L', 'Sumenep', '1993-06-20', 1, 2, '085228505163', NULL, 1, 1, 1, '000c04ff-4da8-42cd-9985-1f67c1ce0b54', 'T'),
('003af9ca-78c0-40d6-8747-f74b958ec887', '3529073010210005', '3529073010210006', 'Rahmah', 'L', 'Riau', '1982-11-06', 1, 3, NULL, NULL, 3, 4, 4, '000c04ff-4da8-42cd-9985-1f67c1ce0b54', 'T'),
('00786def-d38d-447c-9fc2-07af809d5dcc', '3529073010210003', '3529073010210004', 'Nur Fatimah', 'P', 'Sulawesi', '1998-11-10', 3, 3, NULL, NULL, 4, 4, 3, '000c04ff-4da8-42cd-9985-1f67c1ce0b54', 'T');

-- --------------------------------------------------------

--
-- Table structure for table `status_keluarga`
--

CREATE TABLE `status_keluarga` (
  `id` tinyint(4) NOT NULL,
  `keterangan_dasar` varchar(45) NOT NULL,
  `keterangan_kebalikan` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `status_keluarga`
--

INSERT INTO `status_keluarga` (`id`, `keterangan_dasar`, `keterangan_kebalikan`) VALUES
(1, 'Kakek dari ayah', 'Cucu'),
(2, 'Kakek dari ibu', 'Cucu'),
(3, 'Kakek dari pamannya ayah', 'Cucu'),
(4, 'Kakek dari pamannya ibu', 'Cucu'),
(5, 'Nenek dari ayah', 'Cucu'),
(6, 'Nenek dari ibu', 'Cucu'),
(7, 'Nenek dari bibinya ayah', 'Cucu'),
(8, 'Nenek dari bibinya ibu', 'Cucu'),
(9, 'Ayah kandung', 'Anak'),
(10, 'Ayah tiri', 'Anak'),
(11, 'Ayah dari suaminya ibu susuan', 'Anak'),
(12, 'Ibu kandung', 'Anak'),
(13, 'Ibu tiri', 'Anak'),
(14, 'Ibu susuan', 'Anak'),
(15, 'Paman dari ayah', 'Keponakan'),
(16, 'Paman dari ibu', 'Keponakan'),
(17, 'Bibi dari ayah', 'Keponakan'),
(18, 'Bibi dari ibu', 'Keponakan'),
(19, 'Saudara kandung', 'Saudara kandung'),
(20, 'Saudara tiri', 'Saudara tiri'),
(21, 'Saudara seayah', 'Saudara seayah'),
(22, 'Saudara seibu', 'Saudara seibu'),
(23, 'Saudara sesusuan', 'Saudara sesusuan');

-- --------------------------------------------------------

--
-- Table structure for table `user_group`
--

CREATE TABLE `user_group` (
  `id` int(11) NOT NULL,
  `nama` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user_group`
--

INSERT INTO `user_group` (`id`, `nama`) VALUES
(1, 'admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `user_id_group` (`user_id_group`),
  ADD KEY `uuid_person` (`uuid_person`);

--
-- Indexes for table `keluarga`
--
ALTER TABLE `keluarga`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_status_keluarga` (`id_status_keluarga`);

--
-- Indexes for table `pekerjaan`
--
ALTER TABLE `pekerjaan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pendidikan`
--
ALTER TABLE `pendidikan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `penghasilan`
--
ALTER TABLE `penghasilan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `person`
--
ALTER TABLE `person`
  ADD PRIMARY KEY (`uuid`),
  ADD UNIQUE KEY `nik` (`nik`),
  ADD KEY `pembuat` (`pembuat`),
  ADD KEY `id_pendidikan` (`id_pendidikan`),
  ADD KEY `id_pekerjaan` (`id_pekerjaan`),
  ADD KEY `id_penghasilan` (`id_penghasilan`);

--
-- Indexes for table `status_keluarga`
--
ALTER TABLE `status_keluarga`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_group`
--
ALTER TABLE `user_group`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `keluarga`
--
ALTER TABLE `keluarga`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `pekerjaan`
--
ALTER TABLE `pekerjaan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT for table `pendidikan`
--
ALTER TABLE `pendidikan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `penghasilan`
--
ALTER TABLE `penghasilan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `status_keluarga`
--
ALTER TABLE `status_keluarga`
  MODIFY `id` tinyint(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `user_group`
--
ALTER TABLE `user_group`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `admin`
--
ALTER TABLE `admin`
  ADD CONSTRAINT `admin_ibfk_1` FOREIGN KEY (`user_id_group`) REFERENCES `user_group` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `keluarga`
--
ALTER TABLE `keluarga`
  ADD CONSTRAINT `keluarga_ibfk_1` FOREIGN KEY (`id_status_keluarga`) REFERENCES `status_keluarga` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
