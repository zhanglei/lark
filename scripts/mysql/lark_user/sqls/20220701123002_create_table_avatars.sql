DROP TABLE IF EXISTS `user_avatars`;
CREATE TABLE `user_avatars` (
  `uid` bigint NOT NULL COMMENT '用户ID',
  `avatar_small` varchar(255) DEFAULT '' COMMENT '小图 72*62',
  `avatar_medium` varchar(255) DEFAULT '' COMMENT '中图 240*240',
  `avatar_large` varchar(255) DEFAULT '' COMMENT '大图 640*640',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `uid` (`uid`),
  KEY `idx_deletedTs` (`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
