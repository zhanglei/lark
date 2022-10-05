DROP TABLE IF EXISTS `avatars`;
CREATE TABLE `avatars` (
  `avatar_id` bigint NOT NULL,
  `owner_id` bigint NOT NULL COMMENT '用户ID/ChatID',
  `owner_type` tinyint(1) DEFAULT '0' COMMENT '1:用户头像 2:群头像',
  `avatar_small` varchar(64) DEFAULT '' COMMENT '小图 72*62',
  `avatar_medium` varchar(64) DEFAULT '' COMMENT '中图 240*240',
  `avatar_large` varchar(64) DEFAULT '' COMMENT '大图 640*640',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`avatar_id`),
  UNIQUE KEY `avatarId_ownerId` (`avatar_id`,`owner_id`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_ownerType` (`owner_type`),
  KEY `idx_avatarSmall` (`avatar_small`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
