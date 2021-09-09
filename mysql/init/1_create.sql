CREATE DATABASE capsule;

use capsule;

CREATE TABLE User
(
    id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    name varchar(10) NOT NULL UNIQUE,
    token varchar(10) NOT NULL UNIQUE
);

CREATE TABLE Characters
(
    id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    name varchar(20) NOT NULL UNIQUE,
    rarity varchar(20) NOT NULL,
    attack INT NOT NULL,
    defence INT NOT NULL,
    recovery INT NOT NULL
);

CREATE TABLE UserCharacters
(
    id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
    user_id INT NOT NULL,
    character_id INT NOT NULL
);

INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("ミシシッピアカミミガメ", "☆", 3,3,0);
INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("イシガメ", "☆☆", 1,2,4);
INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("スッポン", "☆☆☆", 5,1,2);
INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("クサガメ", "☆☆☆☆", 2,4,3);
INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("カミツキガメ", "☆☆☆☆☆", 7,5,1);
INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("リュウキュウヤマガメ", "☆☆☆☆☆☆", 4,4,7);
INSERT INTO Characters(name, rarity, attack, defence, recovery) VALUES("ウミガメ", "☆☆☆☆☆☆☆", 6,6,6);