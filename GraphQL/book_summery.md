# GraphQL まとめ

## 『初めてのGraphQL ――Webサービスを作って学ぶ新世代API』 を読んで

### 「第1章 GraphQL へようこそ」のまとめ

- GraphQL の利点
    - REST みたいに過剰にデータを取得しすぎることがなくなる
        - 必要なデータだけを取得できる
    - エンドポイントを統合できる
    
### 「第3章 GraphQL の問合せ言語」のまとめ

- 問合せコマンド
    - Query 
        - データの取得コマンド
    - Mutation
        - データの変更コマンド
    - Subscription
        - socket 通信を使ってデータの変更監視コマンド

- エラーハンドリング
    - 問合せでエラーが発生した場合には、errors と言うキーが結果に含まれることになる(通常時は data)
        - errors にエラーの詳細は含まれる

- Query Document
    - 問合せに使用する GraphQL クエリ言語で記述されたテキスト

- ルート型
    - operation の GraphQL の型
    - クエリはルート型

- フィールド
    - クエリするフィールドは、API のスキーマで定義できる
    - 選択セット
        - 必要なフィールドは、 `{}` で指定可能
    - エイリアス付与可能
    - クエリ引数
        - 結果をフィルタリングできる
        
#### 「3.2.1 エッジと接続」のまとめ

- Scalar 型
    - primitive 型みたいなもの
        - 選択セットでの葉
- Object 型
    - 1つ以上のフィールドの集合

### 参考文献

Eve Porcello、Alex Banks　著、尾崎 沙耶、あんどうやすし　訳 (2019/11/13)『初めてのGraphQL ――Webサービスを作って学ぶ新世代API』オライリージャパン


## その他まとめ

### リソースの表し方について

- REST API ではリソースを URL で表していた
    - 例 
        - `/rooms/{id}/comments`
- GraphQL では、ペイロードに付与するスキーマによってリソースを表す
    - 例
         ```
         {
            rooms {
                id
                name
            }
        }
        ```
    - URL のエンドポイントは1つ
        - 例
            - `/`

- 参考
    - [GraphQL を RESTful API と比較しながら実装して理解する - JX通信社エンジニアブログ](https://tech.jxpress.net/entry/graphql-vs-rest)
