FROM mongo:4.2.8

COPY users.json products.json /
CMD mongoimport --host mongodb --db challenge --collection user --authenticationDatabase admin --username root --password dummyPass --drop --file users.json --jsonArray
CMD mongoimport --host mongodb --db challenge --collection product --authenticationDatabase admin --username root --password dummyPass --drop --file products.json --jsonArray
