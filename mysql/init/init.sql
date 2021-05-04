CREATE DATABASE IF NOT EXISTS todo_app;
USE todo_app;
CREATE TABLE IF NOT EXISTS todos (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  title VARBINARY(1024) NOT NULL,
  created_at datetime default current_timestamp
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
-- CREATE DATABASE IF NOT EXISTS todo_app;
-- USE todo_app;
-- CREATE TABLE IF NOT EXISTS todos (
--   id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
--   title VARCHAR(256) NOT NULL,
--   created_at datetime default current_timestamp
-- ) ENGINE = InnoDB DEFAULT CHARSET = utf8;
INSERT INTO todos(title)
VALUES ("テスト勉強");
INSERT INTO todos(title)
VALUES ("テスト勉強");
INSERT INTO todos(title)
VALUES ("テスト勉強");