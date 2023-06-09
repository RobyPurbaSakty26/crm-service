-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: mini_project
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actors`
--

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actors` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `role_id` int unsigned NOT NULL,
  `verified` enum('true','false') DEFAULT NULL,
  `active` enum('true','false') DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_actors_role` (`role_id`),
  CONSTRAINT `fk_actors_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

LOCK TABLES `actors` WRITE;
/*!40000 ALTER TABLE `actors` DISABLE KEYS */;
INSERT INTO `actors` VALUES (1,'JoeSU','1234',2,'true','true','2023-05-29 17:00:00','2023-05-29 17:00:00',NULL),(3,'Joe','',1,'true','true','2023-06-05 14:48:33','2023-06-05 14:48:33',NULL),(4,'Joe','123',1,'true','true','2023-06-05 14:48:58','2023-06-05 14:48:58',NULL),(5,'Joe','123',1,'false','false','2023-06-05 14:50:51','2023-06-05 14:50:51',NULL),(6,'Joe','123',1,'false','false','2023-06-05 14:52:03','2023-06-05 14:52:03',NULL),(7,'Joe','123',1,'false','false','2023-06-05 14:53:52','2023-06-05 14:53:52',NULL),(8,'Joe','$2a$10$VCpJlvw1ZIorlXtltsAI6OkH9mTcGecqOeHkDezAxdKdlJxuxBITa',1,'false','false','2023-06-05 15:04:47','2023-06-05 15:04:47',NULL),(9,'Joe','$2a$10$RoWGBBnWLNk2J8X/5GSv5ew.KNb5.KJTz9oXk17.E79itmkMbW77O',1,'false','false','2023-06-05 15:07:01','2023-06-05 15:07:01',NULL),(10,'Joe','$2a$10$qg9ATXq4Z47Y141Sbr1B6eq.5OY8zM7Fl5Qwf6G3WjCQHPZbIv4aG',1,'false','false','2023-06-05 15:07:25','2023-06-05 15:07:25',NULL),(11,'Joe','$2a$10$.9XTcYhvkFcwRHpTEONTg.YxP0rXNpnYnYggjEr16uCptU2UOuw8S',1,'false','true','2023-06-05 15:08:00','2023-06-05 15:08:00',NULL),(12,'Joe','$2a$10$3DoghSF3DrD6EMcqrQRA9ORPCcsvZbCa.h8mec1xObviFZPksHYXu',1,'false','true','2023-06-05 15:09:00','2023-06-05 15:09:00',NULL),(13,'Joe','$2a$10$G45TW21LJ2GKLdJGZxJVbengS5eSdVEWf53DGaor0beoMWLW/dv7C',1,'false','true','2023-06-05 15:09:20','2023-06-05 15:09:20',NULL),(14,'Joe','$2a$10$Cpiyszsu.BEmawSeOLjj2OB.Y.3rkE1d7cJjQ.HIMQ7fPkysktO2y',1,'false','true','2023-06-05 15:10:22','2023-06-05 15:10:22',NULL),(15,'Joe','$2a$10$ZwHeW0XyMLd2TF9.S09qUurGHB7j8jkP8H7gZw8m.zdMfP7VD.w6m',1,'false','true','2023-06-05 15:10:39','2023-06-05 15:10:39',NULL),(16,'Joe','$2a$10$2j0LWIYeotM9IFsj2OcWzOfQT0ARMRVwXnAp0Zs6wrwSvmvEt/3U.',1,'true','true','2023-06-05 15:11:33','2023-06-05 15:11:33',NULL),(17,'Joe','$2a$10$9T3XVLCLvPgDmPVoguLeo.9/Z3Nu0Z2Du5LXsGDSnubvKAMQka0Jq',1,'true','true','2023-06-05 15:12:04','2023-06-05 15:12:04',NULL),(18,'Joe','$2a$10$JdTwiMoCXs6XYlQB8fE0RO6ThVu2hFD1qRLJHYWJfybZoGyorntzq',1,'false','false','2023-06-05 15:12:30','2023-06-05 15:12:30',NULL),(19,'Joe','$2a$10$ZJjyKLW244KcUroBK7180OS69uQcEmOv0WfumtFWHKLD7hS9zS8EG',1,'false','false','2023-06-05 15:17:13','2023-06-05 15:17:13',NULL),(20,'Salman','$2a$10$2aGyAXYe6LS7KsGd5kL4w.Ri.zTCRp3b0LMU1yLGrTP9BzUwJygTe',1,'false','false','2023-06-05 16:32:37','2023-06-05 16:32:37',NULL),(22,'ali','$2a$10$BkWcHClUx5QMdT0X5uqImeKot2m3LIoszT7wK.i1XINU.g.p/IadK',1,'false','false','2023-06-05 17:27:21','2023-06-05 17:27:21',NULL),(23,'Ahmad','$2a$10$Xu1fszWZOu2VUuNDo8QARuxHdyVCEBLl9JqZhIWT4K/sfsINR3QP6',1,'false','true','2023-06-06 17:33:22','2023-06-08 20:08:55',NULL),(24,'super admin','$2a$10$HW/Bx.QtRaty/v05z9E0COIgkPLk1QyxqgA9256erZkaGTRA.cQom',2,'false','false','2023-06-06 17:59:01','2023-06-06 17:59:01',NULL),(25,'super admin','$2a$10$k20t7kVuMwHbcFHFK.sFu.EioUTA/t/cUoSXfFEajK9j6bRpfq8AW',2,'false','false','2023-06-08 06:04:16','2023-06-08 06:04:16',NULL),(26,'super admin','$2a$10$qyqMLMVsLL3LeccxcNWevO.0DV9NgF53NAA1ruE.tsr16UtupC.WO',2,'false','false','2023-06-08 10:22:12','2023-06-08 10:22:12',NULL),(27,'super admin','$2a$10$hN2HFubpj5Th5wa.fzxWM.3pgXDmRIHZbe38EohCtDeIYTj9HVjm6',2,'false','false','2023-06-08 20:09:01','2023-06-08 20:09:01',NULL);
/*!40000 ALTER TABLE `actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `first_name` varchar(20) DEFAULT NULL,
  `last_name` varchar(20) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  `avatar` varchar(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customers`
--

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers` DISABLE KEYS */;
INSERT INTO `customers` VALUES (1,'Doe ','joerss ','joe ','img.png ','2023-06-03 21:55:26','2023-06-04 03:13:12','2023-06-04 05:50:02'),(2,'Doe ','joerss new','joe ','img.png ','2023-06-03 22:16:58','2023-06-04 03:17:26','2023-06-04 05:50:32'),(3,NULL,'joerss','joe','img.png','2023-06-03 22:17:36','2023-06-03 22:17:36','2023-06-04 05:50:34'),(4,'','joerss','joe','img.png','2023-06-03 22:21:47','2023-06-03 22:21:47','2023-06-04 05:50:36'),(5,'Doe','joerss','joe','img.png','2023-06-03 22:22:12','2023-06-03 22:22:12','2023-06-04 05:50:38'),(6,'Doe','joerss','joe','img.png','2023-06-03 22:28:23','2023-06-03 22:28:23','2023-06-04 05:50:43'),(7,'Doe','joerss','joe','img.png','2023-06-03 23:39:38','2023-06-03 23:39:38','2023-06-04 05:50:45'),(8,'Doe','joerss','joe','img.png','2023-06-03 23:39:48','2023-06-03 23:39:48','2023-06-04 05:50:47'),(9,'Doe','joerss','joe','img.png','2023-06-03 23:39:49','2023-06-03 23:39:49','2023-06-04 05:50:51'),(10,'Doe','joerss','joe','img.png','2023-06-03 23:39:50','2023-06-03 23:39:50','2023-06-04 05:50:56'),(11,'Doe ','joerss new','joe ','img.png ','2023-06-03 23:39:55','2023-06-09 05:09:42',NULL),(12,'Doe','joerss','joe','img.png','2023-06-04 02:23:37','2023-06-04 02:23:37',NULL),(13,'Doe','joerss','joe','img.png','2023-06-04 02:24:01','2023-06-04 02:24:01',NULL),(14,'Doe','joerss','joe','img.png','2023-06-04 02:24:56','2023-06-04 02:24:56',NULL),(15,'Doe','joerss','joe','img.png','2023-06-04 02:44:41','2023-06-04 02:44:41',NULL),(16,'Doe','joerss','joe','img.png','2023-06-04 02:47:10','2023-06-04 02:47:10',NULL),(17,'Doe','joerss','joe','img.png','2023-06-04 03:00:37','2023-06-04 03:00:37',NULL),(18,'Doe','joerss new','joe','img.png new','2023-06-04 03:09:27','2023-06-04 03:09:27',NULL),(19,'Doe','joerss new','joe','img.png new','2023-06-04 03:09:46','2023-06-04 03:09:46',NULL),(20,'Doe ','joerss new','joe ','img.png ','2023-06-04 03:10:00','2023-06-08 20:39:50','2023-06-08 20:40:47'),(21,'Doe ','joerss new','joe ','img.png ','2023-06-03 21:55:26','2023-06-04 03:14:02',NULL),(22,'Doe ','joerss new','joe ','img.png ','2023-06-03 22:16:58','2023-06-04 03:14:14',NULL),(23,'Doe','joerss emai','joe@mail.com','img.png','2023-06-04 16:56:24','2023-06-04 16:56:24',NULL),(24,'Doe','joerss emai','king@mail.com','img.png','2023-06-04 17:19:48','2023-06-04 17:19:48',NULL),(25,'Doe','joerss emai','doeg@mail.com','img.png','2023-06-04 17:26:47','2023-06-04 17:26:47',NULL),(26,'Doe','joerss emai','king@mail.com','img.png','2023-06-04 18:14:05','2023-06-04 18:14:05',NULL),(27,'Doe','joerss emai','king@mail.com','img.png','2023-06-04 18:14:53','2023-06-04 18:14:53',NULL),(28,'Doe','joerss emai','king@mail.com','img.png','2023-06-04 18:14:54','2023-06-04 18:14:54',NULL),(29,'Doe','joerss emai','king@mail.com','img.png','2023-06-06 18:56:38','2023-06-06 18:56:38',NULL),(30,'Doe','joerss emai','king@mail.com','img.png','2023-06-08 06:04:02','2023-06-08 06:04:02',NULL),(31,'Doe','joerss emai','king@mail.com','img.png','2023-06-08 20:39:14','2023-06-08 20:39:14',NULL),(32,'Doe','joerss emai','king@mail.com','img.png','2023-06-09 05:09:30','2023-06-09 05:09:30',NULL);
/*!40000 ALTER TABLE `customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `register`
--

DROP TABLE IF EXISTS `register`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `register` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `admin_id` bigint unsigned DEFAULT NULL,
  `super_admin_id` bigint unsigned DEFAULT NULL,
  `status` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_register_actors` (`admin_id`),
  CONSTRAINT `fk_register_actors` FOREIGN KEY (`admin_id`) REFERENCES `actors` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `register`
--

LOCK TABLES `register` WRITE;
/*!40000 ALTER TABLE `register` DISABLE KEYS */;
/*!40000 ALTER TABLE `register` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'Admin'),(2,'Super Admin');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-09 20:15:31
