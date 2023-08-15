# включено в /mysql/my.cnf
# включаем лог медленных запросов (1-включить; 0-выключить)
# SET global slow_launch_time = 1;
#  путь до файла, в который будут записываться все медленные запросы
# SET global slow_query_log_file = '/var/log/mysql/query_slow.log';
# задаем время, при превышении которого запрос будет считаться медленным (1 секунда);
# SET global slow_query_log = 1;
# очищаем log-файл
# FLUSH LOGS;

CREATE DATABASE IF NOT EXISTS irol CHARACTER SET utf8 COLLATE utf8_general_ci;
#
CREATE USER 'irol_user'@'localhost' IDENTIFIED BY 'irol_user_password';
GRANT SELECT, INSERT, UPDATE, DELETE ON `irol`.* TO 'irol_user'@'localhost';
REVOKE ALL PRIVILEGES ON *.* FROM 'irol_user'@'localhost';
REVOKE GRANT OPTION ON *.* FROM 'irol_user'@'localhost';
GRANT ALL ON irol.* TO 'irol_user'@'localhost' WITH MAX_QUERIES_PER_HOUR 0 MAX_UPDATES_PER_HOUR 0 MAX_CONNECTIONS_PER_HOUR 0 MAX_USER_CONNECTIONS 0;
FLUSH PRIVILEGES;
#
USE irol;
#
CREATE TABLE IF NOT EXISTS users (
  uid INT UNSIGNED NOT NULL AUTO_INCREMENT,
  login VARCHAR(12) NOT NULL,
  psw VARCHAR(30) NOT NULL,
  username VARCHAR(30) NOT NULL,
  usergrp INT UNSIGNED NOT NULL,
  token VARCHAR(50) NOT NULL,
  lastaccess DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`))
  ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;
#
ALTER TABLE `users` ADD UNIQUE(`token`);
#
INSERT INTO `users` (`login`, `username`, `psw`,`usergrp`, `token`) VALUES ('admin', 'Администратор', '12345', '1', 'Ghj54kGE46Bsed');
#
CREATE TABLE IF NOT EXISTS usergrp (
  uid INT UNSIGNED NOT NULL AUTO_INCREMENT,
  grpname VARCHAR(30) NOT NULL,
  access VARCHAR(50) NOT NULL,
   PRIMARY KEY (`uid`)) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

INSERT INTO `usergrp` (`uid`,`grpname`, `access`) VALUES ('1', 'Администратор', '');
