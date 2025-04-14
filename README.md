# Comments Service Test Task

**Описание проекта:**  
Микросервис на Go для хранения и обработки постов и комментариев с использованием GraphQL.  
Система поддерживает древовидную структуру комментариев (вложенность без ограничений) и пагинацию.  
аналогично таким платформам как Хабр и Reddit.

---

## Основной функционал

- Получение списка постов.
- Получение одного поста с комментариями.
- Возможность автором запретить комментарии к своему посту.
- Комментарии без ограничения по уровню вложенности.
- Пагинация для комментариев.
- Пагинация для статей.

---

## Конфигурация

Настройки проекта задаются через `config.yaml`
Пример конфигурации для запуска с PostgreSQL:

```yaml
server:
  host: localhost
  port: 8888
  env: local

repository-mode: postgres

postgres:
  user: postgres
  password: postgres
  name: mydb
  port: 5432
  host: postgres # имя контейнера PostgreSQL
  migrate: true
  sslmode: disable
```

для in-memory 
```yaml
server:
  host: localhost
  port: 8888
  env: local

repository-mode: in-memory
```

## Запуск

```bash
docker build -t comments-service-test-task .
```

для in-memory
```bash
docker run -p 8888:8888 comments-service-test-task:latest
```

для контейнера PostgreSQL
```bash
docker network create comments-network

docker run -d   --name postgres   --network comment-network   -p 5432:5432   -e POSTGRES_USER=postgres   -e POSTGRES_PASSWORD=postgres   -e POSTGRES_DB=mydb   postgres:latest

docker run -d   --name comments-service   --network comment-network   -p 8888:8888   comments-service-test-task:latest
```

## GraphQL API

### Добавить статью
```yaml
mutation {
  createArticle(input: {
    content: "тест1",
    commentPermission: true
  }) {
    ... on ArticleCreateOk {
      article {
        id
        content
        createdAt
        commentPermission
      }
    }
    ... on ArticleCreateBadRequest {
      message
    }
  }
}
```

### Добавить комментарий
```yaml
mutation {
  createComment(input: {
    content: "1 уровень",
    articleId:   "fcedfe1c-b1d5-4499-9218-cd32d7f35c43"
      parentId: null
  }) {
    ... on CommentCreateOk {
      comment {
        id
        content
        createdAt
        articleId
        parentId
        replies {
          id
          content
          createdAt
        }
      }
    }
    ... on CommentCreateBadRequest {
      message
    }
  }
}
```

### Список статей
```yaml
query {
  getList(pageNumber: 1, pageSize: 10) {
    ... on ListArticleGetOk {
      articles {
        id
        content
        createdAt
        commentPermission
      }
      total
    }
    ... on ListArticleGetBadRequest {
      message
    }
  }
}
```

### Статья с комментарием
```yaml
query {
  getArticle(
    articleId: "fcedfe1c-b1d5-4499-9218-cd32d7f35c43"
    commentPage: 1
    commentPageSize: 20
  ) {
    ... on ArticleGetOK {
      article {
        id
        content
        createdAt
        commentPermission
      }
      comments {
        id
        content
        createdAt
        replies {
          id
          content
          createdAt
          replies {
            id
            content
            createdAt
            replies {
              id
              content
              createdAt
              replies {
            id
            content
            createdAt
            replies {
              id
              content
              createdAt
            }
          }
            }
          }
        }
      }
    }
    ... on ArticleGetBadRequest {
      message
    }
  }
}
```

## Структура

```yaml
📄 README.md
💻 cmd
  └── main.go
⚙️ config.yaml
🐳 Dockerfile
📦 go.mod
🔐 go.sum
🔧 gqlgen.yml
🧩 graph
  ├── 🧠 model
  │   └── model.go
  ├── 🔄 resolver
  │   ├── articlemutation_create.resolvers.go
  │   ├── articlequery_get.resolvers.go
  │   ├── commentmutation_create.resolvers.go
  │   ├── listarticlequery_get.resolvers.go
  │   └── resolver.go
  ├── ⏳ runtime
  │   └── runtime.go
  └── 📜 schema
      ├── 📝 models
      │   ├── article.graphql
      │   └── comment.graphql
      ├── 🔨 mutation
      │   ├── article
      │   │   └── articlemutation_create.graphql
      │   └── comment
      │       └── commentmutation_create.graphql
      ├── 🔍 query
      │   └── article
      │       ├── articlequery_get.graphql
      │       └── listarticlequery_get.graphql
      └── scalar.graphql
🛠️ internal
  ├── ⚙️ config
  │   └── config.go
  ├── 🏷️ domain
  │   └── models.go
  ├── 📝 logger
  │   └── logger.go
  ├── 🔄 mappers
  │   └── resolver_mappers.go
  ├── 🗄️ repository
  │   ├── 🧠 inmemory
  │   │   └── inmemory_repository.go
  │   ├── 🎭 mocks
  │   │   └── repository_mock.go
  │   ├── 🐘 postgres
  │   │   ├── postgres_mappers.go
  │   │   ├── postgres_models.go
  │   │   └── postgres_repository.go
  │   ├── repository.go
  │   └── repository_test.go
  └── 🚀 service
      ├── article_service.go
      ├── 🎭 mocks
      │   └── service_mock.go
      ├── service.go
      └── service_test.go

```
