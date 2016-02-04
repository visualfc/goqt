// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

type QLayoutInterface interface {
	QObjectInterface
	Activate() bool
	AddItem(*QLayoutItem)
	AddWidget(QWidgetInterface)
	ContentsMargins() *QMargins
	ContentsRect() *QRect
	Count() int32
	ExpandingDirections() Qt_Orientation
	Geometry() *QRect
	GetContentsMargins(*int32, *int32, *int32, *int32)
	IndexOf(QWidgetInterface) int32
	Invalidate()
	IsEmpty() bool
	IsEnabled() bool
	ItemAt(int32) *QLayoutItem
	Layout() *QLayout
	Margin() int32
	MaximumSize() *QSize
	MenuBar() *QWidget
	MinimumSize() *QSize
	ParentWidget() *QWidget
	RemoveItem(*QLayoutItem)
	RemoveWidget(QWidgetInterface)
	SetAlignment(Qt_AlignmentFlag)
	SetAlignmentWithLayoutAlignment(QLayoutInterface, Qt_AlignmentFlag) bool
	SetAlignmentWithWidgetAlignment(QWidgetInterface, Qt_AlignmentFlag) bool
	SetContentsMargins(*QMargins)
	SetContentsMarginsWithLeftTopRightBottom(int32, int32, int32, int32)
	SetEnabled(bool)
	SetGeometry(*QRect)
	SetMargin(int32)
	SetMenuBar(QWidgetInterface)
	SetSizeConstraint(QLayout_SizeConstraint)
	SetSpacing(int32)
	SizeConstraint() QLayout_SizeConstraint
	Spacing() int32
	TakeAt(int32) *QLayoutItem
	TotalHeightForWidth(int32) int32
	TotalMaximumSize() *QSize
	TotalMinimumSize() *QSize
	TotalSizeHint() *QSize
	Update()
}

type QWidgetInterface interface {
	QObjectInterface
	AcceptDrops() bool
	AccessibleDescription() string
	AccessibleName() string
	Actions() []*QAction
	ActivateWindow()
	AddAction(*QAction)
	AddActions([]*QAction)
	AdjustSize()
	AutoFillBackground() bool
	BackgroundRole() QPalette_ColorRole
	BaseSize() *QSize
	ChildAt(*QPoint) *QWidget
	ChildAtWithXY(int32, int32) *QWidget
	ChildrenRect() *QRect
	ChildrenRegion() *QRegion
	ClearFocus()
	ClearMask()
	Close() bool
	ColorCount() int32
	ContentsMargins() *QMargins
	ContentsRect() *QRect
	ContextMenuPolicy() Qt_ContextMenuPolicy
	CreateWinId()
	Cursor() *QCursor
	Depth() int32
	DevType() int32
	EnsurePolished()
	FocusPolicy() Qt_FocusPolicy
	FocusProxy() *QWidget
	FocusWidget() *QWidget
	Font() *QFont
	FontInfo() *QFontInfo
	FontMetrics() *QFontMetrics
	ForegroundRole() QPalette_ColorRole
	FrameGeometry() *QRect
	FrameSize() *QSize
	Geometry() *QRect
	GetContentsMargins(*int32, *int32, *int32, *int32)
	GrabGesture(Qt_GestureType)
	GrabGestureWithTypeFlags(Qt_GestureType, Qt_GestureFlag)
	GrabKeyboard()
	GrabMouse()
	GrabMouseWithCursor(*QCursor)
	GrabShortcut(*QKeySequence) int32
	GrabShortcutWithKeyContext(*QKeySequence, Qt_ShortcutContext) int32
	GraphicsEffect() *QGraphicsEffect
	GraphicsProxyWidget() *QGraphicsProxyWidget
	HasFocus() bool
	HasMouseTracking() bool
	Height() int32
	HeightForWidth(int32) int32
	HeightMM() int32
	Hide()
	InputMethodHints() Qt_InputMethodHint
	InputMethodQuery(Qt_InputMethodQuery) *QVariant
	InsertAction(*QAction, *QAction)
	InsertActions(*QAction, []*QAction)
	IsActiveWindow() bool
	IsAncestorOf(QWidgetInterface) bool
	IsEnabled() bool
	IsEnabledTo(QWidgetInterface) bool
	IsEnabledToTLW() bool
	IsFullScreen() bool
	IsHidden() bool
	IsLeftToRight() bool
	IsMaximized() bool
	IsMinimized() bool
	IsModal() bool
	IsRightToLeft() bool
	IsTopLevel() bool
	IsVisible() bool
	IsVisibleTo(QWidgetInterface) bool
	IsWindow() bool
	IsWindowModified() bool
	Layout() *QLayout
	LayoutDirection() Qt_LayoutDirection
	Locale() *QLocale
	LogicalDpiX() int32
	LogicalDpiY() int32
	Lower()
	MapFrom(QWidgetInterface, *QPoint) *QPoint
	MapFromGlobal(*QPoint) *QPoint
	MapFromParent(*QPoint) *QPoint
	MapTo(QWidgetInterface, *QPoint) *QPoint
	MapToGlobal(*QPoint) *QPoint
	MapToParent(*QPoint) *QPoint
	Mask() *QRegion
	MaximumHeight() int32
	MaximumSize() *QSize
	MaximumWidth() int32
	MinimumHeight() int32
	MinimumSize() *QSize
	MinimumSizeHint() *QSize
	MinimumWidth() int32
	Move(*QPoint)
	MoveWithXY(int32, int32)
	NativeParentWidget() *QWidget
	NextInFocusChain() *QWidget
	NormalGeometry() *QRect
	OverrideWindowFlags(Qt_WindowType)
	OverrideWindowState(Qt_WindowState)
	PaintEngine() *QPaintEngine
	PaintingActive() bool
	Palette() *QPalette
	ParentWidget() *QWidget
	PhysicalDpiX() int32
	PhysicalDpiY() int32
	Pos() *QPoint
	PreviousInFocusChain() *QWidget
	Raise()
	Rect() *QRect
	ReleaseKeyboard()
	ReleaseMouse()
	ReleaseShortcut(int32)
	RemoveAction(*QAction)
	Render(QPaintDeviceInterface)
	RenderWithPaintDeviceTargetoffsetSourceregionRenderflags(QPaintDeviceInterface, *QPoint, *QRegion, QWidget_RenderFlag)
	RenderWithPainter(*QPainter)
	RenderWithPainterTargetoffsetSourceregionRenderflags(*QPainter, *QPoint, *QRegion, QWidget_RenderFlag)
	Repaint()
	RepaintWithRect(*QRect)
	RepaintWithRegion(*QRegion)
	RepaintWithXYWidthHeight(int32, int32, int32, int32)
	Resize(*QSize)
	ResizeWithWidthHeight(int32, int32)
	RestoreGeometry([]byte) bool
	SaveGeometry() []byte
	ScrollWithDxDy(int32, int32)
	ScrollWithDxDyRect(int32, int32, *QRect)
	SetAcceptDrops(bool)
	SetAccessibleDescription(string)
	SetAccessibleName(string)
	SetAttribute(Qt_WidgetAttribute)
	SetAttributeWithWidgetattributeOn(Qt_WidgetAttribute, bool)
	SetAutoFillBackground(bool)
	SetBackgroundRole(QPalette_ColorRole)
	SetBaseSize(*QSize)
	SetBaseSizeWithBasewBaseh(int32, int32)
	SetContentsMargins(*QMargins)
	SetContentsMarginsWithLeftTopRightBottom(int32, int32, int32, int32)
	SetContextMenuPolicy(Qt_ContextMenuPolicy)
	SetCursor(*QCursor)
	SetDisabled(bool)
	SetEnabled(bool)
	SetFixedHeight(int32)
	SetFixedSize(*QSize)
	SetFixedSizeWithWidthHeight(int32, int32)
	SetFixedWidth(int32)
	SetFocus()
	SetFocusPolicy(Qt_FocusPolicy)
	SetFocusProxy(QWidgetInterface)
	SetFocusWithReason(Qt_FocusReason)
	SetFont(*QFont)
	SetForegroundRole(QPalette_ColorRole)
	SetGeometry(*QRect)
	SetGeometryWithXYWidthHeight(int32, int32, int32, int32)
	SetGraphicsEffect(*QGraphicsEffect)
	SetHidden(bool)
	SetInputMethodHints(Qt_InputMethodHint)
	SetLayout(QLayoutInterface)
	SetLayoutDirection(Qt_LayoutDirection)
	SetLocale(*QLocale)
	SetMask(*QBitmap)
	SetMaskWithRegion(*QRegion)
	SetMaximumHeight(int32)
	SetMaximumSize(*QSize)
	SetMaximumSizeWithMaxwMaxh(int32, int32)
	SetMaximumWidth(int32)
	SetMinimumHeight(int32)
	SetMinimumSize(*QSize)
	SetMinimumSizeWithMinwMinh(int32, int32)
	SetMinimumWidth(int32)
	SetMouseTracking(bool)
	SetPalette(*QPalette)
	SetParentWidget(QWidgetInterface)
	SetParentWidgetWithParentFlags(QWidgetInterface, Qt_WindowType)
	SetShortcutAutoRepeat(int32)
	SetShortcutAutoRepeatWithIdEnable(int32, bool)
	SetShortcutEnabled(int32)
	SetShortcutEnabledWithIdEnable(int32, bool)
	SetSizeIncrement(*QSize)
	SetSizeIncrementWithWidthHeight(int32, int32)
	SetSizePolicy(*QSizePolicy)
	SetSizePolicyWithHorizontalVertical(QSizePolicy_Policy, QSizePolicy_Policy)
	SetStatusTip(string)
	SetStyle(*QStyle)
	SetStyleSheet(string)
	SetToolTip(string)
	SetUpdatesEnabled(bool)
	SetVisible(bool)
	SetWhatsThis(string)
	SetWindowFilePath(string)
	SetWindowFlags(Qt_WindowType)
	SetWindowIcon(*QIcon)
	SetWindowIconText(string)
	SetWindowModality(Qt_WindowModality)
	SetWindowModified(bool)
	SetWindowOpacity(float64)
	SetWindowRole(string)
	SetWindowState(Qt_WindowState)
	SetWindowTitle(string)
	Show()
	ShowFullScreen()
	ShowMaximized()
	ShowMinimized()
	ShowNormal()
	Size() *QSize
	SizeHint() *QSize
	SizeIncrement() *QSize
	SizePolicy() *QSizePolicy
	StackUnder(QWidgetInterface)
	StatusTip() string
	Style() *QStyle
	StyleSheet() string
	TestAttribute(Qt_WidgetAttribute) bool
	ToolTip() string
	TopLevelWidget() *QWidget
	UnderMouse() bool
	UngrabGesture(Qt_GestureType)
	UnsetCursor()
	UnsetLayoutDirection()
	UnsetLocale()
	Update()
	UpdateGeometry()
	UpdateWithRect(*QRect)
	UpdateWithRegion(*QRegion)
	UpdateWithXYWidthHeight(int32, int32, int32, int32)
	UpdatesEnabled() bool
	VisibleRegion() *QRegion
	WhatsThis() string
	Width() int32
	WidthMM() int32
	Window() *QWidget
	WindowFilePath() string
	WindowFlags() Qt_WindowType
	WindowIcon() *QIcon
	WindowIconText() string
	WindowModality() Qt_WindowModality
	WindowOpacity() float64
	WindowRole() string
	WindowState() Qt_WindowState
	WindowTitle() string
	WindowType() Qt_WindowType
	X() int32
	Y() int32
}

type QObjectInterface interface {
	Driver
	BlockSignals(bool) bool
	Children() []*QObject
	DeleteLater()
	DynamicPropertyNames() [][]byte
	Event(*QEvent) bool
	FindChild(string) *QObject
	FindChildren(string) []*QObject
	FindChildrenWithRegexp(*QRegExp) []*QObject
	Inherits(string) bool
	IsWidgetType() bool
	KillTimer(int32)
	MetaObject() *QMetaObject
	ObjectName() string
	Parent() *QObject
	Property(string) *QVariant
	SetObjectName(string)
	SetParent(QObjectInterface)
	SetProperty(string, *QVariant) bool
	SignalsBlocked() bool
	StartTimer(int32) int32
	Tr(string) string
	TrWithSourcetextDisambiguation(string, string) string
}

type QIODeviceInterface interface {
	QObjectInterface
	AtEnd() bool
	BytesAvailable() int64
	BytesToWrite() int64
	CanReadLine() bool
	Close()
	ErrorString() string
	GetChar(*byte) bool
	IsOpen() bool
	IsReadable() bool
	IsSequential() bool
	IsTextModeEnabled() bool
	IsWritable() bool
	Open(QIODevice_OpenModeFlag) bool
	OpenMode() QIODevice_OpenModeFlag
	Peek(int64) []byte
	PeekWithDataMaxlen(*byte, int64) int64
	Pos() int64
	PutChar(byte) bool
	Read(int64) []byte
	ReadAll() []byte
	ReadLine() []byte
	ReadLineWithDataMaxlen(*byte, int64) int64
	ReadLineWithMaxlen(int64) []byte
	ReadWithDataMaxlen(*byte, int64) int64
	Reset() bool
	Seek(int64) bool
	SetTextModeEnabled(bool)
	Size() int64
	UngetChar(byte)
	WaitForBytesWritten(int32) bool
	WaitForReadyRead(int32) bool
	Write([]byte) int64
	WriteWithDataLen(*byte, int64) int64
}

type QPaintDeviceInterface interface {
	Driver
	ColorCount() int32
	Depth() int32
	DevType() int32
	Height() int32
	HeightMM() int32
	LogicalDpiX() int32
	LogicalDpiY() int32
	PaintEngine() *QPaintEngine
	PaintingActive() bool
	PhysicalDpiX() int32
	PhysicalDpiY() int32
	Width() int32
	WidthMM() int32
}

type QAbstractItemModelInterface interface {
	QObjectInterface
	Buddy(*QModelIndex) *QModelIndex
	CanFetchMore(*QModelIndex) bool
	ColumnCount() int32
	ColumnCountWithParent(*QModelIndex) int32
	Data(*QModelIndex) *QVariant
	DataWithIndexRole(*QModelIndex, int32) *QVariant
	DropMimeData(*QMimeData, Qt_DropAction, int32, int32, *QModelIndex) bool
	FetchMore(*QModelIndex)
	Flags(*QModelIndex) Qt_ItemFlag
	HasChildren() bool
	HasChildrenWithParent(*QModelIndex) bool
	HasIndex(int32, int32, *QModelIndex) bool
	HeaderData(int32, Qt_Orientation, int32) *QVariant
	Index(int32, int32, *QModelIndex) *QModelIndex
	InsertColumn(int32) bool
	InsertColumnWithColumnParent(int32, *QModelIndex) bool
	InsertColumns(int32, int32, *QModelIndex) bool
	InsertRow(int32) bool
	InsertRowWithRowParent(int32, *QModelIndex) bool
	InsertRows(int32, int32, *QModelIndex) bool
	ItemData(*QModelIndex) map[int]*QVariant
	Match(*QModelIndex, int32, *QVariant, int32, Qt_MatchFlag) []*QModelIndex
	MimeData([]*QModelIndex) *QMimeData
	MimeTypes() []string
	RemoveColumn(int32) bool
	RemoveColumnWithColumnParent(int32, *QModelIndex) bool
	RemoveColumns(int32, int32, *QModelIndex) bool
	RemoveRow(int32) bool
	RemoveRowWithRowParent(int32, *QModelIndex) bool
	RemoveRows(int32, int32, *QModelIndex) bool
	Revert()
	RowCount() int32
	RowCountWithParent(*QModelIndex) int32
	SetData(*QModelIndex, *QVariant, int32) bool
	SetHeaderData(int32, Qt_Orientation, *QVariant, int32) bool
	SetItemData(*QModelIndex, map[int]*QVariant) bool
	Sibling(int32, int32, *QModelIndex) *QModelIndex
	Sort(int32)
	SortWithColumnOrder(int32, Qt_SortOrder)
	Span(*QModelIndex) *QSize
	Submit() bool
	SupportedDragActions() Qt_DropAction
	SupportedDropActions() Qt_DropAction
}
