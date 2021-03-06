package config
// Read all JSON-files from a directory
import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/itshosted/mcore/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	Extension    = ".json"
	ExtensionDir = ".d"
)

// Return all files available in basedir.
// WARN: Function forces dir to only contain files (dir means error)
func files(basedir string) (map[string]string, error) {
	out := make(map[string]string)

	// Check if directory is named directory.d
	s := strings.Split(basedir, ExtensionDir)
	if len(s) != 2 {
		return out, errors.New(fmt.Sprintf("Config directory must be named 'directory%s'", ExtensionDir))
	}

	// Parse each file in basedir
	err := filepath.Walk(basedir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if basedir == path {
			// Ignore basedir
			return nil
		}

		if f.IsDir() {
			// Only read files
			return errors.New("LoadJsonD does not support directories within directories")
		}

		log.Debug("Found config=%s", path)
		out[f.Name()] = path
		return nil
	})

	return out, err
}

// Read files from basedir and convert JSON to map[filename]type
// Note: Dir should end with .d suffix (Debian/XSNews convention)
// Note: Only .json-files should exist in this dir
// Note: Subdirs not allowed in basedir
func LoadJsonD(basedir string, x interface{}) error {
	// Get list of files
	filelist, err := files(basedir)
	if err != nil {
		return err
	}

	// No files found
	if len(filelist) == 0 {
		return nil
	}

	// Create one big json string where every file is a key
	jsonCollection := []string{}
	for filename, fullpath := range filelist {
		// Only load directory.d/file.Extension
		s := strings.Split(filename, Extension)
		if len(s) != 2 {
			return errors.New(fmt.Sprintf("Invalid file '%s' present in config dir.", filename))
		}

		// Load content from file
		if data, err := ioutil.ReadFile(fullpath); err == nil {
			// Add to our json structure with key "filename"
			jsonCollection = append(jsonCollection, fmt.Sprintf(`"%s": %s`, s[0], data))
		}
	}

	// Unmarshal json
	return json.Unmarshal([]byte(fmt.Sprintf("{%s}", strings.Join(jsonCollection, ","))), &x)
}
