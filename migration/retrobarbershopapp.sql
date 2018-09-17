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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `absen`
--

LOCK TABLES `absen` WRITE;
/*!40000 ALTER TABLE `absen` DISABLE KEYS */;
INSERT INTO `absen` VALUES (4,'2018-07-31','20:46:06','hadir',62,'Tuesday','-6.1793875','106.7348358','sanrish','1',NULL),(5,'2018-07-31','20:49:22','hadir',62,'Tuesday','-6.1787078','106.7354212','sanrish','1',NULL),(6,'2018-07-31','123456','izin',29,'Selasa','4.44343','4.6767','082188352121','sakit',NULL),(7,'2018-07-31','123456','izin',29,'Selasa','4.44343','4.6767','082188352121','acar keluarga',NULL),(10,'2018-07-31','Tuesday 23:53:20 2018-07-31','izin',29,'Tuesday','-6.1793658','106.7348319','082188352121','sakit',NULL),(11,'2018-08-01','0:02:32','izin',29,'Wednesday','-6.1793658','106.7348319','082188352121','suatu hal lain-lain',NULL),(12,'2018-08-01','0:27:54','hadir',64,'Wednesday','-6.1793658','106.7348319','082188352121','',NULL),(13,'2018-08-01','0:28:30','izin',64,'Wednesday','-6.1793658','106.7348319','082188352121','sakit',NULL),(14,'2018-08-04','15:42:26','hadir',62,'Saturday','-6.1793937','106.7348369','sanrish','',NULL),(15,'2018-08-04','15:58:24','hadir',62,'Saturday','22.107531993662175','44.81406499002955','sanrish','',NULL);
/*!40000 ALTER TABLE `absen` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_barang`
--

DROP TABLE IF EXISTS `item_barang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_barang` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item_category` varchar(45) DEFAULT NULL,
  `code_item` int(20) DEFAULT NULL,
  `created` date DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `item_barang_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25220 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_barang`
--

LOCK TABLES `item_barang` WRITE;
/*!40000 ALTER TABLE `item_barang` DISABLE KEYS */;
INSERT INTO `item_barang` VALUES (25155,'kartu tri',2,'2018-09-15',27),(25156,'hs',0,'2018-09-15',27),(25157,'hs',1,'2018-09-15',27),(25158,'hs',0,'2018-09-15',27),(25159,'hs',1,'2018-09-15',27),(25160,'hs',2,'2018-09-15',27),(25161,'hs',3,'2018-09-15',27),(25162,'hs',4,'2018-09-15',27),(25163,'hs',5,'2018-09-15',27),(25164,'hs',6,'2018-09-15',27),(25165,'hs',7,'2018-09-15',27),(25166,'hs',8,'2018-09-15',27),(25167,'hs',9,'2018-09-15',27),(25168,'hs',10,'2018-09-15',27),(25169,'hs',11,'2018-09-15',27),(25170,'hs',12,'2018-09-15',27),(25171,'hs',13,'2018-09-15',27),(25172,'hs',14,'2018-09-15',27),(25173,'hs',15,'2018-09-15',27),(25174,'hs',16,'2018-09-15',27),(25175,'hs',17,'2018-09-15',27),(25176,'hs',18,'2018-09-15',27),(25177,'hs',19,'2018-09-15',27),(25178,'hs',20,'2018-09-15',27),(25179,'hs',21,'2018-09-15',27),(25180,'hs',22,'2018-09-15',27),(25181,'hs',23,'2018-09-15',27),(25182,'hs',24,'2018-09-15',27),(25183,'hs',25,'2018-09-15',27),(25184,'hs',26,'2018-09-15',27),(25185,'hs',27,'2018-09-15',27),(25186,'hs',28,'2018-09-15',27),(25187,'hs',29,'2018-09-15',27),(25188,'hs',30,'2018-09-15',27),(25189,'hs',31,'2018-09-15',27),(25190,'hs',32,'2018-09-15',27),(25191,'hs',33,'2018-09-15',27),(25192,'hs',34,'2018-09-15',27),(25193,'hs',35,'2018-09-15',27),(25194,'hs',36,'2018-09-15',27),(25195,'hs',37,'2018-09-15',27),(25196,'hs',38,'2018-09-15',27),(25197,'hs',39,'2018-09-15',27),(25198,'hs',40,'2018-09-15',27),(25199,'hs',41,'2018-09-15',27),(25200,'hs',42,'2018-09-15',27),(25201,'hs',43,'2018-09-15',27),(25202,'hs',44,'2018-09-15',27),(25203,'hs',45,'2018-09-15',27),(25204,'hs',46,'2018-09-15',27),(25205,'hs',47,'2018-09-15',27),(25206,'hs',48,'2018-09-15',27),(25207,'hs',49,'2018-09-15',27),(25208,'hs',50,'2018-09-15',27),(25209,'hs',51,'2018-09-15',27),(25210,'hs',52,'2018-09-15',27),(25211,'hs',53,'2018-09-15',27),(25212,'hs',54,'2018-09-15',27),(25213,'hs',55,'2018-09-15',27),(25214,'hs',56,'2018-09-15',27),(25215,'hs',57,'2018-09-15',27),(25216,'hs',58,'2018-09-15',27),(25217,'hs',62,'2018-09-16',27),(25218,'hs',60,'2018-09-15',27),(25219,'hs',61,'2018-09-15',27);
/*!40000 ALTER TABLE `item_barang` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_barang_detail`
--

DROP TABLE IF EXISTS `item_barang_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_barang_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(45) DEFAULT NULL,
  `stock` varchar(45) DEFAULT NULL,
  `harga_pokok_barang` int(11) DEFAULT NULL,
  `harga_jual` int(11) DEFAULT NULL,
  `code_category` int(10) DEFAULT NULL,
  `deskripsi` varchar(45) DEFAULT NULL,
  `created` date DEFAULT NULL,
  `update` date DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `link_barang_detail_ibfk_2` (`code_category`),
  CONSTRAINT `item_barang_detail_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `link_barang_detail_ibfk_2` FOREIGN KEY (`code_category`) REFERENCES `item_barang` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=392 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_barang_detail`
--

LOCK TABLES `item_barang_detail` WRITE;
/*!40000 ALTER TABLE `item_barang_detail` DISABLE KEYS */;
INSERT INTO `item_barang_detail` VALUES (38,'5','100',200,2000,25155,'bandar','2018-09-15','2018-09-15',27),(39,'12','1',100,1000,25155,'sangat','2018-09-15','2018-09-15',27),(40,'7','100',1,12,25155,'ying','2018-09-15','2018-09-15',27),(41,'11','13',1,21,25155,'sandar','2018-09-15','2018-09-15',27),(42,'34','11',10000,12000,25155,'jangan','2018-09-15','2018-09-15',27),(43,'152','12',5464,4654,25155,'bersama','2018-09-15','2018-09-15',27),(44,'25','1',81,318,25155,'fi','2018-09-15','2018-09-15',27),(45,'111','1',4846,4615,25155,'hwfh','2018-09-15','2018-09-15',27),(46,'gwh','1',1,111,25155,'52','2018-09-15','2018-09-15',27),(47,'1t','46',64,5646,25155,'vhaja','2018-09-15','2018-09-15',27),(48,'6161','4545',846,45548,25155,'gycyc','2018-09-15','2018-09-15',27),(49,'qggq','543',4843,1218,25155,'vqh','2018-09-15','2018-09-15',27),(50,'1t1t','84643',48464,4546,25155,'vahwu','2018-09-15','2018-09-15',27),(51,'qfqyq','545454',84843,484864,25155,'vwgq','2018-09-15','2018-09-15',27),(52,'15151','12518',84545,4848,25155,'cqgqy','2018-09-15','2018-09-15',27),(53,'25t','12',45434,79737,25155,'bansa','2018-09-15','2018-09-15',27),(54,'taya','15',576767,5767,25155,'gsys','2018-09-15','2018-09-15',27),(55,'151515','12188',1854,6464,25155,'gsus','2018-09-15','2018-09-15',27),(56,'1tq6tq','11',515,45484,25155,'yag61g','2018-09-15','2018-09-15',27),(57,'2ywg','8484',9484,9494,25155,'dhdhs','2018-09-15','2018-09-15',27),(58,'wywyw','1618',9494,6494,25155,'s svs','2018-09-15','2018-09-15',27),(59,'wgwgw','94840',4949,9494,25155,'vsbsb','2018-09-15','2018-09-15',27),(60,'avwvwv','946',9494,4,25155,'vavag','2018-09-15','2018-09-15',27),(61,'2gwg','49',94949,9494,25155,'vava','2018-09-15','2018-09-15',27),(62,'avaga','9',94949,9494948,25155,'agaga','2018-09-15','2018-09-15',27),(63,'25w','54',546,646,25155,'bxb','2018-09-15','2018-09-15',27),(64,'gigig','12',8386838,8686835,25155,'gi','2018-09-15','2018-09-15',27),(65,'1525','2',8757,9767,25155,'basa','2018-09-15','2018-09-15',27),(390,'1','100',2500,5000,25219,'pisang','2018-09-17','2018-09-17',27),(391,'jdjdjd','98679',98979,98979,25219,'xjxjxj','2018-09-17','2018-09-17',27);
/*!40000 ALTER TABLE `item_barang_detail` ENABLE KEYS */;
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
INSERT INTO `user` VALUES (27,'rio latif','0','superadmin','123456',NULL,NULL,NULL,'1533046031024.jpg','1','',NULL,NULL),(28,'safira','0','admin','admin',NULL,NULL,NULL,'1533043129417.jpg','3','',NULL,NULL),(29,'Programmer Jalanan','0','082188352121','123456','-6.1793658','106.7348319',NULL,'1533050213794.jpg','3','1',NULL,NULL),(62,'Akun Demo','0','sanrish','123456','22.107531993662175','44.81406499002955',NULL,'1533371637489.jpg','3','1',NULL,NULL),(64,'visa akun','0','082188352121','astaga','-6.1793658','106.7348319',NULL,'1533045750673.jpg','3','1',NULL,NULL);
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

-- Dump completed on 2018-09-16 22:26:33
