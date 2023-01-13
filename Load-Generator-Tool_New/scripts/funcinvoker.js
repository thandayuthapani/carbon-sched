import http from 'k6/http';
import { check, sleep } from "k6";

import papaparse from 'https://jslib.k6.io/papaparse/5.1.1/index.js';
import {
    SharedArray
} from "k6/data";

const funcData = new SharedArray("another data name", function () {
    return papaparse.parse(open('./data/func_count.csv')).data;
});
import {
    fetchWikiDetails, testK6, pythonTest
} from './functions/fetchData.js'


//NOTE: 1 VU - approax 50 req in 30s
const makeScenarios = function () {
    let scenarios = {}
    //total no of days each min 20162 min 
    // console.log(funcData.length)
    //total no of columns
    // console.log(funcData[1].length)
    //data: days - 14, perday - 24hrs
    //NOTE: Currently we are replicating requests for a single function for 30 min
    // console.log(funcData.length / (14 * 24 * 4))

    // running time (funcData.length-2) i.e 14 days
    for (let i = 1; i <= 15; i++) {
        let timer = 60 * (i - 1)
        let customTime = timer + "s"
        // iterate through all the columns i.e 160 funcData[1].length
        for (let j = 0; j < 1; j++) {
            // let functionname = "f_" + funcData[0][j]
            let functionname = "testPython"
            if (funcData[i][j] > 0) {
                scenarios['scenario' + i + Date.now()] = {
                    executor: 'shared-iterations',
                    startTime: customTime,
                    vus: Math.ceil(funcData[i][j] / 50),
                    iterations: funcData[i][j],
                    exec: functionname,
                    maxDuration: '60s',
                }
            }
        }
    }
    return scenarios
}

export const options = {
    scenarios: makeScenarios()
}

export function S01_FetchData() {
    testK6();
}

export function testPython() {
    pythonTest();
}


