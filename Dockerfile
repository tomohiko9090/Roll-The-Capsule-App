# ベースとなるDockerイメージ指定
FROM golang:1.16.3

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work

EXPOSE 8080
