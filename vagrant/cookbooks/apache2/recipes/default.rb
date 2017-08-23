#
# Cookbook Name:: apache2
# Recipe:: default
#
#
# All rights reserved - Do Not Redistribute
#

if platform?("amazon")
  # amazonの場合、php55を選ぶと強制でapache24なので
  apache_conf_template = 'apache24.conf'
  yum_package 'httpd24' do
    options '--enablerepo=epel,remi-php55'
    action :install
  end
else
  apache_conf_template = 'apache22.conf'
  
  yum_package 'httpd' do
    action :install
  end
end

template apache_conf_template do
  path "/etc/httpd/conf/httpd.conf"
  owner 'root'
  group 'root'
  mode 0644
end

service 'apache2' do
  service_name 'httpd'
  action :enable
end

template 'vhost.conf' do
  path "/etc/httpd/conf.d/vhost.conf"
  source 'vhost.conf.erb'
  owner 'root'
  group 'root'
  mode 0644
  backup false
  variables({
    :port => node["vhost"]["port"],
    :name => node["vhost"]["name"],
    :DocumentRoot => node["vhost"]["DocumentRoot"],
    :ErrorLog => node["vhost"]["ErrorLog"],
    :ServerName => node["vhost"]["ServerName"],
    :CustomLog => node["vhost"]["CustomLog"]
  })
end