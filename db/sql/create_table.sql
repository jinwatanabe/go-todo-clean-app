CREATE DATABASE IF NOT EXISTS godo;

USE godo;

CREATE TABLE IF NOT EXISTS todos (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `done` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
);

INSERT INTO todos (title, done) VALUES
  ('Learn Go', 0),
  ('Learn Vue', 0),
  ('Build a todo app', 0);