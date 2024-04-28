.PHONY: build-lambdas

build-lambda:
	@echo "Building Go Lambda functions..."
	@for dir in $$(find lambdas -type d); do \
		if [ -f $$dir/main.go ]; then \
			echo "Building in $$dir..."; \
			GOOS=linux GOARCH=amd64 go build -o $$dir/bootstrap $$dir/main.go; \
			cd $$dir && zip function.zip bootstrap && cd -; \
			echo "Lambda function in $$dir built and zipped successfully."; \
		fi \
	done


.PHONY: cdk-deploy

deploy: build-lambda
	@echo "Deploying CDK stack..."
	@cdk deploy
