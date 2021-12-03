import React, {useState} from "react";

import Header from "components/Header/Header.js";
import LoginDialog from "./LoginDialog.js";
import SignupDialog from "./SignupDialog.js";
import HeaderLinks from "components/Header/HeaderLinks.js";

export default function Navbar({color, fixed}) {
    const [showLogin, setShowLogin] = useState(false);
    const [showSignup, setShowSignup] = useState(false);

    const toggleLogin = () => {
        setShowLogin(!showLogin);
    }

    const toggleSignup = () => {
        setShowSignup(!showSignup);
    }

    return (
        <div>
             <Header
                brand={<span style={{ fontWeight: 500 }}>Phoenix</span>}
                rightLinks={<HeaderLinks onLogin = {toggleLogin} onSignup = {toggleSignup} />}
                fixed={fixed ? fixed : true}
                color={color ? color : "transparent"}
                changeColorOnScroll={{
                    height: 200,
                    color: "white",
                }}
            />
            <LoginDialog show={showLogin} onLogin = {toggleLogin}/>
            <SignupDialog show={showSignup} onSignup = {toggleSignup}/>
        </div>
    ); 
}