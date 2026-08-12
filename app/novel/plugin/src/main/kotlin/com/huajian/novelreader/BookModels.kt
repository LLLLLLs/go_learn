package com.huajian.novelreader

data class BookState(
    var path: String = "",
    var title: String = "",
    var lastChapterIndex: Int = 0
)

data class NovelPersistedState(
    var books: MutableList<BookState> = mutableListOf(),
    var selectedBookPath: String? = null,
    var nightMode: Boolean = true,
    var lastAddDirectory: String? = null
)

data class Chapter(
    val title: String,
    val start: Int,
    val end: Int
)

data class ParsedBook(
    val content: String,
    val chapters: List<Chapter>
)
