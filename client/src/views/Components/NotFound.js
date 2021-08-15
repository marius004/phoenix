import Navbar from "components/Navbar/Navbar";
import React from "react";

const NotFound = ({message}) => {
    return (<>
        <Navbar color = "white" fixed ={false}/> 
        <div style={{position: "absolute", top: "50%", left: "50%", transform: "translate(-50%, -50%)" }}>
            <img width="350px"  src={require("assets/img/404.png").default}/> 
            <h3>{message ? message : "Pagina cautata nu a fost gasita"}</h3>
        </div> 
    </>);
};

export default NotFound;