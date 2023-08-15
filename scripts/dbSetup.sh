#!/bin/bash
echo "Kindly Enter your mySql username: "
read sqlUsername

echo "Kindly Enter your mySql password: " 
read -s dbPassword

echo "Kindly Enter mySql database name: " 
read  dbName

echo "$sqlUsername $dbPassword $dbName"

mysql -u "$sqlUsername" -p"$dbPassword" -e "CREATE DATABASE IF NOT EXISTS $dbName;"

migrate -path ../database/migration/ -database "mysql://$sqlUsername:$dbPassword@tcp(localhost:3306)/$dbName" -verbose up

read -p "Enter JWT Secret Key: " secretKey

cat << EOF > ../db.yaml
DB_USERNAME: "$sqlUsername"
DB_PASSWORD: "$dbPassword"
DB_HOST: 127.0.0.1:3306
DB_NAME: "$dbName"
JWTSecretKey: "$secretKey"
EOF

echo "db.yaml has been created successfully."

echo "Kindly Enter your Admin password: " 
read -s adminPassword

python3 generate_bcrypt_hash.py "$sqlUsername" "$dbPassword" "$dbName" "$adminPassword"