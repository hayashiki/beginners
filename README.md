# beginners

# 目的 

Goで簡易アプリケーションをステップバイステップで作成する

# アプリ仕様

マッチングECサイトを想定した売買アプリケーションを作成する

- Merchantが登録、ログインができる
- Merchantが登録できる項目は以下
    - 性
    - 名
    - Email
    - 電話番号
- ログインはEmail/Passwordによって行う
- Merchantは商品登録ができる
    - 登録可能な項目は以下とする
        - 商品名(100文字以内)
        - 商品写真(1枚のみ)
        - 商品説明
- 買い手をUserとし、登録、ログインができる
- Userの項目はMerchantと同じでよい（後々それぞれ別属性を追加する予定）
    - 性
    - 名
    - 名前
    - Email

最初のステップはここまでとし、 取引ステップはまた追記する

# step0

スプレッドシートにて
てきとうなサンプルデータに各テーブル(Merchant, Product, User)に6件ほど登録する

# step1-1

main.goを実行する
HelloWorldを出力する

# step1-2

Webサーバをport 8000でたて、
http://localhost:8000で開いてHelloWorldをブラウザ画面から出力する

# step1-3

出力形式をJSONにして、postmanからリクエストを行い、
JSON形式(key-value)でレスポンスを返す

{
"hello": "world"
}

# step2-1

Merchant, Product, UserのStructを作成する
各Structを初期化させ、スレッドシートにいれたデータでフィールドを満たす

# step2-2

postmanからPOSTリクエスト経由で、MerchantのStructを動的に作成する
