CREATE TABLE `users` (
  `userId` int(12) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `userName` varchar(255) NOT NULL,
  `salt` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `isAdmin` BOOLEAN DEFAULT FALSE
) ENGINE=InnoDB;
