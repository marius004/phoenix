import authenticationUtil from "./authentication";

const problemUtil = {
  canUserEditProblem,
}

function canUserEditProblem(problem) {
    if(!authenticationUtil.isUserLoggedIn()) 
        return false;

    if(!problem) 
        return false;
    
    const user = authenticationUtil.getAuthToken()?.user;
    return user.isAdmin || (user?.id !== null && problem?.authorId === user?.id);
}

export default problemUtil;
