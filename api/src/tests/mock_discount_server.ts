import * as grpc from '@grpc/grpc-js';
import {sendUnaryData} from '@grpc/grpc-js/build/src/server-call';
import {DiscountService, IDiscountServer} from "../discount/proto/discount_grpc_pb";
import {Request, Response} from "../discount/proto/discount_pb";

// @ts-ignore
class MockDiscountServer implements IDiscountServer {
    calculate(call: grpc.ServerUnaryCall<Request, Response>, callback: sendUnaryData<Response>) {
        let response = new Response().setPercentage(10).setValueInCents(100)
        callback(null, response);
    }

}

export function ServeMockDiscountServer(port: number): void {
    const server = new grpc.Server();
    // @ts-ignore
    server.addService<IDiscountServer>(DiscountService, new MockDiscountServer());
    const uri = `localhost:${port}`;
    server.bindAsync(uri, grpc.ServerCredentials.createInsecure(), (err, port) => {
        if (err) {
            throw err;
        }
        console.log(`Listening on ${port}`);
        server.start();
    });
}
