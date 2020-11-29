package main

import (
	wails "github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

func main() {

	// Create menu
	myMenu := menu.DefaultMacMenu()

	windowMenu := menu.SubMenu("Test", []*menu.MenuItem{
		menu.Togglefullscreen(),
		menu.Minimize(),
		menu.Zoom(),

		menu.Separator(),

		menu.Copy(),
		menu.Cut(),
		menu.Delete(),

		menu.Separator(),

		menu.Front(),

		menu.SubMenu("Test Submenu", []*menu.MenuItem{
			menu.SubMenu("Accelerators", []*menu.MenuItem{
				menu.TextWithAccelerator("Shift accelerator", "Shift", menu.ShiftAccel("o")),
				menu.TextWithAccelerator("Control accelerator", "Control", menu.ControlAccel("o")),
				menu.TextWithAccelerator("Command accelerator", "Command", menu.CmdOrCtrlAccel("o")),
				menu.TextWithAccelerator("Option accelerator", "Option", menu.OptionOrAltAccel("o")),
				menu.TextWithAccelerator("Backspace", "Backspace", menu.Accel("Backspace")),
				menu.TextWithAccelerator("Tab", "Tab", menu.Accel("Tab")),
				menu.TextWithAccelerator("Return", "Return", menu.Accel("Return")),
				menu.TextWithAccelerator("Escape", "Escape", menu.Accel("Escape")),
				menu.TextWithAccelerator("Left", "Left", menu.Accel("Left")),
				menu.TextWithAccelerator("Right", "Right", menu.Accel("Right")),
				menu.TextWithAccelerator("Up", "Up", menu.Accel("Up")),
				menu.TextWithAccelerator("Down", "Down", menu.Accel("Down")),
				menu.TextWithAccelerator("Space", "Space", menu.Accel("Space")),
				menu.TextWithAccelerator("Delete", "Delete", menu.Accel("Delete")),
				menu.TextWithAccelerator("Home", "Home", menu.Accel("Home")),
				menu.TextWithAccelerator("End", "End", menu.Accel("End")),
				menu.TextWithAccelerator("Page Up", "Page Up", menu.Accel("Page Up")),
				menu.TextWithAccelerator("Page Down", "Page Down", menu.Accel("Page Down")),
			}),
			&menu.MenuItem{
				Label:       "Disabled Menu",
				Type:        menu.TextType,
				Accelerator: menu.ComboAccel("p", menu.CmdOrCtrl, menu.Shift),
				Disabled:    true,
			},
			&menu.MenuItem{
				Label:  "Hidden Menu",
				Type:   menu.TextType,
				Hidden: true,
			},
			&menu.MenuItem{
				ID:          "checkbox-menu",
				Label:       "Checkbox Menu",
				Type:        menu.CheckboxType,
				Accelerator: menu.CmdOrCtrlAccel("l"),
				Checked:     true,
			},
			menu.Separator(),
			menu.Radio("😀 Option 1", "😀option-1", true),
			menu.Radio("😺 Option 2", "option-2", false),
			menu.Radio("❤️ Option 3", "option-3", false),
		}),
	})

	myMenu.Append(windowMenu)

	// Create application with options
	app := wails.CreateAppWithOptions(&options.App{
		Title:     "Kitchen Sink",
		Width:     1024,
		Height:    768,
		MinWidth:  800,
		MinHeight: 600,
		Mac: &mac.Options{
			WebviewIsTransparent:          true,
			WindowBackgroundIsTranslucent: true,
			// Comment out line below to see Window.SetTitle() work
			TitleBar: mac.TitleBarHiddenInset(),
			Menu:     myMenu,
		},
		LogLevel: logger.TRACE,
	})

	app.Bind(&Events{})
	app.Bind(&Logger{})
	app.Bind(&Browser{})
	app.Bind(&System{})
	app.Bind(&Dialog{})
	app.Bind(&Window{})

	app.Run()
}
