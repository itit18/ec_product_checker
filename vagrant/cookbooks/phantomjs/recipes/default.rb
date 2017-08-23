#
# Cookbook Name:: phantomjs
# Recipe:: default
#

# phantomJS用のyumリポジトリ追加
bash "add phantomJS repo" do
  code   "rpm -ivh http://repo.okay.com.mx/centos/6/x86_64/release/okay-release-1-1.noarch.rpm"
  action :run
  not_if { ::File.exists?("/etc/yum.repos.d/okay.repo") }
end

# 依存しているライブラリをインストール
pkg_hash = %w(libxslt libxslt-devel)
pkg_hash.each do |pkg|
  package pkg do
    action :install
    options '--enablerepo=epel,remi'
  end
end

pkg_hash = %w(libxml2 libxml2-devel)
pkg_hash.each do |pkg|
  package pkg do
    action :install
    options '--enablerepo=epel,remi'
  end
end

package 'phantomjs' do
  action :install
  options '--enablerepo=epel,remi'
end
