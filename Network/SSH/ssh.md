# SSH

## 認証

- ユーザーがSSHでclientからserverにログインする際には以下の2種類の認証が行われる
    - ホスト認証
        - ログイン先のserverが第三者のなりすましではなく、意図したログイン先かどうかを確認
    - ユーザー認証
        - ユーザーがそのserverにログインしていいかどうかを確認

参考: [入門OpenSSH / 第3章 OpenSSH のしくみ](http://www.unixuser.org/~euske/doc/openssh/book/chap3.html)

## finger print

- clientからserverへsshで接続を行う時に、server側がなりすましている偽物ではないかを確認するために用いる
- 公開鍵暗号方式を用いる
- server側にprivate keyを置いておいて、clientがserverにアクセスした際にはpublic keyが送られる
- 初めてserverにアクセスした際にclient側の `known_hosts` ファイルにserver側のpublic keyが記録される
- 初回以降はその `known_hosts` ファイルに記載されたpublic keyとserver側からアクセスの度に送信されてくるpublic keyを照会して、一致すれば同様のserver
-  Host keyは以下のタイミングで生成される
    - OpenSSHが初めてインストールされた時
    - computerが初めてboostした時
    - ssh-keygenした時

[known_hostsからエントリを消す - 理系学生日記](https://kiririmode.hatenablog.jp/entry/20171020/1508485674)
[【SSHをはじめて触る人へ】Linuxのホスト間認証とは？](https://eng-entrance.com/linux-ssh-host)
[OpenSSHの警告メッセージを黙らせる | Siguniang's Blog](https://siguniang.wordpress.com/2014/02/28/get-rid-of-openssh-warning-message/)

### Host Keyが存在する理由

- clientは、serverAとSSH接続したいとする
- なりすましする人がserverBを用意したとする
    - client-serverAのやりとりには
        - client: public key、serverA: private key
        - ここで、仮にserverBがserverAになりすまし、serverAのpublic keyをclientに送信して、clientがserverAのpublic keyで暗号化して、serverBに送信したとしてもserverBはその内容を読むことができない
            - clientはserverAのpublic keyで暗号化しているため、serverAの持っているprivate keyでしか復号できず、serverBはserverAのprivate keyを持っていないので
- ここからが、Host Keyが存在する理由
    - serverBがserverAのIPを名乗って、clientがserverBをserverAだと思ってSSH接続したとする
        - これが成功した場合、clientは、serverA だと思って serverBの public keyでserverBとやりとりをしてしまうので、serverAに送るつもりだったものをserverBが取得し、severBのprivate keyで複合することができてしまう
    - しかし、client側がserverAの正しいpublic keyを記録しておき(known_hosts)、serverAにアクセスする度にpublic keyがserverAのものかを確認することで、clientはserverAだと思ってserverBにアクセスすることはなくなる
 
