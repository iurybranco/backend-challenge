#IT INSERT USERS DUMP INTO MONGODB
mongoimport --host mongodb --db challenge --collection user --authenticationDatabase admin --username root --password dummyPass --drop --file dump/users.json --jsonArray
#IT INSERT PRODUCTS DUMP INTO MONGODB
mongoimport --host mongodb --db challenge --collection product --authenticationDatabase admin --username root --password dummyPass --drop --file dump/products.json --jsonArray
