package tipe

import (
	"context"
	"net/http"
)

type Documents interface {
	Create(context.Context, CreateDocumentOptions) error
	Get(context.Context, interface{}, GetDocumentOptions) error
	Update(context.Context, UpdateDocumentOptions) error
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
	payload := map[string]interface{}{
		"name":     opts.Name,
		"skuId":    opts.SkuID,
		"template": opts.Template,
	}

	if opts.Fields != nil {
		payload["fields"] = opts.Fields
	}

	if opts.Refs != nil {
		payload["refs"] = opts.Refs
	}

	// Create request
	req, err := ds.client.newRequest(
		http.MethodPost,
		ds.client.host,
		formatPath(ds.client.project, "createDocument"),
		payload,
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
	// SkuID of the Document
	SkuID string
}

// Get a Tipe document by id
func (ds *docService) Get(ctx context.Context, doc interface{}, opts GetDocumentOptions) error {
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

	resp := &Response{doc}

	// Get the document from Tipe
	if err := ds.client.do(req.WithContext(ctx), resp); err != nil {
		return err
	}

	return nil
}

type UpdateDocumentOptions struct {
	ID       string
	Fields   map[string]interface{}
	Name     string
	Refs     map[string]interface{}
	SkuID    string
	Status   string
	Template string
}

// Create a Tipe Document
func (ds *docService) Update(ctx context.Context, opts UpdateDocumentOptions) error {
	payload := map[string]interface{}{
		"fields":   opts.Fields,
		"id":       opts.ID,
		"name":     opts.Name,
		"refs":     opts.Refs,
		"skuId":    opts.SkuID,
		"template": opts.Template,
	}

	if opts.Status != "" {
		payload["status"] = opts.Status
	}

	// Create request
	req, err := ds.client.newRequest(
		http.MethodPut,
		ds.client.host,
		formatPath(ds.client.project, "updateDocument"),
		payload,
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
