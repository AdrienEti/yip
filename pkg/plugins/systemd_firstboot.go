package plugins

import (
	"fmt"
	"strings"

	"github.com/AdrienEti/yip/pkg/logger"
	"github.com/AdrienEti/yip/pkg/schema"
	"github.com/hashicorp/go-multierror"
	"github.com/twpayne/go-vfs"
)

func SystemdFirstboot(l logger.Interface, s schema.Stage, fs vfs.FS, console Console) error {
	var errs error

	args := []string{}

	for k, v := range s.SystemdFirstBoot {
		args = append(args, fmt.Sprintf("--%s=%s", strings.ToLower(k), v))
	}

	if err := console.RunTemplate(args, "systemd-firstboot %s"); err != nil {
		errs = multierror.Append(errs, err)
	}

	return errs
}
