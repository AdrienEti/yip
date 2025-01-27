package plugins

import (
	"path/filepath"
	"strings"

	"github.com/AdrienEti/yip/pkg/logger"
	"github.com/AdrienEti/yip/pkg/schema"
	"github.com/hashicorp/go-multierror"
	"github.com/twpayne/go-vfs"
)

var (
	procSys = []string{"/proc", "sys"}
)

func Sysctl(l logger.Interface, s schema.Stage, fs vfs.FS, console Console) error {
	var errs error
	for k, v := range s.Sysctl {
		elements := procSys
		elements = append(elements, strings.Split(k, ".")...)
		path := filepath.Join(elements...)
		if err := fs.WriteFile(path, []byte(v), 0644); err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}
