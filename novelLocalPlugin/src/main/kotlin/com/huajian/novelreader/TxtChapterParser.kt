package com.huajian.novelreader

import java.nio.charset.Charset
import java.nio.charset.StandardCharsets
import java.nio.file.Files
import java.nio.file.Path

object TxtChapterParser {
    private val chapterPattern = Regex(
        pattern = """(?m)^\s*(第[0-9一二三四五六七八九十百千万零〇两]+[章节卷回部篇].*|Chapter\s+\d+.*)\s*$""",
        options = setOf(RegexOption.IGNORE_CASE)
    )

    fun parse(file: Path): ParsedBook {
        val content = readTextWithFallback(file)
        val matches = chapterPattern.findAll(content).toList()

        if (matches.isEmpty()) {
            return ParsedBook(content = content, chapters = listOf(Chapter("全文", 0, content.length)))
        }

        val chapters = mutableListOf<Chapter>()
        for (i in matches.indices) {
            val match = matches[i]
            val title = match.value.trim().ifEmpty { "第${i + 1}章" }
            val start = match.range.first
            val end = if (i == matches.lastIndex) content.length else matches[i + 1].range.first
            chapters.add(Chapter(title = title, start = start, end = end))
        }

        return ParsedBook(content = content, chapters = chapters)
    }

    private fun readTextWithFallback(file: Path): String {
        val bytes = Files.readAllBytes(file)
        val utf8 = decode(bytes, StandardCharsets.UTF_8)
        if (replacementRatio(utf8) < 0.002) {
            return utf8
        }

        val gbk = decode(bytes, Charset.forName("GBK"))
        return if (replacementRatio(gbk) < replacementRatio(utf8)) gbk else utf8
    }

    private fun decode(bytes: ByteArray, charset: Charset): String {
        return String(bytes, charset)
    }

    private fun replacementRatio(text: String): Double {
        if (text.isEmpty()) return 0.0
        val replacementCount = text.count { it == '\uFFFD' }
        return replacementCount.toDouble() / text.length.toDouble()
    }
}
