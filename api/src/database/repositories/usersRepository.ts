import {getMongoRepository} from "typeorm";
import User from "../models/user";
import {IUsersRepository} from "./interfaces";

export default class UsersRepository implements IUsersRepository {

    async get(id: number): Promise<User | undefined> {
        return await getMongoRepository(User).findOne({_id: id});
    }

}