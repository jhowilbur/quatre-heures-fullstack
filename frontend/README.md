# Kumojin Frontend pour le projet Events

## Description:

Créez un exemple Typescript Vue.js 3 pour utiliser les API REST, afficher et modifier les données à l'aide d'Axios et de Vue Router.
- Chaque événement a un identifiant, un titre, une description, une date_initiale et une date_finale.
- Nous pouvons créer, récupérer des événements.
- Il y a une barre de recherche pour trouver des événements par identifiant.

#### Cela fait partie d'un défi backend et frontend pour le projet Events en 4 heures.

### Utilisation du menu fixe:
```bash
docker build -t frontend-events-kumojin .
docker run -p 8080:8080 frontend-events-kumojin
```

## Configuration du projet
```
npm install
```

### Compile et recharge à chaud pour le développement
```
npm run serve
```

### Compile et minimise pour la production
```
npm run build
```

### Exécutez vos tests
```
npm run test
```

### Fichiers peluches et correctifs
```
npm run lint
```

FAIRE:
- [ ] Ajouter des tests - Terminer les tests
- [ ] Ajouter plus de styles
- [ ] Ajouter plus de validations
- [ ] Ajouter plus de gestion des erreurs

---

# Kumojin Frontend for the Events project

## Description:

Build a Vue.js 3 Typescript example to consume REST APIs, display and modify data using Axios and Vue Router.
- Each Event has id, title, description, initial_date and final_date.
- We can create, retrieve Events.
- There is a Search bar for finding Events by id.


#### This is part of a backend and frontend challenge for the Events project in 4 hours.

### Using docker:
```bash
docker build -t frontend-events-kumojin .
docker run -p 8080:8080 frontend-events-kumojin
```

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

TODO:
- [ ] Add tests - Finish tests
- [ ] Add more styles
- [ ] Add more validations
- [ ] Add more error handling