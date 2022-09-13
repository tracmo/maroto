package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tracmo/maroto/pkg/consts"
	"github.com/tracmo/maroto/pkg/pdf"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)
	m.SetMaxGridSum(36)

	m.Row(40, func() {
		m.Col(2, func() {
			m.Text("Any Text1")
		})
		m.Col(4, func() {
			m.Text("Any Text2")
		})
		m.Col(6, func() {
			m.Text("Any Text3")
		})
		m.Col(10, func() {
			m.Text("Any Text4")
		})
		m.Col(14, func() {
			m.Text("Any Text5")
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/textgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
