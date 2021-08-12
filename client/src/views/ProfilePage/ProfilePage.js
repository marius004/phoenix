import React, {useState, useEffect} from "react";
import {useParams} from "react-router-dom";
import classNames from "classnames";
import { makeStyles } from "@material-ui/core/styles";

import Footer from "components/Footer/Footer.js";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem.js";
import Parallax from "components/Parallax/Parallax.js";
import Navbar from "components/Navbar/Navbar";

import styles from "assets/jss/material-kit-react/views/profilePage.js";
import userAPI from "api/user";
import { unary } from "lodash";
import userUtil from "util/user";

const useStyles = makeStyles(styles);

export default function ProfilePage() {
    const { username } = useParams();

    const [user, setUser] = useState({})
    const [emailHash, setEmailHash] = useState("")

    const classes = useStyles();

    const imageClasses = classNames(
        classes.imgRaised,
        classes.imgRoundedCircle,
        classes.imgFluid
    );

    const fetchUser = async() => {
        try {
            const user = await userAPI.getByUsername(username);
            setUser(user);

            const emailHash = userUtil.calculateEmailHash(user.email)
            setEmailHash(emailHash);
        } catch(err) {
            console.error(err);
        }
    }

    const getGravatarURI = (imgSize) => {
        return `https://www.gravatar.com/avatar/${emailHash}?s=${imgSize}`;
    }

    useEffect(fetchUser, []);

    return (
        <div>
        <Navbar color="transparent" fixed ={false}/> 
        <Parallax
            small
            filter
            image={require("assets/img/profile-bg.jpg").default}
        />
        <div className={classNames(classes.main, classes.mainRaised)}>
            <div>
            <div className={classes.container}>
                <GridContainer justifyContent="center">
                <GridItem xs={12} sm={12} md={6}>
                    <div className={classes.profile}>
                    <div>
                        <img src={getGravatarURI(500)} alt="..." className={imageClasses} />
                    </div>
                    <div className={classes.name}>
                        <h3 className={classes.title}>{user.username}</h3>
                        <h5 style={{textTransform: "lowercase"}}>{user.email}</h5>
                    </div>
                    </div>
                </GridItem>
                </GridContainer>
                <div className={classes.description}>
                    <p>
                    {user.bio}
                    </p>
                </div>
               {/* TODO */}
            </div>
            </div>
        </div>
        <Footer />
        </div>
    );
}
