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

 Date: 23/05/2019 11:52:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `ticket` int(11) NOT NULL,
  `login` int(11) NOT NULL,
  `symbol` char(16) NOT NULL,
  `cmd` int(11) NOT NULL,
  `volume` decimal(28,2) NOT NULL,
  `open_price` decimal(28,2) NOT NULL,
  `close_price` decimal(28,2) NOT NULL,
  `sl` decimal(28,2) NOT NULL,
  `tp` decimal(28,2) NOT NULL,
  `conv_rate_open` decimal(28,2) NOT NULL COMMENT 'the basic currency rate of the instrument against the deposit currency by the moment of opening a position',
  `conv_rate_close` decimal(28,2) NOT NULL COMMENT 'the basic currency rate of the instrument against the deposit currency by the moment of closing a position',
  `conv_need` tinyint(1) NOT NULL COMMENT 'profit currency is different from deposit currency',
  `conv_symbol` char(16) NOT NULL COMMENT 'the symbol for convert profit to deposit currency',
  `conv_multiply` tinyint(1) NOT NULL COMMENT 'to get deposit currency result, profit formula result need multiply conSymbol''s price or divide',
  `deposit_currency` char(16) NOT NULL,
  `profit` decimal(28,2) NOT NULL COMMENT 'profit from a trade transaction in deposit currency',
  `margin` decimal(28,2) NOT NULL,
  `margin_rate` decimal(28,2) NOT NULL,
  `commision` decimal(28,2) NOT NULL COMMENT 'order commission amount',
  `swap` decimal(28,2) NOT NULL COMMENT 'order swap in the client''s deposit',
  `open_time` datetime NOT NULL COMMENT 'order open time',
  `close_time` datetime NOT NULL COMMENT 'order close time',
  `modify_time` datetime NOT NULL COMMENT 'order modify time',
  `expiration` datetime NOT NULL COMMENT 'pending order expiration time',
  `comment` varchar(255) DEFAULT NULL COMMENT 'comment to an order',
  `reason` int(11) NOT NULL COMMENT 'reason for placing the order, 0-client, 1-expert, 2-dealer, 3-signal, 4-gateway, 5-mobile, 6-web, 7-api',
  `delete_flag` tinyint(1) NOT NULL COMMENT 'marked order as deleted when SL/TP/SO activation',
  PRIMARY KEY (`ticket`,`login`),
  KEY `login` (`login`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;