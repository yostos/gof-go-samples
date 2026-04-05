# GoF Go Samples — 課題

## 概要

README.mdの実装方針に基づき、全23パターンのGo実装を検証した結果、以下の2件が方針と合致しない実装であることを確認した。

---

## 有効な課題

### Issue 1: Singleton — 題材がJava版と完全に異なる

**READMEの分類**: Goイディオムで実装するパターン（依存性の注入）

**問題**: Goイディオム（DI）の採用自体は方針通りだが、実装している型と題材がJava版とまったく異なる。Go版は `Config`, `App`, `AnotherService` というJava版に存在しない独自の型を使っており、「実装する機能はJavaソースを踏襲する」という方針に違反している。

**Java版の内容**:
- `Singleton` クラス（privateコンストラクタ、static instance、`getInstance()`）
- mainで `getInstance()` を3回呼び出し、同一インスタンスであることを確認

**Go版の現状**:
- `Config`, `App`, `AnotherService` というJava版に存在しない型
- DIによるインスタンス共有を実演しているが、題材が別物

**修正方針**:
- Java版の `Singleton` クラスに対応する型を用意する
- Goイディオムとして `sync.Once` によるシングルトンではなくDIで実装するのは方針通り
- mainではJava版と同様に同一インスタンスであることを確認する動作を再現する

---

### Issue 2: Chain of Responsibility — パターン構造が関数型に置き換えられている

**READMEの分類**: パターン構造を忠実に再現するパターン

**問題**: 「忠実に再現」すべきパターンだが、各サポートがクロージャ（`Handler` 関数型）で実装され、`chain()` 関数によりスライスに集約されている。Java版のSupport連鎖（name/next/setNext/resolve によるリンクリスト構造）が失われている。

**Java版の構造**:
- `Support` 抽象クラス (`name`, `next` フィールド, `setNext()`, `support()` テンプレート, 抽象 `resolve()`)
- `NoSupport` クラス (extends Support)
- `LimitSupport` クラス (extends Support)
- `OddSupport` クラス (extends Support)
- `SpecialSupport` クラス (extends Support)
- mainで `setNext()` によるチェーン構築、Trouble を先頭から順次処理

**Go版の現状**:
- `Handler` 関数型とクロージャ（noSupport, limitSupport, oddSupport, specialSupport）
- `chain()` 関数でスライスに集約
- 出力はJava版と同等だが、パターン構造が消えている

**修正方針**:
- `Support` インターフェースを定義（`SetNext(Support) Support`, `Handle(Trouble)`, `Resolve(Trouble) bool`, `Name() string`）
- `BaseSupport` 構造体を作成（`name` と `next Support` フィールド、`SetNext()`, `Handle()` の共通ロジック）
- `NoSupport`, `LimitSupport`, `OddSupport`, `SpecialSupport` 構造体に `BaseSupport` を埋め込み、各 `Resolve()` を実装
- mainで `SetNext()` によるチェーン構築と、Troubleの順次処理を再現
- 出力はJava版と同等にする

---

## 検証済み — 課題なし

### Goイディオムで実装するパターン（11パターン）

以下は「Goイディオムで実装するパターン」に該当し、かつ題材・機能がJava版と同等であることをソースコードレベルで確認したもの。

| パターン | Goの実装手法 | Java版との機能比較 |
|----------|-------------|-------------------|
| Template Method | インターフェース + テンプレート関数 | CharDisplay/StringDisplay、display()で5回表示。同等 |
| Factory Method | ファクトリ関数 | IDCard生成・登録・使用。出力同等 |
| Abstract Factory | PageRenderer + 関数フィールド | Link/Tray/Page、list/div形式のHTML生成。同等 |
| Bridge | DisplayImpl + 関数 | StringDisplayImpl、display/multiDisplay。出力同等 |
| Prototype | 値コピー + prototypesマップ | MessageBox/UnderlinePen、プロトタイプ登録・取得・使用。同等 |
| Memento | Gamer構造体コピー | サイコロゲーム、save/restore。同等 |
| Strategy | 関数型（StrategyFunc/StudyFunc） | じゃんけん、WinningStrategy/ProbStrategy相当、10000回対戦。同等 |
| Command | クロージャスライス | テキストCanvas、描画・履歴・replay・undo。同等 |
| Visitor | type switch | File/Directory、パス付きファイル一覧。同等 |
| Observer | チャネル/ゴルーチン | 乱数生成、DigitObserver(数値)/GraphObserver(\*)表示。同等 |
| Iterator | iter.Seq（Go 1.23） | BookShelf/Book、4冊走査。同等 |

### パターン構造を忠実に再現するパターン（10パターン）

以下は「パターン構造を忠実に再現」すべきパターンで、構造・機能ともに問題がないことを確認したもの。

| パターン | 構造の忠実さ | Java版との機能比較 |
|----------|-------------|-------------------|
| Composite | Entry interface、File、Directory（children + recursive printList） | 同等 |
| Decorator | Display interface、StringDisplay、SideBorder/FullBorder（inner Display） | 同等 |
| Builder | Builder interface、TextBuilder/HTMLBuilder、Construct関数（Director） | 同等 |
| Adapter | Printer interface、Banner（adaptee）、BannerAdapter（embedding） | 同等 |
| Facade | loadDatabase、htmlWriter、makeWelcomePage（facade関数） | 同等 |
| Proxy | Printable interface、Printer（heavyJob）、PrinterProxy（lazy init） | 同等 |
| Mediator | Colleague interface、RadioButton/TextField/Button、LoginDialog（mediator） | 同等 |
| State | State interface、DayState/NightState、SafeFrame（context） | 同等 |
| Flyweight | BigChar（flyweight）、sync.Map（pool）、BigString | 同等 |
| Interpreter | Node interface、Context、ProgramNode/CommandListNode/RepeatNode/PrimitiveNode | 同等 |
