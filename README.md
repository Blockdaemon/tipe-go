# tipe-go

Tipe Golang SDK

```bash
go get github.com/Blockdaemon/tipe-go
```

## Example

```go
import (
	"context"
	"fmt"

	"github.com/Blockdaemon/tipe-go"
)

type Document struct {
	CreatedBy tipe.CreatedBy `json:"createdBy"`
	Fields    struct {
		Description tipe.TextField `json:"description"`
	} `json:"fields"`
	ID       string        `json:"id"`
	Template tipe.Template `json:"template"`
}

func main() {
	client := tipe.New(
		tipe.Project("test"),
		tipe.Key(""),
		tipe.Offline(true),
		tipe.Port(8000),
	)

	doc := &Document{}

	if err := client.Documents.Get(
		context.Background(),
		doc,
		&tipe.GetDocumentOptions{
			SkuID: "MySkuID",
			Depth: 1,
		},
	); err != nil {
		panic(err)
	}

	fmt.Println(doc.Fields.Description.Value)
}
```
