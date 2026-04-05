# Changelog

## [v1.0.0] - 2026-04-06

### Added

- GoF 23デザインパターンのGo実装とJava原典を追加
  - Goイディオムで実装 (12パターン): TemplateMethod, FactoryMethod, AbstractFactory, Bridge, Prototype, Memento, Strategy, Command, Visitor, Observer, Iterator, Singleton
  - パターン構造を忠実に再現 (11パターン): Composite, Decorator, ChainOfResponsibility, Builder, Adapter, Facade, Proxy, Mediator, State, Flyweight, Interpreter
- プロジェクトドキュメント (CLAUDE.md, README.md)
- 課題管理ファイル (Docs/issues.md)
- mise.toml によるツールチェイン設定

### Fixed

- FactoryMethod: 出力メッセージの不要なスペースを除去しJava版と一致させた
- Memento: メッセージとフルーツ名を英語から日本語に修正しJava版と一致させた
- Proxy: メッセージを英語から日本語に修正しJava版と一致させた

### Notes

- Go 1.23以上が必要 (Iterator パターンが `iter.Seq` を使用)
- Java版ソースは『Java言語で学ぶデザインパターン入門 第3版』(結城浩著, MIT License) に基づく
