# AWS Account、Group、Organization周り

## Account系についてまとめる

### AWS Account(root user)

- 最初に作成するAccount
- 完全なアクセス権限
- 日常的に使用してはダメ

[AWS アカウントのルートユーザー - AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_root-user.html)

### IAM 

- AWS のサービスやリソースへのアクセスを管理する機能
- 認証と許可を制御する

参考: [IAM の詳細を理解します。 - AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/intro-structure.html)

#### IAM User

- AWS Accountから払い出される個別のIDを使用する
- 実際に使用するときには、こちらを使用する
- このIAM user毎に権限を付与できる
- 人間だけじゃなく、AWSにアクセスするアプリケーション等がIAM Userになり得る

参考: [IAM ユーザー - AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_users.html)

#### IAM Group

- 複数のIAM userをまとめたもの
- このgroup単位でアクセス権限を管理することができる
- 例えばAdmin groupとか

参考: 
[IAM グループ - AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_groups.html)
[非エンジニアにこそ知ってほしいAWSのアカウント管理 ｜ DevelopersIO](https://dev.classmethod.jp/cloud/aws/aws-beginner-account/)

#### IAM Role

- リーソースに対して、指定したアクセス権限を持つ
- roleに対して与えた権限を持った一時キーをIAM UserやAWS上のサービスに対して渡すことができる

参考: 
[IAMロール徹底理解 〜 AssumeRoleの正体 ｜ DevelopersIO](https://dev.classmethod.jp/cloud/aws/iam-role-and-assumerole/)
[ロールに関する用語と概念 - AWS Identity and Access Management](https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_roles_terms-and-concepts.html)

#### policy

- > 「誰が」「どのAWSサービスの」「どのリソースに対して」「どんな操作を」「許可する(許可しない)」、といったことをJSON形式で記述します。
- > 記述したポリシーをユーザー(IAMユーザー、IAMグループ、IAMロール)や、AWSリソースに関連づけることで、アクセス制御を実現しています。

引用元: [AWS IAMポリシーを理解する ｜ DevelopersIO](https://dev.classmethod.jp/cloud/aws-iam-policy/)

### Organization

- 複数のAWSアカウント(root user)を一限管理するのに使用する
- > AWS Organizations には、お客様のビジネスの予算、セキュリティ、コンプライアンスのニーズをより適切に満たすアカウント管理および一括請求機能が備わっています
    - 引用元: [AWS Organizations とは何ですか? - AWS Organizations](https://docs.aws.amazon.com/ja_jp/organizations/latest/userguide/orgs_introduction.html)

参考: [AWS Organizations を実際に初めてみる第一歩 - Qiita](https://qiita.com/akiray03/items/07e9a786d1cc054f4e22)


#### root

- OUの一番上のものみたいに考えれば良い

#### organization unit(OU)

- 複数のアカウントをまとめるコンテナ
- OUの中にOUと言った形で、再帰的な階層を作ることができる


参考: [AWS Organizations の用語と概念 - AWS Organizations](https://docs.aws.amazon.com/ja_jp/organizations/latest/userguide/orgs_getting-started_concepts.html)
[AWS Organizationsによるマルチアカウント戦略とその実装 - クラウドワークス エンジニアブログ](https://engineer.crowdworks.jp/entry/2018/07/17/103453)
