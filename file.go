package main

import (
	"fmt"
	"math/rand"
	"os"
)

// not power-loss atomic save data func:
func SaveData1(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer func() { // 4. discard the temp file if it still exists
		fp.Close() // not expected to fail
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = fp.Write(data); err != nil { // 1. save to temp file
		return err
	}
	if err = fp.Sync(); err != nil { // 2. fsync
		return err
	}
	err = os.Rename(tmp, path) // 3. replace target
	return err
}

func randomInt() int {
	return rand.Int()
}

// need to fix save data func so that it is powerloss atomic...

// it isnt powerloss atomic because it depends on directory updates, and directory
// updates are not power-loss atomic.

// need to make not directory dependant by calling fsync on the whole directory!
// to do this we need a handle (file descriptor) of the directory??

/* func SaveData2(path string, data []byte) error {
	dir, err := os.OpenFile(path, os.O_RDONLY|os.O_DIRECTORY, 0755)
	if err != nil {
		return err
	}
	defer dir.Close()

	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer func() { // 5. discard the temp file if it still exists
		fp.Close() // not expected to fail
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = fp.Write(data); err != nil { // 1. save to temp file
		return err
	}
	if err = fp.Sync(); err != nil { //2. fsync on file
	}
	if err = dir.Sync(); err != nil { // 3. fsync ON DIRECTORY
		return err
	}
	err = os.Rename(tmp, path) // 4. replace target
	return err
}
*/
