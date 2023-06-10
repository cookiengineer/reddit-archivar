package structs

import "os"
import "path"
import "strings"

type Cache struct {
	Folder string `json:"folder"`
}

func NewCache(folder string) Cache {

	var cache Cache

	if strings.HasPrefix(folder, "/") {

		if strings.HasSuffix(folder, "/") {
			folder = folder[0:len(folder)-1]
		}

		stat, err := os.Stat(folder)

		if err == nil && stat.IsDir() {

			cache.Folder = folder

		} else {

			err2 := os.MkdirAll(folder, 0750)

			if err2 == nil {
				cache.Folder = folder
			}

		}

	} else {

		folder = "/tmp/reddit-archivar"

		if strings.HasSuffix(folder, "/") {
			folder = folder[0:len(folder)-1]
		}

		stat, err := os.Stat(folder)

		if err == nil && stat.IsDir() {

			cache.Folder = folder

		} else {

			err2 := os.MkdirAll(folder, 0750)

			if err2 == nil {
				cache.Folder = folder
			}

		}

	}

	return cache

}

func (cache *Cache) Exists(file string) bool {

	var result bool = false

	if !strings.HasPrefix(file, "/") {
		file = "/" + file
	}

	_, err := os.Stat(cache.Folder+file)

	if err == nil {
		result = true
	}

	return result

}

func (cache *Cache) Write(file string, buffer []byte) bool {

	var result bool = false

	if !strings.HasPrefix(file, "/") {
		file = "/" + file
	}

	folder := path.Dir(file)
	stat1, err1 := os.Stat(cache.Folder+folder)

	if err1 == nil && stat1.IsDir() {

		err2 := os.WriteFile(cache.Folder+file, buffer, 0666)

		if err2 == nil {
			result = true
		}

	} else {

		err2 := os.MkdirAll(cache.Folder+folder, 0755)

		if err2 == nil {

			err3 := os.WriteFile(cache.Folder+file, buffer, 0666)

			if err3 == nil {
				result = true
			}

		}

	}

	return result

}
