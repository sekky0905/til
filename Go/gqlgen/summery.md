# gqlgen まとめ

- Query の引数で Required ではないものに関しては、生成される引数は promitive 型であってもポインタで生成される
	- 例えば `comments(userID: ID, chatRoomID: ID!): [Comment]` のような Query の場合には、userID は string のポインタで生成され、chatRoomID はstring 型で生成される
