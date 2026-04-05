# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**jrnl Project Tag:** gof-go-samples

## プロジェクト概要

GoFの23デザインパターンについて、『Java言語で学ぶデザインパターン入門 第3版』(結城浩著)のJavaサンプルコードと同じ題材をGo言語で実装するサンプル集。付属ブログ記事シリーズの付属コード。

## ビルドと実行

各パターンは独立した `go.mod` を持つ。Go 1.23以上が必要（Iterator が `iter.Seq` を使用）。

```bash
# 個別パターンの実行
cd src/<PatternName>/Go && go run main.go

# Java版の実行（mise で java 環境を構築済みの場合）
cd src/<PatternName>/Java && javac *.java && java Main

# 全パターン一括実行
for dir in src/*/Go; do echo "=== $(basename $(dirname $dir)) ===" && (cd "$dir" && go run main.go) && echo; done
```

一部パターンはコマンドライン引数が必要:
- `AbstractFactory`: `go run main.go <filename.html> <list|div>`
- `Builder`: `go run main.go <text|html>`
- `Flyweight`: `go run main.go <digits>`
- `Facade`: カレントディレクトリに `maildata.txt` が必要

## 実装方針（必ず遵守）

### 基本方針

1. **Javaソースと同等の機能をGo言語で実装する** — 題材と動作はJava版を踏襲する
2. **Go言語が言語レベルで解決しているパターンや非推奨パターンは、Goイディオムや現代の実装手法を採用する** — ただし実装する機能（型の役割、入出力、動作）はJava版と同等にする
3. **上記に該当しないパターンは、パターン構造をGoで忠実に再現する** — インターフェースと構造体の埋め込みでJavaのクラス階層を表現する
4. **Javaソースの機能要素に欠落がないようにする**

### パターン分類

**Goイディオムで実装（12パターン）** — パターン構造は不要だが機能はJava版と同等にする:
Template Method, Factory Method, Abstract Factory, Bridge, Prototype, Memento, Strategy, Command, Visitor, Observer, Iterator, Singleton

**パターン構造を忠実に再現（11パターン）** — 構造と機能の両方をJava版に合わせる:
Composite, Decorator, Chain of Responsibility, Builder, Adapter, Facade, Proxy, Mediator, State, Flyweight, Interpreter

### 修正・新規作成時の注意

- Go版を変更する前に、必ず対応する `src/<PatternName>/Java/` のJavaソースを読み、型の対応関係と機能を確認すること
- 「Goイディオムで実装」のパターンでも、Java版に存在しない独自の型や題材を発明してはならない
- 検証時は Java版とGo版の両方を実行し、出力を比較すること

## リポジトリ構成

```
src/<PatternName>/
  Java/       # 結城浩著のJavaサンプル（MIT License）。変更不可
  Go/
    go.mod    # パターンごとに独立（module example.com/<pattern>）
    main.go   # Go実装（単一ファイル）
Docs/
  issues.md   # 課題管理
```

## 課題管理

`Docs/issues.md` で課題を管理している。修正前に必ず確認すること。
