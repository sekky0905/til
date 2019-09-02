# MySQL の timezone のメモ

## @@system_time_zone, @@global.time_zone, @@time_zone

- 大まかに分けると、 `system_time_zone` と `time_zone` がある
    - 更に `time_zone` には、`global` と `session` があり、`session` の場合には、そのセッションに限った設定になる
- 以下のような感じで、`@@global.time_zone` とか `@@time_zone` に設定されている `SYSTEM` は、`system_time_zone` の値
    - `system_time_zone` が `UTC` で、 `@@global.time_zone` も `@@time_zone` `SYSTEM` なので、全て `UTC` になる。
        - グロバーバルでもセッションごとでも `UTC`
        
```sql
SELECT @@system_time_zone, @@global.time_zone, @@time_zone;
+--------------------+--------------------+-------------+
| @@system_time_zone | @@global.time_zone | @@time_zone |
+--------------------+--------------------+-------------+
| UTC                | SYSTEM             | SYSTEM      |
+--------------------+--------------------+-------------+
```

## 参考

- [time_zone設定の違うMySQLのレプリケーションについて - 角待ちは対空](https://blog.yux3.net/entry/2018/12/02/182113)
- [MySQLのタイムゾーン - @tmtms のメモ](https://tmtms.hatenablog.com/entry/2015/08/22/mysql-timezone)
