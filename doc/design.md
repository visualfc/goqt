# Design

## Enum

* Global

		enum Qt::GlobalColor => ui.Qt_GlobalColor
		
		Qt::white => ui.Qt_white
	
* Class enum

		enum QMessageBox::StandardButton => ui.	QMessageBox_StandardButton
		
		QMessageBox::Save => ui.QMessageBox_Save

## Class Name

### QObject base class 

	QWidget => ui.QWidget

* c++

		QWidget *widget = new QWidget();
		widget->show();

* go
	
		widget := ui.NewWidget()
		widget.Show()

###  Plain class

	QPoint => ui.QPoint

* c++
	
		QPoint pt;
		pt.setX(100);
		pt.setY(100);

* go
	
		pt := ui.NewPoint()
		pt.SetX(100)
		pt.SetY(100)
		
### Class Member Function

	widget->show()	=> 	widget.Show()

### Class Static Function

	QMessageBox::about() => ui.QMessageBoxAbout() or &ui.QMessageBox{}.About()
	
## Signal Slot
	
* c++	

		connect(btn,SIGNAL(clicked()),this,SLOT(onClicked()));
	
* go

		btn.OnClicked(func() {
			//fmt.Println("clicked")
		})	
	

## Event

* c++

		class MyWidget : public QWidget
		
		virtual MyWidget::paintEvent(e *QPaintEvent)
		{
			QWidget::paintEvent(e)
			
			QPainter p(this);			
			//p.draw
		}
		
* go

		type MyWidget struct {
			*ui.QWidget
		}	
		
		func NewMyWidget() *MyWidget {
			w := &MyWidget{}
			w.QWidget = ui.NewWidget()
			w.InstallEventFilter(w)
		}	
		
		func (w *MyWidget) OnPaintEvent(e *ui.QPaintEvent) bool {
			w.PaintEvent() //w.QWidget.PaintEvent()
			
			painter := ui.NewPainterWithPaintDevice(w)
			defer painter.Delete()
			
			//painter.Draw ...	
			
			return true
		}
	
## Rename

### New
	QAction::QAction(QObject*) => func NewAction(parent QObjectInterface) *QAction
	QAction::QAction(QString const&,QObject*) => func NewActionWithTextParent(text string,parent QObjectInterface) *QAction
	QAction::QAction(QIcon const&,QString const&,QObject*) => func NewActionWithIconTextParent(icon *QIcon,text string,parent QObjectInterface) *QAction

### Function

	QWidget::resize(QSize const&) => func (q *QWidget) Resize(value *QSize)
	QWidget::resize(int,int) => func (q *QWidget) ResizeWithWidthHeight(w int,h int) 