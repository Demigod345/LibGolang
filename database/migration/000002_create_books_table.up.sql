CREATE TABLE `books` (
  `bookId` int(12) AUTO_INCREMENT NOT NULL,
  `title` varchar(255) NOT NULL UNIQUE,
  `totalQuantity` int(11)  DEFAULT 1,
  `available` int(11)  DEFAULT 1 CONSTRAINT non_negative_constraint CHECK (available >= 0),
  PRIMARY KEY (`bookId`)
  )ENGINE=InnoDB;