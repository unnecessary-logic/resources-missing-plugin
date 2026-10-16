package output

import (
	"fmt"
	"io"
	"strings"

	"github.com/unnecessary-logic/resources-missing-plugin/internal/scan"
)

const (
	red   = "\033[31m"
	cyan  = "\033[36m"
	reset = "\033[0m"
)

func ScanningAll(w io.Writer) {
	fmt.Fprintln(w, cyan+"Scanning all namespaces..."+reset)
}

func Discovery(w io.Writer, results []scan.Result) {
	for _, find := range results {
		fmt.Fprintf(w, "%sPod name:%s %s\n%sPod namespace:%s %s\n%sOffending %s:%s %s\n",
			red, reset, find.Pod,
			red, reset, find.Namespace,
			red, "Containers", reset, strings.Join(find.Container, ", "))
	}
}
