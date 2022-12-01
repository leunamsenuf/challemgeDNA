# challemgeDNA
BY manuelfunes@yahoo.com.br
VERSION 0.0


Teste – Capgemini
Você foi contratado para desenvolver um projeto em Go, Java, Python ou Javascript (NodeJS),
que vai identificar se uma sequência de letras é válida.
O projeto consiste em desenvolver uma API REST, e disponibilizar um endpoint HTTP POST
"/sequence". Esse endpoint receberá como parâmetro, um JSON com a sequência de letras
(Array de Strings), onde, cada elemento desse array representa uma linha de uma tabela
quadrada de (NxN), como no exemplo abaixo:

```
POST -> /sequence
{
"letters": ["DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", " DDDDUB", "UDBDUH"]
}

Você saberá se uma sequência é válida, se encontrar 2 ou mais sequências de quatro
letras iguais em qualquer direção, horizontal, vertical ou nas diagonais.
As letras da String só podem ser: (B, U, D, H)
A API deve retornar um json com "is_valid": boolean. Caso você identifique uma sequência
válida, deve ser true, caso identifique uma sequência inválida, deve ser false, como no
exemplo abaixo:
HTTP 200
{"is_valid": true}


'''
