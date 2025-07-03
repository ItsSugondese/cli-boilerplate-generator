package file_writer

import (
	"bufio"
	"cli-boilerplate-generator/pkg/utils/string_utils"
	"fmt"
	"os"
	"strings"
)

func ReadAndWriteFromStorageFileToEnumFile(mainFilePath, fileToCreate, module string) {
	input, err := os.Open(mainFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer input.Close()

	output, err := os.Create(fileToCreate)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	snake := strings.ReplaceAll(module, "-", "_") + "_enums"
	upperCamel := string_utils.ToCamelCase(module, '-', true)
	lowerCamel := string_utils.ToCamelCase(module, '-', false)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "?snake", snake)
		line = strings.ReplaceAll(line, "?u", upperCamel)
		line = strings.ReplaceAll(line, "?l", lowerCamel)

		writer.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}

func WriteStringInFile(data, filePath string) error {
	//data := []byte("Hello, this is the content to write to the file.\n")
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
