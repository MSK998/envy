package envy

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

// Load the environment variables from a file
func Load(filename string) error {
	f, err := fileOrDefault(filename)
	if err != nil {
		return err
	}

	envVar, err := LoadFile(f)
	if err != nil {
		return err
	}

	for k, v := range envVar {
		err := os.Setenv(k, v)
		if err != nil {
			continue
		}
	}
	return nil
}

// Will check if file exists, if not it will create a default .env
func fileOrDefault(filename string) (string, error) {
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		file, err := os.Create(".env")
		if err != nil {
			return "", err
		}
		defer file.Close()
		return ".env", nil
	}
	return filename, nil
}

// Load file into a map returned to the user
func LoadFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	envMap, err := ParseFile(file)
	if err != nil {
		return nil, err
	}

	return envMap, nil
}

// Parse file, returning a map of key, value pairs
func ParseFile(r io.Reader) (map[string]string, error) {
	scanner := bufio.NewScanner(r)
	envVar := make(map[string]string)

	for scanner.Scan() {
		kv, err := ParseLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		if !kv.valid {
			continue
		}
		envVar[kv.key] = kv.value
	}

	return envVar, nil
}

// Parses a single line, returning a key value pair
func ParseLine(line string) (kvPair, error) {
	wsRgx := regexp.MustCompile(`\s`)
	clearLine := wsRgx.ReplaceAllString(line, "")

	verify, err := regexp.MatchString(`^[a-zA-z_\-\.]+(=+).+$`, clearLine)
	if err != nil || !verify {
		return kvPair{valid: false}, err
	}

	kvString := strings.Split(clearLine, "=")
	return kvPair{key: strings.ToUpper(kvString[0]), value: kvString[1], valid: true}, nil
}
