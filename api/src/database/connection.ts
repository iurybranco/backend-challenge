import {createConnection} from 'typeorm'

const config:any = {
    "type": "mongodb",
    "host": "mongo.service.com.br",
    "database": "challenge",
    "username": "root",
    "port": 27017,
    "password": "dummyPass",
    "timezone": "Z",
    "entities": [
        "./src/database/models/*.ts",
        "./dist/database/models/*.js"
    ],
    "extra": {
        "authSource": "admin",
        "useUnifiedTopology": true
    }
}
createConnection(config).catch(err => {
    console.log("failed to connect to database: ", err.name)
    process.exit(0);
})

