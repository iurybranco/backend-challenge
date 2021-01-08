import {ErrorRequestHandler} from 'express'
import jsend from "jsend";

const errorHandler: ErrorRequestHandler = (error, req, res, next) => {
    console.log(error)
    return res.status(500).json(jsend.error("data has been lost on server"))
}

export default errorHandler