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
	Count() int
	ExpandingDirections() Qt_Orientation
	Geometry() *QRect
	GetContentsMargins(*int,*int,*int,*int) 
	IndexOf(QWidgetInterface) int
	Invalidate() 
	IsEmpty() bool
	IsEnabled() bool
	ItemAt(int) *QLayoutItem
	Layout() *QLayout
	Margin() int
	MaximumSize() *QSize
	MenuBar() *QWidget
	MinimumSize() *QSize
	ParentWidget() *QWidget
	RemoveItem(*QLayoutItem) 
	RemoveWidget(QWidgetInterface) 
	SetAlignment(Qt_AlignmentFlag) 
	SetAlignmentWithLayoutAlignment(QLayoutInterface,Qt_AlignmentFlag) bool
	SetAlignmentWithWidgetAlignment(QWidgetInterface,Qt_AlignmentFlag) bool
	SetContentsMargins(*QMargins) 
	SetContentsMarginsWithLeftTopRightBottom(int,int,int,int) 
	SetEnabled(bool) 
	SetGeometry(*QRect) 
	SetMargin(int) 
	SetMenuBar(QWidgetInterface) 
	SetSizeConstraint(QLayout_SizeConstraint) 
	SetSpacing(int) 
	SizeConstraint() QLayout_SizeConstraint
	Spacing() int
	TakeAt(int) *QLayoutItem
	TotalHeightForWidth(int) int
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
	ChildAtWithXY(int,int) *QWidget
	ChildrenRect() *QRect
	ChildrenRegion() *QRegion
	ClearFocus() 
	ClearMask() 
	Close() bool
	ColorCount() int
	ContentsMargins() *QMargins
	ContentsRect() *QRect
	ContextMenuPolicy() Qt_ContextMenuPolicy
	CreateWinId() 
	Cursor() *QCursor
	Depth() int
	DevType() int
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
	GetContentsMargins(*int,*int,*int,*int) 
	GrabGesture(Qt_GestureType) 
	GrabGestureWithTypeFlags(Qt_GestureType,Qt_GestureFlag) 
	GrabKeyboard() 
	GrabMouse() 
	GrabMouseWithCursor(*QCursor) 
	GrabShortcut(*QKeySequence) int
	GrabShortcutWithKeyContext(*QKeySequence,Qt_ShortcutContext) int
	GraphicsEffect() *QGraphicsEffect
	GraphicsProxyWidget() *QGraphicsProxyWidget
	HasFocus() bool
	HasMouseTracking() bool
	Height() int
	HeightForWidth(int) int
	HeightMM() int
	Hide() 
	InputMethodHints() Qt_InputMethodHint
	InputMethodQuery(Qt_InputMethodQuery) *QVariant
	InsertAction(*QAction,*QAction) 
	InsertActions(*QAction,[]*QAction) 
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
	LogicalDpiX() int
	LogicalDpiY() int
	Lower() 
	MapFrom(QWidgetInterface,*QPoint) *QPoint
	MapFromGlobal(*QPoint) *QPoint
	MapFromParent(*QPoint) *QPoint
	MapTo(QWidgetInterface,*QPoint) *QPoint
	MapToGlobal(*QPoint) *QPoint
	MapToParent(*QPoint) *QPoint
	Mask() *QRegion
	MaximumHeight() int
	MaximumSize() *QSize
	MaximumWidth() int
	MinimumHeight() int
	MinimumSize() *QSize
	MinimumSizeHint() *QSize
	MinimumWidth() int
	Move(*QPoint) 
	MoveWithXY(int,int) 
	NativeParentWidget() *QWidget
	NextInFocusChain() *QWidget
	NormalGeometry() *QRect
	OverrideWindowFlags(Qt_WindowType) 
	OverrideWindowState(Qt_WindowState) 
	PaintEngine() *QPaintEngine
	PaintingActive() bool
	Palette() *QPalette
	ParentWidget() *QWidget
	PhysicalDpiX() int
	PhysicalDpiY() int
	Pos() *QPoint
	PreviousInFocusChain() *QWidget
	Raise() 
	Rect() *QRect
	ReleaseKeyboard() 
	ReleaseMouse() 
	ReleaseShortcut(int) 
	RemoveAction(*QAction) 
	Render(QPaintDeviceInterface) 
	RenderWithPaintDeviceTargetoffsetSourceregionRenderflags(QPaintDeviceInterface,*QPoint,*QRegion,QWidget_RenderFlag) 
	RenderWithPainter(*QPainter) 
	RenderWithPainterTargetoffsetSourceregionRenderflags(*QPainter,*QPoint,*QRegion,QWidget_RenderFlag) 
	Repaint() 
	RepaintWithRect(*QRect) 
	RepaintWithRegion(*QRegion) 
	RepaintWithXYWidthHeight(int,int,int,int) 
	Resize(*QSize) 
	ResizeWithWidthHeight(int,int) 
	RestoreGeometry([]byte) bool
	SaveGeometry() []byte
	ScrollWithDxDy(int,int) 
	ScrollWithDxDyRect(int,int,*QRect) 
	SetAcceptDrops(bool) 
	SetAccessibleDescription(string) 
	SetAccessibleName(string) 
	SetAttribute(Qt_WidgetAttribute) 
	SetAttributeWithWidgetattributeOn(Qt_WidgetAttribute,bool) 
	SetAutoFillBackground(bool) 
	SetBackgroundRole(QPalette_ColorRole) 
	SetBaseSize(*QSize) 
	SetBaseSizeWithBasewBaseh(int,int) 
	SetContentsMargins(*QMargins) 
	SetContentsMarginsWithLeftTopRightBottom(int,int,int,int) 
	SetContextMenuPolicy(Qt_ContextMenuPolicy) 
	SetCursor(*QCursor) 
	SetDisabled(bool) 
	SetEnabled(bool) 
	SetFixedHeight(int) 
	SetFixedSize(*QSize) 
	SetFixedSizeWithWidthHeight(int,int) 
	SetFixedWidth(int) 
	SetFocus() 
	SetFocusPolicy(Qt_FocusPolicy) 
	SetFocusProxy(QWidgetInterface) 
	SetFocusWithReason(Qt_FocusReason) 
	SetFont(*QFont) 
	SetForegroundRole(QPalette_ColorRole) 
	SetGeometry(*QRect) 
	SetGeometryWithXYWidthHeight(int,int,int,int) 
	SetGraphicsEffect(*QGraphicsEffect) 
	SetHidden(bool) 
	SetInputMethodHints(Qt_InputMethodHint) 
	SetLayout(QLayoutInterface) 
	SetLayoutDirection(Qt_LayoutDirection) 
	SetLocale(*QLocale) 
	SetMask(*QBitmap) 
	SetMaskWithRegion(*QRegion) 
	SetMaximumHeight(int) 
	SetMaximumSize(*QSize) 
	SetMaximumSizeWithMaxwMaxh(int,int) 
	SetMaximumWidth(int) 
	SetMinimumHeight(int) 
	SetMinimumSize(*QSize) 
	SetMinimumSizeWithMinwMinh(int,int) 
	SetMinimumWidth(int) 
	SetMouseTracking(bool) 
	SetPalette(*QPalette) 
	SetParentWidget(QWidgetInterface) 
	SetParentWidgetWithParentFlags(QWidgetInterface,Qt_WindowType) 
	SetShortcutAutoRepeat(int) 
	SetShortcutAutoRepeatWithIdEnable(int,bool) 
	SetShortcutEnabled(int) 
	SetShortcutEnabledWithIdEnable(int,bool) 
	SetSizeIncrement(*QSize) 
	SetSizeIncrementWithWidthHeight(int,int) 
	SetSizePolicy(*QSizePolicy) 
	SetSizePolicyWithHorizontalVertical(QSizePolicy_Policy,QSizePolicy_Policy) 
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
	UpdateWithXYWidthHeight(int,int,int,int) 
	UpdatesEnabled() bool
	VisibleRegion() *QRegion
	WhatsThis() string
	Width() int
	WidthMM() int
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
	X() int
	Y() int
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
	KillTimer(int) 
	MetaObject() *QMetaObject
	ObjectName() string
	Parent() *QObject
	Property(string) *QVariant
	SetObjectName(string) 
	SetParent(QObjectInterface) 
	SetProperty(string,*QVariant) bool
	SignalsBlocked() bool
	StartTimer(int) int
	Tr(string) string
	TrWithSourcetextDisambiguation(string,string) string
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
	PeekWithDataMaxlen(*byte,int64) int64
	Pos() int64
	PutChar(byte) bool
	Read(int64) []byte
	ReadAll() []byte
	ReadLine() []byte
	ReadLineWithDataMaxlen(*byte,int64) int64
	ReadLineWithMaxlen(int64) []byte
	ReadWithDataMaxlen(*byte,int64) int64
	Reset() bool
	Seek(int64) bool
	SetTextModeEnabled(bool) 
	Size() int64
	UngetChar(byte) 
	WaitForBytesWritten(int) bool
	WaitForReadyRead(int) bool
	Write([]byte) int64
	WriteWithDataLen(*byte,int64) int64
}

type QPaintDeviceInterface interface {
	Driver
	ColorCount() int
	Depth() int
	DevType() int
	Height() int
	HeightMM() int
	LogicalDpiX() int
	LogicalDpiY() int
	PaintEngine() *QPaintEngine
	PaintingActive() bool
	PhysicalDpiX() int
	PhysicalDpiY() int
	Width() int
	WidthMM() int
}

type QAbstractItemModelInterface interface {
	QObjectInterface
	Buddy(*QModelIndex) *QModelIndex
	CanFetchMore(*QModelIndex) bool
	ColumnCount() int
	ColumnCountWithParent(*QModelIndex) int
	Data(*QModelIndex) *QVariant
	DataWithIndexRole(*QModelIndex,int) *QVariant
	DropMimeData(*QMimeData,Qt_DropAction,int,int,*QModelIndex) bool
	FetchMore(*QModelIndex) 
	Flags(*QModelIndex) Qt_ItemFlag
	HasChildren() bool
	HasChildrenWithParent(*QModelIndex) bool
	HasIndex(int,int,*QModelIndex) bool
	HeaderData(int,Qt_Orientation,int) *QVariant
	Index(int,int,*QModelIndex) *QModelIndex
	InsertColumn(int) bool
	InsertColumnWithColumnParent(int,*QModelIndex) bool
	InsertColumns(int,int,*QModelIndex) bool
	InsertRow(int) bool
	InsertRowWithRowParent(int,*QModelIndex) bool
	InsertRows(int,int,*QModelIndex) bool
	ItemData(*QModelIndex) map[int]*QVariant
	Match(*QModelIndex,int,*QVariant,int,Qt_MatchFlag) []*QModelIndex
	MimeData([]*QModelIndex) *QMimeData
	MimeTypes() []string
	RemoveColumn(int) bool
	RemoveColumnWithColumnParent(int,*QModelIndex) bool
	RemoveColumns(int,int,*QModelIndex) bool
	RemoveRow(int) bool
	RemoveRowWithRowParent(int,*QModelIndex) bool
	RemoveRows(int,int,*QModelIndex) bool
	Revert() 
	RowCount() int
	RowCountWithParent(*QModelIndex) int
	SetData(*QModelIndex,*QVariant,int) bool
	SetHeaderData(int,Qt_Orientation,*QVariant,int) bool
	SetItemData(*QModelIndex,map[int]*QVariant) bool
	Sibling(int,int,*QModelIndex) *QModelIndex
	Sort(int) 
	SortWithColumnOrder(int,Qt_SortOrder) 
	Span(*QModelIndex) *QSize
	Submit() bool
	SupportedDragActions() Qt_DropAction
	SupportedDropActions() Qt_DropAction
}

