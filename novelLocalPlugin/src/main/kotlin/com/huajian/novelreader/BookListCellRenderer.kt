package com.huajian.novelreader

import java.awt.Component
import javax.swing.DefaultListCellRenderer
import javax.swing.JList

class BookListCellRenderer : DefaultListCellRenderer() {
    override fun getListCellRendererComponent(
        list: JList<*>,
        value: Any?,
        index: Int,
        isSelected: Boolean,
        cellHasFocus: Boolean
    ): Component {
        val c = super.getListCellRendererComponent(list, value, index, false, false)
        if (value is BookState) {
            text = value.title
            toolTipText = value.path
        }
        return c
    }
}
