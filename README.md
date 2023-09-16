# Tarea 1 - API RESTFul

## Dudas
- [ ] Parametro opcional query *

## Access Token - Amadeus API
- [x] Replicar peticion curl en GO
- [x] Parsear `access_token`
- [ ] Refresh Token *
## Busqueda
- [x] Crear tipo struct `AtributosBusqueda`
- [x] Rellenar peticion en el menu y enviar atributos a la API
- [x] Recibir busqueda desde menu en el Server y mandarlos a la API Amadeus.
- [ ] Leer respuesta de Amadeus y enviar respuesta al menu
- [ ] Mostrar tabla de resultados en el menu

### Errores y Bugs

**Amadeus responde con mal formato**

```
- go run server.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/search               --> main.busqueda (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on 0.0.0.0:5000
{"errors":[{"status":400,"code":477,"title":"INVALID FORMAT","detail":"This attribute must be a number","source":{"pointer":"adults","example":"1"}}]}
[GIN] 2023/09/16 - 14:23:50 | 200 | 11.658704375s |       127.0.0.1 | GET      "/api/search?originLocationCode=ARI&destinationLocationCode=SCL&departureDate=2023-12-02&adults=1&includedAirlineCodes=EK&nonStop=true&currencyCode=CLP&travelClass=ECONOMY"
```

## Referencias

### Go
- [JSON Example](https://gobyexample.com/json)
- [Clientes HTTP](https://apuntes.de/golang/clientes-http-uso-del-http-client/#gsc.tab=0)
- [String Interpolation](https://marketsplash.com/tutorials/go/golang-string-interpolation/)
- [How to parse JSON](https://blog.hackajob.com/how-to-parse-json-from-apis-in-golang/)
- [How to parse JSON (No usado)](https://dev.to/billylkc/parse-json-api-response-in-go-10ng)
- [Variables y constantes](https://www.digitalocean.com/community/tutorials/how-to-use-variables-and-constants-in-go-es)

### Gin Web Framework
- [Quickstart](https://gin-gonic.com/docs/quickstart/)
- [Bind query or post data](https://gin-gonic.com/docs/examples/bind-query-or-post/)
- [Tutorial: Restful API with Gin](https://go.dev/doc/tutorial/web-service-gin)
- [Serving data from reader](https://gin-gonic.com/docs/examples/serving-data-from-reader/)
- [IP:PORT Especifico](https://stackoverflow.com/questions/39448905/how-to-run-a-go-gin-server-on-a-specific-ip-address) 
