# VPC エンドポイント

# 前段の話

- そもそもの話としてS3等はVPC外
- よって、VPCエンドポイントを使用せずにS3に繋ぐには一旦インターネットに出ていく必要がある

# VPCエンドポイント

- Direct Connectを使用せずに[PrivateLink](https://aws.amazon.com/jp/privatelink/) を使用するAWSサービスやVPCエンドポイントサービスにVPCをプライベートに接続できる
    - 例えば、VPCエンドポイントを使用せずにprivate subnet上のEC2インスタンスがS3にアクセスしたい時には、NAT等を使用してインターネットアクセスする必要があった
    - VPCエンドポイントを使用すれば、private subnet上のEC2インスタンスからインターネットを経由せず、直接アクセスすることができる
    
参考: 

- [S3 VPCエンドポイントを利用するメリット - Qiita](https://qiita.com/SatoHiroyuki/items/b611485b6ec736e9076f)
- [【新機能】S3がVPCのプライベートサブネットからアクセス可能になりました！ ｜ DevelopersIO](https://dev.classmethod.jp/cloud/vpc-endpoint-for-s3/)
- [AWSでS3を使う場合は必ずVPCエンドポイントも作成しておくクセをつけましょう。 | 株式会社ビットクリア](https://www.bitclear.co.jp/vpcendpoint/)
- [【新機能】S3がVPCのプライベートサブネットからアクセス可能になりました！ ｜ DevelopersIO](https://dev.classmethod.jp/cloud/vpc-endpoint-for-s3/)
