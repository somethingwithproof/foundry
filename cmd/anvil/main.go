// SPDX-FileCopyrightText: 2026 Thomas Vincent <thomasvincent@gmail.com>
// SPDX-License-Identifier: Apache-2.0

// Package main implements the anvil CLI for the Foundry CI/CD engine.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/foundry-ci/foundry/internal/config"
	"github.com/foundry-ci/foundry/internal/exec"
	"github.com/foundry-ci/foundry/internal/i18n"
	"github.com/foundry-ci/foundry/internal/plan"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildDate = "unknown"
)

func extractGlobalLang(args []string) (string, []string) {
	var lang string
	var rest []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--" {
			rest = append(rest, args[i:]...)
			break
		}
		if arg == "--lang" && i+1 < len(args) {
			lang = args[i+1]
			i++
			continue
		}
		if strings.HasPrefix(arg, "--lang=") {
			lang = strings.TrimPrefix(arg, "--lang=")
			continue
		}
		rest = append(rest, arg)
	}
	return lang, rest
}

func main() {
	globalLang, args := extractGlobalLang(os.Args[1:])
	i18n.Init(globalLang)

	if len(args) < 1 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	switch args[0] {
	case "version", "--version":
		cmdVersion(args[1:])
	case "doctor":
		cmdDoctor(args[1:])
	case "plan":
		cmdPlan(args[1:])
	case "run":
		cmdRun(args[1:])
	case "help", "--help", "-h":
		printUsage(os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "%s\n\n", i18n.T("err.unknown_command", args[0]))
		printUsage(os.Stderr)
		os.Exit(1)
	}
}

func printUsage(w io.Writer) {
	_, _ = fmt.Fprintf(w, "%s\n\n%s\n\n%s\n  version    %s\n  doctor     %s\n  plan       %s\n  run        %s\n\n%s\n",
		i18n.T("usage.header"),
		i18n.T("usage.syntax"),
		i18n.T("usage.commands"),
		i18n.T("usage.cmd.version"),
		i18n.T("usage.cmd.doctor"),
		i18n.T("usage.cmd.plan"),
		i18n.T("usage.cmd.run"),
		i18n.T("usage.help_hint"),
	)
}

func setupLogger(jsonOutput bool) {
	var handler slog.Handler
	if jsonOutput {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	slog.SetDefault(slog.New(handler))
}

// --- version ---

func cmdVersion(args []string) {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	jsonOut := fs.Bool("json", false, i18n.T("flag.json"))
	lang := fs.String("lang", "", i18n.T("flag.lang"))
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *lang != "" {
		i18n.Init(*lang)
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(map[string]string{
			"version":    version,
			"commit":     commit,
			"build_date": buildDate,
		})
	} else {
		fmt.Println(i18n.T("version.line", version, commit, buildDate))
	}
}

// --- doctor ---

func cmdDoctor(args []string) {
	fs := flag.NewFlagSet("doctor", flag.ContinueOnError)
	configPath := fs.String("config", ".foundry.yaml", i18n.T("flag.config"))
	lang := fs.String("lang", "", i18n.T("flag.lang"))
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *lang != "" {
		i18n.Init(*lang)
	}

	setupLogger(false)
	allPass := true

	// Check 1: config file exists.
	if _, err := os.Stat(*configPath); err != nil {
		fmt.Println(i18n.T("doctor.config_missing", *configPath))
		allPass = false
	} else {
		fmt.Println(i18n.T("doctor.config_exists", *configPath))
	}

	// Check 2: config parses and validates.
	if _, err := config.Load(*configPath); err != nil {
		fmt.Println(i18n.T("doctor.config_invalid", err))
		allPass = false
	} else {
		fmt.Println(i18n.T("doctor.config_valid"))
	}

	// Check 3: go is available.
	if err := exec.CheckTool("go", "version"); err != nil {
		fmt.Println(i18n.T("doctor.go_missing", err))
		allPass = false
	} else {
		fmt.Println(i18n.T("doctor.go_available"))
	}

	if !allPass {
		os.Exit(1)
	}
	fmt.Printf("\n%s\n", i18n.T("doctor.all_passed"))
}

// --- plan ---

func cmdPlan(args []string) {
	fs := flag.NewFlagSet("plan", flag.ContinueOnError)
	profileName := fs.String("profile", "default", i18n.T("flag.profile"))
	configPath := fs.String("config", ".foundry.yaml", i18n.T("flag.config"))
	jsonOut := fs.Bool("json", false, i18n.T("flag.json"))
	lang := fs.String("lang", "", i18n.T("flag.lang"))
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *lang != "" {
		i18n.Init(*lang)
	}

	setupLogger(*jsonOut)

	p, _ := buildAndWritePlan(*configPath, *profileName)

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(p)
	} else {
		fmt.Println(i18n.T("plan.generated", len(p.Steps), *profileName))
		fmt.Println(i18n.T("plan.execution_order"))
		for i, id := range p.Order {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
		fmt.Println(i18n.T("plan.written"))
	}
}

// --- run ---

func cmdRun(args []string) {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	profileName := fs.String("profile", "default", i18n.T("flag.profile"))
	configPath := fs.String("config", ".foundry.yaml", i18n.T("flag.config"))
	jobs := fs.Int("jobs", 4, i18n.T("flag.jobs"))
	jsonOut := fs.Bool("json", false, i18n.T("flag.json"))
	lang := fs.String("lang", "", i18n.T("flag.lang"))
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	if *lang != "" {
		i18n.Init(*lang)
	}

	setupLogger(*jsonOut)

	p, outDir := buildAndWritePlan(*configPath, *profileName)

	// Execute with signal handling.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	opts := exec.DefaultOptions()
	opts.Jobs = *jobs
	opts.OutDir = outDir

	results, err := exec.Execute(ctx, p, opts)
	if err != nil {
		slog.Error(i18n.T("err.execution_failed"), "error", err)
		os.Exit(1)
	}

	if err := exec.WriteResults(results, outDir); err != nil {
		slog.Error(i18n.T("err.write_results"), "error", err)
		os.Exit(1)
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		_ = enc.Encode(results)
	} else {
		fmt.Printf("\n%s\n", i18n.T("run.summary", results.Status, results.Duration))
		for _, sr := range results.Steps {
			marker := "✓"
			if sr.Status != "success" {
				marker = "✗"
			}
			fmt.Printf("  %s %s [%s] %s\n", marker, sr.ID, sr.Status, sr.Duration)
		}
	}

	if results.Status != "success" {
		os.Exit(1)
	}
}

// --- helpers ---

func buildAndWritePlan(configPath, profileName string) (*plan.Plan, string) {
	cfg, steps, configData := loadAndResolve(configPath, profileName)

	for _, s := range steps {
		if err := cfg.Policy.ValidateStep(s.Type, s.ID); err != nil {
			slog.Error(i18n.T("err.policy_violation"), "error", err)
			os.Exit(1)
		}
	}

	p, err := plan.Build(cfg.Project.Name, profileName, steps, configData)
	if err != nil {
		slog.Error(i18n.T("err.build_plan"), "error", err)
		os.Exit(1)
	}

	outDir := ".foundry/out"
	if writeErr := plan.WritePlan(p, outDir); writeErr != nil {
		slog.Error(i18n.T("err.write_plan"), "error", writeErr)
		os.Exit(1)
	}
	return p, outDir
}

// loadAndResolve loads config, resolves the profile, and returns raw config bytes.
func loadAndResolve(configPath, profileName string) (*config.Config, []config.Step, []byte) {
	cfg, err := config.Load(configPath)
	if err != nil {
		slog.Error(i18n.T("err.load_config"), "path", configPath, "error", err)
		os.Exit(1)
	}

	steps, err := config.ResolveProfile(cfg, profileName)
	if err != nil {
		slog.Error(i18n.T("err.resolve_profile"), "profile", profileName, "error", err)
		os.Exit(1)
	}

	configData, err := config.RawBytes(configPath)
	if err != nil {
		slog.Error(i18n.T("err.read_config"), "error", err)
		os.Exit(1)
	}

	return cfg, steps, configData
}
