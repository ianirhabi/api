-- MySQL dump 10.13  Distrib 5.7.22, for Linux (x86_64)
--
-- Host: 192.168.88.252    Database: retrobarbershop_app
-- ------------------------------------------------------
-- Server version	5.7.22-0ubuntu18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `absen`
--

DROP TABLE IF EXISTS `absen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `absen` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `tanggal` date DEFAULT NULL,
  `waktu` varchar(50) DEFAULT NULL,
  `kehadiran` varchar(60) DEFAULT NULL,
  `id_user` int(5) DEFAULT NULL,
  `hari` varchar(45) DEFAULT NULL,
  `lat` varchar(45) DEFAULT NULL,
  `long` varchar(45) DEFAULT NULL,
  `user` varchar(45) DEFAULT NULL,
  `alasan` varchar(45) DEFAULT NULL,
  `pesan` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `absen_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `absen`
--

LOCK TABLES `absen` WRITE;
/*!40000 ALTER TABLE `absen` DISABLE KEYS */;
INSERT INTO `absen` VALUES (1,'2018-07-29','18:28:36','hadir',29,'Minggu','-6.1793937','106.7348369','082188352121','1',NULL),(2,'2018-07-29','23:58:53','hadir',29,'Minggu','-6.1793658','106.7348319','082188352121','1',NULL),(4,'2018-07-31','20:46:06','hadir',62,'Tuesday','-6.1793875','106.7348358','sanrish','1',NULL),(5,'2018-07-31','20:49:22','hadir',62,'Tuesday','-6.1787078','106.7354212','sanrish','1',NULL),(6,'2018-07-31','123456','izin',29,'Selasa','4.44343','4.6767','082188352121','sakit',NULL),(7,'2018-07-31','123456','izin',29,'Selasa','4.44343','4.6767','082188352121','acar keluarga',NULL),(10,'2018-07-31','Tuesday 23:53:20 2018-07-31','izin',29,'Tuesday','-6.1793658','106.7348319','082188352121','sakit',NULL),(11,'2018-08-01','0:02:32','izin',29,'Wednesday','-6.1793658','106.7348319','082188352121','suatu hal lain-lain',NULL),(12,'2018-08-01','0:27:54','hadir',64,'Wednesday','-6.1793658','106.7348319','082188352121','',NULL),(13,'2018-08-01','0:28:30','izin',64,'Wednesday','-6.1793658','106.7348319','082188352121','sakit',NULL);
/*!40000 ALTER TABLE `absen` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deactive_user`
--

DROP TABLE IF EXISTS `deactive_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `deactive_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deactive_user`
--

LOCK TABLES `deactive_user` WRITE;
/*!40000 ALTER TABLE `deactive_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `deactive_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(45) NOT NULL,
  `notifikasi` varchar(45) DEFAULT NULL,
  `user_retro` varchar(45) DEFAULT NULL,
  `password_retro` varchar(45) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longtitude` varchar(45) DEFAULT NULL,
  `foto` varchar(45) DEFAULT NULL,
  `nama_foto` varchar(40) DEFAULT NULL,
  `usergroup` varchar(45) DEFAULT NULL,
  `ianmonitor` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `created` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (27,'rio latif','0','superadmin','123456',NULL,NULL,NULL,'1533046031024.jpg','1','',NULL,NULL),(28,'safira','0','admin','safira123',NULL,NULL,NULL,'1533043129417.jpg','2','',NULL,NULL),(29,'Programmer Jalanan','0','082188352121','123456','-6.1793658','106.7348319',NULL,'1533050213794.jpg','3','1',NULL,NULL),(62,'Akun Demo','0','sanrish','123456','-6.1787078','106.7354212',NULL,'1533044671213.jpg','3','1',NULL,NULL),(64,'visa akun','0','082188352121','astaga','-6.1793658','106.7348319',NULL,'1533045750673.jpg','3','1',NULL,NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-07-31 13:57:42
