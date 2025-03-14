### Clément DROUET, Gwendal KERBOUL et Luke RIOUX

# Ce projet constitué d'un Front, d'un Back et d'une Base de données est une To Do List

<img width="811" alt="image" src="https://github.com/user-attachments/assets/dd805385-5c4d-4f4e-bea7-17a98d8c14b9" />

# Technologies utilisées:
Front: HTML / CSS + bootstrap / Javascript

Back: Api en GoLang

BDD: Postgresql

### Fonctionnalités:

Avec cette To Do List on peux: **créer, afficher et supprimer des tâches**.

### Routes: 

__/getAllTaches__: permet d'afficher la liste des tâches stockées dans la base

__/addTache__: permet de créer une nouvelle tâche en spécifiant une date et une description

__/deleteTache__: permet de supprimer une tâche précise grâce à son id

# Gestion Docker:
Chaque composant à son **propre conteneur**. Nous en avons donc 3: le front, le back et la bdd.

Notre fichier **docker-compose** orchestre le tout pour démarrer ou stopper tous les conteneurs en même temps.

# Approche Devops / Bonnes pratiques

Pour ce projet, nous nous sommes répartis les tâches au maximum en ayant chacun une "spécialité":

Luke pour le front, Clément pour le back et Gwendal pour la base.

### Github

Le projet est stocké sur Github.

Se concentrer sur l'un des trois élément du projet ne nous a pas empêcher réfléchir et de travailler sur l'ensemble du projet pour garder le fonctionnement en tête et en avoir une compréhension complète.

Nous avons créer des branches pour chaque feature importante et particulièrement pour séparer clairement les devs du front, du back et de la base

Les noms de branches sont le plus **explicite** possible, voici un exemple:

<img width="778" alt="image" src="https://github.com/user-attachments/assets/aaf691a9-64bf-4cad-80df-249312c949d0" />

Après avoir terminé nos devs sur nos branches respectives, nous avions juste à créer une pull request pour fusionner notre travail sur main.

Nous avons instauré une règle simple: **La personne qui valide la pull request ne doit pas être celle qui l'a créée**. Ainsi, chaque personne contrôle et à connaissance du travail effectué par les autres.

Les secrets comme l'url de la base sont stockés dans un fichier **.env** du back 

__(Comme nous avont déplacé le projet de repo en cours, notre graph des dépots est séparé en deux)__
<img width="516" alt="image" src="https://github.com/user-attachments/assets/e4b099d2-9ed8-4687-a206-dfa49a604e56" />


# Comment démarrer le projet en local 

- Cloner le projet avec la commande git clone + l'url du projet
  
- Démarrer docker

- Se placer à la racine du projet dans un terminal

- utiliser la commande "__docker-compose up__" ou "__docker-compose up --build__"

- L'application est accessible à l'adresse suivante __http://localhost__ (pas besoin de spécifier un port) 

__Pour tester seulement le back, on peut faire des requêtes vers http://localhost:8080__





