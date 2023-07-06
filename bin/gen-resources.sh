# Generate types
oapi-codegen \
	-generate "types" \
	-o "internal/common/schema/schema.go" \
	-package "schema" \
	"api/openapi.json"

# Generate spec
oapi-codegen \
	-generate "spec" \
	-o "internal/middleware/openapi/oa_spec_gen.go" \
	-package "openapi" \
	"api/openapi.json"