/*
COMMAND:
	store cols ...
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var colsGroup string
var listGroup bool

// colsCmd represents the cols command
var colsCmd = &cobra.Command{
	Use:   "cols",
	Short: "Manage your collections.",
	Long: `Manage your collections.

EXAMPLE: 
  store cols add https://www.google.com --group=websites
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cols called")
	},
}

/*
COMMAND:
	store cols find ...
*/
var findColsCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a value from the collection.",
	Long: `Find a value from the collection.
It will return matching strings also not only the exact.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called")
	},
}

/*
COMMAND:
	store cols remove ...
*/
var removeColsCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a value from the collection.",
	Long: `Remove a value from the collection.
The value must exist from the collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(colsCmd)

	/* sub-functions */
	colsCmd.AddCommand(findColsCmd, removeColsCmd)

	/* flags */
	colsCmd.PersistentFlags().StringVarP(&colsGroup, "group", "g", "", "the collection name / group")
	cobra.MarkFlagRequired(findColsCmd.PersistentFlags(), "group")
	cobra.MarkFlagRequired(removeColsCmd.PersistentFlags(), "group")

	colsCmd.Flags().BoolVar(&listGroup, "list", false, "list the current collections created")
}