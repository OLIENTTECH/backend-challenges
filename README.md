# 📝 バックエンド採用課題

## 開発環境の構築

本リポジトリの推奨される Go のバージョンは以下の通りです。<br>
インストールについてはこちらの[公式サイト](https://go.dev/doc/manage-install)を参照してください。

```shell
v1.21.7
```

プロジェクトを開始する前に下記を実行してください。
lefthookがインストールされていない場合は、brewやこちらの[github](https://github.com/evilmartians/lefthook)を参考にしてlefthookをインストールしてください。

```shell
lefthook install
```

`.env.local` ファイルを作成し、`.env.example`の内容を設定してください。

```shell
cp .env.example .env.local
```

Docker コンテナをビルドしてください。

```shell
make build
```

## 開発

下記を実行すると、Docker コンテナの起動や停止を行えます。

```shell
# launching docker containers
make up
# down docker containers
make down
```

### ヘルスチェック

コンテナ起動後、以下のコマンドを実行して Go の開発環境準備ができているか確認してください。

```shell
curl http://localhost:8080/health
```

成功すると、OK が返されます。

### DB の準備

PostgreSQL のマイグレーションや、ロールバック、シードデータの挿入は下記のコマンドで実行できます。

```shell
# create tables to the latest version
make migrate-up
# seed mock data
make migrate-seed
# reset tables and seed
make migrate-reset
# rollback all migrations
make migrate-drop
```

詳細は [GOOSE](https://github.com/pressly/goose) を参照してください。

### マイグレーションファイルの作成

マイグレーションファイルや、シードファイルを作成するには下記を実行してください。
実行する際は、`FILE` にファイル名を指定してください。（{} は実行の際は必要ないので注意してください。）

```shell
# create migration file (DDL)
make create-ddl FILE={filename}
# create migration file (DML)
make create-dml FILE={filename}
```

## ユニットテスト

バックエンドのユニットテストを実行するには下記を実行してください。

```shell
make test
```

### モックの作成

テスト用のモックを生成するには下記を実行してください。

```shell
make mockgen source=xxx destination=xxx filename=xxx
```

## Lint
Goファイルのlintを実行するには下記を実行してください。

```shell
make lint
```

## 開発における注意点

- ユースケース層においては必ずユニットテストを実施
- pkgディレクトリ内のコードは今後別のPJで扱うことも考慮してできるだけテストを実施すること
- ドメイン層やユーザインターフェース層においては、可能な限りテストを実施すること

## Architecture

このプロジェクトは、 [レイヤードアーキテクチャ](https://qiita.com/kichion/items/aca19765cb16e7e65946) + [DDD](https://zenn.dev/hisamitsu/articles/2937fc4dd9bd4c)を採用しています。

## Cording Style

コーディングルールとAPIデザインルールは、以下に従うものとします。

- https://sprinkle-hawk-750.notion.site/API-Cording-Style-WIP-b9a2793723db495387a8c29d5dbb6e29
