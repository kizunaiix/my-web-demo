gorun:
	cd cmd && swag init -d ./,../internal/handler
	cd cmd && go run .

gobuild:
	cd cmd && swag init -d ./,../internal/handler
	cd cmd && go build -o ../bin/gin-app-be .

dockerbuild:
	docker build -t kizunaiix/my-web-demo-backend:latest .

dockerrun-dev:
	docker run -itd \
	--name my-web-demo-backend \
	-e ENV=dev \
	-p 9000:9000 \
	kizunaiix/my-web-demo-backend:latest
