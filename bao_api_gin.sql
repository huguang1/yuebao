/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Version : 50712
 Source Host           : localhost
 Source Database       : ppgo_api_demo_gin

 Target Server Version : 50712
 File Encoding         : utf-8

 Date: 10/18/2017 22:21:17 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Records of `bao_user`
-- ----------------------------
DROP TABLE IF EXISTS `bao_user`;
CREATE TABLE `bao_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(150) NOT NULL COMMENT '登录名',
    `password` varchar(64) NOT NULL COMMENT '密码',
    `last_login` datetime(6) DEFAULT NULL COMMENT '最近登录',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `bao_record`
-- ----------------------------
DROP TABLE IF EXISTS `bao_record`;
CREATE TABLE `bao_record` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `type` int(11) NOT NULL COMMENT '转账类型',
    `amount` decimal(15,2) NOT NULL COMMENT '金额',
    `member` varchar(250) DEFAULT NULL COMMENT '会员账号',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `bao_rate`
-- ----------------------------
DROP TABLE IF EXISTS `bao_rate`;
CREATE TABLE `bao_rate` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `interest_rate` decimal(15,5) NOT NULL COMMENT '利率',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `bao_member`
-- ----------------------------
DROP TABLE IF EXISTS `bao_member`;
CREATE TABLE `bao_member` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `account` varchar(250) NOT NULL COMMENT '用户名',
    `password` varchar(250) NOT NULL COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `bao_member_account_2f175939_uniq` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `bao_member`
-- ----------------------------
DROP TABLE IF EXISTS `bao_balance`;
CREATE TABLE `bao_balance` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `balance` decimal(15,2) NOT NULL COMMENT '当前余额',
    `transfer_in` decimal(15,2) NOT NULL COMMENT '今日转入',
    `transfer_out` decimal(15,2) NOT NULL COMMENT '今日转出',
    `interest` decimal(15,2) NOT NULL COMMENT '利息',
    `member` varchar(250) DEFAULT NULL COMMENT '会员账号',
    `is_compute` int(11) NOT NULL COMMENT '是否计算',
    `create_time` datetime(6) DEFAULT NULL COMMENT '时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

BEGIN;
INSERT INTO `bao_user` (`username`, `password`) VALUES ('123', 'python123');
COMMIT;
SET FOREIGN_KEY_CHECKS = 1;
