#
# Cookbook Name:: apache2
# Attributes:: apache
#

default["ServerName"] = "127.0.0.1"

default["vhost"]["name"] = "*"
default["vhost"]["port"] = "80"
default["vhost"]["DocumentRoot"] = "/srv/www/html"
default["vhost"]["ServerName"] = "127.0.0.1"
default["vhost"]["ErrorLog"] = "/var/log/httpd/vhost_error_log"
default["vhost"]["CustomLog"] = "logs/vhost.common-access_log commo"

default["fuel"]["name"] = "*"
default["fuel"]["port"] = "80"
default["fuel"]["DocumentRoot"] = "/srv/www/html/public"
default["fuel"]["ServerName"] = "127.0.0.1"
default["fuel"]["ErrorLog"] = "/var/log/httpd/fuel_error_log"
default["fuel"]["CustomLog"] = "logs/fuel.common-access_log commo"