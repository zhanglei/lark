DROP TABLE IF EXISTS `chats`;
CREATE TABLE `chats` (
  `chat_id` bigint NOT NULL COMMENT 'chat ID',
  `creator_uid` bigint NOT NULL COMMENT '创建者 uid',
  `chat_hash` CHAR(32) NOT NULL COMMENT 'chat hash值',
  `chat_type` tinyint(1) NOT NULL  COMMENT 'chat type 1:私聊/2:群聊',
  `avatar_key` varchar(64) NOT NULL COMMENT '小图 72*62',
  `title` varchar(128) DEFAULT '' COMMENT 'chat标题',
  `about` varchar(255) DEFAULT '' COMMENT '关于',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`chat_id`),
  UNIQUE KEY `chat_id` (`chat_id`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_chatHash` (`chat_hash`),
  KEY `idx_chatType` (`chat_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;