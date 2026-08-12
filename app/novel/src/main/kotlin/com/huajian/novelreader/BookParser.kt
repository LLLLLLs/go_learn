package com.huajian.novelreader

import java.nio.file.Path

object BookParser {
    fun parse(file: Path): ParsedBook {
        return when (file.fileName.toString().substringAfterLast('.', "").lowercase()) {
            "txt" -> TxtChapterParser.parse(file)
            "epub" -> EpubChapterParser.parse(file)
            else -> error("不支持的格式: ${file.fileName}")
        }
    }
}
