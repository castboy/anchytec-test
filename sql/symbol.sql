/*
 Navicat Premium Data Transfer

 Source Server         : mac mysql
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost:3306
 Source Schema         : symbol

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : 65001

 Date: 21/05/2019 17:18:11
*/

-- ----------------------------
-- Table structure for symbol
-- ----------------------------
DROP TABLE IF EXISTS `symbol`;
CREATE TABLE `symbol` (
  `id` int(11) NOT NULL,
  `index` int(11) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(255) NOT NULL,
  `source` varchar(255) DEFAULT NULL,
  `symbol_type` int(11) DEFAULT NULL,
  `digits` int(11) DEFAULT NULL,
  `point` decimal(28,2) NOT NULL,
  `multiply` decimal(28,2) NOT NULL,
  `contract_size` decimal(28,2) NOT NULL,
  `stops_level` int(11) DEFAULT NULL,
  `margin_initial` decimal(28,2) NOT NULL,
  `margin_divider` decimal(28,2) NOT NULL,
  `percentage` decimal(28,2) NOT NULL,
  `profit_mode` int(11) DEFAULT NULL,
  `profit_currency` varchar(255) DEFAULT NULL,
  `margin_mode` int(11) DEFAULT NULL,
  `margin_currency` varchar(255) DEFAULT NULL,
  `leverage` decimal(28,2) NOT NULL,
  `swap_type` int(11) DEFAULT NULL,
  `swap_long` decimal(28,2) NOT NULL,
  `swap_short` decimal(28,2) NOT NULL,
  `swap_3_day` int(11) DEFAULT NULL,
  PRIMARY KEY (`index`),
  UNIQUE KEY (`id`),
  UNIQUE KEY (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `session`;
CREATE TABLE `session` (
  `id` int(11) NOT NULL,
  `symbol` varchar(255) NOT NULL,
  `type` ENUM("quote", "trade") NOT NULL,
  `weekday` ENUM("0", "1", "2", "3", "4", "5", "6") NOT NULL COMMENT '0->Sunday, 1->Monday, ...',
  `time` varchar(255) NOT NULL COMMENT '00:00-20:55',
  KEY (`id`),
  KEY (`symbol`)
  /* FOREIGN KEY (`id`) REFERENCES `symbol` (`id`) ON DELETE CASCADE ON UPDATE CASCADE */
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `holiday`;
CREATE TABLE `holiday` (
  `id` int(11) NOT NULL,
  `symbol` varchar(255) NOT NULL,
  `date` DATE NOT NULL,
  `time` varchar(255) NOT NULL COMMENT '00:00-20:55',
  PRIMARY KEY (`id`)
  /* FOREIGN KEY (`id`) REFERENCES `symbol` (`id`) ON DELETE CASCADE ON UPDATE CASCADE */
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
