-- Create Table Users
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL DEFAULT '',
  `password` VARCHAR(255) NOT NULL DEFAULT '',

  PRIMARY KEY (`id`),
  UNIQUE (`username`)
) ENGINE = INNODB COMMENT='User table';

-- [DDL] Create new table for Role
DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `type` VARCHAR(255) NOT NULL DEFAULT '',
    `rank` INT NOT NULL DEFAULT 0,

    -- Utility columns
    `status` SMALLINT NOT NULL DEFAULT '1',
    `flag` INT NOT NULL DEFAULT '0',
    `meta` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` VARCHAR(255),
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_by` VARCHAR(255),
    `deleted_at`TIMESTAMP,
    `deleted_by` VARCHAR(255),
    PRIMARY KEY (`id`)
);

-- [DDL] Add new column `fk_role_id` in user table
ALTER TABLE `user` ADD `fk_role_id` INT COMMENT 'Foreign Key To role Id' AFTER `id`;
