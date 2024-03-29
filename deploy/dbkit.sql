-- mysql -h127.0.0.1 -P10007 -uroot -p dbkit < dbkit.sql

-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: dbkit
-- ------------------------------------------------------
-- Server version	8.0.28-0ubuntu0.20.04.3

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
-- Table structure for table `target_dsn`
--

DROP TABLE IF EXISTS `target_dsn`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `target_dsn` (
  `tid` int NOT NULL AUTO_INCREMENT,
  `db_type` varchar(50) NOT NULL,
  `db_host` varchar(100) NOT NULL,
  `db_port` int NOT NULL,
  `db_user` varchar(100) NOT NULL,
  `db_pwd` varchar(100) NOT NULL DEFAULT '',
  `db_name` varchar(100) DEFAULT NULL,
  `params` varchar(200) NOT NULL DEFAULT '',
  `state` int DEFAULT '0',
  `version` varchar(100) DEFAULT '-',
  `deleted` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`tid`),
  UNIQUE KEY `target_dsn_tid_uindex` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `target_dsn`
--

LOCK TABLES `target_dsn` WRITE;
/*!40000 ALTER TABLE `target_dsn` DISABLE KEYS */;
INSERT INTO `target_dsn` VALUES (1,'mysql','127.0.0.1',3306,'root','tobeno.1','test','',1,'8.0.28-0ubuntu0.20.04.3',0),(2,'mysql','127.0.0.1',3306,'dinary','tobeno.1','test','',0,'-',0),(3,'mysql','127.0.0.1',3306,'dinary','tobeno.1','test','',0,'-',0),(4,'tidb','127.0.0.1',4000,'root','','test','',1,'5.7.25-TiDB-v5.4.0',0),(5,'tidb','localhost',4000,'root','','test','',0,'-',0),(6,'mariadb','127.0.0.1',3307,'root','tobeno.1','test','',0,'-',0),(7,'mariadb','locahost',3308,'root','','test','',0,'-',0),(8,'mariadb','192.168.0.1',3306,'root','test','test','',0,'-',0);
/*!40000 ALTER TABLE `target_dsn` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_job`
--

DROP TABLE IF EXISTS `test_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_job` (
  `jid` int NOT NULL AUTO_INCREMENT,
  `dsn` varchar(300) DEFAULT NULL,
  `db_name` varchar(50) DEFAULT NULL,
  `target` varchar(50) DEFAULT NULL,
  `oracle` varchar(20) DEFAULT NULL,
  `state` int DEFAULT NULL,
  `time_limit` float DEFAULT NULL,
  `comments` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`jid`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_job`
--

LOCK TABLES `test_job` WRITE;
/*!40000 ALTER TABLE `test_job` DISABLE KEYS */;
INSERT INTO `test_job` VALUES (1,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0,'模拟','2022-05-12 08:34:29',0),(2,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,1.5,'模拟','2022-05-12 08:34:29',0),(3,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',1,1,'模拟','2022-05-12 08:34:29',0),(4,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',-1,2,'模拟','2022-05-12 08:34:29',0),(5,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(6,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(7,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(8,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(9,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(10,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(11,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(12,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(13,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(14,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(15,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(16,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(17,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(18,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(19,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(20,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(21,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0),(22,'user:password@tcp(127.0.0.1:3306)/dbname','dbname','mysql','query',2,0.5,'模拟','2022-05-12 08:34:29',0);
/*!40000 ALTER TABLE `test_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_report`
--

DROP TABLE IF EXISTS `test_report`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_report` (
  `rid` int NOT NULL AUTO_INCREMENT,
  `jid` int NOT NULL,
  `input_stmt` text,
  `input_res` text,
  `oracle_stmt` text,
  `oracle_res` text,
  `category` varchar(50) DEFAULT NULL,
  `report_time` timestamp NULL DEFAULT NULL,
  `state` varchar(20) DEFAULT NULL,
  `url` text,
  `comments` text,
  `deleted` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_report`
--

LOCK TABLES `test_report` WRITE;
/*!40000 ALTER TABLE `test_report` DISABLE KEYS */;
INSERT INTO `test_report` VALUES (1,1,'Order:[1-0, 1-1, 2-0, 2-1, 2-2, 1-2, 1-3, 1-4, 1-5]\nQuery Results:\n	1-0: null\n	1-1: [0, 0, 1, 1, 2, 2]\n	2-0: null\n	2-1: null\n	2-2: null\n	1-2: [0, 0, 1, 1, 2, 2]\n	1-3: null\n	1-4: [10, 0, 10, 1, 10, 2]\n	1-5: null\nFinalState: [10, 0, 10, 1, 10, 2]\nDeadBlock: false','CREATE TABLE t(a INT, b INT);\nINSERT INTO t VALUES(NULL, 1), (2, 2), (NULL, 3)\n\n -- TRANSACTION 1 [ReadCommitted];\nbegin;\nupdate t set a = 10 where 1;\ncommit;\n\n -- TRANSACTION 2 [ReadCommitted];\nbegin;\nupdate t set b = 20 where a;\ncommit;\n\n','Order:[1-0, 1-1, 2-0, 2-1(B), 1-2, 2-1, 2-2]\nQuery Results:\n	1-0: null\n	1-1: null\n	2-0: null\n	2-1: null\n	1-2: null\n	2-1: null\n	2-2: null\nFinalState: [10, 1, 10, 20, 10, 20]\nDeadBlock: false','CREATE TABLE t(a int, b int)\nINSERT INTO t VALUES(0, 0), (1, 1), (2, 2)\n\n -- TRANSACTION 1 [RepeatableRead];\nBEGIN\nSELECT * FROM t\nSELECT * FROM t\nUPDATE t SET a = 10\nSELECT * FROM t\nCOMMIT\n\n -- TRANSACTION 2 [RepeatableRead];\nBEGIN\nUPDATE t SET a = 10 WHERE b = 1\nCOMMIT','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(2,1,'CREATE TABLE t(a INT, b INT);\nINSERT INTO t VALUES(NULL, 1), (2, 2), (NULL, 3)\n\n -- TRANSACTION 1 [ReadCommitted];\nbegin;\nupdate t set a = 10 where 1;\ncommit;\n\n -- TRANSACTION 2 [ReadCommitted];\nbegin;\nupdate t set b = 20 where a;\ncommit;\n\n','Order:[1-0, 1-1, 2-0, 2-1(B), 1-2, 2-1, 2-2]\nQuery Results:\n	1-0: null\n	1-1: null\n	2-0: null\n	2-1: null\n	1-2: null\n	2-1: null\n	2-2: null\nFinalState: [10, 1, 10, 20, 10, 20]\nDeadBlock: false','CREATE TABLE t(a int, b int)\nINSERT INTO t VALUES(0, 0), (1, 1), (2, 2)\n\n -- TRANSACTION 1 [RepeatableRead];\nBEGIN\nSELECT * FROM t\nSELECT * FROM t\nUPDATE t SET a = 10\nSELECT * FROM t\nCOMMIT\n\n -- TRANSACTION 2 [RepeatableRead];\nBEGIN\nUPDATE t SET a = 10 WHERE b = 1\nCOMMIT','Order:[1-0, 1-1, 2-0, 2-1, 2-2, 1-2, 1-3, 1-4, 1-5]\nQuery Results:\n	1-0: null\n	1-1: [0, 0, 1, 1, 2, 2]\n	2-0: null\n	2-1: null\n	2-2: null\n	1-2: [0, 0, 1, 1, 2, 2]\n	1-3: null\n	1-4: [10, 0, 10, 1, 10, 2]\n	1-5: null\nFinalState: [10, 0, 10, 1, 10, 2]\nDeadBlock: false','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(3,2,'input1','input_res1','oracle1','oracle_res1','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(4,2,'input1','input_res1','oracle1','oracle_res1','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(5,3,'input1','input_res1','oracle1','oracle_res1','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(6,5,'input1','input_res1','oracle1','oracle_res1','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(7,5,'input1','input_res1','oracle1','oracle_res1','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0),(8,5,'input1','input_res1','oracle1','oracle_res1','tlp','2022-05-12 08:36:44','verified','http://www.baidu.com','模拟',0);
/*!40000 ALTER TABLE `test_report` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_statistic`
--

DROP TABLE IF EXISTS `test_statistic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_statistic` (
  `sid` int NOT NULL AUTO_INCREMENT,
  `jid` int NOT NULL,
  `sql_count` int DEFAULT NULL,
  `case_count` int DEFAULT NULL,
  `report_count` int DEFAULT NULL,
  `fail_cause` text,
  `end_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_statistic`
--

LOCK TABLES `test_statistic` WRITE;
/*!40000 ALTER TABLE `test_statistic` DISABLE KEYS */;
INSERT INTO `test_statistic` VALUES (1,1,100,20,2,'\"\"','2022-05-12 08:38:55'),(2,2,100,20,2,'\"\"','2022-05-12 08:38:55'),(3,3,100,20,1,'\"\"','2022-05-12 08:38:55'),(4,4,100,20,0,'\"\"','2022-05-12 08:38:55'),(5,5,100,20,3,'\"\"','2022-05-12 08:38:55');
/*!40000 ALTER TABLE `test_statistic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `verify_job`
--

DROP TABLE IF EXISTS `verify_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `verify_job` (
  `jid` int NOT NULL AUTO_INCREMENT,
  `dsn` varchar(100) DEFAULT NULL,
  `db_name` varchar(50) DEFAULT NULL,
  `target` varchar(20) DEFAULT NULL,
  `model` varchar(20) DEFAULT NULL,
  `op` int DEFAULT NULL,
  `state` varchar(20) DEFAULT NULL,
  `comments` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`jid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `verify_job`
--

LOCK TABLES `verify_job` WRITE;
/*!40000 ALTER TABLE `verify_job` DISABLE KEYS */;
INSERT INTO `verify_job` VALUES (1,'user:password@tcp(127.0.0.1:3306)/dbname','db_name','mysql','register',100,'2','模拟','2022-05-12 08:49:41',0),(2,'user:password@tcp(127.0.0.1:3306)/dbname','db_name','mysql','register',100,'2','模拟','2022-05-12 08:49:41',0),(3,'user:password@tcp(127.0.0.1:3306)/dbname','db_name','mysql','register',100,'2','模拟','2022-05-12 08:49:41',0),(4,'user:password@tcp(127.0.0.1:3306)/dbname','db_name','mysql','register',100,'2','模拟','2022-05-12 08:49:41',0),(5,'user:password@tcp(127.0.0.1:3306)/dbname','db_name','mysql','register',100,'1','模拟','2022-05-12 08:49:41',0);
/*!40000 ALTER TABLE `verify_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `verify_report`
--

DROP TABLE IF EXISTS `verify_report`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `verify_report` (
  `rid` int NOT NULL AUTO_INCREMENT,
  `jid` int DEFAULT NULL,
  `pass` int DEFAULT NULL,
  `file_path` varchar(100) DEFAULT NULL,
  `comments` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`rid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `verify_report`
--

LOCK TABLES `verify_report` WRITE;
/*!40000 ALTER TABLE `verify_report` DISABLE KEYS */;
INSERT INTO `verify_report` VALUES (1,1,1,'/1.html','模拟','2022-05-12 08:54:05',0),(2,2,0,'/2.html','模拟','2022-05-12 08:54:40',0),(3,3,-1,'/3.html','模拟','2022-05-12 08:54:40',0),(4,4,1,'/4.html','模拟','2022-05-12 08:54:40',0);
/*!40000 ALTER TABLE `verify_report` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-16 18:58:12
