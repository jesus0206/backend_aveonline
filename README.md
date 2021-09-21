# Servicio para Fichaje electonico

Aveonline desea inicia una unidad de negocio en el segmento farmaceutico, la compañia desea ingresar al mercado con la capacidad de entregar una alta variedad de promociones con duración diaria, semanal o mensual.
El Servicio Rest fue esto en **Golang 1.14**

## Instalando dependencias

```
go mod download
```

## Correr aplicacion

- Copiar y Pegar **example.env** y renombrar a **.env**
- luego configurar las variables de entorno en el archivo **.env** para el correcto funcionamiento.

```
go run main.go
```

## Compilar aplicación

```
go build -o main .
```

## Correr binario

```
./main
```

## generar la imagen docker

```
chmod +x build.sh &&
./build.sh
```
