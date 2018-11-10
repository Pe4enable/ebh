ElecticCarBlockchain



### Building



```
Getting github token
https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/

Build container
docker build --build-arg github_auth_token=<token> --rm -t ech:latest .

Run container
docker run --rm -d -p 8001:8001 --network="host" --name ech ech:latest
 --env-file=config/dev.env

Read logs from container
docker logs ech
```
