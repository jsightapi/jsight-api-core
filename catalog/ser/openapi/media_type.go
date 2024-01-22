package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type mediaType string

// JSIght 0.3 supports: "json", "plainString", "binary"
const (
	MediaTypeJson        mediaType = "application/json"
	MediaTypeTextPlain   mediaType = "text/plain"
	MediaTypeOctetStream mediaType = "application/octet-stream" // TODO: discuss
)

func FormatToMediaType(f catalog.SerializeFormat) mediaType {
	switch f {
	case catalog.SerializeFormatJSON:
		return MediaTypeJson
	case catalog.SerializeFormatPlainString:
		return MediaTypeTextPlain
	case catalog.SerializeFormatBinary:
		return MediaTypeOctetStream
	default:
		panic("Unsupported srializeFormat")
	}
}
