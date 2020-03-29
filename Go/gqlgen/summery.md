# gqlgen まとめ

- Query の引数で Required ではないものに関しては、生成される引数は promitive 型であってもポインタで生成される
	- 例えば `comments(userID: ID, chatRoomID: ID!): [Comment]` のような Query の場合には、userID は string のポインタで生成され、chatRoomID はstring 型で生成される

## Model の定義の仕方のコツ

- 基本的には、自前で構造体として model を定義し、`gqlgen.yml` の `model` の項目ではそれを指定する
    - 構造体で定義する Model には、FK を埋め込んでおく
- `schema.graphql` では、埋め込む
    - それ用の Resolver が生成される
        - [Go + gqlgenを使ったGraphQLアプリケーションサーバーの実装 | MF KESSAI TECH BLOG](https://tech.mfkessai.co.jp/2018/08/go-gqlgen-graphql/)の「子フィールドを取得するResolverを生成させるパターン」の部分
