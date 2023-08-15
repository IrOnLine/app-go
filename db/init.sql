CREATE DATABASE IF NOT EXISTS irol;

#
GRANT SELECT, INSERT, UPDATE, DELETE ON `irol`.* TO 'user_irol'@'mariadb';
REVOKE ALL PRIVILEGES ON *.* FROM 'user_irol'@'mariadb';
REVOKE GRANT OPTION ON *.* FROM 'user_irol'@'mariadb';
GRANT SELECT, INSERT, UPDATE, DELETE, FILE ON *.* TO 'user_irol'@'mariadb' REQUIRE NONE WITH MAX_QUERIES_PER_HOUR 0 MAX_CONNECTIONS_PER_HOUR 0 MAX_UPDATES_PER_HOUR 0 MAX_USER_CONNECTIONS 0;
#
USE irol;
#
CREATE TABLE IF NOT EXISTS `users` (
  `uid` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
  `login` VARCHAR(12) NOT NULL ,
  `psw` VARCHAR(30) NOT NULL ,
  `username` VARCHAR(30) NOT NULL ,
  `usergrp` INT UNSIGNED NOT NULL ,
  `token` VARCHAR(50) NOT NULL ,
  `lastaccess` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`uid`)) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;
#
ALTER TABLE `users` ADD UNIQUE(`token`);
#
INSERT INTO `users` (`login`, `username`, `usergrp`, `token`) VALUES ('admin', 'Администратор', '1', 'Ghj54kGE46Bsed');
#
CREATE TABLE IF NOT EXISTS `usergrp` (
  `uid` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
  `grpname` VARCHAR(30) NOT NULL ,
  `access` VARCHAR(50) NOT NULL ,
   PRIMARY KEY (`uid`)) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;
#
INSERT INTO `usergrp` (`uid`,`grpname`, `access`) VALUES ('1', 'Администратор', '');
