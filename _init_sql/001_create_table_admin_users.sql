-- admin_users
CREATE TABLE IF NOT EXISTS `vtm_db`.`admin_users`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `admin_id` char(36) NOT NULL DEFAULT '' COMMENT 'adminユーザーID（UUIDv4）',
  `name` VARCHAR(255) NOT NULL DEFAULT '名称未設定' COMMENT '名前',
  `password` VARCHAR(255) NOT NULL COMMENT 'password',
  `mail_address` VARCHAR(255) DEFAULT NULL COMMENT 'メールアドレス',
  `status` enum('init', 'active', 'other') NOT NULL DEFAULT 'init' COMMENT 'status',
  `permission` enum('admin', 'manager', 'general') NOT NULL DEFAULT 'general' COMMENT "管理権限",
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_data',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY(`id`),
  UNIQUE KEY `admin_id` (`admin_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
