# SSHのまとめ

## 認証

- ユーザーがSSHでclientからserverにログインする際には以下の2種類の認証が行われる
    - ホスト認証
        - ログイン先のserverが第三者のなりすましではなく、意図したログイン先かどうかを確認
    - ユーザー認証
        - ユーザーがそのserverにログインしていいかどうかを確認
      
参考: [入門OpenSSH / 第3章 OpenSSH のしくみ](http://www.unixuser.org/~euske/doc/openssh/book/chap3.html)

### ユーザー認証

#### Password認証

#### 公開鍵認証

### ホスト認証

- clientからserverへsshで接続を行う時に、server側がなりすましている偽物ではないかを確認するために用いられる
- 公開鍵暗号方式を利用する

#### 仕組み

- server側にserverのprivate keyを置いておいて、clientがserverにアクセスした際にはserverのpublic keyがclientに送られる
- 初めてserverにアクセスした際にclient側の `known_hosts` ファイルにserver側のpublic keyが記録される
- 初回以降はその `known_hosts` ファイルに記載されたpublic keyとserver側からアクセスの度に送信されてくるpublic keyを照会して、一致すれば同様のserverとみなす

[known_hostsからエントリを消す - 理系学生日記](https://kiririmode.hatenablog.jp/entry/20171020/1508485674)
[【SSHをはじめて触る人へ】Linuxのホスト間認証とは？](https://eng-entrance.com/linux-ssh-host)
[OpenSSHの警告メッセージを黙らせる | Siguniang's Blog](https://siguniang.wordpress.com/2014/02/28/get-rid-of-openssh-warning-message/)

#### ホスト認証が存在する理由

- 一言で言うと、clientがSSHの通信でやりとりするserverが意図したserverかどうかを確認するため
- clientは、serverAとSSH接続したいとする
- なりすましする人がserverBを用意したとする
    - client-serverAのやりとりには、client: public key、serverA: private keyが必要
    - ここで、仮にserverBがserverAになりすまし、serverAのpublic keyをclientに送信して、clientがserverAのpublic keyで暗号化して、serverBに送信したとしてもserverBはその内容を読むことができない
        - clientはserverAのpublic keyで暗号化しているため、serverAの持っているprivate keyでしか復号できず、serverBはserverAのprivate keyを持っていないので
- ここからが、Host Keyが存在する理由
    - serverBがserverAのIPを名乗って、clientがserverBをserverAだと思ってSSH接続したとする
        - これが成功した場合、clientは、serverA だと思って serverBの public keyでserverBとやりとりをしてしまうので、serverAに送るつもりだったものをserverBが取得し、severBのprivate keyで複合することができてしまう
    - しかし、client側がserverAの正しいpublic keyを記録しておき(known_hosts)、serverAにアクセスする度にpublic keyがserverAのものかを確認することで、clientはserverAだと思ってserverBにアクセスすることはなくなる
 
### Host keyの設定

- server側の `sshd_config` ファイルの  `HostKey`  の項目にhost keyで使用するprivate keyの場所を記述することで指定する(public keyでの指定も可能)

[sshd_config(5) - OpenBSD manual pages](https://man.openbsd.org/sshd_config#HostKey)

## フォワーディング

### ポートフォワーディング

- 他のTCP上のプロトコルのセッションをSSHセッション上で安全に転送させる機能
- 他のTCP上のプロトコルが安全に通信をできるSSHが提供する通り道をトンネルと言うこともあり、トンネリングと呼ばれることもある

#### 例

##### SSHポートフォワーディングを使用しない例

- SMTPプロトコルを通じてクライアントから、メールサーバへ電子メールを送信するとする
- 通常の場合、SMTPプロトコルを使用してSMTPサーバにアクセスする際には25番ポートを使用するので、TCP上の25番ポートを通じてデータがやりとりされる
- そのまま素直にSMTPを使用してクライアントとSMTPサーバ間でやりとりすると、その間を流れるデータは暗号化されずに平文で流れる

##### SSHポートフォワーディングを使用した例

- クライント側とサーバ側をSSHセッションで結ぶトンネルを作成する
    - その際にクライアント側でその時点で使用されていないポート番号を指定しておく
        - ここでは仮に2000とする
        - これがトンネルの入り口になる
    - クライアントがサーバに対してSMTPでデータを送信する場合、クライアントはこのポート番号2000からデータを送信する
    - そうすると、先ほど作成されたクライント側とサーバ側をSSHセッションで結ぶトンネルを通じてサーバ側にデータが送信される
        - SSHセッションで結ぶトンネルを通じてデータが送信されるので、データは暗号化される
        - クラインアントの2000番ポート→SSHセッション上(トンネル)→サーバ上の25番ポートと言う流れで送信される

#### ローカルフォワード

- sshのクライアント側に存在するアプリケーションのクライアント→sshサーバ側に存在するアプリケーションのサーバへ送信する

#### リモートフォワード

- sshサーバ側に存在するアプリケーションのクライント→sshのクライアント側に存在するアプリケーションのサーバへ送信する

### エージェントフォワーディング

## 多段SSH


## 参考文献

- Daniel J. Barrett, Richard E. Silverman, Robert G. Byrnes　著、小島 肇　監訳、坂井 順行、鹿田 幸治、園田 道夫、高橋 基信、根津 研介、宮本 久仁男、両輪 顕、関谷 麻美　訳(2006/11/22)『実用SSH 第2版―セキュアシェル徹底活用ガイド』オライリー・ジャパン

