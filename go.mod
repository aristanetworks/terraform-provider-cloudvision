module github.com/aristanetworks/terraform-provider-cloudvision

go 1.16

replace github.com/aristanetworks/go-cvprac/v3 v3.0.3 => ./cloudvision/internal/go-cvprac

require (
	github.com/aristanetworks/go-cvprac/v3 v3.0.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.12.0
)

require (
	github.com/mattn/go-colorable v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
)
