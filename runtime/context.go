package runtime

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const metadataHeaderPrefix = "Grpc-Metadata-"

/*
AnnotateContext adds context information such as metadata from the request.

If there are no metadata headers in the request, then the context returned
will be the same context.
*/
func AnnotateContext(ctx context.Context, req *http.Request) context.Context {
	var pairs []string
	for key, vals := range req.Header {
		for _, val := range vals {
			if key == "Authorization" {
				pairs = append(pairs, key, val)
				continue
			}
			if strings.HasPrefix(key, metadataHeaderPrefix) {
				pairs = append(pairs, key[len(metadataHeaderPrefix):], val)
			}
		}
	}

	if len(pairs) == 0 {
		return ctx
	}
	return metadata.NewContext(ctx, metadata.Pairs(pairs...))
}
