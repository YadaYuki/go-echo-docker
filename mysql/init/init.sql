CREATE DATABASE IF NOT EXISTS todo_app;
USE todo_app;
CREATE TABLE IF NOT EXISTS todos (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  title VARCHAR(256) NOT NULL,
  created_at datetime default current_timestamp
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
INSERT INTO todos(id, title)
VALUES (1, "テスト勉強");
INSERT INTO todos(id, title)
VALUES (2, "テスト勉強");
INSERT INTO todos(id, title)
VALUES (3, "テスト勉強");