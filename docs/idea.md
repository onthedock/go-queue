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

---

Más ideas...

En la versión del código del *commit* `fffdaee` tras llamar a la apliación con:

```shell
$ curl -X POST http://localhost:8080/api/v1/add/12/23
{"num1":12,"num2":23,"operation":"add","result":"35"}
```

... se devuelve directamente el resultado al solicitante.

Como la idea es simular algún tipo de cola, podemos, por ejemplo generar un número de *job*, que el usuario puede usar más tarde para consultar el estado del job.

Podemos almacenar el resultado en un documento (un fichero) con formato JSON en el disco o guardarlo en una base de datos. La forma más sencilla es mediante un fichero.

Para que el Id del *job* sea único, usamos el paquete [`uuid`](https://pkg.go.dev/github.com/google/uuid) de Go.

> En [Generate a UUID/GUID in Go (Golang)](https://golangbyexample.com/generate-uuid-guid-golang/) de GoLangByExample tenemos un ejemplo del uso del paquete.

---

Más ideas....

El proceso completo sería:

usuario --> envía petición al servidor para sumar dos números
api server --> comprueba que los datos son correctos
api server --> genera un documento de job (con un jobID)
api server --> devuelve el Id del job al usuario

Además del proceso del api server, que se encarga de interaccionar con el usuario, debería haber otro proceso que revisa los jobs que están pendientes para procesarlos.

La "cola" se puede implementar de varias maneras. Como quiero trabajar con ficheros, el documento creado por el api server cuando recibe la petición del usuario puede tener la extensión *pending*, por ejemplo. Así, el proceso de ejecución de los jobs debe obtener los ficheros con extensión *pending*. Otra opción sería una base de datos, pero excepto si se usa SQLite, todas las bases de datos requerirán un componente adicional que ejecutar, por lo que vamos a hacerlo lo más simple posible.

apiserver --> recibe los datos del usuario, genera un id y lo devuelve al usuario tras generar el fichero.
processor --> procesa la operación, guarda el resultado en el fichero, cambia la extensión. 
cleaner --> elimina los ficheros que ya han sido descargados por el usuario