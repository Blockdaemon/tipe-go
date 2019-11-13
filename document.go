package tipe

import (
	"context"
	"net/http"
)

type Documents interface {
	Create(ctx context.Context, opts CreateDocumentOptions) error
	Get(ctx context.Context, d interface{}, opts GetDocumentOptions) error
}

type docService service

type CreateDocumentOptions struct {
	Fields   map[string]interface{}
	Name     string
	Refs     map[string]interface{}
	SkuID    string
	Template string
}

// Create a Tipe Document
func (ds *docService) Create(ctx context.Context, opts CreateDocumentOptions) error {
	// Create request
	req, err := ds.client.newRequest(
		http.MethodPost,
		ds.client.host,
		formatPath(ds.client.project, "createDocument"),
		map[string]interface{}{
			"fields":   opts.Fields,
			"name":     opts.Name,
			"refs":     opts.Refs,
			"skuId":    opts.SkuID,
			"template": opts.Template,
		},
	)
	if err != nil {
		return err
	}

	resp := &Response{}

	// Get the document from Tipe
	if err := ds.client.do(req.WithContext(ctx), resp); err != nil {
		return err
	}

	return nil
}

type GetDocumentOptions struct {
	// ID of the Document
	ID string
	// Depth is the number of levels to recursively fetch document refs
	Depth int
	// SkuID of the Document
	SkuID string
}

// Get a Tipe document by id
func (ds *docService) Get(ctx context.Context, d interface{}, opts GetDocumentOptions) error {
	payload := map[string]interface{}{}

	var command string

	switch {
	case opts.ID != "":
		command = "documentById"
		payload["id"] = opts.ID
	case opts.SkuID != "":
		command = "documentBySkuId"
		payload["skuId"] = opts.SkuID
	}

	if opts.Depth > 0 {
		payload["depth"] = opts.Depth
	}

	// Create request
	req, err := ds.client.newRequest(
		http.MethodPost,
		ds.client.host,
		formatPath(ds.client.project, command),
		payload,
	)
	if err != nil {
		return err
	}

	resp := &Response{d}

	// Get the document from Tipe
	if err := ds.client.do(req.WithContext(ctx), resp); err != nil {
		return err
	}

	return nil
}
