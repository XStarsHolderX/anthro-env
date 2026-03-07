package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/anthro-env/anthro-env/internal/core"
	"github.com/anthro-env/anthro-env/internal/ui"
)

var version = "dev"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	mgr := core.NewManager()
	if len(args) == 0 {
		return runMenu(mgr)
	}

	switch args[0] {
	case "-v", "--version", "version":
		fmt.Printf("anthro-env %s\n", version)
		return nil
	case "migrate-tokens":
		return runMigrateTokens(mgr)
	case "menu":
		return runMenu(mgr)
	case "init":
		return runInit(mgr)
	case "add", "use", "ls", "current", "rm":
		return runProfile(mgr, args)
	case "doctor":
		return runDoctor(mgr)
	case "hook":
		if len(args) < 2 {
			return errors.New("usage: anthro-env hook <zsh|bash>")
		}
		fmt.Print(core.HookScript(args[1]))
		return nil
	case "env", "export":
		return runExport(mgr)
	case "profile":
		// Backward compatibility for older syntax.
		return runProfile(mgr, args[1:])
	case "help", "-h", "--help":
		printUsage()
		return nil
	default:
		printUsage()
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runInit(mgr *core.Manager) error {
	if err := mgr.EnsureLayout(); err != nil {
		return err
	}

	shell := core.DetectShell(os.Getenv("SHELL"))
	rcFile := core.RCFile(shell)
	if rcFile == "" {
		return fmt.Errorf("unsupported shell: %s", shell)
	}
	if err := core.InstallHook(rcFile, shell); err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Anthro Env initialization")
	fmt.Println("Tip: ANTHROPIC_MODEL is optional. Leave empty to use provider/gateway default model.")
	fmt.Print("Profile name [default]: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "default"
	}
	if !core.ValidProfileName(name) {
		return errors.New("invalid profile name: use letters, numbers, -, _")
	}

	fmt.Print("ANTHROPIC_BASE_URL: ")
	baseURL, _ := reader.ReadString('\n')
	baseURL = strings.TrimSpace(baseURL)

	fmt.Print("ANTHROPIC_MODEL (optional, press Enter to skip): ")
	model, _ := reader.ReadString('\n')
	model = strings.TrimSpace(model)

	fmt.Print("ANTHROPIC_AUTH_TOKEN (stored in Keychain): ")
	token, _ := reader.ReadString('\n')
	token = strings.TrimSpace(token)

	vars := map[string]string{}
	if baseURL != "" {
		vars["ANTHROPIC_BASE_URL"] = baseURL
	}
	if model != "" {
		vars["ANTHROPIC_MODEL"] = model
		vars["ANTHROPIC_SMALL_FAST_MODEL"] = model
		vars["ANTHROPIC_DEFAULT_SONNET_MODEL"] = model
		vars["ANTHROPIC_DEFAULT_OPUS_MODEL"] = model
		vars["ANTHROPIC_DEFAULT_HAIKU_MODEL"] = model
	}

	if err := mgr.SaveProfile(name, vars); err != nil {
		return err
	}
	if token != "" {
		if err := mgr.SaveToken(name, token); err != nil {
			return err
		}
	}
	if err := mgr.UseProfile(name); err != nil {
		return err
	}

	fmt.Printf("Initialized. Active profile: %s\n", name)
	fmt.Printf("Hook installed in: %s\n", rcFile)
	fmt.Printf("Run: source %s\n", rcFile)
	return nil
}

func runMenu(mgr *core.Manager) error {
	profiles, err := mgr.ListProfiles()
	if err != nil {
		return err
	}
	if len(profiles) == 0 {
		return fmt.Errorf("no profiles. run: anthro-env init or anthro-env add <name>")
	}
	sort.Strings(profiles)
	active, _ := mgr.CurrentProfile()

	fmt.Printf("Current profile: %s\n", core.OrDefault(active, "none"))
	fmt.Println("Select a profile:")
	fmt.Println("[0] Exit")
	for i, p := range profiles {
		model, _ := mgr.ProfileModel(p)
		if model == "" {
			model = "-"
		}
		tag := ""
		if p == active {
			tag = " (current"
			if model != "" {
				tag += ", model: " + model
			}
			tag += ")"
		} else {
			tag = " (model: " + model + ")"
		}
		fmt.Printf("[%d] %s%s\n", i+1, p, tag)
	}
	fmt.Print("Enter number: ")
	reader := bufio.NewReader(os.Stdin)
	in, _ := reader.ReadString('\n')
	index, err := ui.ParseMenuSelection(in, len(profiles))
	if err != nil {
		return err
	}
	if index == 0 {
		fmt.Println("Canceled")
		return nil
	}

	name := profiles[index-1]
	if err := mgr.UseProfile(name); err != nil {
		return err
	}
	fmt.Printf("Switched to profile: %s\n", name)
	return nil
}

func runProfile(mgr *core.Manager, args []string) error {
	if len(args) == 0 {
		return errors.New("usage: anthro-env <add|use|ls|current|rm>")
	}

	switch args[0] {
	case "ls":
		profiles, err := mgr.ListProfiles()
		if err != nil {
			return err
		}
		sort.Strings(profiles)
		active, _ := mgr.CurrentProfile()
		for _, p := range profiles {
			if p == active {
				fmt.Printf("%s *\n", p)
			} else {
				fmt.Println(p)
			}
		}
		return nil
	case "current":
		current, err := mgr.CurrentProfile()
		if err != nil || current == "" {
			fmt.Println("none")
			return nil
		}
		fmt.Println(current)
		return nil
	case "use":
		if len(args) < 2 {
			return errors.New("usage: anthro-env use <name>")
		}
		if err := mgr.UseProfile(args[1]); err != nil {
			return err
		}
		fmt.Printf("Switched to profile: %s\n", args[1])
		return nil
	case "add":
		if len(args) < 2 {
			return errors.New("usage: anthro-env add <name>")
		}
		name := args[1]
		if !core.ValidProfileName(name) {
			return errors.New("invalid profile name")
		}
		if err := mgr.EnsureLayout(); err != nil {
			return err
		}
		f := filepath.Join(mgr.ProfilesDir(), name+".env")
		if _, err := os.Stat(f); err == nil {
			return fmt.Errorf("profile exists: %s", name)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("ANTHROPIC_BASE_URL: ")
		baseURL, _ := reader.ReadString('\n')
		baseURL = strings.TrimSpace(baseURL)
		fmt.Println("Tip: ANTHROPIC_MODEL is optional. Leave empty to use provider/gateway default model.")
		fmt.Print("ANTHROPIC_MODEL (optional, press Enter to skip): ")
		model, _ := reader.ReadString('\n')
		model = strings.TrimSpace(model)
		fmt.Print("ANTHROPIC_AUTH_TOKEN (stored in Keychain): ")
		token, _ := reader.ReadString('\n')
		token = strings.TrimSpace(token)
		vars := map[string]string{}
		if baseURL != "" {
			vars["ANTHROPIC_BASE_URL"] = baseURL
		}
		if model != "" {
			vars["ANTHROPIC_MODEL"] = model
			vars["ANTHROPIC_SMALL_FAST_MODEL"] = model
			vars["ANTHROPIC_DEFAULT_SONNET_MODEL"] = model
			vars["ANTHROPIC_DEFAULT_OPUS_MODEL"] = model
			vars["ANTHROPIC_DEFAULT_HAIKU_MODEL"] = model
		}
		if err := mgr.SaveProfile(name, vars); err != nil {
			return err
		}
		if token != "" {
			if err := mgr.SaveToken(name, token); err != nil {
				return err
			}
		}
		fmt.Printf("Added profile: %s\n", name)
		return nil
	case "rm":
		if len(args) < 2 {
			return errors.New("usage: anthro-env rm <name>")
		}
		if err := mgr.RemoveProfile(args[1]); err != nil {
			return err
		}
		fmt.Printf("Removed profile: %s\n", args[1])
		return nil
	default:
		return fmt.Errorf("unknown profile command: %s", args[0])
	}
}

func runDoctor(mgr *core.Manager) error {
	reports := mgr.Doctor()
	for _, r := range reports {
		fmt.Printf("[%s] %s\n", r.Status, r.Message)
	}
	return nil
}

func runMigrateTokens(mgr *core.Manager) error {
	migrated, skipped, err := mgr.MigratePlaintextTokens()
	if err != nil {
		return err
	}
	fmt.Printf("Token migration finished. migrated=%d skipped=%d\n", migrated, skipped)
	if migrated > 0 {
		fmt.Println("Plaintext ANTHROPIC_AUTH_TOKEN has been removed from migrated profile files.")
	}
	return nil
}

func runExport(mgr *core.Manager) error {
	snippet, err := mgr.ExportSnippet()
	if err != nil {
		return err
	}
	fmt.Print(snippet)
	return nil
}

func printUsage() {
	fmt.Println(`anthro-env commands:
  anthro-env -v | --version
  anthro-env migrate-tokens
  anthro-env init
  anthro-env menu
  anthro-env add <name>
  anthro-env use <name>
  anthro-env ls
  anthro-env current
  anthro-env rm <name>
  anthro-env hook <zsh|bash>
  anthro-env doctor`)
}
