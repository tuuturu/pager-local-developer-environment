# pager-local-developer-environment
Local developer environment setup for the [Pager
project](https://github.com/tuuturu?q=pager&type=&language=)

## Prerequisites

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Setup (Linux)

1. `git clone git@github.com:oslokommune/hubdevenv.git && cd hubdevenv`
1. `sudo echo "127.0.0.1       auth" >> /etc/hosts`
	Required due to Keycloak using dynamic base URL in .well-known data

## Usage

1. Add the following line to your hosts file
1. Start developer environment
	```sh
	docker-compose up -d
	```

## Result

* Database (Postgres) available at localhost:5432
	- Stores data for event-service
	- Script for adding provisioned data/setup
		[here](https://github.com/oslokommune/hubdevenv/blob/master/data/postgres/docker-entrypoint-initdb.d/init-databases.sh)

* Authentication provider (Keycloak) available at http://localhost:8080
	- Handles users and authentication
	- Provisioned test users [here](https://github.com/tuuturu/pager-local-developer-environment/blob/main/data/keycloak/README.md)

* Authentication backend (Gatekeeper) available at http://localhost:4554
	- Handles frontend authentication with authentication provider and converting
		token to Authorization header

* Event backend (pager-event-service) available at http://localhost:3000
	- Handles CRUD for events
