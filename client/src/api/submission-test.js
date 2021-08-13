import axios from "axios";
import config from "config";

const submissionTestAPI = {
    getBySubmissionId,
};

function getBySubmissionId(submissionId) {
    return axios.get(`${config.apiUrl}/submissions/${submissionId}/tests`, config.cors)
        .then(res => res.data);
}

export default submissionTestAPI;