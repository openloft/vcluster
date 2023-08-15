package cmd

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/loft-sh/utils/pkg/downloader"
	"github.com/loft-sh/utils/pkg/downloader/commands"
	"github.com/loft-sh/utils/pkg/log"
	"github.com/loft-sh/vcluster/pkg/util/cliconfig"
)

// GetHelmBinaryPath checks for helm binary and downloads if it's not present.
func GetHelmBinaryPath(ctx context.Context, log log.Logger) (string, error) {
	// test for helm
	helmExecutablePath, err := exec.LookPath("helm")
	if err != nil {
		_ = fmt.Errorf("seems like helm is not installed. Helm is required for the creation of a virtual cluster")
		helmExecutablePath, err = downloader.NewDownloader(commands.NewHelmV3Command(), log, cliconfig.VclusterFolder).EnsureCommand(ctx)
		if err != nil {
			return "", fmt.Errorf("error while installing helm: %w", err)
		}
	}
	return helmExecutablePath, nil
}
