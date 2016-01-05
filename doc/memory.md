# Memory Managements

## Go to Qt C++ pointer mapping

* Non QObject based

		pt := ui.NewQPoint() // create new QPoint c++ pointer
		//pt.Delete()		 // delete c++ pointer. 
		//The pt can manual delete or use golang auto gc.
		//When golang auto gc and clean pt, the c++ pointer auto delete.
		
* QObject based	

		widget := ui.NewQWidget()  // create new QWidget c++ pointer
		widget.Delete()			   // delete c++ pointer and children

* System resources

		painter := ui.NewQPainterWithPaintDevice(widget) //create new QPinter c++ pointer, resources must delete
		defer painter.Delete()	//delete c++ pointer and resources
	

## QObject based struct

GoQt QObject memory managements similar to Qt. 

QObjects organize themselves in object trees. When you create a QObject with another object as parent, it's added to the parent's children() list, and is deleted when the parent is. It turns out that this approach fits the needs of GUI objects very well. For example, a QShortcut (keyboard shortcut) is a child of the relevant window, so when the user closes that window, the shorcut is deleted too.

QWidget, the base class of everything that appears on the screen, extends the parent-child relationship. A child normally also becomes a child widget, i.e. it is displayed in its parent's coordinate system and is graphically clipped by its parent's boundaries. For example, when the application deletes a message box after it has been closed, the message box's buttons and label are also deleted, just as we'd want, because the buttons and label are children of the message box.

You can also delete child objects yourself, and they will remove themselves from their parents. For example, when the user removes a toolbar it may lead to the application deleting one of its QToolBar objects, in which case the tool bar's QMainWindow parent would detect the change and reconfigure its screen space accordingly.

## Non QObject struct

Non QObject struct auto gc or manual management.

If struct hold memory only, this can golang gc or manually manage.
If struct hold system resources, you need to manually manage and delete.

* Auto gc for memory only

		pt := ui.NewQPoint() //this can auto gc
		//defer pt.Delete()  
		
* Manual resources managements	

		painter := ui.NewQPainterWithPaintDevice(widget) //must delete
		defer painter.Delete()
 
