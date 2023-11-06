## golang-samples

Go言語の基本構文や各種パッケージのサンプルプログラム集

### インストール
  Go言語のバージョン管理ツールである goenv を導入し Go をインストールする  
[goenvによるインストール](./doc/goenv.md)

### 実行手順

golang-samples 直下で以下のコマンドを実行

```
go run main.go [テストタイプ] [テスト番号]
```
**実行例**
```
go run main.go array 001
```

**テストタイプ一覧**
- [基本構文カテゴリ](./samples/basic)
  - array       ・・・  [配列](./samples/basic/array)
  - slice       ・・・  [スライス](./samples/basic/slice)
  - map         ・・・  [マップ](./samples/basic/map)
  - if          ・・・  [条件分岐(if)](./samples/basic/if)
  - switch      ・・・  [条件分岐(switch)](./samples/basic/switch)
  - for         ・・・  [ループ(for)](./samples/basic/for)
  - func        ・・・  [関数(func)](./samples/basic/func)
  -  struct     ・・・  [構造体(struct)](./samples/basic/struct)
  -  interface  ・・・  [インターフェース(interface)](./samples/basic/interface)
  -  goroutine  ・・・  [ゴルーチン(goroutine)](./samples/basic/goroutine)

- パッケージカテゴリ
  -  csv       ・・・  [encoding/csv](./samples/csv)
  -  errors    ・・・  [errors](./samples/errors)
  -  filepath  ・・・  [filepath](./samples/filepath)
  -  json      ・・・  [encoding/json](./samples/json)
  -  log       ・・・  [log](./samples/log)
  -  os        ・・・  [os](./samples/os)
  -  slices    ・・・  [slices](./samples/slices)
  -  strconv   ・・・  [strconv](./samples/strconv)
  -  strings   ・・・  [strings](./samples/strings)
  -  time      ・・・  [time](./samples/time)
