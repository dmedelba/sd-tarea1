---------Tarea 1 Sistema Distribuidos--------------------------------------------------------------
Integrantes : David Medel A. 201573548-4
	      Macarana Hidalgo A. 201473608-8
---------------------------------------------------------------------------------------------------
Se definen las siguientes maquinas:
1) dist69--> Clientes
2) dist70--> Logistica
3) dist71--> Camiones
4) dist72--> Finanzas
----------------------------------------------------------------------------------------------------

IMPORTANTE: Para correr todos los archivos tiene que estar en la carpeta sd-tarea1/PE1 (cd sd-tarea1/PE1): 

1) dist69--> make cliente
2) dist70--> make logistica
3) dist71--> make camiones
4) dist72--> 
----------------------------------------------------------------------------------------------------
Consideraciones Generales:
1) Antes de  correr el codigo asegurarse de que en la carpeta logistica_files el archivo pedidosGeneral.csv ESTÉ VACÍO
si no lo está porfavor eliminar. 
2) Antes de  correr el codigo asegurarse de que en la carpeta camion_files el archivo pedidos_camiones.csv ESTÉ VACÍO
si no lo está porfavor eliminar 
3) Ademas dentro de la carpeta retail y pyme eliminar todos los archivos csv, SOLO LOS EXTENCIONES CSV

** El orden de ejecución es Logistica-Cliente-Camiones **
* Ingresar ordenes en Clientes para poder ejecutar camiones

1)Camiones:
- la entrega de cada camión es secuencial (se llenan los 3 camiones y se entrega secuencialmente)
- Por tiempo se ha probado solo con 1 camión retail, pero funciona perfectamente, por lo cual según nuestra lógica debería
  funcionar con los demás.

