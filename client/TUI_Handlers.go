package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func TUI_Init() {
    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }
    if err := s.Init(); err != nil {
        log.Fatalf("%+v", err)
    }

    defStyle := tcell.StyleDefault
    s.SetStyle(defStyle)
    s.Clear()
    s.SetContent(0, 0, 'H', nil, defStyle)
    s.SetContent(1, 0, 'i', nil, defStyle)
    s.SetContent(2, 0, '!', nil, defStyle)
    for {
        // Update screen
        s.Show()

        // Poll event
        ev := s.PollEvent()

        // Process event
        switch ev := ev.(type) {
        case *tcell.EventResize:
            s.Sync()
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                s.Fini() //get the terminal out of raw mode
                os.Exit(0)
            }
        }
    }
}
