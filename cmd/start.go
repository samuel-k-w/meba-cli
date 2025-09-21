package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	watchFlag bool
	debugFlag bool
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Long:  "Start the application in production mode or with live-reload",
	Run: func(cmd *cobra.Command, args []string) {
		if watchFlag {
			startWithWatch()
		} else {
			startProduction()
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&watchFlag, "watch", "w", false, "Start with live-reload using Air")
	startCmd.Flags().BoolVar(&debugFlag, "debug", false, "Enable debug mode")
}

func startProduction() {
	fmt.Println("🚀 Starting in production mode...")
	
	// Build first
	if err := buildApp(); err != nil {
		fmt.Printf("❌ Build failed: %v\n", err)
		os.Exit(1)
	}
	
	// Run the binary
	binaryPath := "./server"
	if _, err := os.Stat("./dist/server"); err == nil {
		binaryPath = "./dist/server"
	}
	
	cmd := exec.Command(binaryPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func startWithWatch() {
	fmt.Println("🔥 Starting with live-reload...")
	
	// Check if .air.toml exists, create if not
	if _, err := os.Stat(".air.toml"); os.IsNotExist(err) {
		createAirConfig()
	}
	
	// Auto-generate wire before starting
	fmt.Println("🔧 Generating wire dependencies...")
	wireCmd := exec.Command("wire", "./internal")
	wireCmd.Stdout = os.Stdout
	wireCmd.Stderr = os.Stderr
	wireCmd.Run() // Don't fail if wire fails
	
	// Start air
	args := []string{}
	if debugFlag {
		fmt.Println("🐛 Debug mode enabled")
		args = append(args, "-d")
	}
	
	cmd := exec.Command("air", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Failed to start with Air: %v\n", err)
		fmt.Println("💡 Make sure Air is installed: go install github.com/cosmtrek/air@latest")
		os.Exit(1)
	}
}

func createAirConfig() {
	config := `root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/server"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
`
	
	if err := os.WriteFile(".air.toml", []byte(config), 0644); err != nil {
		fmt.Printf("Warning: Could not create .air.toml: %v\n", err)
	}
}