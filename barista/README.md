# Nom : barista world ou bien le monde du barista

Le but de ce repo est de rassembler au même endroit les différents tests, poc et snippet que j'ai réalisé.
Généralement lorsque j'ai une idée, il me manque toujours un environnement, un projet, quelque chose qui existe déjà sur lequel je peux m'appuyer pour mettre en pratique directement ma pensée. 

J'ai souhaiter mettre en place ce repository à la suite d'un test technique. J'ai voulus apprendre d'avantage la programation en GO, c'est pour cela que le projet `barist world` est né.

Le projet est loin d'être "joli", bien écris, architecturé. Mais il continura d'évolué au fil du temps, tout comme n'importe quel projet à son commencement.

Pour apprendre je dois tester, et pour tester il faut faire, donc je fais 😇.

J'ai pleines de chose à tester, merci de consulter la section Roadmap.


# Liste badges

gitlab, ci, code coverage, guard rails, [more](https://shields.io/).


# Requirements:

- docker >= 17.12.0+
- docker-compose
- golang >= 1.17
- make

# Installation

```bash
git clone https://github.com/ileossa/barista.git
```
Pour démarrer le service en local avec la configuration local, voir la section make->environnement pour plus d'options.

```bash
make run dev
```
Le service va s’exécuter et lancer les différentes images docker nécessaire.

Ouvrir le navigateur à l’adresse [http://localhost:8080/api/v1/ping](http://localhost:8080/api/v1/ping)

# Architecture

mettre une image de l'architecture du project
```
                                 ┌────────┐
                                 │        │
                                 │  HTML  │
                                 │        │
                           ┌────►│  FILES │
                           │     │        │
┌───────────────┐          │     │        │
│               │          │     └────────┘
│               │          │
│    Go         │          │
│               │          │      ┌────────┐             ┌─────────┐
│    API        │          │      │        │             │         │
│               ├──────────┤      │        │             │         │
│               │          │      │  API   │             │  BDD    │
│               │          │      │        │             │         │
│               │          └──────►  V1    │             │         │
│               │                 │        │             │         │
└───────────────┘                 │        │             │         │
                                  │        │             │         │
                                  └────────┘             └─────────┘
```

# Documentation

```bash
- make run [env]  #env: run application with a configuration environement name (like dev, pre-prod,prod)

- make package : build docker image with last version of project.
make publish-images : push latest\'s tag docker image on dockerhub

- make doc : push into publish/documentation branch documentation generate with swag tool

```

# Support
Vous pouvez ouvrir des PR, si vous le souhaitez. 
- Pour me faire des remarques sur mon code
- Proposer des améliorations, guideline, ...
- Juste pour discuter
- Pour pimper à votre sauce 😎

# Roadmap
Je tiens la roadmap directement sur trello à cette adresse : [https://trello.com/b/txyFCUOq](https://trello.com/b/txyFCUOq). 
Je vais tester aussi les solutions de Github et Gitlab. Pour comparer. 

# Erreurs rencontrées
if port 8080 already used

lsof -i :8080 | grep LISTEN kill -9 <pid>

# License
[Mozilla Public License 2.0](https://choosealicense.com/licenses/mpl-2.0/)