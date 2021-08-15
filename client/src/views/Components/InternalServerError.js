import React from "react";
import Navbar from "components/Navbar/Navbar";

const InternalServerError = () => {
    return (<>
        <Navbar color = "white" fixed ={false}/> 
        <div style={{position: "absolute", top: "50%", left: "50%", transform: "translate(-50%, -50%)" }}>
            <img width="350px"  src={require("assets/img/500.jpg").default}/> 
            <h3 style={{textAlign: "center"}}>
                Contacteaza un admin!! <br/>
                <a
                    href="mailto:scarlatmariusstefan2018@gmail.com"
                    style={{textAlign: "center", color: "black", fontSize: "1em", fontWeight: "bold"}}>
                        scarlatmariusstefan2018@gmail.com
                </a>
            </h3>
        </div>
    </>);
};

export default InternalServerError;