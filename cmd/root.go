package cmd

import (
	"os"

	"celvercel/pkg/helpers"
	"celvercel/pkg/prompt"
	"celvercel/pkg/styles"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "celvercel",
	Short: "Oreo meme but with Vercel logo",
	Long:  "Generate the Oreo meme but with the Vercel logo. You can save the result to a file with the -o flag.",
	Run:   runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fileName, error := cmd.Flags().GetString("output")

	if error != nil {
		println(styles.ErrorStyle.Render("An error occurred while trying to get the output file name"))
		os.Exit(1)
	}

	result, words := prompt.CreateForm()

	if fileName != "" {
		helpers.SaveResultToFile(fileName, words, result)
		println("")
		println(styles.SuccessStyle.Render("Result saved to " + fileName))
	}

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("output", "o", "", "Save the result to a file")
}
