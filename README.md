# go-pgdb
Go lang project using PostgresDB


Start the application using the below command:

docker-compose up -d

go run main.go


Using postman make api calls to the application

Create Forum:

Method: Post

URL: http://localhost:8080/api/forum

Sample Body:

{
    "name": "Golang Forum",
    "threads": [
        {
            "title": "Golang Thread",
            "posts": [
                {
                    "title": "Golang Post",
                    "body": "Golang is Simple and Easy"
                }
            ]
        }
    ]
}

Get All Forums:

Method: Get

URL: http://localhost:8080/api/forum/all


To verify the tables in postgres use the following commands:

psql -h localhost -p 5432 -U postgres forum_threads

To run tests in db folder:

go test go-pgdb/db



