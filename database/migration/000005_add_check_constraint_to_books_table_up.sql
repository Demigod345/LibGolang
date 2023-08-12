ALTER TABLE books ADD CONSTRAINT checkQuantity CHECK (available <= totalQuantity);
