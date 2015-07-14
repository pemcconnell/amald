package reports

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
	"github.com/pemcconnell/amald/defs"
)

type ReportAscii struct{}

var (
	buffer bytes.Buffer
)

// Generate creates an HTML string with all the required data
// in place
func (r *ReportAscii) Generate(summaries defs.Summaries) (string, error) {
	log.Debug(summaries)
	output := "\n[ SUMMARIES ]\n"
	table := tablewriter.NewWriter(&buffer)
	table.SetHeader([]string{"URL", "LockedDown", "Status Code", "Status"})
	table.Render()
	output += buffer.String()
	return output, nil
}
