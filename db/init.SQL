CREATE TABLE IF NOT EXISTS tache (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO tache (date, description) VALUES
('2025-03-13', 'Tâche de test 1'),
('2025-03-14', 'Tâche de test 2');
--id s'auto incrémente, et on peut mettre la date ou le texte qu'on veut tant que c'est pas vide.
-- pour test des commandes SQL, il faut faire dans le terminal : docker exec -it postgres_db psql -U user -d mydb
-- ne fonctionne que si le volume est démarré avec : docker-compose up -d
