package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (f *Files) TableOutput(expandedPath string) {
	f.FoundAllFilesInDir(expandedPath)

	const (
		firstColor  = lipgloss.Color("5")
		secondColor = lipgloss.Color("240")
		thirdColor  = lipgloss.Color("250")
	)

	re := lipgloss.NewRenderer(os.Stdout)

	fileNameLen := 0
	for _, name := range f.Name {
		if len(filepath.Base(name)) > fileNameLen {
			fileNameLen = len(filepath.Base(name)) + 4
		}
	}

	var (
		HeaderStyle  = re.NewStyle().Foreground(firstColor).Bold(true).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 1).Width(fileNameLen)
		OddRowStyle  = CellStyle.Foreground(secondColor)
		EvenRowStyle = CellStyle.Foreground(thirdColor)
		BorderStyle  = lipgloss.NewStyle().Foreground(firstColor)
	)

	var rows [][]string
	for i, name := range f.Name {
		rows = append(rows, []string{filepath.Base(name), fmt.Sprintf("%d", f.Lines[i])})
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			if col == 1 {
				style = style.Width(22)
			}

			return style
		}).
		Headers("File name", "Lines").
		Rows(rows...)

	fmt.Println(t)
	fmt.Println("Total lines: ", f.TotalLines)
}

func (f *Files) TreeOutput(expandedPath string) {
	f.FoundAllFilesInDir(expandedPath)
}