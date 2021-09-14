import React from 'react'
import authenticationUtil from 'util/authentication';

export default function RequireAdmin(UnauthorizedComponent, AuthorizedComponent) {

    class AdminRequired extends React.Component {
        state = {
            isAdmin: authenticationUtil.isUserAdmin(),
        }

        render() {
            return !this.state.isAdmin ? <UnauthorizedComponent {...this.props}/> : <AuthorizedComponent {...this.props}/>
        }
    }

    return AdminRequired;
}