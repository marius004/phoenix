import axios from "axios";
import config from "config";

const submissionAPI = {
   getByQuery,
   getById,
   create,
   getByUserAndProblem,
};

function getByQuery(query, page) {
    const offset = page * config.submissionsLimit; 

    if (query == undefined || query == null || query == "")
        return axios.get(`${config.apiUrl}/submissions?limit=${config.submissionsLimit}&offset=${offset}`, config.cors)
            .then(res => res.data)

    return axios.get(`${config.apiUrl}/submissions?${query}&limit=${config.submissionsLimit}&offset=${offset}`, config.cors)
        .then(res => res.data)
}

function getById(id) {
    return axios.get(`${config.apiUrl}/submissions/${id}`, config.cors)
        .then(res => res.data)
}

function getByUserAndProblem(userId, problem) {
    return axios.get(`${config.apiUrl}/submissions?userId=${userId}&problem=${problem}`, config.cors)
        .then(res => res.data);
}

function create(code, language, problemId) {
    return axios.post(`${config.apiUrl}/submissions`, {
        lang: language,
        sourceCode: code,
        problemId: problemId
    }, config.cors)
        .then(res => res.data)
}

export default submissionAPI;