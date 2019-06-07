package resizeimg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

const (
	dstHead = "resize_"
)

func Resizer(srcdir, dstdir string, width, height uint) error {
	dirs := []string{srcdir, dstdir}
	for _, dir := range dirs {
		if fi, err := os.Stat(dir); err != nil {
			return err
		} else if !fi.IsDir() {
			return fmt.Errorf("%s is not a directory", dir)
		}
	}

	fis, err := ioutil.ReadDir(srcdir)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}

		fn := fi.Name()

		switch ext := strings.ToLower(filepath.Ext(fn)); ext {
		case ".gif":
			srcimg, err := decgif(filepath.Join(srcdir, fn))
			if err != err {
				return err
			}
			dstimg := resize.Resize(width, height, srcimg, resize.Lanczos3)
			if err = encgif(filepath.Join(dstdir, dstHead+fn), dstimg); err != nil {
				return err
			}
		case ".jpg":
			srcimg, err := decjpg(filepath.Join(srcdir, fn))
			if err != err {
				return err
			}
			dstimg := resize.Resize(width, height, srcimg, resize.Lanczos3)
			if err = encjpg(filepath.Join(dstdir, dstHead+fn), dstimg); err != nil {
				return err
			}
		case ".png":
			srcimg, err := decpng(filepath.Join(srcdir, fn))
			if err != err {
				return err
			}
			dstimg := resize.Resize(width, height, srcimg, resize.Lanczos3)
			if err = encpng(filepath.Join(dstdir, dstHead+fn), dstimg); err != nil {
				return err
			}
		}
	}

	return nil
}
