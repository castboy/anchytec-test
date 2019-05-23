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

 Date: 23/05/2019 14:26:12
*/

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `login` int(11) NOT NULL,
  `balance` decimal(28,2) NOT NULL,
  `margin` decimal(28,2) NOT NULL,
  `deposit_currency` char(16) NOT NULL,
  `margin_stop_out` decimal(28,2) NOT NULL,
  PRIMARY KEY (`login`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
