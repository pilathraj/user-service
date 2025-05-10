build:
	docker build -t user-service:latest .
run:
	docker-compose up -d
stop:
	docker-compose down
minikube:
	kubectl apply -f k8s/
	kubectl port-forward service/user-service 8084:8084
push:
	docker tag user-service:latest pilathraj/user-service:latest
	docker push pilathraj/user-service:latest
minikube-stop:
	kubectl delete -f k8s/
test:
	go test -v ./... -coverprofile=coverage.out