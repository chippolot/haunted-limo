SET foreign_key_checks = 0;
START TRANSACTION;

CREATE TABLE Stories (
    `Id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,    
    `Story` TEXT,
    `Prompt` TEXT,
    `Timestamp` DATETIME
);

CREATE TABLE Themes (
    `Id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,    
    `Theme` TEXT,
    `Timestamp` DATETIME
);

CREATE TABLE Styles (
    `Id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,    
    `Style` TEXT,
    `Timestamp` DATETIME
);

CREATE TABLE Modifiers (
    `Id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,    
    `Modifier` TEXT,
    `Timestamp` DATETIME
);

COMMIT;