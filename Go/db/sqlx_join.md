# github.com/jmoiron/sqlx で複数のテーブルを JOIN して構造体の埋め込みにロードする際にフィールド名が被る場合には、どうしたらいいか?

## 結論

- 結論としては、SQL 側で `AS` で別名を付与するのが良さそう。
- 同一のフィールド名が存在する場合には、`StructScan` は最初に発見したフィールドを選択する模様
    - ドキュメントでは、以下のように記載されているが、ここでいうフィールドとは構造体のフィールドの `db` タグのことよう(こう考えた根拠は実験で記載)
> StructScan will choose the "first" field encountered which has that name

引用元: 参考: [Illustrated Guide to SQLX](http://jmoiron.github.io/sqlx/#advancedScanning)


## StructScan を考えるための前提知識

### Select

> Select executes a query using the provided Queryer, and StructScans each row into dest, which must be a slice. 
If the slice elements are scannable, then the result set must have only one column. Otherwise, StructScan is used. 
The *sql.Rows are closed automatically. Any placeholder parameters are replaced with supplied args.

【意訳】
`Select` は、与えられた [Queryer](https://godoc.org/github.com/jmoiron/sqlx#Queryer) を利用してクエリを実行する。また、各々の行を `dest` に `StructScan` する。`dest` は、`slice` である必要がある。
もし、 `slice` の要素がスキャンできるものであるならば、結果セットは1つのカラムのみを所持している必要がある。そうでなければ、 `StructScan` が使用される。
`*sql.Rows` は、自動で `close` される。いかなるプレースホルダーのパラメータも、与えられた引数に置き換えられる。

引用元: [sqlx - Select - GoDoc](https://godoc.org/github.com/jmoiron/sqlx#Select)

### StructScan

> StructScan all rows from an sql.Rows or an sqlx.Rows into the dest slice. 
StructScan will scan in the entire rows result, so if you do not want to allocate structs for the entire result, use Queryx and see sqlx.Rows.StructScan. 
If rows is sqlx.Rows, it will use its mapper, otherwise it will use the default.

【意訳】
`sql.Rows` か `sqlx.Rows` から全ての行を `dest` 用の `slice` に `StrctScan` する。
`StructScan` は、全ての行の結果で実行するので、もし構造体を全ての結果に割り当てたくない場合には、`Queryx` を利用して、`sqlx.Rows.StructScan` を見るようにする。
もし、行が `sqlx.Rows` なのであれば、それ自体のマッパーを使用し、そうでないのであれば、デフォルトのマッパーを使用する。

引用元: [sqlx - StructScan - GoDoc](https://godoc.org/github.com/jmoiron/sqlx#StructScan)


## 実験

以下のようなテーブルを定義する。

```sql
CREATE TABLE IF NOT EXISTS `a` (
  `id` INT UNSIGNED NOT NULL,
  `hoge` VARCHAR(10) NOT NULL,
  `foo` VARCHAR(10) NOT NULL,
  PRIMARY KEY(`id`)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `b` (
  `a_id` INT UNSIGNED NOT NULL,
  `hoge` VARCHAR(10) NOT NULL,
  `foo` VARCHAR(10) NOT NULL,
  PRIMARY KEY(`a_id`)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

以下のデータを格納する。

```sql
INSERT INTO a (id, hoge, foo) VALUES (1, 'hogeA', 'fooA');
INSERT INTO b (a_id, hoge, foo) VALUES (1, 'hogeB', 'fooB');
```

下記の `構造体のフィールド名を変更してみる` と `AS を利用してみる` の実行結果から、[ドキュメント](http://jmoiron.github.io/sqlx/#advancedScanning)の以下の文章が指す `"first" field` とは、 `db` タグを付与されたフィールドということがわかる。

> StructScan will choose the "first" field encountered which has that name

引用元: 参考: [Illustrated Guide to SQLX](http://jmoiron.github.io/sqlx/#advancedScanning)

### 何も考えずに構造体を埋め込んでみる

#### コード

```sql
type A struct {
	ID   uint32 `db:"id"`
	Hoge string `db:"hoge"`
	Foo  string `db:"foo"`
}

type B struct {
	AID  uint32 `db:"a_id"`
	Hoge string `db:"hoge"`
	Foo  string `db:"foo"`
}

type C struct {
	*A
	*B
}

func Sample1() ([]*C, error) {
	sql := "SELECT a.id, a.hoge, a.foo, b.a_id, b.hoge, b.foo FROM a LEFT OUTER JOIN b ON a.id = b.a_id"

	var c []*C

	if err := db.Select(&c, sql); err != nil {
		// error handling
	}

	log.Printf("Sample1 => len:%d", len(c))

	for k, v := range c {
		log.Printf("Sample1 => C.A:%+v, C.B:%+v\n", v.A, v.B)
	}

	return c, nil
}
```

#### 実行結果

```
Sample1 => len:1
Sample1 => C.A:&{ID:1 Hoge: Foo:}, C.B:&{AID:1 Hoge:hogeB Foo:fooB}
```

- フィールド名及び `db` タグが被っている `Hoge` と `Foo` は適切に取得されない


### 構造体のフィールド名を変更してみる

#### コード

```go
type A2 struct {
	ID   uint32 `db:"id"`
	Hoge string `db:"hoge"`
	Foo  string `db:"foo"`
}

type B2 struct {
	AID   uint32 `db:"a_id"`
	Hoge2 string `db:"hoge"`
	Foo2  string `db:"foo"`
}

type C2 struct {
	*A2
	*B2
}

func Sample2() ([]*C2, error) {
	sql := "SELECT a.id, a.hoge, a.foo, b.a_id, b.hoge, b.foo FROM a LEFT OUTER JOIN b ON a.id = b.a_id"

	var c []*C2

	if err := db.Select(&c, sql); err != nil {
		// error handling
	}

	log.Printf("Sample2 => len:%d", len(c))

	for k, v := range c {
		log.Printf("Sample2 => C2.A2:%+v, C2.B2:%+v\n", v.A2, v.B2)
	}

	return c, nil
}
```

#### 実行結果

```
Sample2 => len:1
Sample2 => C2.A2:&{ID:1 Hoge: Foo:}, C2.B2:&{AID:1 Hoge2:hogeB Foo2:fooB}
```

- 構造体のフィールド名は異なるものの、`db` タグが被っている `Hoge` と `Foo` は適切に取得されない

### AS を利用してみる

#### コード

```go
type A3 struct {
	ID   uint32 `db:"id"`
	Hoge string `db:"hoge"`
	Foo  string `db:"foo"`
}

type B3 struct {
	AID   uint32 `db:"a_id"`
	Hoge string `db:"b_hoge"`
	Foo  string `db:"b_foo"`
}

type C3 struct {
	*A3
	*B3
}

func Sample3() ([]*C3, error) {
	sql := "SELECT a.id, a.hoge, a.foo, b.a_id, b.hoge AS b_hoge, b.foo AS b_foo FROM a LEFT OUTER JOIN b ON a.id = b.a_id"

	var c []*C3

	if err := db.Select(&c, sql); err != nil {
		// error handling
	}

	log.Printf("Sample3 => len:%d", len(c))

	for k, v := range c {
		log.Printf("%d: Sample3 => C3.A3:%+v, C3.B3:%+v\n", k, v.A3, v.B3)
	}

	return c, nil
}
```

#### 実行結果

```
Sample3 => len:1
Sample3 => C3.A3:&{ID:1 Hoge:hogeA Foo:fooA}, C3.B3:&{AID:1 Hoge:hogeB Foo:fooB}
```

- 構造体のフィールド名は同じでも、`db` タグが異なれば、適切にデータは格納される


