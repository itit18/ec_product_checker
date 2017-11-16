#
# Cookbook Name:: apache2
# Recipe:: eccube
#
# ECCUBE用のヴァーチャルホスト設定
#
# All rights reserved - Do Not Redistribute
#

template 'vhost.conf' do
  path "/etc/httpd/conf.d/vhost.conf"
  source 'vhost.conf.erb'
  owner 'root'
  group 'root'
  mode 0644
  backup false
  variables({
    :port => node["fuel"]["port"],
    :name => node["fuel"]["name"],
    :DocumentRoot => node["fuel"]["DocumentRoot"],
    :ErrorLog => node["fuel"]["ErrorLog"],
    :ServerName => node["fuel"]["ServerName"],
    :CustomLog => node["fuel"]["CustomLog"]
  })
end