CREATE TABLE IF NOT EXISTS requests (
  requestId INT(12) AUTO_INCREMENT PRIMARY KEY,
  bookId INT(12),
  userId INT(12),
  state ENUM('requested', 'issued', 'checkedIn', 'AdminRequest') DEFAULT 'requested',
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (bookId) REFERENCES books(bookId),
  FOREIGN KEY (userId) REFERENCES users(userId)
) ENGINE=InnoDB;
