#
# Cookbook Name:: init
# Recipe:: default
#

# yumの追加リポジトリ
bash "add remi epel" do
  code <<-EOS
    rpm -ivh http://ftp.riken.jp/Linux/fedora/epel/6/x86_64/epel-release-6-8.noarch.rpm
    rpm -ivh http://rpms.famillecollet.com/enterprise/remi-release-6.rpm
  EOS
  user "root"
  not_if { ::File.exists?("/etc/yum.repos.d/remi.repo") }
end

#　適当なパッケージ類のインストール
%w[ vim grep git tree ntp ].each do |pkg|
  package pkg do
    action :install
  end
end

#単にクライアントが欲しいだけなのでここでインストールしちゃう
package "redis" do
 options '--enablerepo=epel,remi'
 action :install
end

bash 'change time zone' do
  code 'cp -p /usr/share/zoneinfo/Japan /etc/localtime'
end

#　言語設定ファイルを置き換え
cookbook_file "i18n_jp" do
  path "/etc/sysconfig/i18n"
  action :create
end

# service登録
service "ntpd" do
  supports :restart => true
  action [ :enable, :start ]
end

# 時間変更した後にcronを再起動しないと時間が反映されない
service 'crond' do
  action :restart
end

