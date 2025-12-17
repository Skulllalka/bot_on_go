package files

import (
	"encoding/gob"
	"os"
	"path/filepath"

	e "github.com/Skulllalka/bot_on_go/lib"
	"github.com/Skulllalka/bot_on_go/storage"
)

const (
	defaultPerm = 0774
)

type Storage struct {
	basePath string
}

func New(path string) *Storage {
	return &Storage{
		basePath: path,
	}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() {
		err = e.WrapIfErr("can't save", err)
	}()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(page)
	if err != nil {
		return err
	}
	filePath = filepath.Join(filePath, fName)
	file, err := os.Create(filePath)
	if err != nil {
		return nil
	}
	defer func() {
		_ = file.Close()
	}()
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
