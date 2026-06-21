#!/usr/bin/env bash

rm /etc/nginx/sites-enabled/podmanproxy.conf
rm /etc/nginx/sites-avaliable/podmanproxy.conf
cp ./configs/nginx/podmanproxy.conf /etc/nginx/sites-avaliable/podmanproxy.conf
sudo ln -s /etc/nginx/sites-available/podmanproxy.conf /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
