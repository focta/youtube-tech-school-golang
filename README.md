# youtube-tech-school-golang


## 教材

https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=1

## ER図(diagram.io)

https://dbdiagram.io/d/648c2441722eb774941309f4


## マイグレーション

```
migrate -path db/migration -database "postgresql://yout:youtpass@localhost:15434/simple_bank?sslmode=disable" -verbose up
```

```
migrate -path db/migration -database "postgresql://yout:youtpass@localhost:15434/simple_bank?sslmode=disable" -verbose down
```