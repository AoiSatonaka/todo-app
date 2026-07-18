package constant

const (
	err_METADATA_PREFIX = "metadata error: "
	err_AUTH_PREFIX     = "authorize error: "
	err_CONVERT_PREFIX  = "convert error: "
)

const (
	// Metadata Error
	ERR_METADATA_NOT_SET = err_METADATA_PREFIX + "Metadata isn't set for Request"

	// Authorize Error
	ERR_AUTH_IDTOKEN_NOT_SET            = err_AUTH_PREFIX + "IDToken isn't set for User"
	ERR_AUTH_HEADER_NOT_SET             = err_AUTH_PREFIX + "Header isn't set for Metadata"
	ERR_AUTH_HEADER_NOT_EXPECTED_FORMAT = err_AUTH_PREFIX + "Hader is in an unexpected format"

	// Convert Error to Protobuf
	ERR_CONVERT_TO_PROTO_PRIORITY  = err_CONVERT_PREFIX + "Failed to convert Priority from domain to protobuf"
	ERR_CONVERT_TO_PROTO_STATUS    = err_CONVERT_PREFIX + "Failed to convert Status from domain to protobuf"
	ERR_CONVERT_TO_PROTO_TASK_INFO = err_CONVERT_PREFIX + "Failed to convert Task_Info from domain to protobuf"
	ERR_CONVERT_TO_PROTO_TASK      = err_CONVERT_PREFIX + "Failed to convert Task from domain to protobuf"

	// Convert Error to Domain
	ERR_CONVERT_TO_DOMAIN_PRIORITY = err_CONVERT_PREFIX + "Failed to convert Priority from protobuf to domain"
	ERR_CONVERT_TO_DOMAIN_STATUS   = err_CONVERT_PREFIX + "Failed to convert Status from protobuf to domain"
	ERR_CONVERT_TO_DOMAIN_TASK     = err_CONVERT_PREFIX + "Failed to convert Task from protobuf to domain"
)
