# Flutter のメモ

## Flutter のデザイン

- [Material Design](https://material.io/design/)
    - iOS のデザイン用の Widgets もないわけではない
    - `Cupertino*` 系の Widgets がそれにあたる

## Widgets 

- UI の部品の単位
- この Widgets を積み重ねて UI を作成していく
- Web で言う所の Component みたいなもの
- 参考: [Introduction to widgets - Flutter](https://flutter.dev/docs/development/ui/widgets-intro)

### Widgets のレンダリング

- `build` メソッドを利用する
    - この中で `Widgets` を定義していく
    - `build` メソッドの第一引数は、 `BuildContext`
        - `Widget Tree` での `Widget` の位置情報を伝える
        
### Widgets 状態

- `Widgets` は `immutable`

#### StatelessWidgets 

- ユーザーとおインタラクティブなやりとりを担当しない `Widgets` のこと 
- 独自の `StatelessWidgets` を定義する際には、[StatelessWidget class](https://api.flutter.dev/flutter/widgets/StatelessWidget-class.html)を `extends` する必要がある

## Scaffold Class

> Implements the basic material design visual layout structure.
  
>  This class provides APIs for showing drawers, snack bars, and bottom sheets.


[意訳] 
マテリアルデザインの視覚レイアウトの基本的な構造を実装している
drawers、snack bars、bottom sheets 向けのAPIを提供している

引用元: [Scaffold class - material library - Dart API](https://api.flutter.dev/flutter/material/Scaffold-class.html)
