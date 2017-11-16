# ec_product_checker

ECサイトを監視して商品が購入可能な状態になったら通知するツール。

Goに慣れるための習作として作成。

# 動作環境

chromeとchromedriverをインストールしている環境で動かしてください。  
MacOS X 10.11、CentOS7.3, Ubuntu16.04にて動作確認済み。

また通知処理にAWS SDKを使っているので以下の環境変数を設定する。

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

参考: https://docs.aws.amazon.com/ja_jp/sdk-for-go/v1/developer-guide/configuring-sdk.html


# config

## General

### TopicArn

AWS SNSを使って通知するので通知先のtopic arnを設定する。
例: arn:aws:sns:ap-northeast-1:xxxxxxxxxxxxxx:topic-name

## Site

監視先サイトの設定。複数設定可。

### Name

サイト名。通知時にしか使わないので適当で良い。

### URL

監視先のURL。商品販売ページのURLを設定。

### Selecter

販売時と売り切れ時で表示が変わる箇所のjQuery Selecterを設定。

### SoldOutMessage

Selecterで設定した箇所が *売り切れの際に* 表示するメッセージをそのまま記載する。  
プログラム中で比較処理に使われるので間違えると正しく動作しません。


