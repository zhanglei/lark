DROP TABLE IF EXISTS `avatars`;
CREATE TABLE `avatars` (
  `owner_id` bigint NOT NULL COMMENT '用户ID/ChatID',
  `owner_type` tinyint(1) NOT NULL COMMENT '1:用户头像 2:群头像',
  `avatar_small` varchar(64) NOT NULL COMMENT '小图 72*62',
  `avatar_medium` varchar(64) NOT NULL COMMENT '中图 240*240',
  `avatar_large` varchar(64) NOT NULL COMMENT '大图 640*640',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`owner_id`),
  UNIQUE KEY `owner_id` (`owner_id`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_ownerType` (`owner_type`),
  KEY `idx_avatarSmall` (`avatar_small`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
