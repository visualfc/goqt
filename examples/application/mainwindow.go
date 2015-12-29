// mainwindow
package main

import "github.com/visualfc/goqt/ui"

type MainWindow struct {
	*ui.QMainWindow
	edit    *ui.QPlainTextEdit
	curFile string
}

func NewMainWindow() *MainWindow {
	w := &MainWindow{}
	w.QMainWindow = ui.NewMainWindow()
	w.InstallEventFilter(w)

	w.edit = ui.NewPlainTextEdit()
	w.SetCentralWidget(w.edit)

	w.createActions()

	w.readSettings()

	w.edit.Document().OnContentsChanged(w.documentWasModified)
	w.setCurrentFile("")

	return w
}

func (w *MainWindow) readSettings() {
	setting := ui.NewSettingsWithOrganizationApplicationParent("GoQt", "Application Example", w)
	defer setting.Delete()
	pos := setting.ValueWithKeyDefaultvalue("pos", ui.NewVariantWithPoint(ui.NewPointWithXposYpos(200, 200))).ToPoint()
	size := setting.ValueWithKeyDefaultvalue("size", ui.NewVariantWithSize(ui.NewSizeWithWidthHeight(400, 400))).ToSize()
	w.Move(pos)
	w.Resize(size)
}

func (w *MainWindow) writeSettings() {
	setting := ui.NewSettingsWithOrganizationApplicationParent("GoQt", "Application Example", nil)
	setting.SetValue("pos", ui.NewVariantWithPoint(w.Pos()))
	setting.SetValue("size", ui.NewVariantWithSize(w.Size()))
}

func (w *MainWindow) maybeSave() bool {
	if w.edit.Document().IsModified() {
		ret := ui.QMessageBoxWarningWithParentTitleTextButtonsStandardbutton(w, "Application",
			"The document has been modified.\nDo youwant to save your changed?",
			ui.QMessageBox_Save|ui.QMessageBox_Discard|ui.QMessageBox_Cancel,
			ui.QMessageBox_Save)
		if ret == ui.QMessageBox_Save {
			return w.save()
		} else if ret == ui.QMessageBox_Cancel {
			return false
		}
	}
	return true
}

func (w *MainWindow) save() bool {
	if w.curFile == "" {
		return w.saveAs()
	}
	return w.saveFile(w.curFile)
}

func (w *MainWindow) saveAs() bool {
	fileName := ui.QFileDialogGetSaveFileName()
	if fileName == "" {
		return false
	}
	return w.saveFile(fileName)
}

func (w *MainWindow) saveFile(fileName string) bool {
	file := ui.NewFileWithName(fileName)
	defer file.Delete()
	if !file.Open(ui.QIODevice_WriteOnly | ui.QIODevice_Text) {
		return false
	}
	file.Write([]byte(w.edit.ToPlainText()))
	w.setCurrentFile(fileName)
	w.StatusBar().ShowMessageWithTextTimeout("File saved", 2000)
	return true
}

func (w *MainWindow) OnCloseEvent(ev *ui.QCloseEvent) bool {
	if w.maybeSave() {
		w.writeSettings()
		ev.Accept()
	} else {
		ev.Ignore()
	}
	return true
}

func (w *MainWindow) createActions() {
	newAct := ui.NewActionWithTextParent(w.Tr("&New"), w)
	newAct.SetIcon(ui.NewIconWithFilename(":/images/new.png"))
	newAct.SetShortcuts(ui.QKeySequence_New)
	newAct.SetStatusTip("Create a new file")
	newAct.OnTriggered(func() { w.newFile() })

	openAct := ui.NewActionWithTextParent("&Open...", w)
	openAct.SetIcon(ui.NewIconWithFilename(":/images/open.png"))
	openAct.SetShortcuts(ui.QKeySequence_Open)
	openAct.SetStatusTip("Open an existing file")
	openAct.OnTriggered(func() { w.open() })

	saveAct := ui.NewActionWithTextParent("&Save", w)
	saveAct.SetIcon(ui.NewIconWithFilename(":/images/save.png"))
	saveAct.SetShortcuts(ui.QKeySequence_Save)
	saveAct.SetStatusTip("Save the document to disk")
	saveAct.OnTriggered(func() { w.save() })

	saveAsAct := ui.NewActionWithTextParent("&SaveAs...", w)
	saveAsAct.SetShortcuts(ui.QKeySequence_SaveAs)
	saveAsAct.SetStatusTip("Save the document under a new name")
	saveAsAct.OnTriggered(func() { w.saveAs() })

	exitAct := ui.NewActionWithTextParent("&Exit", w)
	exitAct.SetShortcuts(ui.QKeySequence_Quit)
	exitAct.SetStatusTip("Exit the application")
	exitAct.OnTriggered(func() { w.Close() })

	copyAct := ui.NewActionWithTextParent("&Copy", w)
	copyAct.SetIcon(ui.NewIconWithFilename(":/images/copy.png"))
	copyAct.SetShortcuts(ui.QKeySequence_Copy)
	copyAct.SetStatusTip("Copy the current selection's contents to the clipboard")
	copyAct.OnTriggered(func() { w.edit.Copy() })
	w.edit.OnCopyAvailable(func(b bool) { copyAct.SetEnabled(b) })

	cutAct := ui.NewActionWithTextParent("&Cut", w)
	cutAct.SetIcon(ui.NewIconWithFilename(":/images/cut.png"))
	cutAct.SetShortcuts(ui.QKeySequence_Cut)
	cutAct.SetStatusTip("Cut the current selection's contents to the clipboard")
	cutAct.OnTriggered(func() { w.edit.Cut() })
	w.edit.OnCopyAvailable(func(b bool) { cutAct.SetEnabled(b) })

	pasteAct := ui.NewActionWithTextParent("&Paste", w)
	pasteAct.SetIcon(ui.NewIconWithFilename(":/images/paste.png"))
	pasteAct.SetShortcuts(ui.QKeySequence_Paste)
	pasteAct.SetStatusTip("Paste the clipboard's contents into the current selection")
	pasteAct.OnTriggered(func() { w.edit.Paste() })

	aboutAct := ui.NewActionWithTextParent("&About", w)
	aboutAct.SetStatusTip("Show the application's About box")
	aboutAct.OnTriggered(func() {
		ui.QMessageBoxAbout(w, "About Application",
			w.Tr("The <b>Application</b> example demonstrates how to "+
				"write modern GUI applications using Qt, with a menu bar, "+
				"toolbars, and a status bar."))
	})
	aboutQtAct := ui.NewActionWithTextParent("About &Qt", w)
	aboutQtAct.SetStatusTip("Show the Qt library's About box")
	aboutQtAct.OnTriggered(func() { ui.QApplicationAboutQt() })

	fileMenu := w.MenuBar().AddMenuWithTitle("&File")
	fileMenu.AddAction(newAct)
	fileMenu.AddAction(openAct)
	fileMenu.AddAction(saveAct)
	fileMenu.AddAction(saveAsAct)
	fileMenu.AddSeparator()
	fileMenu.AddAction(exitAct)

	editMenu := w.MenuBar().AddMenuWithTitle("&Edit")
	editMenu.AddAction(copyAct)
	editMenu.AddAction(cutAct)
	editMenu.AddAction(pasteAct)

	helpMenu := w.MenuBar().AddMenuWithTitle("&Help")
	helpMenu.AddAction(aboutAct)
	helpMenu.AddSeparator()
	helpMenu.AddAction(aboutQtAct)

	fileToolBar := w.AddToolBar("File")
	fileToolBar.AddAction(newAct)
	fileToolBar.AddAction(openAct)
	fileToolBar.AddAction(saveAct)

	editToolBar := w.AddToolBar("Edit")
	editToolBar.AddAction(copyAct)
	editToolBar.AddAction(cutAct)
	editToolBar.AddAction(pasteAct)

	copyAct.SetEnabled(false)
	cutAct.SetEnabled(false)

	w.StatusBar().ShowMessage("Ready")
}

func (w *MainWindow) newFile() {
	if w.maybeSave() {
		w.edit.Clear()
		w.setCurrentFile("")
	}
}

func (w *MainWindow) open() {
	if w.maybeSave() {
		filename := ui.QFileDialogGetOpenFileName()
		if filename != "" {
			w.loadFile(filename)
		}
	}
}

func (w *MainWindow) loadFile(filename string) {
	file := ui.NewFileWithName(filename)
	defer file.Delete()

	if !file.Open(ui.QIODevice_ReadOnly) {
		return
	}
	w.edit.SetPlainText(string(file.ReadAll()))
	w.setCurrentFile(filename)
	w.StatusBar().ShowMessageWithTextTimeout("File loaded", 2000)
}

func (w *MainWindow) documentWasModified() {
	w.SetWindowModified(w.edit.Document().IsModified())
}

func (w *MainWindow) setCurrentFile(filename string) {
	w.curFile = filename
	w.edit.Document().SetModified(false)
	w.SetWindowModified(false)
	shownName := w.curFile
	if shownName == "" {
		shownName = "untitled.txt"
	}
	w.SetWindowFilePath(shownName)
}
