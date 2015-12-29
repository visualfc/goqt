// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStyle_PixelMetric - QStyle::PixelMetric
type QStyle_PixelMetric uint32
const (
	QStyle_PM_ButtonMargin QStyle_PixelMetric = 0
	QStyle_PM_ButtonDefaultIndicator QStyle_PixelMetric = 1
	QStyle_PM_MenuButtonIndicator QStyle_PixelMetric = 2
	QStyle_PM_ButtonShiftHorizontal QStyle_PixelMetric = 3
	QStyle_PM_ButtonShiftVertical QStyle_PixelMetric = 4
	QStyle_PM_DefaultFrameWidth QStyle_PixelMetric = 5
	QStyle_PM_SpinBoxFrameWidth QStyle_PixelMetric = 6
	QStyle_PM_ComboBoxFrameWidth QStyle_PixelMetric = 7
	QStyle_PM_MaximumDragDistance QStyle_PixelMetric = 8
	QStyle_PM_ScrollBarExtent QStyle_PixelMetric = 9
	QStyle_PM_ScrollBarSliderMin QStyle_PixelMetric = 10
	QStyle_PM_SliderThickness QStyle_PixelMetric = 11
	QStyle_PM_SliderControlThickness QStyle_PixelMetric = 12
	QStyle_PM_SliderLength QStyle_PixelMetric = 13
	QStyle_PM_SliderTickmarkOffset QStyle_PixelMetric = 14
	QStyle_PM_SliderSpaceAvailable QStyle_PixelMetric = 15
	QStyle_PM_DockWidgetSeparatorExtent QStyle_PixelMetric = 16
	QStyle_PM_DockWidgetHandleExtent QStyle_PixelMetric = 17
	QStyle_PM_DockWidgetFrameWidth QStyle_PixelMetric = 18
	QStyle_PM_TabBarTabOverlap QStyle_PixelMetric = 19
	QStyle_PM_TabBarTabHSpace QStyle_PixelMetric = 20
	QStyle_PM_TabBarTabVSpace QStyle_PixelMetric = 21
	QStyle_PM_TabBarBaseHeight QStyle_PixelMetric = 22
	QStyle_PM_TabBarBaseOverlap QStyle_PixelMetric = 23
	QStyle_PM_ProgressBarChunkWidth QStyle_PixelMetric = 24
	QStyle_PM_SplitterWidth QStyle_PixelMetric = 25
	QStyle_PM_TitleBarHeight QStyle_PixelMetric = 26
	QStyle_PM_MenuScrollerHeight QStyle_PixelMetric = 27
	QStyle_PM_MenuHMargin QStyle_PixelMetric = 28
	QStyle_PM_MenuVMargin QStyle_PixelMetric = 29
	QStyle_PM_MenuPanelWidth QStyle_PixelMetric = 30
	QStyle_PM_MenuTearoffHeight QStyle_PixelMetric = 31
	QStyle_PM_MenuDesktopFrameWidth QStyle_PixelMetric = 32
	QStyle_PM_MenuBarPanelWidth QStyle_PixelMetric = 33
	QStyle_PM_MenuBarItemSpacing QStyle_PixelMetric = 34
	QStyle_PM_MenuBarVMargin QStyle_PixelMetric = 35
	QStyle_PM_MenuBarHMargin QStyle_PixelMetric = 36
	QStyle_PM_IndicatorWidth QStyle_PixelMetric = 37
	QStyle_PM_IndicatorHeight QStyle_PixelMetric = 38
	QStyle_PM_ExclusiveIndicatorWidth QStyle_PixelMetric = 39
	QStyle_PM_ExclusiveIndicatorHeight QStyle_PixelMetric = 40
	QStyle_PM_CheckListButtonSize QStyle_PixelMetric = 41
	QStyle_PM_CheckListControllerSize QStyle_PixelMetric = 42
	QStyle_PM_DialogButtonsSeparator QStyle_PixelMetric = 43
	QStyle_PM_DialogButtonsButtonWidth QStyle_PixelMetric = 44
	QStyle_PM_DialogButtonsButtonHeight QStyle_PixelMetric = 45
	QStyle_PM_MdiSubWindowFrameWidth QStyle_PixelMetric = 46
	QStyle_PM_MDIFrameWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowFrameWidth
	QStyle_PM_MdiSubWindowMinimizedWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowFrameWidth+1
	QStyle_PM_MDIMinimizedWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth
	QStyle_PM_HeaderMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1
	QStyle_PM_HeaderMarkSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1
	QStyle_PM_HeaderGripMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1
	QStyle_PM_TabBarTabShiftHorizontal QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1
	QStyle_PM_TabBarTabShiftVertical QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1
	QStyle_PM_TabBarScrollButtonWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1
	QStyle_PM_ToolBarFrameWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1
	QStyle_PM_ToolBarHandleExtent QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1
	QStyle_PM_ToolBarItemSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1
	QStyle_PM_ToolBarItemMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ToolBarSeparatorExtent QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ToolBarExtensionExtent QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_SpinBoxSliderHeight QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_DefaultTopLevelMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_DefaultChildMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_DefaultLayoutSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ToolBarIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ListViewIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_IconViewIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_SmallIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LargeIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_FocusFrameVMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_FocusFrameHMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ToolTipLabelFrameWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_CheckBoxLabelSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_TabBarIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_SizeGripSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_DockWidgetTitleMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_MessageBoxIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ButtonIconSize QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_DockWidgetTitleBarButtonMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_RadioButtonLabelSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LayoutLeftMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LayoutTopMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LayoutRightMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LayoutBottomMargin QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LayoutHorizontalSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_LayoutVerticalSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_TabBar_ScrollButtonOverlap QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_TextCursorWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_TabCloseIndicatorWidth QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_TabCloseIndicatorHeight QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_ScrollView_ScrollBarSpacing QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_SubMenuOverlap QStyle_PixelMetric = QStyle_PM_MdiSubWindowMinimizedWidth+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PM_CustomBase QStyle_PixelMetric = 0xf0000000
)
//enum QStyle_SubControl - QStyle::SubControl
type QStyle_SubControl uint32
const (
	QStyle_SC_None QStyle_SubControl = 0x00000000
	QStyle_SC_ScrollBarAddLine QStyle_SubControl = 0x00000001
	QStyle_SC_ScrollBarSubLine QStyle_SubControl = 0x00000002
	QStyle_SC_ScrollBarAddPage QStyle_SubControl = 0x00000004
	QStyle_SC_ScrollBarSubPage QStyle_SubControl = 0x00000008
	QStyle_SC_ScrollBarFirst QStyle_SubControl = 0x00000010
	QStyle_SC_ScrollBarLast QStyle_SubControl = 0x00000020
	QStyle_SC_ScrollBarSlider QStyle_SubControl = 0x00000040
	QStyle_SC_ScrollBarGroove QStyle_SubControl = 0x00000080
	QStyle_SC_SpinBoxUp QStyle_SubControl = 0x00000001
	QStyle_SC_SpinBoxDown QStyle_SubControl = 0x00000002
	QStyle_SC_SpinBoxFrame QStyle_SubControl = 0x00000004
	QStyle_SC_SpinBoxEditField QStyle_SubControl = 0x00000008
	QStyle_SC_ComboBoxFrame QStyle_SubControl = 0x00000001
	QStyle_SC_ComboBoxEditField QStyle_SubControl = 0x00000002
	QStyle_SC_ComboBoxArrow QStyle_SubControl = 0x00000004
	QStyle_SC_ComboBoxListBoxPopup QStyle_SubControl = 0x00000008
	QStyle_SC_SliderGroove QStyle_SubControl = 0x00000001
	QStyle_SC_SliderHandle QStyle_SubControl = 0x00000002
	QStyle_SC_SliderTickmarks QStyle_SubControl = 0x00000004
	QStyle_SC_ToolButton QStyle_SubControl = 0x00000001
	QStyle_SC_ToolButtonMenu QStyle_SubControl = 0x00000002
	QStyle_SC_TitleBarSysMenu QStyle_SubControl = 0x00000001
	QStyle_SC_TitleBarMinButton QStyle_SubControl = 0x00000002
	QStyle_SC_TitleBarMaxButton QStyle_SubControl = 0x00000004
	QStyle_SC_TitleBarCloseButton QStyle_SubControl = 0x00000008
	QStyle_SC_TitleBarNormalButton QStyle_SubControl = 0x00000010
	QStyle_SC_TitleBarShadeButton QStyle_SubControl = 0x00000020
	QStyle_SC_TitleBarUnshadeButton QStyle_SubControl = 0x00000040
	QStyle_SC_TitleBarContextHelpButton QStyle_SubControl = 0x00000080
	QStyle_SC_TitleBarLabel QStyle_SubControl = 0x00000100
	QStyle_SC_Q3ListView QStyle_SubControl = 0x00000001
	QStyle_SC_Q3ListViewBranch QStyle_SubControl = 0x00000002
	QStyle_SC_Q3ListViewExpand QStyle_SubControl = 0x00000004
	QStyle_SC_DialGroove QStyle_SubControl = 0x00000001
	QStyle_SC_DialHandle QStyle_SubControl = 0x00000002
	QStyle_SC_DialTickmarks QStyle_SubControl = 0x00000004
	QStyle_SC_GroupBoxCheckBox QStyle_SubControl = 0x00000001
	QStyle_SC_GroupBoxLabel QStyle_SubControl = 0x00000002
	QStyle_SC_GroupBoxContents QStyle_SubControl = 0x00000004
	QStyle_SC_GroupBoxFrame QStyle_SubControl = 0x00000008
	QStyle_SC_MdiMinButton QStyle_SubControl = 0x00000001
	QStyle_SC_MdiNormalButton QStyle_SubControl = 0x00000002
	QStyle_SC_MdiCloseButton QStyle_SubControl = 0x00000004
	QStyle_SC_CustomBase QStyle_SubControl = 0xf0000000
	QStyle_SC_All QStyle_SubControl = 0xffffffff
)
//enum QStyle_StandardPixmap - QStyle::StandardPixmap
type QStyle_StandardPixmap uint32
const (
	QStyle_SP_TitleBarMenuButton QStyle_StandardPixmap = 0
	QStyle_SP_TitleBarMinButton QStyle_StandardPixmap = 1
	QStyle_SP_TitleBarMaxButton QStyle_StandardPixmap = 2
	QStyle_SP_TitleBarCloseButton QStyle_StandardPixmap = 3
	QStyle_SP_TitleBarNormalButton QStyle_StandardPixmap = 4
	QStyle_SP_TitleBarShadeButton QStyle_StandardPixmap = 5
	QStyle_SP_TitleBarUnshadeButton QStyle_StandardPixmap = 6
	QStyle_SP_TitleBarContextHelpButton QStyle_StandardPixmap = 7
	QStyle_SP_DockWidgetCloseButton QStyle_StandardPixmap = 8
	QStyle_SP_MessageBoxInformation QStyle_StandardPixmap = 9
	QStyle_SP_MessageBoxWarning QStyle_StandardPixmap = 10
	QStyle_SP_MessageBoxCritical QStyle_StandardPixmap = 11
	QStyle_SP_MessageBoxQuestion QStyle_StandardPixmap = 12
	QStyle_SP_DesktopIcon QStyle_StandardPixmap = 13
	QStyle_SP_TrashIcon QStyle_StandardPixmap = 14
	QStyle_SP_ComputerIcon QStyle_StandardPixmap = 15
	QStyle_SP_DriveFDIcon QStyle_StandardPixmap = 16
	QStyle_SP_DriveHDIcon QStyle_StandardPixmap = 17
	QStyle_SP_DriveCDIcon QStyle_StandardPixmap = 18
	QStyle_SP_DriveDVDIcon QStyle_StandardPixmap = 19
	QStyle_SP_DriveNetIcon QStyle_StandardPixmap = 20
	QStyle_SP_DirOpenIcon QStyle_StandardPixmap = 21
	QStyle_SP_DirClosedIcon QStyle_StandardPixmap = 22
	QStyle_SP_DirLinkIcon QStyle_StandardPixmap = 23
	QStyle_SP_FileIcon QStyle_StandardPixmap = 24
	QStyle_SP_FileLinkIcon QStyle_StandardPixmap = 25
	QStyle_SP_ToolBarHorizontalExtensionButton QStyle_StandardPixmap = 26
	QStyle_SP_ToolBarVerticalExtensionButton QStyle_StandardPixmap = 27
	QStyle_SP_FileDialogStart QStyle_StandardPixmap = 28
	QStyle_SP_FileDialogEnd QStyle_StandardPixmap = 29
	QStyle_SP_FileDialogToParent QStyle_StandardPixmap = 30
	QStyle_SP_FileDialogNewFolder QStyle_StandardPixmap = 31
	QStyle_SP_FileDialogDetailedView QStyle_StandardPixmap = 32
	QStyle_SP_FileDialogInfoView QStyle_StandardPixmap = 33
	QStyle_SP_FileDialogContentsView QStyle_StandardPixmap = 34
	QStyle_SP_FileDialogListView QStyle_StandardPixmap = 35
	QStyle_SP_FileDialogBack QStyle_StandardPixmap = 36
	QStyle_SP_DirIcon QStyle_StandardPixmap = 37
	QStyle_SP_DialogOkButton QStyle_StandardPixmap = 38
	QStyle_SP_DialogCancelButton QStyle_StandardPixmap = 39
	QStyle_SP_DialogHelpButton QStyle_StandardPixmap = 40
	QStyle_SP_DialogOpenButton QStyle_StandardPixmap = 41
	QStyle_SP_DialogSaveButton QStyle_StandardPixmap = 42
	QStyle_SP_DialogCloseButton QStyle_StandardPixmap = 43
	QStyle_SP_DialogApplyButton QStyle_StandardPixmap = 44
	QStyle_SP_DialogResetButton QStyle_StandardPixmap = 45
	QStyle_SP_DialogDiscardButton QStyle_StandardPixmap = 46
	QStyle_SP_DialogYesButton QStyle_StandardPixmap = 47
	QStyle_SP_DialogNoButton QStyle_StandardPixmap = 48
	QStyle_SP_ArrowUp QStyle_StandardPixmap = 49
	QStyle_SP_ArrowDown QStyle_StandardPixmap = 50
	QStyle_SP_ArrowLeft QStyle_StandardPixmap = 51
	QStyle_SP_ArrowRight QStyle_StandardPixmap = 52
	QStyle_SP_ArrowBack QStyle_StandardPixmap = 53
	QStyle_SP_ArrowForward QStyle_StandardPixmap = 54
	QStyle_SP_DirHomeIcon QStyle_StandardPixmap = 55
	QStyle_SP_CommandLink QStyle_StandardPixmap = 56
	QStyle_SP_VistaShield QStyle_StandardPixmap = 57
	QStyle_SP_BrowserReload QStyle_StandardPixmap = 58
	QStyle_SP_BrowserStop QStyle_StandardPixmap = 59
	QStyle_SP_MediaPlay QStyle_StandardPixmap = 60
	QStyle_SP_MediaStop QStyle_StandardPixmap = 61
	QStyle_SP_MediaPause QStyle_StandardPixmap = 62
	QStyle_SP_MediaSkipForward QStyle_StandardPixmap = 63
	QStyle_SP_MediaSkipBackward QStyle_StandardPixmap = 64
	QStyle_SP_MediaSeekForward QStyle_StandardPixmap = 65
	QStyle_SP_MediaSeekBackward QStyle_StandardPixmap = 66
	QStyle_SP_MediaVolume QStyle_StandardPixmap = 67
	QStyle_SP_MediaVolumeMuted QStyle_StandardPixmap = 68
	QStyle_SP_CustomBase QStyle_StandardPixmap = 0xf0000000
)
//enum QStyle_StyleHint - QStyle::StyleHint
type QStyle_StyleHint uint32
const (
	QStyle_SH_EtchDisabledText QStyle_StyleHint = 0
	QStyle_SH_DitherDisabledText QStyle_StyleHint = 1
	QStyle_SH_ScrollBar_MiddleClickAbsolutePosition QStyle_StyleHint = 2
	QStyle_SH_ScrollBar_ScrollWhenPointerLeavesControl QStyle_StyleHint = 3
	QStyle_SH_TabBar_SelectMouseType QStyle_StyleHint = 4
	QStyle_SH_TabBar_Alignment QStyle_StyleHint = 5
	QStyle_SH_Header_ArrowAlignment QStyle_StyleHint = 6
	QStyle_SH_Slider_SnapToValue QStyle_StyleHint = 7
	QStyle_SH_Slider_SloppyKeyEvents QStyle_StyleHint = 8
	QStyle_SH_ProgressDialog_CenterCancelButton QStyle_StyleHint = 9
	QStyle_SH_ProgressDialog_TextLabelAlignment QStyle_StyleHint = 10
	QStyle_SH_PrintDialog_RightAlignButtons QStyle_StyleHint = 11
	QStyle_SH_MainWindow_SpaceBelowMenuBar QStyle_StyleHint = 12
	QStyle_SH_FontDialog_SelectAssociatedText QStyle_StyleHint = 13
	QStyle_SH_Menu_AllowActiveAndDisabled QStyle_StyleHint = 14
	QStyle_SH_Menu_SpaceActivatesItem QStyle_StyleHint = 15
	QStyle_SH_Menu_SubMenuPopupDelay QStyle_StyleHint = 16
	QStyle_SH_ScrollView_FrameOnlyAroundContents QStyle_StyleHint = 17
	QStyle_SH_MenuBar_AltKeyNavigation QStyle_StyleHint = 18
	QStyle_SH_ComboBox_ListMouseTracking QStyle_StyleHint = 19
	QStyle_SH_Menu_MouseTracking QStyle_StyleHint = 20
	QStyle_SH_MenuBar_MouseTracking QStyle_StyleHint = 21
	QStyle_SH_ItemView_ChangeHighlightOnFocus QStyle_StyleHint = 22
	QStyle_SH_Widget_ShareActivation QStyle_StyleHint = 23
	QStyle_SH_Workspace_FillSpaceOnMaximize QStyle_StyleHint = 24
	QStyle_SH_ComboBox_Popup QStyle_StyleHint = 25
	QStyle_SH_TitleBar_NoBorder QStyle_StyleHint = 26
	QStyle_SH_Slider_StopMouseOverSlider QStyle_StyleHint = 27
	QStyle_SH_ScrollBar_StopMouseOverSlider QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider
	QStyle_SH_BlinkCursorWhenTextSelected QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1
	QStyle_SH_RichText_FullWidthSelection QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1
	QStyle_SH_Menu_Scrollable QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1
	QStyle_SH_GroupBox_TextLabelVerticalAlignment QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1
	QStyle_SH_GroupBox_TextLabelColor QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1
	QStyle_SH_Menu_SloppySubMenus QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1
	QStyle_SH_Table_GridLineColor QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1
	QStyle_SH_LineEdit_PasswordCharacter QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1
	QStyle_SH_DialogButtons_DefaultButton QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1
	QStyle_SH_ToolBox_SelectedPageTitleBold QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TabBar_PreferNoArrows QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ScrollBar_LeftClickAbsolutePosition QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Q3ListViewExpand_SelectMouseType QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_UnderlineShortcut QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_SpinBox_AnimateButton QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_SpinBox_KeyPressAutoRepeatRate QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_SpinBox_ClickAutoRepeatRate QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Menu_FillScreenWithScroll QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ToolTipLabel_Opacity QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_DrawMenuBarSeparator QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TitleBar_ModifyNotification QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Button_FocusPolicy QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_MenuBar_DismissOnSecondClick QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_MessageBox_UseBorderForButtonSpacing QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TitleBar_AutoRaise QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ToolButton_PopupDelay QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_FocusFrame_Mask QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_RubberBand_Mask QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_WindowFrame_Mask QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_SpinControls_DisableOnBounds QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Dial_BackgroundRole QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ComboBox_LayoutDirection QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_EllipsisLocation QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_ShowDecorationSelected QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_ActivateItemOnSingleClick QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ScrollBar_ContextMenu QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ScrollBar_RollBetweenButtons QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Slider_AbsoluteSetButtons QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Slider_PageSetButtons QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Menu_KeyboardSearch QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TabBar_ElideMode QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_DialogButtonLayout QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ComboBox_PopupFrameStyle QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_MessageBox_TextInteractionFlags QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_DialogButtonBox_ButtonsHaveIcons QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_SpellCheckUnderlineStyle QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_MessageBox_CenterButtons QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Menu_SelectionWrap QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_MovementWithoutUpdatingSelection QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ToolTip_Mask QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_FocusFrame_AboveWidget QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TextControl_FocusIndicatorTextCharFormat QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_WizardStyle QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_ArrowKeysNavigateIntoChildren QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Menu_Mask QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Menu_FlashTriggeredItem QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_Menu_FadeOutOnHide QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_SpinBox_ClickAutoRepeatThreshold QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_PaintAlternatingRowColorsForEmptyArea QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_FormLayoutWrapPolicy QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TabWidget_DefaultTabPosition QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ToolBar_Movable QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_FormLayoutFieldGrowthPolicy QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_FormLayoutFormAlignment QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_FormLayoutLabelAlignment QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ItemView_DrawDelegateFrame QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_TabBar_CloseButtonPosition QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_DockWidget_ButtonsHaveFrame QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_ToolButtonStyle QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_RequestSoftwareInputPanel QStyle_StyleHint = QStyle_SH_Slider_StopMouseOverSlider+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SH_CustomBase QStyle_StyleHint = 0xf0000000
)
//enum QStyle_PrimitiveElement - QStyle::PrimitiveElement
type QStyle_PrimitiveElement uint32
const (
	QStyle_PE_Q3CheckListController QStyle_PrimitiveElement = 0
	QStyle_PE_Q3CheckListExclusiveIndicator QStyle_PrimitiveElement = 1
	QStyle_PE_Q3CheckListIndicator QStyle_PrimitiveElement = 2
	QStyle_PE_Q3DockWindowSeparator QStyle_PrimitiveElement = 3
	QStyle_PE_Q3Separator QStyle_PrimitiveElement = 4
	QStyle_PE_Frame QStyle_PrimitiveElement = 5
	QStyle_PE_FrameDefaultButton QStyle_PrimitiveElement = 6
	QStyle_PE_FrameDockWidget QStyle_PrimitiveElement = 7
	QStyle_PE_FrameFocusRect QStyle_PrimitiveElement = 8
	QStyle_PE_FrameGroupBox QStyle_PrimitiveElement = 9
	QStyle_PE_FrameLineEdit QStyle_PrimitiveElement = 10
	QStyle_PE_FrameMenu QStyle_PrimitiveElement = 11
	QStyle_PE_FrameStatusBar QStyle_PrimitiveElement = 12
	QStyle_PE_FrameStatusBarItem QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar
	QStyle_PE_FrameTabWidget QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1
	QStyle_PE_FrameWindow QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1
	QStyle_PE_FrameButtonBevel QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1
	QStyle_PE_FrameButtonTool QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1
	QStyle_PE_FrameTabBarBase QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1
	QStyle_PE_PanelButtonCommand QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1
	QStyle_PE_PanelButtonBevel QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1
	QStyle_PE_PanelButtonTool QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1
	QStyle_PE_PanelMenuBar QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelToolBar QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelLineEdit QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorArrowDown QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorArrowLeft QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorArrowRight QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorArrowUp QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorBranch QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorButtonDropDown QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorViewItemCheck QStyle_PrimitiveElement = QStyle_PE_FrameStatusBar+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorItemViewItemCheck QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck
	QStyle_PE_IndicatorCheckBox QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1
	QStyle_PE_IndicatorDockWidgetResizeHandle QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1
	QStyle_PE_IndicatorHeaderArrow QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1
	QStyle_PE_IndicatorMenuCheckMark QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1
	QStyle_PE_IndicatorProgressChunk QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1
	QStyle_PE_IndicatorRadioButton QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1
	QStyle_PE_IndicatorSpinDown QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1
	QStyle_PE_IndicatorSpinMinus QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorSpinPlus QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorSpinUp QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorToolBarHandle QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorToolBarSeparator QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelTipLabel QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorTabTear QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelScrollAreaCorner QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_Widget QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorColumnViewArrow QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorItemViewItemDrop QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelItemViewItem QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelItemViewRow QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelStatusBar QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_IndicatorTabClose QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_PanelMenu QStyle_PrimitiveElement = QStyle_PE_IndicatorViewItemCheck+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_PE_CustomBase QStyle_PrimitiveElement = 0xf000000
)
//enum QStyle_ControlElement - QStyle::ControlElement
type QStyle_ControlElement uint32
const (
	QStyle_CE_PushButton QStyle_ControlElement = 0
	QStyle_CE_PushButtonBevel QStyle_ControlElement = 1
	QStyle_CE_PushButtonLabel QStyle_ControlElement = 2
	QStyle_CE_CheckBox QStyle_ControlElement = 3
	QStyle_CE_CheckBoxLabel QStyle_ControlElement = 4
	QStyle_CE_RadioButton QStyle_ControlElement = 5
	QStyle_CE_RadioButtonLabel QStyle_ControlElement = 6
	QStyle_CE_TabBarTab QStyle_ControlElement = 7
	QStyle_CE_TabBarTabShape QStyle_ControlElement = 8
	QStyle_CE_TabBarTabLabel QStyle_ControlElement = 9
	QStyle_CE_ProgressBar QStyle_ControlElement = 10
	QStyle_CE_ProgressBarGroove QStyle_ControlElement = 11
	QStyle_CE_ProgressBarContents QStyle_ControlElement = 12
	QStyle_CE_ProgressBarLabel QStyle_ControlElement = 13
	QStyle_CE_MenuItem QStyle_ControlElement = 14
	QStyle_CE_MenuScroller QStyle_ControlElement = 15
	QStyle_CE_MenuVMargin QStyle_ControlElement = 16
	QStyle_CE_MenuHMargin QStyle_ControlElement = 17
	QStyle_CE_MenuTearoff QStyle_ControlElement = 18
	QStyle_CE_MenuEmptyArea QStyle_ControlElement = 19
	QStyle_CE_MenuBarItem QStyle_ControlElement = 20
	QStyle_CE_MenuBarEmptyArea QStyle_ControlElement = 21
	QStyle_CE_ToolButtonLabel QStyle_ControlElement = 22
	QStyle_CE_Header QStyle_ControlElement = 23
	QStyle_CE_HeaderSection QStyle_ControlElement = 24
	QStyle_CE_HeaderLabel QStyle_ControlElement = 25
	QStyle_CE_Q3DockWindowEmptyArea QStyle_ControlElement = 26
	QStyle_CE_ToolBoxTab QStyle_ControlElement = 27
	QStyle_CE_SizeGrip QStyle_ControlElement = 28
	QStyle_CE_Splitter QStyle_ControlElement = 29
	QStyle_CE_RubberBand QStyle_ControlElement = 30
	QStyle_CE_DockWidgetTitle QStyle_ControlElement = 31
	QStyle_CE_ScrollBarAddLine QStyle_ControlElement = 32
	QStyle_CE_ScrollBarSubLine QStyle_ControlElement = 33
	QStyle_CE_ScrollBarAddPage QStyle_ControlElement = 34
	QStyle_CE_ScrollBarSubPage QStyle_ControlElement = 35
	QStyle_CE_ScrollBarSlider QStyle_ControlElement = 36
	QStyle_CE_ScrollBarFirst QStyle_ControlElement = 37
	QStyle_CE_ScrollBarLast QStyle_ControlElement = 38
	QStyle_CE_FocusFrame QStyle_ControlElement = 39
	QStyle_CE_ComboBoxLabel QStyle_ControlElement = 40
	QStyle_CE_ToolBar QStyle_ControlElement = 41
	QStyle_CE_ToolBoxTabShape QStyle_ControlElement = 42
	QStyle_CE_ToolBoxTabLabel QStyle_ControlElement = 43
	QStyle_CE_HeaderEmptyArea QStyle_ControlElement = 44
	QStyle_CE_ColumnViewGrip QStyle_ControlElement = 45
	QStyle_CE_ItemViewItem QStyle_ControlElement = 46
	QStyle_CE_ShapedFrame QStyle_ControlElement = 47
	QStyle_CE_CustomBase QStyle_ControlElement = 0xf0000000
)
//enum QStyle_ContentsType - QStyle::ContentsType
type QStyle_ContentsType uint32
const (
	QStyle_CT_PushButton QStyle_ContentsType = 0
	QStyle_CT_CheckBox QStyle_ContentsType = 1
	QStyle_CT_RadioButton QStyle_ContentsType = 2
	QStyle_CT_ToolButton QStyle_ContentsType = 3
	QStyle_CT_ComboBox QStyle_ContentsType = 4
	QStyle_CT_Splitter QStyle_ContentsType = 5
	QStyle_CT_Q3DockWindow QStyle_ContentsType = 6
	QStyle_CT_ProgressBar QStyle_ContentsType = 7
	QStyle_CT_MenuItem QStyle_ContentsType = 8
	QStyle_CT_MenuBarItem QStyle_ContentsType = 9
	QStyle_CT_MenuBar QStyle_ContentsType = 10
	QStyle_CT_Menu QStyle_ContentsType = 11
	QStyle_CT_TabBarTab QStyle_ContentsType = 12
	QStyle_CT_Slider QStyle_ContentsType = 13
	QStyle_CT_ScrollBar QStyle_ContentsType = 14
	QStyle_CT_Q3Header QStyle_ContentsType = 15
	QStyle_CT_LineEdit QStyle_ContentsType = 16
	QStyle_CT_SpinBox QStyle_ContentsType = 17
	QStyle_CT_SizeGrip QStyle_ContentsType = 18
	QStyle_CT_TabWidget QStyle_ContentsType = 19
	QStyle_CT_DialogButtons QStyle_ContentsType = 20
	QStyle_CT_HeaderSection QStyle_ContentsType = 21
	QStyle_CT_GroupBox QStyle_ContentsType = 22
	QStyle_CT_MdiControls QStyle_ContentsType = 23
	QStyle_CT_ItemViewItem QStyle_ContentsType = 24
	QStyle_CT_CustomBase QStyle_ContentsType = 0xf0000000
)
//enum QStyle_StateFlag - QStyle::StateFlag
type QStyle_StateFlag uint32
const (
	QStyle_State_None QStyle_StateFlag = 0x00000000
	QStyle_State_Enabled QStyle_StateFlag = 0x00000001
	QStyle_State_Raised QStyle_StateFlag = 0x00000002
	QStyle_State_Sunken QStyle_StateFlag = 0x00000004
	QStyle_State_Off QStyle_StateFlag = 0x00000008
	QStyle_State_NoChange QStyle_StateFlag = 0x00000010
	QStyle_State_On QStyle_StateFlag = 0x00000020
	QStyle_State_DownArrow QStyle_StateFlag = 0x00000040
	QStyle_State_Horizontal QStyle_StateFlag = 0x00000080
	QStyle_State_HasFocus QStyle_StateFlag = 0x00000100
	QStyle_State_Top QStyle_StateFlag = 0x00000200
	QStyle_State_Bottom QStyle_StateFlag = 0x00000400
	QStyle_State_FocusAtBorder QStyle_StateFlag = 0x00000800
	QStyle_State_AutoRaise QStyle_StateFlag = 0x00001000
	QStyle_State_MouseOver QStyle_StateFlag = 0x00002000
	QStyle_State_UpArrow QStyle_StateFlag = 0x00004000
	QStyle_State_Selected QStyle_StateFlag = 0x00008000
	QStyle_State_Active QStyle_StateFlag = 0x00010000
	QStyle_State_Window QStyle_StateFlag = 0x00020000
	QStyle_State_Open QStyle_StateFlag = 0x00040000
	QStyle_State_Children QStyle_StateFlag = 0x00080000
	QStyle_State_Item QStyle_StateFlag = 0x00100000
	QStyle_State_Sibling QStyle_StateFlag = 0x00200000
	QStyle_State_Editing QStyle_StateFlag = 0x00400000
	QStyle_State_KeyboardFocusChange QStyle_StateFlag = 0x00800000
	QStyle_State_ReadOnly QStyle_StateFlag = 0x02000000
	QStyle_State_Small QStyle_StateFlag = 0x04000000
	QStyle_State_Mini QStyle_StateFlag = 0x08000000
)
//enum QStyle_ComplexControl - QStyle::ComplexControl
type QStyle_ComplexControl uint32
const (
	QStyle_CC_SpinBox QStyle_ComplexControl = 0
	QStyle_CC_ComboBox QStyle_ComplexControl = 1
	QStyle_CC_ScrollBar QStyle_ComplexControl = 2
	QStyle_CC_Slider QStyle_ComplexControl = 3
	QStyle_CC_ToolButton QStyle_ComplexControl = 4
	QStyle_CC_TitleBar QStyle_ComplexControl = 5
	QStyle_CC_Q3ListView QStyle_ComplexControl = 6
	QStyle_CC_Dial QStyle_ComplexControl = 7
	QStyle_CC_GroupBox QStyle_ComplexControl = 8
	QStyle_CC_MdiControls QStyle_ComplexControl = 9
	QStyle_CC_CustomBase QStyle_ComplexControl = 0xf0000000
)
//enum QStyle_RequestSoftwareInputPanel - QStyle::RequestSoftwareInputPanel
type QStyle_RequestSoftwareInputPanel uint32
const (
	QStyle_RSIP_OnMouseClickAndAlreadyFocused QStyle_RequestSoftwareInputPanel = 0
	QStyle_RSIP_OnMouseClick QStyle_RequestSoftwareInputPanel = 1
)
//enum QStyle_SubElement - QStyle::SubElement
type QStyle_SubElement uint32
const (
	QStyle_SE_PushButtonContents QStyle_SubElement = 0
	QStyle_SE_PushButtonFocusRect QStyle_SubElement = 1
	QStyle_SE_CheckBoxIndicator QStyle_SubElement = 2
	QStyle_SE_CheckBoxContents QStyle_SubElement = 3
	QStyle_SE_CheckBoxFocusRect QStyle_SubElement = 4
	QStyle_SE_CheckBoxClickRect QStyle_SubElement = 5
	QStyle_SE_RadioButtonIndicator QStyle_SubElement = 6
	QStyle_SE_RadioButtonContents QStyle_SubElement = 7
	QStyle_SE_RadioButtonFocusRect QStyle_SubElement = 8
	QStyle_SE_RadioButtonClickRect QStyle_SubElement = 9
	QStyle_SE_ComboBoxFocusRect QStyle_SubElement = 10
	QStyle_SE_SliderFocusRect QStyle_SubElement = 11
	QStyle_SE_Q3DockWindowHandleRect QStyle_SubElement = 12
	QStyle_SE_ProgressBarGroove QStyle_SubElement = 13
	QStyle_SE_ProgressBarContents QStyle_SubElement = 14
	QStyle_SE_ProgressBarLabel QStyle_SubElement = 15
	QStyle_SE_DialogButtonAccept QStyle_SubElement = 16
	QStyle_SE_DialogButtonReject QStyle_SubElement = 17
	QStyle_SE_DialogButtonApply QStyle_SubElement = 18
	QStyle_SE_DialogButtonHelp QStyle_SubElement = 19
	QStyle_SE_DialogButtonAll QStyle_SubElement = 20
	QStyle_SE_DialogButtonAbort QStyle_SubElement = 21
	QStyle_SE_DialogButtonIgnore QStyle_SubElement = 22
	QStyle_SE_DialogButtonRetry QStyle_SubElement = 23
	QStyle_SE_DialogButtonCustom QStyle_SubElement = 24
	QStyle_SE_ToolBoxTabContents QStyle_SubElement = 25
	QStyle_SE_HeaderLabel QStyle_SubElement = 26
	QStyle_SE_HeaderArrow QStyle_SubElement = 27
	QStyle_SE_TabWidgetTabBar QStyle_SubElement = 28
	QStyle_SE_TabWidgetTabPane QStyle_SubElement = 29
	QStyle_SE_TabWidgetTabContents QStyle_SubElement = 30
	QStyle_SE_TabWidgetLeftCorner QStyle_SubElement = 31
	QStyle_SE_TabWidgetRightCorner QStyle_SubElement = 32
	QStyle_SE_ViewItemCheckIndicator QStyle_SubElement = 33
	QStyle_SE_ItemViewItemCheckIndicator QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator
	QStyle_SE_TabBarTearIndicator QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1
	QStyle_SE_TreeViewDisclosureItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1
	QStyle_SE_LineEditContents QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1
	QStyle_SE_FrameContents QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1
	QStyle_SE_DockWidgetCloseButton QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1
	QStyle_SE_DockWidgetFloatButton QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1
	QStyle_SE_DockWidgetTitleBarText QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1
	QStyle_SE_DockWidgetIcon QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1
	QStyle_SE_CheckBoxLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1
	QStyle_SE_ComboBoxLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_DateTimeEditLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_DialogButtonBoxLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_LabelLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ProgressBarLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_PushButtonLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_RadioButtonLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_SliderLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_SpinBoxLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ToolButtonLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_FrameLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_GroupBoxLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_TabWidgetLayoutItem QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ItemViewItemDecoration QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ItemViewItemText QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ItemViewItemFocusRect QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_TabBarTabLeftButton QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_TabBarTabRightButton QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_TabBarTabText QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ShapedFrameContents QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_ToolBarHandle QStyle_SubElement = QStyle_SE_ViewItemCheckIndicator+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1+1
	QStyle_SE_CustomBase QStyle_SubElement = 0xf0000000
)
//struct QStyle : QStyle
type QStyle struct {
	QObject
}
//QStyle::alignedRect(Qt::LayoutDirection,QFlags<Qt::AlignmentFlag>,QSize const&,QRect const&)	
func QStyleAlignedRect(direction Qt_LayoutDirection,alignment Qt_AlignmentFlag,size *QSize,rectangle *QRect) *QRect {
	var __rv uintptr
	DirectQtDrv(nil,357000,357102,unsafe.Pointer(&direction),unsafe.Pointer(&alignment),Native(size),Native(rectangle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QStyle::alignedRect(Qt::LayoutDirection,QFlags<Qt::AlignmentFlag>,QSize const&,QRect const&)
func (q *QStyle) AlignedRect(direction Qt_LayoutDirection,alignment Qt_AlignmentFlag,size *QSize,rectangle *QRect) *QRect {
	var __rv uintptr
	q.Drv(357000,357102,unsafe.Pointer(&direction),unsafe.Pointer(&alignment),Native(size),Native(rectangle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QStyle::drawItemPixmap(QPainter*,QRect const&,int,QPixmap const&)
func (q *QStyle) DrawItemPixmap(painter *QPainter,rect *QRect,alignment int,pixmap *QPixmap)  {
	q.Drv(357000,357103,Native(painter),Native(rect),unsafe.Pointer(&alignment),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyle::drawItemText(QPainter*,QRect const&,int,QPalette const&,bool,QString const&,QPalette::ColorRole)
func (q *QStyle) DrawItemText(painter *QPainter,rect *QRect,flags int,pal *QPalette,enabled bool,text string,textRole QPalette_ColorRole)  {
	q.Drv(357000,357104,Native(painter),Native(rect),unsafe.Pointer(&flags),Native(pal),unsafe.Pointer(&enabled),unsafe.Pointer(&text),unsafe.Pointer(&textRole),nil,nil,nil,nil,nil)
}	
//QStyle::itemPixmapRect(QRect const&,int,QPixmap const&)
func (q *QStyle) ItemPixmapRect(r *QRect,flags int,pixmap *QPixmap) *QRect {
	var __rv uintptr
	q.Drv(357000,357105,Native(r),unsafe.Pointer(&flags),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QStyle::itemTextRect(QFontMetrics const&,QRect const&,int,bool,QString const&)
func (q *QStyle) ItemTextRect(fm *QFontMetrics,r *QRect,flags int,enabled bool,text string) *QRect {
	var __rv uintptr
	q.Drv(357000,357106,Native(fm),Native(r),unsafe.Pointer(&flags),unsafe.Pointer(&enabled),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QStyle::pixelMetric(QStyle::PixelMetric)
func (q *QStyle) PixelMetric(metric QStyle_PixelMetric) int {
	var __rv int
	q.Drv(357000,357107,unsafe.Pointer(&metric),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::polish(QApplication*)
func (q *QStyle) Polish(value *QApplication)  {
	q.Drv(357000,357108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyle::polish(QPalette&)
func (q *QStyle) PolishWithPalette(value *QPalette)  {
	q.Drv(357000,357109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyle::polish(QWidget*)
func (q *QStyle) PolishWithWidget(value QWidgetInterface)  {
	q.Drv(357000,357110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyle::proxy()
func (q *QStyle) Proxy() *QStyle {
	var __rv uintptr
	q.Drv(357000,357111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QStyle::sliderPositionFromValue(int,int,int,int,bool)	
func QStyleSliderPositionFromValue(min int,max int,val int,space int,upsideDown bool) int {
	var __rv int
	DirectQtDrv(nil,357000,357112,unsafe.Pointer(&min),unsafe.Pointer(&max),unsafe.Pointer(&val),unsafe.Pointer(&space),unsafe.Pointer(&upsideDown),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::sliderPositionFromValue(int,int,int,int,bool)
func (q *QStyle) SliderPositionFromValue(min int,max int,val int,space int,upsideDown bool) int {
	var __rv int
	q.Drv(357000,357112,unsafe.Pointer(&min),unsafe.Pointer(&max),unsafe.Pointer(&val),unsafe.Pointer(&space),unsafe.Pointer(&upsideDown),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::sliderValueFromPosition(int,int,int,int,bool)	
func QStyleSliderValueFromPosition(min int,max int,pos int,space int,upsideDown bool) int {
	var __rv int
	DirectQtDrv(nil,357000,357113,unsafe.Pointer(&min),unsafe.Pointer(&max),unsafe.Pointer(&pos),unsafe.Pointer(&space),unsafe.Pointer(&upsideDown),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::sliderValueFromPosition(int,int,int,int,bool)
func (q *QStyle) SliderValueFromPosition(min int,max int,pos int,space int,upsideDown bool) int {
	var __rv int
	q.Drv(357000,357113,unsafe.Pointer(&min),unsafe.Pointer(&max),unsafe.Pointer(&pos),unsafe.Pointer(&space),unsafe.Pointer(&upsideDown),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::standardIcon(QStyle::StandardPixmap)
func (q *QStyle) StandardIcon(standardIcon QStyle_StandardPixmap) *QIcon {
	var __rv uintptr
	q.Drv(357000,357114,unsafe.Pointer(&standardIcon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QStyle::standardPalette()
func (q *QStyle) StandardPalette() *QPalette {
	var __rv uintptr
	q.Drv(357000,357115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QStyle::standardPixmap(QStyle::StandardPixmap)
func (q *QStyle) StandardPixmap(standardPixmap QStyle_StandardPixmap) *QPixmap {
	var __rv uintptr
	q.Drv(357000,357116,unsafe.Pointer(&standardPixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QStyle::styleHint(QStyle::StyleHint)
func (q *QStyle) StyleHint(stylehint QStyle_StyleHint) int {
	var __rv int
	q.Drv(357000,357117,unsafe.Pointer(&stylehint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::unpolish(QApplication*)
func (q *QStyle) Unpolish(value *QApplication)  {
	q.Drv(357000,357118,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyle::unpolish(QWidget*)
func (q *QStyle) UnpolishWithWidget(value QWidgetInterface)  {
	q.Drv(357000,357119,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyle::visualAlignment(Qt::LayoutDirection,QFlags<Qt::AlignmentFlag>)	
func QStyleVisualAlignment(direction Qt_LayoutDirection,alignment Qt_AlignmentFlag) Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	DirectQtDrv(nil,357000,357120,unsafe.Pointer(&direction),unsafe.Pointer(&alignment),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::visualAlignment(Qt::LayoutDirection,QFlags<Qt::AlignmentFlag>)
func (q *QStyle) VisualAlignment(direction Qt_LayoutDirection,alignment Qt_AlignmentFlag) Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(357000,357120,unsafe.Pointer(&direction),unsafe.Pointer(&alignment),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyle::visualPos(Qt::LayoutDirection,QRect const&,QPoint const&)	
func QStyleVisualPos(direction Qt_LayoutDirection,boundingRect *QRect,logicalPos *QPoint) *QPoint {
	var __rv uintptr
	DirectQtDrv(nil,357000,357121,unsafe.Pointer(&direction),Native(boundingRect),Native(logicalPos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QStyle::visualPos(Qt::LayoutDirection,QRect const&,QPoint const&)
func (q *QStyle) VisualPos(direction Qt_LayoutDirection,boundingRect *QRect,logicalPos *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(357000,357121,unsafe.Pointer(&direction),Native(boundingRect),Native(logicalPos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QStyle::visualRect(Qt::LayoutDirection,QRect const&,QRect const&)	
func QStyleVisualRect(direction Qt_LayoutDirection,boundingRect *QRect,logicalRect *QRect) *QRect {
	var __rv uintptr
	DirectQtDrv(nil,357000,357122,unsafe.Pointer(&direction),Native(boundingRect),Native(logicalRect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QStyle::visualRect(Qt::LayoutDirection,QRect const&,QRect const&)
func (q *QStyle) VisualRect(direction Qt_LayoutDirection,boundingRect *QRect,logicalRect *QRect) *QRect {
	var __rv uintptr
	q.Drv(357000,357122,unsafe.Pointer(&direction),Native(boundingRect),Native(logicalRect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
