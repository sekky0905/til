# GraphQL まとめ

## メモ

### オペレーションの定義とクエリ(query ではなく、オペレーションへの指示という意味で)

- Mutation みたいな大文から始まるやつは、ルートのオペレーション型
    - つまりオペレーションの定義
    - この Mutation type の field にオペレーションを定義する
- mutation みたいな小文字から始まるやつは、クエリ
    - つまり、Mutation で定義したオペレーションへの指示
- 参考: [「GraphQL」徹底入門 ─ RESTとの比較、API・フロント双方の実装から学ぶ - エンジニアHub｜若手Webエンジニアのキャリアを考える！](https://employment.en-japan.com/engineerhub/entry/2018/12/26/103000)

## クエリに対するコストという考え方

- 取得するデータがツリー状になっているため、クライアントが無限にデータを取得できるようにし、サーバがそれに答えようとすると死ぬ
    - ツリーの上流から下流までの合計コストを計算する仕組みを検討すること
    - 参考: [GraphQLサーバをGo言語で作る - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1055228)


## エラーについての考え方

- GraphQL はあくまでも Query Language
    - クライアントから送信されてくるクエリが `syntax error` の場合以外は、エラーなく返すべき
        - DB と SQL 的な考え方
        - REST API は RPC で GraphQL は Query Language だと考えるとそもそも何がエラーなのかという考えが異なる
    - 参考: [GraphQLサーバをGo言語で作る - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1055228)

## クライアント側のクエリってどういう風に考えるのか

- 各々の Component が自分に必要なデータを [Fragment](https://qiita.com/k-boy/items/079d1d4418dc11863a0e) として定義して、上位の Component に渡す
    - 参考: [GraphQLサーバをGo言語で作る - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1055228)

## ツリー構造で取得する考え方が肝

- 定義された関係を任意の深さまで取得することができる
    - 参考: [GraphQLスキーマ設計ガイド - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1576562)
    
## Resolver の考え方

- 各々の要素に対する Resolver の集合が、Query への Resolver になる
- Scheme に対して、全部の要素に Resolver を用意する必要はない( name, age にそれぞれ Resolver を定義するみたいな)
    - 別データの ID に対しては、その ID から別データに変換んするような Resolver は必要
        - ex) `Hoge.FooID` に対しては、 `FooID` を受け取り、 `Foo` のデータを取得してくる `Resolver` が必要
    - 参考: [GraphQLスキーマ設計ガイド - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1576562)

## URL の設計

- GraphQL サーバへのアクセスは単一エンドポイントで、Query でデータの取得を行うが、SPA の場合フロントエンド側で表示する URL については従来のような表示の仕方を取り得ることは大いにあり得る 

## Reference

### 全般系

- [「GraphQL」徹底入門 ─ RESTとの比較、API・フロント双方の実装から学ぶ - エンジニアHub｜若手Webエンジニアのキャリアを考える！](https://employment.en-japan.com/engineerhub/entry/2018/12/26/103000)

### 設計

- [はじめてのGraphQLスキーマ設計 - Speaker Deck](https://speakerdeck.com/rikuson298/hazimetefalsegraphqlsukimashe-ji)
- [GraphQL API設計で気をつけること - Qiita](https://qiita.com/wawoon/items/e24398af912d9f3e0389)
- [GraphQLサーバをGo言語で作る - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1055228)
- [GraphQLスキーマ設計ガイド - わかめの自動売り場 - BOOTH](https://booth.pm/ja/items/1576562)

### GraphQL with Go

- [Using GraphQL with Microservices in Go - Outcrawl](https://outcrawl.com/go-graphql-gateway-microservices/)

#### gqlgen

- [Building GraphQL servers in golang — gqlgen](https://gqlgen.com/getting-started/)
- [Go + gqlgenを使ったGraphQLアプリケーションサーバーの実装 | MF KESSAI TECH BLOG](https://tech.mfkessai.co.jp/2018/08/go-gqlgen-graphql/)

### Dataloader 

- [GraphQL と N+1 SQL 問題と dataloader - Qiita](https://qiita.com/yuku_t/items/2c1735cbf45e75c0bfb8)
- gqlgen での実装
    [Optimizing N+1 database queries using Dataloaders — gqlgen](https://gqlgen.com/reference/dataloaders/)
