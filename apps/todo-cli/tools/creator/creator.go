package creator

import "os"

func CreateDir(dirName string) error {
	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dirName, 0755); err != nil {
				return err
			}
		}
	}

	return nil
}

func CreateFile(fileName string) error {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(fileName); err != nil {
				return err
			}
		}
	}

	return nil
}
