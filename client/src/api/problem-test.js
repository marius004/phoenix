import config from "config";
import axios from "axios";

const testAPI = {
    getProblemTests,
};

function getProblemTests(problemName) {
    return axios.get(`${config.apiUrl}/problems/${problemName}/tests`, cors.config)
        .then(res => res.data);
}

export default testAPI;