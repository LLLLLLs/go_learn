package com.huajian.novelreader

import java.awt.BorderLayout
import java.awt.CardLayout
import java.awt.Color
import java.awt.Dimension
import java.awt.FlowLayout
import java.awt.event.MouseAdapter
import java.awt.event.MouseEvent
import java.io.File
import java.nio.file.Files
import java.nio.file.Paths
import javax.swing.BorderFactory
import javax.swing.DefaultListModel
import javax.swing.JButton
import javax.swing.JCheckBox
import javax.swing.JComboBox
import javax.swing.JFileChooser
import javax.swing.JList
import javax.swing.JOptionPane
import javax.swing.JPanel
import javax.swing.JScrollPane
import javax.swing.JLabel
import javax.swing.JTextArea
import javax.swing.JTextField
import javax.swing.ListSelectionModel
import javax.swing.SwingConstants
import javax.swing.SwingUtilities
import javax.swing.event.DocumentEvent
import javax.swing.event.DocumentListener
import javax.swing.filechooser.FileNameExtensionFilter

class NovelReaderPanel(
    private val stateService: NovelStateService
) : JPanel(BorderLayout()) {

    private val bookListModel = DefaultListModel<BookState>()
    private val bookList = JList(bookListModel)
    private val textArea = JTextArea()
    private val chapterTitleLabel = JLabel("", SwingConstants.CENTER)
    private val chapterSelector = JComboBox<String>()
    private val chooseChapterButton = JButton("选择章节")
    private val prevButton = JButton("上一章")
    private val nextButton = JButton("下一章")
    private val backToShelfButton = JButton("返回书架")
    private val nightModeCheckBox = JCheckBox("夜间模式")

    private val cardLayout = CardLayout()
    private val contentPanel = JPanel(cardLayout)
    private val bookshelfPanel = JPanel(BorderLayout())
    private val readerPanel = JPanel(BorderLayout())

    private val parsedCache = mutableMapOf<String, ParsedBook>()
    private val filteredChapterIndexes = mutableListOf<Int>()

    private var currentBookPath: String? = null
    private var currentChapterIndex: Int = 0
    private var updatingChapterSelector = false

    init {
        setupUi()
        wireEvents()
        loadBooksFromState()
        restoreUiState()
    }

    private fun setupUi() {
        buildBookshelfPanel()
        buildReaderPanel()

        contentPanel.add(bookshelfPanel, CARD_BOOKSHELF)
        contentPanel.add(readerPanel, CARD_READER)
        add(contentPanel, BorderLayout.CENTER)
    }

    private fun buildBookshelfPanel() {
        bookshelfPanel.border = BorderFactory.createEmptyBorder(8, 8, 8, 8)

        bookList.selectionMode = ListSelectionModel.SINGLE_SELECTION
        bookList.cellRenderer = BookListCellRenderer()
        bookshelfPanel.add(JScrollPane(bookList), BorderLayout.CENTER)

        val buttonPanel = JPanel(FlowLayout(FlowLayout.LEFT))
        val addButton = JButton("添加")
        val removeButton = JButton("移除")
        val openButton = JButton("打开")
        buttonPanel.add(addButton)
        buttonPanel.add(removeButton)
        buttonPanel.add(openButton)
        bookshelfPanel.add(buttonPanel, BorderLayout.SOUTH)

        addButton.addActionListener { addBookFromChooser() }
        removeButton.addActionListener { removeSelectedBook() }
        openButton.addActionListener { openSelectedBook() }
    }

    private fun buildReaderPanel() {
        readerPanel.border = BorderFactory.createEmptyBorder(8, 8, 8, 8)
        chapterTitleLabel.font = chapterTitleLabel.font.deriveFont(16f)
        chapterTitleLabel.border = BorderFactory.createEmptyBorder(0, 0, 8, 0)
        readerPanel.add(chapterTitleLabel, BorderLayout.NORTH)

        val bottomPanel = JPanel(FlowLayout(FlowLayout.LEFT))
        chapterSelector.preferredSize = Dimension(280, 28)
        bottomPanel.add(backToShelfButton)
        bottomPanel.add(prevButton)
        bottomPanel.add(nextButton)
        bottomPanel.add(chooseChapterButton)
        bottomPanel.add(chapterSelector)
        bottomPanel.add(nightModeCheckBox)
        readerPanel.add(bottomPanel, BorderLayout.SOUTH)

        textArea.isEditable = false
        textArea.isFocusable = false
        textArea.lineWrap = true
        textArea.wrapStyleWord = true
        textArea.font = textArea.font.deriveFont(16f)
        textArea.margin = java.awt.Insets(12, 12, 12, 12)
        readerPanel.add(JScrollPane(textArea), BorderLayout.CENTER)
    }

    private fun wireEvents() {
        bookList.addMouseListener(object : MouseAdapter() {
            override fun mouseClicked(e: MouseEvent) {
                if (e.clickCount == 2 && SwingUtilities.isLeftMouseButton(e)) {
                    openSelectedBook()
                }
            }
        })

        backToShelfButton.addActionListener {
            showBookshelf()
        }

        prevButton.addActionListener {
            switchChapter(currentChapterIndex - 1)
        }

        nextButton.addActionListener {
            switchChapter(currentChapterIndex + 1)
        }

        chooseChapterButton.addActionListener {
            chooseChapterFromDialog()
        }

        chapterSelector.addActionListener {
            if (updatingChapterSelector) return@addActionListener
            val idx = chapterSelector.selectedIndex
            if (idx >= 0) {
                val actualIndex = filteredChapterIndexes.getOrNull(idx) ?: return@addActionListener
                switchChapter(actualIndex)
            }
        }

        nightModeCheckBox.addActionListener {
            val enabled = nightModeCheckBox.isSelected
            stateService.setNightMode(enabled)
            applyNightMode(enabled)
        }
    }

    private fun loadBooksFromState() {
        bookListModel.clear()
        stateService.books()
            .filter { it.path.isNotBlank() }
            .forEach { saved ->
                if (Files.exists(Paths.get(saved.path))) {
                    bookListModel.addElement(saved)
                }
            }
    }

    private fun restoreUiState() {
        val night = stateService.nightModeEnabled()
        nightModeCheckBox.isSelected = night
        applyNightMode(night)

        val selectedPath = stateService.selectedBookPath()
        if (selectedPath.isNullOrBlank()) {
            showBookshelf()
            return
        }

        val selectedIndex = (0 until bookListModel.size()).firstOrNull { idx ->
            bookListModel.getElementAt(idx).path == selectedPath
        }

        if (selectedIndex == null) {
            stateService.setSelectedBook(null)
            showBookshelf()
            return
        }

        bookList.selectedIndex = selectedIndex
        openBook(selectedPath)
    }

    private fun addBookFromChooser() {
        val chooser = JFileChooser()
        chooser.fileSelectionMode = JFileChooser.FILES_ONLY
        chooser.dialogTitle = "选择 TXT/EPUB 小说文件"
        chooser.fileFilter = FileNameExtensionFilter("小说文件 (*.txt, *.epub)", "txt", "epub")
        stateService.lastAddDirectory()
            ?.let(::File)
            ?.takeIf { it.exists() && it.isDirectory }
            ?.let { chooser.currentDirectory = it }

        val result = chooser.showOpenDialog(this)
        if (result != JFileChooser.APPROVE_OPTION) {
            return
        }

        val file = chooser.selectedFile ?: return
        val extension = file.extension.lowercase()
        if (!file.exists() || (extension != "txt" && extension != "epub")) {
            JOptionPane.showMessageDialog(this, "请选择 .txt 或 .epub 文件")
            return
        }

        val path = file.toPath().toAbsolutePath().toString()
        stateService.setLastAddDirectory(file.parentFile?.absolutePath)
        stateService.upsertBook(path = path, title = file.nameWithoutExtension)
        loadBooksFromState()

        val addedIndex = (0 until bookListModel.size()).firstOrNull { idx ->
            bookListModel.getElementAt(idx).path == path
        } ?: return
        bookList.selectedIndex = addedIndex
    }

    private fun removeSelectedBook() {
        val selected = bookList.selectedValue ?: return
        stateService.removeBook(selected.path)
        parsedCache.remove(selected.path)
        loadBooksFromState()

        if (currentBookPath == selected.path) {
            currentBookPath = null
            chapterTitleLabel.text = ""
            textArea.text = ""
            chapterSelector.removeAllItems()
            showBookshelf()
        }
    }

    private fun openSelectedBook() {
        val selected = bookList.selectedValue ?: return
        openBook(selected.path)
    }

    private fun openBook(path: String) {
        val parsed = runCatching {
            parsedCache.getOrPut(path) { BookParser.parse(Paths.get(path)) }
        }.getOrElse { ex ->
            JOptionPane.showMessageDialog(this, "打开失败: ${ex.message}")
            return
        }

        currentBookPath = path
        stateService.setSelectedBook(path)

        refreshChapterSelector("")

        val startIndex = stateService.getProgress(path).coerceIn(0, parsed.chapters.lastIndex)
        switchChapter(startIndex)
        showReader()
    }

    private fun switchChapter(targetIndex: Int) {
        val path = currentBookPath ?: return
        val parsed = parsedCache[path] ?: return
        if (parsed.chapters.isEmpty()) return

        val safeIndex = targetIndex.coerceIn(0, parsed.chapters.lastIndex)
        currentChapterIndex = safeIndex

        val chapter = parsed.chapters[safeIndex]
        chapterTitleLabel.text = chapter.title
        val chapterText = parsed.content.substring(chapter.start, chapter.end).trim()
        textArea.text = chapterText
        textArea.caretPosition = 0

        updatingChapterSelector = true
        val dropdownIndex = filteredChapterIndexes.indexOf(safeIndex)
        chapterSelector.selectedIndex = dropdownIndex
        updatingChapterSelector = false

        stateService.updateProgress(path, safeIndex)

        prevButton.isEnabled = safeIndex > 0
        nextButton.isEnabled = safeIndex < parsed.chapters.lastIndex
        chooseChapterButton.isEnabled = parsed.chapters.isNotEmpty()
    }

    private fun chooseChapterFromDialog() {
        val path = currentBookPath ?: return
        val parsed = parsedCache[path] ?: return
        if (parsed.chapters.isEmpty()) return

        val searchField = JTextField(24)
        val resultBox = JComboBox<String>()
        val matchIndexes = mutableListOf<Int>()

        fun refill(query: String) {
            val matches = if (query.trim().isEmpty()) {
                parsed.chapters.indices.toList()
            } else {
                findChapterCandidates(query, parsed.chapters)
            }
            matchIndexes.clear()
            matchIndexes.addAll(matches)

            resultBox.removeAllItems()
            matchIndexes.take(200).forEach { idx ->
                resultBox.addItem("${idx + 1}. ${parsed.chapters[idx].title}")
            }
            val selected = matchIndexes.indexOf(currentChapterIndex)
            if (selected >= 0 && selected < resultBox.itemCount) {
                resultBox.selectedIndex = selected
            } else if (resultBox.itemCount > 0) {
                resultBox.selectedIndex = 0
            }
        }

        searchField.document.addDocumentListener(object : DocumentListener {
            override fun insertUpdate(e: DocumentEvent?) = refill(searchField.text)
            override fun removeUpdate(e: DocumentEvent?) = refill(searchField.text)
            override fun changedUpdate(e: DocumentEvent?) = refill(searchField.text)
        })

        refill("")

        val panel = JPanel(BorderLayout(0, 8))
        panel.add(searchField, BorderLayout.NORTH)
        panel.add(resultBox, BorderLayout.SOUTH)

        val result = JOptionPane.showConfirmDialog(
            this,
            panel,
            "章节选择",
            JOptionPane.OK_CANCEL_OPTION,
            JOptionPane.PLAIN_MESSAGE
        )
        if (result != JOptionPane.OK_OPTION) return

        val selected = resultBox.selectedItem as? String ?: return
        val targetIndex = selected.substringBefore('.').trim().toIntOrNull()?.minus(1) ?: return
        if (targetIndex in parsed.chapters.indices) {
            switchChapter(targetIndex)
        }
    }

    private fun refreshChapterSelector(query: String) {
        val path = currentBookPath ?: return
        val parsed = parsedCache[path] ?: return

        val matches = if (query.trim().isEmpty()) {
            parsed.chapters.indices.toList()
        } else {
            findChapterCandidates(query, parsed.chapters)
        }

        filteredChapterIndexes.clear()
        filteredChapterIndexes.addAll(matches)

        updatingChapterSelector = true
        chapterSelector.removeAllItems()
        filteredChapterIndexes.forEach { idx ->
            val chapter = parsed.chapters[idx]
            chapterSelector.addItem("${idx + 1}. ${chapter.title}")
        }
        val selectedDropdownIndex = filteredChapterIndexes.indexOf(currentChapterIndex)
        chapterSelector.selectedIndex = selectedDropdownIndex
        updatingChapterSelector = false
    }

    private fun findChapterCandidates(query: String, chapters: List<Chapter>): List<Int> {
        val raw = query.trim()
        if (raw.isEmpty()) return emptyList()

        val normalized = normalize(raw)
        val queryNumber = raw.toIntOrNull() ?: Regex("(\\d+)").find(raw)?.groupValues?.get(1)?.toIntOrNull()

        val scored = mutableListOf<Pair<Int, Int>>()
        chapters.forEachIndexed { index, chapter ->
            val title = chapter.title
            val titleNorm = normalize(title)
            var score = 0

            if (queryNumber != null) {
                if (index + 1 == queryNumber) score += 1000
                val titleNumber = Regex("(\\d+)").find(title)?.groupValues?.get(1)?.toIntOrNull()
                if (titleNumber != null && titleNumber == queryNumber) score += 600
            }

            if (titleNorm == normalized) score += 900
            if (titleNorm.startsWith(normalized)) score += 700
            if (titleNorm.contains(normalized)) score += 500
            if (isSubsequence(normalized, titleNorm)) score += 200

            if (score > 0) scored.add(index to score)
        }

        return scored
            .sortedWith(compareByDescending<Pair<Int, Int>> { it.second }.thenBy { it.first })
            .map { it.first }
    }

    private fun normalize(value: String): String {
        return value.lowercase().replace(Regex("\\s+"), "")
    }

    private fun isSubsequence(pattern: String, text: String): Boolean {
        if (pattern.isEmpty()) return true
        var i = 0
        var j = 0
        while (i < pattern.length && j < text.length) {
            if (pattern[i] == text[j]) i++
            j++
        }
        return i == pattern.length
    }

    private fun applyNightMode(enabled: Boolean) {
        val textBg = if (enabled) Color(34, 34, 34) else Color.WHITE
        val textFg = if (enabled) Color(220, 220, 220) else Color.BLACK
        val listBg = if (enabled) Color(42, 42, 42) else Color.WHITE
        val listFg = if (enabled) Color(220, 220, 220) else Color.BLACK

        textArea.background = textBg
        textArea.foreground = textFg
        textArea.caretColor = Color(0, 0, 0, 0)
        textArea.selectionColor = textBg
        textArea.selectedTextColor = textFg
        chapterTitleLabel.foreground = textFg

        bookList.selectionBackground = listBg
        bookList.selectionForeground = listFg

        bookList.background = listBg
        bookList.foreground = listFg
        chapterSelector.background = listBg
        chapterSelector.foreground = listFg

        bookshelfPanel.background = listBg
        readerPanel.background = listBg

        SwingUtilities.updateComponentTreeUI(this)
    }

    private fun showBookshelf() {
        cardLayout.show(contentPanel, CARD_BOOKSHELF)
    }

    private fun showReader() {
        cardLayout.show(contentPanel, CARD_READER)
    }

    companion object {
        private const val CARD_BOOKSHELF = "bookshelf"
        private const val CARD_READER = "reader"
    }
}
