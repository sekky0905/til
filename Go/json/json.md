# JSON 系 Tips


- `json:",omitempty"` のように `omitempty"` を付与すると、付与されたフィールドが空だった場合に Marshal の際に出力されないようにすることができる

参考: [golangのenconding/jsonのタグについて - Qiita](https://qiita.com/pside/items/0fc7f85746211371b23e#marshal%E6%99%82%E3%81%AB%E5%87%BA%E5%8A%9B%E3%81%97%E3%81%AA%E3%81%84)
