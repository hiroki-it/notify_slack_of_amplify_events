# notify-slack-of-amplify-events

## 概要

EventBridgeから転送されたAmplifyイベントに情報を追加し，Slackに転送します．

Goは開発者に実装方法を強制させられるため，可読性が高く，後続者が開発しやすいです．

また，静的解析のルールが厳しいため，バグを事前に見つけることができます．

そのため，実装がスケーリングしやすく，特にバグが許されないような基盤部分に適しています．

今回，これらには当てはまりませんが，Goの環境構築からデプロイまでの一連の流れを俯瞰するために，Goを採用しました．

## ディレクトリ構造

本リポジトリのディレクトリ構造は以下の通りに構成しております．

cmdディレクトリの構成は，クリーンアーキテクチャを参考にいたしました．

参考：https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

```
notify-slack-of-amplify-events
├── build   # ビルド処理
├── cmd     # エントリポイントを含む処理
|   ├── controllers # コントローラ
|   ├── entities    # エンティティ
|   └── usecases    # ユースケース
|
└── test
    ├── mock     # モック処理
    ├── testdata # テストデータ（JSON）
    └── unit     # ユニットテスト処理
```

## 環境構築

### 1. RIEのインストール

ローカルで統合テストを行うために，RIEをインストールする必要があります．

RIEにより，ローカルで擬似的にLambdaを再現できます．

いくつか方法が用意されており，Goソースコードの再利用性の観点から，ホストPCにRIEをインストールする方法を採用いたしました．

ホストPCで以下のコマンドを実行します．

```shell
$ mkdir -p ~/.aws-lambda-rie
$ curl -Lo ~/.aws-lambda-rie/aws-lambda-rie https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie
$ chmod +x ~/.aws-lambda-rie/aws-lambda-rie
```

その他のインストール方法につきましては，以下を参考に．

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html#go-image-other

### 2. ビルドイメージ & コンテナ構築

Dockerfileからイメージをビルドし，コンテナを構築します．

イメージのビルドにおいては，マルチステージビルドを採用しております．

```shell
$ docker-compose run -d --rm notify-slack-of-amplify-event
````

### 4. モジュールのインストール

コンテナで，アプリケーションで使用するモジュールをインストールします．

```shell
$ docker-compose run --rm  notify-slack-of-amplify-events go mod download -x
```

## リロード

### ホットリロード

ホスト側のソースコードを修正に合わせて，コンテナ側のアーティファクトを随時更新します．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

```shell
$ docker-compose run --rm notify-slack-of-amplify-events air -c .air.toml
```

### RIEにソースコードを再反映

LambdaのRIEにソースコードを反映するためには，イメージを再ビルドする必要があります．

```shell
$ docker-compose up --build -d
```

## テスト

### ユニットテスト

Goサービスのユニットテストを実行します．

```shell
$ docker-compose run --rm notify-slack-of-amplify-events go test -v -cover ./test/unit/...
```

### 統合テスト

GoサービスからLambdaサービスにPOSTリクエストを送信し，一連の処理をテストします．

ローカルでLambdaの擬似的な再現するために，RIEを使用いたしました．

```shell
$ docker-compose run --rm notify-slack-of-amplify-events go test -v -cover ./test/integration/...
```

## デプロイ

原則として，ローカルPCからソースコードをLambdaにデプロイしないようにします．

CircleCIによるCDにて，これをLambdaにデプロイします．

デプロイツールとして，[Serverless Framework](https://github.com/serverless/serverless) を使用いたしました．

## Amplifyについて

Amplifyを用いて，SSGアプリのCI/CDを構築します．

AmplifyがS3にSSGアプリをデプロイし，このイベントをEventBridgeがキャッチします．

AmplifyのCI/CDにつきましては，[こちらのリポジトリ](https://github.com/hiroki-it/deploy-ssg-to-amplify) を参考に．
