import {Request, Response, Router} from "express";

const routes = Router()

routes.get("/product", function (req: Request, res: Response) {
    return res.status(200).json({"status": "success"})
})

export default routes