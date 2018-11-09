ElecticCarBlockchain



### Building



```
Getting github token
https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/

Build container
docker build --build-arg github_auth_token=<token> --rm -t ech:latest .

Run container
docker run --rm -d --env-file=config/dev.env --name ech ech:latest

Read logs from container
docker logs ech
```
