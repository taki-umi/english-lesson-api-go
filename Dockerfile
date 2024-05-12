# Goの公式イメージを使用
FROM golang:1.18

# 作業ディレクトリを設定
WORKDIR /app

# 依存関係ファイルをコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o /hello-api

# アプリケーションを実行
CMD ["/hello-api"]
