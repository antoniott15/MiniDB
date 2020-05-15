# MiniDb

Proyecto universitario que usa implementacion de b+tree en disco y un pequeño approach de sql lang para crear, seleccionar e insertar a la base de datos.
## Propósito
   Entender y aplicar los algoritmos de almacenamiento de archivos físicos y el acceso concurrente.

## Definiciones previas
- Registro: Es la representación de un conjunto de atributos de un objeto o entiendad en particular.
- Archivo: Es la representación de una secuencia de registros, donde podría existir un orden particular.
- Índice: El índice es el identificador asociado a un registro, donde se detalla su clave primaria y la posición donde se ubica el registro en el archivo de datos.
- Archivo de datos: Es el conjunto de registros representados a través de sus atributos. 
- Archivo de  índice: Es la representacíon del  índice basado en el archivo de datos. Por ejemplo: ”clave,posicion” para el caso del índice usando Random File.
- Memoria principal: La memoria principal es uno de los componentes más importantes del ordenador. Se utiliza durante el tiempo de ejecucíon del programa y con la pérdida de fluido eléctrico se pierde la información. Un ejemplo para referirse a la memoria principal sería la Ramdon Access Memory (RAM). Además, el tiempo de acceso es independiente de la dirección.
- Memoria secundaria: En la memoria secundaria se almacena data donde, a pesar de no existir fluido eléctrico, se conserva. Un ejemplo para refereirse a la memoria secundaria sería el Disco Duro (HDD) o Disco Sólido (SSD). Además, el tiempo de acceso es dependiente de la dirección.

## Resultadoas esperados
Se implemetará el b+ tree en disco y se hará un benchmarking vs random file para probar su efectividad al momento de indexar data. 

## Técnicas implementadas

### B+ Tree
 - Descripción: Se irá guardando en archivos de texto plano en la memoria secundaria. Conforme va incrementando la cantidad de registros los cuales se van ingresando en el programa se van creando los buckets necesarios, estos tambien archivos de texto plano. Se usará una key como medio de indexación.
 
### Random File
 - Descripción: Random File emplea un archivo de índice, basado en un diccionario, clave-posición.La clave indica el registro al que se refiere y posición hace referencia a la posición donde se ubica el registro en el archivo de datos
 
# Resultados experimentales

En el siguiente proyecto se ha realizado un benchamarking entre ambos algoritmos. Para cada uno de ellos se tomó el tiempo entre insercion y búsqueda con tamaños de array con 10, 100, 10000 y 1 000 000 datos respectivamente, consiguiendo los siguientes resultados:

## Inserción





