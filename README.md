# Forum 
### A stock market simulator but for content-creators. Creators will be able to generate tokens (stocks) based on their performance on social media and fans can buy these tokens with an idea to sell it in future at higher price to make profit. 
<br>

## Tech Stack
- Client: Svelte w/Routify
- Server: Go, go-chi router, testify for unit-tests, sqlc for generating go-code from sql, viper for reading config file
- Database: PostgreSQL (ran inside a docker container

## Features
### For Creators:
- Generate tokens based on their social media presence and performace with one-click
- Select a preferred currency which will be used to display various prices accross the platform
- Since it will be later implemented on blockchain, no geographical restriction for any creator
- Creators are also able to buy tokens of different creators but only via a fan

### For Fans:
- Buy tokens of their favorite content-creators
- Select a preferred currency which will be used to display various prices accross the platform
- No geographical barrier. Unlike NYSE or BSE, fans will not be restricted to buy tokens of creators belonging to their country. They can buy tokens of any creator from anywhere on the globe

## Installation
### Requirements
- **Go (any version above 1.13)**
- **Postgres (via docker container)**
- **Latest Nodejs version for svelte**
- **Yarn or NPM to install svelte**
- **Docker**

### Procedure
- Clone this repo.
- Run ```go mod tidy``` to pull all the dependencies.
- Run the ```make postgres``` command. This will pull a docker image & create an instance of it.
- Run ```make createdb``` command. This will create a database in your postgres container.
- Then run ```make migrateup```. This will migrate the database schema in your db. **Note: You need to have golang-migrate install in your system for this to work**
- Start the golang server using the command ```go run cmd/web-services/main.go```.
**Now that you have your backend running, let's pull the frontend and set it up.**
- Get the frontend repo with ```git clone https://github.com/neel229/forum-frontend.git```
- Change your directory to the frontend directory and run ```yarn install```. This will fetch all the node_modules.
- After the step above, start the frontend server using ```yarn run dev```.

### Voila!, you have forum running in your localhost.

