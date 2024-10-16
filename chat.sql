CREATE DATABASE chat;

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uuid` varchar(150) NOT NULL COMMENT 'uuid',
    `username` varchar(191) NOT NULL COMMENT 'user name',
    `nickname` varchar(255) DEFAULT NULL COMMENT 'nickname',
    `email` varchar(80) DEFAULT NULL COMMENT 'mailbox',
    `password` varchar(150) NOT NULL COMMENT 'password',
    `avatar` varchar(250) NOT NULL COMMENT 'avatar',
    `create_at` datetime(3) DEFAULT NULL,
    `update_at` datetime(3) DEFAULT NULL,
    `delete_at` bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `idx_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='User watch';


DROP TABLE IF EXISTS `user_friends`;
CREATE TABLE IF NOT EXISTS `user_friends` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` datetime(3) DEFAULT NULL COMMENT 'Creation time',
    `updated_at` datetime(3) DEFAULT NULL COMMENT 'Update time',
    `deleted_at` bigint unsigned DEFAULT NULL COMMENT 'Delete time stamp',
    `user_id` int DEFAULT NULL COMMENT 'User ID',
    `friend_id` int DEFAULT NULL COMMENT 'Friend ID',
    PRIMARY KEY (`id`),
    KEY `idx_user_friends_user_id` (`user_id`),
    KEY `idx_user_friends_friend_id` (`friend_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT 'Friends information table';


DROP TABLE IF EXISTS `messages`;
CREATE TABLE IF NOT EXISTS `messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` datetime(3) DEFAULT NULL COMMENT 'Creation time',
  `updated_at` datetime(3) DEFAULT NULL COMMENT 'Update time',
  `deleted_at` bigint unsigned DEFAULT NULL COMMENT 'Delete time stamp',
  `from_user_id` int DEFAULT NULL COMMENT 'Sender ID',
  `to_user_id` int DEFAULT NULL COMMENT 'Recipient ID',
  `content` varchar(2500) DEFAULT NULL COMMENT 'Content',
  `url` varchar(350) DEFAULT NULL COMMENT '''File or picture address''',
  `pic` text COMMENT 'Thumbnail',
  `message_type` smallint DEFAULT NULL COMMENT '''Message type: 1 single chat, 2 group chat''',
  `content_type` smallint DEFAULT NULL COMMENT '''Message content type: 1 text, 2 voice, 3 video''',
  PRIMARY KEY (`id`),
  KEY `idx_messages_deleted_at` (`deleted_at`),
  KEY `idx_messages_from_user_id` (`from_user_id`),
  KEY `idx_messages_to_user_id` (`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT 'Message table';


DROP TABLE IF EXISTS `groups`;
CREATE TABLE IF NOT EXISTS  `groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  `user_id` int DEFAULT NULL COMMENT '''Group owner ID''',
  `name` varchar(150) DEFAULT NULL COMMENT '''Group name',
  `notice` varchar(350) DEFAULT NULL COMMENT '''Group announcement',
  `uuid` varchar(150) NOT NULL COMMENT '''uuid''',
  PRIMARY KEY (`id`),
  KEY `idx_groups_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT 'Group table';


DROP TABLE IF EXISTS `group_members`;
CREATE TABLE  IF NOT EXISTS `group_members` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  `user_id` int DEFAULT NULL COMMENT '''User ID''',
  `group_id` int DEFAULT NULL COMMENT '''Group ID''',
  `nickname` varchar(350) DEFAULT NULL COMMENT '''Nick name',
  `mute` smallint DEFAULT NULL COMMENT '''Whether to prohibit''',
  PRIMARY KEY (`id`),
  KEY `idx_group_members_user_id` (`user_id`),
  KEY `idx_group_members_group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT 'Group member table';
