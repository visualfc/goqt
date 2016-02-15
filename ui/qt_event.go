// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

//Installs an event filter filterInterface on this object.
//
//OnEvent(event *QEvent) bool
//
//OnEvent(watched *QObject, event *QEvent) bool
//
//OnAcceptDropsChangeEvent(event *QEvent) bool
//
//OnAcceptDropsChangeEvent(watched *QObject, event *QEvent) bool
//
//OnAccessibilityDescriptionEvent(event *QEvent) bool
//
//OnAccessibilityDescriptionEvent(watched *QObject, event *QEvent) bool
//
//OnAccessibilityHelpEvent(event *QEvent) bool
//
//OnAccessibilityHelpEvent(watched *QObject, event *QEvent) bool
//
//OnAccessibilityPrepareEvent(event *QEvent) bool
//
//OnAccessibilityPrepareEvent(watched *QObject, event *QEvent) bool
//
//OnActionAddedEvent(event *QActionEvent) bool
//
//OnActionAddedEvent(watched *QObject, event *QActionEvent) bool
//
//OnActionChangedEvent(event *QActionEvent) bool
//
//OnActionChangedEvent(watched *QObject, event *QActionEvent) bool
//
//OnActionRemovedEvent(event *QActionEvent) bool
//
//OnActionRemovedEvent(watched *QObject, event *QActionEvent) bool
//
//OnActivateControlEvent(event *QEvent) bool
//
//OnActivateControlEvent(watched *QObject, event *QEvent) bool
//
//OnActivationChangeEvent(event *QEvent) bool
//
//OnActivationChangeEvent(watched *QObject, event *QEvent) bool
//
//OnApplicationActivateEvent(event *QEvent) bool
//
//OnApplicationActivateEvent(watched *QObject, event *QEvent) bool
//
//OnApplicationDeactivateEvent(event *QEvent) bool
//
//OnApplicationDeactivateEvent(watched *QObject, event *QEvent) bool
//
//OnApplicationFontChangeEvent(event *QEvent) bool
//
//OnApplicationFontChangeEvent(watched *QObject, event *QEvent) bool
//
//OnApplicationLayoutDirectionChangeEvent(event *QEvent) bool
//
//OnApplicationLayoutDirectionChangeEvent(watched *QObject, event *QEvent) bool
//
//OnApplicationPaletteChangeEvent(event *QEvent) bool
//
//OnApplicationPaletteChangeEvent(watched *QObject, event *QEvent) bool
//
//OnApplicationWindowIconChangeEvent(event *QEvent) bool
//
//OnApplicationWindowIconChangeEvent(watched *QObject, event *QEvent) bool
//
//OnChildAddedEvent(event *QChildEvent) bool
//
//OnChildAddedEvent(watched *QObject, event *QChildEvent) bool
//
//OnChildPolishedEvent(event *QChildEvent) bool
//
//OnChildPolishedEvent(watched *QObject, event *QChildEvent) bool
//
//OnChildRemovedEvent(event *QChildEvent) bool
//
//OnChildRemovedEvent(watched *QObject, event *QChildEvent) bool
//
//OnClipboardEvent(event *QEvent) bool
//
//OnClipboardEvent(watched *QObject, event *QEvent) bool
//
//OnCloseEvent(event *QCloseEvent) bool
//
//OnCloseEvent(watched *QObject, event *QCloseEvent) bool
//
//OnCloseSoftwareInputPanelEvent(event *QEvent) bool
//
//OnCloseSoftwareInputPanelEvent(watched *QObject, event *QEvent) bool
//
//OnContentsRectChangeEvent(event *QEvent) bool
//
//OnContentsRectChangeEvent(watched *QObject, event *QEvent) bool
//
//OnContextMenuEvent(event *QContextMenuEvent) bool
//
//OnContextMenuEvent(watched *QObject, event *QContextMenuEvent) bool
//
//OnCreateEvent(event *QEvent) bool
//
//OnCreateEvent(watched *QObject, event *QEvent) bool
//
//OnCursorChangeEvent(event *QEvent) bool
//
//OnCursorChangeEvent(watched *QObject, event *QEvent) bool
//
//OnDeactivateControlEvent(event *QEvent) bool
//
//OnDeactivateControlEvent(watched *QObject, event *QEvent) bool
//
//OnDeferredDeleteEvent(event *QEvent) bool
//
//OnDeferredDeleteEvent(watched *QObject, event *QEvent) bool
//
//OnDestroyEvent(event *QEvent) bool
//
//OnDestroyEvent(watched *QObject, event *QEvent) bool
//
//OnDragEnterEvent(event *QDragEnterEvent) bool
//
//OnDragEnterEvent(watched *QObject, event *QDragEnterEvent) bool
//
//OnDragLeaveEvent(event *QEvent) bool
//
//OnDragLeaveEvent(watched *QObject, event *QEvent) bool
//
//OnDragMoveEvent(event *QDragMoveEvent) bool
//
//OnDragMoveEvent(watched *QObject, event *QDragMoveEvent) bool
//
//OnDragResponseEvent(event *QEvent) bool
//
//OnDragResponseEvent(watched *QObject, event *QEvent) bool
//
//OnDropEvent(event *QEvent) bool
//
//OnDropEvent(watched *QObject, event *QEvent) bool
//
//OnDynamicPropertyChangeEvent(event *QEvent) bool
//
//OnDynamicPropertyChangeEvent(watched *QObject, event *QEvent) bool
//
//OnEmbeddingControlEvent(event *QEvent) bool
//
//OnEmbeddingControlEvent(watched *QObject, event *QEvent) bool
//
//OnEnabledChangeEvent(event *QEvent) bool
//
//OnEnabledChangeEvent(watched *QObject, event *QEvent) bool
//
//OnEnterEvent(event *QEvent) bool
//
//OnEnterEvent(watched *QObject, event *QEvent) bool
//
//OnEnterWhatsThisModeEvent(event *QEvent) bool
//
//OnEnterWhatsThisModeEvent(watched *QObject, event *QEvent) bool
//
//OnFileOpenEvent(event *QFileOpenEvent) bool
//
//OnFileOpenEvent(watched *QObject, event *QFileOpenEvent) bool
//
//OnFocusInEvent(event *QFocusEvent) bool
//
//OnFocusInEvent(watched *QObject, event *QFocusEvent) bool
//
//OnFocusOutEvent(event *QFocusEvent) bool
//
//OnFocusOutEvent(watched *QObject, event *QFocusEvent) bool
//
//OnFontChangeEvent(event *QEvent) bool
//
//OnFontChangeEvent(watched *QObject, event *QEvent) bool
//
//OnFutureCallOutEvent(event *QEvent) bool
//
//OnFutureCallOutEvent(watched *QObject, event *QEvent) bool
//
//OnGestureEvent(event *QGestureEvent) bool
//
//OnGestureEvent(watched *QObject, event *QGestureEvent) bool
//
//OnGestureOverrideEvent(event *QGestureEvent) bool
//
//OnGestureOverrideEvent(watched *QObject, event *QGestureEvent) bool
//
//OnGrabKeyboardEvent(event *QEvent) bool
//
//OnGrabKeyboardEvent(watched *QObject, event *QEvent) bool
//
//OnGrabMouseEvent(event *QEvent) bool
//
//OnGrabMouseEvent(watched *QObject, event *QEvent) bool
//
//OnGraphicsSceneContextMenuEvent(event *QGraphicsSceneContextMenuEvent) bool
//
//OnGraphicsSceneContextMenuEvent(watched *QObject, event *QGraphicsSceneContextMenuEvent) bool
//
//OnGraphicsSceneDragEnterEvent(event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDragEnterEvent(watched *QObject, event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDragLeaveEvent(event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDragLeaveEvent(watched *QObject, event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDragMoveEvent(event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDragMoveEvent(watched *QObject, event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDropEvent(event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneDropEvent(watched *QObject, event *QGraphicsSceneDragDropEvent) bool
//
//OnGraphicsSceneHelpEvent(event *QHelpEvent) bool
//
//OnGraphicsSceneHelpEvent(watched *QObject, event *QHelpEvent) bool
//
//OnGraphicsSceneHoverEnterEvent(event *QGraphicsSceneHoverEvent) bool
//
//OnGraphicsSceneHoverEnterEvent(watched *QObject, event *QGraphicsSceneHoverEvent) bool
//
//OnGraphicsSceneHoverLeaveEvent(event *QGraphicsSceneHoverEvent) bool
//
//OnGraphicsSceneHoverLeaveEvent(watched *QObject, event *QGraphicsSceneHoverEvent) bool
//
//OnGraphicsSceneHoverMoveEvent(event *QGraphicsSceneHoverEvent) bool
//
//OnGraphicsSceneHoverMoveEvent(watched *QObject, event *QGraphicsSceneHoverEvent) bool
//
//OnGraphicsSceneMouseDoubleClickEvent(event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMouseDoubleClickEvent(watched *QObject, event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMouseMoveEvent(event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMouseMoveEvent(watched *QObject, event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMousePressEvent(event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMousePressEvent(watched *QObject, event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMouseReleaseEvent(event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMouseReleaseEvent(watched *QObject, event *QGraphicsSceneMouseEvent) bool
//
//OnGraphicsSceneMoveEvent(event *QGraphicsSceneMoveEvent) bool
//
//OnGraphicsSceneMoveEvent(watched *QObject, event *QGraphicsSceneMoveEvent) bool
//
//OnGraphicsSceneResizeEvent(event *QGraphicsSceneResizeEvent) bool
//
//OnGraphicsSceneResizeEvent(watched *QObject, event *QGraphicsSceneResizeEvent) bool
//
//OnGraphicsSceneWheelEvent(event *QGraphicsSceneWheelEvent) bool
//
//OnGraphicsSceneWheelEvent(watched *QObject, event *QGraphicsSceneWheelEvent) bool
//
//OnHelpRequestEvent(event *QEvent) bool
//
//OnHelpRequestEvent(watched *QObject, event *QEvent) bool
//
//OnHideEvent(event *QHideEvent) bool
//
//OnHideEvent(watched *QObject, event *QHideEvent) bool
//
//OnHideToParentEvent(event *QEvent) bool
//
//OnHideToParentEvent(watched *QObject, event *QEvent) bool
//
//OnHoverEnterEvent(event *QHoverEvent) bool
//
//OnHoverEnterEvent(watched *QObject, event *QHoverEvent) bool
//
//OnHoverLeaveEvent(event *QHoverEvent) bool
//
//OnHoverLeaveEvent(watched *QObject, event *QHoverEvent) bool
//
//OnHoverMoveEvent(event *QHoverEvent) bool
//
//OnHoverMoveEvent(watched *QObject, event *QHoverEvent) bool
//
//OnIconDragEvent(event *QIconDragEvent) bool
//
//OnIconDragEvent(watched *QObject, event *QIconDragEvent) bool
//
//OnIconTextChangeEvent(event *QEvent) bool
//
//OnIconTextChangeEvent(watched *QObject, event *QEvent) bool
//
//OnInputMethodEvent(event *QInputMethodEvent) bool
//
//OnInputMethodEvent(watched *QObject, event *QInputMethodEvent) bool
//
//OnKeyPressEvent(event *QKeyEvent) bool
//
//OnKeyPressEvent(watched *QObject, event *QKeyEvent) bool
//
//OnKeyReleaseEvent(event *QKeyEvent) bool
//
//OnKeyReleaseEvent(watched *QObject, event *QKeyEvent) bool
//
//OnKeyboardLayoutChangeEvent(event *QEvent) bool
//
//OnKeyboardLayoutChangeEvent(watched *QObject, event *QEvent) bool
//
//OnLanguageChangeEvent(event *QEvent) bool
//
//OnLanguageChangeEvent(watched *QObject, event *QEvent) bool
//
//OnLayoutDirectionChangeEvent(event *QEvent) bool
//
//OnLayoutDirectionChangeEvent(watched *QObject, event *QEvent) bool
//
//OnLayoutRequestEvent(event *QEvent) bool
//
//OnLayoutRequestEvent(watched *QObject, event *QEvent) bool
//
//OnLeaveEvent(event *QEvent) bool
//
//OnLeaveEvent(watched *QObject, event *QEvent) bool
//
//OnLeaveWhatsThisModeEvent(event *QEvent) bool
//
//OnLeaveWhatsThisModeEvent(watched *QObject, event *QEvent) bool
//
//OnLocaleChangeEvent(event *QEvent) bool
//
//OnLocaleChangeEvent(watched *QObject, event *QEvent) bool
//
//OnMacGLClearDrawableEvent(event *QEvent) bool
//
//OnMacGLClearDrawableEvent(watched *QObject, event *QEvent) bool
//
//OnMacGLWindowChangeEvent(event *QEvent) bool
//
//OnMacGLWindowChangeEvent(watched *QObject, event *QEvent) bool
//
//OnMacSizeChangeEvent(event *QEvent) bool
//
//OnMacSizeChangeEvent(watched *QObject, event *QEvent) bool
//
//OnMenubarUpdatedEvent(event *QEvent) bool
//
//OnMenubarUpdatedEvent(watched *QObject, event *QEvent) bool
//
//OnMetaCallEvent(event *QEvent) bool
//
//OnMetaCallEvent(watched *QObject, event *QEvent) bool
//
//OnModifiedChangeEvent(event *QEvent) bool
//
//OnModifiedChangeEvent(watched *QObject, event *QEvent) bool
//
//OnMouseButtonDblClickEvent(event *QMouseEvent) bool
//
//OnMouseButtonDblClickEvent(watched *QObject, event *QMouseEvent) bool
//
//OnMouseButtonPressEvent(event *QMouseEvent) bool
//
//OnMouseButtonPressEvent(watched *QObject, event *QMouseEvent) bool
//
//OnMouseButtonReleaseEvent(event *QMouseEvent) bool
//
//OnMouseButtonReleaseEvent(watched *QObject, event *QMouseEvent) bool
//
//OnMouseMoveEvent(event *QMouseEvent) bool
//
//OnMouseMoveEvent(watched *QObject, event *QMouseEvent) bool
//
//OnMouseTrackingChangeEvent(event *QEvent) bool
//
//OnMouseTrackingChangeEvent(watched *QObject, event *QEvent) bool
//
//OnMoveEvent(event *QMoveEvent) bool
//
//OnMoveEvent(watched *QObject, event *QMoveEvent) bool
//
//OnNativeGestureEvent(event *QEvent) bool
//
//OnNativeGestureEvent(watched *QObject, event *QEvent) bool
//
//OnNetworkReplyUpdatedEvent(event *QEvent) bool
//
//OnNetworkReplyUpdatedEvent(watched *QObject, event *QEvent) bool
//
//OnNonClientAreaMouseButtonDblClickEvent(event *QEvent) bool
//
//OnNonClientAreaMouseButtonDblClickEvent(watched *QObject, event *QEvent) bool
//
//OnNonClientAreaMouseButtonPressEvent(event *QEvent) bool
//
//OnNonClientAreaMouseButtonPressEvent(watched *QObject, event *QEvent) bool
//
//OnNonClientAreaMouseButtonReleaseEvent(event *QEvent) bool
//
//OnNonClientAreaMouseButtonReleaseEvent(watched *QObject, event *QEvent) bool
//
//OnNonClientAreaMouseMoveEvent(event *QEvent) bool
//
//OnNonClientAreaMouseMoveEvent(watched *QObject, event *QEvent) bool
//
//OnOkRequestEvent(event *QEvent) bool
//
//OnOkRequestEvent(watched *QObject, event *QEvent) bool
//
//OnPaintEvent(event *QPaintEvent) bool
//
//OnPaintEvent(watched *QObject, event *QPaintEvent) bool
//
//OnPaletteChangeEvent(event *QEvent) bool
//
//OnPaletteChangeEvent(watched *QObject, event *QEvent) bool
//
//OnParentAboutToChangeEvent(event *QEvent) bool
//
//OnParentAboutToChangeEvent(watched *QObject, event *QEvent) bool
//
//OnParentChangeEvent(event *QEvent) bool
//
//OnParentChangeEvent(watched *QObject, event *QEvent) bool
//
//OnPolishEvent(event *QEvent) bool
//
//OnPolishEvent(watched *QObject, event *QEvent) bool
//
//OnPolishRequestEvent(event *QEvent) bool
//
//OnPolishRequestEvent(watched *QObject, event *QEvent) bool
//
//OnQueryWhatsThisEvent(event *QEvent) bool
//
//OnQueryWhatsThisEvent(watched *QObject, event *QEvent) bool
//
//OnQuitEvent(event *QEvent) bool
//
//OnQuitEvent(watched *QObject, event *QEvent) bool
//
//OnRequestSoftwareInputPanelEvent(event *QEvent) bool
//
//OnRequestSoftwareInputPanelEvent(watched *QObject, event *QEvent) bool
//
//OnResizeEvent(event *QResizeEvent) bool
//
//OnResizeEvent(watched *QObject, event *QResizeEvent) bool
//
//OnShortcutEvent(event *QShortcutEvent) bool
//
//OnShortcutEvent(watched *QObject, event *QShortcutEvent) bool
//
//OnShortcutOverrideEvent(event *QKeyEvent) bool
//
//OnShortcutOverrideEvent(watched *QObject, event *QKeyEvent) bool
//
//OnShowEvent(event *QShowEvent) bool
//
//OnShowEvent(watched *QObject, event *QShowEvent) bool
//
//OnShowToParentEvent(event *QEvent) bool
//
//OnShowToParentEvent(watched *QObject, event *QEvent) bool
//
//OnShowWindowRequestEvent(event *QEvent) bool
//
//OnShowWindowRequestEvent(watched *QObject, event *QEvent) bool
//
//OnSockActEvent(event *QEvent) bool
//
//OnSockActEvent(watched *QObject, event *QEvent) bool
//
//OnSpeechEvent(event *QEvent) bool
//
//OnSpeechEvent(watched *QObject, event *QEvent) bool
//
//OnStateMachineSignalEvent(event *QStateMachineSignalEvent) bool
//
//OnStateMachineSignalEvent(watched *QObject, event *QStateMachineSignalEvent) bool
//
//OnStateMachineWrappedEvent(event *QStateMachineWrappedEvent) bool
//
//OnStateMachineWrappedEvent(watched *QObject, event *QStateMachineWrappedEvent) bool
//
//OnStatusTipEvent(event *QStatusTipEvent) bool
//
//OnStatusTipEvent(watched *QObject, event *QStatusTipEvent) bool
//
//OnStyleEvent(event *QEvent) bool
//
//OnStyleEvent(watched *QObject, event *QEvent) bool
//
//OnStyleChangeEvent(event *QEvent) bool
//
//OnStyleChangeEvent(watched *QObject, event *QEvent) bool
//
//OnTabletEnterProximityEvent(event *QTabletEvent) bool
//
//OnTabletEnterProximityEvent(watched *QObject, event *QTabletEvent) bool
//
//OnTabletLeaveProximityEvent(event *QTabletEvent) bool
//
//OnTabletLeaveProximityEvent(watched *QObject, event *QTabletEvent) bool
//
//OnTabletMoveEvent(event *QTabletEvent) bool
//
//OnTabletMoveEvent(watched *QObject, event *QTabletEvent) bool
//
//OnTabletPressEvent(event *QTabletEvent) bool
//
//OnTabletPressEvent(watched *QObject, event *QTabletEvent) bool
//
//OnTabletReleaseEvent(event *QTabletEvent) bool
//
//OnTabletReleaseEvent(watched *QObject, event *QTabletEvent) bool
//
//OnThreadChangeEvent(event *QEvent) bool
//
//OnThreadChangeEvent(watched *QObject, event *QEvent) bool
//
//OnTimerEvent(event *QTimerEvent) bool
//
//OnTimerEvent(watched *QObject, event *QTimerEvent) bool
//
//OnToolBarChangeEvent(event *QEvent) bool
//
//OnToolBarChangeEvent(watched *QObject, event *QEvent) bool
//
//OnToolTipEvent(event *QHelpEvent) bool
//
//OnToolTipEvent(watched *QObject, event *QHelpEvent) bool
//
//OnToolTipChangeEvent(event *QEvent) bool
//
//OnToolTipChangeEvent(watched *QObject, event *QEvent) bool
//
//OnTouchBeginEvent(event *QTouchEvent) bool
//
//OnTouchBeginEvent(watched *QObject, event *QTouchEvent) bool
//
//OnTouchEndEvent(event *QTouchEvent) bool
//
//OnTouchEndEvent(watched *QObject, event *QTouchEvent) bool
//
//OnTouchUpdateEvent(event *QTouchEvent) bool
//
//OnTouchUpdateEvent(watched *QObject, event *QTouchEvent) bool
//
//OnUngrabKeyboardEvent(event *QEvent) bool
//
//OnUngrabKeyboardEvent(watched *QObject, event *QEvent) bool
//
//OnUngrabMouseEvent(event *QEvent) bool
//
//OnUngrabMouseEvent(watched *QObject, event *QEvent) bool
//
//OnUpdateLaterEvent(event *QEvent) bool
//
//OnUpdateLaterEvent(watched *QObject, event *QEvent) bool
//
//OnUpdateRequestEvent(event *QEvent) bool
//
//OnUpdateRequestEvent(watched *QObject, event *QEvent) bool
//
//OnUpdateSoftKeysEvent(event *QEvent) bool
//
//OnUpdateSoftKeysEvent(watched *QObject, event *QEvent) bool
//
//OnWhatsThisEvent(event *QEvent) bool
//
//OnWhatsThisEvent(watched *QObject, event *QEvent) bool
//
//OnWhatsThisClickedEvent(event *QEvent) bool
//
//OnWhatsThisClickedEvent(watched *QObject, event *QEvent) bool
//
//OnWheelEvent(event *QWheelEvent) bool
//
//OnWheelEvent(watched *QObject, event *QWheelEvent) bool
//
//OnWinEventActEvent(event *QEvent) bool
//
//OnWinEventActEvent(watched *QObject, event *QEvent) bool
//
//OnWinIdChangeEvent(event *QEvent) bool
//
//OnWinIdChangeEvent(watched *QObject, event *QEvent) bool
//
//OnWindowActivateEvent(event *QEvent) bool
//
//OnWindowActivateEvent(watched *QObject, event *QEvent) bool
//
//OnWindowBlockedEvent(event *QEvent) bool
//
//OnWindowBlockedEvent(watched *QObject, event *QEvent) bool
//
//OnWindowDeactivateEvent(event *QEvent) bool
//
//OnWindowDeactivateEvent(watched *QObject, event *QEvent) bool
//
//OnWindowIconChangeEvent(event *QEvent) bool
//
//OnWindowIconChangeEvent(watched *QObject, event *QEvent) bool
//
//OnWindowStateChangeEvent(event *QWindowStateChangeEvent) bool
//
//OnWindowStateChangeEvent(watched *QObject, event *QWindowStateChangeEvent) bool
//
//OnWindowTitleChangeEvent(event *QEvent) bool
//
//OnWindowTitleChangeEvent(watched *QObject, event *QEvent) bool
//
//OnWindowUnblockedEvent(event *QEvent) bool
//
//OnWindowUnblockedEvent(watched *QObject, event *QEvent) bool
//
//OnZOrderChangeEvent(event *QEvent) bool
//
//OnZOrderChangeEvent(watched *QObject, event *QEvent) bool
//
//OnZeroTimerEventEvent(event *QEvent) bool
//
//OnZeroTimerEventEvent(watched *QObject, event *QEvent) bool
//
//OnUserEvent(event *QEvent) bool
//
//OnUserEvent(watched *QObject, event *QEvent) bool
func (q *QObject) InstallEventFilter(filter interface{}) {
	drvInstallEventFilter(q, filter)
}

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
	case QEvent_ActionAdded:
		if v, ok := i.(interface {
			OnActionAddedEvent(*QObject, *QActionEvent) bool
		}); ok {
			e := &QActionEvent{}
			e.SetDriver(event, 5000, false)
			return v.OnActionAddedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActionAddedEvent(*QActionEvent) bool
		}); ok {
			e := &QActionEvent{}
			e.SetDriver(event, 5000, false)
			return v.OnActionAddedEvent(e)
		}
	case QEvent_ActionChanged:
		if v, ok := i.(interface {
			OnActionChangedEvent(*QObject, *QActionEvent) bool
		}); ok {
			e := &QActionEvent{}
			e.SetDriver(event, 5000, false)
			return v.OnActionChangedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActionChangedEvent(*QActionEvent) bool
		}); ok {
			e := &QActionEvent{}
			e.SetDriver(event, 5000, false)
			return v.OnActionChangedEvent(e)
		}
	case QEvent_ActionRemoved:
		if v, ok := i.(interface {
			OnActionRemovedEvent(*QObject, *QActionEvent) bool
		}); ok {
			e := &QActionEvent{}
			e.SetDriver(event, 5000, false)
			return v.OnActionRemovedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnActionRemovedEvent(*QActionEvent) bool
		}); ok {
			e := &QActionEvent{}
			e.SetDriver(event, 5000, false)
			return v.OnActionRemovedEvent(e)
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
	case QEvent_ChildAdded:
		if v, ok := i.(interface {
			OnChildAddedEvent(*QObject, *QChildEvent) bool
		}); ok {
			e := &QChildEvent{}
			e.SetDriver(event, 11000, false)
			return v.OnChildAddedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnChildAddedEvent(*QChildEvent) bool
		}); ok {
			e := &QChildEvent{}
			e.SetDriver(event, 11000, false)
			return v.OnChildAddedEvent(e)
		}
	case QEvent_ChildPolished:
		if v, ok := i.(interface {
			OnChildPolishedEvent(*QObject, *QChildEvent) bool
		}); ok {
			e := &QChildEvent{}
			e.SetDriver(event, 11000, false)
			return v.OnChildPolishedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnChildPolishedEvent(*QChildEvent) bool
		}); ok {
			e := &QChildEvent{}
			e.SetDriver(event, 11000, false)
			return v.OnChildPolishedEvent(e)
		}
	case QEvent_ChildRemoved:
		if v, ok := i.(interface {
			OnChildRemovedEvent(*QObject, *QChildEvent) bool
		}); ok {
			e := &QChildEvent{}
			e.SetDriver(event, 11000, false)
			return v.OnChildRemovedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnChildRemovedEvent(*QChildEvent) bool
		}); ok {
			e := &QChildEvent{}
			e.SetDriver(event, 11000, false)
			return v.OnChildRemovedEvent(e)
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
	case QEvent_DragLeave:
		if v, ok := i.(interface {
			OnDragLeaveEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDragLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDragLeaveEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDragLeaveEvent(e)
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
	case QEvent_Drop:
		if v, ok := i.(interface {
			OnDropEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDropEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDropEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDropEvent(e)
		}
	case QEvent_DynamicPropertyChange:
		if v, ok := i.(interface {
			OnDynamicPropertyChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDynamicPropertyChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnDynamicPropertyChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnDynamicPropertyChangeEvent(e)
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
	case QEvent_FocusIn:
		if v, ok := i.(interface {
			OnFocusInEvent(*QObject, *QFocusEvent) bool
		}); ok {
			e := &QFocusEvent{}
			e.SetDriver(event, 36000, false)
			return v.OnFocusInEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFocusInEvent(*QFocusEvent) bool
		}); ok {
			e := &QFocusEvent{}
			e.SetDriver(event, 36000, false)
			return v.OnFocusInEvent(e)
		}
	case QEvent_FocusOut:
		if v, ok := i.(interface {
			OnFocusOutEvent(*QObject, *QFocusEvent) bool
		}); ok {
			e := &QFocusEvent{}
			e.SetDriver(event, 36000, false)
			return v.OnFocusOutEvent(obj, e)
		} else if v, ok := i.(interface {
			OnFocusOutEvent(*QFocusEvent) bool
		}); ok {
			e := &QFocusEvent{}
			e.SetDriver(event, 36000, false)
			return v.OnFocusOutEvent(e)
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
			OnGestureOverrideEvent(*QObject, *QGestureEvent) bool
		}); ok {
			e := &QGestureEvent{}
			e.SetDriver(event, 44000, false)
			return v.OnGestureOverrideEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGestureOverrideEvent(*QGestureEvent) bool
		}); ok {
			e := &QGestureEvent{}
			e.SetDriver(event, 44000, false)
			return v.OnGestureOverrideEvent(e)
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
	case QEvent_GraphicsSceneDragEnter:
		if v, ok := i.(interface {
			OnGraphicsSceneDragEnterEvent(*QObject, *QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDragEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDragEnterEvent(*QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDragEnterEvent(e)
		}
	case QEvent_GraphicsSceneDragLeave:
		if v, ok := i.(interface {
			OnGraphicsSceneDragLeaveEvent(*QObject, *QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDragLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDragLeaveEvent(*QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDragLeaveEvent(e)
		}
	case QEvent_GraphicsSceneDragMove:
		if v, ok := i.(interface {
			OnGraphicsSceneDragMoveEvent(*QObject, *QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDragMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDragMoveEvent(*QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDragMoveEvent(e)
		}
	case QEvent_GraphicsSceneDrop:
		if v, ok := i.(interface {
			OnGraphicsSceneDropEvent(*QObject, *QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDropEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneDropEvent(*QGraphicsSceneDragDropEvent) bool
		}); ok {
			e := &QGraphicsSceneDragDropEvent{}
			e.SetDriver(event, 276000, false)
			return v.OnGraphicsSceneDropEvent(e)
		}
	case QEvent_GraphicsSceneHelp:
		if v, ok := i.(interface {
			OnGraphicsSceneHelpEvent(*QObject, *QHelpEvent) bool
		}); ok {
			e := &QHelpEvent{}
			e.SetDriver(event, 48000, false)
			return v.OnGraphicsSceneHelpEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHelpEvent(*QHelpEvent) bool
		}); ok {
			e := &QHelpEvent{}
			e.SetDriver(event, 48000, false)
			return v.OnGraphicsSceneHelpEvent(e)
		}
	case QEvent_GraphicsSceneHoverEnter:
		if v, ok := i.(interface {
			OnGraphicsSceneHoverEnterEvent(*QObject, *QGraphicsSceneHoverEvent) bool
		}); ok {
			e := &QGraphicsSceneHoverEvent{}
			e.SetDriver(event, 279000, false)
			return v.OnGraphicsSceneHoverEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHoverEnterEvent(*QGraphicsSceneHoverEvent) bool
		}); ok {
			e := &QGraphicsSceneHoverEvent{}
			e.SetDriver(event, 279000, false)
			return v.OnGraphicsSceneHoverEnterEvent(e)
		}
	case QEvent_GraphicsSceneHoverLeave:
		if v, ok := i.(interface {
			OnGraphicsSceneHoverLeaveEvent(*QObject, *QGraphicsSceneHoverEvent) bool
		}); ok {
			e := &QGraphicsSceneHoverEvent{}
			e.SetDriver(event, 279000, false)
			return v.OnGraphicsSceneHoverLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHoverLeaveEvent(*QGraphicsSceneHoverEvent) bool
		}); ok {
			e := &QGraphicsSceneHoverEvent{}
			e.SetDriver(event, 279000, false)
			return v.OnGraphicsSceneHoverLeaveEvent(e)
		}
	case QEvent_GraphicsSceneHoverMove:
		if v, ok := i.(interface {
			OnGraphicsSceneHoverMoveEvent(*QObject, *QGraphicsSceneHoverEvent) bool
		}); ok {
			e := &QGraphicsSceneHoverEvent{}
			e.SetDriver(event, 279000, false)
			return v.OnGraphicsSceneHoverMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneHoverMoveEvent(*QGraphicsSceneHoverEvent) bool
		}); ok {
			e := &QGraphicsSceneHoverEvent{}
			e.SetDriver(event, 279000, false)
			return v.OnGraphicsSceneHoverMoveEvent(e)
		}
	case QEvent_GraphicsSceneMouseDoubleClick:
		if v, ok := i.(interface {
			OnGraphicsSceneMouseDoubleClickEvent(*QObject, *QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMouseDoubleClickEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMouseDoubleClickEvent(*QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMouseDoubleClickEvent(e)
		}
	case QEvent_GraphicsSceneMouseMove:
		if v, ok := i.(interface {
			OnGraphicsSceneMouseMoveEvent(*QObject, *QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMouseMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMouseMoveEvent(*QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMouseMoveEvent(e)
		}
	case QEvent_GraphicsSceneMousePress:
		if v, ok := i.(interface {
			OnGraphicsSceneMousePressEvent(*QObject, *QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMousePressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMousePressEvent(*QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMousePressEvent(e)
		}
	case QEvent_GraphicsSceneMouseRelease:
		if v, ok := i.(interface {
			OnGraphicsSceneMouseReleaseEvent(*QObject, *QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMouseReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnGraphicsSceneMouseReleaseEvent(*QGraphicsSceneMouseEvent) bool
		}); ok {
			e := &QGraphicsSceneMouseEvent{}
			e.SetDriver(event, 280000, false)
			return v.OnGraphicsSceneMouseReleaseEvent(e)
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
	case QEvent_HoverEnter:
		if v, ok := i.(interface {
			OnHoverEnterEvent(*QObject, *QHoverEvent) bool
		}); ok {
			e := &QHoverEvent{}
			e.SetDriver(event, 50000, false)
			return v.OnHoverEnterEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHoverEnterEvent(*QHoverEvent) bool
		}); ok {
			e := &QHoverEvent{}
			e.SetDriver(event, 50000, false)
			return v.OnHoverEnterEvent(e)
		}
	case QEvent_HoverLeave:
		if v, ok := i.(interface {
			OnHoverLeaveEvent(*QObject, *QHoverEvent) bool
		}); ok {
			e := &QHoverEvent{}
			e.SetDriver(event, 50000, false)
			return v.OnHoverLeaveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHoverLeaveEvent(*QHoverEvent) bool
		}); ok {
			e := &QHoverEvent{}
			e.SetDriver(event, 50000, false)
			return v.OnHoverLeaveEvent(e)
		}
	case QEvent_HoverMove:
		if v, ok := i.(interface {
			OnHoverMoveEvent(*QObject, *QHoverEvent) bool
		}); ok {
			e := &QHoverEvent{}
			e.SetDriver(event, 50000, false)
			return v.OnHoverMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnHoverMoveEvent(*QHoverEvent) bool
		}); ok {
			e := &QHoverEvent{}
			e.SetDriver(event, 50000, false)
			return v.OnHoverMoveEvent(e)
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
	case QEvent_KeyPress:
		if v, ok := i.(interface {
			OnKeyPressEvent(*QObject, *QKeyEvent) bool
		}); ok {
			e := &QKeyEvent{}
			e.SetDriver(event, 65000, false)
			return v.OnKeyPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnKeyPressEvent(*QKeyEvent) bool
		}); ok {
			e := &QKeyEvent{}
			e.SetDriver(event, 65000, false)
			return v.OnKeyPressEvent(e)
		}
	case QEvent_KeyRelease:
		if v, ok := i.(interface {
			OnKeyReleaseEvent(*QObject, *QKeyEvent) bool
		}); ok {
			e := &QKeyEvent{}
			e.SetDriver(event, 65000, false)
			return v.OnKeyReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnKeyReleaseEvent(*QKeyEvent) bool
		}); ok {
			e := &QKeyEvent{}
			e.SetDriver(event, 65000, false)
			return v.OnKeyReleaseEvent(e)
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
	case QEvent_MouseButtonDblClick:
		if v, ok := i.(interface {
			OnMouseButtonDblClickEvent(*QObject, *QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseButtonDblClickEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseButtonDblClickEvent(*QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseButtonDblClickEvent(e)
		}
	case QEvent_MouseButtonPress:
		if v, ok := i.(interface {
			OnMouseButtonPressEvent(*QObject, *QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseButtonPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseButtonPressEvent(*QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseButtonPressEvent(e)
		}
	case QEvent_MouseButtonRelease:
		if v, ok := i.(interface {
			OnMouseButtonReleaseEvent(*QObject, *QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseButtonReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseButtonReleaseEvent(*QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseButtonReleaseEvent(e)
		}
	case QEvent_MouseMove:
		if v, ok := i.(interface {
			OnMouseMoveEvent(*QObject, *QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnMouseMoveEvent(*QMouseEvent) bool
		}); ok {
			e := &QMouseEvent{}
			e.SetDriver(event, 81000, false)
			return v.OnMouseMoveEvent(e)
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
			OnShortcutOverrideEvent(*QObject, *QKeyEvent) bool
		}); ok {
			e := &QKeyEvent{}
			e.SetDriver(event, 65000, false)
			return v.OnShortcutOverrideEvent(obj, e)
		} else if v, ok := i.(interface {
			OnShortcutOverrideEvent(*QKeyEvent) bool
		}); ok {
			e := &QKeyEvent{}
			e.SetDriver(event, 65000, false)
			return v.OnShortcutOverrideEvent(e)
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
	case QEvent_StateMachineSignal:
		if v, ok := i.(interface {
			OnStateMachineSignalEvent(*QObject, *QStateMachineSignalEvent) bool
		}); ok {
			e := &QStateMachineSignalEvent{}
			e.SetDriver(event, 125000, false)
			return v.OnStateMachineSignalEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStateMachineSignalEvent(*QStateMachineSignalEvent) bool
		}); ok {
			e := &QStateMachineSignalEvent{}
			e.SetDriver(event, 125000, false)
			return v.OnStateMachineSignalEvent(e)
		}
	case QEvent_StateMachineWrapped:
		if v, ok := i.(interface {
			OnStateMachineWrappedEvent(*QObject, *QStateMachineWrappedEvent) bool
		}); ok {
			e := &QStateMachineWrappedEvent{}
			e.SetDriver(event, 126000, false)
			return v.OnStateMachineWrappedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnStateMachineWrappedEvent(*QStateMachineWrappedEvent) bool
		}); ok {
			e := &QStateMachineWrappedEvent{}
			e.SetDriver(event, 126000, false)
			return v.OnStateMachineWrappedEvent(e)
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
	case QEvent_TabletEnterProximity:
		if v, ok := i.(interface {
			OnTabletEnterProximityEvent(*QObject, *QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletEnterProximityEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletEnterProximityEvent(*QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletEnterProximityEvent(e)
		}
	case QEvent_TabletLeaveProximity:
		if v, ok := i.(interface {
			OnTabletLeaveProximityEvent(*QObject, *QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletLeaveProximityEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletLeaveProximityEvent(*QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletLeaveProximityEvent(e)
		}
	case QEvent_TabletMove:
		if v, ok := i.(interface {
			OnTabletMoveEvent(*QObject, *QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletMoveEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletMoveEvent(*QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletMoveEvent(e)
		}
	case QEvent_TabletPress:
		if v, ok := i.(interface {
			OnTabletPressEvent(*QObject, *QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletPressEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletPressEvent(*QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletPressEvent(e)
		}
	case QEvent_TabletRelease:
		if v, ok := i.(interface {
			OnTabletReleaseEvent(*QObject, *QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletReleaseEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTabletReleaseEvent(*QTabletEvent) bool
		}); ok {
			e := &QTabletEvent{}
			e.SetDriver(event, 138000, false)
			return v.OnTabletReleaseEvent(e)
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
	case QEvent_ToolBarChange:
		if v, ok := i.(interface {
			OnToolBarChangeEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnToolBarChangeEvent(obj, e)
		} else if v, ok := i.(interface {
			OnToolBarChangeEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnToolBarChangeEvent(e)
		}
	case QEvent_ToolTip:
		if v, ok := i.(interface {
			OnToolTipEvent(*QObject, *QHelpEvent) bool
		}); ok {
			e := &QHelpEvent{}
			e.SetDriver(event, 48000, false)
			return v.OnToolTipEvent(obj, e)
		} else if v, ok := i.(interface {
			OnToolTipEvent(*QHelpEvent) bool
		}); ok {
			e := &QHelpEvent{}
			e.SetDriver(event, 48000, false)
			return v.OnToolTipEvent(e)
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
	case QEvent_TouchBegin:
		if v, ok := i.(interface {
			OnTouchBeginEvent(*QObject, *QTouchEvent) bool
		}); ok {
			e := &QTouchEvent{}
			e.SetDriver(event, 176000, false)
			return v.OnTouchBeginEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTouchBeginEvent(*QTouchEvent) bool
		}); ok {
			e := &QTouchEvent{}
			e.SetDriver(event, 176000, false)
			return v.OnTouchBeginEvent(e)
		}
	case QEvent_TouchEnd:
		if v, ok := i.(interface {
			OnTouchEndEvent(*QObject, *QTouchEvent) bool
		}); ok {
			e := &QTouchEvent{}
			e.SetDriver(event, 176000, false)
			return v.OnTouchEndEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTouchEndEvent(*QTouchEvent) bool
		}); ok {
			e := &QTouchEvent{}
			e.SetDriver(event, 176000, false)
			return v.OnTouchEndEvent(e)
		}
	case QEvent_TouchUpdate:
		if v, ok := i.(interface {
			OnTouchUpdateEvent(*QObject, *QTouchEvent) bool
		}); ok {
			e := &QTouchEvent{}
			e.SetDriver(event, 176000, false)
			return v.OnTouchUpdateEvent(obj, e)
		} else if v, ok := i.(interface {
			OnTouchUpdateEvent(*QTouchEvent) bool
		}); ok {
			e := &QTouchEvent{}
			e.SetDriver(event, 176000, false)
			return v.OnTouchUpdateEvent(e)
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
	case QEvent_WhatsThisClicked:
		if v, ok := i.(interface {
			OnWhatsThisClickedEvent(*QObject, *QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWhatsThisClickedEvent(obj, e)
		} else if v, ok := i.(interface {
			OnWhatsThisClickedEvent(*QEvent) bool
		}); ok {
			e := &QEvent{}
			e.SetDriver(event, 31000, false)
			return v.OnWhatsThisClickedEvent(e)
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
