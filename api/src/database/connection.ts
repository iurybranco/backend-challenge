import {createConnection} from 'typeorm'


const config:any = {
    "type": "mongodb",
    "host": process.env.DB_HOST ? process.env.DB_HOST : "mongo.service.com.br",
    "database": process.env.DB_DATABASE ? process.env.DB_DATABASE : "challenge",
    "username": process.env.DB_USER ? process.env.DB_USER : "root",
    "port": process.env.DB_PORT ? process.env.DB_PORT : "27017",
    "password": process.env.DB_PASS ? process.env.DB_PASS : "dummyPass",
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

