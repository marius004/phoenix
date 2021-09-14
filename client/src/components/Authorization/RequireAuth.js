import React from 'react'
import authenticationUtil from 'util/authentication';

export default function RequireAuth(UnauthorizedComponent, AuthorizedComponent) {

    class RequireAuthentication extends React.Component {
        state = {
            isAuthenticated: authenticationUtil.isUserLoggedIn(),
        }

        render() {
            return !this.state.isAuthenticated ? <UnauthorizedComponent {...this.props}/> : <AuthorizedComponent {...this.props}/>
        }
    }

    return RequireAuthentication;
}