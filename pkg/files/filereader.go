package files

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

func ReadFile(filePath string) []string {
	_, currentFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}

	file, err := os.Open(path.Join(path.Dir(currentFile), filePath))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			lines = append(lines, string(line))
		}
	}

	return lines
}
