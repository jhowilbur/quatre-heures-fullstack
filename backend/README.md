# Kumojin Backend pour le projet Events

## Description :
Il s'agit du backend pour le projet Events.
C'est une API REST simple qui permet à l'utilisateur de créer et de lire des événements.

- Chaque événement a un identifiant, un titre, une description, une date de début et une date de fin.
- Nous pouvons créer et récupérer des événements.

#### Cela fait partie d'un défi backend et frontend pour le projet Events en 4 heures.

------
## À FAIRE :
- [x] Ajouter un fichier README.md approprié
- [x] Ajouter un fichier .gitignore approprié
- [x] Ajouter un fichier .env approprié
- [x] Ajouter une migration de données
- [x] Préparer le package de configuration
- [x] Préparer les packages de base de données
- [x] Préparer l'entité Événement
- [x] Préparer le dépôt Événement (stockage)
- [x] Préparer le service Événement (logique métier)
- [ ] Ajouter des tests unitaires
- [ ] Ajouter des tests d'intégration - Terminer les tests
- [ ] Ajouter Swagger
- [ ] Ajouter une middleware de sécurité sur les points d'extrémité (JWT)
- [ ] Ajouter Jenkins CI/CD
- [ ] Ajouter Dockerfile
- [ ] Ajouter le déploiement Kubernetes

------

### Comment exécuter :

Tout d'abord, vous devez avoir une base de données PostgreSQL en cours d'exécution. Vous pouvez utiliser Docker pour l'exécuter :
```bash
docker run -e POSTGRES_USER=postgres -e POSTGRES_DB=test -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres
```

Ensuite, vous pouvez utiliser le fichier .env réel avec les variables actuelles et exécuter les migrations de la base de données de l'application :
```bash
go run *.go db migrate
```

Enfin, vous pouvez exécuter l'application :
```bash
go run *.go web
```

---

Utilisation de Docker :
```bash
docker build -t backend-events-kumojin .
docker run -p 8081:8081 backend-events-kumojin
```

Utilisation de Docker Compose : vous devez vous rendre dans le dossier où se trouve le fichier docker-compose.yml et exécuter :
```bash
docker-compose up
```

---

### Comment tester :

### Comment compiler :
Vous pouvez utiliser la commande de compilation :
```bash
go build -o backend-events-kumojin
```

### Comment déployer :
TODO

------
# Autres détails

### Structure des dossiers :
- En suivant les bonnes pratiques de Golang :
  https://github.com/golang-standards/project-layout/tree/master



---



# Kumojin Backend for the Events project

## Description:
This is the backend for the Events project. 
It is a simple REST API that allows the user to create, read events.

- Each Event has id, title, description, initial_date and final_date.
- We can create, retrieve Events.

#### This is part of a challenge backend and frontend Events project in 4 hours.

------
## TODO:
- [x] Add a proper README.md
- [x] Add a proper .gitignore
- [x] Add a proper .env file
- [x] Add data migration
- [x] Prepare the config package
- [x] Prepare the database packages
- [x] Prepare the Event Entity
- [x] Prepare the Event Repository (storage)
- [x] Prepare the Event Service (business logic)
- [ ] Add unit tests
- [ ] Add integration tests - Finish the tests
- [ ] Add Swagger
- [ ] Add security middleware on the endpoints (JWT)
- [ ] Add Jenkins CI/CD
- [ ] Add Dockerfile
- [ ] Add Kubernetes deployment

------

### How to run:

First, you need to have a postgres database running. You can use docker to run it:
```bash
docker run -e POSTGRES_USER=postgres -e POSTGRES_DB=test -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres
```

Then, you can use the actual .env file with the actual variables, and run the application db migrations:
```bash
go run *.go db migrate
```

Finally, you can run the application:
```bash
go run *.go web
```

---

Using docker:
```bash
docker build -t backend-events-kumojin 
docker run -p 8081:8081 backend-events-kumojin
```

Using docker-compose, you need to go to the folder where the docker-compose.yml file is located, and run:
```bash
docker-compose up
```

---

### How to test:

### How to build:
You can use build command:
```bash
go build -o backend-events-kumojin
```

### How to deploy:
TODO

------
# Other details

### Folder structure:
- Following the golang good practices:
https://github.com/golang-standards/project-layout/tree/master