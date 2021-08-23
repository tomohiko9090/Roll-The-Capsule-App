# Turtle Gacha API （CA-Tech-Dojo）
 
「Turtle Gacha API」は、日本に生息するカメをモチーフとしたガチャAPIです。
 
# DEMO
例)ガチャ実行APIの場合  
<br>

1. main.goを実行します。  
<img width="260" alt="スクリーンショット 0003-08-23 午後10 14 53" src="https://user-images.githubusercontent.com/66200485/130453561-33a18e22-d06f-4286-81e5-114ebcbf7692.png">

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

# Features

## Rarely
- ☆1から☆7まである
- レアなキャラほどガチャを引く確率が低くなるように設定 
- 容易に新しいキャラが追加可能  

|キャラ|レア度|攻撃力|防御力|回復力
|:---|:---|:---|:---|:---|
|ミドリガメ|☆|3|3|0|  
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
![スクリーンショット 0003-08-17 午後3 20 34](https://user-images.githubusercontent.com/66200485/129673939-5c5931d7-4a71-4aa6-845e-02bd718cc696.png)

# DirectoryStructure
<img width="266" alt="スクリーンショット 0003-08-23 午後9 52 49" src="https://user-images.githubusercontent.com/66200485/130450575-6b337bd6-afea-4bd5-8ee4-2425ff854ac8.png">  

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

# KeyWords
マッピング : 関連付けを行うこと  
キャスティング : データ型を別の型に変換すること  
リファクター : 内部構造は変えながらもアウトプットは同じにすること  
アッパーキャメルケース : 1文字目大文字  
ローワーキャメル : 1文字目小文字  
 
# Author
 
* 作成者 Hikotomo!
