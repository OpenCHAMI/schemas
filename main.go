package main

import (
	"encoding/json"

	"os"
	"path/filepath"

	"log"

	"github.com/invopop/jsonschema"
	"github.com/openchami/schemas/schemas"
	"github.com/openchami/schemas/schemas/csm"
)

type InventoryRequest struct {
	Header               schemas.Envelope          `json:"header"`
	InventoryDetailArray []schemas.InventoryDetail `json:"inventory_detail_array"`
}

func generateAndWriteSchemas(path string) {
	schemas := map[string]interface{}{

		"Component.json":              &csm.Component{},
		"RedfishEndpoint.json":        &csm.RedfishEndpoint{},
		"InventoryDetailRequest.json": &InventoryRequest{},
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal("Failed to create schema directory")
	}

	for filename, model := range schemas {
		schema := jsonschema.Reflect(model)
		data, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			log.Fatal("Failed to generate JSON schema")
		}
		fullpath := filepath.Join(path, filename)
		if err := os.WriteFile(fullpath, data, 0644); err != nil {
			log.Fatal("Failed to write JSON schema to file")
		}
		log.Println("Schema written")
	}
}

func main() {
	generateAndWriteSchemas("jsonschemas")
}
