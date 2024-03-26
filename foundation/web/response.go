package web

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-json-experiment/json"
)

func RenderHTML(ctx context.Context, w http.ResponseWriter, component templ.Component, statusCode int) error {
	setStatusCode(ctx, statusCode)
	w.WriteHeader(statusCode)

	return component.Render(ctx, w)
}

func RenderXML(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	setStatusCode(ctx, statusCode)

	w.Header().Set("Content-Type", HXMLMime)

	//bytes, err := xml.MarshalIndent(data, "", "  ")
	bytes, err := xml.Marshal(data)
	if err != nil {
		return err
	}

	//fmt.Println("******************************************************************************")
	//fmt.Println(string(bytes))

	_, err = w.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

// RespondJSON converts a Go value to JSON and sends it to the client.
func RespondJSON(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	setStatusCode(ctx, statusCode)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
