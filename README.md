# MyGoCv

golangのコンテナイメージをもとに、gocvをインストールして使える状態の新たなコンテナイメージを作成・実行します。

## 実行環境について

諸事情により、プロキシ環境下での利用を想定して記述しています。  
非プロキシ環境下での利用時は、以下の点の適宜変更が必要です。

### .env

プロキシに関する環境変数を無効化(アンコメントや削除)

### Dockerfile

```
- COPY ./apt.conf /etc/apt/apt.conf
+ COPY ./apt.conf /etc/apt/apt.conf
```

```
- # RUN go get -u -d gocv.io/x/gocv
- # RUN cd $GOPATH/src/gocv.io/x/gocv
- # RUN make install
+ RUN go get -u -d gocv.io/x/gocv
+ RUN cd $GOPATH/src/gocv.io/x/gocv
+ RUN make install
```

## How to use

### 1. 起動

```
$ docker-compose up -d
```

### 2. gocvインストール(Proxy環境化の場合のみ)

```
$ docker-compose exec app sh
% go get -u -d gocv.io/x/gocv
% cd $GOPATH/src/gocv.io/x/gocv
% make install
% [Ctrl+p,Ctrl+q] (コンテナのデタッチ)
```

### 3. コンテナイメージのコミット(Proxy環境化の場合のみ)

```
$ docker commit gocv_app_1 mygocv
```

### 4. 利用コンテナイメージの変更(Proxy環境化の場合のみ)

#### Dockerfile

```
- FROM golang:latest
+ FROM mygocv:latest
```

### 5. goやgocvの実行

コンテナが起動してしまえばshが常駐しているだけなので、  
以下のいずれかで実行します。

```
$ docker-compose exec app 実行したいコマンド

or

$ docker-compose exec app sh
% 実行したいコマンド
```

