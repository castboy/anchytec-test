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

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for symbol
-- ----------------------------
DROP TABLE IF EXISTS `symbol`;
CREATE TABLE `symbol` (
  `id` int(11) NOT NULL,
  `index` int(11) NOT NULL,
  `symbol` varchar(255) NOT NULL,
  `source` varchar(255) DEFAULT NULL,
  `symbol_type` int(11) DEFAULT NULL,
  `digits` int(11) DEFAULT NULL,
  `point` decimal(28,2) NOT NULL,
  `multiply` decimal(28,2) NOT NULL,
  `contract_size` decimal(28,2) NOT NULL,
  `stops_level` int(11) DEFAULT NULL,
  `margin_initial` decimal(28,2) NOT NULL,
  `margin_divider` decimal(28,13) NOT NULL,
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
  PRIMARY KEY (`id`),
  UNIQUE KEY (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `symbol_session`;
CREATE TABLE `symbol_session` (
  `id` int(11) NOT NULL,
  `symbol` varchar(255) NOT NULL,
  `type` ENUM("quote", "trade") NOT NULL,
  `weekday` ENUM("0", "1", "2", "3", "4", "5", "6") NOT NULL, /* 0->sunday, ... */
  `time` varchar(255) NOT NULL, /* example: 00:00-20:55 */
  PRIMARY KEY (`id`),
  FOREIGN KEY (`id`) REFERENCES `symbol` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `holiday`;
CREATE TABLE `holiday` (
  `id` int(11) NOT NULL,
  `symbol` varchar(255) NOT NULL,
  `date` DATE NOT NULL,
  `time` varchar(255) NOT NULL, /* example: 00:00-20:55 */
  PRIMARY KEY (`id`),
  FOREIGN KEY (`id`) REFERENCES `symbol` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;

/* example */
insert into `symbol_session`(id, symbol, type, weekday, time) values(4, 'AUDCHF', 'quote', '0', '00:00-00:00');
insert into `holiday`(id, symbol, date, time) values(4, 'AUDCHF', '2019-05-21', '00:00-00:00');
