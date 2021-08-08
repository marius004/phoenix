import axios from "axios";
import md5 from "md5";
import config from "../config";
import authenticationService from "./authentication.service";

const userService = {
    getUserId,
    getGravatarEmailHash,
    getByUserName,
    calculateEmailHash,
    getUser,
}

function getByUserName(username) {
    return axios.get(`${config.apiUrl}/users/${username}`)
}

function getUserId() {
    const authToken = authenticationService.getAuthToken();
    return authToken?.user?.id
}

function getUser() {
    const authToken = authenticationService.getAuthToken();
    return authToken?.user
}

function getGravatarEmailHash(id) {
    return axios.get(`${config.apiUrl}/users/${id}/gravatar`);
}

function calculateEmailHash(email) {
    email = email.toLowerCase();
    email = email.trim();

    const hash = md5(email);
    return hash
} 

export default userService;