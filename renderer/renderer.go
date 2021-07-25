package renderer

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/fikriauliya/yamlexplorer/entity"
)

type Renderer interface {
	Render(table *entity.Table) error
}

type renderer struct {
	data *entity.Table
}

func NewRenderer() Renderer {
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	return &renderer{}
}

func (r *renderer) Render(table *entity.Table) error {
	r.data = table

	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManager(r)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		return err
	}
	return nil
}

func leftAlign(s string, width int) string {
	return fmt.Sprintf("%-*s", width, s)
}

func (r *renderer) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	viewName := "table"

	widths, err := Resize(r.data, maxX-(len(r.data.Header)*2)-1) // space for <space>|
	if err != nil {
		return err
	}

	v, err := g.SetView(viewName, 0, 0, maxX, maxY, 0)
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		return err
	}

	y := 0
	v.SetWritePos(0, y)

	for i, header := range r.data.Header {
		v.WriteString(trim(leftAlign(header, widths[i]), widths[i]))
		v.WriteString(" |")
	}

	y = 1
	v.SetWritePos(0, y)
	for i := range r.data.Header {
		if widths[i] > 0 {
			v.WriteString(strings.Repeat("-", widths[i]))
		}
		v.WriteString(" |")
	}

	y = 2
	for i, row := range r.data.Body {
		v.SetWritePos(0, y+i)
		for j, cell := range row {
			v.WriteString(trim(leftAlign(cell, widths[j]), widths[j]))
			v.WriteString(" |")
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
