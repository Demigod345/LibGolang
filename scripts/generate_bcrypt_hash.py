#!/usr/bin/env python3

from passlib.hash import bcrypt
import mysql.connector
import sys

if len(sys.argv) != 5:
    print("Usage: python generate_bcrypt_hash_and_insert.py <db_user> <db_password> <password>")
    sys.exit(1)

db_user = sys.argv[1]
db_password = sys.argv[2]
db_name = sys.argv[3]
password = sys.argv[4]

# Generate bcrypt hash
hashed_password = bcrypt.using(rounds=4, ident='2a').hash(password)

# Connect to MySQL database
db_connection = mysql.connector.connect(
    host="localhost",
    user=db_user,
    password=db_password,
    database=db_name  # Replace with your database name
)

# Create a cursor
cursor = db_connection.cursor()

# Insert the hash into the database
insert_query = "INSERT INTO users (username, hash, isAdmin) VALUES (%s, %s, true)"
user_data = ("Admin", hashed_password)
cursor.execute(insert_query, user_data)

# Commit the changes and close the connection
db_connection.commit()
cursor.close()
db_connection.close()

print("Password hash inserted into the database.")
