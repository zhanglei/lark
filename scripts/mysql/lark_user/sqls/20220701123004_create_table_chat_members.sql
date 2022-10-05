DROP TABLE IF EXISTS `chat_members`;
CREATE TABLE `chat_members` (
  `chat_id` bigint NOT NULL COMMENT 'chat ID',
  `chat_hash` CHAR(32) COMMENT 'chat hash值',
  `chat_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'chat type',
  `uid` bigint NOT NULL COMMENT '用户ID',
  `mute` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否开启免打扰',
  `display_name` VARCHAR(64) NOT NULL COMMENT '显示名称',
  `avatar_key` varchar(50) NOT NULL DEFAULT '' COMMENT '头像 72*72',
  `sync` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否同步用户信息',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'chat状态',
  `platform` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1:iOS 2:安卓',
  `server_id` int NOT NULL DEFAULT '0' COMMENT '服务器ID',
  `settings` varchar(512) NOT NULL DEFAULT '' COMMENT '用户设置',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`chat_id`,`uid`),
  UNIQUE KEY `chatId_uid` (`chat_id`,`uid`),
  KEY `idx_uid_sync` (`uid`,`sync`),
  KEY `idx_status` (`status`),
  KEY `idx_deletedTs` (`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*
 修改`users`.`nickname`,当 `chat_members`.`sync`=1 ,需要同步修改`chat_members`.`display_name`
 修改`user_avatars`.`avatar_*`,需要同步修改`chat_members`.`avatar_key`
 需要更新缓存信息
 */
