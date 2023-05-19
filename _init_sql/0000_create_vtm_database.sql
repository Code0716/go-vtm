CREATE DATABASE IF NOT EXISTS `vtm_db`;

-- admin_users
CREATE TABLE IF NOT EXISTS `vtm_db`.`admin_users`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `admin_id` char(36) NOT NULL DEFAULT '' COMMENT 'adminユーザーID(UUID)',
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

-- address
CREATE TABLE IF NOT EXISTS  `vtm_db`.`address` (
  `id` char(36) NOT NULL DEFAULT '' COMMENT '住所のID(UUID)',
  `post_code` varchar(8) NOT NULL COMMENT '郵便番号',
  `address` char(255) NOT NULL COMMENT '住所',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- department
CREATE TABLE IF NOT EXISTS  `vtm_db`.`department` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` char(255) NOT NULL COMMENT '所属部署/所属店舗名',
  `address_id` char(36) NOT NULL DEFAULT '' COMMENT '住所のID(UUID)',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`address_id`) REFERENCES `address` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- members
CREATE TABLE IF NOT EXISTS `vtm_db`.`members`(
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `member_id` char(36) NOT NULL DEFAULT '' COMMENT 'member ID(UUID)',
  `name` VARCHAR(255) NOT NULL DEFAULT '名称未設定' COMMENT '名前',
  `birthday` datetime DEFAULT NULL COMMENT '誕生日',
  `password` VARCHAR(255) DEFAULT NULL COMMENT 'password',
  `mail_address` VARCHAR(255) DEFAULT NULL COMMENT 'メールアドレス',
  `phone_number` VARCHAR(255) DEFAULT NULL COMMENT '電話番号',
  `status` enum('init', 'active', 'other') NOT NULL DEFAULT 'init' COMMENT 'status',
  `role` enum('manager','accountant','common') NOT NULL COMMENT 'memberのstatus manager:店長,accountant:経理,一般:common',
  `employment_status` enum('hourly', 'monthly', 'day', 'other') NOT NULL DEFAULT 'hourly' COMMENT '雇用形態 時給:hourly, 月給:monthly, 日払い（spot）:day, その他(経営者など):other',
  `unit_price` SMALLINT UNSIGNED DEFAULT NULL COMMENT '月給/単価',
  `department_id` char(36) NOT NULL DEFAULT '' COMMENT '所属部署/所属店舗のID(UUID)',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY(`id`),
  UNIQUE KEY `member_id` (`member_id`),
  FOREIGN KEY (`department_id`) REFERENCES `department` (`address_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- attendances
CREATE TABLE IF NOT EXISTS  `vtm_db`.`attendances` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `admin_user_id` char(36) DEFAULT NULL COMMENT 'admin_user ID(UUID)',
  `member_id` char(36) DEFAULT NULL COMMENT 'member ID(UUID)',
  `start_time` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT '勤務開始日時',
  `end_time` datetime DEFAULT NULL COMMENT '勤務終了日時',
  `break_time` datetime DEFAULT NULL COMMENT '休憩時間',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY (`id`),
  KEY `attend_id_index` (`start_time`,`member_id`),
  KEY `member_id` (`member_id`),
  FOREIGN KEY (`admin_user_id`) REFERENCES `admin_users` (`admin_id`),
  FOREIGN KEY (`member_id`) REFERENCES `members` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- invoice
CREATE TABLE IF NOT EXISTS  `vtm_db`.`invoice` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `billing_date` datetime NOT NULL COMMENT '請求日付',
  `admin_user_id` char(36) DEFAULT NULL COMMENT 'admin_user ID(UUID)',
  `member_id` char(36) DEFAULT NULL COMMENT 'member ID(UUID)',
  `authorizer_id` char(36) NOT NULL COMMENT '承認者 ID(UUID)',
  `billing_amount`int(11) unsigned DEFAULT NULL COMMENT '請求額',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`admin_user_id`) REFERENCES `attendances` (`admin_user_id`),
  FOREIGN KEY (`member_id`) REFERENCES `attendances` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
