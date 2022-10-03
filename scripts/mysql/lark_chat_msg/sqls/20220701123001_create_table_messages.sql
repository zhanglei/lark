/*
 Navicat Premium Data Transfer

 Source Server         : lark_msg
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:13307
 Source Schema         : lark_msg

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 09/08/2022 18:24:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `srv_msg_id` bigint NOT NULL COMMENT '服务端消息号',
  `cli_msg_id` bigint NOT NULL DEFAULT '0' COMMENT '客户端消息号',
  `sender_id` bigint NOT NULL DEFAULT '0' COMMENT '发送者uid',
  `receiver_id` bigint NOT NULL DEFAULT '0' COMMENT '接收者uid',
  `sender_platform` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发送者平台',
  `chat_id` bigint NOT NULL DEFAULT '0' COMMENT '会话ID',
  `chat_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '会话类型',
  `seq_id` int NOT NULL DEFAULT '0' COMMENT '消息唯一ID',
  `msg_from` int NOT NULL DEFAULT '0' COMMENT '消息来源',
  `msg_type` int NOT NULL DEFAULT '0' COMMENT '消息类型',
  `body` text CHARACTER SET utf8mb4 NOT NULL COMMENT '消息本体',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '消息状态',
  `sent_ts` bigint NOT NULL DEFAULT '0' COMMENT '客户端本地发送时间',
  `srv_ts` bigint NOT NULL DEFAULT '0' COMMENT '服务端接收消息的时间',
  `updated_ts` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_ts` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`srv_msg_id`),
  UNIQUE KEY `srv_msg_id` (`srv_msg_id`),
  UNIQUE KEY `chatId_seqId` (`chat_id`,`seq_id`),
  KEY `idx_chatId_seqId_deletedTs` (`chat_id`,`seq_id`,`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
