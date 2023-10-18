#!/bin/bash

isInstalled() {
    if command -v "$1" &> /dev/null; then
        echo "$1 is installed. Continuing..."
    else 
        echo "$1 is not installed. Aborting..."
        exit 1
    fi
}

echo $'Checking for dependencies..\n'

isInstalled "python3"
isInstalled "go"
isInstalled "mysql"
isInstalled "migrate"

echo $'\nAll dependencies are installed. \n'

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

sudo apt install python3-pip
pip install passlib mysql-connector-python

python3 generateBcryptHash.py "$sqlUsername" "$dbPassword" "$dbName" "$adminPassword"

go mod vendor
go mod tidy
echo "Setup has been completed."

echo $'run "./make.sh" to start the server'
