# go-pgdb
Go lang project using PostgresDB


Start the application using the below command:

go run main.go


Using postman make api calls to the application

Create Forum:

Method: Post

URL: http://localhost:8080/api/forum

Sample Body:

{
    "name": "AWS Forum",
    "threads": [
        {
            "title": "AWS Thread",
            "posts": [
                {
                    "title": "AWS Post",
                    "body": "AWS is Awesome"
                }
            ]
        }
    ]
}

Get All Forums:

Method: Get

URL: http://localhost:8080/api/forum/all


To verify the tables in postgres use the following commands:

docker-compose up -d

psql -h localhost -p 5432 -U postgres forum_threads



