-- attendance
CREATE TABLE IF NOT EXISTS `vtm_db`.`attendance`(
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'date',
  `member_id` char(36) NOT NULL COMMENT 'member ID（UUIDv4）',
  `status` enum('work','rest') NOT NULL COMMENT 'status',
  `started_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'started_at',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted_date',
  PRIMARY KEY(`id`),
  INDEX attend_id_index (`date`,`member_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
