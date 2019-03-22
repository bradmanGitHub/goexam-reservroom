-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.2.7-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------


-- Dumping database structure for bookingdb
CREATE DATABASE IF NOT EXISTS `bookingdb` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `bookingdb`;

-- Dumping structure for table bookingdb.tb_booking
CREATE TABLE IF NOT EXISTS `tb_booking` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(500) NOT NULL,
  `ROOM` varchar(500) NOT NULL,
  `START` datetime NOT NULL,
  `END` datetime NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
