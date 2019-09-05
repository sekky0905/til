# AWS Budgets

- これまでAWSの使用料金がしきい値を超えたら、アラートを通知するというようなことを行うには以下の設定を行なっていた
    - 請求アラートの有効化
    - 請求アラームの作成(CloudWatch)
    - 参考: [AWS の予想請求額をモニタリングする請求アラームの作成 - Amazon CloudWatch](https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html)
- 「請求アラートの有効化」に関しては、一度オンにすると、オフに切り替えられない
- そこで、[AWS Budgets](https://aws.amazon.com/jp/aws-cost-management/aws-budgets/)を使用するともっとシンプルに設定した予算のしきい値を超えるとアラートを発信ということが可能になる
    - 設定自体は [予算の作成 - AWS 請求情報とコスト管理](https://docs.aws.amazon.com/ja_jp/awsaccountbilling/latest/aboutv2/budgets-create.html)を参考にすると簡単に作成できる
    - Eメールでの通知の他、[Amazon SNS](https://aws.amazon.com/jp/sns/?whats-new-cards.sort-by=item.additionalFields.postDateTime&whats-new-cards.sort-order=desc)で `Topic` と `Subscription` を作成し、それに紐づけることができる
        - `Lambda` 等を噛ませれば `Slack` 等にも通知をさせることができるはず
