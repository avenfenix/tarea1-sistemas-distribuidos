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

**Amadeus responde error en formato**

```
{"errors":[{"status":400,"code":477,"title":"INVALID FORMAT","detail":"This attribute must be a number","source":{"pointer":"adults","example":"1"}}]}
```
**Fix:** En la direccion url tenia ```&adults=%s```debio ser ```&adults=%d``` para parsear correctamente numeros enteros.


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
