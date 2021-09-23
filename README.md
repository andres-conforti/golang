# golang

TRABAJO PRACTICO ENTREABLE PARA EL SEMINARIO DE GOLANG.

INTEGRANTES: CONFORTI ANDRES, JORGE EUGENIA.

El trabajo consistia en transformar un string en una estructura con sus correspondientes datos.
Para hacerlo dividimos la funcion en 2 partes.
Un main que llama a la funcion y dentro de la carpeta model, un entities donde fue modulizado el programa.

Luego de crear el modulo con el que trabajamos, e importar en todos los archivos las librerias necesarias, comenzamos por el main.
Primero, asignamos el valor de la cadena a utilizar, luego creamos una funcion que transforme 2 variables (una que controle el error y otra que guarde la cadena.) y se hace un llamado a la funcion en model encargada de revisar la cadena.
Luego segun el resultado obtendremos una cadena completa en forma de estructura como se pidio o una vacia sumada a su error el cual sera ejecutada por el Panic.

En segundo lugar, tenemos el archivo entities, en el model, donde trabajamos con las funciones necesarias.
Comenzamos creando la estructura de la cadena como se pidio, con un tipo, un lenght y un valor.
Luego creamos una funcion encargada de crear la cadena con 3 valores (2 string y 1 int) los cuales seran enviados al llamarla si todo esta bien.
Despues de esto, vamos a la funcion que llamamos desde main, la cual sera la encargada de llamar a las demas encargadas de revisar cualquier error que tenga la cadena, o simplemente retornarla completa.
Esta funcion del tipo Struct que creamos, y error, puede ser descripta de la siguiente manera:
Toma el valor del string que se le envia del main (por ejemplo "TX03ABC")
Luego revisa que la cadena cuente con minimo 4 valores.
Transforma el lenght de la cadena que es "03" a un entero
Este entero es comparado con la cantidad de caracteres que tiene el Value de nuestra cadena y si el tipo es valido (TX o NN)
Asigno a un par de variables con los valores VALUE y TIPO
Se checkea que estos sean validos, revisando que TIPO sea TX/NN y VALUE sea correspondiente a numeros o texto.
Cabe aclarar que tubimos en cuenta que cuando el LENGHT sea 00, y el VALUE nil, lo tomamos como correcto, ya que creemos que al ser de lenght 0, no deberia corresponderle un VALUE, lo cual no significa que deba ser una cadena invalida.
En el caso que no halla algun problema, se creara la estructura con un Return que llame a la funcion de creacion de la cadena.
En caso contrario, se devolvera una estructura vacia, acompañada de un error, tal y como se describio en el main.

En el caso del testing, para el main_test.go, solo efectuamos el llamado al main, ya que creimos que era mas que suficiente.
Y en el caso del entities_test.go, primero se importaron las librerias del assert con el comando "go get github.com/stretchr/testify/assert" y se tomo de referencia el snippet que hay en el github del tpe, donde modificamos los valores del mismo para que se adecue a nuestro trabajo, agregamos unos valores de prueba mas, y agregamos unos casos de prueba simple.

Los testeos ejecutados con el comando:
go test ./... -count=1 -cover
Generamos los archivos out y su html, con los comandos:
go test ./... -coverprofile=out.test
go tool cover -html out -o out.html