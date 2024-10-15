-- MySQL dump 10.13  Distrib 8.0.32, for macos13 (arm64)
--
-- Host: 127.0.0.1    Database: Security
-- ------------------------------------------------------
-- Server version	8.0.32

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
-- Current Database: `Security`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `Security` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `Security`;

--
-- Table structure for table `Agent`
--

DROP TABLE IF EXISTS `Agent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Agent` (
  `id` bigint NOT NULL,
  `username` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `tel_num` varchar(45) DEFAULT NULL,
  `create` int DEFAULT NULL,
  `salt` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `amount` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  `role` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Agent`
--

LOCK TABLES `Agent` WRITE;
/*!40000 ALTER TABLE `Agent` DISABLE KEYS */;
INSERT INTO `Agent` VALUES (964522009411117057,'yunfeng','81fc5d761bf9beff44fedc9fc022a66c','12312312345',1713188766,'DDnngGKkpFQBUuaozsFRXVSrjIqxboIH','111@222.333',0,1,'agent');
/*!40000 ALTER TABLE `Agent` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `BankCard`
--

DROP TABLE IF EXISTS `BankCard`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `BankCard` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `bcard_num` varchar(45) DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_id_idx` (`user_id`),
  CONSTRAINT `fk_userid_4` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `BankCard`
--

LOCK TABLES `BankCard` WRITE;
/*!40000 ALTER TABLE `BankCard` DISABLE KEYS */;
/*!40000 ALTER TABLE `BankCard` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Bill_Details`
--

DROP TABLE IF EXISTS `Bill_Details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Bill_Details` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `amount` bigint DEFAULT NULL,
  `service_id` bigint DEFAULT NULL,
  `service` varchar(45) DEFAULT NULL,
  `payee` varchar(45) DEFAULT NULL,
  `payer` varchar(45) DEFAULT NULL,
  `explain` varchar(45) DEFAULT NULL,
  `create` int DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `agent` varchar(45) DEFAULT NULL,
  `ddos_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_userid_7_idx` (`user_id`),
  KEY `fk_serviceid_idx` (`service_id`),
  KEY `fk_serviceid2_idx` (`ddos_id`),
  CONSTRAINT `fk_serviceid` FOREIGN KEY (`service_id`) REFERENCES `ScdnService` (`id`),
  CONSTRAINT `fk_serviceid2` FOREIGN KEY (`ddos_id`) REFERENCES `DDoSService` (`id`),
  CONSTRAINT `fk_userid_7` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Bill_Details`
--

LOCK TABLES `Bill_Details` WRITE;
/*!40000 ALTER TABLE `Bill_Details` DISABLE KEYS */;
INSERT INTO `Bill_Details` VALUES (974622774388088833,957493600380858369,0,974622695435517953,NULL,'','ccc','',1715596977,'0','',NULL);
/*!40000 ALTER TABLE `Bill_Details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `CloudEyeList`
--

DROP TABLE IF EXISTS `CloudEyeList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `CloudEyeList` (
  `id` bigint NOT NULL,
  `ce_id` bigint DEFAULT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `content` varchar(45) DEFAULT NULL,
  `task_type` json DEFAULT NULL,
  `monitor_type` json DEFAULT NULL,
  `start_count` bigint DEFAULT NULL,
  `order_money` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  `sell_status` int DEFAULT NULL,
  `zyks_money` bigint DEFAULT NULL,
  `zyzk_money` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_ceid_idx` (`ce_id`),
  CONSTRAINT `fk_ceid` FOREIGN KEY (`ce_id`) REFERENCES `DexunCloudEyeList` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CloudEyeList`
--

LOCK TABLES `CloudEyeList` WRITE;
/*!40000 ALTER TABLE `CloudEyeList` DISABLE KEYS */;
/*!40000 ALTER TABLE `CloudEyeList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ConfigList`
--

DROP TABLE IF EXISTS `ConfigList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ConfigList` (
  `id` bigint NOT NULL,
  `domain_id` bigint DEFAULT NULL,
  `load_balancing` varchar(45) DEFAULT NULL,
  `overload_redirect_url` varchar(45) DEFAULT NULL,
  `overload_status_code` varchar(45) DEFAULT NULL,
  `overload_type` varchar(45) DEFAULT NULL,
  `port` varchar(45) DEFAULT NULL,
  `protocol` varchar(45) DEFAULT NULL,
  `redirect` varchar(45) DEFAULT NULL,
  `server` varchar(45) DEFAULT NULL,
  `uri_forward` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dmid_idx` (`domain_id`),
  CONSTRAINT `fk_dmid` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ConfigList`
--

LOCK TABLES `ConfigList` WRITE;
/*!40000 ALTER TABLE `ConfigList` DISABLE KEYS */;
INSERT INTO `ConfigList` VALUES (977017044195274753,977016768593276929,'0','','','1','','http','false','','');
/*!40000 ALTER TABLE `ConfigList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DDoSCombos`
--

DROP TABLE IF EXISTS `DDoSCombos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DDoSCombos` (
  `id` bigint NOT NULL,
  `combo_id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  `yms` varchar(45) DEFAULT NULL,
  `ccfy` varchar(45) DEFAULT NULL,
  `gjwaf` varchar(45) DEFAULT NULL,
  `ywll` varchar(45) DEFAULT NULL,
  `gjcs` varchar(45) DEFAULT NULL,
  `pro_flow` bigint DEFAULT NULL,
  `ddos_hh` varchar(45) DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `complete_state` bigint DEFAULT NULL,
  `firewall_state` bigint DEFAULT NULL,
  `waf_state` bigint DEFAULT NULL,
  `group_type` bigint DEFAULT NULL,
  `zyks_money` bigint DEFAULT NULL,
  `zyzk_money` bigint DEFAULT NULL,
  `sell_status` int DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `port_num` bigint DEFAULT NULL,
  `ywdk_num` bigint DEFAULT NULL,
  `xl` varchar(45) DEFAULT NULL,
  `zfdks` varchar(45) DEFAULT NULL,
  `fhyms` varchar(45) DEFAULT NULL,
  `ywdk` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dxddosid_idx` (`combo_id`),
  CONSTRAINT `fk_dxddosid` FOREIGN KEY (`combo_id`) REFERENCES `DexunDDoSCombos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DDoSCombos`
--

LOCK TABLES `DDoSCombos` WRITE;
/*!40000 ALTER TABLE `DDoSCombos` DISABLE KEYS */;
INSERT INTO `DDoSCombos` VALUES (972333746217025537,972059802433499137,'7765d8ba-a1c8-49db-9535-cab4c3260b56','基础版',1999,1599,'','20','','','',0,'100GB峰值',50,2,1,1,1,5997,4797,1,'dexun',100,100,'BGP多线','50','50','100'),(972333746225414145,972059802438762497,'093ba3f6-617e-4d2c-840e-18014845d88a','专业版',2999,2399,'','30','','','',0,'200GB峰值',50,2,1,1,1,8997,7197,1,'dexun',100,100,'BGP多线','50','50','100');
/*!40000 ALTER TABLE `DDoSCombos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DDoSDomains`
--

DROP TABLE IF EXISTS `DDoSDomains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DDoSDomains` (
  `id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DDoSDomains`
--

LOCK TABLES `DDoSDomains` WRITE;
/*!40000 ALTER TABLE `DDoSDomains` DISABLE KEYS */;
/*!40000 ALTER TABLE `DDoSDomains` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DDoSService`
--

DROP TABLE IF EXISTS `DDoSService`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DDoSService` (
  `id` bigint NOT NULL,
  `combo_id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `u_user_id` bigint DEFAULT NULL,
  `server_ip` varchar(45) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `product_sitename` varchar(45) DEFAULT NULL,
  `stat_time` varchar(45) DEFAULT NULL,
  `end_time` varchar(45) DEFAULT NULL,
  `site_start` bigint DEFAULT NULL,
  `ddos_hh` varchar(45) DEFAULT NULL,
  `ks_start` bigint DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `recharge_domain` bigint DEFAULT NULL,
  `port_num` bigint DEFAULT NULL,
  `recharge_port` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `agent` varchar(45) DEFAULT NULL,
  `pro_type` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_ddosid_idx` (`combo_id`),
  CONSTRAINT `fk_ddosid` FOREIGN KEY (`combo_id`) REFERENCES `DDoSCombos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DDoSService`
--

LOCK TABLES `DDoSService` WRITE;
/*!40000 ALTER TABLE `DDoSService` DISABLE KEYS */;
/*!40000 ALTER TABLE `DDoSService` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAccCDNRank`
--

DROP TABLE IF EXISTS `DexunAccCDNRank`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAccCDNRank` (
  `id` bigint NOT NULL,
  `client_ip` varchar(45) DEFAULT NULL,
  `count_sum` bigint DEFAULT NULL,
  `order_id` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAccCDNRank`
--

LOCK TABLES `DexunAccCDNRank` WRITE;
/*!40000 ALTER TABLE `DexunAccCDNRank` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAccCDNRank` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAccLogs`
--

DROP TABLE IF EXISTS `DexunAccLogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAccLogs` (
  `id` bigint NOT NULL,
  `cachehit` varchar(45) DEFAULT NULL,
  `clientip` varchar(45) DEFAULT NULL,
  `clientport` varchar(45) DEFAULT NULL,
  `clientregion` varchar(45) DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `createdat` varchar(45) DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `domain_id` varchar(45) DEFAULT NULL,
  `form` varchar(45) DEFAULT NULL,
  `dx_id` bigint DEFAULT NULL,
  `instanceid` bigint DEFAULT NULL,
  `localaddr` varchar(45) DEFAULT NULL,
  `localip` varchar(45) DEFAULT NULL,
  `localport` bigint DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `nodeid` varchar(45) DEFAULT NULL,
  `packagesize` bigint DEFAULT NULL,
  `remoteaddr` varchar(45) DEFAULT NULL,
  `responsesize` bigint DEFAULT NULL,
  `responsestatuscode` bigint DEFAULT NULL,
  `timerangeend` varchar(45) DEFAULT NULL,
  `timerrangestart` varchar(45) DEFAULT NULL,
  `url` varchar(45) DEFAULT NULL,
  `wblist` varchar(45) DEFAULT NULL,
  `Accept` varchar(45) DEFAULT NULL,
  `Accept-Encoding` varchar(45) DEFAULT NULL,
  `Accept-Language` varchar(45) DEFAULT NULL,
  `Authorization` varchar(45) DEFAULT NULL,
  `Cache-Control` varchar(45) DEFAULT NULL,
  `Connection` varchar(45) DEFAULT NULL,
  `Pragma` varchar(45) DEFAULT NULL,
  `Purpose` varchar(45) DEFAULT NULL,
  `Referer` varchar(45) DEFAULT NULL,
  `Ugrade-Insecure-Requests` varchar(45) DEFAULT NULL,
  `User-Agent` varchar(45) DEFAULT NULL,
  `X-Forwarded-Host` varchar(45) DEFAULT NULL,
  `X-Forwarded-Port` varchar(45) DEFAULT NULL,
  `X-Forwarded-Proto` varchar(45) DEFAULT NULL,
  `X-Forwarded-Server` varchar(45) DEFAULT NULL,
  `X-Real-Ip` varchar(45) DEFAULT NULL,
  `Accept-Ranges` varchar(45) DEFAULT NULL,
  `Content-Encoding` varchar(45) DEFAULT NULL,
  `Content-Length` varchar(45) DEFAULT NULL,
  `Content-Type` varchar(45) DEFAULT NULL,
  `Date` varchar(45) DEFAULT NULL,
  `Etag` varchar(45) DEFAULT NULL,
  `Last-Modified` varchar(45) DEFAULT NULL,
  `Server` varchar(45) DEFAULT NULL,
  `Vary` varchar(45) DEFAULT NULL,
  `Www-Authenticate` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAccLogs`
--

LOCK TABLES `DexunAccLogs` WRITE;
/*!40000 ALTER TABLE `DexunAccLogs` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAccLogs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAreaAccessCon`
--

DROP TABLE IF EXISTS `DexunAreaAccessCon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAreaAccessCon` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `regions` json DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAreaAccessCon`
--

LOCK TABLES `DexunAreaAccessCon` WRITE;
/*!40000 ALTER TABLE `DexunAreaAccessCon` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAreaAccessCon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAreaRankStats`
--

DROP TABLE IF EXISTS `DexunAreaRankStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAreaRankStats` (
  `id` bigint NOT NULL,
  `order_id` bigint DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `total_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid6_idx` (`order_id`),
  CONSTRAINT `fk_orderid22` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAreaRankStats`
--

LOCK TABLES `DexunAreaRankStats` WRITE;
/*!40000 ALTER TABLE `DexunAreaRankStats` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAreaRankStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAreaStats`
--

DROP TABLE IF EXISTS `DexunAreaStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAreaStats` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `source` varchar(45) DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid7_idx` (`order_id`),
  CONSTRAINT `fk_orderid21` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAreaStats`
--

LOCK TABLES `DexunAreaStats` WRITE;
/*!40000 ALTER TABLE `DexunAreaStats` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAreaStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAtkInterStats`
--

DROP TABLE IF EXISTS `DexunAtkInterStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAtkInterStats` (
  `id` bigint NOT NULL,
  `order_id` bigint DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `total_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid5_idx` (`order_id`),
  CONSTRAINT `fk_orderid23` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAtkInterStats`
--

LOCK TABLES `DexunAtkInterStats` WRITE;
/*!40000 ALTER TABLE `DexunAtkInterStats` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAtkInterStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunATKLogLists`
--

DROP TABLE IF EXISTS `DexunATKLogLists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunATKLogLists` (
  `id` bigint NOT NULL,
  `clientip` varchar(45) DEFAULT NULL,
  `clientregion` varchar(45) DEFAULT NULL,
  `targeturl` varchar(45) DEFAULT NULL,
  `nodeid` varchar(45) DEFAULT NULL,
  `httpmethod` varchar(45) DEFAULT NULL,
  `attacktype` varchar(45) DEFAULT NULL,
  `attackinfo` varchar(45) DEFAULT NULL,
  `protecttype` varchar(45) DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `requestinfo` varchar(45) DEFAULT NULL,
  `domainid` varchar(45) DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `localip` varchar(45) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `timerangestart` varchar(45) DEFAULT NULL,
  `timerangeend` varchar(45) DEFAULT NULL,
  `instanceid` varchar(45) DEFAULT NULL,
  `clientport` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunATKLogLists`
--

LOCK TABLES `DexunATKLogLists` WRITE;
/*!40000 ALTER TABLE `DexunATKLogLists` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunATKLogLists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAtkLogs`
--

DROP TABLE IF EXISTS `DexunAtkLogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAtkLogs` (
  `id` bigint NOT NULL,
  `attackinfo` varchar(45) DEFAULT NULL,
  `attacktype` varchar(45) DEFAULT NULL,
  `clientip` varchar(45) DEFAULT NULL,
  `clientport` bigint DEFAULT NULL,
  `clientregion` varchar(45) DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `httpmethod` varchar(45) DEFAULT NULL,
  `al_id` bigint DEFAULT NULL,
  `instanceid` bigint DEFAULT NULL,
  `localip` varchar(45) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `protecttype` varchar(45) DEFAULT NULL,
  `requestinfo` varchar(45) DEFAULT NULL,
  `targeturl` varchar(45) DEFAULT NULL,
  `timerangeend` varchar(45) DEFAULT NULL,
  `timerangestart` varchar(45) DEFAULT NULL,
  `domainid` varchar(45) DEFAULT NULL,
  `nodeid` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid17_idx` (`order_id`),
  CONSTRAINT `fk_orderid17` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAtkLogs`
--

LOCK TABLES `DexunAtkLogs` WRITE;
/*!40000 ALTER TABLE `DexunAtkLogs` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunAtkLogs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunAtkStats`
--

DROP TABLE IF EXISTS `DexunAtkStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunAtkStats` (
  `id` bigint NOT NULL,
  `time` bigint DEFAULT NULL,
  `order_id` bigint DEFAULT NULL,
  `total_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid8_idx` (`order_id`),
  CONSTRAINT `fk_orderid20` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunAtkStats`
--

LOCK TABLES `DexunAtkStats` WRITE;
/*!40000 ALTER TABLE `DexunAtkStats` DISABLE KEYS */;
INSERT INTO `DexunAtkStats` VALUES (974638130131841025,1715597037,974622695435517953,0),(974882500435431425,1715655300,974622695435517953,0),(974912523491364865,1715662448,974622695435517953,0),(974912636447596545,1715662485,974622695435517953,0),(974914781941387265,1715662996,974622695435517953,0),(974927841004384257,1715666110,974622695435517953,0),(974930102781857793,1715666649,974622695435517953,0),(974945242665259009,1715670253,974622695435517953,0),(974962596225286145,1715674396,974622695435517953,0),(974964404545863681,1715674827,974622695435517953,0),(974965650928635905,1715675124,974622695435517953,0),(974965964911222785,1715675199,974622695435517953,0),(974976106893197313,1715677617,974622695435517953,0),(974980391383027713,1715678638,974622695435517953,0),(974980928172879873,1715678766,974622695435517953,0),(974983637026570241,1715679412,974622695435517953,0),(974996397908742145,1715682455,974622695435517953,0),(974997738815807489,1715682775,974622695435517953,0),(974999189363441665,1715683121,974622695435517953,0),(974999682051084289,1715683238,974622695435517953,0),(975000025663430657,1715683320,974622695435517953,0),(975000613064245249,1715683460,974622695435517953,0),(975015714807259137,1715687060,974622695435517953,0),(977013913023627265,1716163468,974622695435517953,0),(977015244998393857,1716163786,974622695435517953,0),(977016619870560257,1716164113,974622695435517953,0),(977019148367515649,1716164716,974622695435517953,0),(977029840355229697,1716167265,974622695435517953,0),(977044942252154881,1716170866,974622695435517953,0),(977091863683010561,1716182053,974622695435517953,0),(977093111035830273,1716182350,974622695435517953,0),(977094367979036673,1716182650,974622695435517953,0),(977095151798841345,1716182837,974622695435517953,0),(977105537141182465,1716185313,974622695435517953,0),(977106885489750017,1716185634,974622695435517953,0),(977108181373816833,1716185943,974622695435517953,0),(977115944821559297,1716187794,974622695435517953,0),(977122222580391937,1716189291,974622695435517953,0),(977123502077566977,1716189596,974622695435517953,0),(977124518855086081,1716189839,974622695435517953,0),(977129149163003905,1716190943,974622695435517953,0),(977130689279754241,1716191310,974622695435517953,0),(977145794803519489,1716194910,974622695435517953,0),(977374911295918081,1716249537,974622695435517953,0),(977382658069450753,1716251384,974622695435517953,0),(977397760288120833,1716254984,974622695435517953,0);
/*!40000 ALTER TABLE `DexunAtkStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunBWInstance`
--

DROP TABLE IF EXISTS `DexunBWInstance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunBWInstance` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `type` bigint DEFAULT NULL,
  `ip_list` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunBWInstance`
--

LOCK TABLES `DexunBWInstance` WRITE;
/*!40000 ALTER TABLE `DexunBWInstance` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunBWInstance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunBWSingle`
--

DROP TABLE IF EXISTS `DexunBWSingle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunBWSingle` (
  `id` bigint NOT NULL,
  `bws_id` bigint DEFAULT NULL,
  `uuid` varchar(45) DEFAULT NULL,
  `domain_id` varchar(45) DEFAULT NULL,
  `type` bigint DEFAULT NULL,
  `ip_list` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunBWSingle`
--

LOCK TABLES `DexunBWSingle` WRITE;
/*!40000 ALTER TABLE `DexunBWSingle` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunBWSingle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunBWStats`
--

DROP TABLE IF EXISTS `DexunBWStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunBWStats` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `ip` varchar(45) DEFAULT NULL,
  `bw_list_type` varchar(45) DEFAULT NULL,
  `total_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid2_idx` (`order_id`),
  CONSTRAINT `fk_orderid27` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunBWStats`
--

LOCK TABLES `DexunBWStats` WRITE;
/*!40000 ALTER TABLE `DexunBWStats` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunBWStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunCache`
--

DROP TABLE IF EXISTS `DexunCache`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunCache` (
  `id` bigint NOT NULL,
  `cache_id` bigint DEFAULT NULL,
  `dd_uuid` varchar(45) DEFAULT NULL,
  `cache_uuid` varchar(45) DEFAULT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `active` bigint DEFAULT NULL,
  `urlmode` varchar(45) DEFAULT NULL,
  `cachemode` varchar(45) DEFAULT NULL,
  `cachepath` varchar(45) DEFAULT NULL,
  `cacheextensions` varchar(45) DEFAULT NULL,
  `cachereg` varchar(45) DEFAULT NULL,
  `timeout` varchar(45) DEFAULT NULL,
  `weight` varchar(45) DEFAULT NULL,
  `createtime` varchar(45) DEFAULT NULL,
  `updatetime` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunCache`
--

LOCK TABLES `DexunCache` WRITE;
/*!40000 ALTER TABLE `DexunCache` DISABLE KEYS */;
INSERT INTO `DexunCache` VALUES (968328631017332737,0,'','ed417d2caf755c7b80bb8a385360fc81','5bbf31323cf9578995fe546eb216adb2dx73dx23y',1,'full','ext','','zip','','60','100','2024-04-26 09:49:18','2024-04-26 09:49:18');
/*!40000 ALTER TABLE `DexunCache` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunCacheModel`
--

DROP TABLE IF EXISTS `DexunCacheModel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunCacheModel` (
  `id` bigint NOT NULL,
  `cm_uuid` varchar(45) DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `dd_type` bigint DEFAULT NULL,
  `cache_name` varchar(45) DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  `urlmode` varchar(45) DEFAULT NULL,
  `cachemode` varchar(45) DEFAULT NULL,
  `cachepath` varchar(45) DEFAULT NULL,
  `cacheextensions` varchar(45) DEFAULT NULL,
  `cachereg` varchar(45) DEFAULT NULL,
  `timeout` bigint DEFAULT NULL,
  `weight` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunCacheModel`
--

LOCK TABLES `DexunCacheModel` WRITE;
/*!40000 ALTER TABLE `DexunCacheModel` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunCacheModel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunCC`
--

DROP TABLE IF EXISTS `DexunCC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunCC` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `cs_active` varchar(45) DEFAULT NULL,
  `policy` varchar(45) DEFAULT NULL,
  `url` varchar(45) DEFAULT NULL,
  `rate` varchar(45) DEFAULT NULL,
  `waitseconds` bigint DEFAULT NULL,
  `blockminutes` varchar(45) DEFAULT NULL,
  `redirectlocation` varchar(45) DEFAULT NULL,
  `global_concurrent` varchar(45) DEFAULT NULL,
  `waitpolicyminutes` varchar(45) DEFAULT NULL,
  `redirectwaitseconds` varchar(45) DEFAULT NULL,
  `count` varchar(45) DEFAULT NULL,
  `block_time` varchar(45) DEFAULT NULL,
  `block_active` varchar(45) DEFAULT NULL,
  `rr_active` varchar(45) DEFAULT NULL,
  `rr_rate` varchar(45) DEFAULT NULL,
  `ur_url` varchar(45) DEFAULT NULL,
  `ur_rate` varchar(45) DEFAULT NULL,
  `cookieName` varchar(45) DEFAULT NULL,
  `excludeExt` varchar(45) DEFAULT NULL,
  `concurrency` varchar(45) DEFAULT NULL,
  `r_blockMinutes` varchar(45) DEFAULT NULL,
  `whiteMinutes` varchar(45) DEFAULT NULL,
  `challengeLimit` varchar(45) DEFAULT NULL,
  `protectMinutes` varchar(45) DEFAULT NULL,
  `challengeMethods` json DEFAULT NULL,
  `challengePolicy` varchar(45) DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  `use_default` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunCC`
--

LOCK TABLES `DexunCC` WRITE;
/*!40000 ALTER TABLE `DexunCC` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunCC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunCloudEyeList`
--

DROP TABLE IF EXISTS `DexunCloudEyeList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunCloudEyeList` (
  `id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `content` varchar(45) DEFAULT NULL,
  `task_type` json DEFAULT NULL,
  `monitor_type` json DEFAULT NULL,
  `start_count` bigint DEFAULT NULL,
  `order_money` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunCloudEyeList`
--

LOCK TABLES `DexunCloudEyeList` WRITE;
/*!40000 ALTER TABLE `DexunCloudEyeList` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunCloudEyeList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunCombos`
--

DROP TABLE IF EXISTS `DexunCombos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunCombos` (
  `id` bigint NOT NULL,
  `uuid` varchar(45) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  `yms` varchar(45) DEFAULT NULL,
  `ccfy` varchar(45) DEFAULT NULL,
  `gjwaf` varchar(45) DEFAULT NULL,
  `ywll` varchar(45) DEFAULT NULL,
  `gjcs` varchar(45) DEFAULT NULL,
  `pro_flow` bigint DEFAULT NULL,
  `ddos_hh` varchar(45) DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `complete_state` bigint DEFAULT NULL,
  `firewall_state` bigint DEFAULT NULL,
  `waf_state` bigint DEFAULT NULL,
  `group_type` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunCombos`
--

LOCK TABLES `DexunCombos` WRITE;
/*!40000 ALTER TABLE `DexunCombos` DISABLE KEYS */;
INSERT INTO `DexunCombos` VALUES (962974346923986945,'9d8a6928-149b-43da-8cc4-9482610305df','体验版',99,79,'1','1','2','100',NULL,100,'10GB峰值',1,2,1,0,NULL,1),(962974346928787457,'cd16fc25-a897-4623-a8b6-7237b6a833d7','入门版',599,479,'10','5','2','1000',NULL,1000,'100GB峰值',10,2,1,0,NULL,1),(962974346933051393,'25fc50c0-0b53-4336-b6c5-6abf6b50ed85','基础版',999,799,'30','20','2','3600',NULL,3600,'100GB峰值',30,2,1,0,NULL,1),(962974346935853057,'32aee6cf-50e7-4172-b2ae-42d0e5f0ce6d','专业版',1999,1599,'50','30','1','5000',NULL,5000,'200GB峰值',50,2,1,1,NULL,1),(962974346939039745,'f4b907d1-0524-4a07-b0cf-030d659e4db1','商务版',4999,3999,'50','80','1','10000',NULL,10000,'400GB峰值',50,2,1,1,NULL,1),(962974346939113473,'c3a7f0c2-531f-4ac7-a374-1633d3cad920','企业版',2999,2399,'50','50','1','7200',NULL,7200,'300GB峰值',50,2,1,1,NULL,1),(962974346942824449,'61cca695-414e-45f6-818f-ca8d6e8e8996','旗舰版',9999,7999,'50','100','1','15000',NULL,15000,'600GB峰值',50,2,1,1,NULL,1);
/*!40000 ALTER TABLE `DexunCombos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunDDoSCombos`
--

DROP TABLE IF EXISTS `DexunDDoSCombos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunDDoSCombos` (
  `id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  `yms` varchar(45) DEFAULT NULL,
  `ccfy` varchar(45) DEFAULT NULL,
  `gjwaf` varchar(45) DEFAULT NULL,
  `ywll` varchar(45) DEFAULT NULL,
  `gjcs` varchar(45) DEFAULT NULL,
  `pro_flow` bigint DEFAULT NULL,
  `ddos_hh` varchar(45) DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `complete_state` bigint DEFAULT NULL,
  `firewall_state` bigint DEFAULT NULL,
  `waf_state` bigint DEFAULT NULL,
  `group_type` bigint DEFAULT NULL,
  `xl` varchar(45) DEFAULT NULL,
  `zfdks` varchar(45) DEFAULT NULL,
  `fhyms` varchar(45) DEFAULT NULL,
  `ywdk` varchar(45) DEFAULT NULL,
  `port_num` bigint DEFAULT NULL,
  `ywdk_num` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunDDoSCombos`
--

LOCK TABLES `DexunDDoSCombos` WRITE;
/*!40000 ALTER TABLE `DexunDDoSCombos` DISABLE KEYS */;
INSERT INTO `DexunDDoSCombos` VALUES (972059802433499137,'7765d8ba-a1c8-49db-9535-cab4c3260b56','基础版',1999,1599,'','20','','','',0,'100GB峰值',50,2,1,1,1,'BGP多线','50','50','100',100,100),(972059802438762497,'093ba3f6-617e-4d2c-840e-18014845d88a','专业版',2999,2399,'','30','','','',0,'200GB峰值',50,2,1,1,1,'BGP多线','50','50','100',100,100);
/*!40000 ALTER TABLE `DexunDDoSCombos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunDomainCert`
--

DROP TABLE IF EXISTS `DexunDomainCert`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunDomainCert` (
  `id` bigint NOT NULL,
  `dc_id` bigint DEFAULT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `status` bigint DEFAULT NULL,
  `ssl_always` bigint DEFAULT NULL,
  `hsts` bigint DEFAULT NULL,
  `cert_name` varchar(45) DEFAULT NULL,
  `cert` varchar(255) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  `desc` varchar(45) DEFAULT NULL,
  `createtime` varchar(45) DEFAULT NULL,
  `updatetime` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunDomainCert`
--

LOCK TABLES `DexunDomainCert` WRITE;
/*!40000 ALTER TABLE `DexunDomainCert` DISABLE KEYS */;
INSERT INTO `DexunDomainCert` VALUES (968353809279934465,0,'5bbf31323cf9578995fe546eb216adb2dx73dx23y',1,1,1,'www.cdn12.com','aaaa','aaaa','2024-04-26 11:32:13','2024-04-26 11:32:13','2024-04-26 11:32:13');
/*!40000 ALTER TABLE `DexunDomainCert` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunDomainFlow`
--

DROP TABLE IF EXISTS `DexunDomainFlow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunDomainFlow` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `request_size` bigint DEFAULT NULL,
  `response_size` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid10_idx` (`order_id`),
  CONSTRAINT `fk_orderid18` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunDomainFlow`
--

LOCK TABLES `DexunDomainFlow` WRITE;
/*!40000 ALTER TABLE `DexunDomainFlow` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunDomainFlow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunDomainQuery`
--

DROP TABLE IF EXISTS `DexunDomainQuery`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunDomainQuery` (
  `id` bigint NOT NULL,
  `cache_calls` bigint DEFAULT NULL,
  `cache_rate` bigint DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `total_calls` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunDomainQuery`
--

LOCK TABLES `DexunDomainQuery` WRITE;
/*!40000 ALTER TABLE `DexunDomainQuery` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunDomainQuery` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunDomains`
--

DROP TABLE IF EXISTS `DexunDomains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunDomains` (
  `id` bigint NOT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `primary_domain` varchar(45) DEFAULT NULL,
  `domain_status` bigint DEFAULT NULL,
  `domain_record` varchar(45) DEFAULT NULL,
  `port` varchar(45) DEFAULT NULL,
  `server` varchar(45) DEFAULT NULL,
  `protocol` varchar(45) DEFAULT NULL,
  `overload_type` varchar(45) DEFAULT NULL,
  `overload_redirect_url` varchar(45) DEFAULT NULL,
  `overload_redirect_code` varchar(45) DEFAULT NULL,
  `load_balancing` varchar(45) DEFAULT NULL,
  `redirect` varchar(45) DEFAULT NULL,
  `uri_forward` varchar(45) DEFAULT NULL,
  `sni` varchar(45) DEFAULT NULL,
  `sni_port` varchar(45) DEFAULT NULL,
  `weight` varchar(45) DEFAULT NULL,
  `address` varchar(45) DEFAULT NULL,
  `sni_protocol` varchar(45) DEFAULT NULL,
  `concurrent` varchar(45) DEFAULT NULL,
  `four_layers_config` varchar(45) DEFAULT NULL,
  `cache_file_size_limit` bigint DEFAULT NULL,
  `cache_total_size_limit` bigint DEFAULT NULL,
  `cache_config` json DEFAULT NULL,
  `cache_active` bigint DEFAULT NULL,
  `white_num` bigint DEFAULT NULL,
  `use_flow` bigint DEFAULT NULL,
  `createtime` varchar(45) DEFAULT NULL,
  `updatetime` varchar(45) DEFAULT NULL,
  `access_active` varchar(45) DEFAULT NULL,
  `grouping` varchar(45) DEFAULT NULL,
  `is_filings` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunDomains`
--

LOCK TABLES `DexunDomains` WRITE;
/*!40000 ALTER TABLE `DexunDomains` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunDomains` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunHttpPackStats`
--

DROP TABLE IF EXISTS `DexunHttpPackStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunHttpPackStats` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `time` bigint DEFAULT NULL,
  `total_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid4_idx` (`order_id`),
  CONSTRAINT `fk_orderid24` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunHttpPackStats`
--

LOCK TABLES `DexunHttpPackStats` WRITE;
/*!40000 ALTER TABLE `DexunHttpPackStats` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunHttpPackStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunInterceptStats`
--

DROP TABLE IF EXISTS `DexunInterceptStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunInterceptStats` (
  `id` bigint NOT NULL,
  `order_id` bigint DEFAULT NULL,
  `app_cc` bigint DEFAULT NULL,
  `cc` bigint DEFAULT NULL,
  `ip_black` bigint DEFAULT NULL,
  `referer` bigint DEFAULT NULL,
  `url_black` bigint DEFAULT NULL,
  `web_protect` bigint DEFAULT NULL,
  `other` bigint DEFAULT NULL,
  `area_acc` bigint DEFAULT NULL,
  `safe_acc` bigint DEFAULT NULL,
  `pre_acc` bigint DEFAULT NULL,
  `create_time` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid9_idx` (`order_id`),
  CONSTRAINT `fk_orderid19` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunInterceptStats`
--

LOCK TABLES `DexunInterceptStats` WRITE;
/*!40000 ALTER TABLE `DexunInterceptStats` DISABLE KEYS */;
INSERT INTO `DexunInterceptStats` VALUES (974638128866009089,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715600637),(974712111789731841,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715618276),(974880908610179073,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715658521),(974881257083080705,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715658604),(974882498275594241,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715658900),(974912523496620033,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715666058),(974912634298413057,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715666085),(974914778543923201,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715666596),(974927838854541313,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715669710),(974930100200083457,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715670249),(974945242655969281,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715673859),(974962593747238913,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715677996),(974964402489917441,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715678427),(974965648346312705,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715678724),(974965962334126081,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715678799),(974976104322678785,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715681217),(974980388833820673,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715682238),(974980925730656257,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715682366),(974983635308441601,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715683012),(974996396312231937,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715686055),(974997737321586689,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715686375),(974999187972800513,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715686721),(974999680704114689,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715686838),(975000024279994369,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715686920),(975000611343958017,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715687060),(975015713620160513,974622695435517953,0,0,0,0,0,0,0,0,0,0,1715690661),(977013911564214273,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716167068),(977015243605639169,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716167386),(977016618193281025,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716167713),(977019146914742273,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716168316),(977029838993899521,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716170865),(977044941906612225,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716174466),(977091862304157697,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716185653),(977093109498064897,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716185950),(977094366327394305,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716186250),(977095150428733441,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716186437),(977105535607599105,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716188913),(977106883286376449,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716189234),(977108179936305153,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716189543),(977115943377940481,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716191394),(977122167328370689,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716192878),(977122220931108865,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716192891),(977123500358713345,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716193196),(977124517536825345,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716193439),(977129147893448705,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716194543),(977130687725449217,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716194910),(977145795082752001,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716198512),(977374909879869441,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716253137),(977382656708435969,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716254984),(977397759819661313,974622695435517953,0,0,0,0,0,0,0,0,0,0,1716258585);
/*!40000 ALTER TABLE `DexunInterceptStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunLeechlink`
--

DROP TABLE IF EXISTS `DexunLeechlink`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunLeechlink` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `dd_uuid` varchar(45) DEFAULT NULL,
  `pro_type` varchar(45) DEFAULT NULL,
  `type` tinyint DEFAULT NULL,
  `active` tinyint DEFAULT NULL,
  `domians` json DEFAULT NULL,
  `allow_empty` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunLeechlink`
--

LOCK TABLES `DexunLeechlink` WRITE;
/*!40000 ALTER TABLE `DexunLeechlink` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunLeechlink` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunLineChartStats`
--

DROP TABLE IF EXISTS `DexunLineChartStats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunLineChartStats` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `time` bigint DEFAULT NULL,
  `response_size` bigint DEFAULT NULL,
  `request_size` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid3_idx` (`order_id`),
  CONSTRAINT `fk_orderid25` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunLineChartStats`
--

LOCK TABLES `DexunLineChartStats` WRITE;
/*!40000 ALTER TABLE `DexunLineChartStats` DISABLE KEYS */;
INSERT INTO `DexunLineChartStats` VALUES (974638137096187905,974622695435517953,1715427839,0,0),(974638137111040001,974622695435517953,1715431439,0,0),(974638137119719425,974622695435517953,1715435039,0,0),(974638137124712449,974622695435517953,1715438639,0,0),(974638137127026689,974622695435517953,1715442239,0,0),(974638137137045505,974622695435517953,1715445839,0,0),(974638137147834369,974622695435517953,1715449439,0,0),(974638137152253953,974622695435517953,1715453039,0,0),(974638137159069697,974622695435517953,1715456639,0,0),(974638137162256385,974622695435517953,1715460239,0,0),(974638137165795329,974622695435517953,1715463839,0,0),(974638137169346561,974622695435517953,1715467439,0,0),(974638137176891393,974622695435517953,1715471039,0,0),(974638137178968065,974622695435517953,1715478239,0,0),(974638137179172865,974622695435517953,1715474639,0,0),(974638137181814785,974622695435517953,1715481839,0,0),(974638137189560321,974622695435517953,1715485439,0,0),(974638137189584897,974622695435517953,1715489039,0,0),(974638137192206337,974622695435517953,1715492639,0,0),(974638137195253761,974622695435517953,1715499839,0,0),(974638137196933121,974622695435517953,1715496239,0,0),(974638137198563329,974622695435517953,1715507039,0,0),(974638137200947201,974622695435517953,1715503439,0,0),(974638137205444609,974622695435517953,1715510639,0,0),(974638137206845441,974622695435517953,1715517839,0,0),(974638137208492033,974622695435517953,1715521439,0,0),(974638137209544705,974622695435517953,1715514239,0,0),(974638137217646593,974622695435517953,1715525039,0,0),(974638137224302593,974622695435517953,1715528639,0,0),(974638137230659585,974622695435517953,1715532239,0,0),(974638137235234817,974622695435517953,1715539439,0,0),(974638137235382273,974622695435517953,1715535839,0,0),(974638137237917697,974622695435517953,1715543039,0,0),(974638137244114945,974622695435517953,1715550239,0,0),(974638137244143617,974622695435517953,1715546639,0,0),(974638137246629889,974622695435517953,1715553839,0,0),(974638137250115585,974622695435517953,1715561039,0,0),(974638137251799041,974622695435517953,1715557439,0,0),(974638137255960577,974622695435517953,1715564639,0,0),(974638137259896833,974622695435517953,1715568239,0,0),(974638137262501889,974622695435517953,1715571839,0,0),(974638137269317633,974622695435517953,1715575439,0,0),(974638137270927361,974622695435517953,1715579039,0,0),(974638137272291329,974622695435517953,1715582639,0,0),(974638137277341697,974622695435517953,1715586239,0,0),(974638137279504385,974622695435517953,1715593439,0,0),(974638137281114113,974622695435517953,1715589839,0,0),(974638137283698689,974622695435517953,1715597039,0,0),(974882511840706561,974622695435517953,1715486103,0,0),(974882511853473793,974622695435517953,1715489703,0,0),(974882511862542337,974622695435517953,1715493303,0,0),(974882511872192513,974622695435517953,1715496903,0,0),(974882511877001217,974622695435517953,1715500503,0,0),(974882511883173889,974622695435517953,1715504103,0,0),(974882511887331329,974622695435517953,1715507703,0,0),(974882511894941697,974622695435517953,1715511303,0,0),(974882511899013121,974622695435517953,1715514903,0,0),(974882511899639809,974622695435517953,1715518503,0,0),(974882511906697217,974622695435517953,1715522103,0,0),(974882511908184065,974622695435517953,1715525703,0,0),(974882511909425153,974622695435517953,1715529303,0,0),(974882511914500097,974622695435517953,1715532903,0,0),(974882511916851201,974622695435517953,1715536503,0,0),(974882511922249729,974622695435517953,1715540103,0,0),(974882511926067201,974622695435517953,1715543703,0,0),(974882511930945537,974622695435517953,1715547303,0,0),(974882511935205377,974622695435517953,1715554503,0,0),(974882511935909889,974622695435517953,1715550903,0,0),(974882511940632577,974622695435517953,1715558103,0,0),(974882511941402625,974622695435517953,1715561703,0,0),(974882511947411457,974622695435517953,1715565303,0,0),(974882511951060993,974622695435517953,1715568903,0,0),(974882511951794177,974622695435517953,1715572503,0,0),(974882511957737473,974622695435517953,1715576103,0,0),(974882511960281089,974622695435517953,1715579703,0,0),(974882511962734593,974622695435517953,1715583303,0,0),(974882511966982145,974622695435517953,1715586903,0,0),(974882511971110913,974622695435517953,1715590503,0,0),(974882511977750529,974622695435517953,1715594103,0,0),(974882511978033153,974622695435517953,1715597703,0,0),(974882511982698497,974622695435517953,1715601303,0,0),(974882511984041985,974622695435517953,1715604903,0,0),(974882511988719617,974622695435517953,1715612103,0,0),(974882511988895745,974622695435517953,1715608503,0,0),(974882511992119297,974622695435517953,1715615703,0,0),(974882511995854849,974622695435517953,1715619303,0,0),(974882512000348161,974622695435517953,1715626503,0,0),(974882512002093057,974622695435517953,1715622903,0,0),(974882512004685825,974622695435517953,1715630103,0,0),(974882512011694081,974622695435517953,1715637303,0,0),(974882512012550145,974622695435517953,1715633703,0,0),(974882512013381633,974622695435517953,1715640903,0,0),(974882512018903041,974622695435517953,1715644503,0,0),(974882512022208513,974622695435517953,1715648103,0,0),(974882512023932929,974622695435517953,1715651703,0,0),(974882512028966913,974622695435517953,1715655303,0,0),(974912523495051265,974622695435517953,1715493258,0,0),(974912523554918401,974622695435517953,1715496858,0,0),(974912523565453313,974622695435517953,1715500458,0,0),(974912523566940161,974622695435517953,1715504058,0,0),(974912523579236353,974622695435517953,1715507658,0,0),(974912523591315457,974622695435517953,1715511258,0,0),(974912523592257537,974622695435517953,1715514858,0,0),(974912523600027649,974622695435517953,1715518458,0,0),(974912523608858625,974622695435517953,1715522058,0,0),(974912523616440321,974622695435517953,1715525658,0,0),(974912523625693185,974622695435517953,1715529258,0,0),(974912523636944897,974622695435517953,1715532858,0,0),(974912523643957249,974622695435517953,1715536458,0,0),(974912523651534849,974622695435517953,1715540058,0,0),(974912523654516737,974622695435517953,1715543658,0,0),(974912523660439553,974622695435517953,1715547258,0,0),(974912523664519169,974622695435517953,1715550858,0,0),(974912523669331969,974622695435517953,1715554458,0,0),(974912523674714113,974622695435517953,1715558058,0,0),(974912523678281729,974622695435517953,1715561658,0,0),(974912523683078145,974622695435517953,1715565258,0,0),(974912523684954113,974622695435517953,1715568858,0,0),(974912523688390657,974622695435517953,1715572458,0,0),(974912523696234497,974622695435517953,1715576058,0,0),(974912523697627137,974622695435517953,1715579658,0,0),(974912523702128641,974622695435517953,1715583258,0,0),(974912523706953729,974622695435517953,1715590458,0,0),(974912523707502593,974622695435517953,1715586858,0,0),(974912523712008193,974622695435517953,1715594058,0,0),(974912523721474049,974622695435517953,1715597658,0,0),(974912523722625025,974622695435517953,1715604858,0,0),(974912523725008897,974622695435517953,1715601258,0,0),(974912523727040513,974622695435517953,1715608458,0,0),(974912523732307969,974622695435517953,1715612058,0,0),(974912523736752129,974622695435517953,1715615658,0,0),(974912523739107329,974622695435517953,1715619258,0,0),(974912523742642177,974622695435517953,1715622858,0,0),(974912523749437441,974622695435517953,1715626458,0,0),(974912523750985729,974622695435517953,1715633658,0,0),(974912523752751105,974622695435517953,1715630058,0,0),(974912523760513025,974622695435517953,1715640858,0,0),(974912523760939009,974622695435517953,1715637258,0,0),(974912523766870017,974622695435517953,1715644458,0,0),(974912523769098241,974622695435517953,1715648058,0,0),(974912523770646529,974622695435517953,1715651658,0,0),(974912523773030401,974622695435517953,1715655258,0,0),(974912523776819201,974622695435517953,1715658858,0,0),(974912523779330049,974622695435517953,1715662458,0,0),(974912890135035905,974622695435517953,1715493346,0,0),(974912890147831809,974622695435517953,1715496946,0,0),(974912890152099841,974622695435517953,1715500546,0,0),(974912890152820737,974622695435517953,1715504146,0,0),(974912890158669825,974622695435517953,1715507746,0,0),(974912890163802113,974622695435517953,1715511346,0,0),(974912890170036225,974622695435517953,1715514946,0,0),(974912890176630785,974622695435517953,1715522146,0,0),(974912890177097729,974622695435517953,1715518546,0,0),(974912890178449409,974622695435517953,1715525746,0,0),(974912890184015873,974622695435517953,1715532946,0,0),(974912890184749057,974622695435517953,1715529346,0,0),(974912890190819329,974622695435517953,1715540146,0,0),(974912890194599937,974622695435517953,1715536546,0,0),(974912890197901313,974622695435517953,1715543746,0,0),(974912890199904257,974622695435517953,1715547346,0,0),(974912890204450817,974622695435517953,1715550946,0,0),(974912890204848129,974622695435517953,1715554546,0,0),(974912890211110913,974622695435517953,1715558146,0,0),(974912890214281217,974622695435517953,1715565346,0,0),(974912890214912001,974622695435517953,1715561746,0,0),(974912890216476673,974622695435517953,1715568946,0,0),(974912890222256129,974622695435517953,1715576146,0,0),(974912890222366721,974622695435517953,1715572546,0,0),(974912890226511873,974622695435517953,1715579746,0,0),(974912890228682753,974622695435517953,1715583346,0,0),(974912890230665217,974622695435517953,1715586946,0,0),(974912890233573377,974622695435517953,1715594146,0,0),(974912890236317697,974622695435517953,1715590546,0,0),(974912890237042689,974622695435517953,1715597746,0,0),(974912890244890625,974622695435517953,1715601346,0,0),(974912890246975489,974622695435517953,1715608546,0,0),(974912890247942145,974622695435517953,1715604946,0,0),(974912890253144065,974622695435517953,1715612146,0,0),(974912890253803521,974622695435517953,1715615746,0,0),(974912890257829889,974622695435517953,1715619346,0,0),(974912890258583553,974622695435517953,1715622946,0,0),(974912890265075713,974622695435517953,1715626546,0,0),(974912890267852801,974622695435517953,1715630146,0,0),(974912890269605889,974622695435517953,1715633746,0,0),(974912890272829441,974622695435517953,1715640946,0,0),(974912890273120257,974622695435517953,1715637346,0,0),(974912890274795521,974622695435517953,1715644546,0,0),(974912890276921345,974622695435517953,1715648146,0,0),(974912890280677377,974622695435517953,1715651746,0,0),(974912890285326337,974622695435517953,1715655346,0,0),(974912890286792705,974622695435517953,1715658946,0,0),(974912890291884033,974622695435517953,1715662546,0,0),(974914794770169857,974622695435517953,1715493800,0,0),(974914794789363713,974622695435517953,1715497400,0,0),(974914794797510657,974622695435517953,1715504600,0,0),(974914794797916161,974622695435517953,1715501000,0,0),(974914794803556353,974622695435517953,1715508200,0,0),(974914794813284353,974622695435517953,1715511800,0,0),(974914794817826817,974622695435517953,1715515400,0,0),(974914794821517313,974622695435517953,1715519000,0,0),(974914794830704641,974622695435517953,1715522600,0,0),(974914794832826369,974622695435517953,1715529800,0,0),(974914794834481153,974622695435517953,1715526200,0,0),(974914794839957505,974622695435517953,1715533400,0,0),(974914794844725249,974622695435517953,1715537000,0,0),(974914794844946433,974622695435517953,1715540600,0,0),(974914794857091073,974622695435517953,1715544200,0,0),(974914794859823105,974622695435517953,1715547800,0,0),(974914794862194689,974622695435517953,1715551400,0,0),(974914794869100545,974622695435517953,1715555000,0,0),(974914794872778753,974622695435517953,1715558600,0,0),(974914794876829697,974622695435517953,1715565800,0,0),(974914794877726721,974622695435517953,1715562200,0,0),(974914794880593921,974622695435517953,1715569400,0,0),(974914794886045697,974622695435517953,1715573000,0,0),(974914794888384513,974622695435517953,1715576600,0,0),(974914794894082049,974622695435517953,1715580200,0,0),(974914794897604609,974622695435517953,1715583800,0,0),(974914794899419137,974622695435517953,1715587400,0,0),(974914794908479489,974622695435517953,1715591000,0,0),(974914794914205697,974622695435517953,1715594600,0,0),(974914794919034881,974622695435517953,1715598200,0,0),(974914794922545153,974622695435517953,1715601800,0,0),(974914794926616577,974622695435517953,1715605400,0,0),(974914794929147905,974622695435517953,1715612600,0,0),(974914794929496065,974622695435517953,1715609000,0,0),(974914794936750081,974622695435517953,1715616200,0,0),(974914794938494977,974622695435517953,1715619800,0,0),(974914794945110017,974622695435517953,1715623400,0,0),(974914794949513217,974622695435517953,1715627000,0,0),(974914794950807553,974622695435517953,1715630600,0,0),(974914794951720961,974622695435517953,1715634200,0,0),(974914794955636737,974622695435517953,1715637800,0,0),(974914794958548993,974622695435517953,1715641400,0,0),(974914794960584705,974622695435517953,1715645000,0,0),(974914794964615169,974622695435517953,1715648600,0,0),(974914794967547905,974622695435517953,1715652200,0,0),(974914794970501121,974622695435517953,1715655800,0,0),(974914794972381185,974622695435517953,1715659400,0,0),(974914794976141313,974622695435517953,1715663000,0,0),(974927852600553473,974622695435517953,1715496913,0,0),(974927852618969089,974622695435517953,1715500513,0,0),(974927852622364673,974622695435517953,1715504113,0,0),(974927852629434369,974622695435517953,1715507713,0,0),(974927852632846337,974622695435517953,1715511313,0,0),(974927852635324417,974622695435517953,1715514913,0,0),(974927852637364225,974622695435517953,1715518513,0,0),(974927852641288193,974622695435517953,1715522113,0,0),(974927852646563841,974622695435517953,1715525713,0,0),(974927852650385409,974622695435517953,1715529313,0,0),(974927852651114497,974622695435517953,1715532913,0,0),(974927852659965953,974622695435517953,1715536513,0,0),(974927852673056769,974622695435517953,1715540113,0,0),(974927852676063233,974622695435517953,1715543713,0,0),(974927852681908225,974622695435517953,1715547313,0,0),(974927852684496897,974622695435517953,1715550913,0,0),(974927852689170433,974622695435517953,1715554513,0,0),(974927852693954561,974622695435517953,1715558113,0,0),(974927852696653825,974622695435517953,1715561713,0,0),(974927852696793089,974622695435517953,1715565313,0,0),(974927852700479489,974622695435517953,1715568913,0,0),(974927852705034241,974622695435517953,1715572513,0,0),(974927852709126145,974622695435517953,1715576113,0,0),(974927852710912001,974622695435517953,1715579713,0,0),(974927852712988673,974622695435517953,1715583313,0,0),(974927852719046657,974622695435517953,1715586913,0,0),(974927852724178945,974622695435517953,1715590513,0,0),(974927852728782849,974622695435517953,1715594113,0,0),(974927852732850177,974622695435517953,1715597713,0,0),(974927852733538305,974622695435517953,1715601313,0,0),(974927852736876545,974622695435517953,1715604913,0,0),(974927852743589889,974622695435517953,1715612113,0,0),(974927852744327169,974622695435517953,1715608513,0,0),(974927852745580545,974622695435517953,1715615713,0,0),(974927852749893633,974622695435517953,1715622913,0,0),(974927852752920577,974622695435517953,1715619313,0,0),(974927852754247681,974622695435517953,1715626513,0,0),(974927852755124225,974622695435517953,1715630113,0,0),(974927852761751553,974622695435517953,1715633713,0,0),(974927852761780225,974622695435517953,1715640913,0,0),(974927852763353089,974622695435517953,1715637313,0,0),(974927852768890881,974622695435517953,1715644513,0,0),(974927852769353729,974622695435517953,1715648113,0,0),(974927852772175873,974622695435517953,1715651713,0,0),(974927852772618241,974622695435517953,1715655313,0,0),(974927852775792641,974622695435517953,1715658913,0,0),(974927852778557441,974622695435517953,1715666113,0,0),(974927852779479041,974622695435517953,1715662513,0,0),(974930122549878785,974622695435517953,1715497454,0,0),(974930122557427713,974622695435517953,1715501054,0,0),(974930122562600961,974622695435517953,1715504654,0,0),(974930122565152769,974622695435517953,1715508254,0,0),(974930122568208385,974622695435517953,1715511854,0,0),(974930122573934593,974622695435517953,1715515454,0,0),(974930122574237697,974622695435517953,1715519054,0,0),(974930122578161665,974622695435517953,1715522654,0,0),(974930122584543233,974622695435517953,1715526254,0,0),(974930122587590657,974622695435517953,1715533454,0,0),(974930122587893761,974622695435517953,1715529854,0,0),(974930122589855745,974622695435517953,1715537054,0,0),(974930122595454977,974622695435517953,1715540654,0,0),(974930122597666817,974622695435517953,1715547854,0,0),(974930122598772737,974622695435517953,1715544254,0,0),(974930122605711361,974622695435517953,1715551454,0,0),(974930122608955393,974622695435517953,1715555054,0,0),(974930122610561025,974622695435517953,1715558654,0,0),(974930122614829057,974622695435517953,1715565854,0,0),(974930122615402497,974622695435517953,1715562254,0,0),(974930122618683393,974622695435517953,1715569454,0,0),(974930122624221185,974622695435517953,1715573054,0,0),(974930122627010561,974622695435517953,1715576654,0,0),(974930122632347649,974622695435517953,1715580254,0,0),(974930122635608065,974622695435517953,1715583854,0,0),(974930122640281601,974622695435517953,1715587454,0,0),(974930122644647937,974622695435517953,1715591054,0,0),(974930122646159361,974622695435517953,1715594654,0,0),(974930122654285825,974622695435517953,1715598254,0,0),(974930122657042433,974622695435517953,1715601854,0,0),(974930122658824193,974622695435517953,1715605454,0,0),(974930122660679681,974622695435517953,1715609054,0,0),(974930122665127937,974622695435517953,1715612654,0,0),(974930122674630657,974622695435517953,1715619854,0,0),(974930122674946049,974622695435517953,1715616254,0,0),(974930122679103489,974622695435517953,1715623454,0,0),(974930122682863617,974622695435517953,1715627054,0,0),(974930122686537729,974622695435517953,1715634254,0,0),(974930122688442369,974622695435517953,1715630654,0,0),(974930122692243457,974622695435517953,1715641454,0,0),(974930122692829185,974622695435517953,1715637854,0,0),(974930122695909377,974622695435517953,1715645054,0,0),(974930122697154561,974622695435517953,1715648654,0,0),(974930122701299713,974622695435517953,1715652254,0,0),(974930122703519745,974622695435517953,1715655854,0,0),(974930122705776641,974622695435517953,1715659454,0,0),(974930122706952193,974622695435517953,1715666654,0,0),(974930122707873793,974622695435517953,1715663054,0,0),(974945242669903873,974622695435517953,1715501059,0,0),(974945242703417345,974622695435517953,1715504659,0,0),(974945242711498753,974622695435517953,1715508259,0,0),(974945242715578369,974622695435517953,1715511859,0,0),(974945242722201601,974622695435517953,1715515459,0,0),(974945242729992193,974622695435517953,1715519059,0,0),(974945242734014465,974622695435517953,1715522659,0,0),(974945242737332225,974622695435517953,1715526259,0,0),(974945242746830849,974622695435517953,1715529859,0,0),(974945242760753153,974622695435517953,1715533459,0,0),(974945242770391041,974622695435517953,1715537059,0,0),(974945242779361281,974622695435517953,1715540659,0,0),(974945242780418049,974622695435517953,1715544259,0,0),(974945242788155393,974622695435517953,1715547859,0,0),(974945242793824257,974622695435517953,1715551459,0,0),(974945242800549889,974622695435517953,1715555059,0,0),(974945242804244481,974622695435517953,1715562259,0,0),(974945242805551105,974622695435517953,1715558659,0,0),(974945242810867713,974622695435517953,1715565859,0,0),(974945242812264449,974622695435517953,1715569459,0,0),(974945242819686401,974622695435517953,1715573059,0,0),(974945242821062657,974622695435517953,1715576659,0,0),(974945242822852609,974622695435517953,1715580259,0,0),(974945242825912321,974622695435517953,1715583859,0,0),(974945242829537281,974622695435517953,1715587459,0,0),(974945242830139393,974622695435517953,1715591059,0,0),(974945242834386945,974622695435517953,1715594659,0,0),(974945242839609345,974622695435517953,1715601859,0,0),(974945242839736321,974622695435517953,1715598259,0,0),(974945242842976257,974622695435517953,1715605459,0,0),(974945242844291073,974622695435517953,1715609059,0,0),(974945242845626369,974622695435517953,1715616259,0,0),(974945242847125505,974622695435517953,1715612659,0,0),(974945242849570817,974622695435517953,1715619859,0,0),(974945242855387137,974622695435517953,1715627059,0,0),(974945242857234433,974622695435517953,1715623459,0,0),(974945242859954177,974622695435517953,1715630659,0,0),(974945242865692673,974622695435517953,1715634259,0,0),(974945242866520065,974622695435517953,1715637859,0,0),(974945242868039681,974622695435517953,1715641459,0,0),(974945242870779905,974622695435517953,1715645059,0,0),(974945242874834945,974622695435517953,1715648659,0,0),(974945242877521921,974622695435517953,1715652259,0,0),(974945242879025153,974622695435517953,1715655859,0,0),(974945242881921025,974622695435517953,1715659459,0,0),(974945242884509697,974622695435517953,1715663059,0,0),(974945242885746689,974622695435517953,1715666659,0,0),(974945242887749633,974622695435517953,1715670259,0,0),(974962610420158465,974622695435517953,1715505200,0,0),(974962610436591617,974622695435517953,1715508800,0,0),(974962610439884801,974622695435517953,1715512400,0,0),(974962610446024705,974622695435517953,1715516000,0,0),(974962610451415041,974622695435517953,1715519600,0,0),(974962610453303297,974622695435517953,1715526800,0,0),(974962610455056385,974622695435517953,1715523200,0,0),(974962610459246593,974622695435517953,1715530400,0,0),(974962610460237825,974622695435517953,1715534000,0,0),(974962610465570817,974622695435517953,1715537600,0,0),(974962610465841153,974622695435517953,1715541200,0,0),(974962610474655745,974622695435517953,1715544800,0,0),(974962610477174785,974622695435517953,1715548400,0,0),(974962610479546369,974622695435517953,1715552000,0,0),(974962610483310593,974622695435517953,1715555600,0,0),(974962610485731329,974622695435517953,1715559200,0,0),(974962610490875905,974622695435517953,1715562800,0,0),(974962610496233473,974622695435517953,1715566400,0,0),(974962610498170881,974622695435517953,1715570000,0,0),(974962610502770689,974622695435517953,1715573600,0,0),(974962610509234177,974622695435517953,1715577200,0,0),(974962610511626241,974622695435517953,1715580800,0,0),(974962610517590017,974622695435517953,1715584400,0,0),(974962610518163457,974622695435517953,1715588000,0,0),(974962610521567233,974622695435517953,1715591600,0,0),(974962610525450241,974622695435517953,1715595200,0,0),(974962610528141313,974622695435517953,1715598800,0,0),(974962610531774465,974622695435517953,1715606000,0,0),(974962610532413441,974622695435517953,1715602400,0,0),(974962610538418177,974622695435517953,1715609600,0,0),(974962610542419969,974622695435517953,1715613200,0,0),(974962610545311745,974622695435517953,1715616800,0,0),(974962610548678657,974622695435517953,1715620400,0,0),(974962610555940865,974622695435517953,1715624000,0,0),(974962610558308353,974622695435517953,1715627600,0,0),(974962610560286721,974622695435517953,1715631200,0,0),(974962610564231169,974622695435517953,1715634800,0,0),(974962610565742593,974622695435517953,1715642000,0,0),(974962610565894145,974622695435517953,1715638400,0,0),(974962610570276865,974622695435517953,1715649200,0,0),(974962610570940417,974622695435517953,1715645600,0,0),(974962610573643777,974622695435517953,1715652800,0,0),(974962610577584129,974622695435517953,1715656400,0,0),(974962610583183361,974622695435517953,1715660000,0,0),(974962610588114945,974622695435517953,1715663600,0,0),(974962610590679041,974622695435517953,1715667200,0,0),(974962610591916033,974622695435517953,1715670800,0,0),(974962610595176449,974622695435517953,1715674400,0,0),(974964415093706753,974622695435517953,1715505630,0,0),(974964415106568193,974622695435517953,1715509230,0,0),(974964415113027585,974622695435517953,1715512830,0,0),(974964415115059201,974622695435517953,1715516430,0,0),(974964415120691201,974622695435517953,1715520030,0,0),(974964415125733377,974622695435517953,1715523630,0,0),(974964415132082177,974622695435517953,1715527230,0,0),(974964415135977473,974622695435517953,1715530830,0,0),(974964415141941249,974622695435517953,1715534430,0,0),(974964415145054209,974622695435517953,1715538030,0,0),(974964415148535809,974622695435517953,1715541630,0,0),(974964415153700865,974622695435517953,1715545230,0,0),(974964415157366785,974622695435517953,1715548830,0,0),(974964415160840193,974622695435517953,1715552430,0,0),(974964415170580481,974622695435517953,1715556030,0,0),(974964415175577601,974622695435517953,1715559630,0,0),(974964415177121793,974622695435517953,1715563230,0,0),(974964415177367553,974622695435517953,1715566830,0,0),(974964415182950401,974622695435517953,1715570430,0,0),(974964415187251201,974622695435517953,1715574030,0,0),(974964415191146497,974622695435517953,1715577630,0,0),(974964415196094465,974622695435517953,1715581230,0,0),(974964415200788481,974622695435517953,1715584830,0,0),(974964415201943553,974622695435517953,1715592030,0,0),(974964415204524033,974622695435517953,1715588430,0,0),(974964415207190529,974622695435517953,1715595630,0,0),(974964415213322241,974622695435517953,1715599230,0,0),(974964415215345665,974622695435517953,1715602830,0,0),(974964415220776961,974622695435517953,1715606430,0,0),(974964415223013377,974622695435517953,1715613630,0,0),(974964415225757697,974622695435517953,1715610030,0,0),(974964415230640129,974622695435517953,1715617230,0,0),(974964415233998849,974622695435517953,1715620830,0,0),(974964415236960257,974622695435517953,1715628030,0,0),(974964415236964353,974622695435517953,1715624430,0,0),(974964415241490433,974622695435517953,1715631630,0,0),(974964415241768961,974622695435517953,1715635230,0,0),(974964415247659009,974622695435517953,1715638830,0,0),(974964415250604033,974622695435517953,1715642430,0,0),(974964415253553153,974622695435517953,1715649630,0,0),(974964415255797761,974622695435517953,1715646030,0,0),(974964415259668481,974622695435517953,1715653230,0,0),(974964415261683713,974622695435517953,1715656830,0,0),(974964415263903745,974622695435517953,1715660430,0,0),(974964415265488897,974622695435517953,1715667630,0,0),(974964415268515841,974622695435517953,1715664030,0,0),(974964415269953537,974622695435517953,1715671230,0,0),(974964415275290625,974622695435517953,1715674830,0,0),(974965662960644097,974622695435517953,1715505928,0,0),(974965662978838529,974622695435517953,1715509528,0,0),(974965662982148097,974622695435517953,1715513128,0,0),(974965662988738561,974622695435517953,1715516728,0,0),(974965662993018881,974622695435517953,1715520328,0,0),(974965662997630977,974622695435517953,1715523928,0,0),(974965663000248321,974622695435517953,1715527528,0,0),(974965663010570241,974622695435517953,1715531128,0,0),(974965663012020225,974622695435517953,1715538328,0,0),(974965663015047169,974622695435517953,1715534728,0,0),(974965663015636993,974622695435517953,1715541928,0,0),(974965663021318145,974622695435517953,1715545528,0,0),(974965663025713153,974622695435517953,1715549128,0,0),(974965663030743041,974622695435517953,1715552728,0,0),(974965663033458689,974622695435517953,1715559928,0,0),(974965663034880001,974622695435517953,1715556328,0,0),(974965663042658305,974622695435517953,1715563528,0,0),(974965663045742593,974622695435517953,1715567128,0,0),(974965663053037569,974622695435517953,1715570728,0,0),(974965663055044609,974622695435517953,1715574328,0,0),(974965663056687105,974622695435517953,1715577928,0,0),(974965663059984385,974622695435517953,1715581528,0,0),(974965663069634561,974622695435517953,1715585128,0,0),(974965663070294017,974622695435517953,1715588728,0,0),(974965663075311617,974622695435517953,1715592328,0,0),(974965663076630529,974622695435517953,1715595928,0,0),(974965663079264257,974622695435517953,1715599528,0,0),(974965663085113345,974622695435517953,1715603128,0,0),(974965663088406529,974622695435517953,1715606728,0,0),(974965663091691521,974622695435517953,1715610328,0,0),(974965663091982337,974622695435517953,1715613928,0,0),(974965663096307713,974622695435517953,1715617528,0,0),(974965663102341121,974622695435517953,1715624728,0,0),(974965663102861313,974622695435517953,1715621128,0,0),(974965663104995329,974622695435517953,1715628328,0,0),(974965663108431873,974622695435517953,1715631928,0,0),(974965663111495681,974622695435517953,1715635528,0,0),(974965663112450049,974622695435517953,1715639128,0,0),(974965663116591105,974622695435517953,1715642728,0,0),(974965663121424385,974622695435517953,1715646328,0,0),(974965663125114881,974622695435517953,1715649928,0,0),(974965663128936449,974622695435517953,1715653528,0,0),(974965663130218497,974622695435517953,1715657128,0,0),(974965663136595969,974622695435517953,1715660728,0,0),(974965663141433345,974622695435517953,1715664328,0,0),(974965663142526977,974622695435517953,1715667928,0,0),(974965663145500673,974622695435517953,1715671528,0,0),(974965663148855297,974622695435517953,1715675128,0,0),(974965975914573825,974622695435517953,1715506002,0,0),(974965975915847681,974622695435517953,1715509602,0,0),(974965975922032641,974622695435517953,1715513202,0,0),(974965975923707905,974622695435517953,1715516802,0,0),(974965975927578625,974622695435517953,1715524002,0,0),(974965975929036801,974622695435517953,1715520402,0,0),(974965975935090689,974622695435517953,1715527602,0,0),(974965975936188417,974622695435517953,1715531202,0,0),(974965975938904065,974622695435517953,1715534802,0,0),(974965975940796417,974622695435517953,1715542002,0,0),(974965975943032833,974622695435517953,1715538402,0,0),(974965975943827457,974622695435517953,1715545602,0,0),(974965975944818689,974622695435517953,1715549202,0,0),(974965975947476993,974622695435517953,1715552802,0,0),(974965975950577665,974622695435517953,1715556402,0,0),(974965975955767297,974622695435517953,1715563602,0,0),(974965975956381697,974622695435517953,1715560002,0,0),(974965975957094401,974622695435517953,1715570802,0,0),(974965975959347201,974622695435517953,1715567202,0,0),(974965975964217345,974622695435517953,1715574402,0,0),(974965975965761537,974622695435517953,1715581602,0,0),(974965975966605313,974622695435517953,1715578002,0,0),(974965975969701889,974622695435517953,1715588802,0,0),(974965975972777985,974622695435517953,1715585202,0,0),(974965975973654529,974622695435517953,1715592402,0,0),(974965975979192321,974622695435517953,1715596002,0,0),(974965975980539905,974622695435517953,1715599602,0,0),(974965975983611905,974622695435517953,1715603202,0,0),(974965975986114561,974622695435517953,1715606802,0,0),(974965975986302977,974622695435517953,1715614002,0,0),(974965975987138561,974622695435517953,1715610402,0,0),(974965975990771713,974622695435517953,1715621202,0,0),(974965975990824961,974622695435517953,1715617602,0,0),(974965975995396097,974622695435517953,1715624802,0,0),(974965975995678721,974622695435517953,1715628402,0,0),(974965976000167937,974622695435517953,1715635602,0,0),(974965976001191937,974622695435517953,1715632002,0,0),(974965976006791169,974622695435517953,1715639202,0,0),(974965976006938625,974622695435517953,1715642802,0,0),(974965976010035201,974622695435517953,1715646402,0,0),(974965976011005953,974622695435517953,1715650002,0,0),(974965976011767809,974622695435517953,1715653602,0,0),(974965976017035265,974622695435517953,1715660802,0,0),(974965976017305601,974622695435517953,1715657202,0,0),(974965976020856833,974622695435517953,1715668002,0,0),(974965976021889025,974622695435517953,1715664402,0,0),(974965976025800705,974622695435517953,1715671602,0,0),(974965976026591233,974622695435517953,1715675202,0,0),(974976119791525889,974622695435517953,1715508421,0,0),(974976119811162113,974622695435517953,1715512021,0,0),(974976119821422593,974622695435517953,1715515621,0,0),(974976119825637377,974622695435517953,1715519221,0,0),(974976119827996673,974622695435517953,1715522821,0,0),(974976119834357761,974622695435517953,1715526421,0,0),(974976119837753345,974622695435517953,1715530021,0,0),(974976119840591873,974622695435517953,1715533621,0,0),(974976119847370753,974622695435517953,1715537221,0,0),(974976119851458561,974622695435517953,1715540821,0,0),(974976119853785089,974622695435517953,1715544421,0,0),(974976119858446337,974622695435517953,1715548021,0,0),(974976119860563969,974622695435517953,1715551621,0,0),(974976119865077761,974622695435517953,1715555221,0,0),(974976119867002881,974622695435517953,1715558821,0,0),(974976119872339969,974622695435517953,1715562421,0,0),(974976119877742593,974622695435517953,1715566021,0,0),(974976119884353537,974622695435517953,1715569621,0,0),(974976119886151681,974622695435517953,1715576821,0,0),(974976119889014785,974622695435517953,1715573221,0,0),(974976119893295105,974622695435517953,1715580421,0,0),(974976119894245377,974622695435517953,1715584021,0,0),(974976119899611137,974622695435517953,1715587621,0,0),(974976119902597121,974622695435517953,1715591221,0,0),(974976119908007937,974622695435517953,1715598421,0,0),(974976119908208641,974622695435517953,1715594821,0,0),(974976119912980481,974622695435517953,1715602021,0,0),(974976119915548673,974622695435517953,1715605621,0,0),(974976119926599681,974622695435517953,1715609221,0,0),(974976119928598529,974622695435517953,1715612821,0,0),(974976119931891713,974622695435517953,1715616421,0,0),(974976119932919809,974622695435517953,1715620021,0,0),(974976119938916353,974622695435517953,1715627221,0,0),(974976119939350529,974622695435517953,1715623621,0,0),(974976119942639617,974622695435517953,1715630821,0,0),(974976119945015297,974622695435517953,1715638021,0,0),(974976119947481089,974622695435517953,1715634421,0,0),(974976119950602241,974622695435517953,1715645221,0,0),(974976119951937537,974622695435517953,1715641621,0,0),(974976119952547841,974622695435517953,1715648821,0,0),(974976119956738049,974622695435517953,1715652421,0,0),(974976119960977409,974622695435517953,1715659621,0,0),(974976119964418049,974622695435517953,1715656021,0,0),(974976119967236097,974622695435517953,1715666821,0,0),(974976119967260673,974622695435517953,1715663221,0,0),(974976119970967553,974622695435517953,1715670421,0,0),(974976119973343233,974622695435517953,1715674021,0,0),(974976119976181761,974622695435517953,1715677621,0,0),(974980403637506049,974622695435517953,1715509442,0,0),(974980403654144001,974622695435517953,1715513042,0,0),(974980403662278657,974622695435517953,1715516642,0,0),(974980403672215553,974622695435517953,1715520242,0,0),(974980403689111553,974622695435517953,1715523842,0,0),(974980403696332801,974622695435517953,1715527442,0,0),(974980403704725505,974622695435517953,1715531042,0,0),(974980403709837313,974622695435517953,1715534642,0,0),(974980403720454145,974622695435517953,1715538242,0,0),(974980403726000129,974622695435517953,1715541842,0,0),(974980403730518017,974622695435517953,1715545442,0,0),(974980403740979201,974622695435517953,1715549042,0,0),(974980403748769793,974622695435517953,1715552642,0,0),(974980403755720705,974622695435517953,1715556242,0,0),(974980403760680961,974622695435517953,1715559842,0,0),(974980403764563969,974622695435517953,1715563442,0,0),(974980403769634817,974622695435517953,1715567042,0,0),(974980403772276737,974622695435517953,1715570642,0,0),(974980403779125249,974622695435517953,1715574242,0,0),(974980403781681153,974622695435517953,1715577842,0,0),(974980403785740289,974622695435517953,1715581442,0,0),(974980403788546049,974622695435517953,1715585042,0,0),(974980403793596417,974622695435517953,1715588642,0,0),(974980403798212609,974622695435517953,1715595842,0,0),(974980403798999041,974622695435517953,1715592242,0,0),(974980403801186305,974622695435517953,1715599442,0,0),(974980403807584257,974622695435517953,1715603042,0,0),(974980403810267137,974622695435517953,1715606642,0,0),(974980403814612993,974622695435517953,1715610242,0,0),(974980403823017985,974622695435517953,1715613842,0,0),(974980403825565697,974622695435517953,1715617442,0,0),(974980403829891073,974622695435517953,1715621042,0,0),(974980403832569857,974622695435517953,1715624642,0,0),(974980403834269697,974622695435517953,1715631842,0,0),(974980403835564033,974622695435517953,1715628242,0,0),(974980403844100097,974622695435517953,1715639042,0,0),(974980403846062081,974622695435517953,1715635442,0,0),(974980403847794689,974622695435517953,1715642642,0,0),(974980403853438977,974622695435517953,1715646242,0,0),(974980403856109569,974622695435517953,1715653442,0,0),(974980403858513921,974622695435517953,1715649842,0,0),(974980403860738049,974622695435517953,1715657042,0,0),(974980403866210305,974622695435517953,1715664242,0,0),(974980403867521025,974622695435517953,1715660642,0,0),(974980403871096833,974622695435517953,1715667842,0,0),(974980403871178753,974622695435517953,1715671442,0,0),(974980403872755713,974622695435517953,1715675042,0,0),(974980403879739393,974622695435517953,1715678642,0,0),(974980941445480449,974622695435517953,1715509570,0,0),(974980941453135873,974622695435517953,1715513170,0,0),(974980941459927041,974622695435517953,1715516770,0,0),(974980941462020097,974622695435517953,1715520370,0,0),(974980941463998465,974622695435517953,1715523970,0,0),(974980941464834049,974622695435517953,1715527570,0,0),(974980941466386433,974622695435517953,1715531170,0,0),(974980941470609409,974622695435517953,1715534770,0,0),(974980941471350785,974622695435517953,1715538370,0,0),(974980941474398209,974622695435517953,1715541970,0,0),(974980941480509441,974622695435517953,1715545570,0,0),(974980941482094593,974622695435517953,1715549170,0,0),(974980941483003905,974622695435517953,1715552770,0,0),(974980941488914433,974622695435517953,1715556370,0,0),(974980941491957761,974622695435517953,1715563570,0,0),(974980941492670465,974622695435517953,1715559970,0,0),(974980941494001665,974622695435517953,1715567170,0,0),(974980941498494977,974622695435517953,1715570770,0,0),(974980941507796993,974622695435517953,1715574370,0,0),(974980941512814593,974622695435517953,1715577970,0,0),(974980941515132929,974622695435517953,1715585170,0,0),(974980941517099009,974622695435517953,1715581570,0,0),(974980941520523265,974622695435517953,1715588770,0,0),(974980941525966849,974622695435517953,1715592370,0,0),(974980941527707649,974622695435517953,1715595970,0,0),(974980941529559041,974622695435517953,1715599570,0,0),(974980941535305729,974622695435517953,1715603170,0,0),(974980941535862785,974622695435517953,1715610370,0,0),(974980941536759809,974622695435517953,1715606770,0,0),(974980941545123841,974622695435517953,1715613970,0,0),(974980941550186497,974622695435517953,1715621170,0,0),(974980941551165441,974622695435517953,1715617570,0,0),(974980941555003393,974622695435517953,1715624770,0,0),(974980941557698561,974622695435517953,1715628370,0,0),(974980941562073089,974622695435517953,1715635570,0,0),(974980941562277889,974622695435517953,1715631970,0,0),(974980941565333505,974622695435517953,1715642770,0,0),(974980941566455809,974622695435517953,1715639170,0,0),(974980941569089537,974622695435517953,1715646370,0,0),(974980941571051521,974622695435517953,1715653570,0,0),(974980941573124097,974622695435517953,1715649970,0,0),(974980941575344129,974622695435517953,1715660770,0,0),(974980941575651329,974622695435517953,1715657170,0,0),(974980941578092545,974622695435517953,1715667970,0,0),(974980941578252289,974622695435517953,1715664370,0,0),(974980941582331905,974622695435517953,1715678770,0,0),(974980941583265793,974622695435517953,1715671570,0,0),(974980941583556609,974622695435517953,1715675170,0,0),(974983644462977025,974622695435517953,1715510215,0,0),(974983644481482753,974622695435517953,1715513815,0,0),(974983644488744961,974622695435517953,1715517415,0,0),(974983644494987265,974622695435517953,1715521015,0,0),(974983644500041729,974622695435517953,1715524615,0,0),(974983644501135361,974622695435517953,1715528215,0,0),(974983644506947585,974622695435517953,1715531815,0,0),(974983644509757441,974622695435517953,1715535415,0,0),(974983644509945857,974622695435517953,1715539015,0,0),(974983644513673217,974622695435517953,1715542615,0,0),(974983644517171201,974622695435517953,1715546215,0,0),(974983644522491905,974622695435517953,1715549815,0,0),(974983644528435201,974622695435517953,1715553415,0,0),(974983644534353921,974622695435517953,1715557015,0,0),(974983644541685761,974622695435517953,1715560615,0,0),(974983644548501505,974622695435517953,1715564215,0,0),(974983644553564161,974622695435517953,1715567815,0,0),(974983644553936897,974622695435517953,1715571415,0,0),(974983644555796481,974622695435517953,1715575015,0,0),(974983644562378753,974622695435517953,1715578615,0,0),(974983644571267073,974622695435517953,1715582215,0,0),(974983644579389441,974622695435517953,1715585815,0,0),(974983644580683777,974622695435517953,1715589415,0,0),(974983644587012097,974622695435517953,1715593015,0,0),(974983644589461505,974622695435517953,1715596615,0,0),(974983644596109313,974622695435517953,1715600215,0,0),(974983644596473857,974622695435517953,1715603815,0,0),(974983644600815617,974622695435517953,1715607415,0,0),(974983644604166145,974622695435517953,1715611015,0,0),(974983644606889985,974622695435517953,1715614615,0,0),(974983644606910465,974622695435517953,1715618215,0,0),(974983644611084289,974622695435517953,1715621815,0,0),(974983644614021121,974622695435517953,1715629015,0,0),(974983644615303169,974622695435517953,1715625415,0,0),(974983644621680641,974622695435517953,1715632615,0,0),(974983644624322561,974622695435517953,1715636215,0,0),(974983644635328513,974622695435517953,1715639815,0,0),(974983644635906049,974622695435517953,1715643415,0,0),(974983644641792001,974622695435517953,1715647015,0,0),(974983644642619393,974622695435517953,1715650615,0,0),(974983644643680257,974622695435517953,1715657815,0,0),(974983644646109185,974622695435517953,1715654215,0,0),(974983644649283585,974622695435517953,1715661415,0,0),(974983644649684993,974622695435517953,1715665015,0,0),(974983644653903873,974622695435517953,1715668615,0,0),(974983644656185345,974622695435517953,1715672215,0,0),(974983644656476161,974622695435517953,1715675815,0,0),(974983644662206465,974622695435517953,1715679415,0,0),(974996406069231617,974622695435517953,1715513257,0,0),(974996406088822785,974622695435517953,1715516857,0,0),(974996406092726273,974622695435517953,1715520457,0,0),(974996406094155777,974622695435517953,1715527657,0,0),(974996406097018881,974622695435517953,1715524057,0,0),(974996406100606977,974622695435517953,1715531257,0,0),(974996406102659073,974622695435517953,1715534857,0,0),(974996406109454337,974622695435517953,1715538457,0,0),(974996406112927745,974622695435517953,1715542057,0,0),(974996406115459073,974622695435517953,1715545657,0,0),(974996406123761665,974622695435517953,1715549257,0,0),(974996406127247361,974622695435517953,1715552857,0,0),(974996406127824897,974622695435517953,1715556457,0,0),(974996406131789825,974622695435517953,1715560057,0,0),(974996406136369153,974622695435517953,1715563657,0,0),(974996406142074881,974622695435517953,1715567257,0,0),(974996406144147457,974622695435517953,1715570857,0,0),(974996406145802241,974622695435517953,1715574457,0,0),(974996406149971969,974622695435517953,1715578057,0,0),(974996406153658369,974622695435517953,1715581657,0,0),(974996406158503937,974622695435517953,1715585257,0,0),(974996406161285121,974622695435517953,1715588857,0,0),(974996406168956929,974622695435517953,1715592457,0,0),(974996406170202113,974622695435517953,1715596057,0,0),(974996406173782017,974622695435517953,1715599657,0,0),(974996406180605953,974622695435517953,1715603257,0,0),(974996406182940673,974622695435517953,1715606857,0,0),(974996406188679169,974622695435517953,1715610457,0,0),(974996406190604289,974622695435517953,1715614057,0,0),(974996406197301249,974622695435517953,1715617657,0,0),(974996406202548225,974622695435517953,1715621257,0,0),(974996406205861889,974622695435517953,1715624857,0,0),(974996406209978369,974622695435517953,1715628457,0,0),(974996406211502081,974622695435517953,1715632057,0,0),(974996406215069697,974622695435517953,1715635657,0,0),(974996406221398017,974622695435517953,1715642857,0,0),(974996406221484033,974622695435517953,1715639257,0,0),(974996406226337793,974622695435517953,1715646457,0,0),(974996406228721665,974622695435517953,1715653657,0,0),(974996406230458369,974622695435517953,1715650057,0,0),(974996406232612865,974622695435517953,1715657257,0,0),(974996406236794881,974622695435517953,1715660857,0,0),(974996406244409345,974622695435517953,1715664457,0,0),(974996406244614145,974622695435517953,1715668057,0,0),(974996406245789697,974622695435517953,1715671657,0,0),(974996406250541057,974622695435517953,1715675257,0,0),(974996406255325185,974622695435517953,1715678857,0,0),(974996406255783937,974622695435517953,1715682457,0,0),(974997747410698241,974622695435517953,1715513577,0,0),(974997747416428545,974622695435517953,1715517177,0,0),(974997747425017857,974622695435517953,1715520777,0,0),(974997747430375425,974622695435517953,1715527977,0,0),(974997747431321601,974622695435517953,1715524377,0,0),(974997747438727169,974622695435517953,1715531577,0,0),(974997747443109889,974622695435517953,1715535177,0,0),(974997747443339265,974622695435517953,1715538777,0,0),(974997747449044993,974622695435517953,1715545977,0,0),(974997747451973633,974622695435517953,1715542377,0,0),(974997747453366273,974622695435517953,1715549577,0,0),(974997747460218881,974622695435517953,1715553177,0,0),(974997747464953857,974622695435517953,1715556777,0,0),(974997747469598721,974622695435517953,1715560377,0,0),(974997747471351809,974622695435517953,1715563977,0,0),(974997747475365889,974622695435517953,1715571177,0,0),(974997747478179841,974622695435517953,1715567577,0,0),(974997747479465985,974622695435517953,1715574777,0,0),(974997747489046529,974622695435517953,1715578377,0,0),(974997747491323905,974622695435517953,1715585577,0,0),(974997747491356673,974622695435517953,1715581977,0,0),(974997747495448577,974622695435517953,1715589177,0,0),(974997747502628865,974622695435517953,1715592777,0,0),(974997747504893953,974622695435517953,1715596377,0,0),(974997747505627137,974622695435517953,1715599977,0,0),(974997747513524225,974622695435517953,1715603577,0,0),(974997747518889985,974622695435517953,1715607177,0,0),(974997747523911681,974622695435517953,1715610777,0,0),(974997747528445953,974622695435517953,1715614377,0,0),(974997747533565953,974622695435517953,1715617977,0,0),(974997747540205569,974622695435517953,1715621577,0,0),(974997747541995521,974622695435517953,1715628777,0,0),(974997747542523905,974622695435517953,1715625177,0,0),(974997747546603521,974622695435517953,1715632377,0,0),(974997747553808385,974622695435517953,1715635977,0,0),(974997747555958785,974622695435517953,1715639577,0,0),(974997747560206337,974622695435517953,1715643177,0,0),(974997747564285953,974622695435517953,1715646777,0,0),(974997747569766401,974622695435517953,1715650377,0,0),(974997747570573313,974622695435517953,1715653977,0,0),(974997747577004033,974622695435517953,1715661177,0,0),(974997747577778177,974622695435517953,1715657577,0,0),(974997747579305985,974622695435517953,1715664777,0,0),(974997747583692801,974622695435517953,1715668377,0,0),(974997747588296705,974622695435517953,1715671977,0,0),(974997747593285633,974622695435517953,1715675577,0,0),(974997747602890753,974622695435517953,1715679177,0,0),(974997747603972097,974622695435517953,1715682777,0,0),(974999197259505665,974622695435517953,1715513923,0,0),(974999197273108481,974622695435517953,1715517523,0,0),(974999197285175297,974622695435517953,1715521123,0,0),(974999197287571457,974622695435517953,1715524723,0,0),(974999197296836609,974622695435517953,1715528323,0,0),(974999197300690945,974622695435517953,1715531923,0,0),(974999197308477441,974622695435517953,1715535523,0,0),(974999197311533057,974622695435517953,1715539123,0,0),(974999197315301377,974622695435517953,1715542723,0,0),(974999197323034625,974622695435517953,1715546323,0,0),(974999197325889537,974622695435517953,1715549923,0,0),(974999197329244161,974622695435517953,1715553523,0,0),(974999197339762689,974622695435517953,1715557123,0,0),(974999197343080449,974622695435517953,1715560723,0,0),(974999197344342017,974622695435517953,1715564323,0,0),(974999197355311105,974622695435517953,1715567923,0,0),(974999197358940161,974622695435517953,1715571523,0,0),(974999197369393153,974622695435517953,1715575123,0,0),(974999197373652993,974622695435517953,1715578723,0,0),(974999197374517249,974622695435517953,1715582323,0,0),(974999197381378049,974622695435517953,1715585923,0,0),(974999197383512065,974622695435517953,1715589523,0,0),(974999197386366977,974622695435517953,1715593123,0,0),(974999197387984897,974622695435517953,1715596723,0,0),(974999197393174529,974622695435517953,1715600323,0,0),(974999197395197953,974622695435517953,1715607523,0,0),(974999197395533825,974622695435517953,1715603923,0,0),(974999197398937601,974622695435517953,1715614723,0,0),(974999197399040001,974622695435517953,1715611123,0,0),(974999197405339649,974622695435517953,1715618323,0,0),(974999197410033665,974622695435517953,1715621923,0,0),(974999197410304001,974622695435517953,1715625523,0,0),(974999197415354369,974622695435517953,1715629123,0,0),(974999197418598401,974622695435517953,1715632723,0,0),(974999197423042561,974622695435517953,1715636323,0,0),(974999197425344513,974622695435517953,1715639923,0,0),(974999197430439937,974622695435517953,1715643523,0,0),(974999197436162049,974622695435517953,1715647123,0,0),(974999197437661185,974622695435517953,1715650723,0,0),(974999197437796353,974622695435517953,1715654323,0,0),(974999197441617921,974622695435517953,1715657923,0,0),(974999197442789377,974622695435517953,1715661523,0,0),(974999197445001217,974622695435517953,1715665123,0,0),(974999197450579969,974622695435517953,1715668723,0,0),(974999197455581185,974622695435517953,1715672323,0,0),(974999197456683009,974622695435517953,1715675923,0,0),(974999197458210817,974622695435517953,1715679523,0,0),(974999197464981505,974622695435517953,1715683123,0,0),(974999688878690305,974622695435517953,1715514040,0,0),(974999688894873601,974622695435517953,1715517640,0,0),(974999688899358721,974622695435517953,1715521240,0,0),(974999688905433089,974622695435517953,1715524840,0,0),(974999688910626817,974622695435517953,1715528440,0,0),(974999688922374145,974622695435517953,1715532040,0,0),(974999688924311553,974622695435517953,1715535640,0,0),(974999688929775617,974622695435517953,1715539240,0,0),(974999688933826561,974622695435517953,1715542840,0,0),(974999688939614209,974622695435517953,1715550040,0,0),(974999688940781569,974622695435517953,1715546440,0,0),(974999688943894529,974622695435517953,1715553640,0,0),(974999688947494913,974622695435517953,1715557240,0,0),(974999688952950785,974622695435517953,1715560840,0,0),(974999688953241601,974622695435517953,1715564440,0,0),(974999688954667009,974622695435517953,1715568040,0,0),(974999688962834433,974622695435517953,1715571640,0,0),(974999688964562945,974622695435517953,1715575240,0,0),(974999688967356417,974622695435517953,1715578840,0,0),(974999688974872577,974622695435517953,1715582440,0,0),(974999688978505729,974622695435517953,1715586040,0,0),(974999688982188033,974622695435517953,1715589640,0,0),(974999688986869761,974622695435517953,1715593240,0,0),(974999688990212097,974622695435517953,1715596840,0,0),(974999688990273537,974622695435517953,1715600440,0,0),(974999688994963457,974622695435517953,1715604040,0,0),(974999688999972865,974622695435517953,1715607640,0,0),(974999689000960001,974622695435517953,1715611240,0,0),(974999689005998081,974622695435517953,1715614840,0,0),(974999689009590273,974622695435517953,1715618440,0,0),(974999689012838401,974622695435517953,1715622040,0,0),(974999689013374977,974622695435517953,1715625640,0,0),(974999689021427713,974622695435517953,1715629240,0,0),(974999689023426561,974622695435517953,1715632840,0,0),(974999689027727361,974622695435517953,1715636440,0,0),(974999689032364033,974622695435517953,1715640040,0,0),(974999689034653697,974622695435517953,1715643640,0,0),(974999689035177985,974622695435517953,1715647240,0,0),(974999689043820545,974622695435517953,1715650840,0,0),(974999689045245953,974622695435517953,1715654440,0,0),(974999689049366529,974622695435517953,1715658040,0,0),(974999689052635137,974622695435517953,1715661640,0,0),(974999689054793729,974622695435517953,1715665240,0,0),(974999689060102145,974622695435517953,1715672440,0,0),(974999689062526977,974622695435517953,1715668840,0,0),(974999689067175937,974622695435517953,1715679640,0,0),(974999689067372545,974622695435517953,1715676040,0,0),(974999689070694401,974622695435517953,1715683240,0,0),(975000034615005185,974622695435517953,1715514122,0,0),(975000034625187841,974622695435517953,1715517722,0,0),(975000034632151041,974622695435517953,1715521322,0,0),(975000034640277505,974622695435517953,1715524922,0,0),(975000034649464833,974622695435517953,1715528522,0,0),(975000034651713537,974622695435517953,1715532122,0,0),(975000034661261313,974622695435517953,1715535722,0,0),(975000034664792065,974622695435517953,1715539322,0,0),(975000034672259073,974622695435517953,1715542922,0,0),(975000034674872321,974622695435517953,1715546522,0,0),(975000034681987073,974622695435517953,1715550122,0,0),(975000034683117569,974622695435517953,1715553722,0,0),(975000034688589825,974622695435517953,1715557322,0,0),(975000034699116545,974622695435517953,1715560922,0,0),(975000034703323137,974622695435517953,1715564522,0,0),(975000034707226625,974622695435517953,1715568122,0,0),(975000034710110209,974622695435517953,1715571722,0,0),(975000034714198017,974622695435517953,1715575322,0,0),(975000034719903745,974622695435517953,1715578922,0,0),(975000034721095681,974622695435517953,1715586122,0,0),(975000034721869825,974622695435517953,1715582522,0,0),(975000034727030785,974622695435517953,1715589722,0,0),(975000034728869889,974622695435517953,1715593322,0,0),(975000034734874625,974622695435517953,1715596922,0,0),(975000034743803905,974622695435517953,1715600522,0,0),(975000034749431809,974622695435517953,1715604122,0,0),(975000034752167937,974622695435517953,1715607722,0,0),(975000034753036289,974622695435517953,1715611322,0,0),(975000034753933313,974622695435517953,1715614922,0,0),(975000034758115329,974622695435517953,1715622122,0,0),(975000034761273345,974622695435517953,1715618522,0,0),(975000034766368769,974622695435517953,1715625722,0,0),(975000034768474113,974622695435517953,1715632922,0,0),(975000034769362945,974622695435517953,1715629322,0,0),(975000034772410369,974622695435517953,1715636522,0,0),(975000034772877313,974622695435517953,1715640122,0,0),(975000034777223169,974622695435517953,1715647322,0,0),(975000034777698305,974622695435517953,1715643722,0,0),(975000034781458433,974622695435517953,1715650922,0,0),(975000034784022529,974622695435517953,1715654522,0,0),(975000034784497665,974622695435517953,1715658122,0,0),(975000034787889153,974622695435517953,1715665322,0,0),(975000034791481345,974622695435517953,1715661722,0,0),(975000034792869889,974622695435517953,1715672522,0,0),(975000034795470849,974622695435517953,1715668922,0,0),(975000034798338049,974622695435517953,1715676122,0,0),(975000034798661633,974622695435517953,1715679722,0,0),(975000034802319361,974622695435517953,1715683322,0,0),(975000620364234753,974622695435517953,1715514262,0,0),(975000620371845121,974622695435517953,1715521462,0,0),(975000620372504577,974622695435517953,1715517862,0,0),(975000620379287553,974622695435517953,1715525062,0,0),(975000620379607041,974622695435517953,1715528662,0,0),(975000620384960513,974622695435517953,1715535862,0,0),(975000620387667969,974622695435517953,1715532262,0,0),(975000620390420481,974622695435517953,1715539462,0,0),(975000620393467905,974622695435517953,1715543062,0,0),(975000620398395393,974622695435517953,1715550262,0,0),(975000620399751169,974622695435517953,1715546662,0,0),(975000620404232193,974622695435517953,1715553862,0,0),(975000620407169025,974622695435517953,1715557462,0,0),(975000620412657665,974622695435517953,1715561062,0,0),(975000620414275585,974622695435517953,1715564662,0,0),(975000620421255169,974622695435517953,1715568262,0,0),(975000620423360513,974622695435517953,1715575462,0,0),(975000620423540737,974622695435517953,1715571862,0,0),(975000620427698177,974622695435517953,1715579062,0,0),(975000620433022977,974622695435517953,1715582662,0,0),(975000620436758529,974622695435517953,1715586262,0,0),(975000620438028289,974622695435517953,1715589862,0,0),(975000620438913025,974622695435517953,1715593462,0,0),(975000620443983873,974622695435517953,1715597062,0,0),(975000620449886209,974622695435517953,1715600662,0,0),(975000620452327425,974622695435517953,1715604262,0,0),(975000620456452097,974622695435517953,1715607862,0,0),(975000620458483713,974622695435517953,1715611462,0,0),(975000620461678593,974622695435517953,1715615062,0,0),(975000620463448065,974622695435517953,1715618662,0,0),(975000620470730753,974622695435517953,1715622262,0,0),(975000620473266177,974622695435517953,1715629462,0,0),(975000620473458689,974622695435517953,1715625862,0,0),(975000620477198337,974622695435517953,1715633062,0,0),(975000620482981889,974622695435517953,1715640262,0,0),(975000620483043329,974622695435517953,1715636662,0,0),(975000620488876033,974622695435517953,1715647462,0,0),(975000620490911745,974622695435517953,1715643862,0,0),(975000620492730369,974622695435517953,1715651062,0,0),(975000620493496321,974622695435517953,1715654662,0,0),(975000620499755009,974622695435517953,1715658262,0,0),(975000620502835201,974622695435517953,1715661862,0,0),(975000620504903681,974622695435517953,1715665462,0,0),(975000620505243649,974622695435517953,1715672662,0,0),(975000620508319745,974622695435517953,1715669062,0,0),(975000620510380033,974622695435517953,1715676262,0,0),(975000620511719425,974622695435517953,1715683462,0,0),(975000620512460801,974622695435517953,1715679862,0,0),(975015721671913473,974622695435517953,1715517862,0,0),(975015721687855105,974622695435517953,1715521462,0,0),(975015721693589505,974622695435517953,1715525062,0,0),(975015721696681985,974622695435517953,1715528662,0,0),(975015721705562113,974622695435517953,1715532262,0,0),(975015721710862337,974622695435517953,1715535862,0,0),(975015721712427009,974622695435517953,1715539462,0,0),(975015721715286017,974622695435517953,1715543062,0,0),(975015721725517825,974622695435517953,1715546662,0,0),(975015721731719169,974622695435517953,1715550262,0,0),(975015721732046849,974622695435517953,1715553862,0,0),(975015721737666561,974622695435517953,1715557462,0,0),(975015721740570625,974622695435517953,1715561062,0,0),(975015721748639745,974622695435517953,1715564662,0,0),(975015721751945217,974622695435517953,1715568262,0,0),(975015721756770305,974622695435517953,1715571862,0,0),(975015721760690177,974622695435517953,1715575462,0,0),(975015721762947073,974622695435517953,1715579062,0,0),(975015721763069953,974622695435517953,1715582662,0,0),(975015721767329793,974622695435517953,1715586262,0,0),(975015721772609537,974622695435517953,1715593462,0,0),(975015721773637633,974622695435517953,1715589862,0,0),(975015721776304129,974622695435517953,1715597062,0,0),(975015721778270209,974622695435517953,1715600662,0,0),(975015721789657089,974622695435517953,1715604262,0,0),(975015721794658305,974622695435517953,1715607862,0,0),(975015721802932225,974622695435517953,1715611462,0,0),(975015721805828097,974622695435517953,1715615062,0,0),(975015721814839297,974622695435517953,1715618662,0,0),(975015721819770881,974622695435517953,1715622262,0,0),(975015721824092161,974622695435517953,1715625862,0,0),(975015721826357249,974622695435517953,1715629462,0,0),(975015721830158337,974622695435517953,1715633062,0,0),(975015721832296449,974622695435517953,1715636662,0,0),(975015721833205761,974622695435517953,1715640262,0,0),(975015721840128001,974622695435517953,1715643862,0,0),(975015721840291841,974622695435517953,1715647462,0,0),(975015721842184193,974622695435517953,1715651062,0,0),(975015721842724865,974622695435517953,1715654662,0,0),(975015721845391361,974622695435517953,1715658262,0,0),(975015721850384385,974622695435517953,1715665462,0,0),(975015721852833793,974622695435517953,1715661862,0,0),(975015721856647169,974622695435517953,1715672662,0,0),(975015721856901121,974622695435517953,1715669062,0,0),(975015721859821569,974622695435517953,1715676262,0,0),(975015721862320129,974622695435517953,1715679862,0,0),(975015721866629121,974622695435517953,1715683462,0,0),(975015721873154049,974622695435517953,1715687062,0,0),(977013920325144577,974622695435517953,1715994270,0,0),(977013920334614529,974622695435517953,1715997870,0,0),(977013920337207297,974622695435517953,1716005070,0,0),(977013920339554305,974622695435517953,1716001470,0,0),(977013920342425601,974622695435517953,1716008670,0,0),(977013920344715265,974622695435517953,1716015870,0,0),(977013920345280513,974622695435517953,1716012270,0,0),(977013920351457281,974622695435517953,1716023070,0,0),(977013920351875073,974622695435517953,1716019470,0,0),(977013920353259521,974622695435517953,1716026670,0,0),(977013920354435073,974622695435517953,1716030270,0,0),(977013920358969345,974622695435517953,1716033870,0,0),(977013920359583745,974622695435517953,1716037470,0,0),(977013920362635265,974622695435517953,1716044670,0,0),(977013920364040193,974622695435517953,1716048270,0,0),(977013920364130305,974622695435517953,1716041070,0,0),(977013920367665153,974622695435517953,1716055470,0,0),(977013920368451585,974622695435517953,1716051870,0,0),(977013920370597889,974622695435517953,1716062670,0,0),(977013920370917377,974622695435517953,1716059070,0,0),(977013920374472705,974622695435517953,1716069870,0,0),(977013920375369729,974622695435517953,1716073470,0,0),(977013920376446977,974622695435517953,1716066270,0,0),(977013920379371521,974622695435517953,1716080670,0,0),(977013920381902849,974622695435517953,1716077070,0,0),(977013920382746625,974622695435517953,1716084270,0,0),(977013920384839681,974622695435517953,1716087870,0,0),(977013920387661825,974622695435517953,1716091470,0,0),(977013920391139329,974622695435517953,1716098670,0,0),(977013920393465857,974622695435517953,1716095070,0,0),(977013920397307905,974622695435517953,1716105870,0,0),(977013920398053377,974622695435517953,1716102270,0,0),(977013920400666625,974622695435517953,1716109470,0,0),(977013920402055169,974622695435517953,1716113070,0,0),(977013920403980289,974622695435517953,1716116670,0,0),(977013920408801281,974622695435517953,1716123870,0,0),(977013920409288705,974622695435517953,1716120270,0,0),(977013920412139521,974622695435517953,1716131070,0,0),(977013920412692481,974622695435517953,1716134670,0,0),(977013920415432705,974622695435517953,1716127470,0,0),(977013920417681409,974622695435517953,1716141870,0,0),(977013920419540993,974622695435517953,1716138270,0,0),(977013920421203969,974622695435517953,1716149070,0,0),(977013920422985729,974622695435517953,1716145470,0,0),(977013920424501249,974622695435517953,1716156270,0,0),(977013920428261377,974622695435517953,1716152670,0,0),(977013920428494849,974622695435517953,1716159870,0,0),(977013920431534081,974622695435517953,1716163470,0,0),(977015252578451457,974622695435517953,1715994588,0,0),(977015252591095809,974622695435517953,1715998188,0,0),(977015252598329345,974622695435517953,1716001788,0,0),(977015252605743105,974622695435517953,1716005388,0,0),(977015252609978369,974622695435517953,1716008988,0,0),(977015252611846145,974622695435517953,1716012588,0,0),(977015252617736193,974622695435517953,1716016188,0,0),(977015252623888385,974622695435517953,1716019788,0,0),(977015252629602305,974622695435517953,1716023388,0,0),(977015252635729921,974622695435517953,1716026988,0,0),(977015252636856321,974622695435517953,1716030588,0,0),(977015252642127873,974622695435517953,1716034188,0,0),(977015252647223297,974622695435517953,1716037788,0,0),(977015252654305281,974622695435517953,1716041388,0,0),(977015252659191809,974622695435517953,1716044988,0,0),(977015252664897537,974622695435517953,1716048588,0,0),(977015252665569281,974622695435517953,1716052188,0,0),(977015252671442945,974622695435517953,1716055788,0,0),(977015252677292033,974622695435517953,1716059388,0,0),(977015252678815745,974622695435517953,1716062988,0,0),(977015252685340673,974622695435517953,1716066588,0,0),(977015252693737473,974622695435517953,1716070188,0,0),(977015252695179265,974622695435517953,1716073788,0,0),(977015252699750401,974622695435517953,1716077388,0,0),(977015252702584833,974622695435517953,1716080988,0,0),(977015252704370689,974622695435517953,1716084588,0,0),(977015252708192257,974622695435517953,1716091788,0,0),(977015252709912577,974622695435517953,1716088188,0,0),(977015252715298817,974622695435517953,1716095388,0,0),(977015252716441601,974622695435517953,1716102588,0,0),(977015252719194113,974622695435517953,1716098988,0,0),(977015252720279553,974622695435517953,1716106188,0,0),(977015252725129217,974622695435517953,1716109788,0,0),(977015252725608449,974622695435517953,1716113388,0,0),(977015252729933825,974622695435517953,1716116988,0,0),(977015252732788737,974622695435517953,1716120588,0,0),(977015252735156225,974622695435517953,1716124188,0,0),(977015252737171457,974622695435517953,1716127788,0,0),(977015252742574081,974622695435517953,1716134988,0,0),(977015252742598657,974622695435517953,1716131388,0,0),(977015252747350017,974622695435517953,1716138588,0,0),(977015252751843329,974622695435517953,1716142188,0,0),(977015252752760833,974622695435517953,1716145788,0,0),(977015252754440193,974622695435517953,1716149388,0,0),(977015252757274625,974622695435517953,1716152988,0,0),(977015252757778433,974622695435517953,1716160188,0,0),(977015252759814145,974622695435517953,1716156588,0,0),(977015252764258305,974622695435517953,1716163788,0,0),(977016627252121601,974622695435517953,1715994916,0,0),(977016627265220609,974622695435517953,1715998516,0,0),(977016627272937473,974622695435517953,1716002116,0,0),(977016627280166913,974622695435517953,1716005716,0,0),(977016627282485249,974622695435517953,1716009316,0,0),(977016627289944065,974622695435517953,1716012916,0,0),(977016627298037761,974622695435517953,1716016516,0,0),(977016627298324481,974622695435517953,1716020116,0,0),(977016627303763969,974622695435517953,1716023716,0,0),(977016627309748225,974622695435517953,1716027316,0,0),(977016627313836033,974622695435517953,1716030916,0,0),(977016627317211137,974622695435517953,1716034516,0,0),(977016627321106433,974622695435517953,1716038116,0,0),(977016627324968961,974622695435517953,1716041716,0,0),(977016627330715649,974622695435517953,1716045316,0,0),(977016627332857857,974622695435517953,1716048916,0,0),(977016627335872513,974622695435517953,1716056116,0,0),(977016627339743233,974622695435517953,1716052516,0,0),(977016627341156353,974622695435517953,1716063316,0,0),(977016627342172161,974622695435517953,1716059716,0,0),(977016627346325505,974622695435517953,1716066916,0,0),(977016627349938177,974622695435517953,1716070516,0,0),(977016627353968641,974622695435517953,1716074116,0,0),(977016627356651521,974622695435517953,1716077716,0,0),(977016627356954625,974622695435517953,1716084916,0,0),(977016627360641025,974622695435517953,1716081316,0,0),(977016627361218561,974622695435517953,1716088516,0,0),(977016627368378369,974622695435517953,1716092116,0,0),(977016627368521729,974622695435517953,1716095716,0,0),(977016627376279553,974622695435517953,1716099316,0,0),(977016627380584449,974622695435517953,1716102916,0,0),(977016627383488513,974622695435517953,1716106516,0,0),(977016627389812737,974622695435517953,1716110116,0,0),(977016627392962561,974622695435517953,1716113716,0,0),(977016627393245185,974622695435517953,1716117316,0,0),(977016627398103041,974622695435517953,1716120916,0,0),(977016627401220097,974622695435517953,1716124516,0,0),(977016627405770753,974622695435517953,1716128116,0,0),(977016627406512129,974622695435517953,1716131716,0,0),(977016627407458305,974622695435517953,1716135316,0,0),(977016627415031809,974622695435517953,1716138916,0,0),(977016627417931777,974622695435517953,1716146116,0,0),(977016627419635713,974622695435517953,1716142516,0,0),(977016627421425665,974622695435517953,1716149716,0,0),(977016627421757441,974622695435517953,1716153316,0,0),(977016627426365441,974622695435517953,1716156916,0,0),(977016627428810753,974622695435517953,1716160516,0,0),(977016627432140801,974622695435517953,1716164116,0,0),(977019156966068225,974622695435517953,1715995519,0,0),(977019156983283713,974622695435517953,1715999119,0,0),(977019156989435905,974622695435517953,1716002719,0,0),(977019156995637249,974622695435517953,1716006319,0,0),(977019157003386881,974622695435517953,1716009919,0,0),(977019157008560129,974622695435517953,1716013519,0,0),(977019157011701761,974622695435517953,1716017119,0,0),(977019157020401665,974622695435517953,1716020719,0,0),(977019157027008513,974622695435517953,1716024319,0,0),(977019157030457345,974622695435517953,1716027919,0,0),(977019157034020865,974622695435517953,1716031519,0,0),(977019157038981121,974622695435517953,1716035119,0,0),(977019157042745345,974622695435517953,1716038719,0,0),(977019157046865921,974622695435517953,1716042319,0,0),(977019157052456961,974622695435517953,1716045919,0,0),(977019157057970177,974622695435517953,1716049519,0,0),(977019157062537217,974622695435517953,1716053119,0,0),(977019157064400897,974622695435517953,1716056719,0,0),(977019157069524993,974622695435517953,1716060319,0,0),(977019157072420865,974622695435517953,1716063919,0,0),(977019157079035905,974622695435517953,1716067519,0,0),(977019157082021889,974622695435517953,1716071119,0,0),(977019157085683713,974622695435517953,1716074719,0,0),(977019157089972225,974622695435517953,1716078319,0,0),(977019157093920769,974622695435517953,1716081919,0,0),(977019157094232065,974622695435517953,1716085519,0,0),(977019157098622977,974622695435517953,1716089119,0,0),(977019157104234497,974622695435517953,1716096319,0,0),(977019157104365569,974622695435517953,1716092719,0,0),(977019157107355649,974622695435517953,1716099919,0,0),(977019157112893441,974622695435517953,1716103519,0,0),(977019157115457537,974622695435517953,1716107119,0,0),(977019157118025729,974622695435517953,1716110719,0,0),(977019157124034561,974622695435517953,1716114319,0,0),(977019157129416705,974622695435517953,1716117919,0,0),(977019157130600449,974622695435517953,1716125119,0,0),(977019157132079105,974622695435517953,1716121519,0,0),(977019157137170433,974622695435517953,1716128719,0,0),(977019157142028289,974622695435517953,1716132319,0,0),(977019157143818241,974622695435517953,1716135919,0,0),(977019157148520449,974622695435517953,1716139519,0,0),(977019157153939457,974622695435517953,1716146719,0,0),(977019157154430977,974622695435517953,1716143119,0,0),(977019157156212737,974622695435517953,1716150319,0,0),(977019157161766913,974622695435517953,1716153919,0,0),(977019157165232129,974622695435517953,1716157519,0,0),(977019157165387777,974622695435517953,1716161119,0,0),(977019157172088833,974622695435517953,1716164719,0,0),(977029847912263681,974622695435517953,1715998068,0,0),(977029847931326465,974622695435517953,1716001668,0,0),(977029847937097729,974622695435517953,1716005268,0,0),(977029847937265665,974622695435517953,1716008868,0,0),(977029847947988993,974622695435517953,1716012468,0,0),(977029847957864449,974622695435517953,1716016068,0,0),(977029847961567233,974622695435517953,1716019668,0,0),(977029847964393473,974622695435517953,1716023268,0,0),(977029847971168257,974622695435517953,1716026868,0,0),(977029847976640513,974622695435517953,1716030468,0,0),(977029847984959489,974622695435517953,1716034068,0,0),(977029847989628929,974622695435517953,1716037668,0,0),(977029847992532993,974622695435517953,1716041268,0,0),(977029848000770049,974622695435517953,1716044868,0,0),(977029848005951489,974622695435517953,1716048468,0,0),(977029848011829249,974622695435517953,1716052068,0,0),(977029848018706433,974622695435517953,1716055668,0,0),(977029848021258241,974622695435517953,1716059268,0,0),(977029848025628673,974622695435517953,1716062868,0,0),(977029848030867457,974622695435517953,1716066468,0,0),(977029848035471361,974622695435517953,1716070068,0,0),(977029848040132609,974622695435517953,1716073668,0,0),(977029848044953601,974622695435517953,1716077268,0,0),(977029848047333377,974622695435517953,1716080868,0,0),(977029848052981761,974622695435517953,1716084468,0,0),(977029848055353345,974622695435517953,1716088068,0,0),(977029848059768833,974622695435517953,1716095268,0,0),(977029848061460481,974622695435517953,1716091668,0,0),(977029848063569921,974622695435517953,1716098868,0,0),(977029848067850241,974622695435517953,1716102468,0,0),(977029848071979009,974622695435517953,1716106068,0,0),(977029848072204289,974622695435517953,1716109668,0,0),(977029848075739137,974622695435517953,1716113268,0,0),(977029848080302081,974622695435517953,1716120468,0,0),(977029848083017729,974622695435517953,1716116868,0,0),(977029848088608769,974622695435517953,1716124068,0,0),(977029848092721153,974622695435517953,1716127668,0,0),(977029848096399361,974622695435517953,1716131268,0,0),(977029848097136641,974622695435517953,1716134868,0,0),(977029848101928961,974622695435517953,1716138468,0,0),(977029848102379521,974622695435517953,1716142068,0,0),(977029848105111553,974622695435517953,1716145668,0,0),(977029848112381953,974622695435517953,1716149268,0,0),(977029848113410049,974622695435517953,1716152868,0,0),(977029848120926209,974622695435517953,1716156468,0,0),(977029848123682817,974622695435517953,1716160068,0,0),(977029848129761281,974622695435517953,1716163668,0,0),(977029848130908161,974622695435517953,1716167268,0,0),(977044949619752961,974622695435517953,1716001668,0,0),(977044949628551169,974622695435517953,1716005268,0,0),(977044949633302529,974622695435517953,1716008868,0,0),(977044949635387393,974622695435517953,1716012468,0,0),(977044949642559489,974622695435517953,1716016068,0,0),(977044949649350657,974622695435517953,1716019668,0,0),(977044949652430849,974622695435517953,1716023268,0,0),(977044949657042945,974622695435517953,1716030468,0,0),(977044949658533889,974622695435517953,1716026868,0,0),(977044949660655617,974622695435517953,1716034068,0,0),(977044949664583681,974622695435517953,1716037668,0,0),(977044949668360193,974622695435517953,1716041268,0,0),(977044949675388929,974622695435517953,1716044868,0,0),(977044949676302337,974622695435517953,1716048468,0,0),(977044949679853569,974622695435517953,1716052068,0,0),(977044949680852993,974622695435517953,1716055668,0,0),(977044949687791617,974622695435517953,1716062868,0,0),(977044949687803905,974622695435517953,1716059268,0,0),(977044949690552321,974622695435517953,1716066468,0,0),(977044949692452865,974622695435517953,1716073668,0,0),(977044949695311873,974622695435517953,1716070068,0,0),(977044949698682881,974622695435517953,1716077268,0,0),(977044949701246977,974622695435517953,1716080868,0,0),(977044949701881857,974622695435517953,1716084468,0,0),(977044949705629697,974622695435517953,1716088068,0,0),(977044949711036417,974622695435517953,1716091668,0,0),(977044949713653761,974622695435517953,1716095268,0,0),(977044949720694785,974622695435517953,1716098868,0,0),(977044949722230785,974622695435517953,1716102468,0,0),(977044949726150657,974622695435517953,1716106068,0,0),(977044949730828289,974622695435517953,1716109668,0,0),(977044949732864001,974622695435517953,1716113268,0,0),(977044949734453249,974622695435517953,1716116868,0,0),(977044949740179457,974622695435517953,1716120468,0,0),(977044949744549889,974622695435517953,1716124068,0,0),(977044949750382593,974622695435517953,1716127668,0,0),(977044949755015169,974622695435517953,1716131268,0,0),(977044949756379137,974622695435517953,1716134868,0,0),(977044949764640769,974622695435517953,1716138468,0,0),(977044949770072065,974622695435517953,1716142068,0,0),(977044949773561857,974622695435517953,1716149268,0,0),(977044949773570049,974622695435517953,1716145668,0,0),(977044949778223105,974622695435517953,1716152868,0,0),(977044949782945793,974622695435517953,1716160068,0,0),(977044949783138305,974622695435517953,1716156468,0,0),(977044949786796033,974622695435517953,1716163668,0,0),(977044949790429185,974622695435517953,1716170868,0,0),(977044949791588353,974622695435517953,1716167268,0,0),(977091871285751809,974622695435517953,1716012855,0,0),(977091871295475713,974622695435517953,1716016455,0,0),(977091871299960833,974622695435517953,1716020055,0,0),(977091871309185025,974622695435517953,1716023655,0,0),(977091871312748545,974622695435517953,1716027255,0,0),(977091871324352513,974622695435517953,1716030855,0,0),(977091871327473665,974622695435517953,1716034455,0,0),(977091871333236737,974622695435517953,1716038055,0,0),(977091871335387137,974622695435517953,1716045255,0,0),(977091871336742913,974622695435517953,1716041655,0,0),(977091871338037249,974622695435517953,1716048855,0,0),(977091871342841857,974622695435517953,1716052455,0,0),(977091871344046081,974622695435517953,1716056055,0,0),(977091871349370881,974622695435517953,1716059655,0,0),(977091871351775233,974622695435517953,1716066855,0,0),(977091871352176641,974622695435517953,1716063255,0,0),(977091871354433537,974622695435517953,1716074055,0,0),(977091871355834369,974622695435517953,1716070455,0,0),(977091871359987713,974622695435517953,1716081255,0,0),(977091871360229377,974622695435517953,1716077655,0,0),(977091871363457025,974622695435517953,1716088455,0,0),(977091871365926913,974622695435517953,1716084855,0,0),(977091871368318977,974622695435517953,1716092055,0,0),(977091871369302017,974622695435517953,1716099255,0,0),(977091871369547777,974622695435517953,1716095655,0,0),(977091871371620353,974622695435517953,1716102855,0,0),(977091871376920577,974622695435517953,1716106455,0,0),(977091871377137665,974622695435517953,1716110055,0,0),(977091871381426177,974622695435517953,1716113655,0,0),(977091871383724033,974622695435517953,1716117255,0,0),(977091871383863297,974622695435517953,1716120855,0,0),(977091871385649153,974622695435517953,1716128055,0,0),(977091871387672577,974622695435517953,1716124455,0,0),(977091871388336129,974622695435517953,1716135255,0,0),(977091871392055297,974622695435517953,1716131655,0,0),(977091871394115585,974622695435517953,1716146055,0,0),(977091871394770945,974622695435517953,1716138855,0,0),(977091871396286465,974622695435517953,1716142455,0,0),(977091871396990977,974622695435517953,1716153255,0,0),(977091871399862273,974622695435517953,1716149655,0,0),(977091871405920257,974622695435517953,1716160455,0,0),(977091871406931969,974622695435517953,1716156855,0,0),(977091871410147329,974622695435517953,1716167655,0,0),(977091871410651137,974622695435517953,1716164055,0,0),(977091871413219329,974622695435517953,1716174855,0,0),(977091871416287233,974622695435517953,1716171255,0,0),(977091871419613185,974622695435517953,1716178455,0,0),(977091871420542977,974622695435517953,1716182055,0,0),(977093118037155841,974622695435517953,1716013152,0,0),(977093118044192769,974622695435517953,1716016752,0,0),(977093118055735297,974622695435517953,1716020352,0,0),(977093118058524673,974622695435517953,1716023952,0,0),(977093118060494849,974622695435517953,1716027552,0,0),(977093118063341569,974622695435517953,1716031152,0,0),(977093118067388417,974622695435517953,1716034752,0,0),(977093118071640065,974622695435517953,1716041952,0,0),(977093118073241601,974622695435517953,1716038352,0,0),(977093118077292545,974622695435517953,1716045552,0,0),(977093118078545921,974622695435517953,1716052752,0,0),(977093118078668801,974622695435517953,1716049152,0,0),(977093118081912833,974622695435517953,1716056352,0,0),(977093118089785345,974622695435517953,1716059952,0,0),(977093118090682369,974622695435517953,1716067152,0,0),(977093118092206081,974622695435517953,1716063552,0,0),(977093118096433153,974622695435517953,1716070752,0,0),(977093118099169281,974622695435517953,1716074352,0,0),(977093118099210241,974622695435517953,1716077952,0,0),(977093118104190977,974622695435517953,1716081552,0,0),(977093118106206209,974622695435517953,1716085152,0,0),(977093118109818881,974622695435517953,1716088752,0,0),(977093118114533377,974622695435517953,1716092352,0,0),(977093118115287041,974622695435517953,1716095952,0,0),(977093118115979265,974622695435517953,1716099552,0,0),(977093118120837121,974622695435517953,1716103152,0,0),(977093118121394177,974622695435517953,1716106752,0,0),(977093118124384257,974622695435517953,1716113952,0,0),(977093118124826625,974622695435517953,1716110352,0,0),(977093118130016257,974622695435517953,1716117552,0,0),(977093118133485569,974622695435517953,1716124752,0,0),(977093118134571009,974622695435517953,1716121152,0,0),(977093118139781121,974622695435517953,1716128352,0,0),(977093118140641281,974622695435517953,1716131952,0,0),(977093118147592193,974622695435517953,1716135552,0,0),(977093118151282689,974622695435517953,1716139152,0,0),(977093118153248769,974622695435517953,1716142752,0,0),(977093118160515073,974622695435517953,1716146352,0,0),(977093118164316161,974622695435517953,1716149952,0,0),(977093118167597057,974622695435517953,1716153552,0,0),(977093118170562561,974622695435517953,1716157152,0,0),(977093118175756289,974622695435517953,1716160752,0,0),(977093118180880385,974622695435517953,1716164352,0,0),(977093118184005633,974622695435517953,1716171552,0,0),(977093118184161281,974622695435517953,1716167952,0,0),(977093118189842433,974622695435517953,1716175152,0,0),(977093118191259649,974622695435517953,1716182352,0,0),(977093118191910913,974622695435517953,1716178752,0,0),(977094375365128193,974622695435517953,1716013452,0,0),(977094375374168065,974622695435517953,1716017052,0,0),(977094375380127745,974622695435517953,1716024252,0,0),(977094375380275201,974622695435517953,1716020652,0,0),(977094375385808897,974622695435517953,1716027852,0,0),(977094375385845761,974622695435517953,1716031452,0,0),(977094375389491201,974622695435517953,1716035052,0,0),(977094375391170561,974622695435517953,1716038652,0,0),(977094375393697793,974622695435517953,1716042252,0,0),(977094375394189313,974622695435517953,1716045852,0,0),(977094375396478977,974622695435517953,1716049452,0,0),(977094375398416385,974622695435517953,1716053052,0,0),(977094375401578497,974622695435517953,1716056652,0,0),(977094375403429889,974622695435517953,1716060252,0,0),(977094375406264321,974622695435517953,1716067452,0,0),(977094375408263169,974622695435517953,1716063852,0,0),(977094375409369089,974622695435517953,1716071052,0,0),(977094375410118657,974622695435517953,1716074652,0,0),(977094375413301249,974622695435517953,1716078252,0,0),(977094375417860097,974622695435517953,1716081852,0,0),(977094375421280257,974622695435517953,1716089052,0,0),(977094375422857217,974622695435517953,1716085452,0,0),(977094375427796993,974622695435517953,1716092652,0,0),(977094375429791745,974622695435517953,1716099852,0,0),(977094375430041601,974622695435517953,1716096252,0,0),(977094375437385729,974622695435517953,1716103452,0,0),(977094375440121857,974622695435517953,1716107052,0,0),(977094375441473537,974622695435517953,1716110652,0,0),(977094375442051073,974622695435517953,1716114252,0,0),(977094375442972673,974622695435517953,1716117852,0,0),(977094375447633921,974622695435517953,1716121452,0,0),(977094375448932353,974622695435517953,1716125052,0,0),(977094375452471297,974622695435517953,1716128652,0,0),(977094375455789057,974622695435517953,1716135852,0,0),(977094375455797249,974622695435517953,1716132252,0,0),(977094375461158913,974622695435517953,1716139452,0,0),(977094375463292929,974622695435517953,1716143052,0,0),(977094375467945985,974622695435517953,1716146652,0,0),(977094375472865281,974622695435517953,1716153852,0,0),(977094375474335745,974622695435517953,1716150252,0,0),(977094375477854209,974622695435517953,1716157452,0,0),(977094375483600897,974622695435517953,1716161052,0,0),(977094375485595649,974622695435517953,1716164652,0,0),(977094375485894657,974622695435517953,1716168252,0,0),(977094375490056193,974622695435517953,1716171852,0,0),(977094375494270977,974622695435517953,1716175452,0,0),(977094375498735617,974622695435517953,1716179052,0,0),(977094375502106625,974622695435517953,1716182652,0,0),(977095158711242753,974622695435517953,1716013639,0,0),(977095158716354561,974622695435517953,1716017239,0,0),(977095158725492737,974622695435517953,1716020839,0,0),(977095158730756097,974622695435517953,1716024439,0,0),(977095158734508033,974622695435517953,1716028039,0,0),(977095158735888385,974622695435517953,1716031639,0,0),(977095158740115457,974622695435517953,1716035239,0,0),(977095158744440833,974622695435517953,1716038839,0,0),(977095158749106177,974622695435517953,1716042439,0,0),(977095158749696001,974622695435517953,1716046039,0,0),(977095158754836481,974622695435517953,1716049639,0,0),(977095158758588417,974622695435517953,1716053239,0,0),(977095158762311681,974622695435517953,1716056839,0,0),(977095158766514177,974622695435517953,1716060439,0,0),(977095158774046721,974622695435517953,1716064039,0,0),(977095158775631873,974622695435517953,1716067639,0,0),(977095158779174913,974622695435517953,1716071239,0,0),(977095158783172609,974622695435517953,1716074839,0,0),(977095158789980161,974622695435517953,1716078439,0,0),(977095158791372801,974622695435517953,1716082039,0,0),(977095158796718081,974622695435517953,1716085639,0,0),(977095158801600513,974622695435517953,1716089239,0,0),(977095158804656129,974622695435517953,1716092839,0,0),(977095158809464833,974622695435517953,1716096439,0,0),(977095158814097409,974622695435517953,1716100039,0,0),(977095158817878017,974622695435517953,1716103639,0,0),(977095158822195201,974622695435517953,1716107239,0,0),(977095158824837121,974622695435517953,1716110839,0,0),(977095158830092289,974622695435517953,1716114439,0,0),(977095158837129217,974622695435517953,1716118039,0,0),(977095158838112257,974622695435517953,1716121639,0,0),(977095158843219969,974622695435517953,1716125239,0,0),(977095158847766529,974622695435517953,1716128839,0,0),(977095158851928065,974622695435517953,1716132439,0,0),(977095158852300801,974622695435517953,1716136039,0,0),(977095158857838593,974622695435517953,1716139639,0,0),(977095158861389825,974622695435517953,1716143239,0,0),(977095158864707585,974622695435517953,1716146839,0,0),(977095158868774913,974622695435517953,1716150439,0,0),(977095158873858049,974622695435517953,1716154039,0,0),(977095158877179905,974622695435517953,1716157639,0,0),(977095158881996801,974622695435517953,1716161239,0,0),(977095158883418113,974622695435517953,1716164839,0,0),(977095158890311681,974622695435517953,1716168439,0,0),(977095158897999873,974622695435517953,1716172039,0,0),(977095158902370305,974622695435517953,1716179239,0,0),(977095158903058433,974622695435517953,1716175639,0,0),(977095158906839041,974622695435517953,1716182839,0,0),(977105545311563777,974622695435517953,1716016115,0,0),(977105545327153153,974622695435517953,1716019715,0,0),(977105545333010433,974622695435517953,1716023315,0,0),(977105545342181377,974622695435517953,1716026915,0,0),(977105545346981889,974622695435517953,1716030515,0,0),(977105545350074369,974622695435517953,1716034115,0,0),(977105545352237057,974622695435517953,1716037715,0,0),(977105545359581185,974622695435517953,1716041315,0,0),(977105545361453057,974622695435517953,1716044915,0,0),(977105545367887873,974622695435517953,1716048515,0,0),(977105545368686593,974622695435517953,1716052115,0,0),(977105545372774401,974622695435517953,1716059315,0,0),(977105545373110273,974622695435517953,1716055715,0,0),(977105545378193409,974622695435517953,1716062915,0,0),(977105545382313985,974622695435517953,1716066515,0,0),(977105545385414657,974622695435517953,1716070115,0,0),(977105545387446273,974622695435517953,1716073715,0,0),(977105545393557505,974622695435517953,1716077315,0,0),(977105545394876417,974622695435517953,1716084515,0,0),(977105545397022721,974622695435517953,1716080915,0,0),(977105545398763521,974622695435517953,1716091715,0,0),(977105545399255041,974622695435517953,1716088115,0,0),(977105545406836737,974622695435517953,1716095315,0,0),(977105545407672321,974622695435517953,1716098915,0,0),(977105545411198977,974622695435517953,1716106115,0,0),(977105545412509697,974622695435517953,1716102515,0,0),(977105545414594561,974622695435517953,1716109715,0,0),(977105545417117697,974622695435517953,1716113315,0,0),(977105545420877825,974622695435517953,1716120515,0,0),(977105545421332481,974622695435517953,1716116915,0,0),(977105545426612225,974622695435517953,1716127715,0,0),(977105545426644993,974622695435517953,1716124115,0,0),(977105545428307969,974622695435517953,1716134915,0,0),(977105545429430273,974622695435517953,1716131315,0,0),(977105545434480641,974622695435517953,1716138515,0,0),(977105545439170561,974622695435517953,1716142115,0,0),(977105545439965185,974622695435517953,1716145715,0,0),(977105545444761601,974622695435517953,1716152915,0,0),(977105545445879809,974622695435517953,1716149315,0,0),(977105545448243201,974622695435517953,1716160115,0,0),(977105545450590209,974622695435517953,1716156515,0,0),(977105545452834817,974622695435517953,1716163715,0,0),(977105545456693249,974622695435517953,1716167315,0,0),(977105545460019201,974622695435517953,1716170915,0,0),(977105545461641217,974622695435517953,1716174515,0,0),(977105545465417729,974622695435517953,1716178115,0,0),(977105545470525441,974622695435517953,1716185315,0,0),(977105545470545921,974622695435517953,1716181715,0,0),(977106895111663617,974622695435517953,1716016437,0,0),(977106895131156481,974622695435517953,1716020037,0,0),(977106895137955841,974622695435517953,1716023637,0,0),(977106895149572097,974622695435517953,1716027237,0,0),(977106895153008641,974622695435517953,1716030837,0,0),(977106895157379073,974622695435517953,1716034437,0,0),(977106895158489089,974622695435517953,1716038037,0,0),(977106895166152705,974622695435517953,1716041637,0,0),(977106895167164417,974622695435517953,1716048837,0,0),(977106895168368641,974622695435517953,1716045237,0,0),(977106895179075585,974622695435517953,1716052437,0,0),(977106895181070337,974622695435517953,1716056037,0,0),(977106895183622145,974622695435517953,1716063237,0,0),(977106895184592897,974622695435517953,1716059637,0,0),(977106895188869121,974622695435517953,1716066837,0,0),(977106895194374145,974622695435517953,1716070437,0,0),(977106895195983873,974622695435517953,1716074037,0,0),(977106895197294593,974622695435517953,1716077637,0,0),(977106895199498241,974622695435517953,1716081237,0,0),(977106895201697793,974622695435517953,1716084837,0,0),(977106895202361345,974622695435517953,1716088437,0,0),(977106895205720065,974622695435517953,1716092037,0,0),(977106895206576129,974622695435517953,1716095637,0,0),(977106895211360257,974622695435517953,1716099237,0,0),(977106895213117441,974622695435517953,1716106437,0,0),(977106895213539329,974622695435517953,1716102837,0,0),(977106895217111041,974622695435517953,1716110037,0,0),(977106895221215233,974622695435517953,1716113637,0,0),(977106895222538241,974622695435517953,1716117237,0,0),(977106895224082433,974622695435517953,1716120837,0,0),(977106895225892865,974622695435517953,1716124437,0,0),(977106895233699841,974622695435517953,1716128037,0,0),(977106895236751361,974622695435517953,1716131637,0,0),(977106895238950913,974622695435517953,1716135237,0,0),(977106895243407361,974622695435517953,1716142437,0,0),(977106895246204929,974622695435517953,1716138837,0,0),(977106895246458881,974622695435517953,1716146037,0,0),(977106895251877889,974622695435517953,1716149637,0,0),(977106895251972097,974622695435517953,1716153237,0,0),(977106895255539713,974622695435517953,1716156837,0,0),(977106895258165249,974622695435517953,1716160437,0,0),(977106895261253633,974622695435517953,1716164037,0,0),(977106895264206849,974622695435517953,1716171237,0,0),(977106895264681985,974622695435517953,1716167637,0,0),(977106895267938305,974622695435517953,1716178437,0,0),(977106895268827137,974622695435517953,1716174837,0,0),(977106895274147841,974622695435517953,1716182037,0,0),(977106895276290049,974622695435517953,1716185637,0,0),(977108188163059713,974622695435517953,1716016745,0,0),(977108188176715777,974622695435517953,1716020345,0,0),(977108188185403393,974622695435517953,1716023945,0,0),(977108188188979201,974622695435517953,1716027545,0,0),(977108188194783233,974622695435517953,1716031145,0,0),(977108188196605953,974622695435517953,1716034745,0,0),(977108188199501825,974622695435517953,1716038345,0,0),(977108188206149633,974622695435517953,1716041945,0,0),(977108188211589121,974622695435517953,1716045545,0,0),(977108188213477377,974622695435517953,1716049145,0,0),(977108188216602625,974622695435517953,1716052745,0,0),(977108188220719105,974622695435517953,1716056345,0,0),(977108188224724993,974622695435517953,1716059945,0,0),(977108188234645505,974622695435517953,1716063545,0,0),(977108188240883713,974622695435517953,1716067145,0,0),(977108188242821121,974622695435517953,1716070745,0,0),(977108188247425025,974622695435517953,1716074345,0,0),(977108188255604737,974622695435517953,1716077945,0,0),(977108188261556225,974622695435517953,1716081545,0,0),(977108188265062401,974622695435517953,1716085145,0,0),(977108188270559233,974622695435517953,1716088745,0,0),(977108188273242113,974622695435517953,1716092345,0,0),(977108188278771713,974622695435517953,1716095945,0,0),(977108188282933249,974622695435517953,1716099545,0,0),(977108188286148609,974622695435517953,1716103145,0,0),(977108188290248705,974622695435517953,1716106745,0,0),(977108188298584065,974622695435517953,1716110345,0,0),(977108188301344769,974622695435517953,1716113945,0,0),(977108188306030593,974622695435517953,1716117545,0,0),(977108188309618689,974622695435517953,1716121145,0,0),(977108188316225537,974622695435517953,1716124745,0,0),(977108188320063489,974622695435517953,1716128345,0,0),(977108188323299329,974622695435517953,1716131945,0,0),(977108188327772161,974622695435517953,1716135545,0,0),(977108188331036673,974622695435517953,1716142745,0,0),(977108188333539329,974622695435517953,1716139145,0,0),(977108188335702017,974622695435517953,1716149945,0,0),(977108188336381953,974622695435517953,1716146345,0,0),(977108188340772865,974622695435517953,1716153545,0,0),(977108188343013377,974622695435517953,1716160745,0,0),(977108188345503745,974622695435517953,1716157145,0,0),(977108188347035649,974622695435517953,1716167945,0,0),(977108188347326465,974622695435517953,1716164345,0,0),(977108188352581633,974622695435517953,1716171545,0,0),(977108188357120001,974622695435517953,1716175145,0,0),(977108188359462913,974622695435517953,1716178745,0,0),(977108188360630273,974622695435517953,1716182345,0,0),(977108188365410305,974622695435517953,1716185945,0,0),(977115951956160513,974622695435517953,1716018596,0,0),(977115951960965121,974622695435517953,1716022196,0,0),(977115951967281153,974622695435517953,1716025796,0,0),(977115951972511745,974622695435517953,1716029396,0,0),(977115951975747585,974622695435517953,1716032996,0,0),(977115951979737089,974622695435517953,1716036596,0,0),(977115951988670465,974622695435517953,1716040196,0,0),(977115951990489089,974622695435517953,1716043796,0,0),(977115951995703297,974622695435517953,1716047396,0,0),(977115952002154497,974622695435517953,1716050996,0,0),(977115952006033409,974622695435517953,1716054596,0,0),(977115952007262209,974622695435517953,1716061796,0,0),(977115952009142273,974622695435517953,1716058196,0,0),(977115952015990785,974622695435517953,1716065396,0,0),(977115952020729857,974622695435517953,1716068996,0,0),(977115952026677249,974622695435517953,1716072596,0,0),(977115952030433281,974622695435517953,1716076196,0,0),(977115952032182273,974622695435517953,1716079796,0,0),(977115952036421633,974622695435517953,1716083396,0,0),(977115952038502401,974622695435517953,1716090596,0,0),(977115952038522881,974622695435517953,1716086996,0,0),(977115952040681473,974622695435517953,1716097796,0,0),(977115952043819009,974622695435517953,1716094196,0,0),(977115952044843009,974622695435517953,1716101396,0,0),(977115952049250305,974622695435517953,1716108596,0,0),(977115952052027393,974622695435517953,1716104996,0,0),(977115952054829057,974622695435517953,1716112196,0,0),(977115952055349249,974622695435517953,1716115796,0,0),(977115952058720257,974622695435517953,1716119396,0,0),(977115952059248641,974622695435517953,1716122996,0,0),(977115952063545345,974622695435517953,1716126596,0,0),(977115952066670593,974622695435517953,1716130196,0,0),(977115952069267457,974622695435517953,1716133796,0,0),(977115952071991297,974622695435517953,1716144596,0,0),(977115952072916993,974622695435517953,1716140996,0,0),(977115952073261057,974622695435517953,1716137396,0,0),(977115952076128257,974622695435517953,1716151796,0,0),(977115952077737985,974622695435517953,1716148196,0,0),(977115952079474689,974622695435517953,1716155396,0,0),(977115952082407425,974622695435517953,1716158996,0,0),(977115952083517441,974622695435517953,1716162596,0,0),(977115952087814145,974622695435517953,1716166196,0,0),(977115952090574849,974622695435517953,1716169796,0,0),(977115952093851649,974622695435517953,1716176996,0,0),(977115952094019585,974622695435517953,1716173396,0,0),(977115952098852865,974622695435517953,1716184196,0,0),(977115952099016705,974622695435517953,1716180596,0,0),(977115952100622337,974622695435517953,1716187796,0,0),(977122230836490241,974622695435517953,1716020093,0,0),(977122230854115329,974622695435517953,1716023693,0,0),(977122230855315457,974622695435517953,1716027293,0,0),(977122230859886593,974622695435517953,1716030893,0,0),(977122230867828737,974622695435517953,1716034493,0,0),(977122230879510529,974622695435517953,1716038093,0,0),(977122230884220929,974622695435517953,1716041693,0,0),(977122230885478401,974622695435517953,1716045293,0,0),(977122230891503617,974622695435517953,1716048893,0,0),(977122230892814337,974622695435517953,1716056093,0,0),(977122230896422913,974622695435517953,1716052493,0,0),(977122230898626561,974622695435517953,1716059693,0,0),(977122230901293057,974622695435517953,1716066893,0,0),(977122230904160257,974622695435517953,1716063293,0,0),(977122230905978881,974622695435517953,1716070493,0,0),(977122230906380289,974622695435517953,1716074093,0,0),(977122230909698049,974622695435517953,1716077693,0,0),(977122230917599233,974622695435517953,1716081293,0,0),(977122230920007681,974622695435517953,1716088493,0,0),(977122230921543681,974622695435517953,1716084893,0,0),(977122230922469377,974622695435517953,1716092093,0,0),(977122230922805249,974622695435517953,1716095693,0,0),(977122230927761409,974622695435517953,1716102893,0,0),(977122230929780737,974622695435517953,1716099293,0,0),(977122230930857985,974622695435517953,1716106493,0,0),(977122230931406849,974622695435517953,1716110093,0,0),(977122230936322049,974622695435517953,1716113693,0,0),(977122230936481793,974622695435517953,1716117293,0,0),(977122230941581313,974622695435517953,1716120893,0,0),(977122230942253057,974622695435517953,1716124493,0,0),(977122230943174657,974622695435517953,1716128093,0,0),(977122230945439745,974622695435517953,1716131693,0,0),(977122230949068801,974622695435517953,1716138893,0,0),(977122230951395329,974622695435517953,1716135293,0,0),(977122230951546881,974622695435517953,1716146093,0,0),(977122230954422273,974622695435517953,1716142493,0,0),(977122230956539905,974622695435517953,1716149693,0,0),(977122230958940161,974622695435517953,1716153293,0,0),(977122230963838977,974622695435517953,1716156893,0,0),(977122230967549953,974622695435517953,1716160493,0,0),(977122230970429441,974622695435517953,1716164093,0,0),(977122230976176129,974622695435517953,1716167693,0,0),(977122230976434177,974622695435517953,1716171293,0,0),(977122230980018177,974622695435517953,1716174893,0,0),(977122230980231169,974622695435517953,1716178493,0,0),(977122230982250497,974622695435517953,1716182093,0,0),(977122230985261057,974622695435517953,1716185693,0,0),(977122230987833345,974622695435517953,1716189293,0,0),(977123509445423105,974622695435517953,1716020398,0,0),(977123509462171649,974622695435517953,1716023998,0,0),(977123509465296897,974622695435517953,1716027598,0,0),(977123509467856897,974622695435517953,1716031198,0,0),(977123509471830017,974622695435517953,1716034798,0,0),(977123509479604225,974622695435517953,1716038398,0,0),(977123509483061249,974622695435517953,1716041998,0,0),(977123509485830145,974622695435517953,1716045598,0,0),(977123509493956609,974622695435517953,1716049198,0,0),(977123509499609089,974622695435517953,1716052798,0,0),(977123509504425985,974622695435517953,1716056398,0,0),(977123509509386241,974622695435517953,1716059998,0,0),(977123509512716289,974622695435517953,1716063598,0,0),(977123509520371713,974622695435517953,1716067198,0,0),(977123509529866241,974622695435517953,1716070798,0,0),(977123509533986817,974622695435517953,1716077998,0,0),(977123509534076929,974622695435517953,1716074398,0,0),(977123509538557953,974622695435517953,1716081598,0,0),(977123509540298753,974622695435517953,1716085198,0,0),(977123509546844161,974622695435517953,1716088798,0,0),(977123509548490753,974622695435517953,1716092398,0,0),(977123509558460417,974622695435517953,1716095998,0,0),(977123509562564609,974622695435517953,1716099598,0,0),(977123509564129281,974622695435517953,1716103198,0,0),(977123509567954945,974622695435517953,1716106798,0,0),(977123509571411969,974622695435517953,1716110398,0,0),(977123509573095425,974622695435517953,1716117598,0,0),(977123509573263361,974622695435517953,1716113998,0,0),(977123509579059201,974622695435517953,1716121198,0,0),(977123509581611009,974622695435517953,1716128398,0,0),(977123509582770177,974622695435517953,1716124798,0,0),(977123509585563649,974622695435517953,1716131998,0,0),(977123509590851585,974622695435517953,1716135598,0,0),(977123509592829953,974622695435517953,1716139198,0,0),(977123509595459585,974622695435517953,1716142798,0,0),(977123509600301057,974622695435517953,1716146398,0,0),(977123509603741697,974622695435517953,1716153598,0,0),(977123509605617665,974622695435517953,1716149998,0,0),(977123509606780929,974622695435517953,1716157198,0,0),(977123509612425217,974622695435517953,1716160798,0,0),(977123509616291841,974622695435517953,1716167998,0,0),(977123509616930817,974622695435517953,1716164398,0,0),(977123509620969473,974622695435517953,1716171598,0,0),(977123509623156737,974622695435517953,1716175198,0,0),(977123509627092993,974622695435517953,1716178798,0,0),(977123509632266241,974622695435517953,1716185998,0,0),(977123509632389121,974622695435517953,1716182398,0,0),(977123509637255169,974622695435517953,1716189598,0,0),(977124526249021441,974622695435517953,1716020641,0,0),(977124526262013953,974622695435517953,1716024241,0,0),(977124526268190721,974622695435517953,1716027841,0,0),(977124526273355777,974622695435517953,1716031441,0,0),(977124526278361089,974622695435517953,1716035041,0,0),(977124526281302017,974622695435517953,1716038641,0,0),(977124526286553089,974622695435517953,1716042241,0,0),(977124526289612801,974622695435517953,1716045841,0,0),(977124526293159937,974622695435517953,1716049441,0,0),(977124526298337281,974622695435517953,1716053041,0,0),(977124526301298689,974622695435517953,1716056641,0,0),(977124526309326849,974622695435517953,1716060241,0,0),(977124526310522881,974622695435517953,1716063841,0,0),(977124526318899201,974622695435517953,1716067441,0,0),(977124526323425281,974622695435517953,1716071041,0,0),(977124526325473281,974622695435517953,1716074641,0,0),(977124526329905153,974622695435517953,1716078241,0,0),(977124526333739009,974622695435517953,1716081841,0,0),(977124526338596865,974622695435517953,1716085441,0,0),(977124526342832129,974622695435517953,1716092641,0,0),(977124526342950913,974622695435517953,1716089041,0,0),(977124526344511489,974622695435517953,1716096241,0,0),(977124526347931649,974622695435517953,1716099841,0,0),(977124526350376961,974622695435517953,1716103441,0,0),(977124526354857985,974622695435517953,1716110641,0,0),(977124526354903041,974622695435517953,1716107041,0,0),(977124526359441409,974622695435517953,1716114241,0,0),(977124526360387585,974622695435517953,1716117841,0,0),(977124526366445569,974622695435517953,1716121441,0,0),(977124526367453185,974622695435517953,1716125041,0,0),(977124526370074625,974622695435517953,1716128641,0,0),(977124526372990977,974622695435517953,1716132241,0,0),(977124526378299393,974622695435517953,1716135841,0,0),(977124526380249089,974622695435517953,1716139441,0,0),(977124526382280705,974622695435517953,1716143041,0,0),(977124526383734785,974622695435517953,1716146641,0,0),(977124526388043777,974622695435517953,1716150241,0,0),(977124526388678657,974622695435517953,1716153841,0,0),(977124526389288961,974622695435517953,1716157441,0,0),(977124526391103489,974622695435517953,1716161041,0,0),(977124526394429441,974622695435517953,1716164641,0,0),(977124526401327105,974622695435517953,1716168241,0,0),(977124526403682305,974622695435517953,1716175441,0,0),(977124526404820993,974622695435517953,1716171841,0,0),(977124526408114177,974622695435517953,1716179041,0,0),(977124526412005377,974622695435517953,1716186241,0,0),(977124526412058625,974622695435517953,1716182641,0,0),(977124526417989633,974622695435517953,1716189841,0,0),(977129156503085057,974622695435517953,1716021745,0,0),(977129156515250177,974622695435517953,1716025345,0,0),(977129156523495425,974622695435517953,1716028945,0,0),(977129156531171329,974622695435517953,1716032545,0,0),(977129156534902785,974622695435517953,1716036145,0,0),(977129156541460481,974622695435517953,1716039745,0,0),(977129156548370433,974622695435517953,1716043345,0,0),(977129156550688769,974622695435517953,1716046945,0,0),(977129156552323073,974622695435517953,1716050545,0,0),(977129156554035201,974622695435517953,1716054145,0,0),(977129156559458305,974622695435517953,1716057745,0,0),(977129156564152321,974622695435517953,1716061345,0,0),(977129156566261761,974622695435517953,1716064945,0,0),(977129156572307457,974622695435517953,1716068545,0,0),(977129156576051201,974622695435517953,1716072145,0,0),(977129156580315137,974622695435517953,1716079345,0,0),(977129156581752833,974622695435517953,1716075745,0,0),(977129156582580225,974622695435517953,1716082945,0,0),(977129156588367873,974622695435517953,1716086545,0,0),(977129156589305857,974622695435517953,1716090145,0,0),(977129156594204673,974622695435517953,1716097345,0,0),(977129156594343937,974622695435517953,1716093745,0,0),(977129156597702657,974622695435517953,1716100945,0,0),(977129156598042625,974622695435517953,1716104545,0,0),(977129156600000513,974622695435517953,1716108145,0,0),(977129156601577473,974622695435517953,1716111745,0,0),(977129156603273217,974622695435517953,1716115345,0,0),(977129156603478017,974622695435517953,1716118945,0,0),(977129156608405505,974622695435517953,1716126145,0,0),(977129156608880641,974622695435517953,1716122545,0,0),(977129156612206593,974622695435517953,1716133345,0,0),(977129156612890625,974622695435517953,1716129745,0,0),(977129156616097793,974622695435517953,1716136945,0,0),(977129156617682945,974622695435517953,1716144145,0,0),(977129156618014721,974622695435517953,1716140545,0,0),(977129156620451841,974622695435517953,1716151345,0,0),(977129156622577665,974622695435517953,1716147745,0,0),(977129156624609281,974622695435517953,1716154945,0,0),(977129156627427329,974622695435517953,1716158545,0,0),(977129156630581249,974622695435517953,1716162145,0,0),(977129156633681921,974622695435517953,1716165745,0,0),(977129156640395265,974622695435517953,1716169345,0,0),(977129156641902593,974622695435517953,1716172945,0,0),(977129156643999745,974622695435517953,1716176545,0,0),(977129156647399425,974622695435517953,1716180145,0,0),(977129156650831873,974622695435517953,1716187345,0,0),(977129156652941313,974622695435517953,1716183745,0,0),(977129156655325185,974622695435517953,1716190945,0,0),(977130697015226369,974622695435517953,1716022112,0,0),(977130697031864321,974622695435517953,1716025712,0,0),(977130697036423169,974622695435517953,1716029312,0,0),(977130697036824577,974622695435517953,1716036512,0,0),(977130697039323137,974622695435517953,1716032912,0,0),(977130697043197953,974622695435517953,1716043712,0,0),(977130697044361217,974622695435517953,1716040112,0,0),(977130697048338433,974622695435517953,1716047312,0,0),(977130697048547329,974622695435517953,1716050912,0,0),(977130697053057025,974622695435517953,1716054512,0,0),(977130697053143041,974622695435517953,1716058112,0,0),(977130697056727041,974622695435517953,1716061712,0,0),(977130697057886209,974622695435517953,1716068912,0,0),(977130697060651009,974622695435517953,1716065312,0,0),(977130697063997441,974622695435517953,1716072512,0,0),(977130697068539905,974622695435517953,1716076112,0,0),(977130697073893377,974622695435517953,1716079712,0,0),(977130697077420033,974622695435517953,1716083312,0,0),(977130697082523649,974622695435517953,1716086912,0,0),(977130697084133377,974622695435517953,1716090512,0,0),(977130697090011137,974622695435517953,1716094112,0,0),(977130697092939777,974622695435517953,1716097712,0,0),(977130697098248193,974622695435517953,1716101312,0,0),(977130697101258753,974622695435517953,1716104912,0,0),(977130697103355905,974622695435517953,1716108512,0,0),(977130697104396289,974622695435517953,1716112112,0,0),(977130697106427905,974622695435517953,1716115712,0,0),(977130697111764993,974622695435517953,1716119312,0,0),(977130697111912449,974622695435517953,1716122912,0,0),(977130697113849857,974622695435517953,1716126512,0,0),(977130697118519297,974622695435517953,1716133712,0,0),(977130697118896129,974622695435517953,1716130112,0,0),(977130697120231425,974622695435517953,1716137312,0,0),(977130697120690177,974622695435517953,1716140912,0,0),(977130697121890305,974622695435517953,1716144512,0,0),(977130697127989249,974622695435517953,1716148112,0,0),(977130697128554497,974622695435517953,1716151712,0,0),(977130697129021441,974622695435517953,1716158912,0,0),(977130697133125633,974622695435517953,1716155312,0,0),(977130697135415297,974622695435517953,1716162512,0,0),(977130697138442241,974622695435517953,1716166112,0,0),(977130697143779329,974622695435517953,1716169712,0,0),(977130697148088321,974622695435517953,1716173312,0,0),(977130697149161473,974622695435517953,1716176912,0,0),(977130697150750721,974622695435517953,1716180512,0,0),(977130697154338817,974622695435517953,1716184112,0,0),(977130697154412545,974622695435517953,1716187712,0,0),(977130697161170945,974622695435517953,1716191312,0,0),(977145798063235073,974622695435517953,1716025712,0,0),(977145798077599745,974622695435517953,1716029312,0,0),(977145798080909313,974622695435517953,1716032912,0,0),(977145798083928065,974622695435517953,1716036512,0,0),(977145798088220673,974622695435517953,1716040112,0,0),(977145798091464705,974622695435517953,1716043712,0,0),(977145798098976769,974622695435517953,1716047312,0,0),(977145798102499329,974622695435517953,1716054512,0,0),(977145798102798337,974622695435517953,1716050912,0,0),(977145798104133633,974622695435517953,1716058112,0,0),(977145798106832897,974622695435517953,1716061712,0,0),(977145798110461953,974622695435517953,1716065312,0,0),(977145798111522817,974622695435517953,1716068912,0,0),(977145798112862209,974622695435517953,1716076112,0,0),(977145798114217985,974622695435517953,1716072512,0,0),(977145798116876289,974622695435517953,1716083312,0,0),(977145798119419905,974622695435517953,1716079712,0,0),(977145798123286529,974622695435517953,1716086912,0,0),(977145798126567425,974622695435517953,1716094112,0,0),(977145798128848897,974622695435517953,1716090512,0,0),(977145798132101121,974622695435517953,1716101312,0,0),(977145798132248577,974622695435517953,1716097712,0,0),(977145798134169601,974622695435517953,1716112112,0,0),(977145798134218753,974622695435517953,1716108512,0,0),(977145798137151489,974622695435517953,1716104912,0,0),(977145798137950209,974622695435517953,1716119312,0,0),(977145798139830273,974622695435517953,1716115712,0,0),(977145798142541825,974622695435517953,1716126512,0,0),(977145798143500289,974622695435517953,1716122912,0,0),(977145798146686977,974622695435517953,1716130112,0,0),(977145798153166849,974622695435517953,1716133712,0,0),(977145798153924609,974622695435517953,1716137312,0,0),(977145798155034625,974622695435517953,1716140912,0,0),(977145798156156929,974622695435517953,1716144512,0,0),(977145798158737409,974622695435517953,1716151712,0,0),(977145798161039361,974622695435517953,1716148112,0,0),(977145798162673665,974622695435517953,1716158912,0,0),(977145798164979713,974622695435517953,1716155312,0,0),(977145798170427393,974622695435517953,1716162512,0,0),(977145798170558465,974622695435517953,1716166112,0,0),(977145798171238401,974622695435517953,1716173312,0,0),(977145798172180481,974622695435517953,1716169712,0,0),(977145798178516993,974622695435517953,1716180512,0,0),(977145798179295233,974622695435517953,1716176912,0,0),(977145798179729409,974622695435517953,1716184112,0,0),(977145798180904961,974622695435517953,1716187712,0,0),(977145798184230913,974622695435517953,1716194912,0,0),(977145798186070017,974622695435517953,1716191312,0,0),(977374919148429313,974622695435517953,1716080339,0,0),(977374919155302401,974622695435517953,1716083939,0,0),(977374919159025665,974622695435517953,1716087539,0,0),(977374919163871233,974622695435517953,1716091139,0,0),(977374919169048577,974622695435517953,1716094739,0,0),(977374919171350529,974622695435517953,1716098339,0,0),(977374919174168577,974622695435517953,1716105539,0,0),(977374919178166273,974622695435517953,1716101939,0,0),(977374919181627393,974622695435517953,1716109139,0,0),(977374919184240641,974622695435517953,1716112739,0,0),(977374919187185665,974622695435517953,1716119939,0,0),(977374919190429697,974622695435517953,1716116339,0,0),(977374919191912449,974622695435517953,1716123539,0,0),(977374919195856897,974622695435517953,1716130739,0,0),(977374919197822977,974622695435517953,1716127139,0,0),(977374919201419265,974622695435517953,1716134339,0,0),(977374919202050049,974622695435517953,1716137939,0,0),(977374919205801985,974622695435517953,1716141539,0,0),(977374919206772737,974622695435517953,1716145139,0,0),(977374919208742913,974622695435517953,1716152339,0,0),(977374919211446273,974622695435517953,1716148739,0,0),(977374919215099905,974622695435517953,1716159539,0,0),(977374919215558657,974622695435517953,1716155939,0,0),(977374919220817921,974622695435517953,1716163139,0,0),(977374919225593857,974622695435517953,1716166739,0,0),(977374919226191873,974622695435517953,1716170339,0,0),(977374919229751297,974622695435517953,1716177539,0,0),(977374919231946753,974622695435517953,1716173939,0,0),(977374919233712129,974622695435517953,1716181139,0,0),(977374919235272705,974622695435517953,1716184739,0,0),(977374919240196097,974622695435517953,1716188339,0,0),(977374919242317825,974622695435517953,1716191939,0,0),(977374919243296769,974622695435517953,1716195539,0,0),(977374919246688257,974622695435517953,1716199139,0,0),(977374919247368193,974622695435517953,1716202739,0,0),(977374919249907713,974622695435517953,1716206339,0,0),(977374919254679553,974622695435517953,1716209939,0,0),(977374919257128961,974622695435517953,1716213539,0,0),(977374919260250113,974622695435517953,1716220739,0,0),(977374919261671425,974622695435517953,1716217139,0,0),(977374919264395265,974622695435517953,1716224339,0,0),(977374919266234369,974622695435517953,1716227939,0,0),(977374919268233217,974622695435517953,1716231539,0,0),(977374919272017921,974622695435517953,1716235139,0,0),(977374919274745857,974622695435517953,1716238739,0,0),(977374919276711937,974622695435517953,1716245939,0,0),(977374919277051905,974622695435517953,1716242339,0,0),(977374919282302977,974622695435517953,1716249539,0,0),(977382666200461313,974622695435517953,1716082186,0,0),(977382666214944769,974622695435517953,1716085786,0,0),(977382666222776321,974622695435517953,1716089386,0,0),(977382666231693313,974622695435517953,1716092986,0,0),(977382666235518977,974622695435517953,1716096586,0,0),(977382666241536001,974622695435517953,1716100186,0,0),(977382666245824513,974622695435517953,1716103786,0,0),(977382666249121793,974622695435517953,1716107386,0,0),(977382666253410305,974622695435517953,1716110986,0,0),(977382666256433153,974622695435517953,1716114586,0,0),(977382666260992001,974622695435517953,1716118186,0,0),(977382666265726977,974622695435517953,1716121786,0,0),(977382666267611137,974622695435517953,1716125386,0,0),(977382666268979201,974622695435517953,1716128986,0,0),(977382666279956481,974622695435517953,1716132586,0,0),(977382666280394753,974622695435517953,1716136186,0,0),(977382666287116289,974622695435517953,1716143386,0,0),(977382666287230977,974622695435517953,1716139786,0,0),(977382666290454529,974622695435517953,1716146986,0,0),(977382666295689217,974622695435517953,1716154186,0,0),(977382666296213505,974622695435517953,1716150586,0,0),(977382666300715009,974622695435517953,1716157786,0,0),(977382666304221185,974622695435517953,1716161386,0,0),(977382666308263937,974622695435517953,1716164986,0,0),(977382666320019457,974622695435517953,1716168586,0,0),(977382666322509825,974622695435517953,1716172186,0,0),(977382666325725185,974622695435517953,1716175786,0,0),(977382666333536257,974622695435517953,1716179386,0,0),(977382666334498817,974622695435517953,1716182986,0,0),(977382666334990337,974622695435517953,1716186586,0,0),(977382666340782081,974622695435517953,1716190186,0,0),(977382666344796161,974622695435517953,1716193786,0,0),(977382666346053633,974622695435517953,1716197386,0,0),(977382666347597825,974622695435517953,1716204586,0,0),(977382666349453313,974622695435517953,1716200986,0,0),(977382666354016257,974622695435517953,1716208186,0,0),(977382666356572161,974622695435517953,1716215386,0,0),(977382666359115777,974622695435517953,1716211786,0,0),(977382666360815617,974622695435517953,1716218986,0,0),(977382666363289601,974622695435517953,1716222586,0,0),(977382666365652993,974622695435517953,1716229786,0,0),(977382666367574017,974622695435517953,1716226186,0,0),(977382666370031617,974622695435517953,1716236986,0,0),(977382666371358721,974622695435517953,1716233386,0,0),(977382666374561793,974622695435517953,1716244186,0,0),(977382666374983681,974622695435517953,1716240586,0,0),(977382666379505665,974622695435517953,1716247786,0,0),(977382666383806465,974622695435517953,1716251386,0,0),(977397767694274561,974622695435517953,1716085786,0,0),(977397767700434945,974622695435517953,1716089386,0,0),(977397767707807745,974622695435517953,1716092986,0,0),(977397767714201601,974622695435517953,1716096586,0,0),(977397767725547521,974622695435517953,1716100186,0,0),(977397767730057217,974622695435517953,1716103786,0,0),(977397767734214657,974622695435517953,1716107386,0,0),(977397767739437057,974622695435517953,1716110986,0,0),(977397767742734337,974622695435517953,1716114586,0,0),(977397767742955521,974622695435517953,1716118186,0,0),(977397767746850817,974622695435517953,1716121786,0,0),(977397767751557121,974622695435517953,1716128986,0,0),(977397767751823361,974622695435517953,1716125386,0,0),(977397767757950977,974622695435517953,1716132586,0,0),(977397767760711681,974622695435517953,1716139786,0,0),(977397767760850945,974622695435517953,1716136186,0,0),(977397767765389313,974622695435517953,1716143386,0,0),(977397767768543233,974622695435517953,1716150586,0,0),(977397767769747457,974622695435517953,1716146986,0,0),(977397767771967489,974622695435517953,1716154186,0,0),(977397767775911937,974622695435517953,1716157786,0,0),(977397767778885633,974622695435517953,1716161386,0,0),(977397767780614145,974622695435517953,1716164986,0,0),(977397767783776257,974622695435517953,1716168586,0,0),(977397767784976385,974622695435517953,1716172186,0,0),(977397767788691457,974622695435517953,1716175786,0,0),(977397767792689153,974622695435517953,1716182986,0,0),(977397767795146753,974622695435517953,1716179386,0,0),(977397767797309441,974622695435517953,1716186586,0,0),(977397767798845441,974622695435517953,1716190186,0,0),(977397767802368001,974622695435517953,1716193786,0,0),(977397767806459905,974622695435517953,1716197386,0,0),(977397767808700417,974622695435517953,1716200986,0,0),(977397767812931585,974622695435517953,1716204586,0,0),(977397767814201345,974622695435517953,1716211786,0,0),(977397767814496257,974622695435517953,1716208186,0,0),(977397767818051585,974622695435517953,1716215386,0,0),(977397767822995457,974622695435517953,1716218986,0,0),(977397767832342529,974622695435517953,1716226186,0,0),(977397767833178113,974622695435517953,1716222586,0,0),(977397767835971585,974622695435517953,1716233386,0,0),(977397767836602369,974622695435517953,1716229786,0,0),(977397767838584833,974622695435517953,1716236986,0,0),(977397767843287041,974622695435517953,1716244186,0,0),(977397767843622913,974622695435517953,1716240586,0,0),(977397767848931329,974622695435517953,1716251386,0,0),(977397767849324545,974622695435517953,1716247786,0,0),(977397767854653441,974622695435517953,1716254986,0,0);
/*!40000 ALTER TABLE `DexunLineChartStats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunPicRC`
--

DROP TABLE IF EXISTS `DexunPicRC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunPicRC` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `active` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunPicRC`
--

LOCK TABLES `DexunPicRC` WRITE;
/*!40000 ALTER TABLE `DexunPicRC` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunPicRC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunPreAccessCon`
--

DROP TABLE IF EXISTS `DexunPreAccessCon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunPreAccessCon` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `config` json DEFAULT NULL,
  `active` tinyint DEFAULT NULL,
  `type` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunPreAccessCon`
--

LOCK TABLES `DexunPreAccessCon` WRITE;
/*!40000 ALTER TABLE `DexunPreAccessCon` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunPreAccessCon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunSafeAccessCon`
--

DROP TABLE IF EXISTS `DexunSafeAccessCon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunSafeAccessCon` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `config` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunSafeAccessCon`
--

LOCK TABLES `DexunSafeAccessCon` WRITE;
/*!40000 ALTER TABLE `DexunSafeAccessCon` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunSafeAccessCon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunSSLList`
--

DROP TABLE IF EXISTS `DexunSSLList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunSSLList` (
  `id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `p_type_id` bigint DEFAULT NULL,
  `p_type_name` varchar(45) DEFAULT NULL,
  `ssl_name` varchar(45) DEFAULT NULL,
  `ssl_type` varchar(45) DEFAULT NULL,
  `domain_type` bigint DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `ssl_code` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `ey_money` bigint DEFAULT NULL,
  `p_note` varchar(45) DEFAULT NULL,
  `term` varchar(45) DEFAULT NULL,
  `s_type` bigint DEFAULT NULL,
  `market_money` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunSSLList`
--

LOCK TABLES `DexunSSLList` WRITE;
/*!40000 ALTER TABLE `DexunSSLList` DISABLE KEYS */;
INSERT INTO `DexunSSLList` VALUES (972473959082360833,'68b21ac5-8fb9-43af-9fc2-950a185eb8a4',1,'Sectigo','PositiveSSL 证书','DV',1,1,'PositiveSSL',29,0,'','1年',1,199),(972474337089781761,'60f60ccc-ff69-471d-a813-cbf75e0d3425',1,'Sectigo','PositiveSSL 通配符SSL证书','DV',3,1,'WildPositiveSSL',688,0,'','1年',1,2000),(972474337099784193,'bd5fd8d1-b34f-433d-a199-cdeca5601aac',1,'Sectigo','PositiveSSL 多域名SSL证书','DV',2,3,'SanPositiveSSL',299,100,'','1年',1,1500),(972474337104293889,'cb353de9-6f14-4998-8e00-6271bea0c717',1,'Sectigo','PositiveSSL 多域名通配符SSL证书','DV',3,2,'SanWildPositiveSSL',999,499,'','1年',1,3500),(972474418523697153,'0c40fcda-1ca2-45b0-8511-84785a24b244',1,'Sectigo','企业 OV SSL证书','OV',1,1,'SectigoOV',1000,0,'','1年',0,0),(972474418529435649,'60979f08-0c0e-4cb1-a72b-58500930bfb0',1,'Sectigo','企业 OV 多域名SSL证书','OV',2,3,'sanSectigoOV',2000,665,'','1年',0,0),(972474418531966977,'7caab045-8b3e-453c-8b24-331e3ed2132b',1,'Sectigo','企业 OV 通配符SSL证书','OV',3,1,'wildSectigoOV',2800,0,'','1年',0,0),(972474418538606593,'7894669c-409c-4e4c-b1ae-3c405aa535f8',1,'Sectigo','企业 OV 多域名通配符SSL证书','OV',3,2,'SanwildSectigoOV',7500,3750,'','1年',0,0),(972474418543550465,'8076fa5b-8b32-442d-83ea-cb4e0452ff2c',1,'Sectigo','PositiveSSL EV 证书','EV',1,1,'PositiveSSLEV',1200,0,'','1年',0,0),(972474418546782209,'39705d1d-4935-4bea-80c3-3f17c08a8ef1',1,'Sectigo','PositiveSSL EV 多域名SSL证书','EV',2,3,'PositiveSSLEVMD',2500,835,'','1年',0,100);
/*!40000 ALTER TABLE `DexunSSLList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunTotalFlow`
--

DROP TABLE IF EXISTS `DexunTotalFlow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunTotalFlow` (
  `id` bigint NOT NULL,
  `order_id` bigint DEFAULT NULL,
  `request_bandwidth_peak` bigint DEFAULT NULL,
  `requests` bigint DEFAULT NULL,
  `response_bandwidth_peak` bigint DEFAULT NULL,
  `total_request_flows` bigint DEFAULT NULL,
  `total_response_flows` bigint DEFAULT NULL,
  `unidentified_attack` bigint DEFAULT NULL,
  `create_time` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid_idx` (`order_id`),
  CONSTRAINT `fk_orderid26` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunTotalFlow`
--

LOCK TABLES `DexunTotalFlow` WRITE;
/*!40000 ALTER TABLE `DexunTotalFlow` DISABLE KEYS */;
INSERT INTO `DexunTotalFlow` VALUES (974638140044861441,974622695435517953,0,0,0,0,0,0,1715600640),(974712121640493057,974622695435517953,0,0,0,0,0,0,1715618279),(974880920002813953,974622695435517953,0,0,0,0,0,0,1715658523),(974881268750860289,974622695435517953,0,0,0,0,0,0,1715658606),(974882516186083329,974622695435517953,0,0,0,0,0,0,1715658904),(974912523495178241,974622695435517953,0,0,0,0,0,0,1715666058),(974912894826856449,974622695435517953,0,0,0,0,0,0,1715666147),(974914800348540929,974622695435517953,0,0,0,0,0,0,1715666601),(974927858043916289,974622695435517953,0,0,0,0,0,0,1715669714),(974930129669169153,974622695435517953,0,0,0,0,0,0,1715670256),(974945242673242113,974622695435517953,0,0,0,0,0,0,1715673859),(974962615554654209,974622695435517953,0,0,0,0,0,0,1715678001),(974964420360114177,974622695435517953,0,0,0,0,0,0,1715678431),(974965668110991361,974622695435517953,0,0,0,0,0,0,1715678729),(974965980806324225,974622695435517953,0,0,0,0,0,0,1715678803),(974976125800751105,974622695435517953,0,0,0,0,0,0,1715681222),(974980408796614657,974622695435517953,0,0,0,0,0,0,1715682243),(974980948005384193,974622695435517953,0,0,0,0,0,0,1715682372),(974983647771377665,974622695435517953,0,0,0,0,0,0,1715683015),(974996409060360193,974622695435517953,0,0,0,0,0,0,1715686058),(974997750843224065,974622695435517953,0,0,0,0,0,0,1715686378),(974999200420724737,974622695435517953,0,0,0,0,0,0,1715686723),(974999691764727809,974622695435517953,0,0,0,0,0,0,1715686841),(975000037624983553,974622695435517953,0,0,0,0,0,0,1715686923),(975000623794552833,974622695435517953,0,0,0,0,0,0,1715687063),(975015725380837377,974622695435517953,0,0,0,0,0,0,1715690663),(977013924187189249,974622695435517953,0,0,0,0,0,0,1716167071),(977015256233050113,974622695435517953,0,0,0,0,0,0,1716167389),(977016630342692865,974622695435517953,0,0,0,0,0,0,1716167716),(977019160395902977,974622695435517953,0,0,0,0,0,0,1716168319),(977029851343286273,974622695435517953,0,0,0,0,0,0,1716170868),(977044953140408321,974622695435517953,0,0,0,0,0,0,1716174469),(977091874720530433,974622695435517953,0,0,0,0,0,0,1716185656),(977093121563799553,974622695435517953,0,0,0,0,0,0,1716185953),(977094378404651009,974622695435517953,0,0,0,0,0,0,1716186253),(977095161701621761,974622695435517953,0,0,0,0,0,0,1716186440),(977105548740743169,974622695435517953,0,0,0,0,0,0,1716188916),(977106898276782081,974622695435517953,0,0,0,0,0,0,1716189238),(977108191133110273,974622695435517953,0,0,0,0,0,0,1716189546),(977115954884816897,974622695435517953,0,0,0,0,0,0,1716191397),(977122233882198017,974622695435517953,0,0,0,0,0,0,1716192894),(977123512819642369,974622695435517953,0,0,0,0,0,0,1716193199),(977124529053790209,974622695435517953,0,0,0,0,0,0,1716193441),(977129159523942401,974622695435517953,0,0,0,0,0,0,1716194545),(977130700446703617,974622695435517953,0,0,0,0,0,0,1716194913),(977145801378058241,974622695435517953,0,0,0,0,0,0,1716198513),(977374922576863233,974622695435517953,0,0,0,0,0,0,1716253140),(977382669591756801,974622695435517953,0,0,0,0,0,0,1716254987),(977397770798505985,974622695435517953,0,0,0,0,0,0,1716258587);
/*!40000 ALTER TABLE `DexunTotalFlow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunURLBWLists`
--

DROP TABLE IF EXISTS `DexunURLBWLists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunURLBWLists` (
  `id` bigint NOT NULL,
  `bw_id` bigint DEFAULT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `type` bigint DEFAULT NULL,
  `path` varchar(45) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `active` bigint DEFAULT NULL,
  `uuid` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunURLBWLists`
--

LOCK TABLES `DexunURLBWLists` WRITE;
/*!40000 ALTER TABLE `DexunURLBWLists` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunURLBWLists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `DexunWordRC`
--

DROP TABLE IF EXISTS `DexunWordRC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DexunWordRC` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `gzip` tinyint DEFAULT NULL,
  `active` tinyint DEFAULT NULL,
  `keywords` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DexunWordRC`
--

LOCK TABLES `DexunWordRC` WRITE;
/*!40000 ALTER TABLE `DexunWordRC` DISABLE KEYS */;
/*!40000 ALTER TABLE `DexunWordRC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Enterprise`
--

DROP TABLE IF EXISTS `Enterprise`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Enterprise` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `corp_name` varchar(45) DEFAULT NULL,
  `reg_num` varchar(45) DEFAULT NULL,
  `lgman` varchar(45) DEFAULT NULL,
  `lgperson_num` varchar(45) DEFAULT NULL,
  `corp_address` varchar(45) DEFAULT NULL,
  `corp_doc` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `lgperson_front` varchar(255) DEFAULT NULL,
  `lgperson_back` varchar(255) DEFAULT NULL,
  `del_time` int DEFAULT NULL,
  `create` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_id_idx` (`user_id`),
  CONSTRAINT `fk_userid_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Enterprise`
--

LOCK TABLES `Enterprise` WRITE;
/*!40000 ALTER TABLE `Enterprise` DISABLE KEYS */;
/*!40000 ALTER TABLE `Enterprise` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Logs`
--

DROP TABLE IF EXISTS `Logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Logs` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `req_url` varchar(45) DEFAULT NULL,
  `origin_url` varchar(45) DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  `request` json DEFAULT NULL,
  `req_time` int DEFAULT NULL,
  `create` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_userid_6_idx` (`user_id`),
  CONSTRAINT `fk_userid_6` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Logs`
--

LOCK TABLES `Logs` WRITE;
/*!40000 ALTER TABLE `Logs` DISABLE KEYS */;
INSERT INTO `Logs` VALUES (957486818344816641,956929351825580033,'','http://192.168.2.112:4000','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0',NULL,0,1711511446),(957486990813528065,956929351825580033,'','http://192.168.2.112:4000','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0',NULL,0,1711511488);
/*!40000 ALTER TABLE `Logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PackageService`
--

DROP TABLE IF EXISTS `PackageService`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PackageService` (
  `id` bigint NOT NULL,
  `package_id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `protect_note` varchar(45) DEFAULT NULL,
  `protect_lv` varchar(45) DEFAULT NULL,
  `protect_price` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  `sell_status` int DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  `sell_price` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_uuid_2_idx` (`package_id`),
  CONSTRAINT `fk_uuid_2` FOREIGN KEY (`package_id`) REFERENCES `ScdnPackage` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PackageService`
--

LOCK TABLES `PackageService` WRITE;
/*!40000 ALTER TABLE `PackageService` DISABLE KEYS */;
INSERT INTO `PackageService` VALUES (0,0,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(962229000846303233,962194988101332993,'1310f592-2418-422e-be70-9927a67ebe35','特惠500G','500',50,1,1,'DeXun','flow',0),(962229000850497537,962194988105076737,'d51223cf-e8fb-4e7d-9721-f7c041a3e3b2','特惠1000G','1000',100,1,1,'DeXun','flow',0),(962229000854691841,962194988111028225,'ce0920f0-6323-4be1-8e2b-bdadc98271b2','特惠10000G','10000',1000,1,1,'DeXun','flow',0),(962229000854694515,962194988112424961,'74852d04-91b7-4347-868b-52449502f5c9','特惠5000G','5000',500,1,1,'DeXun','flow',0);
/*!40000 ALTER TABLE `PackageService` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Person`
--

DROP TABLE IF EXISTS `Person`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Person` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `real_name` varchar(45) DEFAULT NULL,
  `card_id` varchar(45) DEFAULT NULL,
  `sex` varchar(45) DEFAULT NULL,
  `birthday` varchar(45) DEFAULT NULL,
  `city` varchar(45) DEFAULT NULL,
  `status` int DEFAULT '0',
  `card_front` varchar(255) DEFAULT NULL,
  `card_back` varchar(255) DEFAULT NULL,
  `del_time` int DEFAULT NULL,
  `create` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_id_idx` (`user_id`),
  CONSTRAINT `fk_userid_5` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Person`
--

LOCK TABLES `Person` WRITE;
/*!40000 ALTER TABLE `Person` DISABLE KEYS */;
INSERT INTO `Person` VALUES (957207141331554305,956929351825580033,'1','1','','','',1,'http://localhost:8088/cert/262623042626/hjdgLcGQoDIKOJdOlesWkwXeyiRCtutb/1711441380802077000.jpg','http://localhost:8088/cert/262623042626/hjdgLcGQoDIKOJdOlesWkwXeyiRCtutb/1711441427171891000.jpg',0,1711444766),(957494999259328513,957493600380858369,'1','1','','','',1,'http://localhost:8088/cert/272721122727/gnXCBYhiAqILlViSImvNhllTioWnGYLo/1711513262840531000.png','http://localhost:8088/cert/272721122727/gnXCBYhiAqILlViSImvNhllTioWnGYLo/1711513302917860000.png',0,1711513397);
/*!40000 ALTER TABLE `Person` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnAreaAccessCon`
--

DROP TABLE IF EXISTS `ScdnAreaAccessCon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnAreaAccessCon` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `order_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `da_id` bigint NOT NULL,
  `regions` json DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid11_idx` (`order_id`),
  KEY `fk_domainid10_idx` (`domain_id`),
  CONSTRAINT `fk_domainid10` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid11` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnAreaAccessCon`
--

LOCK TABLES `ScdnAreaAccessCon` WRITE;
/*!40000 ALTER TABLE `ScdnAreaAccessCon` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnAreaAccessCon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnATKLogLists`
--

DROP TABLE IF EXISTS `ScdnATKLogLists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnATKLogLists` (
  `id` bigint NOT NULL,
  `status` tinyint DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `di_id` bigint NOT NULL,
  `clientip` varchar(45) DEFAULT NULL,
  `clientport` bigint DEFAULT NULL,
  `clientregion` varchar(45) DEFAULT NULL,
  `targeturl` varchar(45) DEFAULT NULL,
  `nodeid` varchar(45) DEFAULT NULL,
  `httpmethod` varchar(45) DEFAULT NULL,
  `attackinfo` varchar(45) DEFAULT NULL,
  `attacktype` varchar(45) DEFAULT NULL,
  `protecttype` varchar(45) DEFAULT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `requestinfo` varchar(45) DEFAULT NULL,
  `domainid` varchar(45) DEFAULT NULL,
  `count` varchar(45) DEFAULT NULL,
  `localip` varchar(45) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `timerangeend` varchar(45) DEFAULT NULL,
  `timerangestart` varchar(45) DEFAULT NULL,
  `instanceid` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dpid_idx` (`di_id`),
  KEY `fk_orderid16_idx` (`order_id`),
  KEY `fk-domain15_idx` (`domain_id`),
  CONSTRAINT `fk-domain15` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_diid` FOREIGN KEY (`di_id`) REFERENCES `DexunATKLogLists` (`id`),
  CONSTRAINT `fk_orderid16` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnATKLogLists`
--

LOCK TABLES `ScdnATKLogLists` WRITE;
/*!40000 ALTER TABLE `ScdnATKLogLists` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnATKLogLists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnBWInstance`
--

DROP TABLE IF EXISTS `ScdnBWInstance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnBWInstance` (
  `id` bigint NOT NULL,
  `bw_id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `type` bigint DEFAULT NULL,
  `ip_list` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_bwid_idx` (`bw_id`),
  KEY `fk_domainid4_idx` (`domain_id`),
  KEY `fk_orderid5_idx` (`order_id`),
  CONSTRAINT `fk_bwid` FOREIGN KEY (`bw_id`) REFERENCES `DexunBWInstance` (`id`),
  CONSTRAINT `fk_domainid4` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid5` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnBWInstance`
--

LOCK TABLES `ScdnBWInstance` WRITE;
/*!40000 ALTER TABLE `ScdnBWInstance` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnBWInstance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnBWSingle`
--

DROP TABLE IF EXISTS `ScdnBWSingle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnBWSingle` (
  `id` bigint NOT NULL,
  `bws_id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `type` bigint DEFAULT NULL,
  `ip_list` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_bwid2_idx` (`bws_id`),
  KEY `fk_orderid6_idx` (`order_id`),
  KEY `fk_domainid5_idx` (`domain_id`),
  CONSTRAINT `fk_bwid2` FOREIGN KEY (`bws_id`) REFERENCES `DexunBWSingle` (`id`),
  CONSTRAINT `fk_domainid5` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid6` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnBWSingle`
--

LOCK TABLES `ScdnBWSingle` WRITE;
/*!40000 ALTER TABLE `ScdnBWSingle` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnBWSingle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnCache`
--

DROP TABLE IF EXISTS `ScdnCache`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnCache` (
  `id` bigint NOT NULL,
  `dxcache_id` bigint DEFAULT NULL,
  `cache_id` bigint DEFAULT NULL,
  `dd_uuid` varchar(45) DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `cache_uuid` varchar(45) DEFAULT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint NOT NULL,
  `active` bigint DEFAULT NULL,
  `urlmode` varchar(45) DEFAULT NULL,
  `cachemode` varchar(45) DEFAULT NULL,
  `cachepath` varchar(45) DEFAULT NULL,
  `cacheextensions` varchar(45) DEFAULT NULL,
  `cachereg` varchar(45) DEFAULT NULL,
  `timeout` varchar(45) DEFAULT NULL,
  `weight` varchar(45) DEFAULT NULL,
  `createtime` varchar(45) DEFAULT NULL,
  `updatetime` varchar(45) DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid_2_idx` (`order_id`),
  KEY `fk_domainid_idx` (`domain_id`),
  CONSTRAINT `fk_domainid2` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid_2` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnCache`
--

LOCK TABLES `ScdnCache` WRITE;
/*!40000 ALTER TABLE `ScdnCache` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnCache` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnCacheModel`
--

DROP TABLE IF EXISTS `ScdnCacheModel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnCacheModel` (
  `id` bigint NOT NULL,
  `cm_id` bigint NOT NULL,
  `dd_id` bigint NOT NULL,
  `dd_type` bigint DEFAULT NULL,
  `cache_name` varchar(45) DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  `urlmode` varchar(45) DEFAULT NULL,
  `cachemode` varchar(45) DEFAULT NULL,
  `cachepath` varchar(45) DEFAULT NULL,
  `cacheextensions` varchar(45) DEFAULT NULL,
  `cachereg` varchar(45) DEFAULT NULL,
  `timeout` bigint DEFAULT NULL,
  `weight` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_cmid_idx` (`cm_id`),
  KEY `fk_orderid3_idx` (`dd_id`),
  CONSTRAINT `fk_cmid` FOREIGN KEY (`cm_id`) REFERENCES `DexunCacheModel` (`id`),
  CONSTRAINT `fk_orderid3` FOREIGN KEY (`dd_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnCacheModel`
--

LOCK TABLES `ScdnCacheModel` WRITE;
/*!40000 ALTER TABLE `ScdnCacheModel` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnCacheModel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnCC`
--

DROP TABLE IF EXISTS `ScdnCC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnCC` (
  `id` bigint NOT NULL,
  `cc_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `cs_active` varchar(45) DEFAULT NULL,
  `policy` varchar(45) DEFAULT NULL,
  `url` varchar(45) DEFAULT NULL,
  `rate` varchar(45) DEFAULT NULL,
  `waitseconds` varchar(45) DEFAULT NULL,
  `blockminutes` varchar(45) DEFAULT NULL,
  `redirectlocation` varchar(45) DEFAULT NULL,
  `global_concurrent` varchar(45) DEFAULT NULL,
  `waitpolicyminutes` varchar(45) DEFAULT NULL,
  `redirectwaitseconds` varchar(45) DEFAULT NULL,
  `count` varchar(45) DEFAULT NULL,
  `block_time` varchar(45) DEFAULT NULL,
  `block_active` varchar(45) DEFAULT NULL,
  `rr_active` varchar(45) DEFAULT NULL,
  `rr_rate` varchar(45) DEFAULT NULL,
  `ur_url` varchar(45) DEFAULT NULL,
  `ur_rate` varchar(45) DEFAULT NULL,
  `cookieName` varchar(45) DEFAULT NULL,
  `excludeExt` varchar(45) DEFAULT NULL,
  `concurrency` varchar(45) DEFAULT NULL,
  `r_blockMinutes` varchar(45) DEFAULT NULL,
  `whiteMinutes` varchar(45) DEFAULT NULL,
  `challengeLimit` varchar(45) DEFAULT NULL,
  `protectMinutes` varchar(45) DEFAULT NULL,
  `challengeMethods` json DEFAULT NULL,
  `challengePolicy` varchar(45) DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  `use_default` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid8_idx` (`order_id`),
  KEY `fk_domain7_idx` (`domain_id`),
  CONSTRAINT `fk_domain7` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid8` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnCC`
--

LOCK TABLES `ScdnCC` WRITE;
/*!40000 ALTER TABLE `ScdnCC` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnCC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnCombos`
--

DROP TABLE IF EXISTS `ScdnCombos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnCombos` (
  `id` bigint NOT NULL,
  `combo_id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `yms` varchar(45) DEFAULT NULL,
  `ccfy` varchar(45) DEFAULT NULL,
  `gjwaf` varchar(45) DEFAULT NULL,
  `ywll` varchar(45) DEFAULT NULL,
  `pro_flow` bigint DEFAULT NULL,
  `ddos_hh` varchar(45) DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  `complete_state` bigint DEFAULT NULL,
  `firewall_state` bigint DEFAULT NULL,
  `waf_state` bigint DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `zyzk_money` varchar(45) DEFAULT NULL,
  `zyks_money` varchar(45) DEFAULT NULL,
  `sell_status` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dtid_idx` (`combo_id`),
  CONSTRAINT `fk_dtid` FOREIGN KEY (`combo_id`) REFERENCES `DexunCombos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnCombos`
--

LOCK TABLES `ScdnCombos` WRITE;
/*!40000 ALTER TABLE `ScdnCombos` DISABLE KEYS */;
INSERT INTO `ScdnCombos` VALUES (962976024390311937,962974346923986945,'9d8a6928-149b-43da-8cc4-9482610305df','体验版',99,'1','1','2','100',100,'10GB峰值',1,79,2,1,0,'Dexun',1,'0','297',1),(962976024394506241,962974346928787457,'cd16fc25-a897-4623-a8b6-7237b6a833d7','入门版',599,'10','5','2','1000',1000,'100GB峰值',10,479,2,1,0,'Dexun',1,'0','1797',1),(962976024398700545,962974346933051393,'25fc50c0-0b53-4336-b6c5-6abf6b50ed85','基础版',999,'30','20','2','3600',3600,'100GB峰值',30,799,2,1,0,'Dexun',1,'0','2997',1),(962976024402894849,962974346935853057,'32aee6cf-50e7-4172-b2ae-42d0e5f0ce6d','专业版',1999,'50','30','1','5000',5000,'200GB峰值',50,1599,2,1,1,'Dexun',1,'0','5997',1),(962976024407089153,962974346939039745,'f4b907d1-0524-4a07-b0cf-030d659e4db1','商务版',4999,'50','80','1','10000',10000,'400GB峰值',50,3999,2,1,1,'Dexun',1,'0','14997',1),(962976024407089753,962974346939113473,'c3a7f0c2-531f-4ac7-a374-1633d3cad920','企业版',2999,'50','50','1','7200',7200,'300GB峰值',50,2399,2,1,1,'Dexun',1,'0','8997',1),(962976024411283457,962974346942824449,'61cca695-414e-45f6-818f-ca8d6e8e8996','旗舰版',9999,'50','100','1','15000',15000,'600GB峰值',50,7999,2,1,1,'Dexun',1,'0','29997',1);
/*!40000 ALTER TABLE `ScdnCombos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnDomainCert`
--

DROP TABLE IF EXISTS `ScdnDomainCert`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnDomainCert` (
  `id` bigint NOT NULL,
  `ddc_id` bigint DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `status` bigint DEFAULT NULL,
  `ssl_always` bigint DEFAULT NULL,
  `hsts` bigint DEFAULT NULL,
  `cert_name` varchar(45) DEFAULT NULL,
  `cert` varchar(255) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  `desc` varchar(45) DEFAULT NULL,
  `createtime` varchar(45) DEFAULT NULL,
  `updatetime` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid4_idx` (`order_id`),
  KEY `fk_domainid2_idx` (`domain_id`),
  CONSTRAINT `fk_domainid3` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid4` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnDomainCert`
--

LOCK TABLES `ScdnDomainCert` WRITE;
/*!40000 ALTER TABLE `ScdnDomainCert` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnDomainCert` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnDomains`
--

DROP TABLE IF EXISTS `ScdnDomains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnDomains` (
  `id` bigint NOT NULL,
  `domain` varchar(45) DEFAULT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint DEFAULT NULL,
  `order_id` bigint DEFAULT NULL,
  `primary_domain` varchar(45) DEFAULT NULL,
  `domain_status` bigint DEFAULT NULL,
  `domain_record` varchar(45) DEFAULT NULL,
  `four_layers_config` json DEFAULT NULL,
  `cache_file_size_limit` bigint DEFAULT NULL,
  `cache_total_size_limit` bigint DEFAULT NULL,
  `cache_config` json DEFAULT NULL,
  `cache_active` bigint DEFAULT NULL,
  `white_num` bigint DEFAULT NULL,
  `use_flow` bigint DEFAULT NULL,
  `createtime` varchar(45) DEFAULT NULL,
  `updatetime` varchar(45) DEFAULT NULL,
  `access_active` varchar(45) DEFAULT NULL,
  `grouping` varchar(45) DEFAULT NULL,
  `is_filing` varchar(45) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `user_id` bigint NOT NULL,
  `cname` varchar(255) DEFAULT NULL,
  `source_addresses` json DEFAULT NULL,
  `waf_switch` tinyint DEFAULT NULL,
  `waf_file` tinyint DEFAULT NULL,
  `waf_code` tinyint DEFAULT NULL,
  `waf_session` tinyint DEFAULT NULL,
  `waf_shellshock` tinyint DEFAULT NULL,
  `waf_zombie` tinyint DEFAULT NULL,
  `waf_metadata` tinyint DEFAULT NULL,
  `waf_sql` tinyint DEFAULT NULL,
  `waf_proxy` tinyint DEFAULT NULL,
  `waf_xss` tinyint DEFAULT NULL,
  `ddosdd_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_userid_idx` (`user_id`),
  KEY `fk_ddid_idx` (`order_id`),
  KEY `fk_ddosddid_idx` (`ddosdd_id`),
  CONSTRAINT `fk_ddid` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`),
  CONSTRAINT `fk_ddosddid` FOREIGN KEY (`ddosdd_id`) REFERENCES `DDoSService` (`id`),
  CONSTRAINT `fk_userid` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnDomains`
--

LOCK TABLES `ScdnDomains` WRITE;
/*!40000 ALTER TABLE `ScdnDomains` DISABLE KEYS */;
INSERT INTO `ScdnDomains` VALUES (977016768593276929,'www.cdn23.com','e5e4c486149d594eac0adaf71225647edx73dx23y',0,974622695435517953,'cnameo.com',1,NULL,'[]',2000,20000,'[]',0,0,0,'2024-05-20 09:15:57','2024-05-20 09:16:02','2','','0',1,957493600380858369,'e5e4c486149d594eac0adaf71225647edx73dx23y.cnameo.com',NULL,0,0,0,0,0,0,0,0,0,0,NULL);
/*!40000 ALTER TABLE `ScdnDomains` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnHeat`
--

DROP TABLE IF EXISTS `ScdnHeat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnHeat` (
  `id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `url` varchar(45) DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_domainid_idx` (`domain_id`),
  CONSTRAINT `fk_domainid` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnHeat`
--

LOCK TABLES `ScdnHeat` WRITE;
/*!40000 ALTER TABLE `ScdnHeat` DISABLE KEYS */;
INSERT INTO `ScdnHeat` VALUES (977019391706193921,977016768593276929,'223.83.67.35',1);
/*!40000 ALTER TABLE `ScdnHeat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnLeechlink`
--

DROP TABLE IF EXISTS `ScdnLeechlink`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnLeechlink` (
  `id` bigint NOT NULL,
  `dl_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `dd_uuid` varchar(45) DEFAULT NULL,
  `pro_type` varchar(45) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  `domians` json DEFAULT NULL,
  `allow_empty` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid9_idx` (`order_id`),
  KEY `fk_domain8_idx` (`domain_id`),
  CONSTRAINT `fk_domain8` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid9` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnLeechlink`
--

LOCK TABLES `ScdnLeechlink` WRITE;
/*!40000 ALTER TABLE `ScdnLeechlink` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnLeechlink` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnPackage`
--

DROP TABLE IF EXISTS `ScdnPackage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnPackage` (
  `id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `protect_note` varchar(45) DEFAULT NULL,
  `protect_lv` bigint DEFAULT NULL,
  `protect_price` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnPackage`
--

LOCK TABLES `ScdnPackage` WRITE;
/*!40000 ALTER TABLE `ScdnPackage` DISABLE KEYS */;
INSERT INTO `ScdnPackage` VALUES (0,NULL,NULL,NULL,NULL,NULL,NULL,NULL),(962194988101332993,'1310f592-2418-422e-be70-9927a67ebe35','特惠500G',500,50,1,'DeXun','flow'),(962194988105076737,'d51223cf-e8fb-4e7d-9721-f7c041a3e3b2','特惠1000G',1000,100,1,'DeXun','flow'),(962194988111028225,'ce0920f0-6323-4be1-8e2b-bdadc98271b2','特惠10000G',10000,1000,1,'DeXun','flow'),(962194988112424961,'74852d04-91b7-4347-868b-52449502f5c9','特惠5000G',5000,500,1,'DeXun','flow');
/*!40000 ALTER TABLE `ScdnPackage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnPicRC`
--

DROP TABLE IF EXISTS `ScdnPicRC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnPicRC` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `order_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `dp_id` bigint NOT NULL,
  `active` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_domain12_idx` (`domain_id`),
  KEY `fk_orderid13_idx` (`order_id`),
  CONSTRAINT `fk_domain12` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid13` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnPicRC`
--

LOCK TABLES `ScdnPicRC` WRITE;
/*!40000 ALTER TABLE `ScdnPicRC` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnPicRC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnPreAccessCon`
--

DROP TABLE IF EXISTS `ScdnPreAccessCon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnPreAccessCon` (
  `id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint NOT NULL,
  `order_uuid` varchar(45) DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `da_id` bigint NOT NULL,
  `action` varchar(45) DEFAULT NULL,
  `active` bigint DEFAULT NULL,
  `check_list` varchar(45) DEFAULT NULL,
  `m_item` varchar(45) DEFAULT NULL,
  `m_value` varchar(45) DEFAULT NULL,
  `m_operate` varchar(45) DEFAULT NULL,
  `m_value_xs` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_daid_idx` (`da_id`),
  KEY `fk_orderid10_idx` (`order_id`),
  KEY `fk_domainid9_idx` (`domain_id`),
  CONSTRAINT `fk_daid` FOREIGN KEY (`da_id`) REFERENCES `DexunPreAccessCon` (`id`),
  CONSTRAINT `fk_domainid9` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid10` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnPreAccessCon`
--

LOCK TABLES `ScdnPreAccessCon` WRITE;
/*!40000 ALTER TABLE `ScdnPreAccessCon` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnPreAccessCon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnSafeAccessCon`
--

DROP TABLE IF EXISTS `ScdnSafeAccessCon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnSafeAccessCon` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `order_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `ds_id` bigint NOT NULL,
  `password` varchar(45) DEFAULT NULL,
  `url` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid14_idx` (`order_id`),
  KEY `fk_domainid13_idx` (`domain_id`),
  CONSTRAINT `fk_domainid13` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid14` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnSafeAccessCon`
--

LOCK TABLES `ScdnSafeAccessCon` WRITE;
/*!40000 ALTER TABLE `ScdnSafeAccessCon` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnSafeAccessCon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnService`
--

DROP TABLE IF EXISTS `ScdnService`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnService` (
  `id` bigint NOT NULL,
  `combo_id` bigint NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `tc_name` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `pro_flow` bigint DEFAULT NULL,
  `ddos_hh` varchar(45) DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `zk_money` bigint DEFAULT NULL,
  `source` varchar(45) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `zyzk_money` bigint DEFAULT NULL,
  `zyks_money` bigint DEFAULT NULL,
  `months` bigint DEFAULT NULL,
  `actua_flow` bigint DEFAULT NULL,
  `end_time` varchar(45) DEFAULT NULL,
  `ks_start` bigint DEFAULT NULL,
  `product_sitename` varchar(45) DEFAULT NULL,
  `recharge_flow` bigint DEFAULT NULL,
  `recharge_domain` bigint DEFAULT NULL,
  `server_ip` varchar(255) DEFAULT NULL,
  `stat_time` varchar(45) DEFAULT NULL,
  `total_flow` bigint DEFAULT NULL,
  `u_user_id` bigint DEFAULT NULL,
  `site_stat` bigint DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `agent` varchar(45) DEFAULT NULL,
  `package_id` bigint DEFAULT NULL,
  `username` varchar(45) DEFAULT NULL,
  `pro_type` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_uuidid_idx` (`combo_id`),
  KEY `fk_pkid_idx` (`package_id`),
  CONSTRAINT `fk_pkid` FOREIGN KEY (`package_id`) REFERENCES `PackageService` (`id`),
  CONSTRAINT `fk_uuidid` FOREIGN KEY (`combo_id`) REFERENCES `ScdnCombos` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnService`
--

LOCK TABLES `ScdnService` WRITE;
/*!40000 ALTER TABLE `ScdnService` DISABLE KEYS */;
INSERT INTO `ScdnService` VALUES (974622695435517953,962976024390311937,'f281cf51-3e29-4e45-85fe-1131f22ec820','体验版',99,100,'10GB峰值',1,0,'Dexun',0,0,297,1,0,'2024-06-13',1,'Dx_6641ee6e71ad7',0,0,'122.228.84.70;122.228.86.217;101.67.12.217;120.193.39.217;110.42.101.41;110.42.101.45;110.42.101.28','2024-05-13',0,28145,1,957493600380858369,'',0,'ccc',9);
/*!40000 ALTER TABLE `ScdnService` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnURLBWLists`
--

DROP TABLE IF EXISTS `ScdnURLBWLists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnURLBWLists` (
  `id` bigint NOT NULL,
  `bwl_id` bigint DEFAULT NULL,
  `order_id` bigint NOT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  `path` varchar(45) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `active` bigint DEFAULT NULL,
  `uuid` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid7_idx` (`order_id`),
  KEY `fk_domainid6_idx` (`domain_id`),
  CONSTRAINT `fk_domainid6` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid7` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnURLBWLists`
--

LOCK TABLES `ScdnURLBWLists` WRITE;
/*!40000 ALTER TABLE `ScdnURLBWLists` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnURLBWLists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ScdnWordRC`
--

DROP TABLE IF EXISTS `ScdnWordRC`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ScdnWordRC` (
  `id` bigint NOT NULL,
  `order_id` bigint NOT NULL,
  `order_uuid` varchar(45) DEFAULT NULL,
  `domain_id` bigint NOT NULL,
  `domain_uuid` varchar(45) DEFAULT NULL,
  `dw_id` bigint NOT NULL,
  `gzip` varchar(45) DEFAULT NULL,
  `active` varchar(45) DEFAULT NULL,
  `keywords` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orderid12_idx` (`order_id`),
  KEY `fk_domainid11_idx` (`domain_id`),
  CONSTRAINT `fk_domainid11` FOREIGN KEY (`domain_id`) REFERENCES `ScdnDomains` (`id`),
  CONSTRAINT `fk_orderid12` FOREIGN KEY (`order_id`) REFERENCES `ScdnService` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ScdnWordRC`
--

LOCK TABLES `ScdnWordRC` WRITE;
/*!40000 ALTER TABLE `ScdnWordRC` DISABLE KEYS */;
/*!40000 ALTER TABLE `ScdnWordRC` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SourceAddress`
--

DROP TABLE IF EXISTS `SourceAddress`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SourceAddress` (
  `id` bigint NOT NULL,
  `config_list_id` bigint DEFAULT NULL,
  `address` varchar(45) DEFAULT NULL,
  `concurrent` varchar(45) DEFAULT NULL,
  `port` varchar(45) DEFAULT NULL,
  `protocol` varchar(45) DEFAULT NULL,
  `sni` varchar(45) DEFAULT NULL,
  `weight` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_clid_idx` (`config_list_id`),
  CONSTRAINT `fk_clid` FOREIGN KEY (`config_list_id`) REFERENCES `ConfigList` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SourceAddress`
--

LOCK TABLES `SourceAddress` WRITE;
/*!40000 ALTER TABLE `SourceAddress` DISABLE KEYS */;
INSERT INTO `SourceAddress` VALUES (977017044204634113,977017044195274753,'223.83.67.35','','8080','','','100');
/*!40000 ALTER TABLE `SourceAddress` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SSLList`
--

DROP TABLE IF EXISTS `SSLList`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SSLList` (
  `id` bigint NOT NULL,
  `ssl_id` bigint DEFAULT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `p_type_id` bigint DEFAULT NULL,
  `p_type_name` varchar(45) DEFAULT NULL,
  `ssl_name` varchar(45) DEFAULT NULL,
  `ssl_type` varchar(45) DEFAULT NULL,
  `domain_type` bigint DEFAULT NULL,
  `domain_num` bigint DEFAULT NULL,
  `ssl_code` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `ey_money` bigint DEFAULT NULL,
  `p_note` varchar(45) DEFAULT NULL,
  `term` varchar(45) DEFAULT NULL,
  `s_type` bigint DEFAULT NULL,
  `market_money` bigint DEFAULT NULL,
  `sell_status` int DEFAULT NULL,
  `zyks_money` bigint DEFAULT NULL,
  `zyey_money` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sslid_idx` (`ssl_id`),
  CONSTRAINT `fk_sslid` FOREIGN KEY (`ssl_id`) REFERENCES `DexunSSLList` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SSLList`
--

LOCK TABLES `SSLList` WRITE;
/*!40000 ALTER TABLE `SSLList` DISABLE KEYS */;
INSERT INTO `SSLList` VALUES (972677135296086017,972473959082360833,'68b21ac5-8fb9-43af-9fc2-950a185eb8a4',1,'Sectigo','PositiveSSL 证书','DV',1,1,'PositiveSSL',29,0,'','1年',1,199,1,87,NULL),(972677135325446145,972474337089781761,'60f60ccc-ff69-471d-a813-cbf75e0d3425',1,'Sectigo','PositiveSSL 通配符SSL证书','DV',3,1,'WildPositiveSSL',688,0,'','1年',1,2000,1,2064,NULL),(972677135342223361,972474337099784193,'bd5fd8d1-b34f-433d-a199-cdeca5601aac',1,'Sectigo','PositiveSSL 多域名SSL证书','DV',2,3,'SanPositiveSSL',299,100,'','1年',1,1500,1,897,300),(972677135350611969,972474337104293889,'cb353de9-6f14-4998-8e00-6271bea0c717',1,'Sectigo','PositiveSSL 多域名通配符SSL证书','DV',3,2,'SanWildPositiveSSL',999,499,'','1年',1,3500,1,2997,1497),(972677135363194881,972474418523697153,'0c40fcda-1ca2-45b0-8511-84785a24b244',1,'Sectigo','企业 OV SSL证书','OV',1,1,'SectigoOV',1000,0,'','1年',0,0,1,3000,NULL),(972677135371583489,972474418529435649,'60979f08-0c0e-4cb1-a72b-58500930bfb0',1,'Sectigo','企业 OV 多域名SSL证书','OV',2,3,'sanSectigoOV',2000,665,'','1年',0,0,1,6000,1995),(972677135379972097,972474418531966977,'7caab045-8b3e-453c-8b24-331e3ed2132b',1,'Sectigo','企业 OV 通配符SSL证书','OV',3,1,'wildSectigoOV',2800,0,'','1年',0,0,1,8400,NULL),(972677135396749313,972474418538606593,'7894669c-409c-4e4c-b1ae-3c405aa535f8',1,'Sectigo','企业 OV 多域名通配符SSL证书','OV',3,2,'SanwildSectigoOV',7500,3750,'','1年',0,0,1,22500,11250),(972677135405137921,972474418543550465,'8076fa5b-8b32-442d-83ea-cb4e0452ff2c',1,'Sectigo','PositiveSSL EV 证书','EV',1,1,'PositiveSSLEV',1200,0,'','1年',0,0,1,3600,NULL),(972677135413526529,972474418546782209,'39705d1d-4935-4bea-80c3-3f17c08a8ef1',1,'Sectigo','PositiveSSL EV 多域名SSL证书','EV',2,3,'PositiveSSLEVMD',2500,835,'','1年',0,100,1,7500,2505);
/*!40000 ALTER TABLE `SSLList` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `SSLService`
--

DROP TABLE IF EXISTS `SSLService`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `SSLService` (
  `id` bigint NOT NULL,
  `ssl_id` bigint DEFAULT NULL,
  `admin_info` varchar(45) DEFAULT NULL,
  `com_status` bigint DEFAULT NULL,
  `dns_host` varchar(45) DEFAULT NULL,
  `dns_type` varchar(45) DEFAULT NULL,
  `dns_value` varchar(255) DEFAULT NULL,
  `domain_list` varchar(45) DEFAULT NULL,
  `domain_num` varchar(45) DEFAULT NULL,
  `domain_type` varchar(45) DEFAULT NULL,
  `file_name` varchar(45) DEFAULT NULL,
  `file_value` varchar(255) DEFAULT NULL,
  `order_start` bigint DEFAULT NULL,
  `org_info` varchar(45) DEFAULT NULL,
  `p_method` varchar(45) DEFAULT NULL,
  `p_type_name` varchar(45) DEFAULT NULL,
  `setup_server` bigint DEFAULT NULL,
  `ssl_code` varchar(45) DEFAULT NULL,
  `ssl_csr` varchar(2048) DEFAULT NULL,
  `ssl_key` varchar(2048) DEFAULT NULL,
  `ssl_pem` varchar(45) DEFAULT NULL,
  `ssl_type` varchar(45) DEFAULT NULL,
  `tech_info` varchar(45) DEFAULT NULL,
  `uuid` varchar(45) DEFAULT NULL,
  `xufei_orderid` varchar(45) DEFAULT NULL,
  `yn_prove` bigint DEFAULT NULL,
  `yn_replace` bigint DEFAULT NULL,
  `yn_xufei` varchar(45) DEFAULT NULL,
  `z_domain` varchar(45) DEFAULT NULL,
  `create` varchar(45) DEFAULT NULL,
  `ca_num` varchar(45) DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `agent` varchar(45) DEFAULT NULL,
  `order_id` varchar(45) DEFAULT NULL,
  `order_name` varchar(45) DEFAULT NULL,
  `z_domain_list` varchar(45) DEFAULT NULL,
  `ks_money` bigint DEFAULT NULL,
  `ed_time` varchar(45) DEFAULT NULL,
  `date_diff` varchar(45) DEFAULT NULL,
  `img` varchar(45) DEFAULT NULL,
  `ssl_name` varchar(45) DEFAULT NULL,
  `zyks_money` bigint DEFAULT NULL,
  `ssl_type_detail` varchar(45) DEFAULT NULL,
  `setup_service_detail` varchar(45) DEFAULT NULL,
  `admin_name` varchar(45) DEFAULT NULL,
  `admin_tel` varchar(45) DEFAULT NULL,
  `admin_email` varchar(45) DEFAULT NULL,
  `admin_job` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sslid2_idx` (`ssl_id`),
  CONSTRAINT `fk_sslid2` FOREIGN KEY (`ssl_id`) REFERENCES `SSLList` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `SSLService`
--

LOCK TABLES `SSLService` WRITE;
/*!40000 ALTER TABLE `SSLService` DISABLE KEYS */;
INSERT INTO `SSLService` VALUES (974966692268544001,972677135296086017,'方;逸文;13327806324;529187451@qq.com;工程师',3,'_68cf6bd25910a65d2595f20e9dd99739','CNAME','475bd532dc57a9c606dc007f023e806c.16bee2646e7afb331881720f4d8b4126.ueqf4vnk9jv1or.sectigo.com','','1','1','68CF6BD25910A65D2595F20E9DD99739.txt','475bd532dc57a9c606dc007f023e806c16bee2646e7afb331881720f4d8b4126\nsectigo.com\nueqf4vnk9jv1or',4,'','dns','Sectigo',0,'PositiveSSL','-----BEGIN CERTIFICATE REQUEST-----\nMIIC5TCCAc0CAQAwdzEWMBQGA1UEAwwNd3d3LmJhaWR1LmNvbTEWMBQGA1UECgwN\nd3d3LmJhaWR1LmNvbTEWMBQGA1UECwwNd3d3LmJhaWR1LmNvbTELMAkGA1UEBhMC\nQ04xDzANBgNVBAcMBuWNl+S6rDEPMA0GA1UECAwG5rGf6IuPMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhzYDseWRmnUHndk67oevFjsgzuBlvcN43Y26\nyQHMxGAztj4A+Ss6S38n3OsTNifxBoptvaKPh6nleMDKmtdhC+P9Fwe1SdkG/JrZ\nt/qHazx4J+D0T2CY6VOSm/RPhALm+zSwZgjXorMsn7kO2NOMW+Yu+0PYr3Jg5Lm0\nUu9YKiR5VMv6hIjPB1zuLe+uTs2AV7tgSR9ruxuGyg8I2barbMKQRSPqIFii6mrw\nUvt1RnvBVKw12A+vby6oikCjz2BVYEjVStTnqtr5B1ZtoRfdq+FWlZxGivabME56\n9LGJbP709eOncoDRqw9Rj/b5z8wdwRJ/l3fNhWVG4Gs2tWruZQIDAQABoCkwJwYJ\nKoZIhvcNAQkOMRowGDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkqhkiG9w0B\nAQsFAAOCAQEAJ/Thf9N8si+CCAKnEBDXSLdYRSMqVL5zDXGgYdaGTovKBO2wmQnL\nWF9wosEL3TStJ5p2fj+p8wayHYgIcLsT3RYr4iXKdAgOA85+rcuPCvP3XswXYCMs\nuGy9xLN2QnrGYOLd+RKpP+R5/HNUqkGibyoLtHgCrzJeOPr/+rVBtv0r+Fdgg35i\nFHrzGgNpVZrnoiaJWEgRgRi8K+vItP/xXThdB5uVFn1cfk3dog8Xhb0eIJQcgGBo\nNQ2ngiMb49lp8qFv5C5oKwrkBaIjthJKVQ91KT3dpbEaqi76kX7hRgcDbrAKNlNu\neMysRRKsxid4NnV2MrHEeHxoaRN7dyKhAA==\n-----END CERTIFICATE REQUEST-----','-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCHNgOx5ZGadQed\n2Truh68WOyDO4GW9w3jdjbrJAczEYDO2PgD5KzpLfyfc6xM2J/EGim29oo+HqeV4\nwMqa12EL4/0XB7VJ2Qb8mtm3+odrPHgn4PRPYJjpU5Kb9E+EAub7NLBmCNeisyyf\nuQ7Y04xb5i77Q9ivcmDkubRS71gqJHlUy/qEiM8HXO4t765OzYBXu2BJH2u7G4bK\nDwjZtqtswpBFI+ogWKLqavBS+3VGe8FUrDXYD69vLqiKQKPPYFVgSNVK1Oeq2vkH\nVm2hF92r4VaVnEaK9pswTnr0sYls/vT146dygNGrD1GP9vnPzB3BEn+Xd82FZUbg\naza1au5lAgMBAAECggEAA5J8HiRb0VZ9Ge9qZBP/Eu4DKj+uAea65o+TezLwELsH\nv9BMo7z556orzohfZWH6lGC/4jVwXDvtZ47DS43Knawraw61n0CS+WhtMe47+auu\nCnKWe/vWn9TurzQcpPbw2MEXbRH06qzGn8PXsR7P3fWxP7P/uKRd3YRj8q8p8DpT\nipiUH+Sqa99eTuk6wAz0CEJiaTC7n3NfBMMxgf1s8MAUwL9APPSHN1o64aciPWMB\nogGxn9sCkLm07ieyuwmxgifw7884s71MKuuxDNFZD9XkiPakJe0eIuS59Ooxv1yj\nAFpH4KlYZCBZcWD2bPD0n/cwEQBXnuxzupipFs6iWQKBgQC72ZC8AtfO86cHlVbP\nNM0MuHejOn1Ggi/UhfItly+bViOFIBjrFRFxYMHpogV2XIcNAP5CmVtkFMg2Fjfu\nM/HMyMtZp2ABnnWyR/gRxqDsIHJPIr/wKwXwB3dLDUAOhwcip584Eu+r7vQbL1oK\nFs9NdBoFM2S7SD+/Ee5iZWJe7wKBgQC4Q6XjHiMRr5rpPjTkgPpWTW1VdkG7mgHH\nKSU0LD0UY/HqF2m+vO/2R6pc0AH5aB5NRY7txhjhIGbY8u7aI44DOWBC+wroIzf+\nUj9QaCf0TMOuLjDTSyJdXEzVLPcl/TD5KmerqlWAU1gVdbx/jprkU4D7Sldr605g\nMDpQuM7H6wKBgQCjvQfIqd28mWsdVBHf+RFGpkA6OlBaNj5EzCCFlsO7bcD2WD0A\nJFQY9JLx9/U074AUWKpIGO6tdOPzKTjFW+fHbbq6wgcQQjQrG0pjDDTWvY2F5y9Q\nw9+gEpHsDHeqcg5JKbzCHtdkkZubWpHsO2elBq9rmxmRmW9fN90HTYIKZwKBgQCt\nkDLaSpHX3+P33Xtu3TpwPkBFJr3l2rvgABamHSFvanD9Ag699jPFdF25TMj/Rx57\nxJBL0bOvopids75eRO1IhivxC4yLelkMuiYfM+ymhgQvthyReg0liYJprhAifWo6\n2MGMBkbBOEAY/qyOLiIh/bHZgX31DlDl86eKT3me/wKBgDKdVhm3UyiSRGcTl/+S\nshgsM5/V6xYLyCsM30HVntYmDjSpmE0r1uBT8r0qodcD9kFgk9Ak5JuxiDILIk50\nC5IU2JbfvMnwVHyIccAbhXH0q3vK6WymANGAVRfDNRXRd/jmsDg1Ca9NuQj9+3IZ\ndRoSWh8k9i77RsGN+dXlI3AN\n-----END PRIVATE KEY-----','','DV','','47a0a641-26c8-4e85-bf2e-32e2f946400b','',0,0,'0','www.baidu.com','2024-05-14 17:28:38','2192055950',957493600380858369,'','CS2405136FDY','已撤单','www.baidu.com',23,'-','0','/static/get_img/Sectigo.png','PositiveSSL 证书',69,'DV 域名型 SSL','无','方逸文','13327806324','529187451@qq.com','工程师'),(974966692286955521,972677135296086017,'方;逸文;13327806324;529187451@qq.com;工程师',3,'_a00226890bff1a8c869b03dc935714df','CNAME','aa7d5a197d78dbdcd7a0a976247ce47a.0da6ca1636257f17a9b8c37ecf4bfd23.k3gcmp8nskv3vd.sectigo.com','','1','1','A00226890BFF1A8C869B03DC935714DF.txt','aa7d5a197d78dbdcd7a0a976247ce47a0da6ca1636257f17a9b8c37ecf4bfd23\nsectigo.com\nk3gcmp8nskv3vd',4,'','dns','Sectigo',0,'PositiveSSL','-----BEGIN CERTIFICATE REQUEST-----\nMIIC5TCCAc0CAQAwdzEWMBQGA1UEAwwNd3d3LmJhaWR1LmNvbTEWMBQGA1UECgwN\nd3d3LmJhaWR1LmNvbTEWMBQGA1UECwwNd3d3LmJhaWR1LmNvbTELMAkGA1UEBhMC\nQ04xDzANBgNVBAcMBuWNl+S6rDEPMA0GA1UECAwG5rGf6IuPMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsqRljOe5zdh5wTgMOgTs1lL0ROSXF1O0grJ7\nYwDuprn54UwY1Dmpc1wzcM00wslAXJkaAv/8SDG/EZjwwDu/Pd8JFa+8Sldad7+V\n+LEdcuAe6jmrTA60DTIC6hq3Ti8gNePibflyYKNJdyWYhhywlO8OwbX9W+6kua1h\ntvmN0wn4wHmV/ajxz9Ppb++95rOXrw3ANlGr1gKRyrzC2+N9wufb7xuq8Ytv4o7q\n756r4sXgtvFMr1yQp5nw9aJY+rAYUEFfjk354pnhTAlLKcaMq89W1VQHf9vnUqHB\nEV4k0DOSFfytftRRJ9fosgeb+RTyrwt8NiMwnqfhiTtLwRjD4QIDAQABoCkwJwYJ\nKoZIhvcNAQkOMRowGDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkqhkiG9w0B\nAQsFAAOCAQEAESiXNmGVlgAZ4NbVO6ci+yMzlgWE+Q1oZRFAtFZ8p7fHNQChFGR2\nAV/gswy64xGf3V8sfAcmHih5JsERjgJzrHlxr1yIb7fSwii5cfKqyR23cxEnHMTa\npV6luVIjdsueI3mYgIDr7gKoXlbbbTqhqJggtPmYqjfX9MyjAd2rLk65EUOyyHSC\nt5Pa2I7WPN9oe5/Qw2OztjMPJqoqDNNHfCVJHz7Rq0FrYGf5DVDU0pngoq/5Uuvv\njoK7u4Dspa8G01Rbe+oL8NwBscwFNDJz3CLrCaFgHj8L2Y7DGhpNsTIgAZWHnHF/\nvGrSf04REOYZMs6p3w56dNwRxdw0mZ3szA==\n-----END CERTIFICATE REQUEST-----','-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCypGWM57nN2HnB\nOAw6BOzWUvRE5JcXU7SCsntjAO6mufnhTBjUOalzXDNwzTTCyUBcmRoC//xIMb8R\nmPDAO7893wkVr7xKV1p3v5X4sR1y4B7qOatMDrQNMgLqGrdOLyA14+Jt+XJgo0l3\nJZiGHLCU7w7Btf1b7qS5rWG2+Y3TCfjAeZX9qPHP0+lv773ms5evDcA2UavWApHK\nvMLb433C59vvG6rxi2/ijurvnqvixeC28UyvXJCnmfD1olj6sBhQQV+OTfnimeFM\nCUspxoyrz1bVVAd/2+dSocERXiTQM5IV/K1+1FEn1+iyB5v5FPKvC3w2IzCep+GJ\nO0vBGMPhAgMBAAECggEANw6456URnHIwEE98YTNcZS/a7hK/yGY5d6CXq6tBYmxm\nKbxg2KPO6GmNxyHPtYr+RavAGxWXRQ7j8wH+jVi8t85yR9dKQmVK0iZfi+7Wgy21\nCX5tn3rWnkyAHnPOdA/NLx7Fjhx31nXs8gZJzGeyOqy1ERnMdNWHulxIBfNC9TwT\n/pnVSnQQPfHneJYJGLl++7QqEiK0zBZJWmehL/kwtI5S3//5fpnt3IipL6Kt29TG\nR1DdtUROyuRH5gXTHb0OlicL+MEgE5GEta7HdHDF3uJLqXUeCvN440n4ACWchNJC\nJM5BQOp64G81IwxkfB6iVknAp3AmzI0IhA23TtNHewKBgQC6qO9L3JYnnhaz1jw7\nltquJARkViS5T5wsZV+5rz+BZFnPhL1b2750feLucZ4Sb/m/00VBaagsX3v+GwAu\nwxCQ5mad/7GhTZXxFsDCWpbiiJbnSortDfihSoh5S6yGjYZAE/RLV0iePO/s4y6T\nl1oFGVDMy+DgUVMBfSw6uRCsDwKBgQD1AP0Wa88H8WPqD+622xqgGzDuPW9rCLPU\ndSY828Y6sRBIf/RYpZpucACVd1TBmAazcKh51yQYWRY589bx5G1uFngeEzHDsHqs\nUPHnijdy7t4WR+M2UIkRIeyc3yrlRl/X5cBkpi1V8NFQnT2VSYB+glSJxS+vmM1f\n1w1ocENhDwKBgQCxN98qOV5ectu/FnHNaaZpm0yILrljL1BpXj4KP3ad1LboLXDT\nG8ixNSwVks2vV5zZIS8psbS3nFhOozgDLbQ2vmlDLRZqvQzU/vM4E2YZ4fDT25QX\nCSK7j9YOw2gjv45snAVe2pnZkqWAVRlNi28uVCDAr/jrQIr8aOlrP3WyxQKBgQC1\nzHQATSMkq7QOHKBvwwlwwHyAa8nfTGfUQNGM2y+uoXCJu5ieptN/G1Dzl8ammJ5w\nTCpXrpK0FnbtrchTWcLxq26aMnf801lTwASgYOlgOWtGcTViOkRbGCcxCFAWLTA1\nTeCDTuPZH4bSILtvqUjVlvQP/uexjtrmQPdMBfGZdQKBgFBi8cEKqGpOgG0vw30F\nMENgzWq5t4fvG4jYZFHxkRN178/TlCFnuAhplil2Y899T6nrBjSzNFZ6Vk3oTuRO\nI+M3/oeTEd0i7b3xYAIMbURsR1fobrsA9Y/bZ96/riuzUM1qMjoaaO9VtoypeXyY\nWIrsX9xSr97ZWSvCgRRhfYzg\n-----END PRIVATE KEY-----','','DV','','ea2d0f1e-559f-4d00-8d65-1161a02b768e','',0,0,'0','www.baidu.com','2024-05-14 17:28:38','2192055950',957493600380858369,'','CS240514QZ69','已撤单','www.baidu.com',23,'-','0','/static/get_img/Sectigo.png','PositiveSSL 证书',69,'DV 域名型 SSL','无','方逸文','13327806324','529187451@qq.com','工程师'),(974966692297646081,972677135296086017,'方;逸文;13327806324;529187451@qq.com;工程师',3,'_c5e1249a0add851284c9b12ea2c79c3c','CNAME','3effece76da1415843ce8c8f3a110056.4d20c692d8bf4698b266e772b7195f73.2ufme6y1sro4b2.sectigo.com','','1','1','C5E1249A0ADD851284C9B12EA2C79C3C.txt','3effece76da1415843ce8c8f3a1100564d20c692d8bf4698b266e772b7195f73\nsectigo.com\n2ufme6y1sro4b2',4,'','dns','Sectigo',0,'PositiveSSL','-----BEGIN CERTIFICATE REQUEST-----\nMIIC5TCCAc0CAQAwdzEWMBQGA1UEAwwNd3d3LmJhaWR1LmNvbTEWMBQGA1UECgwN\nd3d3LmJhaWR1LmNvbTEWMBQGA1UECwwNd3d3LmJhaWR1LmNvbTELMAkGA1UEBhMC\nQ04xDzANBgNVBAcMBuWNl+S6rDEPMA0GA1UECAwG5rGf6IuPMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt8v5oJ2y3jG2jNA9H+dI7igDBV+QOr8BjgZA\nb1i7jQeACUH+if9Lyl3FBFc53t62oOFBnQZIT4k38Zo3Z4lOgzkaYXZluqC25566\n81/zZkK3Eezr9ZmlhhzbBgPLN3HTdIf97V1+Kb4nBt/JlnhtNa/wjQ1L5pv1ZTQA\nmgkpxG1zhSgLe+AD8VOaZLAWPSEnaEPaKdg2Cc5a1riDZOoy/EQKz/ezlXklnf0j\n5hce9wZiU50tBC/mirZ73efKrGuhDMEHfxHMNYqBAFnqwf4rd4/x3HFQkknMLB5X\nJBX+ebIPkI4DAydCseo4xToHg6qSEEnaypWvxOQXmwfa/VxXbwIDAQABoCkwJwYJ\nKoZIhvcNAQkOMRowGDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkqhkiG9w0B\nAQsFAAOCAQEAQv58qHW3Dj+RtYfKCz+BZt2HtKRqHTshwXkC4Clt2VICgBZlYVHv\nbdyXf9YojQcbYRjofP0ShV6Y+lzvibcXqIFSqekzQYRo68+v/cpI0D4jyZQ9Icge\nuIJRBKdNdYhDzGXaf7VCvNXuoKq+FCev/D1ZmfO9JFTr61ZwUJC2kJFOrlcddG4p\nenfXmdgk6vkL0fRfR8tMYe1X8A8B3to07KkfLrn8sQxRLRLg/Fx4VOwqPX0teLV0\nFStnVD8RT9ozJgWcdYdqPOMgvTd0iVIGf3VmiHuERWXRWn5FyyvRaqtQ85E4Fqmz\nMg328XERPqw6dPYkPyPgOSIYvc9grRixVg==\n-----END CERTIFICATE REQUEST-----','-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC3y/mgnbLeMbaM\n0D0f50juKAMFX5A6vwGOBkBvWLuNB4AJQf6J/0vKXcUEVzne3rag4UGdBkhPiTfx\nmjdniU6DORphdmW6oLbnnrrzX/NmQrcR7Ov1maWGHNsGA8s3cdN0h/3tXX4pvicG\n38mWeG01r/CNDUvmm/VlNACaCSnEbXOFKAt74APxU5pksBY9ISdoQ9op2DYJzlrW\nuINk6jL8RArP97OVeSWd/SPmFx73BmJTnS0EL+aKtnvd58qsa6EMwQd/Ecw1ioEA\nWerB/it3j/HccVCSScwsHlckFf55sg+QjgMDJ0Kx6jjFOgeDqpIQSdrKla/E5Beb\nB9r9XFdvAgMBAAECggEACuTbsHH4cqQWsmT7YoM2dutHNo/SADFnGbB0LEVhpHvP\n0nynZi2MI/wH2/iBlRy1TXbciT6glZpIBxmBdDEWxr9hOWqnn/a7sKHfaeKDjkmh\nCT52Q/Ht9jOKpeIervQfuPYAjouumTHMpWMpUzJjIr4Bb5Lcr1gKJQv+hyLI53wN\nib1NOlBxoS14sKBMGweHKbaNQiEJ6MTqL4aSyYLPQGeboo6mFsODSoOYxnef+7PO\nxpYuLiZ60CVHEMnexj3BPHSGsWofqCdpeG78MdZbHjGd5H1Z1zIV4W1Vy6exgFZk\nvQ2vBsJkWvuTvG3PuusZGGyv6fR26+g3fDHR2yY1cQKBgQDn7HZhGKUN+I/QlMfa\n8KMt9sA7QbIt9P/S5igaYo9TC+6FghE8l/O9XXe7d6bkJ/G9FwGVAXRBUkWSs38a\nMWka6QSVZlw8y395TuM6OcvUxlLLlVb6kKwqIMo5LGUyDtK3L9srIiZLuEdZ0Q95\nZDSaZ9p0PCg+8KzDJ9uXdtuxEQKBgQDK4IF+IEtoRbuXlnDM21xiWFM3hrEEsy7i\nVFtnUredqMXRFPD/6NIlyYH+r5cMwQ6d5UGeYafYDmwFIxkTEGCAyaEm8Ch1EVVL\nYr7YSH+DkafoF1+DJZIYTmdQJz2GIUQKOGdx7kng+MivTjgvxIasQ3QGQ0Xwp/xz\n5DWV8GeAfwKBgCKn2fSeWyOgTdUsNkUICko3AfiY04g8gPHBxKJUsfgF42l9yxeT\n2S8lxVt1ALI9KclIo8rhQb71DUJjog2G7p6/zRiKCCRpgC3fVOMBeezcici0EyAz\n+BX/elhawvBc+K247/YdDz3nVnocXESWbanY9PLDfnwROK+cyBp0/1URAoGAWB3V\nAKQkPID1Lq0QB1pLsQdt9ZMR9dmhxnofTSxXRioU6XAwEObdx73TMXywncCIiboW\nmyUNUsFI155b+LTHRYjN7uymldpWcqzL+YLbp5ivFhZ4zKyX3OjI5L52cXinZc0D\n2S9HdWh5OfM+Yjj4yX9uuH2lthYwh2GaiKpnOh8CgYEAiryk5ygc13uoSlYwu1Le\nNTOwzdivThy9w/3pyvj3jZ0NkqF4+ZyMpTeMDNxgUybvXgap+bRXQod0fb8bt4qg\nXWC5Uc4UUJqqCVv0xrHKeEub5ifl9ceVbL6YRWoxbrTTUYP6stbL/uz+MbNEfeJO\nEsrnbswmusw5R3UL7nSLPls=\n-----END PRIVATE KEY-----','','DV','','a207317e-05be-403a-b6fb-046be0cd6101','',0,0,'0','www.baidu.com','2024-05-14 17:28:38','2192055950',957493600380858369,'','CS240514CEYD','已撤单','www.baidu.com',23,'-','0','/static/get_img/Sectigo.png','PositiveSSL 证书',69,'DV 域名型 SSL','无','方逸文','13327806324','529187451@qq.com','工程师'),(974966692300881921,972677135296086017,'方;逸文;13327806324;529187451@qq.com;工程师',3,'_64368906469b89932f0660a14541636e','CNAME','35dbe9b90f2fe0751f4e09cdbe104aff.05a5d13fb315dae1156e84d8488b6638.hsf79zzehnm4v7.sectigo.com','','1','1','64368906469B89932F0660A14541636E.txt','35dbe9b90f2fe0751f4e09cdbe104aff05a5d13fb315dae1156e84d8488b6638\nsectigo.com\nhsf79zzehnm4v7',4,'','dns','Sectigo',0,'PositiveSSL','-----BEGIN CERTIFICATE REQUEST-----\nMIIC5TCCAc0CAQAwdzEWMBQGA1UEAwwNd3d3LmJhaWR1LmNvbTEWMBQGA1UECgwN\nd3d3LmJhaWR1LmNvbTEWMBQGA1UECwwNd3d3LmJhaWR1LmNvbTELMAkGA1UEBhMC\nQ04xDzANBgNVBAcMBuWNl+S6rDEPMA0GA1UECAwG5rGf6IuPMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEArrOeMyVfqtRjk4yWSdpP8z4Ua9HLYjClXfyV\n6fnr0HG/NHzK2XtAZEwU9HvZ7ENwd/CGZ8RnVw0Zz4OwCSAFUKPveKtgGncCN/fr\nlTyY9H00t0pb2RuPGhITsWYHGSs+WlNy2pbkwSRuIHO0o5uK/WBhZvG+u0B16vrt\nHw+hWZvATEN3weARbekN3VBH7ECxBO+zsC3yXVzDthwsjHi/0+CiV9cjmBMZRiGD\ntAWOjhww9NGs6l3bBf4g3FXd6xwh1/ZEdyNHpD3cW77rKN0s9LXbM5hUNWsTXSy/\ntCSagWBx4DO1alZ7KF7IecK62aW0TVZ+X7zEsf1oLJiAJsHtFQIDAQABoCkwJwYJ\nKoZIhvcNAQkOMRowGDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkqhkiG9w0B\nAQsFAAOCAQEAXEC+uLQ8lmkQY4irS97WEEFDienljsmbvtaJ3fFJvBbn86uFC/Hk\nnIjJFuarxU9F1zH8rBfNiCX5z1RyJFHoeh9eGXBnLqb0C7BwgCo7C6RTboZX2twK\nG4k9YuO1jk53icwaVLUHKrtoBgvO/3F91nGevfjsaWiFvBY4hBbAydqYcCi1h+kT\nORZaKASXUVr6zYPkRd4fQhcrQ9Xw/k0ba1kkSLRx/ZrQOMkVoU24xVht91cStHtB\nbaUZ76b1BOyE9R8tbjSABfnC7mA/JauIjPiGVF2ZdixCeX3hil1EDlSrP9m3CvOT\nfPsrE3oGIySIb1M5oUshghRzalRJI74zZA==\n-----END CERTIFICATE REQUEST-----','-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCus54zJV+q1GOT\njJZJ2k/zPhRr0ctiMKVd/JXp+evQcb80fMrZe0BkTBT0e9nsQ3B38IZnxGdXDRnP\ng7AJIAVQo+94q2AadwI39+uVPJj0fTS3SlvZG48aEhOxZgcZKz5aU3LaluTBJG4g\nc7Sjm4r9YGFm8b67QHXq+u0fD6FZm8BMQ3fB4BFt6Q3dUEfsQLEE77OwLfJdXMO2\nHCyMeL/T4KJX1yOYExlGIYO0BY6OHDD00azqXdsF/iDcVd3rHCHX9kR3I0ekPdxb\nvuso3Sz0tdszmFQ1axNdLL+0JJqBYHHgM7VqVnsoXsh5wrrZpbRNVn5fvMSx/Wgs\nmIAmwe0VAgMBAAECggEAGpjS8XNzrKTYIhsJuGWpvV0Kq92NSr4gNv5g7nWznpDq\nrLmBAno+s+ZXXF9MrqkXwtLWC489smWWs96XGtSKqpvCtgrVP0I2DsJKVfabTMGR\nGISUFioZb70nzV68nurT+9lvh3LokbUTpRIogq4IC4o0WgFc6dInAabBGfT9ug/5\nOuPlpfmoh5Q6VuE7eXTARVbZ92ROCP+eLSl6Xd1/sXitMM7DlkKqhoXJ35aDdF4d\nJ5FwT1q/JABGkKdGFBI08I5JYBmR0g/dXCarN8LYtyzOWKMifBGdwlEd7Poc0O+v\nWF1gCqvZb77Fhz7yMm/dCCgx9J/O4xZcAl+kwx3fsQKBgQDFTXGhu1x1Vp0907Lt\ntRZcWKa5I5B0ik3iLL+Aoids4Aifoc+gaRBsVjM1Ou5PAaO5DrnuJWUQAQkRLCxK\nzjcav4XQix16bZCq3NnFGUIAX3MRLcrmaWqJoN/RG2kRliiQZaxnLHfmz/qjENN+\njakKS41TO3TRs+Qt91NgUJh7PQKBgQDirOQq4PtXZi/i5Kpqeke6i3CmiiUlUroh\nCELb8BGGS12bMgI3gWqlaigAJ+zsuoZCQbOHe3lMxuf8qizqw37sKTNm6pmeKrCL\nLo9djC+D0Zt5fAik0I3DOTFUtkN81TDrmW14f68v124bbbhoC9dR77D4a8hJ8wVa\nXTPNSfo2uQKBgC5Ztywre9kE9AOmTA6T0CD5opEzzafTJeDQalypu2FUcDIwLTqd\n9AG/bmA+6+UbmX7L+tn7ZvaPh7XFDtK0BzX9cE5BoDc2rhm1sGMhP+QeiEWs4FGo\nyiJy7KGPMuuyvuzNNec3ByPuEA9m0IiWdvO4NActnI/3Fs/Fnoyt4mRxAoGAPZ7y\niBWRPjkxW2T+TuPfC/9MKHn1jyfmIYS1aETi/rcOQ1pc7+nmrAEnzqML7W7ngKuL\nupD+cLjB6BX7D/cFUf4BU0EDcdP5Gl0V7EBnylX5BhWJ5AQuBZxOpNDHl+/81I9p\nCxDq9v6BmENDlkVhy4x1d5MVZ8UA94fOxUgPFUECgYBY8J+b+B2+KLpvD/FHY0v9\nP8oCHdocBaTzj+8dVJK8Hnph/7C00hLFQKduSUoWHwEtjmVsWxF82j9S5xvww58x\n3wc9ovCnPA69PC7B7R3LSRYyDndbfMowuYvPTZbt5o2wKgDH637too9AF0decGM9\nL3r3pq61kabxSP+Iwi9Iiw==\n-----END PRIVATE KEY-----','','DV','','34dae912-724f-418a-ae43-8019296f50d6','',0,0,'0','www.baidu.com','2024-05-14 17:28:38','2192055950',957493600380858369,'','CS2405148E3W','已撤单','www.baidu.com',23,'-','0','/static/get_img/Sectigo.png','PositiveSSL 证书',69,'DV 域名型 SSL','无','方逸文','13327806324','529187451@qq.com','工程师'),(974966692306350081,972677135296086017,'方;逸文;13327806324;529187451@qq.com;工程师',3,'_bf54cb9a709bdf8d74df22446f79d9a6','CNAME','c9818032b6710cd2662ab2683a146d60.dd2de1c3734fbd1ea9f82453ae4d578e.isszonsp77p3i7.sectigo.com','','1','1','BF54CB9A709BDF8D74DF22446F79D9A6.txt','c9818032b6710cd2662ab2683a146d60dd2de1c3734fbd1ea9f82453ae4d578e\nsectigo.com\nisszonsp77p3i7',4,'','dns','Sectigo',0,'PositiveSSL','-----BEGIN CERTIFICATE REQUEST-----\nMIIC5TCCAc0CAQAwdzEWMBQGA1UEAwwNd3d3LmJhaWR1LmNvbTEWMBQGA1UECgwN\nd3d3LmJhaWR1LmNvbTEWMBQGA1UECwwNd3d3LmJhaWR1LmNvbTELMAkGA1UEBhMC\nQ04xDzANBgNVBAcMBuWNl+S6rDEPMA0GA1UECAwG5rGf6IuPMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs/I5HFsAt75awnLAsw/Ey+0EZ1Ff6vibzzb+\nSGwaWvLlMDcj6mbNWa7lBAjkoT/Zx5gLcJ64rrOViIPFP0rfSqgyKxQk7Od4gtam\nFF/JRM9Lq69j3qalgooKwA+/NbrlKC99SEWutqTzZsk/pa+tkwVA/j49k+O/0f98\nwu54NAMNhwxqq3JL+TYB4OUmEDGvzj6Irvpce8DswW/gD+fFuyYk4Bxrbog0AxTC\nmL/a6CAJfKty3xptepevIbF7lBMR9eIJ+Ea/ySwT2z3K6cxiZbXd9g8mRGZO76DW\nnVF6isYjjRxmjBuAyydM3IxlAg2k0CjqRuPPMLS6gnvYouJ5UwIDAQABoCkwJwYJ\nKoZIhvcNAQkOMRowGDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkqhkiG9w0B\nAQsFAAOCAQEAWnnlj2tZvmR8eY9Qx2G5AlsTJhaqV7HphNigZjB9epN1y+poq5r2\nzOhknMzAoCIIrypwQTN1oXRH6iHEwaVpzj9TnHkSt6ZE/hhoRYOJvZLQoBZjlLPo\n4f5O0QXDWjCkdC+3KJeYlCQpu50/Mu6mjG6jwTY6qnjrkKo+QXf1rbPlc6WrIUSf\nmLwxO8lo5/7NnXHP1Sr+uSfB5x4fXgYJZX1snHZAIZVoTFOGCd40UI/3pmWNodFg\nSH5GOMC5AP4b9uO3TJLGSQFN6IDRcuFSKAAw4n+zIKNPS9UpzKCFmw067ALJYrje\n4xkUhARwNP9co9YXrg2l3vZSXIbv5dj+ig==\n-----END CERTIFICATE REQUEST-----','-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCz8jkcWwC3vlrC\ncsCzD8TL7QRnUV/q+JvPNv5IbBpa8uUwNyPqZs1ZruUECOShP9nHmAtwnrius5WI\ng8U/St9KqDIrFCTs53iC1qYUX8lEz0urr2PepqWCigrAD781uuUoL31IRa62pPNm\nyT+lr62TBUD+Pj2T47/R/3zC7ng0Aw2HDGqrckv5NgHg5SYQMa/OPoiu+lx7wOzB\nb+AP58W7JiTgHGtuiDQDFMKYv9roIAl8q3LfGm16l68hsXuUExH14gn4Rr/JLBPb\nPcrpzGJltd32DyZEZk7voNadUXqKxiONHGaMG4DLJ0zcjGUCDaTQKOpG488wtLqC\ne9ii4nlTAgMBAAECggEAEMjaPOPcKEdwLc0IXmttHtxn2b9EZCV2Wxi1FUpIFw2r\nBRlPl1CpfOEMaZjuwn+zL6PmBUY40crcetRXtLSmKU+RBmWloBdjLUkaSrI4ktuH\nImeEuELxE0EPCaUuX7OkJmBvnispxc2TCpZhAnJljV4jFtP1aHMI9GyRVZlqyucJ\nuw5HV9r74gWLeKiDvNU/FF8GPoWVNnnx3WdFiwcwQguXs4CgyAvxB+Oxwsn2BYd4\nf+SvdnTv7pblr5pT/6wvH/aw3maHovjUbgQ7QXib4p91T+DnM/jm0rKed/C5DO9Z\ntovkyNVxvudRAWiPr8eoHi6uiZqlxuRWP2Zd74lU6QKBgQDuofyHVywdS8UOMtRV\nYh6+Hd0yJPxlgSrqdcFr1qJnJ0s76FktSMqvHa5RBF5M0+qrfqWirhIiz+uzvb8G\ngLnFjdsXraFNRbjzKcwA/OlV5lrO/2wNLUxzfh3Kauy5mlw5vdFMljcmuWWfD3l7\nGjck18jUXkLBbu+92/+tG56QnwKBgQDBCtXphYJJxbZoIlHLDcQMt16XF4v76wAy\ne4rRbneca+joa1+6lJpWoMYQ0j8Y7P11eHdUWjk7EOcvl19mCpazNegsVKNvhOVa\nnebThyEOHvrozGpEJWCDju8AiAanfdxWsl9F0XLWmZBvKcSZER18wJoums/mM5bK\nKBJUkqUWzQKBgQCokLhnigZZc2pxSyp+Xd9FgKb9gu7fbMy57hCBco4pu7IDUW+c\n82xpItg91o4eSxvbpIywS7H38VOYFhLYLxoLGNeI9Fea1AN/kSjQEddpnFzNF9Rl\nwnfcpVu53qn1ubon95PfO8Udj5zfLPktIimN1vfAPA1/wjOLs9n5mCb8FQKBgHLI\n96NQoMUvDCRla89akx2gl10uQi18cfSSo1ue3P6k56vz1MtD+3XRSZTM/3nZyWhl\nRkdhErC9f0o8Lxyk5TEQt8pMVZZFJn+4bn2O1yfnYfi1O+bjdr0ja5gcSiCj5TWk\nAKSqQye5zjlbU6xUkf0RJVza7w4J1PrF0wh43jNtAoGAS6oclx1paYfqlOspxIG/\ntcxqSd1VxP1MzYao+Ue5oqeo0Ii4DsE/aZdKgXJ6fNa63FYYHhllefiyaN5ITbjA\nUnPvoCpV5LpEstpICD6t2aHG0T1ZTorLAoKZfoxCevavZr24R+SMH/DXBqOLP0zN\nFo8p4BJz6A6eli9LhMCC7dw=\n-----END PRIVATE KEY-----','','DV','','60a179b6-1d22-48e5-9d6c-cba5fc18251c','',0,0,'0','www.baidu.com','2024-05-14 17:28:38','2192055950',957493600380858369,'','CS240514KSNM','已撤单','www.baidu.com',23,'-','0','/static/get_img/Sectigo.png','PositiveSSL 证书',69,'DV 域名型 SSL','无','方逸文','13327806324','529187451@qq.com','工程师'),(974966692314955777,972677135296086017,'方;逸文;13327806324;529187451@qq.com;工程师',3,'_46f8b5dc67761674b0c090276b25ddb1','CNAME','6b2df1af49de01ef781bc8089db6eae1.af1ab49b3103e47f27c8e2ae59a67b06.mdvcdo24k8lbjy.sectigo.com','','1','1','46F8B5DC67761674B0C090276B25DDB1.txt','6b2df1af49de01ef781bc8089db6eae1af1ab49b3103e47f27c8e2ae59a67b06\nsectigo.com\nmdvcdo24k8lbjy',0,'','dns','Sectigo',0,'PositiveSSL','-----BEGIN CERTIFICATE REQUEST-----\nMIIC5TCCAc0CAQAwdzEWMBQGA1UEAwwNd3d3LmJhaWR1LmNvbTEWMBQGA1UECgwN\nd3d3LmJhaWR1LmNvbTEWMBQGA1UECwwNd3d3LmJhaWR1LmNvbTELMAkGA1UEBhMC\nQ04xDzANBgNVBAcMBuWNl+S6rDEPMA0GA1UECAwG5rGf6IuPMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs5uu/MvaIBlD9CXKb+KyshktZMeyLO43pvu4\nyu/PfWMeEQxnqCXlRxzMCz3SDSD530/5B38RlMcqf/30M07zAD4diNSu1pMlrSQP\n34viDVyY7QDcsBZRB9rEK4ogQuV0AqJ5JqIR3KpbQw8PclYUDLQcXd7SivZgboRK\nRYzLyQ98XeifoOzvQZhlW4F6GxQ62E+/oIrEKK2Dh6IPX4sOaF0nGN0l/nWs9vdJ\nFJllEDQHR2Wd7T3EpS3rm599KDl3ZgLySzCkDA8xpilIF9vMLTFpNSuK4xykQrWO\nWVJF2mNUAu6r4ENlHlPLCmJPimYU9MU7MJpQ0EQbsSE0aBIzUQIDAQABoCkwJwYJ\nKoZIhvcNAQkOMRowGDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkqhkiG9w0B\nAQsFAAOCAQEAjMd0KBPI8X40B7GMlE5j8bG8RS/6xJLeJAWZo3zrQ7g0Vu4tk8qP\ngUuGY5RC5tS9PmsueerV+RVzpOJ07BA0Q+22gKlNnPue6bOR90biSUsb4reSLpuc\nCBSzP42cIjHMF88Jb3xJNUwDEa4nHxkbIc7EWiaqXwfbPoYYTeuzepYwozA5HsAU\nsFh4NBFuJMP7xBlXUw6tYe2yxCbucF+ToIu0f4dif6RRi61lhQJbemim/boNkh4t\ni0QZgD/0KZt1En1XcNk6tX+AoaRxfpSaP8rQbwT7HSZt9+NqWqf+yd9cP7/Pxjgr\n5HQjRcUgiozSLBT/BZxNy7iae8DzvEf4tg==\n-----END CERTIFICATE REQUEST-----','-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCzm678y9ogGUP0\nJcpv4rKyGS1kx7Is7jem+7jK7899Yx4RDGeoJeVHHMwLPdINIPnfT/kHfxGUxyp/\n/fQzTvMAPh2I1K7WkyWtJA/fi+INXJjtANywFlEH2sQriiBC5XQConkmohHcqltD\nDw9yVhQMtBxd3tKK9mBuhEpFjMvJD3xd6J+g7O9BmGVbgXobFDrYT7+gisQorYOH\nog9fiw5oXScY3SX+daz290kUmWUQNAdHZZ3tPcSlLeubn30oOXdmAvJLMKQMDzGm\nKUgX28wtMWk1K4rjHKRCtY5ZUkXaY1QC7qvgQ2UeU8sKYk+KZhT0xTswmlDQRBux\nITRoEjNRAgMBAAECggEAAgSKeXgPzgMCpqxDAzFEiSgrwKFWsNcHRIa463y4s8Oo\ni4/8XSD526vo5epuoOuhKv9SwKEOkUGQimoPs4DcTYElI0uiYDI8ft60F/KSGkzT\nQfxbvlWt2Nz+DbAqTx4VSM9K9aJl+XjneBAkAAiUBpWCaRKPAnUknbE6RWBAzNQt\nIL7ZXevkOY3dPstgzCAwi3vyhpwha3owGWQaG6u7wf4HBJg5nntZ4I4Q3mk3kXac\n2INfLEiluDmU3/zWCSbLqCTHosshrDA2afqBxGDQQ72SXRJ5+2VDHOklx6+3PGUS\ndpNZZDIsQubznR8F4/YMjbZzVGipM3CZRrKowTY1HQKBgQDmO8PKtXNPczJ2XPbn\nkoaS7J3tWuZ+tC1ZX2FREWJMSaL4/AjXV6jylAzzkB8IVMZaezitxqX0awxz02HW\nxjezcwODst0slQALXhjo+V7t41EXIxST05IeRfMhDxHK1hErz1KhvbupyG/GZzHx\nf7OZj/87hz7DqN31lEVOCc115QKBgQDHtX49kqH+k1gCECUCfnlt94T2rK7Hddqh\nC8IBDqux9cW85CfFD6TgqryjEHzNrKzDcRZb+RsSbEUSv7xLX+RzY5cb/SfBEWT/\nchPDYsjK+NSMnALz+c+23fvRv5rTH1XFOjG9RpUiJOBy1egggMPcv2/W2tKEKlZX\nCVKom6bw/QKBgHXj6G+l1fO60v4UwED3oH0fOAbP4vLCI+59joFKW4Egu5iuxYAu\n3M3JzY/yHlHPxvmR2wXwkA3FusycUp79RIxYulX3gpVPfoRTnIG/H0LEgUNNrT28\nujdSLhqhcsTgNE6wjGlRlARuI/393W8Bqt9ZAqJisFrT30PetlYtYs2ZAoGBAJFB\n3VSDtvgny7eUbpblzwTBKCETcPz7J85DcxS5ywBhtzWw7YxT2KThCtvAPkaK9g4h\nzxWOlrZLseH1O8vHL68OHPbqp2OydzeljeJrk4iufErZnvw+E2r7AxKIvuAYxpXx\nwgPdttNi5y4fj/s0LUH+rCYme0BOWwZrbYfdzwsxAoGBAL4k283cxY0nKSUv0VDl\n5vtF6NekIL/cKvKbrVOZEIouYwvNnEZZUjE6ICHLh0vuUMWWdxGGZyPXDo8vHkjZ\nzqv2bolqem/V3X9pCQL4qjPF2a60p/ZUEqMkt//yeqOz3Gcbzb+/LC+1kZhUFjS2\nOMfTTGWd0MRVMz4owbzYXxOd\n-----END PRIVATE KEY-----','','DV','','b4e94d82-aff8-49e9-aa2b-eaddc035f143','',0,0,'0','www.baidu.com','2024-05-14 17:28:38','2192055950',957493600380858369,'','CS240514J7FK','待验证','www.baidu.com',23,'-','0','/static/get_img/Sectigo.png','PositiveSSL 证书',69,'DV 域名型 SSL','无','方逸文','13327806324','529187451@qq.com','工程师');
/*!40000 ALTER TABLE `SSLService` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `id` bigint NOT NULL,
  `username` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `salt` varchar(45) DEFAULT NULL,
  `tel_num` varchar(45) NOT NULL,
  `email` varchar(45) DEFAULT NULL,
  `balance` bigint DEFAULT NULL,
  `create` int NOT NULL,
  `del_time` int DEFAULT NULL,
  `enabled` tinyint DEFAULT '0',
  `cert_type` varchar(45) DEFAULT NULL,
  `role` varchar(45) DEFAULT 'user',
  `agent_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_agent_idx` (`agent_id`),
  CONSTRAINT `fk_agent` FOREIGN KEY (`agent_id`) REFERENCES `Agent` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (954581118729203713,'aaa','511ca6f2bb2ef94b173995329a1f6ade','ymySYmcQIkWGKivvgpeaiobEpcwQjGZj','12312312345','qqq@qq.qqq',0,1710818673,0,0,'','user',NULL),(954713408296226817,'admin','c22e4065885f15796cbb4d975be4aed7','VfuQnEQUPjTPskJTbcHWUrSpoWPoxAKS','12312312345','qqq@qq.qqq',0,1710850214,0,0,'','admin',NULL),(956929351825580033,'bbb','99504b44401ad3136347a70a51326af8','tTpGqiGqdJaXTOqJLlAvlztYXLcIxqgy','12312312345','bbb@111.11',0,1711378536,0,0,'person','user',NULL),(957493600380858369,'ccc','8bd8c3334b5a3b79976a15e7328bc36e','gnXCBYhiAqILlViSImvNhllTioWnGYLo','12312312345','111@222.333',0,1711513063,0,0,'person','user',964522009411117057),(959984219784978433,'ddd','02f3a8db350de97a60c7505f54b7c92d','ZJznqnQcHReqxTdOSXpMYeXmoqbMMUuX','12312312345','111@222.com',0,1712106873,0,0,'','user',NULL);
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-05-21 18:32:52
