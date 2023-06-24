-- users
CREATE TABLE IF NOT EXISTS `vtm_db`.`users`(
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` char(36) NOT NULL DEFAULT '' COMMENT 'user ID(UUID)',
  `name` VARCHAR(255) NOT NULL DEFAULT '名称未設定' COMMENT '名前',
  `birthday` datetime DEFAULT NULL COMMENT '誕生日',
  `mail_address` VARCHAR(255) DEFAULT NULL COMMENT 'メールアドレス',
  `phone_number` VARCHAR(255) DEFAULT NULL COMMENT '電話番号',
  `status` enum('init', 'active', 'other') NOT NULL DEFAULT 'init' COMMENT 'status',
  `role` enum('admin','manager','accountant','common') NOT NULL COMMENT 'userのrole admin:管理者 manager:店長,accountant:経理,一般:common',
  `employment_status` enum('annual', 'monthly', 'hourly', 'day', 'other') NOT NULL DEFAULT 'hourly' COMMENT '雇用形態 年俸:annual, 月給:monthly, 日払い（spot）:day, 時給:hourly, その他(経営者など):other',
  `unit_price` SMALLINT UNSIGNED DEFAULT NULL COMMENT '月給/単価',
  `department_id` char(36) DEFAULT NULL COMMENT '所属部署/所属店舗のID(UUID)',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY(`id`),
  UNIQUE KEY `user_id` (`user_id`),
  FOREIGN KEY (`department_id`) REFERENCES `department` (`address_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
