# UI

## 開発の進め方

以下の5つのステップ
1. UIをコンポーネント階層に分割する
2. Reactで静的バージョンを構築する
3. UIステートの最小かつ完全な表現を見つける
4. 状態をどこに置くべきかを特定する
5. 逆データフローを追加する

以下から抜粋し日本語化
https://react.dev/learn/thinking-in-react

## ReactのComponentの粒度について

- Webがインタラクティブになっていくにつれて、ロジックがコンテンツを決定することが多くなった。
- つまり、JavaScriptがHTMLに対して責任を持つようになった
- そのためReactでは、レンダリングロジックとマークアップが同じ場所（コンポーネント）に同居させている

> But as the Web became more interactive, logic increasingly determined content. JavaScript was in charge of the HTML! This is why in React, rendering logic and markup live together in the same place—components.

https://react.dev/learn/writing-markup-with-jsx

## ReactのStateはimmutableであるべき

### そもそものJavaScriptのデータ型のimmutable/mutable

- string, number, boolean等のプリミティブ型はそもそもimmutableである
- object, array等の複合型はmutableである

### ReactのStateはimmutableであるべき

- state setting functionを使わなければ、Reactはオブジェクトが変更されたことを検知できない 
- レンダリングでアクセスできるステートの値は、読み取り専用として扱うべき
- 再レンダリングをトリガーするには、新しいオブジェクトを作成してステートstate setting functionに渡す

refs: https://react.dev/learn/updating-objects-in-state