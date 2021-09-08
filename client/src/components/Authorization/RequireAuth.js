import React from 'react'
import authenticationUtil from 'util/authentication';
import Unauthorized from 'views/Components/Unauthorized';

export default function RequireAuth(ComposedComponent) {

    class RequireAuthentication extends React.Component {
        state = {
            isAuthenticated: authenticationUtil.isUserLoggedIn(),
        }

        render() {
            return !this.state.isAuthenticated ? <Unauthorized/> : <ComposedComponent {...this.props}/>
        }
    }

    return RequireAuthentication;
}