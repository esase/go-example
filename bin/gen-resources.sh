# Generate types
oapi-codegen \
	-generate "types" \
	-o "internal/suppliers/base/models/models.go" \
	-package "models" \
	"api/openapi.json"

# Generate spec
oapi-codegen \
	-generate "spec" \
	-o "internal/middleware/oa_spec_gen.go" \
	-package "middleware" \
	"api/openapi.json"