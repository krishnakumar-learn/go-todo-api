CREATE DATABASE go_todo_api;
CREATE USER 'todo'@'%' IDENTIFIED BY 'todopass';
GRANT ALL PRIVILEGES ON go_todo_api.* TO 'todo'@'%';
FLUSH PRIVILEGES;