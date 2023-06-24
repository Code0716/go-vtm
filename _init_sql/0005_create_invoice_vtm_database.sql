-- invoice
CREATE TABLE IF NOT EXISTS  `vtm_db`.`invoice` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` char(36) DEFAULT NULL COMMENT 'user ID(UUID)',
  `authorizer_id` char(36) NOT NULL COMMENT '承認者 ID(UUID)',
  `billing_date` datetime NOT NULL COMMENT '請求日付',
  `billing_amount`int(11) unsigned DEFAULT NULL COMMENT '請求額',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `attendances` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
