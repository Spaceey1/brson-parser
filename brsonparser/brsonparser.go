package brsonparser

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/andybalholm/brotli"
	"go.mongodb.org/mongo-driver/bson"
)

var brsonHeader = []byte{70, 114, 68, 84, 0, 0, 0, 0, 3}

func ReadBrson(data []byte) (map[string]any, error) {
	if len(data) < 9 || !bytes.Equal(data[:9], brsonHeader) {
		return nil, fmt.Errorf("invalid BRSON header")
	}
	// BRSON header is skipped
	compressed := data[9:]

	br := brotli.NewReader(bytes.NewReader(compressed))
	decompressed, err := io.ReadAll(br)
	if err != nil {
		return nil, fmt.Errorf("brotli decompression failed: %w", err)
	}

	var doc map[string]any
	if err := bson.Unmarshal(decompressed, &doc); err != nil {
		return nil, fmt.Errorf("bson unmarshal failed: %w", err)
	}

	return doc, nil
}

func WriteBrson(doc map[string]interface{}) ([]byte, error) {
	bsonData, err := bson.Marshal(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal BSON: %w", err)
	}

	var compressedBuf bytes.Buffer
	writer := brotli.NewWriter(&compressedBuf)
	if _, err := writer.Write(bsonData); err != nil {
		return nil, fmt.Errorf("brotli compression failed: %w", err)
	}
	writer.Close()

	final := append(brsonHeader, compressedBuf.Bytes()...)
	return final, nil
}

func WriteBrsonToFile(doc map[string]interface{}, name string) error {
	bytes, err := WriteBrson(doc);
	if err != nil {
		return err
	}
	err = os.WriteFile(name, bytes, 0644);
	return err
}

func ReadBrsonFromFile(name string) (map[string]any, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	data, err := ReadBrson(file)
	return data, err;
}
