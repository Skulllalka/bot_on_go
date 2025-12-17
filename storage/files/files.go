package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	e "github.com/Skulllalka/bot_on_go/lib"
	"github.com/Skulllalka/bot_on_go/storage"
)

const (
	defaultPerm = 0774
)

var errNoSavedPages = errors.New("page not saved")

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

func (s Storage) Remove(p *storage.Page) error {

	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("can't remove file", err)
	}

	filePath := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(filePath); err != nil {
		return e.Wrap(fmt.Sprintf("can't remove file with name: %s", filePath), err)
	}

	return nil
}

func (s *Storage) IsExists(page *storage.Page) (bool, error) {
	fileName, err := fileName(page)
	if err != nil {
		return false, e.Wrap("can't reach file", err)
	}
	filePath := filepath.Join(s.basePath, page.UserName, fileName)
	switch _, err := os.Stat(filePath); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, e.Wrap(fmt.Sprintf("can't reach the file %s", filePath), err)
	}
	return true, nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	filePath := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, e.Wrap("can't open dir", err)
	}
	if len(files) == 0 {
		return nil, errNoSavedPages
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))
	file := files[n]
	return s.decodePage(filepath.Join(filePath, file.Name()))
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("cannot open file", err)
	}
	defer func() {
		_ = f.Close()
	}()

	var p storage.Page
	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("cant decode file", err)
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
