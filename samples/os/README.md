## osパッケージ

各OSに依存する機能を内部に隠ぺいしプラットフォームに依存しないAPIを提供するパッケージ。  
ファイルやディレクトリ・環境変数の操作などの機能を含む。

### フォルダ操作
1. [カレントディレクトのパスを取得(Getwd)](./os_sample_001.go)
2. [ディレクトリの移動(Chdir)](./os_sample_002.go)
3. [ディレクトの作成(Mkdir)](./os_sample_003.go)
4. [ディレクトの一括作成(MkdirAll)](./os_sample_004.go)
5. [テンポラリディレクトリの作成(MkdirTemp)](./os_sample_005.go)

### ファイル操作
6. [ファイルを開く(Open)、ファイルを閉じる(Close)](./os_sample_006.go)
7. [書き込み用にファイルを開く(Create)](./os_sample_007.go)
8. [ファイルの読み込み(Read、ReadAt)](./os_sample_008.go)
9. [ファイルの読み込み(ReadFile)](./os_sample_009.go)
10. [ファイルの書き込み(Write、WriteAt)](./os_sample_010.go)
11. [ファイルの書き込み(WriteFile)](./os_sample_011.go)
12. [ファイル、ディレクトリの削除(Remove)](./os_sample_012.go)
13. [ファイル名の変更と移動(Rename)](./os_sample_013.go)
14. [ファイルの存在確認(stat)](./os_sample_014.go)
15. [シンボリックリンクの作成(Symlink)](./os_sample_014.go)
16. [シンボリックリンクのリンク先を読み込む(Readlink)](./os_sample_015.go)

### ホスト名
17. [ホスト名の取得(Hostname)](./os_sample_016.go)

### 環境変数操作
18. [環境変数の取得(Getenv)](./os_sample_017.go)
19. [環境変数の取得(Environ)](./os_sample_018.go)
20. [環境変数の設定(Setenv)](./os_sample_019.go)
21. [環境変数の削除(Unsetenv)](./os_sample_020.go)
22. [環境変数の存在チェック(LookupEnv)](./os_sample_021.go)
