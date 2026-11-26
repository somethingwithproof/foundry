// SPDX-FileCopyrightText: 2026 Thomas Vincent <thomasvincent@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/foundry-ci/foundry/internal/i18n"
)

func TestExtractGlobalLang(t *testing.T) {
	lang, rest := extractGlobalLang([]string{"--lang", "de", "version", "--json"})
	if lang != "de" {
		t.Fatalf("lang = %q, want de", lang)
	}
	if len(rest) != 2 || rest[0] != "version" || rest[1] != "--json" {
		t.Fatalf("rest = %v, want [version --json]", rest)
	}

	lang, rest = extractGlobalLang([]string{"plan", "--lang=ja", "--profile", "ci"})
	if lang != "ja" {
		t.Fatalf("lang = %q, want ja", lang)
	}
	if len(rest) != 3 || rest[0] != "plan" {
		t.Fatalf("rest = %v", rest)
	}
}

func TestPrintUsageLocalized(t *testing.T) {
	t.Cleanup(func() {
		i18n.Init("en")
	})

	i18n.Init("en")
	var enBuf bytes.Buffer
	printUsage(&enBuf)
	if !strings.Contains(enBuf.String(), "Usage: anvil <command> [flags]") {
		t.Fatalf("unexpected English usage:\n%s", enBuf.String())
	}

	i18n.Init("fr")
	var frBuf bytes.Buffer
	printUsage(&frBuf)
	if frBuf.String() == enBuf.String() {
		t.Fatal("expected French usage output to differ from English")
	}
}
