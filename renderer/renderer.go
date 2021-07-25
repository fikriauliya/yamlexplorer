package renderer

import (
	"errors"
	"fmt"
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

func trim(s string, max int) string {
	return s[:max]
}

func (r *renderer) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	viewName := "table"
	columnSize := 50

	v, err := g.SetView(viewName, 0, 0, maxX, maxY, 0)
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		return err
	}

	y := 0
	for i, header := range r.data.Header {
		v.SetWritePos(i*columnSize, y)
		v.WriteString(trim(leftAlign(header, columnSize-1), columnSize-1))
		v.WriteString("|")
	}

	y = 1
	for i := range r.data.Header {
		v.SetWritePos(i*columnSize, y)
		v.WriteString(strings.Repeat("-", columnSize-1))
		v.WriteString("|")
	}

	y = 2
	for i, row := range r.data.Body {
		for j, cell := range row {
			v.SetWritePos(j*columnSize, y+i)
			v.WriteString(trim(leftAlign(cell, columnSize-1), columnSize-1))
			v.WriteString("|")
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
