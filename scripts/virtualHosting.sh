#!/bin/bash

echo "Kindly Enter your domain name: "
read domainName

echo "Kindly Enter your email: "
read email

sudo apt install apache2
sudo a2enmod proxy proxy_http
cd /etc/apache2/sites-available

sudo tee $domainName.conf > /dev/null <<EOF
    <VirtualHost *:80>
        ServerName $domainName
        ServerAdmin $email
        ProxyPreserveHost On
        ProxyPass / http://127.0.0.1:8000/
        ProxyPassReverse / http://127.0.0.1:8000/
        TransferLog /var/log/apache2/${domain}_access.log
        ErrorLog /var/log/apache2/${domain}_error.log
    </VirtualHost>
EOF

sudo a2ensite $domainName.conf
sudo a2dissite 000-default.conf
sudo apache2ctl configtest
echo "127.0.0.1  $domainName" | sudo tee -a /etc/hosts
sudo systemctl restart apache2
open http://$domainName

echo $'\nVirtual hosting has been done\n'
echo "Open http://$domainName in your browser." 