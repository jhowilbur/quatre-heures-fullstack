# Application de gestion d'événements

## Description

L'application de gestion d'événements est une application web full-stack qui permet aux utilisateurs de gérer des événements. Elle est composée d'une application frontend Vue.js 3 et d'une API backend REST en Golang. Les utilisateurs peuvent créer, récupérer et rechercher des événements en utilisant l'interface utilisateur intuitive fournie par le frontend.

Chaque fichier a ses propres détails et des instructions spécifiques pour l'exécution, mais vous pouvez également utiliser docker-compose situé dans le répertoire principal.

#### Cela fait partie d'un défi backend et frontend pour le projet Events en 4 heures.

## Fonctionnalités

- Créer de nouveaux événements avec un identifiant unique, un titre, une description, une date de début et une date de fin.
- Récupérer les événements existants à partir de l'API backend.
- Interface utilisateur conviviale développée avec Vue.js 3 et TypeScript.
- Intégration transparente entre le frontend et le backend en utilisant Axios pour la communication avec l'API.
- Navigation entre différentes sections de l'application en utilisant Vue Router.
- Backend développé avec le langage de programmation Go.
- Utilisation de la bibliothèque Echo pour un framework web léger et rapide pour la création d'API RESTful.
- Utilisation du pilote PostgreSQL pur en Go pour le package database/sql.
- Utilisation de Cobra, une bibliothèque puissante pour les interfaces en ligne de commande (CLI) pour la création d'applications basées sur des commandes.

## Pour commencer

Pour commencer avec l'application de gestion d'événements, suivez les étapes ci-dessous :
```
docker-compose up
```

## API Backend

L'API backend fournit les endpoints nécessaires pour créer, récupérer et rechercher des événements. Elle est développée en utilisant Go et suit les principes RESTful et SOLID.

### Endpoints

- `POST /events` - Créer un nouvel événement. Accepte une charge utile JSON avec les détails de l'événement.
- `GET /events` - Récupérer tous les événements.
- `GET /events/{id}` - Récupérer un événement spécifique par son identifiant.


---

# Event Management Application

## Description

The Event Management Application is a full-stack web application that allows users to manage events. It consists of a Vue.js 3 frontend application and a backend Golang REST API. Users can create, retrieve, and search for events using the intuitive user interface provided by the frontend.

Each file has its own details and how to run each one specifically, but you can also use docker-compose located in the main path.

#### This is part of a backend and frontend challenge for the Events project in 4 hours.


## Features

- Create new events with a unique ID, title, description, initial date, and final date.
- Retrieve existing events from the backend API.
- User-friendly interface built with Vue.js 3 and TypeScript.
- Seamless integration between the frontend and backend using Axios for API communication.
- Navigation between different sections of the application using Vue Router.
- Backend built with Go programming language
- Echo lib for lightweight and fast web framework for building RESTful APIs.
- Pure Go PostgreSQL driver for database/sql package.
- Cobra for powerful CLI (Command-Line Interface) library for building command-based applications.

## Getting Started

To get started with the Event Management Application, follow the steps below:
```
docker-compose up
```

## Backend API

The backend API provides the necessary endpoints to create, retrieve, and search for events. It is built with Golang and follows RESTful and SOLID principles.

### Endpoints

- `POST /events` - Create a new event. Accepts a JSON payload with the event details.
- `GET /events` - Retrieve all events.
- `GET /events/{id}` - Retrieve a specific event by its ID.

