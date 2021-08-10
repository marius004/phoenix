import axios from "axios";
import config from "config";

const submissionService = {
   getSubmissions,
};

function getSubmissions(query, page) {
    const limit  = 25;
    const offset = page * limit; 

    if (query == undefined || query == null || query == "")
        return axios.get(`${config.apiUrl}/submissions?limit=${limit}&offset=${offset}`)
    
    return axios.get(`${config.apiUrl}/submissions?${query}&limit=${limit}&offset=${offset}`)
}

export default submissionService;