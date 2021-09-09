# Turtle Gacha API （CA-Tech-Dojo）  
<img width="545" alt="スクリーンショット 0003-09-09 午後6 04 29" src="https://user-images.githubusercontent.com/66200485/132656878-c815ce65-2bd9-4d64-a6c6-2f13f22da37a.png">  
「Turtle Gacha API」は、日本に生息するカメをモチーフとしたガチャAPIです。
 
# DEMO
例)ガチャ実行APIの場合  
<br>

1. docker-compose up --buildを実行します。  
![スクリーンショット 0003-09-09 午後9 58 31](https://user-images.githubusercontent.com/66200485/132690086-fafc51bd-1126-48ec-b5a1-b6f21d33c871.png)  
![スクリーンショット 0003-09-09 午後9 56 58](https://user-images.githubusercontent.com/66200485/132689850-abffd748-cc5b-478f-9ac7-fe433f5a1a53.png)

<br>

2. Headerに「token」, Bodyに「times」(ガチャを引きたい回数)を入れてリクエストします。     
<img width="351" alt="スクリーンショット 0003-08-23 午後10 03 16" src="https://user-images.githubusercontent.com/66200485/130452035-bfca7dc4-9ca5-4fc3-b5a8-b376e251d67d.png">
<img width="608" alt="スクリーンショット 0003-08-23 午後10 05 55" src="https://user-images.githubusercontent.com/66200485/130452347-60a97578-c06a-4ac3-a473-af7e8e10b0b7.png">

<br>

3. 当たったキャラクターがJSON形式でレスポンスされます。  
<img width="378" alt="スクリーンショット 0003-08-23 午後10 08 26" src="https://user-images.githubusercontent.com/66200485/130452706-9fbbafa7-966f-4d9a-9e7f-ab902499651f.png">

<br>

 ![IMG_0223](https://user-images.githubusercontent.com/66200485/128810672-bc73e645-3abb-410c-bc3b-20dd6d759883.JPG)  
 イメージ画像  
 (この画像は実装されていません。)  
 (僕のうちで飼っているカメです。)

# Features

## Rarely
- レアなキャラほどガチャを引く確率が低くなるように設定 
- アルゴリズム上、容易に新しいキャラが追加可能  

|キャラ|レア度|攻撃力|防御力|回復力
|:---|:---|:---|:---|:---|
|ミシシッピアカミミガメ|☆|3|3|0|  
|イシガメ|☆☆|1|2|4|
|スッポン|☆☆☆|5|1|2|  
|クサガメ|☆☆☆☆|2|4|3|  
|カミツキガメ|☆☆☆☆☆|7|5|1|  
|リュウキュウヤマガメ|☆☆☆☆☆☆|4|4|7|  
|?|☆☆☆☆☆☆☆|6|6|6|   
 
## Function
1. **ユーザー関連API**  
    1. **ユーザー情報作成API(POST)**  
http://127.0.0.1:8080/user/create  
ユーザ情報を作成します。  
ユーザの名前情報をリクエストで受け取り、ユーザIDと認証用のトークンを生成しデータベースへ保存します。  
    1. **ユーザー情報取得API(GET)**  
http://127.0.0.1:8080/user/get  
ユーザ情報を取得します。  
ユーザの認証と特定の処理はリクエストヘッダのtokenを読み取ってデータベースに照会をします。  
    1. **ユーザー情報更新API(PUT)**  
http://127.0.0.1:8080/user/update  
ユーザ情報の更新をします。  
<br>

2. **ガチャ実行API(POST)**    
http://127.0.0.1:8080/gacha/draw  
ガチャを引いてキャラクターを取得します。  
獲得したキャラクターはユーザ所持キャラクターテーブルへ保存します。  
同じ種類のキャラクターでもユーザは複数所持することができます。  
<br>

3. **ユーザー所持キャラクター一覧取得API(GET)**  
http://127.0.0.1:8080/character/list  
ユーザが所持しているキャラクター一覧情報を取得します。

## RDB
![スクリーンショット 0003-08-29 午後1 26 42](https://user-images.githubusercontent.com/66200485/131238407-957730f7-0f0d-48d9-a491-30e53a401522.png)

# Directory Structure
![スクリーンショット 0003-09-09 午後10 06 55](https://user-images.githubusercontent.com/66200485/132691392-1f106255-a327-4c51-85bb-31a58a76b3b0.png)

## Model-View-Controller
- M : model層  
DBへアクセスしたり、構造体を作成します。  
- V : view = handler層  
tokenからユーザー情報を読み取ったり、paramからIDを取得したりするなど、クライアントのリクエストとレスポンスを行います。　　
- C : controller層  
tokenを生成したり、確率に応じてキャラを引いたりするなど、ビジネスロジックを組み立てます。  

# Requirement
 
**言語**：Golang 1.16.3  
**フレームワーク**：echo v3.3.10  
**開発環境**：MacOS  
**DB**：MySQL  
**ライブラリ**：  
"github.com/go-sql-driver/mysql"  
"github.com/labstack/echo"  

# Installation

Requirementで列挙したライブラリのインストール方法
 
```bash
go get github.com/go-sqlt-driver/mysql
go get github.com/labstack/echo
```
 
# Usage
 
1. このリポジトリをclone
2. TurtleGachaAPI_MVCmodelのディレクトリに移動
3. main.goを実行
4. HeaderやBodyにKeyとValueを入れてリクエスト

```bash
git clone https://github.com/tomohiko9090/CA_Tech_Dojo.git
cd CA-Tech-Dojo
docker-compose up --build
```
 
# Memo

 綺麗なコーディングを行う上で教わったアドバイスをメモしていきます。
 
## Golang(API)
- 基本的に１単語でファイル名を付けること
- ファイルの１文字目に大文字は使わないこと
- handlerでmodelをよむの×(処理の流れ hendler -> controller -> model　とし、hendlerからmodelに跨がないようにする)
- InsertUser, CreateUserのようにわかりやすい関数名にする
- 「panic」にするとサーバーが落ちてしまうため、エラーハンドリングを行う
- Golandの設定でwatchのgo fmt とgo imports で自動でリフォーマットできるようにする
- 頭文字が大文字のものはPublicとなり、外部packageから参照が可能になる。また、小文字のものはPrivateとなり、外部packageから参照が不可能になる。
- modelでグローバル変数を使いますのはアンチパターン
- Goではスネークケースではなく、キャメルケースを使う
- エラー時、ステータスコードは重要であるため必ず行う
- エラーログはfmt出力でなくlog出力を使用すること(logならファイル保存可能)
- 配列を使用したfor文では、forrを使用すること

## MySQL(DB)
- APIでは一度発行したDBコネクションを使い倒す(毎回接続するのはアンチパターン)
- idなどで,通し番号を付けたい時は、「AUTO_INCREMENT」を使用する。
- UserテーブルとCharacterテーブルの中間テーブルはUserCharacterテーブルという名称になる。

## GitHub
- マスターブランチ(pullリク)は、コメントアウトは少なくし、第三者がコードをみる時に必要なもののみにする。コメントアウトがたくさん入ったものはデベロップブランチへ。

## Docker
- docker-composeに書く「ports」について
"(ホスト側のポート)：(コンテナ側ポート)"  
(i) 自分のターミナルからコンテナにアクセスする場合   
→ホスト側のポート番号を使う  
(ii) コンテナからコンテナ にアクセスする場合  
→コンテナ側のポート番号を使う  
- Dockerファイルに書いてあるFROM〜  
Dockerファイルでは、層を重ねて自分独自のイメージを作成しているが、FROM〜によって、最初に層の土台となるDockerイメージをDockerHubからダウンロードしている。  
作成されたDockerイメージは、DockerHubにあげることもできる。

# KeyWords
マッピング : 関連付けを行うこと  
キャスティング : データ型を別の型に変換すること  
リファクター : 内部構造は変えながらもアウトプットは同じにすること  
アッパーキャメルケース : 1文字目大文字  
ローワーキャメル : 1文字目小文字  
 
# Next
1. Dockerでの環境構築
2. トランザクションを意識した開発にする

# Author
 
* 作成者 Hikotomo!
