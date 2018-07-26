-- MySQL dump 10.13  Distrib 5.7.22, for Linux (x86_64)
--
-- Host: 18.188.151.231    Database: retrobarbershop_app
-- ------------------------------------------------------
-- Server version	5.7.22-0ubuntu0.16.04.1

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
  `notif` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `absen_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `absen`
--

LOCK TABLES `absen` WRITE;
/*!40000 ALTER TABLE `absen` DISABLE KEYS */;
INSERT INTO `absen` VALUES (80,'2018-07-22','23:48:57','hadir',26,'Sunday','-6.1793658','106.7348319','082188352121','1'),(81,'2018-07-22','23:49:47','hadir',27,'Sunday','-6.1793937','106.7348369','admin','1'),(82,'2018-07-23','0:04:48','hadir',26,'Monday','-6.1793658','106.7348319','082188352121','1'),(83,'2018-07-23','1:32:34','hadir',26,'Monday','-6.1793658','106.7348319','082188352121','1'),(84,'2018-07-23','1:33:02','hadir',26,'Monday','-6.1793658','106.7348319','082188352121','1'),(85,'2018-07-23','19:49:13','hadir',26,'Monday','-6.1753749','106.7381895','082188352121','1'),(86,'2018-07-23','19:59:00','hadir',26,'Monday','-6.1793658','106.7348319','082188352121','1'),(87,'2018-07-23','19:59:13','hadir',27,'Monday','-6.177623','106.7363024','admin','1'),(88,'2018-07-23','22:14:15','hadir',27,'Monday','-6.1793658','106.7348319','admin','1'),(89,'2018-07-25','0:20:24','hadir',26,'Wednesday','-6.1793658','106.7348319','082188352121','1'),(90,'2018-07-25','0:30:42','hadir',27,'Wednesday','-6.1793658','106.7348319','superadmin','1'),(91,'2018-07-25','0:31:58','hadir',27,'Wednesday','-6.1793658','106.7348319','superadmin','1');
/*!40000 ALTER TABLE `absen` ENABLE KEYS */;
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
  `usergroup` varchar(45) DEFAULT NULL,
  `user_retro` varchar(45) DEFAULT NULL,
  `password_retro` varchar(45) DEFAULT NULL,
  `latitude` varchar(45) DEFAULT NULL,
  `longtitude` varchar(45) DEFAULT NULL,
  `foto` varchar(45) DEFAULT NULL,
  `nama_foto` varchar(40) DEFAULT 'null',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (26,'andrian latif','1','082188352121','ian123456',NULL,NULL,NULL,'1532501809432.jpg'),(27,'sysadmin','1','superadmin','123456',NULL,NULL,NULL,'1532350956151.jpg');
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

-- Dump completed on 2018-07-25 14:09:00
