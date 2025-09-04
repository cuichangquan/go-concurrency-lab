# go-concurrency-lab
go-concurrency-lab My experiments with Go concurrency (goroutines, channels, worker pools, benchmarks). I am learning by doing and sharing my results here.  Feel free to give feedback or open an issue!


# Golang install

🔹 1. asdf で管理（最も一般的・おすすめ）
asdf は Node, Python, Go など複数言語のバージョンを統一的に管理できます。
# asdf のインストール
```
brew install asdf
asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
asdf list all golang
asdf install golang 1.25.1
asdf global golang 1.25.1
cd ~/go-concurrency-lab
asdf set golang 1.25.1
asdf current golang
go version
```
