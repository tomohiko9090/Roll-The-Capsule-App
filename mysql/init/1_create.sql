-- golang_dbという名前のデータベースを作成
CREATE DATABASE golang_db;
-- golang_dbをアクティブ
use golang_db;
-- usersテーブルを作成。名前とパスワード
CREATE TABLE users (
    id INT(11) AUTO_INCREMENT NOT NULL,
    name VARCHAR(64) NOT NULL,
    password CHAR(30) NOT NULL,
    PRIMARY KEY (id)
);
-- usersテーブルに２つレコードを追加
INSERT INTO users (name, password) VALUES ("gophar", "5555");
INSERT INTO users (name, password) VALUES ("octcat", "0000");