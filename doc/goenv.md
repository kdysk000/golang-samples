### goenvのインストール
```
git clone https://github.com/syndbg/goenv.git ~/.goenv
```

### 環境変数の設定
~/.bashrc に設定しておく
.bashrc の最後に以下を追加
```
export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"
export PATH="$GOROOT/bin:$PATH"
export PATH="$PATH:$GOPATH/bin"
```

### Goのインストール
```
goenv install <バージョン>
```

※インストールできるGoのバージョン一覧は以下コマンドで確認できる
```
goenv install -l
```

- 使用するGoのバージョン切り替え
```
goenv global <バージョン>
```
- 特定のディレクトで使用するGoのバージョンを設定する場合
```
goenv local <バージョン>
```

