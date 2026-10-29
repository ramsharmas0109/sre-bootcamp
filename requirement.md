requirement for task-1
1. make a feature branch from main branch.
2. we need image with base version.
3. we can build image.
4. we will prepare a CD pipeline
5. here we need one runner.
6. here we will use that built image to run the app.
7. we need one caddy server.









Main branch apna code hai, 

as a developer, I write a feature in a new branch, when I push that branch it should give me a docker iamge. 

as a devops, I should be able to deploy that new docker image on to my vultr instance. 

also the application should be exposed via caddy on IP address.



req:
1. Pipeline that will build docker image. 
    1. docker image should be pushed to ghcr or dockerhub (let's pick ghcr)
    2. in order to push image to ghcr (token for authentication), we need to create GH secret.
    3. the pipeline should trigger on push event.
2. Vultr instance
    1. Need connectivity from local via IP address and port using SSH. We need to store this information in the GH secrets as well.
    2. Vultr instance should have docker, docker-compose etc installed so that we can run the docker image.
    3. To ensure higher security we will run our self hosted runner.
        3.1 This self hosted runner will run on my local machine. 
    4. The pipleine should first do a docker pull, and run the server using docker-compose.
    5. The vultr instance should have the docker-compose file as well. Which means we need to clone this repo in that machine. Or we can just copy paste the file. (the later is not a good approach), we will go ahead with the first one. 
    6. In order to clone the repo, we need to pass tokens etc for authentication. 
    7. Once the docker container is up, we need some smoke test to ensure the deployment is successful.
    8. Once the deployment is done, we can and should inform the stake holders.
3. Once the deployment for application sever is done, my caddy server should reflect the changes. For that we need to keep caddy information in docker compose along with caddyfile in the repo.
4. Doing all the above will not gaurantee my application is up and running, why? 
    1. Becasue I have not taken care of my application dependencies. 
        Dependencices like DB, My app exects a DB to be running with certain tables & logical DB created. (DB migration should be done)
    2. In order to connect to DB, we can use docker compose setup, only catch is how are we going to pass the values to the container? 
        2.1 We can use github secrets, and create an env file that will be passed to the container at runtime. 
        2.2. We can use a proper secret management tool like vault. 
            2.2.1 But for vault to run we need to maintain it & manage it onto our vultr server. 
        2.3 For our use case let's use GH secrets. 


