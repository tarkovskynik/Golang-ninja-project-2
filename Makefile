
mongo-up:
	@docker run --name mongodb -d -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=qwerty123 -p 27017:27017 mongo

.PHONY: mongo-up