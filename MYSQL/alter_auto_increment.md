# MySQLで AUTO INCREMENT をn番からの採番にする

`AUTO INCREMENT` をカラムに指定すると、1から自動で採番にしてくれるようになるが、1からではなく例えば500から採番してほしいような場合の操作方法

## 操作

### やり方

```sql
ALTER TABLE テーブル名 AUTO_INCREMENT = 採番開始したい数字;
```

### 実際にやってみる

`hoges` table を作成

```sql
CREATE TABLE IF NOT EXISTS `hoges` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  PRIMARY KEY(`id`)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

 `AUTO INCREMENT` の変更前にデータを1つ入力

```sql
INSERT INTO hoges (name) VALUES ('foo');
```

入力されたデータを確認

```sql
SELECT * FROM hoges;
```

結果

```sql
+----+------+
| id | name |
+----+------+
|  1 | foo  |
+----+------+
```

`hoges` table の `AUTO_INCREMENT` を500に変更 

```sql
ALTER TABLE hoges AUTO_INCREMENT = 500;
```

 `AUTO INCREMENT` の変更後にデータを1つ入力

```sql
INSERT INTO hoges (name) values ('bar');
```

確認

```sql
SELECT * FROM hoges;
```

結果( `id = 500` のデータが作成されている)

```sql
+-----+------+
| id  | name |
+-----+------+
|   1 | foo  |
| 500 | bar  |
+-----+------+
```

