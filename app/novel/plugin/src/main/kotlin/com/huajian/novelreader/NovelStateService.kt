package com.huajian.novelreader

import com.intellij.openapi.components.PersistentStateComponent
import com.intellij.openapi.components.State
import com.intellij.openapi.components.Storage

@State(name = "NovelReaderState", storages = [Storage("novelReader.xml")])
class NovelStateService : PersistentStateComponent<NovelPersistedState> {
    private var persistedState = NovelPersistedState()

    override fun getState(): NovelPersistedState = persistedState

    override fun loadState(state: NovelPersistedState) {
        persistedState = state
    }

    fun books(): List<BookState> = persistedState.books.toList()

    fun upsertBook(path: String, title: String) {
        val existing = persistedState.books.firstOrNull { it.path == path }
        if (existing != null) {
            existing.title = title
            return
        }
        persistedState.books.add(BookState(path = path, title = title, lastChapterIndex = 0))
    }

    fun removeBook(path: String) {
        persistedState.books.removeIf { it.path == path }
        if (persistedState.selectedBookPath == path) {
            persistedState.selectedBookPath = null
        }
    }

    fun setSelectedBook(path: String?) {
        persistedState.selectedBookPath = path
    }

    fun selectedBookPath(): String? = persistedState.selectedBookPath

    fun updateProgress(path: String, chapterIndex: Int) {
        persistedState.books.firstOrNull { it.path == path }?.lastChapterIndex = chapterIndex.coerceAtLeast(0)
    }

    fun getProgress(path: String): Int {
        return persistedState.books.firstOrNull { it.path == path }?.lastChapterIndex ?: 0
    }

    fun nightModeEnabled(): Boolean = persistedState.nightMode

    fun setNightMode(enabled: Boolean) {
        persistedState.nightMode = enabled
    }

    fun lastAddDirectory(): String? = persistedState.lastAddDirectory

    fun setLastAddDirectory(path: String?) {
        persistedState.lastAddDirectory = path
    }
}
