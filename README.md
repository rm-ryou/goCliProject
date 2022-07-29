# 概要
指定されたディレクトリ内にあるjsonファイルをid順にソート(挿入ソート、マージソート、クイックソート)、計測して指定のファイルに出力するプログラムです。

### 環境構築
簡易的なDockerfileを用いてイメージ、コンテナを作成して、コンテナ(ubuntu:18.04)を起動させます。
```bash
make build
make run
```
### 使い方
実行ファイルの作成
```bash
make
./json_sorter -type quick -json ./misc/data/user.json -output res.json
```
引数のオプション
```bash
-type   Sort type
-json   Path of input file
-output Path of output file
```
makeを用いての実行
```bash
make type (quick, merge or insertion)
```
その他make機能の表示
```bash
make help
```
