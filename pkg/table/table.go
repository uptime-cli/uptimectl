package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

func Print(header []string, body [][]string) {
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
			Borders: tw.BorderNone,
			Symbols: tw.NewSymbolCustom("kube").WithColumn("\t"),
			Settings: tw.Settings{
				Lines:      tw.LinesNone,
				Separators: tw.Separators{BetweenColumns: tw.On},
			},
		})),
		tablewriter.WithPadding(tw.PaddingNone),
		tablewriter.WithHeaderAlignment(tw.AlignLeft),
		tablewriter.WithRowAlignment(tw.AlignLeft),
		tablewriter.WithHeaderAutoWrap(tw.WrapNone),
		tablewriter.WithRowAutoWrap(tw.WrapNone),
		tablewriter.WithHeaderAutoFormat(tw.On),
	)

	if len(header) > 0 {
		table.Header(header)
	}
	if len(body) > 0 {
		_ = table.Bulk(body)
	}

	_ = table.Render()
}
