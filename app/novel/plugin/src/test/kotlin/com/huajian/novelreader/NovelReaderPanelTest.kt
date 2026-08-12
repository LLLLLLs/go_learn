package com.huajian.novelreader

import java.io.File
import java.lang.reflect.Method
import java.nio.charset.StandardCharsets
import java.nio.file.Files
import java.util.zip.ZipEntry
import java.util.zip.ZipOutputStream
import javax.swing.JButton
import javax.swing.JFrame
import javax.swing.JList
import javax.swing.JScrollPane
import javax.swing.SwingUtilities
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertTrue

class NovelReaderPanelTest {
    @Test
    fun chapterSelectorButtonTracksCurrentChapter() {
        val bookFile = createTestEpub(chapterCount = 21)

        val panel = NovelReaderPanel(NovelStateService())
        invokeOnEdt {
            callPrivateString(panel, "openBook", bookFile.absolutePath)
            callPrivateInt(panel, "switchChapter", 20)
        }

        val button = getPrivateField(panel, "chapterSelectorButton") as JButton
        assertTrue(
            button.text.startsWith("21."),
            "chapter selector button should point to current chapter, actual text: ${button.text}"
        )
    }

    @Test
    fun chapterSelectorPopupTracksHighChapter() {
        val bookFile = createTestEpub(chapterCount = 1_000)

        val panel = NovelReaderPanel(NovelStateService())
        val frame = JFrame("test").apply {
            contentPane.add(panel)
            setSize(800, 600)
            setLocationRelativeTo(null)
        }

        invokeOnEdt {
            frame.isVisible = true
            callPrivateString(panel, "openBook", bookFile.absolutePath)
            callPrivateInt(panel, "switchChapter", 199)
            callPrivateNoArg(panel, "showChapterSelectorPopup")
        }

        val popupListAt200 = getPrivateField(panel, "chapterSelectorPopupList") as JList<*>
        assertEquals(199, popupListAt200.selectedIndex, "popup should select chapter 200")
        val scrollPaneAt200 = popupListAt200.parent.parent as JScrollPane
        assertTrue(scrollPaneAt200.viewport.viewPosition.y > 0, "popup should scroll when selecting chapter 200")

        invokeOnEdt {
            callPrivateInt(panel, "switchChapter", 999)
            callPrivateNoArg(panel, "showChapterSelectorPopup")
        }

        val popupListAt1000 = getPrivateField(panel, "chapterSelectorPopupList") as JList<*>
        assertEquals(999, popupListAt1000.selectedIndex, "popup should select chapter 1000")
        val scrollPaneAt1000 = popupListAt1000.parent.parent as JScrollPane
        assertTrue(scrollPaneAt1000.viewport.viewPosition.y > 0, "popup should scroll when selecting chapter 1000")

        val button = getPrivateField(panel, "chapterSelectorButton") as JButton
        assertTrue(button.text.startsWith("1000."), "chapter selector button should point to chapter 1000, actual text: ${button.text}")

        invokeOnEdt {
            frame.dispose()
        }
    }

    private fun createTestEpub(chapterCount: Int): File {
        val file = Files.createTempFile("novel-reader-test-", ".epub").toFile().apply {
            deleteOnExit()
        }

        ZipOutputStream(file.outputStream().buffered()).use { zip ->
            zip.writeEntry(
                "META-INF/container.xml",
                """
                    <container>
                      <rootfiles>
                        <rootfile full-path="OEBPS/content.opf"/>
                      </rootfiles>
                    </container>
                """.trimIndent()
            )

            val manifest = (1..chapterCount).joinToString("\n") { chapter ->
                "<item id=\"chapter$chapter\" href=\"chapter$chapter.xhtml\"/>"
            }
            val spine = (1..chapterCount).joinToString("\n") { chapter ->
                "<itemref idref=\"chapter$chapter\"/>"
            }
            zip.writeEntry(
                "OEBPS/content.opf",
                """
                    <package>
                      <manifest>
                        $manifest
                      </manifest>
                      <spine>
                        $spine
                      </spine>
                    </package>
                """.trimIndent()
            )

            for (chapter in 1..chapterCount) {
                zip.writeEntry(
                    "OEBPS/chapter$chapter.xhtml",
                    """
                        <html>
                          <head><title>$chapter. Test chapter</title></head>
                          <body><p>Chapter $chapter content.</p></body>
                        </html>
                    """.trimIndent()
                )
            }
        }

        return file
    }

    private fun ZipOutputStream.writeEntry(name: String, content: String) {
        putNextEntry(ZipEntry(name))
        write(content.toByteArray(StandardCharsets.UTF_8))
        closeEntry()
    }

    private fun invokeOnEdt(block: () -> Unit) {
        if (SwingUtilities.isEventDispatchThread()) {
            block()
        } else {
            SwingUtilities.invokeAndWait(block)
        }
    }

    private fun callPrivateString(target: Any, name: String, arg: String) {
        val method: Method = target.javaClass.getDeclaredMethod(name, String::class.java)
        method.isAccessible = true
        method.invoke(target, arg)
    }

    private fun callPrivateInt(target: Any, name: String, arg: Int) {
        val method: Method = target.javaClass.getDeclaredMethod(name, Int::class.javaPrimitiveType)
        method.isAccessible = true
        method.invoke(target, arg)
    }

    private fun callPrivateNoArg(target: Any, name: String) {
        val method: Method = target.javaClass.getDeclaredMethod(name)
        method.isAccessible = true
        method.invoke(target)
    }

    private fun getPrivateField(target: Any, name: String): Any? {
        val field = target.javaClass.getDeclaredField(name)
        field.isAccessible = true
        return field.get(target)
    }
}
