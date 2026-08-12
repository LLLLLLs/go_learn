package com.huajian.novelreader

import java.io.InputStream
import java.nio.charset.StandardCharsets
import java.nio.file.Path
import java.util.zip.ZipEntry
import java.util.zip.ZipFile
import javax.xml.parsers.DocumentBuilderFactory
import org.w3c.dom.Element

object EpubChapterParser {
    fun parse(file: Path): ParsedBook {
        ZipFile(file.toFile()).use { zip ->
            val opfPath = findOpfPath(zip) ?: error("EPUB 缺少 OPF 清单")
            val opfDoc = parseXml(zip.readEntryText(opfPath) ?: error("OPF 文件读取失败: $opfPath"))
            val packageElement = opfDoc.documentElement
            val opfDir = opfPath.substringBeforeLast('/', "")

            val manifest = readManifest(packageElement)
            val spine = readSpine(packageElement)

            val chapters = mutableListOf<Chapter>()
            val content = StringBuilder()

            val orderedItems = spine.mapNotNull { idref ->
                manifest[idref]?.let { href -> resolvePath(opfDir, href) }
            }

            val chapterSources = if (orderedItems.isNotEmpty()) {
                orderedItems
            } else {
                // Fallback for malformed EPUBs without spine.
                zip.entries().asSequence()
                    .map(ZipEntry::getName)
                    .filter { it.endsWith(".xhtml", true) || it.endsWith(".html", true) || it.endsWith(".htm", true) }
                    .sorted()
                    .toList()
            }

            for ((index, itemPath) in chapterSources.withIndex()) {
                val rawHtml = zip.readEntryText(itemPath) ?: continue
                val title = extractTitle(rawHtml).ifBlank { "第${index + 1}章" }
                val plainText = htmlToPlainText(rawHtml).trim()
                if (plainText.isEmpty()) continue

                val start = content.length
                content.append(title).append("\n\n").append(plainText).append("\n\n")
                val end = content.length
                chapters.add(Chapter(title = title, start = start, end = end))
            }

            if (chapters.isEmpty()) {
                error("EPUB 没有可读章节内容")
            }

            return ParsedBook(content = content.toString(), chapters = chapters)
        }
    }

    private fun findOpfPath(zip: ZipFile): String? {
        val containerPath = "META-INF/container.xml"
        val fromContainer = zip.readEntryText(containerPath)?.let { containerXml ->
            val doc = parseXml(containerXml)
            val rootFiles = doc.getElementsByTagName("rootfile")
            if (rootFiles.length > 0) {
                (rootFiles.item(0) as? Element)?.getAttribute("full-path")
            } else {
                null
            }
        }
        if (!fromContainer.isNullOrBlank()) return fromContainer

        return zip.entries().asSequence()
            .map(ZipEntry::getName)
            .firstOrNull { it.endsWith(".opf", true) }
    }

    private fun readManifest(packageElement: Element): Map<String, String> {
        val manifestElements = packageElement.getElementsByTagName("item")
        val result = mutableMapOf<String, String>()
        for (i in 0 until manifestElements.length) {
            val item = manifestElements.item(i) as? Element ?: continue
            val id = item.getAttribute("id")
            val href = item.getAttribute("href")
            if (id.isNotBlank() && href.isNotBlank()) {
                result[id] = href
            }
        }
        return result
    }

    private fun readSpine(packageElement: Element): List<String> {
        val itemRefs = packageElement.getElementsByTagName("itemref")
        val result = mutableListOf<String>()
        for (i in 0 until itemRefs.length) {
            val itemRef = itemRefs.item(i) as? Element ?: continue
            val idref = itemRef.getAttribute("idref")
            if (idref.isNotBlank()) {
                result.add(idref)
            }
        }
        return result
    }

    private fun resolvePath(parent: String, href: String): String {
        if (href.startsWith("/")) return href.removePrefix("/")
        if (parent.isBlank()) return href
        return "$parent/$href"
    }

    private fun parseXml(text: String) = DocumentBuilderFactory.newInstance().apply {
        isNamespaceAware = false
        setFeature("http://apache.org/xml/features/disallow-doctype-decl", true)
    }.newDocumentBuilder().parse(text.byteInputStream())

    private fun extractTitle(html: String): String {
        val titlePattern = Regex("(?is)<title[^>]*>(.*?)</title>")
        val hPattern = Regex("(?is)<h[1-3][^>]*>(.*?)</h[1-3]>")
        val title = titlePattern.find(html)?.groupValues?.get(1)
            ?: hPattern.find(html)?.groupValues?.get(1)
            ?: ""
        return cleanupHtmlText(title)
    }

    private fun htmlToPlainText(html: String): String {
        var text = html
        text = text.replace(Regex("(?is)<script[^>]*>.*?</script>"), "")
        text = text.replace(Regex("(?is)<style[^>]*>.*?</style>"), "")
        text = text.replace(Regex("(?i)<br\\s*/?>"), "\n")
        text = text.replace(Regex("(?i)</p\\s*>"), "\n\n")
        text = text.replace(Regex("(?s)<[^>]+>"), "")
        text = cleanupHtmlText(text)
        text = text.replace(Regex("[ \\t\\x0B\\f\\r]+"), " ")
        text = text.replace(Regex("\\n{3,}"), "\n\n")
        return text
    }

    private fun cleanupHtmlText(value: String): String {
        return value
            .replace("&nbsp;", " ")
            .replace("&amp;", "&")
            .replace("&lt;", "<")
            .replace("&gt;", ">")
            .replace("&quot;", "\"")
            .replace("&#39;", "'")
            .replace(Regex("&#(\\d+);")) { match ->
                match.groupValues[1].toIntOrNull()?.toChar()?.toString() ?: ""
            }
            .replace(Regex("&#x([0-9a-fA-F]+);")) { match ->
                match.groupValues[1].toIntOrNull(16)?.toChar()?.toString() ?: ""
            }
            .trim()
    }

    private fun ZipFile.readEntryText(path: String): String? {
        val entry = getEntry(path) ?: return null
        return getInputStream(entry).use(InputStream::readBytes).toString(StandardCharsets.UTF_8)
    }
}
