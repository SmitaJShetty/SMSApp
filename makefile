
bld: 
		cd cmd/ && go build -o ../build/smssvc 

build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 go build -a -installsuffix cgo -o build/smssvc
	redis-server &
	./smssvc


clean: ## Remove all the temporary and build files
	go clean

run: bld
	build/smssvc

npm-run: 
	cd web/app/sms-app
	npm run start

npm-plint:
	cd web/app/sms-app
	npm run prettier && npm run lint

