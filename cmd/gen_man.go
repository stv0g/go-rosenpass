// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func genMan(cmd *cobra.Command, args []string) error {
	if err := os.MkdirAll(genManOpts.Path, 0o755); err != nil {
		return err
	}

	if err := doc.GenManTreeFromOpts(rootCmd, genManOpts); err != nil {
		return err
	}

	return nil
}
