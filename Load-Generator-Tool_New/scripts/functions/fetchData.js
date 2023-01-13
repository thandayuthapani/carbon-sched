import http from 'k6/http';
import { check, sleep } from "k6";

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

export function pythonTest() {
    const response = http.get('http://helloworld-python.default.127.0.0.1.sslip.io');
    check(response, { "status is 200": (r) => r.status === 200 });
    sleep(.290);
}

export default {
    fetchWikiDetails, testK6, pythonTest
}