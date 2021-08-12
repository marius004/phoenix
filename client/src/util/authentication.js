import Cookies from "js-cookie";
import base64 from "base-64";

const authenticationUtil = {
    getAuthToken,
    isUserLoggedIn,
}

function getAuthToken() {
    const data = Cookies.get("auth-token");

    if (data === undefined)
        return null;

    try {
        const decoded = base64.decode(data);
        return JSON.parse(decoded); 
    }catch(e) {
        console.error(e);
        return null;
    }
}

function isUserLoggedIn() {
    const authToken = getAuthToken();
    return authToken != null && authToken.token != null;
}

export default authenticationUtil;
