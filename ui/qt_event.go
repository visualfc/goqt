// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

func drvNewObject(obj uintptr) *QObject {
	_obj := &QObject{}
	_obj.SetDriver(obj, 316000, false)
	return _obj
}
func drvEventCall(i interface{}, obj *QObject, evid uint32, event uintptr) bool {
	if v, ok := i.(interface {
		OnEvent(*QObject, *QEvent) bool
	}); ok {
		e := &QEvent{}
		e.SetDriver(event, 31000, false)
		if v.OnEvent(obj, e) {
			return true
		}
	} else if v, ok := i.(interface {
		OnEvent(*QEvent) bool
	}); ok {
		e := &QEvent{}
		e.SetDriver(event, 31000, false)
		if v.OnEvent(e) {
			return true
		}
	}
	switch QEvent_Type(evid) {
	case QEvent_Timer:
		if v, ok := i.(interface {
			OnTimerEvent(*QObject, *QTimerEvent) bool
		}); ok {
			e := &QTimerEvent{}
			e.SetDriver(event, 173000, false)
			return v.OnTimerEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTimerEvent(*QTimerEvent) bool
		}); ok {
			e := &QTimerEvent{}
			e.SetDriver(event, 173000, false)
			return v.OnTimerEvent(e)
		}
	case QEvent_MouseButtonPress:
		if v, ok := i.(interface {
			OnMouseButtonPressEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseButtonPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseButtonPressEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseButtonPressEvent(e)
		}
	case QEvent_MouseButtonRelease:
		if v, ok := i.(interface {
			OnMouseButtonReleaseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseButtonReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseButtonReleaseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseButtonReleaseEvent(e)
		}
	case QEvent_MouseButtonDblClick:
		if v, ok := i.(interface {
			OnMouseButtonDblClickEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseButtonDblClickEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseButtonDblClickEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseButtonDblClickEvent(e)
		}
	case QEvent_MouseMove:
		if v, ok := i.(interface {
			OnMouseMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseMoveEvent(e)
		}
	case QEvent_KeyPress:
		if v, ok := i.(interface {
			OnKeyPressEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnKeyPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnKeyPressEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnKeyPressEvent(e)
		}
	case QEvent_KeyRelease:
		if v, ok := i.(interface {
			OnKeyReleaseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnKeyReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnKeyReleaseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnKeyReleaseEvent(e)
		}
	case QEvent_FocusIn:
		if v, ok := i.(interface {
			OnFocusInEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFocusInEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFocusInEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFocusInEvent(e)
		}
	case QEvent_FocusOut:
		if v, ok := i.(interface {
			OnFocusOutEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFocusOutEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFocusOutEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFocusOutEvent(e)
		}
	case QEvent_Enter:
		if v, ok := i.(interface {
			OnEnterEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnEnterEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEnterEvent(e)
		}
	case QEvent_Leave:
		if v, ok := i.(interface {
			OnLeaveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnLeaveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLeaveEvent(e)
		}
	case QEvent_Paint:
		if v, ok := i.(interface {
			OnPaintEvent(*QObject, *QPaintEvent) bool
		}); ok {
			e := &QPaintEvent{}
			e.SetDriver(event, 86000, false)
			return v.OnPaintEvent(obj, e)
		} else if v, ok := i.(interface {
			OnPaintEvent(*QPaintEvent) bool
		}); ok {
			e := &QPaintEvent{}
			e.SetDriver(event, 86000, false)
			return v.OnPaintEvent(e)
		}
	case QEvent_Move:
		if v, ok := i.(interface {
			OnMoveEvent(*QObject, *QMoveEvent) bool
		}); ok {
			e := &QMoveEvent{}
			e.SetDriver(event, 82000, false)
			return v.OnMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMoveEvent(*QMoveEvent) bool
		}); ok {
			e := &QMoveEvent{}
			e.SetDriver(event, 82000, false)
			return v.OnMoveEvent(e)
		}
	case QEvent_Resize:
		if v, ok := i.(interface {
			OnResizeEvent(*QObject, *QResizeEvent) bool
		}); ok {
			e := &QResizeEvent{}
			e.SetDriver(event, 116000, false)
			return v.OnResizeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnResizeEvent(*QResizeEvent) bool
		}); ok {
			e := &QResizeEvent{}
			e.SetDriver(event, 116000, false)
			return v.OnResizeEvent(e)
		}
	case QEvent_Create:
		if v, ok := i.(interface {
			OnCreateEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnCreateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnCreateEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnCreateEvent(e)
		}
	case QEvent_Destroy:
		if v, ok := i.(interface {
			OnDestroyEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDestroyEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDestroyEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDestroyEvent(e)
		}
	case QEvent_Show:
		if v, ok := i.(interface {
			OnShowEvent(*QObject, *QShowEvent) bool
		}); ok {
			e := &QShowEvent{}
			e.SetDriver(event, 119000, false)
			return v.OnShowEvent(obj, e)
		} else if v, ok := i.(interface {
			OnShowEvent(*QShowEvent) bool
		}); ok {
			e := &QShowEvent{}
			e.SetDriver(event, 119000, false)
			return v.OnShowEvent(e)
		}
	case QEvent_Hide:
		if v, ok := i.(interface {
			OnHideEvent(*QObject, *QHideEvent) bool
		}); ok {
			e := &QHideEvent{}
			e.SetDriver(event, 49000, false)
			return v.OnHideEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHideEvent(*QHideEvent) bool
		}); ok {
			e := &QHideEvent{}
			e.SetDriver(event, 49000, false)
			return v.OnHideEvent(e)
		}
	case QEvent_Close:
		if v, ok := i.(interface {
			OnCloseEvent(*QObject, *QCloseEvent) bool
		}); ok {
			e := &QCloseEvent{}
			e.SetDriver(event, 12000, false)
			return v.OnCloseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnCloseEvent(*QCloseEvent) bool
		}); ok {
			e := &QCloseEvent{}
			e.SetDriver(event, 12000, false)
			return v.OnCloseEvent(e)
		}
	case QEvent_Quit:
		if v, ok := i.(interface {
			OnQuitEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnQuitEvent(obj, e)
		} else if v, ok := i.(interface {
			OnQuitEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnQuitEvent(e)
		}
	case QEvent_ParentChange:
		if v, ok := i.(interface {
			OnParentChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnParentChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnParentChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnParentChangeEvent(e)
		}
	case QEvent_ParentAboutToChange:
		if v, ok := i.(interface {
			OnParentAboutToChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnParentAboutToChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnParentAboutToChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnParentAboutToChangeEvent(e)
		}
	case QEvent_ThreadChange:
		if v, ok := i.(interface {
			OnThreadChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnThreadChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnThreadChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnThreadChangeEvent(e)
		}
	case QEvent_WindowActivate:
		if v, ok := i.(interface {
			OnWindowActivateEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowActivateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowActivateEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowActivateEvent(e)
		}
	case QEvent_WindowDeactivate:
		if v, ok := i.(interface {
			OnWindowDeactivateEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowDeactivateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowDeactivateEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowDeactivateEvent(e)
		}
	case QEvent_ShowToParent:
		if v, ok := i.(interface {
			OnShowToParentEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnShowToParentEvent(obj, e)
		} else if v, ok := i.(interface {
			OnShowToParentEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnShowToParentEvent(e)
		}
	case QEvent_HideToParent:
		if v, ok := i.(interface {
			OnHideToParentEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHideToParentEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHideToParentEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHideToParentEvent(e)
		}
	case QEvent_Wheel:
		if v, ok := i.(interface {
			OnWheelEvent(*QObject, *QWheelEvent) bool
		}); ok {
			e := &QWheelEvent{}
			e.SetDriver(event, 191000, false)
			return v.OnWheelEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWheelEvent(*QWheelEvent) bool
		}); ok {
			e := &QWheelEvent{}
			e.SetDriver(event, 191000, false)
			return v.OnWheelEvent(e)
		}
	case QEvent_WindowTitleChange:
		if v, ok := i.(interface {
			OnWindowTitleChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowTitleChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowTitleChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowTitleChangeEvent(e)
		}
	case QEvent_WindowIconChange:
		if v, ok := i.(interface {
			OnWindowIconChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowIconChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowIconChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowIconChangeEvent(e)
		}
	case QEvent_ApplicationWindowIconChange:
		if v, ok := i.(interface {
			OnApplicationWindowIconChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationWindowIconChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnApplicationWindowIconChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationWindowIconChangeEvent(e)
		}
	case QEvent_ApplicationFontChange:
		if v, ok := i.(interface {
			OnApplicationFontChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationFontChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnApplicationFontChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationFontChangeEvent(e)
		}
	case QEvent_ApplicationLayoutDirectionChange:
		if v, ok := i.(interface {
			OnApplicationLayoutDirectionChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationLayoutDirectionChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnApplicationLayoutDirectionChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationLayoutDirectionChangeEvent(e)
		}
	case QEvent_ApplicationPaletteChange:
		if v, ok := i.(interface {
			OnApplicationPaletteChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationPaletteChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnApplicationPaletteChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationPaletteChangeEvent(e)
		}
	case QEvent_PaletteChange:
		if v, ok := i.(interface {
			OnPaletteChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnPaletteChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnPaletteChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnPaletteChangeEvent(e)
		}
	case QEvent_Clipboard:
		if v, ok := i.(interface {
			OnClipboardEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnClipboardEvent(obj, e)
		} else if v, ok := i.(interface {
			OnClipboardEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnClipboardEvent(e)
		}
	case QEvent_Speech:
		if v, ok := i.(interface {
			OnSpeechEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnSpeechEvent(obj, e)
		} else if v, ok := i.(interface {
			OnSpeechEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnSpeechEvent(e)
		}
	case QEvent_MetaCall:
		if v, ok := i.(interface {
			OnMetaCallEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMetaCallEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMetaCallEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMetaCallEvent(e)
		}
	case QEvent_SockAct:
		if v, ok := i.(interface {
			OnSockActEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnSockActEvent(obj, e)
		} else if v, ok := i.(interface {
			OnSockActEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnSockActEvent(e)
		}
	case QEvent_WinEventAct:
		if v, ok := i.(interface {
			OnWinEventActEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWinEventActEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWinEventActEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWinEventActEvent(e)
		}
	case QEvent_DeferredDelete:
		if v, ok := i.(interface {
			OnDeferredDeleteEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDeferredDeleteEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDeferredDeleteEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDeferredDeleteEvent(e)
		}
	case QEvent_DragEnter:
		if v, ok := i.(interface {
			OnDragEnterEvent(*QObject, *QDragEnterEvent) bool
		}); ok {
			e := &QDragEnterEvent{}
			e.SetDriver(event, 24000, false)
			return v.OnDragEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDragEnterEvent(*QDragEnterEvent) bool
		}); ok {
			e := &QDragEnterEvent{}
			e.SetDriver(event, 24000, false)
			return v.OnDragEnterEvent(e)
		}
	case QEvent_DragMove:
		if v, ok := i.(interface {
			OnDragMoveEvent(*QObject, *QDragMoveEvent) bool
		}); ok {
			e := &QDragMoveEvent{}
			e.SetDriver(event, 26000, false)
			return v.OnDragMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDragMoveEvent(*QDragMoveEvent) bool
		}); ok {
			e := &QDragMoveEvent{}
			e.SetDriver(event, 26000, false)
			return v.OnDragMoveEvent(e)
		}
	case QEvent_DragLeave:
		if v, ok := i.(interface {
			OnDragLeaveEvent(*QObject, *QDragLeaveEvent) bool
		}); ok {
			e := &QDragLeaveEvent{}
			e.SetDriver(event, 25000, false)
			return v.OnDragLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDragLeaveEvent(*QDragLeaveEvent) bool
		}); ok {
			e := &QDragLeaveEvent{}
			e.SetDriver(event, 25000, false)
			return v.OnDragLeaveEvent(e)
		}
	case QEvent_Drop:
		if v, ok := i.(interface {
			OnDropEvent(*QObject, *QDropEvent) bool
		}); ok {
			e := &QDropEvent{}
			e.SetDriver(event, 27000, false)
			return v.OnDropEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDropEvent(*QDropEvent) bool
		}); ok {
			e := &QDropEvent{}
			e.SetDriver(event, 27000, false)
			return v.OnDropEvent(e)
		}
	case QEvent_DragResponse:
		if v, ok := i.(interface {
			OnDragResponseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDragResponseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDragResponseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDragResponseEvent(e)
		}
	case QEvent_ChildAdded:
		if v, ok := i.(interface {
			OnChildAddedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnChildAddedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnChildAddedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnChildAddedEvent(e)
		}
	case QEvent_ChildPolished:
		if v, ok := i.(interface {
			OnChildPolishedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnChildPolishedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnChildPolishedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnChildPolishedEvent(e)
		}
	case QEvent_ChildRemoved:
		if v, ok := i.(interface {
			OnChildRemovedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnChildRemovedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnChildRemovedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnChildRemovedEvent(e)
		}
	case QEvent_ShowWindowRequest:
		if v, ok := i.(interface {
			OnShowWindowRequestEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnShowWindowRequestEvent(obj, e)
		} else if v, ok := i.(interface {
			OnShowWindowRequestEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnShowWindowRequestEvent(e)
		}
	case QEvent_PolishRequest:
		if v, ok := i.(interface {
			OnPolishRequestEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnPolishRequestEvent(obj, e)
		} else if v, ok := i.(interface {
			OnPolishRequestEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnPolishRequestEvent(e)
		}
	case QEvent_Polish:
		if v, ok := i.(interface {
			OnPolishEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnPolishEvent(obj, e)
		} else if v, ok := i.(interface {
			OnPolishEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnPolishEvent(e)
		}
	case QEvent_LayoutRequest:
		if v, ok := i.(interface {
			OnLayoutRequestEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLayoutRequestEvent(obj, e)
		} else if v, ok := i.(interface {
			OnLayoutRequestEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLayoutRequestEvent(e)
		}
	case QEvent_UpdateRequest:
		if v, ok := i.(interface {
			OnUpdateRequestEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUpdateRequestEvent(obj, e)
		} else if v, ok := i.(interface {
			OnUpdateRequestEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUpdateRequestEvent(e)
		}
	case QEvent_UpdateLater:
		if v, ok := i.(interface {
			OnUpdateLaterEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUpdateLaterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnUpdateLaterEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUpdateLaterEvent(e)
		}
	case QEvent_EmbeddingControl:
		if v, ok := i.(interface {
			OnEmbeddingControlEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEmbeddingControlEvent(obj, e)
		} else if v, ok := i.(interface {
			OnEmbeddingControlEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEmbeddingControlEvent(e)
		}
	case QEvent_ActivateControl:
		if v, ok := i.(interface {
			OnActivateControlEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActivateControlEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActivateControlEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActivateControlEvent(e)
		}
	case QEvent_DeactivateControl:
		if v, ok := i.(interface {
			OnDeactivateControlEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDeactivateControlEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDeactivateControlEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDeactivateControlEvent(e)
		}
	case QEvent_ContextMenu:
		if v, ok := i.(interface {
			OnContextMenuEvent(*QObject, *QContextMenuEvent) bool
		}); ok {
			e := &QContextMenuEvent{}
			e.SetDriver(event, 16000, false)
			return v.OnContextMenuEvent(obj, e)
		} else if v, ok := i.(interface {
			OnContextMenuEvent(*QContextMenuEvent) bool
		}); ok {
			e := &QContextMenuEvent{}
			e.SetDriver(event, 16000, false)
			return v.OnContextMenuEvent(e)
		}
	case QEvent_InputMethod:
		if v, ok := i.(interface {
			OnInputMethodEvent(*QObject, *QInputMethodEvent) bool
		}); ok {
			e := &QInputMethodEvent{}
			e.SetDriver(event, 58000, false)
			return v.OnInputMethodEvent(obj, e)
		} else if v, ok := i.(interface {
			OnInputMethodEvent(*QInputMethodEvent) bool
		}); ok {
			e := &QInputMethodEvent{}
			e.SetDriver(event, 58000, false)
			return v.OnInputMethodEvent(e)
		}
	case QEvent_AccessibilityPrepare:
		if v, ok := i.(interface {
			OnAccessibilityPrepareEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAccessibilityPrepareEvent(obj, e)
		} else if v, ok := i.(interface {
			OnAccessibilityPrepareEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAccessibilityPrepareEvent(e)
		}
	case QEvent_TabletMove:
		if v, ok := i.(interface {
			OnTabletMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletMoveEvent(e)
		}
	case QEvent_LocaleChange:
		if v, ok := i.(interface {
			OnLocaleChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLocaleChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnLocaleChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLocaleChangeEvent(e)
		}
	case QEvent_LanguageChange:
		if v, ok := i.(interface {
			OnLanguageChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLanguageChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnLanguageChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLanguageChangeEvent(e)
		}
	case QEvent_LayoutDirectionChange:
		if v, ok := i.(interface {
			OnLayoutDirectionChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLayoutDirectionChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnLayoutDirectionChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLayoutDirectionChangeEvent(e)
		}
	case QEvent_Style:
		if v, ok := i.(interface {
			OnStyleEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStyleEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStyleEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStyleEvent(e)
		}
	case QEvent_TabletPress:
		if v, ok := i.(interface {
			OnTabletPressEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletPressEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletPressEvent(e)
		}
	case QEvent_TabletRelease:
		if v, ok := i.(interface {
			OnTabletReleaseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletReleaseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletReleaseEvent(e)
		}
	case QEvent_OkRequest:
		if v, ok := i.(interface {
			OnOkRequestEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnOkRequestEvent(obj, e)
		} else if v, ok := i.(interface {
			OnOkRequestEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnOkRequestEvent(e)
		}
	case QEvent_HelpRequest:
		if v, ok := i.(interface {
			OnHelpRequestEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHelpRequestEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHelpRequestEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHelpRequestEvent(e)
		}
	case QEvent_IconDrag:
		if v, ok := i.(interface {
			OnIconDragEvent(*QObject, *QIconDragEvent) bool
		}); ok {
			e := &QIconDragEvent{}
			e.SetDriver(event, 52000, false)
			return v.OnIconDragEvent(obj, e)
		} else if v, ok := i.(interface {
			OnIconDragEvent(*QIconDragEvent) bool
		}); ok {
			e := &QIconDragEvent{}
			e.SetDriver(event, 52000, false)
			return v.OnIconDragEvent(e)
		}
	case QEvent_FontChange:
		if v, ok := i.(interface {
			OnFontChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFontChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFontChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFontChangeEvent(e)
		}
	case QEvent_EnabledChange:
		if v, ok := i.(interface {
			OnEnabledChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEnabledChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnEnabledChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEnabledChangeEvent(e)
		}
	case QEvent_ActivationChange:
		if v, ok := i.(interface {
			OnActivationChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActivationChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActivationChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActivationChangeEvent(e)
		}
	case QEvent_StyleChange:
		if v, ok := i.(interface {
			OnStyleChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStyleChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStyleChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStyleChangeEvent(e)
		}
	case QEvent_IconTextChange:
		if v, ok := i.(interface {
			OnIconTextChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnIconTextChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnIconTextChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnIconTextChangeEvent(e)
		}
	case QEvent_ModifiedChange:
		if v, ok := i.(interface {
			OnModifiedChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnModifiedChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnModifiedChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnModifiedChangeEvent(e)
		}
	case QEvent_MouseTrackingChange:
		if v, ok := i.(interface {
			OnMouseTrackingChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseTrackingChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseTrackingChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMouseTrackingChangeEvent(e)
		}
	case QEvent_WindowBlocked:
		if v, ok := i.(interface {
			OnWindowBlockedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowBlockedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowBlockedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowBlockedEvent(e)
		}
	case QEvent_WindowUnblocked:
		if v, ok := i.(interface {
			OnWindowUnblockedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowUnblockedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowUnblockedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWindowUnblockedEvent(e)
		}
	case QEvent_WindowStateChange:
		if v, ok := i.(interface {
			OnWindowStateChangeEvent(*QObject, *QWindowStateChangeEvent) bool
		}); ok {
			e := &QWindowStateChangeEvent{}
			e.SetDriver(event, 193000, false)
			return v.OnWindowStateChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWindowStateChangeEvent(*QWindowStateChangeEvent) bool
		}); ok {
			e := &QWindowStateChangeEvent{}
			e.SetDriver(event, 193000, false)
			return v.OnWindowStateChangeEvent(e)
		}
	case QEvent_ToolTip:
		if v, ok := i.(interface {
			OnToolTipEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnToolTipEvent(obj, e)
		} else if v, ok := i.(interface {
			OnToolTipEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnToolTipEvent(e)
		}
	case QEvent_WhatsThis:
		if v, ok := i.(interface {
			OnWhatsThisEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWhatsThisEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWhatsThisEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWhatsThisEvent(e)
		}
	case QEvent_StatusTip:
		if v, ok := i.(interface {
			OnStatusTipEvent(*QObject, *QStatusTipEvent) bool
		}); ok {
			e := &QStatusTipEvent{}
			e.SetDriver(event, 128000, false)
			return v.OnStatusTipEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStatusTipEvent(*QStatusTipEvent) bool
		}); ok {
			e := &QStatusTipEvent{}
			e.SetDriver(event, 128000, false)
			return v.OnStatusTipEvent(e)
		}
	case QEvent_ActionChanged:
		if v, ok := i.(interface {
			OnActionChangedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActionChangedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActionChangedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActionChangedEvent(e)
		}
	case QEvent_ActionAdded:
		if v, ok := i.(interface {
			OnActionAddedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActionAddedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActionAddedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActionAddedEvent(e)
		}
	case QEvent_ActionRemoved:
		if v, ok := i.(interface {
			OnActionRemovedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActionRemovedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActionRemovedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnActionRemovedEvent(e)
		}
	case QEvent_FileOpen:
		if v, ok := i.(interface {
			OnFileOpenEvent(*QObject, *QFileOpenEvent) bool
		}); ok {
			e := &QFileOpenEvent{}
			e.SetDriver(event, 35000, false)
			return v.OnFileOpenEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFileOpenEvent(*QFileOpenEvent) bool
		}); ok {
			e := &QFileOpenEvent{}
			e.SetDriver(event, 35000, false)
			return v.OnFileOpenEvent(e)
		}
	case QEvent_Shortcut:
		if v, ok := i.(interface {
			OnShortcutEvent(*QObject, *QShortcutEvent) bool
		}); ok {
			e := &QShortcutEvent{}
			e.SetDriver(event, 118000, false)
			return v.OnShortcutEvent(obj, e)
		} else if v, ok := i.(interface {
			OnShortcutEvent(*QShortcutEvent) bool
		}); ok {
			e := &QShortcutEvent{}
			e.SetDriver(event, 118000, false)
			return v.OnShortcutEvent(e)
		}
	case QEvent_ShortcutOverride:
		if v, ok := i.(interface {
			OnShortcutOverrideEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnShortcutOverrideEvent(obj, e)
		} else if v, ok := i.(interface {
			OnShortcutOverrideEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnShortcutOverrideEvent(e)
		}
	case QEvent_WhatsThisClicked:
		if v, ok := i.(interface {
			OnWhatsThisClickedEvent(*QObject, *QWhatsThisClickedEvent) bool
		}); ok {
			e := &QWhatsThisClickedEvent{}
			e.SetDriver(event, 190000, false)
			return v.OnWhatsThisClickedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWhatsThisClickedEvent(*QWhatsThisClickedEvent) bool
		}); ok {
			e := &QWhatsThisClickedEvent{}
			e.SetDriver(event, 190000, false)
			return v.OnWhatsThisClickedEvent(e)
		}
	case QEvent_ToolBarChange:
		if v, ok := i.(interface {
			OnToolBarChangeEvent(*QObject, *QToolBarChangeEvent) bool
		}); ok {
			e := &QToolBarChangeEvent{}
			e.SetDriver(event, 174000, false)
			return v.OnToolBarChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnToolBarChangeEvent(*QToolBarChangeEvent) bool
		}); ok {
			e := &QToolBarChangeEvent{}
			e.SetDriver(event, 174000, false)
			return v.OnToolBarChangeEvent(e)
		}
	case QEvent_ApplicationActivate:
		if v, ok := i.(interface {
			OnApplicationActivateEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationActivateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnApplicationActivateEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationActivateEvent(e)
		}
	case QEvent_ApplicationDeactivate:
		if v, ok := i.(interface {
			OnApplicationDeactivateEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationDeactivateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnApplicationDeactivateEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnApplicationDeactivateEvent(e)
		}
	case QEvent_QueryWhatsThis:
		if v, ok := i.(interface {
			OnQueryWhatsThisEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnQueryWhatsThisEvent(obj, e)
		} else if v, ok := i.(interface {
			OnQueryWhatsThisEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnQueryWhatsThisEvent(e)
		}
	case QEvent_EnterWhatsThisMode:
		if v, ok := i.(interface {
			OnEnterWhatsThisModeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEnterWhatsThisModeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnEnterWhatsThisModeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnEnterWhatsThisModeEvent(e)
		}
	case QEvent_LeaveWhatsThisMode:
		if v, ok := i.(interface {
			OnLeaveWhatsThisModeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLeaveWhatsThisModeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnLeaveWhatsThisModeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnLeaveWhatsThisModeEvent(e)
		}
	case QEvent_ZOrderChange:
		if v, ok := i.(interface {
			OnZOrderChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnZOrderChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnZOrderChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnZOrderChangeEvent(e)
		}
	case QEvent_HoverEnter:
		if v, ok := i.(interface {
			OnHoverEnterEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHoverEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHoverEnterEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHoverEnterEvent(e)
		}
	case QEvent_HoverLeave:
		if v, ok := i.(interface {
			OnHoverLeaveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHoverLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHoverLeaveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHoverLeaveEvent(e)
		}
	case QEvent_HoverMove:
		if v, ok := i.(interface {
			OnHoverMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHoverMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHoverMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnHoverMoveEvent(e)
		}
	case QEvent_AccessibilityHelp:
		if v, ok := i.(interface {
			OnAccessibilityHelpEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAccessibilityHelpEvent(obj, e)
		} else if v, ok := i.(interface {
			OnAccessibilityHelpEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAccessibilityHelpEvent(e)
		}
	case QEvent_AccessibilityDescription:
		if v, ok := i.(interface {
			OnAccessibilityDescriptionEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAccessibilityDescriptionEvent(obj, e)
		} else if v, ok := i.(interface {
			OnAccessibilityDescriptionEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAccessibilityDescriptionEvent(e)
		}
	case QEvent_AcceptDropsChange:
		if v, ok := i.(interface {
			OnAcceptDropsChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAcceptDropsChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnAcceptDropsChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnAcceptDropsChangeEvent(e)
		}
	case QEvent_MenubarUpdated:
		if v, ok := i.(interface {
			OnMenubarUpdatedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMenubarUpdatedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMenubarUpdatedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMenubarUpdatedEvent(e)
		}
	case QEvent_ZeroTimerEvent:
		if v, ok := i.(interface {
			OnZeroTimerEventEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnZeroTimerEventEvent(obj, e)
		} else if v, ok := i.(interface {
			OnZeroTimerEventEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnZeroTimerEventEvent(e)
		}
	case QEvent_GraphicsSceneMouseMove:
		if v, ok := i.(interface {
			OnGraphicsSceneMouseMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMouseMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMouseMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMouseMoveEvent(e)
		}
	case QEvent_GraphicsSceneMousePress:
		if v, ok := i.(interface {
			OnGraphicsSceneMousePressEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMousePressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMousePressEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMousePressEvent(e)
		}
	case QEvent_GraphicsSceneMouseRelease:
		if v, ok := i.(interface {
			OnGraphicsSceneMouseReleaseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMouseReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMouseReleaseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMouseReleaseEvent(e)
		}
	case QEvent_GraphicsSceneMouseDoubleClick:
		if v, ok := i.(interface {
			OnGraphicsSceneMouseDoubleClickEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMouseDoubleClickEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMouseDoubleClickEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneMouseDoubleClickEvent(e)
		}
	case QEvent_GraphicsSceneContextMenu:
		if v, ok := i.(interface {
			OnGraphicsSceneContextMenuEvent(*QObject, *QGraphicsSceneContextMenuEvent) bool
		}); ok {
			e := &QGraphicsSceneContextMenuEvent{}
			e.SetDriver(event, 275000, false)
			return v.OnGraphicsSceneContextMenuEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneContextMenuEvent(*QGraphicsSceneContextMenuEvent) bool
		}); ok {
			e := &QGraphicsSceneContextMenuEvent{}
			e.SetDriver(event, 275000, false)
			return v.OnGraphicsSceneContextMenuEvent(e)
		}
	case QEvent_GraphicsSceneHoverEnter:
		if v, ok := i.(interface {
			OnGraphicsSceneHoverEnterEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneHoverEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHoverEnterEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneHoverEnterEvent(e)
		}
	case QEvent_GraphicsSceneHoverMove:
		if v, ok := i.(interface {
			OnGraphicsSceneHoverMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneHoverMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHoverMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneHoverMoveEvent(e)
		}
	case QEvent_GraphicsSceneHoverLeave:
		if v, ok := i.(interface {
			OnGraphicsSceneHoverLeaveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneHoverLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHoverLeaveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneHoverLeaveEvent(e)
		}
	case QEvent_GraphicsSceneHelp:
		if v, ok := i.(interface {
			OnGraphicsSceneHelpEvent(*QObject, *QGraphicsSceneHelpEvent) bool
		}); ok {
			e := &QGraphicsSceneHelpEvent{}
			e.SetDriver(event, 278000, false)
			return v.OnGraphicsSceneHelpEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHelpEvent(*QGraphicsSceneHelpEvent) bool
		}); ok {
			e := &QGraphicsSceneHelpEvent{}
			e.SetDriver(event, 278000, false)
			return v.OnGraphicsSceneHelpEvent(e)
		}
	case QEvent_GraphicsSceneDragEnter:
		if v, ok := i.(interface {
			OnGraphicsSceneDragEnterEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDragEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDragEnterEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDragEnterEvent(e)
		}
	case QEvent_GraphicsSceneDragMove:
		if v, ok := i.(interface {
			OnGraphicsSceneDragMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDragMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDragMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDragMoveEvent(e)
		}
	case QEvent_GraphicsSceneDragLeave:
		if v, ok := i.(interface {
			OnGraphicsSceneDragLeaveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDragLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDragLeaveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDragLeaveEvent(e)
		}
	case QEvent_GraphicsSceneDrop:
		if v, ok := i.(interface {
			OnGraphicsSceneDropEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDropEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDropEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGraphicsSceneDropEvent(e)
		}
	case QEvent_GraphicsSceneWheel:
		if v, ok := i.(interface {
			OnGraphicsSceneWheelEvent(*QObject, *QGraphicsSceneWheelEvent) bool
		}); ok {
			e := &QGraphicsSceneWheelEvent{}
			e.SetDriver(event, 283000, false)
			return v.OnGraphicsSceneWheelEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneWheelEvent(*QGraphicsSceneWheelEvent) bool
		}); ok {
			e := &QGraphicsSceneWheelEvent{}
			e.SetDriver(event, 283000, false)
			return v.OnGraphicsSceneWheelEvent(e)
		}
	case QEvent_KeyboardLayoutChange:
		if v, ok := i.(interface {
			OnKeyboardLayoutChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnKeyboardLayoutChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnKeyboardLayoutChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnKeyboardLayoutChangeEvent(e)
		}
	case QEvent_DynamicPropertyChange:
		if v, ok := i.(interface {
			OnDynamicPropertyChangeEvent(*QObject, *QDynamicPropertyChangeEvent) bool
		}); ok {
			e := &QDynamicPropertyChangeEvent{}
			e.SetDriver(event, 28000, false)
			return v.OnDynamicPropertyChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDynamicPropertyChangeEvent(*QDynamicPropertyChangeEvent) bool
		}); ok {
			e := &QDynamicPropertyChangeEvent{}
			e.SetDriver(event, 28000, false)
			return v.OnDynamicPropertyChangeEvent(e)
		}
	case QEvent_TabletEnterProximity:
		if v, ok := i.(interface {
			OnTabletEnterProximityEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletEnterProximityEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletEnterProximityEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletEnterProximityEvent(e)
		}
	case QEvent_TabletLeaveProximity:
		if v, ok := i.(interface {
			OnTabletLeaveProximityEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletLeaveProximityEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletLeaveProximityEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTabletLeaveProximityEvent(e)
		}
	case QEvent_NonClientAreaMouseMove:
		if v, ok := i.(interface {
			OnNonClientAreaMouseMoveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnNonClientAreaMouseMoveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseMoveEvent(e)
		}
	case QEvent_NonClientAreaMouseButtonPress:
		if v, ok := i.(interface {
			OnNonClientAreaMouseButtonPressEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseButtonPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnNonClientAreaMouseButtonPressEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseButtonPressEvent(e)
		}
	case QEvent_NonClientAreaMouseButtonRelease:
		if v, ok := i.(interface {
			OnNonClientAreaMouseButtonReleaseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseButtonReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnNonClientAreaMouseButtonReleaseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseButtonReleaseEvent(e)
		}
	case QEvent_NonClientAreaMouseButtonDblClick:
		if v, ok := i.(interface {
			OnNonClientAreaMouseButtonDblClickEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseButtonDblClickEvent(obj, e)
		} else if v, ok := i.(interface {
			OnNonClientAreaMouseButtonDblClickEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNonClientAreaMouseButtonDblClickEvent(e)
		}
	case QEvent_MacSizeChange:
		if v, ok := i.(interface {
			OnMacSizeChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMacSizeChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMacSizeChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMacSizeChangeEvent(e)
		}
	case QEvent_ContentsRectChange:
		if v, ok := i.(interface {
			OnContentsRectChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnContentsRectChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnContentsRectChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnContentsRectChangeEvent(e)
		}
	case QEvent_MacGLWindowChange:
		if v, ok := i.(interface {
			OnMacGLWindowChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMacGLWindowChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMacGLWindowChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMacGLWindowChangeEvent(e)
		}
	case QEvent_FutureCallOut:
		if v, ok := i.(interface {
			OnFutureCallOutEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFutureCallOutEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFutureCallOutEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnFutureCallOutEvent(e)
		}
	case QEvent_GraphicsSceneResize:
		if v, ok := i.(interface {
			OnGraphicsSceneResizeEvent(*QObject, *QGraphicsSceneResizeEvent) bool
		}); ok {
			e := &QGraphicsSceneResizeEvent{}
			e.SetDriver(event, 282000, false)
			return v.OnGraphicsSceneResizeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneResizeEvent(*QGraphicsSceneResizeEvent) bool
		}); ok {
			e := &QGraphicsSceneResizeEvent{}
			e.SetDriver(event, 282000, false)
			return v.OnGraphicsSceneResizeEvent(e)
		}
	case QEvent_GraphicsSceneMove:
		if v, ok := i.(interface {
			OnGraphicsSceneMoveEvent(*QObject, *QGraphicsSceneMoveEvent) bool
		}); ok {
			e := &QGraphicsSceneMoveEvent{}
			e.SetDriver(event, 281000, false)
			return v.OnGraphicsSceneMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMoveEvent(*QGraphicsSceneMoveEvent) bool
		}); ok {
			e := &QGraphicsSceneMoveEvent{}
			e.SetDriver(event, 281000, false)
			return v.OnGraphicsSceneMoveEvent(e)
		}
	case QEvent_CursorChange:
		if v, ok := i.(interface {
			OnCursorChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnCursorChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnCursorChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnCursorChangeEvent(e)
		}
	case QEvent_ToolTipChange:
		if v, ok := i.(interface {
			OnToolTipChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnToolTipChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnToolTipChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnToolTipChangeEvent(e)
		}
	case QEvent_NetworkReplyUpdated:
		if v, ok := i.(interface {
			OnNetworkReplyUpdatedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNetworkReplyUpdatedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnNetworkReplyUpdatedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNetworkReplyUpdatedEvent(e)
		}
	case QEvent_GrabMouse:
		if v, ok := i.(interface {
			OnGrabMouseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGrabMouseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGrabMouseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGrabMouseEvent(e)
		}
	case QEvent_UngrabMouse:
		if v, ok := i.(interface {
			OnUngrabMouseEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUngrabMouseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnUngrabMouseEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUngrabMouseEvent(e)
		}
	case QEvent_GrabKeyboard:
		if v, ok := i.(interface {
			OnGrabKeyboardEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGrabKeyboardEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGrabKeyboardEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGrabKeyboardEvent(e)
		}
	case QEvent_UngrabKeyboard:
		if v, ok := i.(interface {
			OnUngrabKeyboardEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUngrabKeyboardEvent(obj, e)
		} else if v, ok := i.(interface {
			OnUngrabKeyboardEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUngrabKeyboardEvent(e)
		}
	case QEvent_MacGLClearDrawable:
		if v, ok := i.(interface {
			OnMacGLClearDrawableEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMacGLClearDrawableEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMacGLClearDrawableEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnMacGLClearDrawableEvent(e)
		}
	case QEvent_StateMachineSignal:
		if v, ok := i.(interface {
			OnStateMachineSignalEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStateMachineSignalEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStateMachineSignalEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStateMachineSignalEvent(e)
		}
	case QEvent_StateMachineWrapped:
		if v, ok := i.(interface {
			OnStateMachineWrappedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStateMachineWrappedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStateMachineWrappedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnStateMachineWrappedEvent(e)
		}
	case QEvent_TouchBegin:
		if v, ok := i.(interface {
			OnTouchBeginEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTouchBeginEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTouchBeginEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTouchBeginEvent(e)
		}
	case QEvent_TouchUpdate:
		if v, ok := i.(interface {
			OnTouchUpdateEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTouchUpdateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTouchUpdateEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTouchUpdateEvent(e)
		}
	case QEvent_TouchEnd:
		if v, ok := i.(interface {
			OnTouchEndEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTouchEndEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTouchEndEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnTouchEndEvent(e)
		}
	case QEvent_NativeGesture:
		if v, ok := i.(interface {
			OnNativeGestureEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNativeGestureEvent(obj, e)
		} else if v, ok := i.(interface {
			OnNativeGestureEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnNativeGestureEvent(e)
		}
	case QEvent_RequestSoftwareInputPanel:
		if v, ok := i.(interface {
			OnRequestSoftwareInputPanelEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnRequestSoftwareInputPanelEvent(obj, e)
		} else if v, ok := i.(interface {
			OnRequestSoftwareInputPanelEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnRequestSoftwareInputPanelEvent(e)
		}
	case QEvent_CloseSoftwareInputPanel:
		if v, ok := i.(interface {
			OnCloseSoftwareInputPanelEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnCloseSoftwareInputPanelEvent(obj, e)
		} else if v, ok := i.(interface {
			OnCloseSoftwareInputPanelEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnCloseSoftwareInputPanelEvent(e)
		}
	case QEvent_UpdateSoftKeys:
		if v, ok := i.(interface {
			OnUpdateSoftKeysEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUpdateSoftKeysEvent(obj, e)
		} else if v, ok := i.(interface {
			OnUpdateSoftKeysEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUpdateSoftKeysEvent(e)
		}
	case QEvent_WinIdChange:
		if v, ok := i.(interface {
			OnWinIdChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWinIdChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWinIdChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWinIdChangeEvent(e)
		}
	case QEvent_Gesture:
		if v, ok := i.(interface {
			OnGestureEvent(*QObject, *QGestureEvent) bool
		}); ok {
			e := &QGestureEvent{}
			e.SetDriver(event, 44000, false)
			return v.OnGestureEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGestureEvent(*QGestureEvent) bool
		}); ok {
			e := &QGestureEvent{}
			e.SetDriver(event, 44000, false)
			return v.OnGestureEvent(e)
		}
	case QEvent_GestureOverride:
		if v, ok := i.(interface {
			OnGestureOverrideEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGestureOverrideEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGestureOverrideEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnGestureOverrideEvent(e)
		}
	default:
		if v, ok := i.(interface {
			OnUserEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUserEvent(obj, e)
		} else if v, ok := i.(interface {
			OnUserEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnUserEvent(e)
		}
	}
	return false
}
