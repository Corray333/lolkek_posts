# Posts handling microservise for Lolkek application
---
## Microservice structure:
├───cmd\  
&emsp;├───main.go\
└───internal\
&emsp;├───database\
&emsp;&emsp;├───database.go\
&emsp;├───models\
&emsp;&emsp;├───models.go\
&emsp;└───posts\
&emsp;&emsp;├───posts.go\
├───.env\
└───.girignore\
└───go.mod\
└───go.sum\
└───README.md\
___
## Packges:
1. */posts_ms/internal/database*: initializes connection to database "posts_db"
2. */posts_ms/internal/models*: includes models of database's documents
3. */posts_ms/internal/posts*: includes methods for post CRUD