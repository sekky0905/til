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
- 初めてserverにアクセスした際にclient側の `known_hosts` ファイルにserver側のpublic keyが記録される
-  Host keyは以下のタイミングで生成される
    - OpenSSHが初めてインストールされた時
    - computerが初めてboostした時
    - ssh-keygenした時

[known_hostsからエントリを消す - 理系学生日記](https://kiririmode.hatenablog.jp/entry/20171020/1508485674)
[【SSHをはじめて触る人へ】Linuxのホスト間認証とは？](https://eng-entrance.com/linux-ssh-host)
[OpenSSHの警告メッセージを黙らせる | Siguniang's Blog](https://siguniang.wordpress.com/2014/02/28/get-rid-of-openssh-warning-message/)
