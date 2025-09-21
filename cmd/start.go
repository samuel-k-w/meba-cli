package cmd

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	watch bool
	debug bool
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the meba application",
	Long:  `Start the meba application in production or development mode with hot reload.`,
	Run: func(cmd *cobra.Command, args []string) {
		if watch {
			startWithAir()
		} else {
			startProduction()
		}
	},
}

func startWithAir() {
	// Check if .air.toml exists
	if _, err := os.Stat(".air.toml"); os.IsNotExist(err) {
		color.Yellow("⚠️  .air.toml not found, creating default configuration...")
		createAirConfig()
	}

	color.Blue("🚀 Starting development server with hot reload...")
	
	cmd := exec.Command("air")
	if debug {
		cmd.Args = append(cmd.Args, "-d")
	}
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		color.Red("Error starting with air: %v", err)
		color.Yellow("Make sure 'air' is installed: go install github.com/cosmtrek/air@latest")
		os.Exit(1)
	}
}

func startProduction() {
	color.Blue("🚀 Starting production server...")
	
	// Build first
	buildCmd := exec.Command("go", "build", "-o", "bin/server", "./cmd/server")
	if err := buildCmd.Run(); err != nil {
		color.Red("Error building application: %v", err)
		os.Exit(1)
	}

	// Run the built binary
	runCmd := exec.Command("./bin/server")
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	runCmd.Stdin = os.Stdin

	if err := runCmd.Run(); err != nil {
		color.Red("Error running application: %v", err)
		os.Exit(1)
	}
}

func createAirConfig() {
	airConfig := `root = "."
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

	if err := os.WriteFile(".air.toml", []byte(airConfig), 0644); err != nil {
		color.Red("Error creating .air.toml: %v", err)
		return
	}
	
	// Create tmp directory
	os.MkdirAll("tmp", 0755)
	color.Green("✅ Created .air.toml configuration")
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVar(&watch, "watch", false, "Start with hot reload using air")
	startCmd.Flags().BoolVar(&debug, "debug", false, "Start in debug mode")
}