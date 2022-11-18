# Go Rest API

## 構成

- レイヤードアーキテクチャ + DDD
- DB 環境は Docker にて構築
- Push,PL のタイミングで Lint チェック(GitHub Actions)

## 参考図

![DDD](https://i0.wp.com/mintaku-blog.net/mintaku/wp-content/uploads/2020/07/5.png?w=1600&ssl=1)

## API 一覧

| 機能         | メソッド | URI     | 権限 |
| ------------ | -------- | ------- | ---- |
| 文字列を返す | GET      | api/v1/ | なし |

| 機能     | メソッド | URI           | 権限 |
| -------- | -------- | ------------- | ---- |
| 新規登録 | POST     | api/v1/signup | なし |
| ログイン | POST     | api/v1/signin | なし |

| 機能               | メソッド | URI               | 権限 |
| ------------------ | -------- | ----------------- | ---- |
| 習慣を登録する     | POST     | api/v1/create     | 有り |
| 習慣を削除する     | DELETE   | api/v1/delete/:id | 有り |
| 習慣を更新する     | POST     | api/v1/update     | 有り |
| 習慣を全て取得する | POST     | api/v1/get        | 有り |

## 使用準備

1, ルートディレクトリに.env ファイルを用意する。

```shell
    touch .env
```

2, ローカルから MySQL のコンテナを作成

```shell
    docker-compose up -d # コンテナの作成 -d -> バックグランドで実行
    docker exec -it mysql_db bash # コンテナに入る
```

3, 動かす -> "Start Server!"が表示されたら OK

```shell
    go run main.go
```
