#

## Funcionamient (alto nivel)

customer -> submits job -> API
API -> creates jobOrder in DB or queue (status: nil)
Worker Manager -> Picks jobOrders
  -> Creates job for each jobOrder
JOB -> does the job
  -> Updates the jobOrder with result + updates status
CUSTOMER -> gets the completed job and status is changed to "retrieved"

## Componentes

1. API server: El API Server se encarga de gestionar las peticiones del usuario. A través del API Server, el usuario puede crear un nuevo *job*, consultar el estado de un *job* u obtener el resultado de un *job*.
1. Worker: El worker es un proceso sin fin (bucle infinito) que consulta la base de datos en busca de nuevos jobs y que lanza *go routines* que calculan el resultado de una *jobOrder* (o *job*). La idea es que se lanza la *go routine* sin esperar a que finalice o comprobar si falla. El primero caso, no esperamos que finalice porque el *worker manager* es un bucle sin fin. Si la *go routine* falla, el *job* no se actualiza y en el siguiente ciclo del *worker manager* se volverá a generar un *worker* para realizar el job. Usar un *channel* permitiría controlar el número de *go routines* en ejecución de manera simultánea (por si fuera necesario limitarlas, pero en nuestro caso, no será necesario).

### API Server

- `/api/v1/add/:num1/:num2` Genera un job para calcular la suma de `:num1` y `:num2`.
- `/api/v1/job/:id` Consulta el estado del job generado.

Para cada uno de los jobs, tenemos:

- `jobId` identificador único del job (usaremos el `id` de la base de datos o se puede genera una cadena *random* para un nombre de fichero.)
- `status` estado del job: `pending`, `wip`, `completed`, `failed`
- `LastUpdated` si el estado es `pending` o `wip` pero no se ha actualizado en `timeout` segundos, lanzar un nuevo *worker*.
- `CompletedOn` si el estado es `completed` o `failed`, indicar cuándo se ha dejado de trabajar en el job.

### API Server - implementación

Usaremos el *framework* Gin para el API Server.

Empezamos creando el *default router*:

```go
r:= gin.Default()
```

Importamos el paquete con `go get github.com/gin-gonic/gin` o usamos `go mod tidy`:

> En mi caso, ya dispongo del paquete descargado, por lo que sólo es necesario usarlo en `main()` y ejecutar `go mod tidy`
