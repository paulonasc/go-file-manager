package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files",
	Short: "Lists files in directory.",
	Long:  `Lists all files in current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set colors
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()

		fostatus, _ := cmd.Flags().GetBool("fo")

		// Reads from current directory hence the ./
		// files returns a slice ([]fs.FileInfo) of memory addresses
		// files[0] returns a single file and it's information containing memory addresses
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}
		sort.Slice(files, func(i, j int) bool {
			return files[i].Size() > files[j].Size()
		})

		for _, f := range files {
			if fostatus {
				fmt.Printf("%s %s bytes \n", green(f.Name()), red(f.Size()))
			} else {
				if f.IsDir() {
					fmt.Printf("%s ", blue("DIRECTORY"))
				}
				fmt.Printf("%s %s bytes \n", green(f.Name()), red(f.Size()))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listFilesCmd.Flags().BoolP("fo", "f", false, "Toggle for files only.")
}

// list-files: all files & directories
// list-files -f: files only
// list-files
