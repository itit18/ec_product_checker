#
# Cookbook Name:: nodejs
# Recipe:: gulp
#

bash 'install gulp' do
  code <<-EOH
    npm install -g gulp
  EOH
  user 'root'
end
