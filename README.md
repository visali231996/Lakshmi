what is in this repositiry:
In main.go:
1.i have done the declarations of all data types(int,char,float 32).i also used different loops like for loop, switch case. used defer to print those at the last.
2. i used examples for for loop,switch and if statements
3.i used pointer examples
4. i have created an arrays,structures and slices
5. i have called all the functions in the another folder

datatype declaration:
var a int=6 or a:=6
as:="fqef"
ad :=3.45 or var bb float 32 =5.575
for i:=0;i<6;i++{
cndition}
for i:=range 10{condition}
defer fmt.Println("dfsrtgrtfb")

Identifier or variable:
declaration: var variablename type= value  or variablename:=value
variablename:=value..........can not be declared globally
default values:
int: 0
string:
float32:0.0000000
boolean:false
fmt.Println("type of the string %T\n",variablename)

ARRAY.......This datatype is having fixed capacity storing same data values....used for matrix creations, game boards, storing exam scores
slices.........having variable capacity..can be used when we dont knpow the exact quantity,... in case of college admissions..noting  seismic waves,weather conditions
method....is having same usage as function but 

channel.....
declarations:  var Channel_name chan Type or channel_name:= make(chan Type)

select statement: are mainly used in channels

select {
case v1: <-c1
case v2:<-c2
}

SQL Commands
•	SELECT * FROM Tablename;
•	SELECT DISTINCT colname FROM Tablename ;--- gives only distinct col values from table
•	SELECT COUNT(DISTINCT Country) FROM Customers;-----counts didtinct countries from table
•	SELECT  * FROM customers WHERE orderid>60
•	SELECT * FROM Customers OREDER BY price DESC
•	SELECT column1, column2, ...
FROM table_name
WHERE condition1 AND condition2 AND condition3 ...;
•	SELECT * FROM Customers
WHERE Country = 'Spain' AND (CustomerName LIKE 'G%' OR CustomerName LIKE 'R%');
•	SELECT * FROM Customers
WHERE NOT Country = 'Spain';
