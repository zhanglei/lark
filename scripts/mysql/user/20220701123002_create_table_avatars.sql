DROP TABLE IF EXISTS `user_avatars`;
CREATE TABLE `user_avatars` (
  `uid` bigint NOT NULL COMMENT '用户ID 系统生成',
  `avatar_small` varchar(255) DEFAULT '' COMMENT '小图72',
  `avatar_medium` varchar(255) DEFAULT '' COMMENT '中图240',
  `avatar_large` varchar(255) DEFAULT '' COMMENT '大图640',
  `avatar_origin` varchar(255) DEFAULT '' COMMENT '原始图',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `uid` (`uid`),
  KEY `idx_deletedTs` (`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
