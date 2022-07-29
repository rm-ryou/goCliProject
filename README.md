# goCliProject
指定されたディレクトリ内にあるjsonファイルをソートして指定のファイルに出力するプログラムです。

環境構築(ubuntu:18.04)
```bash
make build
make run
```
### 使い方
実行ファイルの作成
```bash
make
```
ソートの種類の選択
```bash
make quick
make merge
make insertion
```
または、
```
./json_sorter -type (sort type) -json (Path of input file) -output (Path to output file)
```
