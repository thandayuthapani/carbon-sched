import http from 'k6/http';
import grpc from 'k6/net/grpc';
import { check, sleep } from "k6";

const client = new grpc.Client();
client.load(['./'], 'faasfunc.proto');

export function fetchWikiDetails() {
    console.log("fetchWikiDetails getting executed!")
    const response = http.get('https://en.wikipedia.org/wiki/Germany');
    check(response, { "status is 200": (r) => r.status === 200 });
    sleep(.300);
}

export function testK6() {
    //console.log("fetchtestK6 getting executed!")
    //const response = http.get('http://test.k6.io', { headers: { Accepts: "application/json" } });
    sleep(.400);
    const response = { status: 200 }
    check(response, { "status is 200": (r) => r.status === 200 });
    sleep(.300);
};

//export function pythonTest() {
//    const response = http.get('http://helloworld-python.default.127.0.0.1.sslip.io');
//    check(response, { "status is 200": (r) => r.status === 200 });
//    sleep(.290);
//}

export function pythonTest() {
    client.connect('http://helloworld-python.default.127.0.0.1.sslip.io', {
        plaintext: true
    });

    const data = { name: 'helloworld-python' };
    const response = client.invoke('Funcbench.FaaSFunc/InvokeFunc', data);

    check(response, {
        'status is OK': (r) => r && r.status === grpc.StatusOK,
    });

    console.log(JSON.stringify(response.message));
    client.close();
    sleep(.300);
}

export default {
    fetchWikiDetails, testK6, pythonTest
}