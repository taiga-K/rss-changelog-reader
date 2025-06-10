# rss-changelog-reader

GitHub ChangelogのRSSを購読し、当日投稿分のみをフィルタリングして取得するGoアプリケーション

## 機能

- [gofeed](https://github.com/mmcdole/gofeed)ライブラリを使用したRSSパースィング
- GitHub Copilot Changelogから当日投稿された記事のみを抽出
- シンプルなコマンドライン実行

## 使用方法

```bash
# ビルド
go build .

# 実行
./rss-changelog-reader
```

## テスト

```bash
go test -v
```

## 取得対象

- RSS URL: `https://github.blog/changelog/label/copilot/feed/`
- フィルタ条件: 実行日（当日）に投稿された記事のみ

## 依存関係

- [gofeed](https://github.com/mmcdole/gofeed) - RSSパーサーライブラリ
