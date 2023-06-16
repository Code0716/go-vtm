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
