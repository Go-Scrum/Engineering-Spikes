build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/slack main.go

run:
	# build the package
	make build
	cp serverless.default.yaml serverless.yml
	# generate sam template for local development
	make run-local

deploy:
	make build

	cp serverless.default.yaml serverless.yml
	sls deploy --stage $(STAGE) -v --region $(REGION)

remove:
	sls remove --stage $(STAGE) -v --region $(REGION)

run-local:
	# run sam server locally
	#sam local start-api
	sls offline start --useDocker

