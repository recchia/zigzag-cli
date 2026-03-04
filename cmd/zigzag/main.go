package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/recchia/zigzag-cli/internal/downloader"
	"github.com/recchia/zigzag-cli/internal/eas"
	"github.com/recchia/zigzag-cli/internal/ui"
	"github.com/spf13/cobra"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		// We don't want to fail if .env is missing, as the variable might be set in the environment
	}

	var rootCmd = &cobra.Command{
		Use:   "zigzag",
		Short: "Zigzag is a CLI to download EAS builds",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Fetching builds from EAS...")
			runner := eas.RealCommandRunner{}
			builds, err := eas.ListBuilds(runner)
			if err != nil {
				log.Fatalf("Error fetching builds: %v", err)
			}

			selectedBuild, err := ui.SelectBuild(builds)
			if err != nil {
				log.Fatalf("Error selecting build: %v", err)
			}

			if selectedBuild.Artifacts.BuildUrl == "" {
				log.Fatal("Selected build has no artifact URL")
			}

			outputDir := os.Getenv("ZIGZAG_OUTPUT_DIR")
			if outputDir == "" {
				outputDir = "." // Fallback to current directory
			}

			ext := ".apk"
			if selectedBuild.Platform == "ios" {
				ext = ".tar.gz" // or .ipa depending on distribution
			}

			defaultFileName := fmt.Sprintf("%s-%s%s", selectedBuild.Platform, selectedBuild.BuildProfile, ext)
			fileName, err := ui.AskFileName(defaultFileName)
			if err != nil {
				log.Fatalf("Error getting file name: %v", err)
			}

			fmt.Printf("Downloading build to %s...\n", outputDir)
			finalPath, err := downloader.Download(selectedBuild.Artifacts.BuildUrl, outputDir, fileName)
			if err != nil {
				log.Fatalf("Error downloading build: %v", err)
			}

			fmt.Printf("Successfully downloaded to: %s\n", finalPath)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
