-- members
CREATE TABLE IF NOT EXISTS `vtm_db`.`members`(
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `member_id` char(36) NOT NULL DEFAULT '' COMMENT 'member ID（UUIDv4）',
  `name` VARCHAR(255) NOT NULL DEFAULT '名称未設定' COMMENT '名前',
  `password` VARCHAR(255) DEFAULT NULL COMMENT 'password',
  `phone_number` VARCHAR(255) DEFAULT NULL COMMENT '電話番号',
  `status` enum('init', 'active', 'other') NOT NULL DEFAULT 'init' COMMENT 'status',
  `hourly_price` SMALLINT UNSIGNED DEFAULT NULL COMMENT '時間単価',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_data',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY(`id`),
  UNIQUE KEY `member_id` (`member_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
