> docker run --rm -v $(pwd):$(pwd) -w $(pwd) redocly/openapi-cli lint openapi.yml --skip-rule no-invalid-media-type-examples --skip-rule info-license
> docker run --rm -v $(pwd):$(pwd) -w $(pwd) openapitools/openapi-generator-cli:v5.2.0 generate -i openapi.yml -g typescript-axios -o output/ts-axios
> docker run --rm -v $(pwd):$(pwd) -w $(pwd) openapitools/openapi-generator-cli:v5.2.0 generate -i openapi.yml -g go-echo-server -p enumClassPrefix=true -o output/go-echo-server
