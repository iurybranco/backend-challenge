import {createConnection} from 'typeorm'

createConnection().catch(err => {
    console.log("failed to connect to database: ", err)
})