-- Create Table Users
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL DEFAULT '',
  `password` VARCHAR(255) NOT NULL DEFAULT '',

  PRIMARY KEY (`id`),
  UNIQUE (`username`)
) ENGINE = INNODB COMMENT='User table';


-- Create table link
DROP TABLE IF EXISTS `link`;
CREATE TABLE IF NOT EXISTS `link`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL DEFAULT '',
    `address` VARCHAR(255) NOT NULL DEFAULT '',

    `user_id` INT ,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) 
) ENGINE = INNODB COMMENT='link table';
