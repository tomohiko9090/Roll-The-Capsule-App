# CA Tech Dojo
## 制作物名：「カメガチャ」  


↑僕の家に住むかめさん

## ディレクトリ構造
<img width="348" alt="スクリーンショット 0003-08-06 午前11 47 15" src="https://user-images.githubusercontent.com/66200485/128448965-d7221aab-bba7-4bb4-9451-61a886ca71d8.png">

## MVCモデル
- M : model層  
DBへアクセスしたり、構造体を作成するところ  
- V : view = handler層  
tokenからユーザー情報を読み取ったり、paramからidを取得したりするなど、クライアントのリクエストとレスポンスを行うところ  
- C : controller層  
tokenを作ったり、確率に応じてキャラを引いたりするなどのビジネスロジックを組み立てるところ  
 
## 機能一覧
1. ユーザー関連API  
    1.1. ユーザー情報作成API(POST) http://127.0.0.1:8080/user/create  
    1.2. ユーザー情報取得API(GET) http://127.0.0.1:8080/user/get  
    1.3. ユーザー情報更新API(PUT) http://127.0.0.1:8080/user/update  
1. ガチャ実行API(POST) http://127.0.0.1:8080/gacha/draw  
1. ユーザ所持キャラクター一覧取得API(GET) http://127.0.0.1:8080/character/list  

## カメのレア度
- ☆1から☆7まである
- レアなキャラほどガチャを引く確率が低くなるように設定
- MySQLより、新しいキャラを追加することも可能   
 
|キャラ|レア度|  
|:---|:---|  
|ミドリガメ|☆|  
|イシガメ|☆☆|  
|スッポン|☆☆☆|  
|クサガメ|☆☆☆☆|  
|カミツキガメ|☆☆☆☆☆|  
|リュウキュウヤマガメ|☆☆☆☆☆☆|  
|?|☆☆☆☆☆☆☆|  

## アプリの使用方法
1. このリポジトリをclone
2. 以下を実行
以下を実行
```
工事中
```

## 使用技術詳細
**言語**：Golang 1.16.3  
**フレームワーク**：echo v3.3.10  
**開発期間**:4週間  
**開発環境**：MacOS, Docker  
**外部サーバー(OS)**：vultur(ubuntu)  
**DB**：MySQL  





# カメさんガチャAPI（CA_Tech_Dojo）
 
カメさんガチャAPIは、日本に生息するカメをモチーフとしたキャラクターのガチャを行うことができるAPIです。
 
# DEMO

<example>
No.1
![IMG_0223](https://user-images.githubusercontent.com/66200485/127758779-ed5f97f5-f406-414d-b137-75d913fcff27.JPG)  
名前:ミドリガメ
レア度：☆
攻撃力：3
防御力：3
回復力:0

# Features
 
"hoge"のセールスポイントや差別化などを説明する
 
# Requirement
 
"hoge"を動かすのに必要なライブラリなどを列挙する
 
* huga 3.5.2
* hogehuga 1.0.2
 
# Installation
 
Requirementで列挙したライブラリなどのインストール方法を説明する
 
```bash
pip install huga_package
```
 
# Usage
 
DEMOの実行方法など、"hoge"の基本的な使い方を説明する
 
```bash
git clone https://github.com/hoge/~
cd examples
python demo.py
```
 
# Note
 
注意点などがあれば書く
 
# Author
 
作成情報を列挙する
 
* 作成者
* 所属
* E-mail
 
# License
ライセンスを明示する
 
"hoge" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
 
社内向けなら社外秘であることを明示してる
 
"hoge" is Confidential.
