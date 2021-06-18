# API連携

## 認証について

いくつかの認証方式がある。
[Salesforceで適切なOAuthフローを選べるようするには - TerraSkyBase | テラスカイを支える人とテクノロジーの情報を発信する基地局](https://base.terrasky.co.jp/articles/ecEXW) を参考にして判断する

### JWT

[サーバ間インテグレーション用の OAuth 2.0 JWT ベアラーフロー](https://help.salesforce.com/articleView?id=sf.remoteaccess_oauth_jwt_flow.htm&type=5)
[OAuthのJWTベアラーフローを用いてSalesforceに接続する - TerraSkyBase | テラスカイを支える人とテクノロジーの情報を発信する基地局](https://base.terrasky.co.jp/articles/iRj2p)

## ユーザーについて

APIで外部システムと連携する際には、API用のユーザーを作成するのが良さそう
[安全な Salesforce API ユーザの作成](https://help.salesforce.com/articleView?id=000331470&type=1&mode=1)

## その他

### 環境変数に格納する

pem等を環境変数に格納したいケースもまあまあある。
そんな場合には、以下のような手順で行うと良い。
- 一行化して環境変数に格納する
- メモリ上に持ってきた時に、改行を復活させる

refs: [[RS256] JWTでRSA秘密鍵を環境変数で処理したい [Javascript] - Qiita](https://qiita.com/sho7650/items/1dd65a1db785f902a2d6)
