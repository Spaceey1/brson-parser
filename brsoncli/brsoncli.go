package brsoncli

import (
	"brsonparser/brsonparser"
	"encoding/json/v2"
	"errors"
	"fmt"
	"os"
	"strings"
)

var outputPath = ""
var inputPath = ""
var encoding = false;
var decoding = false;

func parseFlag(flag string, index int) int{
	switch flag {
	case "-o":
		outputPath = os.Args[index + 1]
		return 2
	case "-d":
		inputPath = os.Args[index + 1]
		decoding = true
		return 2
	case "-e":
		inputPath = os.Args[index + 1]
		encoding = true;
		return 2
	default:
		return 1
	}
}

func Run() error {
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			i += parseFlag(arg, i) - 1
		}	
	}
	if encoding == decoding {
		fmt.Println()
		return errors.New("Specify -d to decode a brson into json or -e to encode")
	}
	if encoding {
		file, err := os.ReadFile(inputPath)
		if err != nil{
			return err
		}
		var data map[string]any
		err = json.Unmarshal(file, &data)
		if err != nil {
			return err
		}
		brsonparser.WriteBrsonToFile(data, outputPath)
	}
	if decoding {
		data, err := brsonparser.ReadBrsonFromFile(inputPath)
		if err != nil {
			return err
		}
		jsonbytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		os.WriteFile(outputPath, jsonbytes, 0644);
	}
	return nil
}
