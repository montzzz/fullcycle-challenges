package report

import (
	"fmt"
	"io"
	"os"

	"github.com/montzzzzz/challenges/stress-test/internal/dto"
)

func Print(r dto.Report) {
	PrintTo(r, os.Stdout)
}

func PrintTo(r dto.Report, w io.Writer) {
	fmt.Fprintln(w, "===== Stress Test Report =====")
	fmt.Fprintf(w, "Total execution time: %v\n", r.TotalTime)
	fmt.Fprintf(w, "Total requests: %d\n", r.Total)
	fmt.Fprintf(w, "Successful (200): %d\n", r.Success)
	fmt.Fprintln(w, "Status code distribution:")

	for code, count := range r.ByStatus {
		fmt.Fprintf(w, "  %d: %d\n", code, count)
	}
	fmt.Fprintln(w, "============================")
}
