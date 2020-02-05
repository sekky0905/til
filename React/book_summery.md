# 『React.js&Next.js超入門』 読書メモ

## Chapter 2 『JSX をマスターしよう!』

- DOM
    - HTML のタグを JavaScript から扱えるようにするためのもの
    - HTML のタグを JavaScript の object として表す
        - 属性情報等もその中に入っている
    - Element
        - HTML のタグ( `<p>` 等)を扱う object
    - Node 
        - タグを構成する要素
        - Element も Node に含まれる
- JSX
    - Render で表示できるのは、1つの Element だけ
    - 第一級オブジェクトとして使用できる


## Chapter3

## Chapter4 Redux で値を管理しよう

- 状態管理用のライブラリ
    - vuex みたいなもの
- Redux の仕組み
    - Store 
        - データの保管庫
        - ここで保管されるデータは State と呼ばれる
    - Provider
        - Store を Component に受け渡すためのもの
    - Reducer 
        - Store 内の State を変更するためのもの
            - vuex で言うところの Mutation みたいなもの
        - `action.type` 毎に条件分岐を行い、それに結びついた `ReduceAction` を呼び出す
        - ReduceAction 
            - `Reducer` から呼び出される具体的な処理を行う関数
            - 必ず新しい `State` を返す

- Redux の使用
    - connect 
        - Provider の中に Component を組み込んで Component で Store を使用する
        - connect で Component に Store を接続する
            - `connect(設定)(Component)` といった形
            - ここでの設定は、state から Component で使用するデータを絞って返すこ関数を渡す
                - connect で渡される state は Component の this.props に組み込まれる
    - doAction
        - doAction メソッドに this をバインド
        - JSX 側で onClick とかのイベントに this.doAction を紐付け、Component の JavaScript 側で doAction のメソッドを定義する
            - 上記のようにすることで、イベント発火時に doAction で定義したメソッドが実行されるようになる
    - Action
        -  state を変更する際に行うことに関する情報が詰まった Object
        - type と言うプロパティを持ち、これによって Reducer の処理が分岐する
    - Action Creator
        - `dispatch` の際に引数として渡す `Action` を作成する関数
    - Dispatch 
        - doAction 内で `this.props.dispatch(action)` といった形で呼び出される
        - dispatch が呼び出されて、Action が送信されると、 Store の Reducer が呼び出されて、処理が分岐し、必要な処理が実行される


## 参考文献

掌田 津耶乃 (2019/3/8)『React.js&Next.js超入門』 秀和システム
