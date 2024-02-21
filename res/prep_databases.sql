SET foreign_key_checks = 0;
START TRANSACTION;

CREATE TABLE Stories (
    `Id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,    
    `Story` TEXT,
    `Prompt` TEXT,
    `StoryType` INT NOT NULL,
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
    `StoryType` INT NOT NULL,
    `Timestamp` DATETIME
);

CREATE TABLE Modifiers (
    `Id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,    
    `Modifier` TEXT,
    `StoryType` INT NOT NULL,
    `Timestamp` DATETIME
);

CREATE INDEX IX_Stories_StoryTypes ON Stories (StoryType);
CREATE INDEX IX_Styles_StoryTypes ON Styles (StoryType);
CREATE INDEX IX_Modifiers_StoryTypes ON Modifiers (StoryType);

COMMIT;