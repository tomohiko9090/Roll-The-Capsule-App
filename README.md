# Turtle Gacha API（CA_Tech_Dojo）
 
「Turtle Gacha API」は、日本に生息するカメをモチーフとしたガチャAPIです。
 
# DEMO
例)ガチャ実行APIの場合  
<br>
1. Headerに「token」, Bodyに「times」(ガチャを引きたい回数)をリクエストします。     
<img width="392" alt="スクリーンショット 0003-08-17 午後4 42 55" src="https://user-images.githubusercontent.com/66200485/129684900-d4c290f9-c15a-4c1b-b588-00ae73098c67.png">  
<br>

2. 当たったキャラクターがJSON形式でレスポンスされます。  
<img width="436" alt="スクリーンショット 0003-08-17 午後4 42 15" src="https://user-images.githubusercontent.com/66200485/129684802-f4606add-9058-4471-9369-1f3d69d4b099.png">  
<br>

イメージ画像(この画像は実装されていません。)  
 ![IMG_0223](https://user-images.githubusercontent.com/66200485/128810672-bc73e645-3abb-410c-bc3b-20dd6d759883.JPG)  

# Features

## カメのレア度
- ☆1から☆7まである
- レアなキャラほどガチャを引く確率が低くなるように設定 
- 容易に新しいキャラを追加することができる  

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

## リレーショナルデータベース
![スクリーンショット 0003-08-17 午後3 20 34](https://user-images.githubusercontent.com/66200485/129673939-5c5931d7-4a71-4aa6-845e-02bd718cc696.png)

## MVCモデル
- M : model層  
DBへアクセスしたり、構造体を作成します。  
- V : view = handler層  
tokenからユーザー情報を読み取ったり、paramからIDを取得したりするなど、クライアントのリクエストとレスポンスを行います。　　
- C : controller層  
tokenを生成したり、確率に応じてキャラを引いたりするなど、ビジネスロジックを組み立てます。  

## ディレクトリ構造
<img width="170" alt="スクショ" src="https://user-images.githubusercontent.com/66200485/129654476-34336048-69a7-4f6b-9c58-263aabc9883d.png">

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
- Goではスネークケースではなく、キャメルケースを使う
 
# Author
 
* 作成者 Hikotomo!
