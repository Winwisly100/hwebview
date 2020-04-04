package hwebview

import (
	"fmt"
	"log"

	"github.com/asticode/go-astilectron"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "hwebview"

// HwebviewPlugin implements flutter.Plugin and handles method.
type HwebviewPlugin struct {
	Options        astilectron.Options
	WindowsOptions *astilectron.WindowOptions
}

var _ flutter.Plugin = &HwebviewPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *HwebviewPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("openWebview", p.handleWebview)
	return nil
}

func (p *HwebviewPlugin) handleWebview(arguments interface{}) (reply interface{}, err error) {

	url := arguments.(string)
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, p.Options)
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}
	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow(url, p.WindowsOptions); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern
	a.Wait()

	return "go-flutter " + flutter.PlatformVersion, nil
}
