[System Design Document](https://docs.google.com/document/d/1d1_5EIB0zfsGsthD9zHwn3oEDjJtaj_oIhpRw4auhVI/edit?usp=sharing)

# how to run the code

1. install golang
2. build the code using

```shell
go build -o kvstore cmd/main.go
```

3. install [goreman](https://github.com/mattn/goreman)
4. Run the code using
```shell
goreman start
```
This will start three processes using configurations in Procfile


5. Adding a new node to the cluster.
Call addnode API, pass the node ID and address, and run ./start3.sh to start another node

6. Removing a node
Stop the node, and call removenode API.

# [Postman API](https://documenter.getpostman.com/view/10246941/Szf26B31?version=latest)
