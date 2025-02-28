const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync('./greeter.proto', {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const helloProto = grpc.loadPackageDefinition(packageDefinition).greeter;

function main() {
    const client = new helloProto.Greeter('localhost:50051', grpc.credentials.createInsecure());

    const name = process.argv.length >= 3 ? process.argv[2] : 'world from nodejs';
    client.SayHello({ name: name }, function(err, response) {
        if (err) {
            console.error(err);
            return;
        }
        console.log('Greeting:', response.message);
    });
}

main();
