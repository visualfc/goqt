// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui	

//enum Qt_ShortcutContext - Qt::ShortcutContext
type Qt_ShortcutContext uint32
const (
	Qt_WidgetShortcut Qt_ShortcutContext = 0
	Qt_WindowShortcut Qt_ShortcutContext = 1
	Qt_ApplicationShortcut Qt_ShortcutContext = 2
	Qt_WidgetWithChildrenShortcut Qt_ShortcutContext = 3
)

//enum Qt_CheckState - Qt::CheckState
type Qt_CheckState uint32
const (
	Qt_Unchecked Qt_CheckState = 0
	Qt_PartiallyChecked Qt_CheckState = 1
	Qt_Checked Qt_CheckState = 2
)

//enum Qt_FocusPolicy - Qt::FocusPolicy
type Qt_FocusPolicy uint32
const (
	Qt_NoFocus Qt_FocusPolicy = 0
	Qt_TabFocus Qt_FocusPolicy = 0x1
	Qt_ClickFocus Qt_FocusPolicy = 0x2
	Qt_StrongFocus Qt_FocusPolicy = Qt_TabFocus|Qt_ClickFocus|0x8
	Qt_WheelFocus Qt_FocusPolicy = Qt_StrongFocus|0x4
)

//enum Qt_GestureType - Qt::GestureType
type Qt_GestureType uint32
const (
	Qt_TapGesture Qt_GestureType = 1
	Qt_TapAndHoldGesture Qt_GestureType = 2
	Qt_PanGesture Qt_GestureType = 3
	Qt_PinchGesture Qt_GestureType = 4
	Qt_SwipeGesture Qt_GestureType = 5
	Qt_CustomGesture Qt_GestureType = 0x0100
	Qt_LastGestureType Qt_GestureType = ^Qt_GestureType(0)
)

//enum Qt_DropAction - Qt::DropAction
type Qt_DropAction uint32
const (
	Qt_CopyAction Qt_DropAction = 0x1
	Qt_MoveAction Qt_DropAction = 0x2
	Qt_LinkAction Qt_DropAction = 0x4
	Qt_ActionMask Qt_DropAction = 0xff
	Qt_TargetMoveAction Qt_DropAction = 0x8002
	Qt_IgnoreAction Qt_DropAction = 0x0
)

//enum Qt_WindowType - Qt::WindowType
type Qt_WindowType uint32
const (
	Qt_Widget Qt_WindowType = 0x00000000
	Qt_Window Qt_WindowType = 0x00000001
	Qt_Dialog Qt_WindowType = 0x00000002|Qt_Window
	Qt_Sheet Qt_WindowType = 0x00000004|Qt_Window
	Qt_Drawer Qt_WindowType = 0x00000006|Qt_Window
	Qt_Popup Qt_WindowType = 0x00000008|Qt_Window
	Qt_Tool Qt_WindowType = 0x0000000a|Qt_Window
	Qt_ToolTip Qt_WindowType = 0x0000000c|Qt_Window
	Qt_SplashScreen Qt_WindowType = 0x0000000e|Qt_Window
	Qt_Desktop Qt_WindowType = 0x00000010|Qt_Window
	Qt_SubWindow Qt_WindowType = 0x00000012
	Qt_WindowType_Mask Qt_WindowType = 0x000000ff
	Qt_MSWindowsFixedSizeDialogHint Qt_WindowType = 0x00000100
	Qt_MSWindowsOwnDC Qt_WindowType = 0x00000200
	Qt_X11BypassWindowManagerHint Qt_WindowType = 0x00000400
	Qt_FramelessWindowHint Qt_WindowType = 0x00000800
	Qt_WindowTitleHint Qt_WindowType = 0x00001000
	Qt_WindowSystemMenuHint Qt_WindowType = 0x00002000
	Qt_WindowMinimizeButtonHint Qt_WindowType = 0x00004000
	Qt_WindowMaximizeButtonHint Qt_WindowType = 0x00008000
	Qt_WindowMinMaxButtonsHint Qt_WindowType = Qt_WindowMinimizeButtonHint|Qt_WindowMaximizeButtonHint
	Qt_WindowContextHelpButtonHint Qt_WindowType = 0x00010000
	Qt_WindowShadeButtonHint Qt_WindowType = 0x00020000
	Qt_WindowStaysOnTopHint Qt_WindowType = 0x00040000
	Qt_CustomizeWindowHint Qt_WindowType = 0x02000000
	Qt_WindowStaysOnBottomHint Qt_WindowType = 0x04000000
	Qt_WindowCloseButtonHint Qt_WindowType = 0x08000000
	Qt_MacWindowToolBarButtonHint Qt_WindowType = 0x10000000
	Qt_BypassGraphicsProxyWidget Qt_WindowType = 0x20000000
	Qt_WindowOkButtonHint Qt_WindowType = 0x00080000
	Qt_WindowCancelButtonHint Qt_WindowType = 0x00100000
	Qt_WindowSoftkeysVisibleHint Qt_WindowType = 0x40000000
	Qt_WindowSoftkeysRespondHint Qt_WindowType = 0x80000000
)

//enum Qt_DateFormat - Qt::DateFormat
type Qt_DateFormat uint32
const (
	Qt_TextDate Qt_DateFormat = 0
	Qt_ISODate Qt_DateFormat = 1
	Qt_SystemLocaleDate Qt_DateFormat = 2
	Qt_LocalDate Qt_DateFormat = Qt_SystemLocaleDate
	Qt_LocaleDate Qt_DateFormat = Qt_SystemLocaleDate+1
	Qt_SystemLocaleShortDate Qt_DateFormat = Qt_SystemLocaleDate+1+1
	Qt_SystemLocaleLongDate Qt_DateFormat = Qt_SystemLocaleDate+1+1+1
	Qt_DefaultLocaleShortDate Qt_DateFormat = Qt_SystemLocaleDate+1+1+1+1
	Qt_DefaultLocaleLongDate Qt_DateFormat = Qt_SystemLocaleDate+1+1+1+1+1
)

//enum Qt_TextFormat - Qt::TextFormat
type Qt_TextFormat uint32
const (
	Qt_PlainText Qt_TextFormat = 0
	Qt_RichText Qt_TextFormat = 1
	Qt_AutoText Qt_TextFormat = 2
	Qt_LogText Qt_TextFormat = 3
)

//enum Qt_PenJoinStyle - Qt::PenJoinStyle
type Qt_PenJoinStyle uint32
const (
	Qt_MiterJoin Qt_PenJoinStyle = 0x00
	Qt_BevelJoin Qt_PenJoinStyle = 0x40
	Qt_RoundJoin Qt_PenJoinStyle = 0x80
	Qt_SvgMiterJoin Qt_PenJoinStyle = 0x100
	Qt_MPenJoinStyle Qt_PenJoinStyle = 0x1c0
)

//enum Qt_GestureFlag - Qt::GestureFlag
type Qt_GestureFlag uint32
const (
	Qt_DontStartGestureOnChildren Qt_GestureFlag = 0x01
	Qt_ReceivePartialGestures Qt_GestureFlag = 0x02
	Qt_IgnoredGesturesPropagateToParent Qt_GestureFlag = 0x04
)

//enum Qt_CaseSensitivity - Qt::CaseSensitivity
type Qt_CaseSensitivity uint32
const (
	Qt_CaseInsensitive Qt_CaseSensitivity = 0
	Qt_CaseSensitive Qt_CaseSensitivity = 1
)

//enum Qt_EventPriority - Qt::EventPriority
type Qt_EventPriority int32
const (
	Qt_HighEventPriority Qt_EventPriority = 1
	Qt_NormalEventPriority Qt_EventPriority = 0
	Qt_LowEventPriority Qt_EventPriority = -1
)

//enum Qt_ImageConversionFlag - Qt::ImageConversionFlag
type Qt_ImageConversionFlag uint32
const (
	Qt_ColorMode_Mask Qt_ImageConversionFlag = 0x00000003
	Qt_AutoColor Qt_ImageConversionFlag = 0x00000000
	Qt_ColorOnly Qt_ImageConversionFlag = 0x00000003
	Qt_MonoOnly Qt_ImageConversionFlag = 0x00000002
	Qt_AlphaDither_Mask Qt_ImageConversionFlag = 0x0000000c
	Qt_ThresholdAlphaDither Qt_ImageConversionFlag = 0x00000000
	Qt_OrderedAlphaDither Qt_ImageConversionFlag = 0x00000004
	Qt_DiffuseAlphaDither Qt_ImageConversionFlag = 0x00000008
	Qt_NoAlpha Qt_ImageConversionFlag = 0x0000000c
	Qt_Dither_Mask Qt_ImageConversionFlag = 0x00000030
	Qt_DiffuseDither Qt_ImageConversionFlag = 0x00000000
	Qt_OrderedDither Qt_ImageConversionFlag = 0x00000010
	Qt_ThresholdDither Qt_ImageConversionFlag = 0x00000020
	Qt_DitherMode_Mask Qt_ImageConversionFlag = 0x000000c0
	Qt_AutoDither Qt_ImageConversionFlag = 0x00000000
	Qt_PreferDither Qt_ImageConversionFlag = 0x00000040
	Qt_AvoidDither Qt_ImageConversionFlag = 0x00000080
	Qt_NoOpaqueDetection Qt_ImageConversionFlag = 0x00000100
	Qt_NoFormatConversion Qt_ImageConversionFlag = 0x00000200
)

//enum Qt_GestureState - Qt::GestureState
type Qt_GestureState uint32
const (
	Qt_NoGesture Qt_GestureState = 0
	Qt_GestureStarted Qt_GestureState = 1
	Qt_GestureUpdated Qt_GestureState = 2
	Qt_GestureFinished Qt_GestureState = 3
	Qt_GestureCanceled Qt_GestureState = 4
)

//enum Qt_BGMode - Qt::BGMode
type Qt_BGMode uint32
const (
	Qt_TransparentMode Qt_BGMode = 0
	Qt_OpaqueMode Qt_BGMode = 1
)

//enum Qt_ConnectionType - Qt::ConnectionType
type Qt_ConnectionType uint32
const (
	Qt_AutoConnection Qt_ConnectionType = 0
	Qt_DirectConnection Qt_ConnectionType = 1
	Qt_QueuedConnection Qt_ConnectionType = 2
	Qt_AutoCompatConnection Qt_ConnectionType = 3
	Qt_BlockingQueuedConnection Qt_ConnectionType = 4
	Qt_UniqueConnection Qt_ConnectionType = 0x80
)

//enum Qt_ToolBarArea - Qt::ToolBarArea
type Qt_ToolBarArea uint32
const (
	Qt_LeftToolBarArea Qt_ToolBarArea = 0x1
	Qt_RightToolBarArea Qt_ToolBarArea = 0x2
	Qt_TopToolBarArea Qt_ToolBarArea = 0x4
	Qt_BottomToolBarArea Qt_ToolBarArea = 0x8
	Qt_ToolBarArea_Mask Qt_ToolBarArea = 0xf
	Qt_AllToolBarAreas Qt_ToolBarArea = Qt_ToolBarArea_Mask
	Qt_NoToolBarArea Qt_ToolBarArea = 0
)

//enum Qt_CoordinateSystem - Qt::CoordinateSystem
type Qt_CoordinateSystem uint32
const (
	Qt_DeviceCoordinates Qt_CoordinateSystem = 0
	Qt_LogicalCoordinates Qt_CoordinateSystem = 1
)

//enum Qt_SizeMode - Qt::SizeMode
type Qt_SizeMode uint32
const (
	Qt_AbsoluteSize Qt_SizeMode = 0
	Qt_RelativeSize Qt_SizeMode = 1
)

//enum Qt_FocusReason - Qt::FocusReason
type Qt_FocusReason uint32
const (
	Qt_MouseFocusReason Qt_FocusReason = 0
	Qt_TabFocusReason Qt_FocusReason = 1
	Qt_BacktabFocusReason Qt_FocusReason = 2
	Qt_ActiveWindowFocusReason Qt_FocusReason = 3
	Qt_PopupFocusReason Qt_FocusReason = 4
	Qt_ShortcutFocusReason Qt_FocusReason = 5
	Qt_MenuBarFocusReason Qt_FocusReason = 6
	Qt_OtherFocusReason Qt_FocusReason = 7
	Qt_NoFocusReason Qt_FocusReason = 8
)

//enum Qt_Modifier - Qt::Modifier
type Qt_Modifier uint32
const (
	Qt_META Qt_Modifier = Qt_Modifier(Qt_MetaModifier)
	Qt_SHIFT Qt_Modifier = Qt_Modifier(Qt_ShiftModifier)
	Qt_CTRL Qt_Modifier = Qt_Modifier(Qt_ControlModifier)
	Qt_ALT Qt_Modifier = Qt_Modifier(Qt_AltModifier)
	Qt_MODIFIER_MASK Qt_Modifier = Qt_Modifier(Qt_KeyboardModifierMask)
	Qt_UNICODE_ACCEL Qt_Modifier = 0x00000000
)

//enum Qt_TileRule - Qt::TileRule
type Qt_TileRule uint32
const (
	Qt_StretchTile Qt_TileRule = 0
	Qt_RepeatTile Qt_TileRule = 1
	Qt_RoundTile Qt_TileRule = 2
)

//enum Qt_WhiteSpaceMode - Qt::WhiteSpaceMode
type Qt_WhiteSpaceMode int32
const (
	Qt_WhiteSpaceNormal Qt_WhiteSpaceMode = 0
	Qt_WhiteSpacePre Qt_WhiteSpaceMode = 1
	Qt_WhiteSpaceNoWrap Qt_WhiteSpaceMode = 2
	Qt_WhiteSpaceModeUndefined Qt_WhiteSpaceMode = -1
)

//enum Qt_AspectRatioMode - Qt::AspectRatioMode
type Qt_AspectRatioMode uint32
const (
	Qt_IgnoreAspectRatio Qt_AspectRatioMode = 0
	Qt_KeepAspectRatio Qt_AspectRatioMode = 1
	Qt_KeepAspectRatioByExpanding Qt_AspectRatioMode = 2
)

//enum Qt_SizeHint - Qt::SizeHint
type Qt_SizeHint uint32
const (
	Qt_MinimumSize Qt_SizeHint = 0
	Qt_PreferredSize Qt_SizeHint = 1
	Qt_MaximumSize Qt_SizeHint = 2
	Qt_MinimumDescent Qt_SizeHint = 3
	Qt_NSizeHints Qt_SizeHint = 4
)

//enum Qt_AlignmentFlag - Qt::AlignmentFlag
type Qt_AlignmentFlag uint32
const (
	Qt_AlignLeft Qt_AlignmentFlag = 0x0001
	Qt_AlignLeading Qt_AlignmentFlag = Qt_AlignLeft
	Qt_AlignRight Qt_AlignmentFlag = 0x0002
	Qt_AlignTrailing Qt_AlignmentFlag = Qt_AlignRight
	Qt_AlignHCenter Qt_AlignmentFlag = 0x0004
	Qt_AlignJustify Qt_AlignmentFlag = 0x0008
	Qt_AlignAbsolute Qt_AlignmentFlag = 0x0010
	Qt_AlignHorizontal_Mask Qt_AlignmentFlag = Qt_AlignLeft|Qt_AlignRight|Qt_AlignHCenter|Qt_AlignJustify|Qt_AlignAbsolute
	Qt_AlignTop Qt_AlignmentFlag = 0x0020
	Qt_AlignBottom Qt_AlignmentFlag = 0x0040
	Qt_AlignVCenter Qt_AlignmentFlag = 0x0080
	Qt_AlignVertical_Mask Qt_AlignmentFlag = Qt_AlignTop|Qt_AlignBottom|Qt_AlignVCenter
	Qt_AlignCenter Qt_AlignmentFlag = Qt_AlignVCenter|Qt_AlignHCenter
)

//enum Qt_ContextMenuPolicy - Qt::ContextMenuPolicy
type Qt_ContextMenuPolicy uint32
const (
	Qt_NoContextMenu Qt_ContextMenuPolicy = 0
	Qt_DefaultContextMenu Qt_ContextMenuPolicy = 1
	Qt_ActionsContextMenu Qt_ContextMenuPolicy = 2
	Qt_CustomContextMenu Qt_ContextMenuPolicy = 3
	Qt_PreventContextMenu Qt_ContextMenuPolicy = 4
)

//enum Qt_DockWidgetArea - Qt::DockWidgetArea
type Qt_DockWidgetArea uint32
const (
	Qt_LeftDockWidgetArea Qt_DockWidgetArea = 0x1
	Qt_RightDockWidgetArea Qt_DockWidgetArea = 0x2
	Qt_TopDockWidgetArea Qt_DockWidgetArea = 0x4
	Qt_BottomDockWidgetArea Qt_DockWidgetArea = 0x8
	Qt_DockWidgetArea_Mask Qt_DockWidgetArea = 0xf
	Qt_AllDockWidgetAreas Qt_DockWidgetArea = Qt_DockWidgetArea_Mask
	Qt_NoDockWidgetArea Qt_DockWidgetArea = 0
)

//enum Qt_UIEffect - Qt::UIEffect
type Qt_UIEffect uint32
const (
	Qt_UI_General Qt_UIEffect = 0
	Qt_UI_AnimateMenu Qt_UIEffect = 1
	Qt_UI_FadeMenu Qt_UIEffect = 2
	Qt_UI_AnimateCombo Qt_UIEffect = 3
	Qt_UI_AnimateTooltip Qt_UIEffect = 4
	Qt_UI_FadeTooltip Qt_UIEffect = 5
	Qt_UI_AnimateToolBox Qt_UIEffect = 6
)

//enum Qt_Initialization - Qt::Initialization
type Qt_Initialization uint32
const (
	Qt_Uninitialized Qt_Initialization = 0
)

//enum Qt_AnchorPoint - Qt::AnchorPoint
type Qt_AnchorPoint uint32
const (
	Qt_AnchorLeft Qt_AnchorPoint = 0
	Qt_AnchorHorizontalCenter Qt_AnchorPoint = 0
	Qt_AnchorRight Qt_AnchorPoint = 1
	Qt_AnchorTop Qt_AnchorPoint = 2
	Qt_AnchorVerticalCenter Qt_AnchorPoint = 3
	Qt_AnchorBottom Qt_AnchorPoint = 4
)

//enum Qt_GlobalColor - Qt::GlobalColor
type Qt_GlobalColor uint32
const (
	Qt_color0 Qt_GlobalColor = 0
	Qt_color1 Qt_GlobalColor = 1
	Qt_black Qt_GlobalColor = 2
	Qt_white Qt_GlobalColor = 3
	Qt_darkGray Qt_GlobalColor = 4
	Qt_gray Qt_GlobalColor = 5
	Qt_lightGray Qt_GlobalColor = 6
	Qt_red Qt_GlobalColor = 7
	Qt_green Qt_GlobalColor = 8
	Qt_blue Qt_GlobalColor = 9
	Qt_cyan Qt_GlobalColor = 10
	Qt_magenta Qt_GlobalColor = 11
	Qt_yellow Qt_GlobalColor = 12
	Qt_darkRed Qt_GlobalColor = 13
	Qt_darkGreen Qt_GlobalColor = 14
	Qt_darkBlue Qt_GlobalColor = 15
	Qt_darkCyan Qt_GlobalColor = 16
	Qt_darkMagenta Qt_GlobalColor = 17
	Qt_darkYellow Qt_GlobalColor = 18
	Qt_transparent Qt_GlobalColor = 19
)

//enum Qt_KeyboardModifier - Qt::KeyboardModifier
type Qt_KeyboardModifier uint32
const (
	Qt_NoModifier Qt_KeyboardModifier = 0x00000000
	Qt_ShiftModifier Qt_KeyboardModifier = 0x02000000
	Qt_ControlModifier Qt_KeyboardModifier = 0x04000000
	Qt_AltModifier Qt_KeyboardModifier = 0x08000000
	Qt_MetaModifier Qt_KeyboardModifier = 0x10000000
	Qt_KeypadModifier Qt_KeyboardModifier = 0x20000000
	Qt_GroupSwitchModifier Qt_KeyboardModifier = 0x40000000
	Qt_KeyboardModifierMask Qt_KeyboardModifier = 0xfe000000
)

//enum Qt_NavigationMode - Qt::NavigationMode
type Qt_NavigationMode uint32
const (
	Qt_NavigationModeNone Qt_NavigationMode = 0
	Qt_NavigationModeKeypadTabOrder Qt_NavigationMode = 1
	Qt_NavigationModeKeypadDirectional Qt_NavigationMode = 2
	Qt_NavigationModeCursorAuto Qt_NavigationMode = 3
	Qt_NavigationModeCursorForceVisible Qt_NavigationMode = 4
)

//enum Qt_ItemDataRole - Qt::ItemDataRole
type Qt_ItemDataRole uint32
const (
	Qt_DisplayRole Qt_ItemDataRole = 0
	Qt_DecorationRole Qt_ItemDataRole = 1
	Qt_EditRole Qt_ItemDataRole = 2
	Qt_ToolTipRole Qt_ItemDataRole = 3
	Qt_StatusTipRole Qt_ItemDataRole = 4
	Qt_WhatsThisRole Qt_ItemDataRole = 5
	Qt_FontRole Qt_ItemDataRole = 6
	Qt_TextAlignmentRole Qt_ItemDataRole = 7
	Qt_BackgroundColorRole Qt_ItemDataRole = 8
	Qt_BackgroundRole Qt_ItemDataRole = 8
	Qt_TextColorRole Qt_ItemDataRole = 9
	Qt_ForegroundRole Qt_ItemDataRole = 9
	Qt_CheckStateRole Qt_ItemDataRole = 10
	Qt_AccessibleTextRole Qt_ItemDataRole = 11
	Qt_AccessibleDescriptionRole Qt_ItemDataRole = 12
	Qt_SizeHintRole Qt_ItemDataRole = 13
	Qt_DisplayPropertyRole Qt_ItemDataRole = 27
	Qt_DecorationPropertyRole Qt_ItemDataRole = 28
	Qt_ToolTipPropertyRole Qt_ItemDataRole = 29
	Qt_StatusTipPropertyRole Qt_ItemDataRole = 30
	Qt_WhatsThisPropertyRole Qt_ItemDataRole = 31
	Qt_UserRole Qt_ItemDataRole = 32
)

//enum Qt_ScrollBarPolicy - Qt::ScrollBarPolicy
type Qt_ScrollBarPolicy uint32
const (
	Qt_ScrollBarAsNeeded Qt_ScrollBarPolicy = 0
	Qt_ScrollBarAlwaysOff Qt_ScrollBarPolicy = 1
	Qt_ScrollBarAlwaysOn Qt_ScrollBarPolicy = 2
)

//enum Qt_InputMethodHint - Qt::InputMethodHint
type Qt_InputMethodHint uint32
const (
	Qt_ImhNone Qt_InputMethodHint = 0x0
	Qt_ImhHiddenText Qt_InputMethodHint = 0x1
	Qt_ImhNoAutoUppercase Qt_InputMethodHint = 0x2
	Qt_ImhPreferNumbers Qt_InputMethodHint = 0x4
	Qt_ImhPreferUppercase Qt_InputMethodHint = 0x8
	Qt_ImhPreferLowercase Qt_InputMethodHint = 0x10
	Qt_ImhNoPredictiveText Qt_InputMethodHint = 0x20
	Qt_ImhDigitsOnly Qt_InputMethodHint = 0x10000
	Qt_ImhFormattedNumbersOnly Qt_InputMethodHint = 0x20000
	Qt_ImhUppercaseOnly Qt_InputMethodHint = 0x40000
	Qt_ImhLowercaseOnly Qt_InputMethodHint = 0x80000
	Qt_ImhDialableCharactersOnly Qt_InputMethodHint = 0x100000
	Qt_ImhEmailCharactersOnly Qt_InputMethodHint = 0x200000
	Qt_ImhUrlCharactersOnly Qt_InputMethodHint = 0x400000
	Qt_ImhExclusiveInputMask Qt_InputMethodHint = 0xffff0000
)

//enum Qt_AnchorAttribute - Qt::AnchorAttribute
type Qt_AnchorAttribute uint32
const (
	Qt_AnchorName Qt_AnchorAttribute = 0
	Qt_AnchorHref Qt_AnchorAttribute = 1
)

//enum Qt_WindowModality - Qt::WindowModality
type Qt_WindowModality uint32
const (
	Qt_NonModal Qt_WindowModality = 0
	Qt_WindowModal Qt_WindowModality = 1
	Qt_ApplicationModal Qt_WindowModality = 2
)

//enum Qt_SortOrder - Qt::SortOrder
type Qt_SortOrder uint32
const (
	Qt_AscendingOrder Qt_SortOrder = 0
	Qt_DescendingOrder Qt_SortOrder = 1
)

//enum Qt_PenStyle - Qt::PenStyle
type Qt_PenStyle uint32
const (
	Qt_NoPen Qt_PenStyle = 0
	Qt_SolidLine Qt_PenStyle = 1
	Qt_DashLine Qt_PenStyle = 2
	Qt_DotLine Qt_PenStyle = 3
	Qt_DashDotLine Qt_PenStyle = 4
	Qt_DashDotDotLine Qt_PenStyle = 5
	Qt_CustomDashLine Qt_PenStyle = 6
	Qt_MPenStyle Qt_PenStyle = 0x0f
)

//enum Qt_ItemFlag - Qt::ItemFlag
type Qt_ItemFlag uint32
const (
	Qt_NoItemFlags Qt_ItemFlag = 0
	Qt_ItemIsSelectable Qt_ItemFlag = 1
	Qt_ItemIsEditable Qt_ItemFlag = 2
	Qt_ItemIsDragEnabled Qt_ItemFlag = 4
	Qt_ItemIsDropEnabled Qt_ItemFlag = 8
	Qt_ItemIsUserCheckable Qt_ItemFlag = 16
	Qt_ItemIsEnabled Qt_ItemFlag = 32
	Qt_ItemIsTristate Qt_ItemFlag = 64
)

//enum Qt_Axis - Qt::Axis
type Qt_Axis uint32
const (
	Qt_XAxis Qt_Axis = 0
	Qt_YAxis Qt_Axis = 1
	Qt_ZAxis Qt_Axis = 2
)

//enum Qt_TransformationMode - Qt::TransformationMode
type Qt_TransformationMode uint32
const (
	Qt_FastTransformation Qt_TransformationMode = 0
	Qt_SmoothTransformation Qt_TransformationMode = 1
)

//enum Qt_WindowFrameSection - Qt::WindowFrameSection
type Qt_WindowFrameSection uint32
const (
	Qt_NoSection Qt_WindowFrameSection = 0
	Qt_LeftSection Qt_WindowFrameSection = 1
	Qt_TopLeftSection Qt_WindowFrameSection = 2
	Qt_TopSection Qt_WindowFrameSection = 3
	Qt_TopRightSection Qt_WindowFrameSection = 4
	Qt_RightSection Qt_WindowFrameSection = 5
	Qt_BottomRightSection Qt_WindowFrameSection = 6
	Qt_BottomSection Qt_WindowFrameSection = 7
	Qt_BottomLeftSection Qt_WindowFrameSection = 8
	Qt_TitleBarArea Qt_WindowFrameSection = 9
)

//enum Qt_HitTestAccuracy - Qt::HitTestAccuracy
type Qt_HitTestAccuracy uint32
const (
	Qt_ExactHit Qt_HitTestAccuracy = 0
	Qt_FuzzyHit Qt_HitTestAccuracy = 1
)

//enum Qt_CursorShape - Qt::CursorShape
type Qt_CursorShape uint32
const (
	Qt_ArrowCursor Qt_CursorShape = 0
	Qt_UpArrowCursor Qt_CursorShape = 1
	Qt_CrossCursor Qt_CursorShape = 2
	Qt_WaitCursor Qt_CursorShape = 3
	Qt_IBeamCursor Qt_CursorShape = 4
	Qt_SizeVerCursor Qt_CursorShape = 5
	Qt_SizeHorCursor Qt_CursorShape = 6
	Qt_SizeBDiagCursor Qt_CursorShape = 7
	Qt_SizeFDiagCursor Qt_CursorShape = 8
	Qt_SizeAllCursor Qt_CursorShape = 9
	Qt_BlankCursor Qt_CursorShape = 10
	Qt_SplitVCursor Qt_CursorShape = 11
	Qt_SplitHCursor Qt_CursorShape = 12
	Qt_PointingHandCursor Qt_CursorShape = 13
	Qt_ForbiddenCursor Qt_CursorShape = 14
	Qt_WhatsThisCursor Qt_CursorShape = 15
	Qt_BusyCursor Qt_CursorShape = 16
	Qt_OpenHandCursor Qt_CursorShape = 17
	Qt_ClosedHandCursor Qt_CursorShape = 18
	Qt_DragCopyCursor Qt_CursorShape = 19
	Qt_DragMoveCursor Qt_CursorShape = 20
	Qt_DragLinkCursor Qt_CursorShape = 21
	Qt_LastCursor Qt_CursorShape = Qt_DragLinkCursor
	Qt_BitmapCursor Qt_CursorShape = 24
	Qt_CustomCursor Qt_CursorShape = 25
)

//enum Qt_ItemSelectionMode - Qt::ItemSelectionMode
type Qt_ItemSelectionMode uint32
const (
	Qt_ContainsItemShape Qt_ItemSelectionMode = 0x0
	Qt_IntersectsItemShape Qt_ItemSelectionMode = 0x1
	Qt_ContainsItemBoundingRect Qt_ItemSelectionMode = 0x2
	Qt_IntersectsItemBoundingRect Qt_ItemSelectionMode = 0x3
)

//enum Qt_Orientation - Qt::Orientation
type Qt_Orientation uint32
const (
	Qt_Horizontal Qt_Orientation = 0x1
	Qt_Vertical Qt_Orientation = 0x2
)

//enum Qt_InputMethodQuery - Qt::InputMethodQuery
type Qt_InputMethodQuery uint32
const (
	Qt_ImMicroFocus Qt_InputMethodQuery = 0
	Qt_ImFont Qt_InputMethodQuery = 1
	Qt_ImCursorPosition Qt_InputMethodQuery = 2
	Qt_ImSurroundingText Qt_InputMethodQuery = 3
	Qt_ImCurrentSelection Qt_InputMethodQuery = 4
	Qt_ImMaximumTextLength Qt_InputMethodQuery = 5
	Qt_ImAnchorPosition Qt_InputMethodQuery = 6
)

//enum Qt_TimeSpec - Qt::TimeSpec
type Qt_TimeSpec uint32
const (
	Qt_LocalTime Qt_TimeSpec = 0
	Qt_UTC Qt_TimeSpec = 1
	Qt_OffsetFromUTC Qt_TimeSpec = 2
)

//enum Qt_ArrowType - Qt::ArrowType
type Qt_ArrowType uint32
const (
	Qt_NoArrow Qt_ArrowType = 0
	Qt_UpArrow Qt_ArrowType = 1
	Qt_DownArrow Qt_ArrowType = 2
	Qt_LeftArrow Qt_ArrowType = 3
	Qt_RightArrow Qt_ArrowType = 4
)

//enum Qt_FillRule - Qt::FillRule
type Qt_FillRule uint32
const (
	Qt_OddEvenFill Qt_FillRule = 0
	Qt_WindingFill Qt_FillRule = 1
)

//enum Qt_MaskMode - Qt::MaskMode
type Qt_MaskMode uint32
const (
	Qt_MaskInColor Qt_MaskMode = 0
	Qt_MaskOutColor Qt_MaskMode = 1
)

//enum Qt_WindowState - Qt::WindowState
type Qt_WindowState uint32
const (
	Qt_WindowNoState Qt_WindowState = 0x00000000
	Qt_WindowMinimized Qt_WindowState = 0x00000001
	Qt_WindowMaximized Qt_WindowState = 0x00000002
	Qt_WindowFullScreen Qt_WindowState = 0x00000004
	Qt_WindowActive Qt_WindowState = 0x00000008
)

//enum Qt_ToolBarAreaSizes - Qt::ToolBarAreaSizes
type Qt_ToolBarAreaSizes uint32
const (
	Qt_NToolBarAreas Qt_ToolBarAreaSizes = 4
)

//enum Qt_Corner - Qt::Corner
type Qt_Corner uint32
const (
	Qt_TopLeftCorner Qt_Corner = 0x00000
	Qt_TopRightCorner Qt_Corner = 0x00001
	Qt_BottomLeftCorner Qt_Corner = 0x00002
	Qt_BottomRightCorner Qt_Corner = 0x00003
)

//enum Qt_DayOfWeek - Qt::DayOfWeek
type Qt_DayOfWeek uint32
const (
	Qt_Monday Qt_DayOfWeek = 1
	Qt_Tuesday Qt_DayOfWeek = 2
	Qt_Wednesday Qt_DayOfWeek = 3
	Qt_Thursday Qt_DayOfWeek = 4
	Qt_Friday Qt_DayOfWeek = 5
	Qt_Saturday Qt_DayOfWeek = 6
	Qt_Sunday Qt_DayOfWeek = 7
)

//enum Qt_ClipOperation - Qt::ClipOperation
type Qt_ClipOperation uint32
const (
	Qt_NoClip Qt_ClipOperation = 0
	Qt_ReplaceClip Qt_ClipOperation = 1
	Qt_IntersectClip Qt_ClipOperation = 2
	Qt_UniteClip Qt_ClipOperation = 3
)

//enum Qt_LayoutDirection - Qt::LayoutDirection
type Qt_LayoutDirection uint32
const (
	Qt_LeftToRight Qt_LayoutDirection = 0
	Qt_RightToLeft Qt_LayoutDirection = 1
	Qt_LayoutDirectionAuto Qt_LayoutDirection = 2
)

//enum Qt_ToolButtonStyle - Qt::ToolButtonStyle
type Qt_ToolButtonStyle uint32
const (
	Qt_ToolButtonIconOnly Qt_ToolButtonStyle = 0
	Qt_ToolButtonTextOnly Qt_ToolButtonStyle = 1
	Qt_ToolButtonTextBesideIcon Qt_ToolButtonStyle = 2
	Qt_ToolButtonTextUnderIcon Qt_ToolButtonStyle = 3
	Qt_ToolButtonFollowStyle Qt_ToolButtonStyle = 4
)

//enum Qt_DockWidgetAreaSizes - Qt::DockWidgetAreaSizes
type Qt_DockWidgetAreaSizes uint32
const (
	Qt_NDockWidgetAreas Qt_DockWidgetAreaSizes = 4
)

//enum Qt_Key - Qt::Key
type Qt_Key uint32
const (
	Qt_Key_Escape Qt_Key = 0x01000000
	Qt_Key_Tab Qt_Key = 0x01000001
	Qt_Key_Backtab Qt_Key = 0x01000002
	Qt_Key_Backspace Qt_Key = 0x01000003
	Qt_Key_Return Qt_Key = 0x01000004
	Qt_Key_Enter Qt_Key = 0x01000005
	Qt_Key_Insert Qt_Key = 0x01000006
	Qt_Key_Delete Qt_Key = 0x01000007
	Qt_Key_Pause Qt_Key = 0x01000008
	Qt_Key_Print Qt_Key = 0x01000009
	Qt_Key_SysReq Qt_Key = 0x0100000a
	Qt_Key_Clear Qt_Key = 0x0100000b
	Qt_Key_Home Qt_Key = 0x01000010
	Qt_Key_End Qt_Key = 0x01000011
	Qt_Key_Left Qt_Key = 0x01000012
	Qt_Key_Up Qt_Key = 0x01000013
	Qt_Key_Right Qt_Key = 0x01000014
	Qt_Key_Down Qt_Key = 0x01000015
	Qt_Key_PageUp Qt_Key = 0x01000016
	Qt_Key_PageDown Qt_Key = 0x01000017
	Qt_Key_Shift Qt_Key = 0x01000020
	Qt_Key_Control Qt_Key = 0x01000021
	Qt_Key_Meta Qt_Key = 0x01000022
	Qt_Key_Alt Qt_Key = 0x01000023
	Qt_Key_CapsLock Qt_Key = 0x01000024
	Qt_Key_NumLock Qt_Key = 0x01000025
	Qt_Key_ScrollLock Qt_Key = 0x01000026
	Qt_Key_F1 Qt_Key = 0x01000030
	Qt_Key_F2 Qt_Key = 0x01000031
	Qt_Key_F3 Qt_Key = 0x01000032
	Qt_Key_F4 Qt_Key = 0x01000033
	Qt_Key_F5 Qt_Key = 0x01000034
	Qt_Key_F6 Qt_Key = 0x01000035
	Qt_Key_F7 Qt_Key = 0x01000036
	Qt_Key_F8 Qt_Key = 0x01000037
	Qt_Key_F9 Qt_Key = 0x01000038
	Qt_Key_F10 Qt_Key = 0x01000039
	Qt_Key_F11 Qt_Key = 0x0100003a
	Qt_Key_F12 Qt_Key = 0x0100003b
	Qt_Key_F13 Qt_Key = 0x0100003c
	Qt_Key_F14 Qt_Key = 0x0100003d
	Qt_Key_F15 Qt_Key = 0x0100003e
	Qt_Key_F16 Qt_Key = 0x0100003f
	Qt_Key_F17 Qt_Key = 0x01000040
	Qt_Key_F18 Qt_Key = 0x01000041
	Qt_Key_F19 Qt_Key = 0x01000042
	Qt_Key_F20 Qt_Key = 0x01000043
	Qt_Key_F21 Qt_Key = 0x01000044
	Qt_Key_F22 Qt_Key = 0x01000045
	Qt_Key_F23 Qt_Key = 0x01000046
	Qt_Key_F24 Qt_Key = 0x01000047
	Qt_Key_F25 Qt_Key = 0x01000048
	Qt_Key_F26 Qt_Key = 0x01000049
	Qt_Key_F27 Qt_Key = 0x0100004a
	Qt_Key_F28 Qt_Key = 0x0100004b
	Qt_Key_F29 Qt_Key = 0x0100004c
	Qt_Key_F30 Qt_Key = 0x0100004d
	Qt_Key_F31 Qt_Key = 0x0100004e
	Qt_Key_F32 Qt_Key = 0x0100004f
	Qt_Key_F33 Qt_Key = 0x01000050
	Qt_Key_F34 Qt_Key = 0x01000051
	Qt_Key_F35 Qt_Key = 0x01000052
	Qt_Key_Super_L Qt_Key = 0x01000053
	Qt_Key_Super_R Qt_Key = 0x01000054
	Qt_Key_Menu Qt_Key = 0x01000055
	Qt_Key_Hyper_L Qt_Key = 0x01000056
	Qt_Key_Hyper_R Qt_Key = 0x01000057
	Qt_Key_Help Qt_Key = 0x01000058
	Qt_Key_Direction_L Qt_Key = 0x01000059
	Qt_Key_Direction_R Qt_Key = 0x01000060
	Qt_Key_Space Qt_Key = 0x20
	Qt_Key_Any Qt_Key = Qt_Key_Space
	Qt_Key_Exclam Qt_Key = 0x21
	Qt_Key_QuoteDbl Qt_Key = 0x22
	Qt_Key_NumberSign Qt_Key = 0x23
	Qt_Key_Dollar Qt_Key = 0x24
	Qt_Key_Percent Qt_Key = 0x25
	Qt_Key_Ampersand Qt_Key = 0x26
	Qt_Key_Apostrophe Qt_Key = 0x27
	Qt_Key_ParenLeft Qt_Key = 0x28
	Qt_Key_ParenRight Qt_Key = 0x29
	Qt_Key_Asterisk Qt_Key = 0x2a
	Qt_Key_Plus Qt_Key = 0x2b
	Qt_Key_Comma Qt_Key = 0x2c
	Qt_Key_Minus Qt_Key = 0x2d
	Qt_Key_Period Qt_Key = 0x2e
	Qt_Key_Slash Qt_Key = 0x2f
	Qt_Key_0 Qt_Key = 0x30
	Qt_Key_1 Qt_Key = 0x31
	Qt_Key_2 Qt_Key = 0x32
	Qt_Key_3 Qt_Key = 0x33
	Qt_Key_4 Qt_Key = 0x34
	Qt_Key_5 Qt_Key = 0x35
	Qt_Key_6 Qt_Key = 0x36
	Qt_Key_7 Qt_Key = 0x37
	Qt_Key_8 Qt_Key = 0x38
	Qt_Key_9 Qt_Key = 0x39
	Qt_Key_Colon Qt_Key = 0x3a
	Qt_Key_Semicolon Qt_Key = 0x3b
	Qt_Key_Less Qt_Key = 0x3c
	Qt_Key_Equal Qt_Key = 0x3d
	Qt_Key_Greater Qt_Key = 0x3e
	Qt_Key_Question Qt_Key = 0x3f
	Qt_Key_At Qt_Key = 0x40
	Qt_Key_A Qt_Key = 0x41
	Qt_Key_B Qt_Key = 0x42
	Qt_Key_C Qt_Key = 0x43
	Qt_Key_D Qt_Key = 0x44
	Qt_Key_E Qt_Key = 0x45
	Qt_Key_F Qt_Key = 0x46
	Qt_Key_G Qt_Key = 0x47
	Qt_Key_H Qt_Key = 0x48
	Qt_Key_I Qt_Key = 0x49
	Qt_Key_J Qt_Key = 0x4a
	Qt_Key_K Qt_Key = 0x4b
	Qt_Key_L Qt_Key = 0x4c
	Qt_Key_M Qt_Key = 0x4d
	Qt_Key_N Qt_Key = 0x4e
	Qt_Key_O Qt_Key = 0x4f
	Qt_Key_P Qt_Key = 0x50
	Qt_Key_Q Qt_Key = 0x51
	Qt_Key_R Qt_Key = 0x52
	Qt_Key_S Qt_Key = 0x53
	Qt_Key_T Qt_Key = 0x54
	Qt_Key_U Qt_Key = 0x55
	Qt_Key_V Qt_Key = 0x56
	Qt_Key_W Qt_Key = 0x57
	Qt_Key_X Qt_Key = 0x58
	Qt_Key_Y Qt_Key = 0x59
	Qt_Key_Z Qt_Key = 0x5a
	Qt_Key_BracketLeft Qt_Key = 0x5b
	Qt_Key_Backslash Qt_Key = 0x5c
	Qt_Key_BracketRight Qt_Key = 0x5d
	Qt_Key_AsciiCircum Qt_Key = 0x5e
	Qt_Key_Underscore Qt_Key = 0x5f
	Qt_Key_QuoteLeft Qt_Key = 0x60
	Qt_Key_BraceLeft Qt_Key = 0x7b
	Qt_Key_Bar Qt_Key = 0x7c
	Qt_Key_BraceRight Qt_Key = 0x7d
	Qt_Key_AsciiTilde Qt_Key = 0x7e
	Qt_Key_nobreakspace Qt_Key = 0x0a0
	Qt_Key_exclamdown Qt_Key = 0x0a1
	Qt_Key_cent Qt_Key = 0x0a2
	Qt_Key_sterling Qt_Key = 0x0a3
	Qt_Key_currency Qt_Key = 0x0a4
	Qt_Key_yen Qt_Key = 0x0a5
	Qt_Key_brokenbar Qt_Key = 0x0a6
	Qt_Key_section Qt_Key = 0x0a7
	Qt_Key_diaeresis Qt_Key = 0x0a8
	Qt_Key_copyright Qt_Key = 0x0a9
	Qt_Key_ordfeminine Qt_Key = 0x0aa
	Qt_Key_guillemotleft Qt_Key = 0x0ab
	Qt_Key_notsign Qt_Key = 0x0ac
	Qt_Key_hyphen Qt_Key = 0x0ad
	Qt_Key_registered Qt_Key = 0x0ae
	Qt_Key_macron Qt_Key = 0x0af
	Qt_Key_degree Qt_Key = 0x0b0
	Qt_Key_plusminus Qt_Key = 0x0b1
	Qt_Key_twosuperior Qt_Key = 0x0b2
	Qt_Key_threesuperior Qt_Key = 0x0b3
	Qt_Key_acute Qt_Key = 0x0b4
	Qt_Key_mu Qt_Key = 0x0b5
	Qt_Key_paragraph Qt_Key = 0x0b6
	Qt_Key_periodcentered Qt_Key = 0x0b7
	Qt_Key_cedilla Qt_Key = 0x0b8
	Qt_Key_onesuperior Qt_Key = 0x0b9
	Qt_Key_masculine Qt_Key = 0x0ba
	Qt_Key_guillemotright Qt_Key = 0x0bb
	Qt_Key_onequarter Qt_Key = 0x0bc
	Qt_Key_onehalf Qt_Key = 0x0bd
	Qt_Key_threequarters Qt_Key = 0x0be
	Qt_Key_questiondown Qt_Key = 0x0bf
	Qt_Key_Agrave Qt_Key = 0x0c0
	Qt_Key_Aacute Qt_Key = 0x0c1
	Qt_Key_Acircumflex Qt_Key = 0x0c2
	Qt_Key_Atilde Qt_Key = 0x0c3
	Qt_Key_Adiaeresis Qt_Key = 0x0c4
	Qt_Key_Aring Qt_Key = 0x0c5
	Qt_Key_AE Qt_Key = 0x0c6
	Qt_Key_Ccedilla Qt_Key = 0x0c7
	Qt_Key_Egrave Qt_Key = 0x0c8
	Qt_Key_Eacute Qt_Key = 0x0c9
	Qt_Key_Ecircumflex Qt_Key = 0x0ca
	Qt_Key_Ediaeresis Qt_Key = 0x0cb
	Qt_Key_Igrave Qt_Key = 0x0cc
	Qt_Key_Iacute Qt_Key = 0x0cd
	Qt_Key_Icircumflex Qt_Key = 0x0ce
	Qt_Key_Idiaeresis Qt_Key = 0x0cf
	Qt_Key_ETH Qt_Key = 0x0d0
	Qt_Key_Ntilde Qt_Key = 0x0d1
	Qt_Key_Ograve Qt_Key = 0x0d2
	Qt_Key_Oacute Qt_Key = 0x0d3
	Qt_Key_Ocircumflex Qt_Key = 0x0d4
	Qt_Key_Otilde Qt_Key = 0x0d5
	Qt_Key_Odiaeresis Qt_Key = 0x0d6
	Qt_Key_multiply Qt_Key = 0x0d7
	Qt_Key_Ooblique Qt_Key = 0x0d8
	Qt_Key_Ugrave Qt_Key = 0x0d9
	Qt_Key_Uacute Qt_Key = 0x0da
	Qt_Key_Ucircumflex Qt_Key = 0x0db
	Qt_Key_Udiaeresis Qt_Key = 0x0dc
	Qt_Key_Yacute Qt_Key = 0x0dd
	Qt_Key_THORN Qt_Key = 0x0de
	Qt_Key_ssharp Qt_Key = 0x0df
	Qt_Key_division Qt_Key = 0x0f7
	Qt_Key_ydiaeresis Qt_Key = 0x0ff
	Qt_Key_AltGr Qt_Key = 0x01001103
	Qt_Key_Multi_key Qt_Key = 0x01001120
	Qt_Key_Codeinput Qt_Key = 0x01001137
	Qt_Key_SingleCandidate Qt_Key = 0x0100113c
	Qt_Key_MultipleCandidate Qt_Key = 0x0100113d
	Qt_Key_PreviousCandidate Qt_Key = 0x0100113e
	Qt_Key_Mode_switch Qt_Key = 0x0100117e
	Qt_Key_Kanji Qt_Key = 0x01001121
	Qt_Key_Muhenkan Qt_Key = 0x01001122
	Qt_Key_Henkan Qt_Key = 0x01001123
	Qt_Key_Romaji Qt_Key = 0x01001124
	Qt_Key_Hiragana Qt_Key = 0x01001125
	Qt_Key_Katakana Qt_Key = 0x01001126
	Qt_Key_Hiragana_Katakana Qt_Key = 0x01001127
	Qt_Key_Zenkaku Qt_Key = 0x01001128
	Qt_Key_Hankaku Qt_Key = 0x01001129
	Qt_Key_Zenkaku_Hankaku Qt_Key = 0x0100112a
	Qt_Key_Touroku Qt_Key = 0x0100112b
	Qt_Key_Massyo Qt_Key = 0x0100112c
	Qt_Key_Kana_Lock Qt_Key = 0x0100112d
	Qt_Key_Kana_Shift Qt_Key = 0x0100112e
	Qt_Key_Eisu_Shift Qt_Key = 0x0100112f
	Qt_Key_Eisu_toggle Qt_Key = 0x01001130
	Qt_Key_Hangul Qt_Key = 0x01001131
	Qt_Key_Hangul_Start Qt_Key = 0x01001132
	Qt_Key_Hangul_End Qt_Key = 0x01001133
	Qt_Key_Hangul_Hanja Qt_Key = 0x01001134
	Qt_Key_Hangul_Jamo Qt_Key = 0x01001135
	Qt_Key_Hangul_Romaja Qt_Key = 0x01001136
	Qt_Key_Hangul_Jeonja Qt_Key = 0x01001138
	Qt_Key_Hangul_Banja Qt_Key = 0x01001139
	Qt_Key_Hangul_PreHanja Qt_Key = 0x0100113a
	Qt_Key_Hangul_PostHanja Qt_Key = 0x0100113b
	Qt_Key_Hangul_Special Qt_Key = 0x0100113f
	Qt_Key_Dead_Grave Qt_Key = 0x01001250
	Qt_Key_Dead_Acute Qt_Key = 0x01001251
	Qt_Key_Dead_Circumflex Qt_Key = 0x01001252
	Qt_Key_Dead_Tilde Qt_Key = 0x01001253
	Qt_Key_Dead_Macron Qt_Key = 0x01001254
	Qt_Key_Dead_Breve Qt_Key = 0x01001255
	Qt_Key_Dead_Abovedot Qt_Key = 0x01001256
	Qt_Key_Dead_Diaeresis Qt_Key = 0x01001257
	Qt_Key_Dead_Abovering Qt_Key = 0x01001258
	Qt_Key_Dead_Doubleacute Qt_Key = 0x01001259
	Qt_Key_Dead_Caron Qt_Key = 0x0100125a
	Qt_Key_Dead_Cedilla Qt_Key = 0x0100125b
	Qt_Key_Dead_Ogonek Qt_Key = 0x0100125c
	Qt_Key_Dead_Iota Qt_Key = 0x0100125d
	Qt_Key_Dead_Voiced_Sound Qt_Key = 0x0100125e
	Qt_Key_Dead_Semivoiced_Sound Qt_Key = 0x0100125f
	Qt_Key_Dead_Belowdot Qt_Key = 0x01001260
	Qt_Key_Dead_Hook Qt_Key = 0x01001261
	Qt_Key_Dead_Horn Qt_Key = 0x01001262
	Qt_Key_Back Qt_Key = 0x01000061
	Qt_Key_Forward Qt_Key = 0x01000062
	Qt_Key_Stop Qt_Key = 0x01000063
	Qt_Key_Refresh Qt_Key = 0x01000064
	Qt_Key_VolumeDown Qt_Key = 0x01000070
	Qt_Key_VolumeMute Qt_Key = 0x01000071
	Qt_Key_VolumeUp Qt_Key = 0x01000072
	Qt_Key_BassBoost Qt_Key = 0x01000073
	Qt_Key_BassUp Qt_Key = 0x01000074
	Qt_Key_BassDown Qt_Key = 0x01000075
	Qt_Key_TrebleUp Qt_Key = 0x01000076
	Qt_Key_TrebleDown Qt_Key = 0x01000077
	Qt_Key_MediaPlay Qt_Key = 0x01000080
	Qt_Key_MediaStop Qt_Key = 0x01000081
	Qt_Key_MediaPrevious Qt_Key = 0x01000082
	Qt_Key_MediaNext Qt_Key = 0x01000083
	Qt_Key_MediaRecord Qt_Key = 0x01000084
	Qt_Key_MediaPause Qt_Key = 0x1000085
	Qt_Key_MediaTogglePlayPause Qt_Key = 0x1000086
	Qt_Key_HomePage Qt_Key = 0x01000090
	Qt_Key_Favorites Qt_Key = 0x01000091
	Qt_Key_Search Qt_Key = 0x01000092
	Qt_Key_Standby Qt_Key = 0x01000093
	Qt_Key_OpenUrl Qt_Key = 0x01000094
	Qt_Key_LaunchMail Qt_Key = 0x010000a0
	Qt_Key_LaunchMedia Qt_Key = 0x010000a1
	Qt_Key_Launch0 Qt_Key = 0x010000a2
	Qt_Key_Launch1 Qt_Key = 0x010000a3
	Qt_Key_Launch2 Qt_Key = 0x010000a4
	Qt_Key_Launch3 Qt_Key = 0x010000a5
	Qt_Key_Launch4 Qt_Key = 0x010000a6
	Qt_Key_Launch5 Qt_Key = 0x010000a7
	Qt_Key_Launch6 Qt_Key = 0x010000a8
	Qt_Key_Launch7 Qt_Key = 0x010000a9
	Qt_Key_Launch8 Qt_Key = 0x010000aa
	Qt_Key_Launch9 Qt_Key = 0x010000ab
	Qt_Key_LaunchA Qt_Key = 0x010000ac
	Qt_Key_LaunchB Qt_Key = 0x010000ad
	Qt_Key_LaunchC Qt_Key = 0x010000ae
	Qt_Key_LaunchD Qt_Key = 0x010000af
	Qt_Key_LaunchE Qt_Key = 0x010000b0
	Qt_Key_LaunchF Qt_Key = 0x010000b1
	Qt_Key_MonBrightnessUp Qt_Key = 0x010000b2
	Qt_Key_MonBrightnessDown Qt_Key = 0x010000b3
	Qt_Key_KeyboardLightOnOff Qt_Key = 0x010000b4
	Qt_Key_KeyboardBrightnessUp Qt_Key = 0x010000b5
	Qt_Key_KeyboardBrightnessDown Qt_Key = 0x010000b6
	Qt_Key_PowerOff Qt_Key = 0x010000b7
	Qt_Key_WakeUp Qt_Key = 0x010000b8
	Qt_Key_Eject Qt_Key = 0x010000b9
	Qt_Key_ScreenSaver Qt_Key = 0x010000ba
	Qt_Key_WWW Qt_Key = 0x010000bb
	Qt_Key_Memo Qt_Key = 0x010000bc
	Qt_Key_LightBulb Qt_Key = 0x010000bd
	Qt_Key_Shop Qt_Key = 0x010000be
	Qt_Key_History Qt_Key = 0x010000bf
	Qt_Key_AddFavorite Qt_Key = 0x010000c0
	Qt_Key_HotLinks Qt_Key = 0x010000c1
	Qt_Key_BrightnessAdjust Qt_Key = 0x010000c2
	Qt_Key_Finance Qt_Key = 0x010000c3
	Qt_Key_Community Qt_Key = 0x010000c4
	Qt_Key_AudioRewind Qt_Key = 0x010000c5
	Qt_Key_BackForward Qt_Key = 0x010000c6
	Qt_Key_ApplicationLeft Qt_Key = 0x010000c7
	Qt_Key_ApplicationRight Qt_Key = 0x010000c8
	Qt_Key_Book Qt_Key = 0x010000c9
	Qt_Key_CD Qt_Key = 0x010000ca
	Qt_Key_Calculator Qt_Key = 0x010000cb
	Qt_Key_ToDoList Qt_Key = 0x010000cc
	Qt_Key_ClearGrab Qt_Key = 0x010000cd
	Qt_Key_Close Qt_Key = 0x010000ce
	Qt_Key_Copy Qt_Key = 0x010000cf
	Qt_Key_Cut Qt_Key = 0x010000d0
	Qt_Key_Display Qt_Key = 0x010000d1
	Qt_Key_DOS Qt_Key = 0x010000d2
	Qt_Key_Documents Qt_Key = 0x010000d3
	Qt_Key_Excel Qt_Key = 0x010000d4
	Qt_Key_Explorer Qt_Key = 0x010000d5
	Qt_Key_Game Qt_Key = 0x010000d6
	Qt_Key_Go Qt_Key = 0x010000d7
	Qt_Key_iTouch Qt_Key = 0x010000d8
	Qt_Key_LogOff Qt_Key = 0x010000d9
	Qt_Key_Market Qt_Key = 0x010000da
	Qt_Key_Meeting Qt_Key = 0x010000db
	Qt_Key_MenuKB Qt_Key = 0x010000dc
	Qt_Key_MenuPB Qt_Key = 0x010000dd
	Qt_Key_MySites Qt_Key = 0x010000de
	Qt_Key_News Qt_Key = 0x010000df
	Qt_Key_OfficeHome Qt_Key = 0x010000e0
	Qt_Key_Option Qt_Key = 0x010000e1
	Qt_Key_Paste Qt_Key = 0x010000e2
	Qt_Key_Phone Qt_Key = 0x010000e3
	Qt_Key_Calendar Qt_Key = 0x010000e4
	Qt_Key_Reply Qt_Key = 0x010000e5
	Qt_Key_Reload Qt_Key = 0x010000e6
	Qt_Key_RotateWindows Qt_Key = 0x010000e7
	Qt_Key_RotationPB Qt_Key = 0x010000e8
	Qt_Key_RotationKB Qt_Key = 0x010000e9
	Qt_Key_Save Qt_Key = 0x010000ea
	Qt_Key_Send Qt_Key = 0x010000eb
	Qt_Key_Spell Qt_Key = 0x010000ec
	Qt_Key_SplitScreen Qt_Key = 0x010000ed
	Qt_Key_Support Qt_Key = 0x010000ee
	Qt_Key_TaskPane Qt_Key = 0x010000ef
	Qt_Key_Terminal Qt_Key = 0x010000f0
	Qt_Key_Tools Qt_Key = 0x010000f1
	Qt_Key_Travel Qt_Key = 0x010000f2
	Qt_Key_Video Qt_Key = 0x010000f3
	Qt_Key_Word Qt_Key = 0x010000f4
	Qt_Key_Xfer Qt_Key = 0x010000f5
	Qt_Key_ZoomIn Qt_Key = 0x010000f6
	Qt_Key_ZoomOut Qt_Key = 0x010000f7
	Qt_Key_Away Qt_Key = 0x010000f8
	Qt_Key_Messenger Qt_Key = 0x010000f9
	Qt_Key_WebCam Qt_Key = 0x010000fa
	Qt_Key_MailForward Qt_Key = 0x010000fb
	Qt_Key_Pictures Qt_Key = 0x010000fc
	Qt_Key_Music Qt_Key = 0x010000fd
	Qt_Key_Battery Qt_Key = 0x010000fe
	Qt_Key_Bluetooth Qt_Key = 0x010000ff
	Qt_Key_WLAN Qt_Key = 0x01000100
	Qt_Key_UWB Qt_Key = 0x01000101
	Qt_Key_AudioForward Qt_Key = 0x01000102
	Qt_Key_AudioRepeat Qt_Key = 0x01000103
	Qt_Key_AudioRandomPlay Qt_Key = 0x01000104
	Qt_Key_Subtitle Qt_Key = 0x01000105
	Qt_Key_AudioCycleTrack Qt_Key = 0x01000106
	Qt_Key_Time Qt_Key = 0x01000107
	Qt_Key_Hibernate Qt_Key = 0x01000108
	Qt_Key_View Qt_Key = 0x01000109
	Qt_Key_TopMenu Qt_Key = 0x0100010a
	Qt_Key_PowerDown Qt_Key = 0x0100010b
	Qt_Key_Suspend Qt_Key = 0x0100010c
	Qt_Key_ContrastAdjust Qt_Key = 0x0100010d
	Qt_Key_LaunchG Qt_Key = 0x0100010e
	Qt_Key_LaunchH Qt_Key = 0x0100010f
	Qt_Key_MediaLast Qt_Key = 0x0100ffff
	Qt_Key_Select Qt_Key = 0x01010000
	Qt_Key_Yes Qt_Key = 0x01010001
	Qt_Key_No Qt_Key = 0x01010002
	Qt_Key_Cancel Qt_Key = 0x01020001
	Qt_Key_Printer Qt_Key = 0x01020002
	Qt_Key_Execute Qt_Key = 0x01020003
	Qt_Key_Sleep Qt_Key = 0x01020004
	Qt_Key_Play Qt_Key = 0x01020005
	Qt_Key_Zoom Qt_Key = 0x01020006
	Qt_Key_Context1 Qt_Key = 0x01100000
	Qt_Key_Context2 Qt_Key = 0x01100001
	Qt_Key_Context3 Qt_Key = 0x01100002
	Qt_Key_Context4 Qt_Key = 0x01100003
	Qt_Key_Call Qt_Key = 0x01100004
	Qt_Key_Hangup Qt_Key = 0x01100005
	Qt_Key_Flip Qt_Key = 0x01100006
	Qt_Key_ToggleCallHangup Qt_Key = 0x01100007
	Qt_Key_VoiceDial Qt_Key = 0x01100008
	Qt_Key_LastNumberRedial Qt_Key = 0x01100009
	Qt_Key_Camera Qt_Key = 0x01100020
	Qt_Key_CameraFocus Qt_Key = 0x01100021
	Qt_Key_unknown Qt_Key = 0x01ffffff
)

//enum Qt_ApplicationAttribute - Qt::ApplicationAttribute
type Qt_ApplicationAttribute uint32
const (
	Qt_AA_ImmediateWidgetCreation Qt_ApplicationAttribute = 0
	Qt_AA_MSWindowsUseDirect3DByDefault Qt_ApplicationAttribute = 1
	Qt_AA_DontShowIconsInMenus Qt_ApplicationAttribute = 2
	Qt_AA_NativeWindows Qt_ApplicationAttribute = 3
	Qt_AA_DontCreateNativeWidgetSiblings Qt_ApplicationAttribute = 4
	Qt_AA_MacPluginApplication Qt_ApplicationAttribute = 5
	Qt_AA_DontUseNativeMenuBar Qt_ApplicationAttribute = 6
	Qt_AA_MacDontSwapCtrlAndMeta Qt_ApplicationAttribute = 7
	Qt_AA_S60DontConstructApplicationPanes Qt_ApplicationAttribute = 8
	Qt_AA_S60DisablePartialScreenInputMode Qt_ApplicationAttribute = 9
	Qt_AA_AttributeCount Qt_ApplicationAttribute = 9+1
)

//enum Qt_TextFlag - Qt::TextFlag
type Qt_TextFlag uint32
const (
	Qt_TextSingleLine Qt_TextFlag = 0x0100
	Qt_TextDontClip Qt_TextFlag = 0x0200
	Qt_TextExpandTabs Qt_TextFlag = 0x0400
	Qt_TextShowMnemonic Qt_TextFlag = 0x0800
	Qt_TextWordWrap Qt_TextFlag = 0x1000
	Qt_TextWrapAnywhere Qt_TextFlag = 0x2000
	Qt_TextDontPrint Qt_TextFlag = 0x4000
	Qt_TextIncludeTrailingSpaces Qt_TextFlag = 0x08000000
	Qt_TextHideMnemonic Qt_TextFlag = 0x8000
	Qt_TextJustificationForced Qt_TextFlag = 0x10000
	Qt_TextForceLeftToRight Qt_TextFlag = 0x20000
	Qt_TextForceRightToLeft Qt_TextFlag = 0x40000
	Qt_TextLongestVariant Qt_TextFlag = 0x80000
	Qt_TextBypassShaping Qt_TextFlag = 0x100000
)

//enum Qt_BrushStyle - Qt::BrushStyle
type Qt_BrushStyle uint32
const (
	Qt_NoBrush Qt_BrushStyle = 0
	Qt_SolidPattern Qt_BrushStyle = 1
	Qt_Dense1Pattern Qt_BrushStyle = 2
	Qt_Dense2Pattern Qt_BrushStyle = 3
	Qt_Dense3Pattern Qt_BrushStyle = 4
	Qt_Dense4Pattern Qt_BrushStyle = 5
	Qt_Dense5Pattern Qt_BrushStyle = 6
	Qt_Dense6Pattern Qt_BrushStyle = 7
	Qt_Dense7Pattern Qt_BrushStyle = 8
	Qt_HorPattern Qt_BrushStyle = 9
	Qt_VerPattern Qt_BrushStyle = 10
	Qt_CrossPattern Qt_BrushStyle = 11
	Qt_BDiagPattern Qt_BrushStyle = 12
	Qt_FDiagPattern Qt_BrushStyle = 13
	Qt_DiagCrossPattern Qt_BrushStyle = 14
	Qt_LinearGradientPattern Qt_BrushStyle = 15
	Qt_RadialGradientPattern Qt_BrushStyle = 16
	Qt_ConicalGradientPattern Qt_BrushStyle = 17
	Qt_TexturePattern Qt_BrushStyle = 24
)

//enum Qt_WidgetAttribute - Qt::WidgetAttribute
type Qt_WidgetAttribute uint32
const (
	Qt_WA_Disabled Qt_WidgetAttribute = 0
	Qt_WA_UnderMouse Qt_WidgetAttribute = 1
	Qt_WA_MouseTracking Qt_WidgetAttribute = 2
	Qt_WA_ContentsPropagated Qt_WidgetAttribute = 3
	Qt_WA_OpaquePaintEvent Qt_WidgetAttribute = 4
	Qt_WA_NoBackground Qt_WidgetAttribute = Qt_WA_OpaquePaintEvent
	Qt_WA_StaticContents Qt_WidgetAttribute = 5
	Qt_WA_LaidOut Qt_WidgetAttribute = 7
	Qt_WA_PaintOnScreen Qt_WidgetAttribute = 8
	Qt_WA_NoSystemBackground Qt_WidgetAttribute = 9
	Qt_WA_UpdatesDisabled Qt_WidgetAttribute = 10
	Qt_WA_Mapped Qt_WidgetAttribute = 11
	Qt_WA_MacNoClickThrough Qt_WidgetAttribute = 12
	Qt_WA_PaintOutsidePaintEvent Qt_WidgetAttribute = 13
	Qt_WA_InputMethodEnabled Qt_WidgetAttribute = 14
	Qt_WA_WState_Visible Qt_WidgetAttribute = 15
	Qt_WA_WState_Hidden Qt_WidgetAttribute = 16
	Qt_WA_ForceDisabled Qt_WidgetAttribute = 32
	Qt_WA_KeyCompression Qt_WidgetAttribute = 33
	Qt_WA_PendingMoveEvent Qt_WidgetAttribute = 34
	Qt_WA_PendingResizeEvent Qt_WidgetAttribute = 35
	Qt_WA_SetPalette Qt_WidgetAttribute = 36
	Qt_WA_SetFont Qt_WidgetAttribute = 37
	Qt_WA_SetCursor Qt_WidgetAttribute = 38
	Qt_WA_NoChildEventsFromChildren Qt_WidgetAttribute = 39
	Qt_WA_WindowModified Qt_WidgetAttribute = 41
	Qt_WA_Resized Qt_WidgetAttribute = 42
	Qt_WA_Moved Qt_WidgetAttribute = 43
	Qt_WA_PendingUpdate Qt_WidgetAttribute = 44
	Qt_WA_InvalidSize Qt_WidgetAttribute = 45
	Qt_WA_MacBrushedMetal Qt_WidgetAttribute = 46
	Qt_WA_MacMetalStyle Qt_WidgetAttribute = Qt_WA_MacBrushedMetal
	Qt_WA_CustomWhatsThis Qt_WidgetAttribute = 47
	Qt_WA_LayoutOnEntireRect Qt_WidgetAttribute = 48
	Qt_WA_OutsideWSRange Qt_WidgetAttribute = 49
	Qt_WA_GrabbedShortcut Qt_WidgetAttribute = 50
	Qt_WA_TransparentForMouseEvents Qt_WidgetAttribute = 51
	Qt_WA_PaintUnclipped Qt_WidgetAttribute = 52
	Qt_WA_SetWindowIcon Qt_WidgetAttribute = 53
	Qt_WA_NoMouseReplay Qt_WidgetAttribute = 54
	Qt_WA_DeleteOnClose Qt_WidgetAttribute = 55
	Qt_WA_RightToLeft Qt_WidgetAttribute = 56
	Qt_WA_SetLayoutDirection Qt_WidgetAttribute = 57
	Qt_WA_NoChildEventsForParent Qt_WidgetAttribute = 58
	Qt_WA_ForceUpdatesDisabled Qt_WidgetAttribute = 59
	Qt_WA_WState_Created Qt_WidgetAttribute = 60
	Qt_WA_WState_CompressKeys Qt_WidgetAttribute = 61
	Qt_WA_WState_InPaintEvent Qt_WidgetAttribute = 62
	Qt_WA_WState_Reparented Qt_WidgetAttribute = 63
	Qt_WA_WState_ConfigPending Qt_WidgetAttribute = 64
	Qt_WA_WState_Polished Qt_WidgetAttribute = 66
	Qt_WA_WState_DND Qt_WidgetAttribute = 67
	Qt_WA_WState_OwnSizePolicy Qt_WidgetAttribute = 68
	Qt_WA_WState_ExplicitShowHide Qt_WidgetAttribute = 69
	Qt_WA_ShowModal Qt_WidgetAttribute = 70
	Qt_WA_MouseNoMask Qt_WidgetAttribute = 71
	Qt_WA_GroupLeader Qt_WidgetAttribute = 72
	Qt_WA_NoMousePropagation Qt_WidgetAttribute = 73
	Qt_WA_Hover Qt_WidgetAttribute = 74
	Qt_WA_InputMethodTransparent Qt_WidgetAttribute = 75
	Qt_WA_QuitOnClose Qt_WidgetAttribute = 76
	Qt_WA_KeyboardFocusChange Qt_WidgetAttribute = 77
	Qt_WA_AcceptDrops Qt_WidgetAttribute = 78
	Qt_WA_DropSiteRegistered Qt_WidgetAttribute = 79
	Qt_WA_ForceAcceptDrops Qt_WidgetAttribute = Qt_WA_DropSiteRegistered
	Qt_WA_WindowPropagation Qt_WidgetAttribute = 80
	Qt_WA_NoX11EventCompression Qt_WidgetAttribute = 81
	Qt_WA_TintedBackground Qt_WidgetAttribute = 82
	Qt_WA_X11OpenGLOverlay Qt_WidgetAttribute = 83
	Qt_WA_AlwaysShowToolTips Qt_WidgetAttribute = 84
	Qt_WA_MacOpaqueSizeGrip Qt_WidgetAttribute = 85
	Qt_WA_SetStyle Qt_WidgetAttribute = 86
	Qt_WA_SetLocale Qt_WidgetAttribute = 87
	Qt_WA_MacShowFocusRect Qt_WidgetAttribute = 88
	Qt_WA_MacNormalSize Qt_WidgetAttribute = 89
	Qt_WA_MacSmallSize Qt_WidgetAttribute = 90
	Qt_WA_MacMiniSize Qt_WidgetAttribute = 91
	Qt_WA_LayoutUsesWidgetRect Qt_WidgetAttribute = 92
	Qt_WA_StyledBackground Qt_WidgetAttribute = 93
	Qt_WA_MSWindowsUseDirect3D Qt_WidgetAttribute = 94
	Qt_WA_CanHostQMdiSubWindowTitleBar Qt_WidgetAttribute = 95
	Qt_WA_MacAlwaysShowToolWindow Qt_WidgetAttribute = 96
	Qt_WA_StyleSheet Qt_WidgetAttribute = 97
	Qt_WA_ShowWithoutActivating Qt_WidgetAttribute = 98
	Qt_WA_X11BypassTransientForHint Qt_WidgetAttribute = 99
	Qt_WA_NativeWindow Qt_WidgetAttribute = 100
	Qt_WA_DontCreateNativeAncestors Qt_WidgetAttribute = 101
	Qt_WA_MacVariableSize Qt_WidgetAttribute = 102
	Qt_WA_DontShowOnScreen Qt_WidgetAttribute = 103
	Qt_WA_X11NetWmWindowTypeDesktop Qt_WidgetAttribute = 104
	Qt_WA_X11NetWmWindowTypeDock Qt_WidgetAttribute = 105
	Qt_WA_X11NetWmWindowTypeToolBar Qt_WidgetAttribute = 106
	Qt_WA_X11NetWmWindowTypeMenu Qt_WidgetAttribute = 107
	Qt_WA_X11NetWmWindowTypeUtility Qt_WidgetAttribute = 108
	Qt_WA_X11NetWmWindowTypeSplash Qt_WidgetAttribute = 109
	Qt_WA_X11NetWmWindowTypeDialog Qt_WidgetAttribute = 110
	Qt_WA_X11NetWmWindowTypeDropDownMenu Qt_WidgetAttribute = 111
	Qt_WA_X11NetWmWindowTypePopupMenu Qt_WidgetAttribute = 112
	Qt_WA_X11NetWmWindowTypeToolTip Qt_WidgetAttribute = 113
	Qt_WA_X11NetWmWindowTypeNotification Qt_WidgetAttribute = 114
	Qt_WA_X11NetWmWindowTypeCombo Qt_WidgetAttribute = 115
	Qt_WA_X11NetWmWindowTypeDND Qt_WidgetAttribute = 116
	Qt_WA_MacFrameworkScaled Qt_WidgetAttribute = 117
	Qt_WA_SetWindowModality Qt_WidgetAttribute = 118
	Qt_WA_WState_WindowOpacitySet Qt_WidgetAttribute = 119
	Qt_WA_TranslucentBackground Qt_WidgetAttribute = 120
	Qt_WA_AcceptTouchEvents Qt_WidgetAttribute = 121
	Qt_WA_WState_AcceptedTouchBeginEvent Qt_WidgetAttribute = 122
	Qt_WA_TouchPadAcceptSingleTouchEvents Qt_WidgetAttribute = 123
	Qt_WA_MergeSoftkeys Qt_WidgetAttribute = 124
	Qt_WA_MergeSoftkeysRecursively Qt_WidgetAttribute = 125
	Qt_WA_LockPortraitOrientation Qt_WidgetAttribute = 128
	Qt_WA_LockLandscapeOrientation Qt_WidgetAttribute = 129
	Qt_WA_AutoOrientation Qt_WidgetAttribute = 130
	Qt_WA_X11DoNotAcceptFocus Qt_WidgetAttribute = 132
	Qt_WA_SymbianNoSystemRotation Qt_WidgetAttribute = 133
	Qt_WA_AttributeCount Qt_WidgetAttribute = 133+1
)

//enum Qt_TouchPointState - Qt::TouchPointState
type Qt_TouchPointState uint32
const (
	Qt_TouchPointPressed Qt_TouchPointState = 0x01
	Qt_TouchPointMoved Qt_TouchPointState = 0x02
	Qt_TouchPointStationary Qt_TouchPointState = 0x04
	Qt_TouchPointReleased Qt_TouchPointState = 0x08
	Qt_TouchPointStateMask Qt_TouchPointState = 0x0f
	Qt_TouchPointPrimary Qt_TouchPointState = 0x10
)

//enum Qt_TextInteractionFlag - Qt::TextInteractionFlag
type Qt_TextInteractionFlag uint32
const (
	Qt_NoTextInteraction Qt_TextInteractionFlag = 0
	Qt_TextSelectableByMouse Qt_TextInteractionFlag = 1
	Qt_TextSelectableByKeyboard Qt_TextInteractionFlag = 2
	Qt_LinksAccessibleByMouse Qt_TextInteractionFlag = 4
	Qt_LinksAccessibleByKeyboard Qt_TextInteractionFlag = 8
	Qt_TextEditable Qt_TextInteractionFlag = 16
	Qt_TextEditorInteraction Qt_TextInteractionFlag = Qt_TextSelectableByMouse|Qt_TextSelectableByKeyboard|Qt_TextEditable
	Qt_TextBrowserInteraction Qt_TextInteractionFlag = Qt_TextSelectableByMouse|Qt_LinksAccessibleByMouse|Qt_LinksAccessibleByKeyboard
)

//enum Qt_MouseButton - Qt::MouseButton
type Qt_MouseButton uint32
const (
	Qt_NoButton Qt_MouseButton = 0x00000000
	Qt_LeftButton Qt_MouseButton = 0x00000001
	Qt_RightButton Qt_MouseButton = 0x00000002
	Qt_MidButton Qt_MouseButton = 0x00000004
	Qt_MiddleButton Qt_MouseButton = Qt_MidButton
	Qt_XButton1 Qt_MouseButton = 0x00000008
	Qt_XButton2 Qt_MouseButton = 0x00000010
	Qt_MouseButtonMask Qt_MouseButton = 0x000000ff
)

//enum Qt_MatchFlag - Qt::MatchFlag
type Qt_MatchFlag uint32
const (
	Qt_MatchExactly Qt_MatchFlag = 0
	Qt_MatchContains Qt_MatchFlag = 1
	Qt_MatchStartsWith Qt_MatchFlag = 2
	Qt_MatchEndsWith Qt_MatchFlag = 3
	Qt_MatchRegExp Qt_MatchFlag = 4
	Qt_MatchWildcard Qt_MatchFlag = 5
	Qt_MatchFixedString Qt_MatchFlag = 8
	Qt_MatchCaseSensitive Qt_MatchFlag = 16
	Qt_MatchWrap Qt_MatchFlag = 32
	Qt_MatchRecursive Qt_MatchFlag = 64
)

//enum Qt_PenCapStyle - Qt::PenCapStyle
type Qt_PenCapStyle uint32
const (
	Qt_FlatCap Qt_PenCapStyle = 0x00
	Qt_SquareCap Qt_PenCapStyle = 0x10
	Qt_RoundCap Qt_PenCapStyle = 0x20
	Qt_MPenCapStyle Qt_PenCapStyle = 0x30
)

//enum Qt_TextElideMode - Qt::TextElideMode
type Qt_TextElideMode uint32
const (
	Qt_ElideLeft Qt_TextElideMode = 0
	Qt_ElideRight Qt_TextElideMode = 1
	Qt_ElideMiddle Qt_TextElideMode = 2
	Qt_ElideNone Qt_TextElideMode = 3
)

