# Generate swagger documentation

You must install : [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)

[doc installation](https://goswagger.io/install.html)

execute : 
```
go get -u github.com/swaggo/swag/cmd/swag
swag init && go build main.go
```

Il est possible de mettre un api/route dans l'application et d'ajouter la documentation dans le binaire. Mais cela alourdi l'executable. 
Le but ici, et d'avoir un executable et la documentation séparé. 
Je vais mettre la doc sur une branche spécifique, par exemple _doc/v1_
Comme ça plusa facile de la publier, attentions elle doit être mise à jour à chaque pipeline/build.

Offical base documentation for
swagger+Go: [https://github.com/swaggo/swag/blob/master/README.md](https://github.com/swaggo/gin-swagger)



```
./swag init
```