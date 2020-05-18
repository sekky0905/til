# 普段よく使うGolandのコマンド一覧
社内でjetBrain社製のIDEの勉強会があったので、普段[GoLand](https://www.jetbrains.com/go/)でよく使っているコマンドをまとめた。

## コマンドライン立ち上げ

`alt + F12`

## 定義元へジャンプ

`command + トラックパッド`

## 変数使用先へジャンプ

`command + option + F7`

## 前へに操作していたところにジャンプ

`command + option + 左`

## 次へに操作していたところにジャンプ

`command + option + 右`

## 関数/メソッドにする

`command + option + M`

## window追加

`shift + F4`

## 縦の連続選択

`shift + option + トラックパッドで縦に選択`

縦全体に space を削除するとかのに使うと便利

## 再定義

`shift + F6`

## テスト自動作成

`command + shift + T`

## ifとか自動で描いてくれる

`command + option + T`

## interfaceの実装先へ

`command + option + B`

## 空の構造体を作る

以下みたいなので `{}` の部分で赤いマーク => `fill struct recursively`

```go
	tests := []struct {
		name    string
		r       *hogeRepository
		args    args
		want    *model.Hoge
		wantErr bool
	}{
		{}
	}
```

## コード保存時に自動でgo fmtやgometalinterを効かせる様に設定する
[Golandでコード保存時に自動でgo fmtやgometalinterを効かせる様に設定する際のメモ - Qiita](https://qiita.com/Sekky0905/items/c40e3375c72c003e8b1d)

# その他plugin
## ignore
.gitignoreとか自動で作ってくれる
[.ignore - Plugins | JetBrains](https://plugins.jetbrains.com/plugin/7495--ignore)

## Makefile Support
[Makefile support - Plugins | JetBrains](https://plugins.jetbrains.com/plugin/9333-makefile-support)

## Yaml
[YAML/Ansible support - Plugins | JetBrains](https://plugins.jetbrains.com/plugin/7792-yaml-ansible-support)

## Toml
[Toml - Plugins | JetBrains](https://plugins.jetbrains.com/plugin/8195-toml)

## Plant UML
[PlantUML integration - Plugins | JetBrains](https://plugins.jetbrains.com/plugin/7017-plantuml-integration)

## Rest Client 

- [Intellijのeditor-based Rest Clientをおためし - Qiita](https://qiita.com/U_UU/items/2efbe23156e5ec860bed)

## 大文字/小文字

`Command + Shift + u`

# 参考にさせていただいた記事
[GoLand: A Clever IDE to Go by JetBrains](https://www.jetbrains.com/go/?fromMenu)

[メルカリ社員100人に聞いたGoLandの使い方 / JetBrains Night Tokyo 2018 - Speaker Deck](https://speakerdeck.com/vvakame/jetbrains-night-tokyo-2018)

[忙しい人のためのIntelliJ IDEAショートカット集（´-`） - Qiita](https://qiita.com/yoppe/items/f7cbeb825c071691d3f2)




