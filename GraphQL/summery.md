# GraphQL まとめ

## メモ

### オペレーションの定義とクエリ(query ではなく、オペレーションへの指示という意味で)

- Mutation みたいな大文から始まるやつは、ルートのオペレーション型
    - つまりオペレーションの定義
    - この Mutation type の field にオペレーションを定義する
- mutation みたいな小文字から始まるやつは、クエリ
    - つまり、Mutation で定義したオペレーションへの指示
- 参考: [「GraphQL」徹底入門 ─ RESTとの比較、API・フロント双方の実装から学ぶ - エンジニアHub｜若手Webエンジニアのキャリアを考える！](https://employment.en-japan.com/engineerhub/entry/2018/12/26/103000)

## Reference

### 全般系

- [「GraphQL」徹底入門 ─ RESTとの比較、API・フロント双方の実装から学ぶ - エンジニアHub｜若手Webエンジニアのキャリアを考える！](https://employment.en-japan.com/engineerhub/entry/2018/12/26/103000)

### 設計

- [はじめてのGraphQLスキーマ設計 - Speaker Deck](https://speakerdeck.com/rikuson298/hazimetefalsegraphqlsukimashe-ji)
- [GraphQL API設計で気をつけること - Qiita](https://qiita.com/wawoon/items/e24398af912d9f3e0389)

### GraphQL with Go

- [Using GraphQL with Microservices in Go - Outcrawl](https://outcrawl.com/go-graphql-gateway-microservices/)

#### gqlgen

- [Building GraphQL servers in golang — gqlgen](https://gqlgen.com/getting-started/)
- [Go + gqlgenを使ったGraphQLアプリケーションサーバーの実装 | MF KESSAI TECH BLOG](https://tech.mfkessai.co.jp/2018/08/go-gqlgen-graphql/)


