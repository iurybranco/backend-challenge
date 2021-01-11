import chai, {expect} from 'chai';
import chaiHttp from 'chai-http';
import app from '../server';
import {createConnection} from 'typeorm'
import Product from "../database/models/product";
import {ServeMockDiscountServer} from "./mock_discount_server";

chai.use(chaiHttp);
chai.should();

describe("test get products", () => {
    before(async () => {
        const config = getMockDbConfig();
        await createConnection(config)
        ServeMockDiscountServer(3005)
    });
    it("should get all products from a existent user", (done) => {
        chai.request(app)
            .get('/product',)
            .query({
                userId: 1
            })
            .end((err, res) => {
                res.should.have.status(200);
                res.body.data.should.have.length(2);
                res.body.data.map((product: Product) => {
                    let mockProduct = getProductsWithDiscountMock().find((mockProduct: Product) => {
                        return mockProduct._id == product._id
                    })
                    expect(mockProduct).should.not.be.undefined
                    assertObject(product, mockProduct)
                })
                done()
            });
    });
    it("should get all products from a nonexistent user", (done) => {
        chai.request(app)
            .get('/product',)
            .query({
                userId: 5
            })
            .end((err, res) => {
                res.should.have.status(200);
                res.body.data.should.have.length(2);
                res.body.data.map((product: Product) => {
                    let mockProduct = getProductsWithoutDiscountMock().find((mockProduct: Product) => {
                        return mockProduct._id == product._id
                    })
                    expect(mockProduct).should.not.be.undefined
                    assertObject(product, mockProduct)
                })
                done()
            });
    });
});

function assertObject(object: Product, expected: Product) {
    object._id.should.be.eq(expected._id)
    object.price_in_cents.should.be.eq(expected.price_in_cents)
    object.title.should.be.eq(expected.title)
    object.description.should.be.eq(expected.description)
    object.discount.percentage.should.be.eq(expected.discount.percentage)
    object.discount.value_in_cents.should.be.eq(expected.discount.value_in_cents)
}

function getProductsWithoutDiscountMock(): Product[] {
    return [
        {
            "_id": 1,
            "price_in_cents": 1500,
            "title": "Soap",
            "description": "A smelly soap to take a shower",
            "discount": {
                "percentage": 0.00,
                "value_in_cents": 0
            }
        },
        {
            "_id": 2,
            "price_in_cents": 7000,
            "title": "Soda",
            "description": "The best soda ever",
            "discount": {
                "percentage": 0.00,
                "value_in_cents": 0
            }
        }
    ]
}

function getProductsWithDiscountMock(): Product[] {
    return [
        {
            "_id": 1,
            "price_in_cents": 1500,
            "title": "Soap",
            "description": "A smelly soap to take a shower",
            "discount": {
                "percentage": 12.00,
                "value_in_cents": 1000
            }
        },
        {
            "_id": 2,
            "price_in_cents": 7000,
            "title": "Soda",
            "description": "The best soda ever",
            "discount": {
                "percentage": 12.00,
                "value_in_cents": 1000
            }
        }
    ]
}

function getMockDbConfig() {
    const config: any = {
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
    return config;
}