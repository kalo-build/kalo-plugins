module github.com/kalo-build/kalo-plugins/plugins/openapi-sdk

go 1.24.2

require (
	github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go v0.0.0-00010101000000-000000000000
	github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-ts v0.0.0-00010101000000-000000000000
)

require (
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/pb33f/jsonpath v0.8.1 // indirect
	github.com/pb33f/libopenapi v0.34.2 // indirect
	github.com/pb33f/ordered-map/v2 v2.3.0 // indirect
	go.yaml.in/yaml/v4 v4.0.0-rc.4 // indirect
	golang.org/x/sync v0.19.0 // indirect
)

replace (
	github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go => ../openapi-sdk-go
	github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-ts => ../openapi-sdk-ts
)
