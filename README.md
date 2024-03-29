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

httpをHandleするメソッドはhandler.goファイルを作成して呼び出す

# step1-2

Webサーバをport 8000でたて、
http://localhost:8000で開いてHelloWorldをブラウザ画面から出力する

# step1-3

出力形式をJSONにして、postmanからリクエストを行い、
JSON形式(key-value)でレスポンスを返す

encoding/json パッケージを使ってJSONデータをエンコードする

```
{
    "hello": "world"
}
```

# step2-1

Merchant, Product, UserのStructを作成する
各Structを初期化させ、スレッドシートにいれたデータでフィールドを満たす

https://qiita.com/cotrpepe/items/b8e7f70f27813a846431

# step2-2

postmanからPOSTリクエスト経由で、MerchantのStructを動的に作成する

- requestBody
```
{
    "name": "hoge",
    "email": "hoge@example.com"
}
```

- response
```
{
    "merchant": {
        "ID": 1,
        "Email": "hoge@example.com",
        "Name": "hoge",
        "PhotoURL": ""
    },
    "success": true
}
```

# step3-1

sqlite3を利用してデータの永続化を行う

sqliteはmysqlなどのDBと違ってインストールする必要がなく、
ローカルにファイル形式で保存されるデータベースなので使い勝手がよい

Goでsqliteのデータベースにつないでデータベース操作するには
github.com/mattn/go-sqlite3のクライアントライブラリを利用する

# step3-2

DBスキーマを作成する
step0で作成した定義に基づいて作成する

ノリとしては以下
```go
var schema = `
CREATE TABLE IF NOT EXISTS merchants
(
	id INTEGER PRIMARY KEY,
	name TEXT,
	email TEXT,
	photo_url TEXT,
	timestamp DATETIME
)`

func prepareSchema(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("could not create http_requests table: %w", err)
	}
	return nil
}
```

では、どこで実行するか？
通常はコマンドJobでスキーマ作成を行うが、その解説をしていると本筋からそれるので
今回はWebサーバのエンドポイント経由で作成してみることにする

`/dbinit` というエンドポイントからスキーマを作成してみよう

# step3-3

作成したスキーマ（テーブル）にデータを挿入する

```go
    merchant := Merchant{
        Email:    "hayashiki@gmail.com",
        Name:     "aioue",
        PhotoURL: "https://hoge.com",
    }
    
    ctx := context.Background()
    result := db.MustExecContext(ctx,
        "INSERT INTO merchants(email, name, photo_url)\nVALUES (?,?,?)",
        merchant.Email, merchant.Name, merchant.PhotoURL)
```

# step3-4

作成したデータを抽出する（複数）
エンドポイントは"/merchants/list"とする
RestfulのURL設計としては適切ではないのだが、 一旦はこれで。

# step4-1

[chi](https://github.com/go-chi/chi)を導入する
net/httpのハンドラー周りを薄くラップしたライブラリとなる


# step4-2

- 指定したIDのMerchantのデータを取得する
- 指定したIDのMerchantのデータを編集する
  例えば名前を変更する（そのようなユースケースはあまり実態に即してないけど）

# step4-3

- Merchantの新規登録のhtml form画面を作成する
- 
