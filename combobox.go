/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 */

package winc

import (
	"syscall"
	"unsafe"

	"github.com/samuel-jimenez/winc/w32"
)

type ComboBox struct {
	ControlBase
	onSelectedChange EventManager
}

func NewComboBox(parent Controller) *ComboBox {
	control := new(ComboBox)

	control.InitControl("COMBOBOX", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_TABSTOP|w32.WS_VSCROLL|w32.CBS_DROPDOWN)
	RegMsgHandler(control)

	control.SetFont(DefaultFont)
	control.SetSize(200, 400)
	return control
}

func NewComboBoxWithFlags(parent Controller, style uint) *ComboBox {
	control := new(ComboBox)

	control.InitControl("COMBOBOX", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_TABSTOP|w32.WS_VSCROLL|w32.CBS_DROPDOWN|style)
	RegMsgHandler(control)

	control.SetFont(DefaultFont)
	control.SetSize(200, 400)
	return control
}

func NewListComboBox(parent Controller) *ComboBox {
	control := new(ComboBox)

	control.InitControl("COMBOBOX", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_TABSTOP|w32.WS_VSCROLL|w32.CBS_DROPDOWNLIST)
	RegMsgHandler(control)

	control.SetFont(DefaultFont)
	control.SetSize(200, 400)
	return control
}

func (control *ComboBox) DeleteAllItems() bool {
	return w32.SendMessage(control.hwnd, w32.CB_RESETCONTENT, 0, 0) == w32.TRUE
}

func (control *ComboBox) AddItem(str string) bool {
	lp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	return w32.SendMessage(control.hwnd, w32.CB_ADDSTRING, 0, lp) != w32.CB_ERR
}

func (control *ComboBox) InsertItem(index int, str string) bool {
	lp := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	return w32.SendMessage(control.hwnd, w32.CB_INSERTSTRING, uintptr(index), lp) != w32.CB_ERR
}

func (control *ComboBox) DeleteItem(index int) bool {
	return w32.SendMessage(control.hwnd, w32.CB_DELETESTRING, uintptr(index), 0) != w32.CB_ERR
}

func (control *ComboBox) SelectedItem() int {
	return int(int32(w32.SendMessage(control.hwnd, w32.CB_GETCURSEL, 0, 0)))
}

func (control *ComboBox) SetSelectedItem(value int) bool {
	return int(int32(w32.SendMessage(control.hwnd, w32.CB_SETCURSEL, uintptr(value), 0))) == value
}

func (control *ComboBox) OnSelectedChange() *EventManager {
	return &control.onSelectedChange
}

// Message processer
func (control *ComboBox) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_COMMAND:
		code := w32.HIWORD(uint32(wparam))

		switch code {
		case w32.CBN_SELCHANGE:
			control.onSelectedChange.Fire(NewEvent(control, nil))
		}
	}
	return w32.DefWindowProc(control.hwnd, msg, wparam, lparam)
	//return control.W32Control.WndProc(msg, wparam, lparam)
}
