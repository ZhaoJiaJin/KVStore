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
