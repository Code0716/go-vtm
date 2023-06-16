-- attendances
CREATE TABLE IF NOT EXISTS  `vtm_db`.`attendances` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` char(36) DEFAULT NULL COMMENT 'user ID(UUID)',
  `start_time` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT '勤務開始日時',
  `end_time` datetime DEFAULT NULL COMMENT '勤務終了日時',
  `break_time` time DEFAULT NULL COMMENT '休憩時間',
  `created_at` datetime NOT NULL NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_date',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated_date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY (`id`),
  KEY `attend_id_index` (`start_time`,`user_id`),
  KEY `user_id` (`user_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
