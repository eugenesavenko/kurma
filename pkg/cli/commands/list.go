// Copyright 2015-2016 Apcera Inc. All rights reserved.

package commands

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/apcera/kurma/pkg/apiclient"
	"github.com/apcera/kurma/pkg/cli"
	"github.com/apcera/termtables"
	"github.com/spf13/cobra"
)

var (
	ListCmd = &cobra.Command{
		Use:   "list",
		Short: "List running pods",
		Run:   cmdList,
	}
)

func init() {
	cli.RootCmd.AddCommand(ListCmd)
}

func cmdList(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		fmt.Printf("Invalid command options specified.\n")
		os.Exit(1)
	}

	pods, err := cli.GetClient().ListPods()
	if err != nil {
		fmt.Printf("Failed to get list of pods: %v\n", err)
		os.Exit(1)
	}

	// create the table
	table := termtables.CreateTable()

	table.AddHeaders("UUID", "Name", "Apps", "State", "IP(s)")
	sort.Sort(sortedPods(pods))

	for n, pod := range pods {
		ips := make([]string, 0)
		for _, net := range pod.Networks {
			if net.IP4 != nil {
				ips = append(ips, net.IP4.IP.IP.String())
			}
		}

		for i, app := range pod.Pod.Apps {
			if i == 0 {
				table.AddRow(pod.UUID, pod.Name, app.Name.String(), pod.State, strings.Join(ips, " "))
			} else {
				table.AddRow("", "", app.Name.String(), "")
			}
		}
		if n < len(pods)-1 {
			table.AddSeparator()
		}
	}
	fmt.Printf("%s", table.Render())

}

type sortedPods []*apiclient.Pod

func (a sortedPods) Len() int      { return len(a) }
func (a sortedPods) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortedPods) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}
