# Basic Pattern

[AWSを学ぶために最初に構築するアーキテクチャパターン５選 - log4ketancho](https://www.ketancho.net/entry/2018/05/07/080000)の王道Webパターンのような感じ。

## 作成の簡単な流れと構造

- VPCを作成する
    - IPアドレスの範囲を決定する
- VPC内にSubnetを構築する(VPCをsubnetに分ける)
    - VPC領域に割り当てたCIDRブロックをサイラに分割してより小さなCIDRブロックを作成する
        - この分割されたCIDRブロックがSubnet
            - Public Subnet
                - インターネットからのアクセスを許容する
                - Webサーバ等
            - Private Subnet
                - インターネットからのアクセスを許容しない
                - DBサーバ等
- Internet Gatewayを設定する
    - Public Subnetをインターネットに接続するため
    - Default Gatewayを設定する
        - 設定されている送信先IPアドレス以外のパケットの送り先
            - 参考: [デフォルトゲートウェイ (default gateway)とは｜「分かりそう」で「分からない」でも「分かった」気になれるIT用語辞典](https://wa3.i-3-i.info/word12053.html)
        - 0.0.0.0/0をInternet Gatewayに設定する
- Security groupを構築する
- Webサーバを構築する(EC2)
    - 設定項目
        - リージョン
        - AMI
        - VPC
        - subnet
            - Public Subnetに構築する
        - Storage
        - Security group
- RDSを構築する
    - Private Subnetに構築する
        

## 参考

### 参考文献

- 玉川憲  (著), 片山暁雄  (著), 今井雄太 (著) (2017/4/13)『Amazon Web Services 基礎からのネットワーク&サーバー構築 改訂版』日経BP

### 参考記事

- [AWSを学ぶために最初に構築するアーキテクチャパターン５選 - log4ketancho](https://www.ketancho.net/entry/2018/05/07/080000)
- [デフォルトゲートウェイ (default gateway)とは｜「分かりそう」で「分からない」でも「分かった」気になれるIT用語辞典](https://wa3.i-3-i.info/word12053.html)
