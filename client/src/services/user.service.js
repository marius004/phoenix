import axios from "axios";
import config from "../config";
import authenticationService from "./authentication.service";

const userService = {
    getUserId,
    getGravatarEmailHash,
}

function getUserId() {
    const authToken = authenticationService.getAuthToken();
    return authToken?.user?.id
}

function getGravatarEmailHash(id) {
    return axios.get(`${config.apiUrl}/users/${id}/gravatar`);
}

export default userService;