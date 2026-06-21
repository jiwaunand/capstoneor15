#!/usr/bin/env bash

cp ./configs/nginx/podmanproxy.conf /etc/nginx/sites-avaliable/podmanproxy.conf
sudo ln -s /etc/nginx/sites-available/podmanproxy.conf /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
