import config from "config";
import axios from "axios";

const testAPI = {
    getProblemTests,
    createProblemTest,
};

function getProblemTests(problemName) {
    return axios.get(`${config.apiUrl}/problems/${problemName}/tests`, 
        config.cors).then(res => res.data);
}

function createProblemTest(problemName, score, input, output) {
    return axios.post(`${config.apiUrl}/problems/${problemName}/tests`, {
        score,
        input: input,
        output: output,
    }, config.cors).then(res => res.data);
}

export default testAPI;