DROP TABLE IF EXISTS `chat_requests`;
CREATE TABLE `chat_requests` (
  `request_id` bigint NOT NULL COMMENT 'request ID',
  `request_ts` bigint DEFAULT '0' COMMENT '申请时间',
  `chat_type` tinyint(1) DEFAULT '0' COMMENT '1:私聊/2:群聊',
  `initiator_uid` bigint NOT NULL DEFAULT '0' COMMENT '发起人 UID',
  `target_id` bigint NOT NULL DEFAULT '0'  COMMENT '被邀请人UID/群ID',
  `request_msg` varchar(255) NOT NULL DEFAULT '' COMMENT 'request消息',
  `handler_uid` bigint DEFAULT '0' COMMENT '处理人 UID',
  `handle_result` tinyint(1) DEFAULT '0' COMMENT '结果',
  `handle_msg` varchar(255) DEFAULT '' COMMENT '处理消息',
  `handled_ts` bigint DEFAULT '0' COMMENT '处理时间',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`request_id`),
  UNIQUE KEY `request_id` (`request_id`),
  KEY `id_chatType_initiatorUid_handleResult` (`chat_type`,`initiator_uid`,`handle_result`),
  KEY `id_chatType_targetId_handleResult` (`chat_type`,`target_id`,`handle_result`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;