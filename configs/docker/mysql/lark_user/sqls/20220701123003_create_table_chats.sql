DROP TABLE IF EXISTS `chats`;
CREATE TABLE `chats` (
  `chat_id` bigint NOT NULL COMMENT 'chat ID',
  `chat_type` tinyint(1) DEFAULT '0' COMMENT 'chat type',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`chat_id`),
  UNIQUE KEY `chat_id` (`chat_id`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_chatType` (`chat_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;