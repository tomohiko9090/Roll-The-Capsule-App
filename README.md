# Turtle Gacha API（CA_Tech_Dojo）
 
「Turtle Gacha API」は、日本に生息するカメをモチーフとしたキャラクターのガチャAPIです。
 
# DEMO
 ![IMG_0223](https://user-images.githubusercontent.com/66200485/128810672-bc73e645-3abb-410c-bc3b-20dd6d759883.JPG)  
キャラクターNo.1   
名前 :ミドリガメ  
レア度 ：☆  
攻撃力 ：3  
防御力 ：3  
回復力 :0  

# Features

## カメのレア度
- ☆1から☆7まである
- レアなキャラほどガチャを引く確率が低くなるように設定 
- MySQL Workbenchなどを使用し、新しいキャラを追加することも可能  

|キャラ|レア度|攻撃力|防御力|回復力
|:---|:---|:---|:---|:---|
|ミドリガメ|☆|3|3|0|  
|イシガメ|☆☆|1|2|4|
|スッポン|☆☆☆|5|1|2|  
|クサガメ|☆☆☆☆|2|4|3|  
|カミツキガメ|☆☆☆☆☆|7|5|1|  
|リュウキュウヤマガメ|☆☆☆☆☆☆|4|4|7|  
|?|☆☆☆☆☆☆☆|6|6|6|   
 
## 機能
1. **ユーザー関連API**  
1.1. **ユーザー情報作成API(POST)**  
http://127.0.0.1:8080/user/create  
ユーザ情報を作成します。  
ユーザの名前情報をリクエストで受け取り、ユーザIDと認証用のトークンを生成しデータベースへ保存します。  
1.2. **ユーザー情報取得API(GET)**  
http://127.0.0.1:8080/user/get  
ユーザ情報を取得します。  
ユーザの認証と特定の処理はリクエストヘッダのtokenを読み取ってデータベースに照会をします。  
1.3. **ユーザー情報更新API(PUT)**  
http://127.0.0.1:8080/user/update  
ユーザ情報の更新をします。  
1. **ガチャ実行API(POST)**  
http://127.0.0.1:8080/gacha/draw  
ガチャを引いてキャラクターを取得します。  
獲得したキャラクターはユーザ所持キャラクターテーブルへ保存します。  
同じ種類のキャラクターでもユーザは複数所持することができます。  
1. **ユーザ所持キャラクター一覧取得API(GET)**  
http://127.0.0.1:8080/character/list  
ユーザが所持しているキャラクター一覧情報を取得します。

## MVCモデル
- M : model層  
DBへアクセスしたり、構造体を作成するところ  
- V : view = handler層  
tokenからユーザー情報を読み取ったり、paramからidを取得したりするなど、クライアントのリクエストとレスポンスを行うところ  
- C : controller層  
tokenを作ったり、確率に応じてキャラを引いたりするなどのビジネスロジックを組み立てるところ  

## ディレクトリ構造
<img width="348" alt="スクリーンショット 0003-08-06 午前11 47 15" src="https://user-images.githubusercontent.com/66200485/128448965-d7221aab-bba7-4bb4-9451-61a886ca71d8.png">

# Requirement
 
**言語**：Golang 1.16.3  
**フレームワーク**：echo v3.3.10  
**開発環境**：MacOS
**DB**：MySQL  
**ライブラリ**：  
"github.com/go-sql-driver/mysql"  
"github.com/labstack/echo"  

# Installation

Requirementで列挙したライブラリなどのインストール方法を説明する
 
```bash
go get github.com/go-sqlt-driver/mysql
go get github.com/labstack/echo
```
 
# Usage
 
1. このリポジトリをcloneする
2. TurtleGachaAPI_MVCmodelのディレクトリに移動する
3. go run main.goする
4. Postmanなどのアプリを使って、HeaderやBodyにKeyとValueを入れてリクエストする

```bash
git clone https://github.com/tomohiko9090/CA_Tech_Dojo.git
cd TurtleGachaAPI_MVCmodel
go run main.go
```
 
# Memo
 綺麗なコーディングを行う上で教わったアドバイスをメモしていきます。
- 基本的に１単語でファイル名を付けること
- ファイルの１文字目に大文字は使わないこと
- APIでは一度発行したDBコネクションを使い倒す(毎回接続するのはアンチパターン)
- handlerでmodelをよむの×(処理の流れ hendler -> controller -> model　とし、hendlerからmodelに跨がないようにする)
- InsertUser, CreateUserのようにわかりやすい関数内にする
- 「panic」にするとサーバーが落ちてしまうため、エラーハンドリングを行う
- Golandの設定でwatchのgo fmt とgo imports で自動でリフォーマットできるようにする
- 頭文字が大文字のものはPublicとなり、外部packageから参照が可能になる。また、小文字のものはPrivateとなり、外部packageから参照が不可能になる。
 
# Author
 
* 作成者 Hikotomo!
