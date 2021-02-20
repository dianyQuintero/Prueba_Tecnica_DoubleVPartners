# Prueba_Tecnica_DoubleVPartners
En este repositorio se encuentra la prueba técnica para el cargo de Software Developer en Double V Partners

## Requerimientos
- Docker
- Go
- PostgreSQL
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
  
  Para la instalación de la librería golang-migrate seguir las siguientes instrucciones:
  ### Para MacOS:
  Ejecutar el siguiente comando en la consola: brew install golang-migrate

  ### Para Windows
  1. Abrir una consola de PowerShell
  2. Instalar scoop con el siguiente comando: iwr -useb get.scoop.sh | iex
  3. Instalar la librería golang-migrate con el siguiente comando: scoop install migrate

## Uso
1. Clonar el repositorio: git clone https://github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners.git
2. Ingresar a la ubicación en donde se clonó el repositorio
3. Crear e iniciar los servicios ejecutando el comando: docker-compose up --build
4. Los archivos de migración de la base de datos están en db/migrations. Ejecutarlos usando migrate de la siguiente manera:
   migrate -path db/migrations -database "postgres://postgres:diany990925@localhost:5432/ticketsDatabase?sslmode=disable" -verbose up
   
   (este comando se ejecuta en una consola de power shell después de haber ingresado a la ubicación en donde se clonó el repositorio)
5. Probar el API desde postman o con comandos curl
    - Crear un ticket:
    ![image](https://user-images.githubusercontent.com/42391925/108585185-075d6200-7315-11eb-8be6-41d2312a9775.png)
    (Solo se requiere enviar por parámetro el nombre del usuario ya que el id se asigna automaticamente, la fecha de creación se asigna según el día en el que se haga la request, la fecha de actualización se asigna automaticamente igual que la de creación, el atributo "estado_abierto" es un booleano que se asigna automaticamente en true al crear el ticket, para esto asumí que el estado de un ticket recien creado siempre va a ser abierto es decir true)
    
    - Obtener todos los tickets:
    ![image](https://user-images.githubusercontent.com/42391925/108585301-c154ce00-7315-11eb-8b4d-36ab3a803134.png)
    - Obtener un ticket en específico (localhost:8080/items/{idTicket}):
    ![image](https://user-images.githubusercontent.com/42391925/108585342-0a0c8700-7316-11eb-858c-3a6d09f30c83.png)
    - Actualizar la información de un ticket (localhost:8080/items/{idTicket}):
    ![image](https://user-images.githubusercontent.com/42391925/108585391-55bf3080-7316-11eb-901a-77d451ac8cd8.png)
    (Se actualizan los atributos de usuario, fecha_actualización y estado_abierto)
    - Eliminar un ticket (localhost:8080/items/{idTicket}):
    ![image](https://user-images.githubusercontent.com/42391925/108585439-9454eb00-7316-11eb-9e0d-1d1b1884fa4d.png)
    
    ![image](https://user-images.githubusercontent.com/42391925/108585447-a8005180-7316-11eb-8059-c5160bdebde4.png)
    
    
  6. Ejecutar:
     migrate -path db/migrations -database "postgres://postgres:diany990925@localhost:5432/ticketsDatabase?sslmode=disable" -verbose down
     
     (Solo si se necesita eiliminar la tabla que contiene los tickets)




    
    
