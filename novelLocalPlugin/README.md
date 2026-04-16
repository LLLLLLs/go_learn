# reader (GoLand Plugin)

一个 GoLand/IntelliJ 平台插件，用于阅读本地 `txt/epub` 小说。
当前构建目标：`GoLand 2026.1`（`261.*`），默认使用本机 `/Applications/GoLand.app` 作为运行时。

## 已实现功能

- 本地书架：添加/移除 `txt/epub` 书籍
- 阅读进度保存：按书保存最近章节
- 章节切换：上一章/下一章 + 章节下拉框跳转
- 夜间模式：阅读区和书架列表切换暗色

## 工程结构

- `src/main/kotlin/com/huajian/novelreader`：核心代码
- `src/main/resources/META-INF/plugin.xml`：插件注册

## 使用方式

1. 在项目根目录生成/使用 Gradle Wrapper（建议 8.13+）：
   - `gradle wrapper --gradle-version 8.13`
2. 启动沙盒 IDE 调试插件：
   - `./gradlew runIde`
3. 打包插件：
   - `./gradlew buildPlugin`

启动后在右侧 Tool Window 找到 `reader`。

## 说明

- 章节识别默认支持：
  - 中文：`第xx章/节/卷/回/部/篇`
  - 英文：`Chapter 1` 形式
- EPUB：按 spine 顺序解析章节文档，支持常见 `xhtml/html` 结构。
- 编码读取优先 UTF-8，疑似乱码时回退 GBK。
- 构建使用 `org.jetbrains.intellij.platform` 2.x（兼容新版 Gradle，避免 `DefaultArtifactPublicationSet` 类加载错误）。
