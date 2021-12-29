dev:
	go run .

provisionnig:
	docker-compose up -d

destroy:
	docker-compose down

test:
	go test

package:
	docker build -t ileossa/barista:latest .

publish-images:
	docker push ileossa/barista:latest

doc:
	swag i

publish-doc:
	echo "Not implemented"
