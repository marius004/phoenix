import axios from "axios";
import config from "../config";

const evaluatorService = {
    createSubmission,
    getSubmissions,
}

function createSubmission(code, language, problemId) {
    return axios.post(`${config.apiUrl}/submissions`, {
        lang: language,
        sourceCode: code,
        problemId: problemId
    }, config.cors);
}

function getSubmissions(userId, problemId) {
    return axios.get(`${config.apiUrl}/submissions?userId=${userId}&problemId=${problemId}`, config.cors)
}

export default evaluatorService;