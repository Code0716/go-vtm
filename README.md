# go-vtm

知人の caffe 用勤怠管理システムを勉強がてら API を作る。

完成予定は未定。
まずは API から作るようにする。

Golang,Echo,Gorm

### 開発手順

基本的に Makefile の手順を実行すれば問題ない

必要なパッケージの導入、インストールする。

.env を調整

`docker compose up -d`で database を立ち上げる

`make up`で server 起動

test は db が立ち上がっている状態で、`make test`コマンドを実行。
