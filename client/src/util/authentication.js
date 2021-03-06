import Cookies from "js-cookie";
import base64 from "base-64";

const authenticationUtil = {
    getAuthToken,
    isUserLoggedIn,
    isUserAdmin,
    isUserProposer,
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

function isUserAdmin() {
    if (!isUserLoggedIn())
        return false;

    const authToken = getAuthToken();
    return authToken?.user?.isAdmin == true;
}

function isUserProposer() {
    if (!isUserLoggedIn())
        return false;

    const authToken = getAuthToken();
    return (authToken?.user?.isAdmin === true || authToken?.user?.isProposer === true);
}

export default authenticationUtil;
