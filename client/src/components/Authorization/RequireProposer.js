import React from 'react'
import authenticationUtil from 'util/authentication';

export default function RequireProposer(UnauthorizedComponent, AuthorizedComponent) {

    class ProposerRequired extends React.Component {
        state = {
            isProposer: authenticationUtil.isUserProposer(),
        }

        render() {
            return !this.state.isProposer ? <UnauthorizedComponent {...this.props}/> : <AuthorizedComponent {...this.props}/>
        }
    }

    return ProposerRequired;
}