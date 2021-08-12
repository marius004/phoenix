import authenticationUtil from "./authentication";
import md5 from "md5";

const userUtil = {
    calculateEmailHash,
    getUser,
    getUserId,
}

function calculateEmailHash(email) {
    email = email.toLowerCase();
    email = email.trim();

    const hash = md5(email);
    return hash
} 

function getUser() {
    const authToken = authenticationUtil.getAuthToken();
    return authToken?.user
}

function getUserId() {
    const authToken = authenticationUtil.getAuthToken();
    return authToken?.user?.id
}

export default userUtil;