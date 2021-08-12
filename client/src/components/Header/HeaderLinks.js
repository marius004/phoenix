import React from "react";
import { Link } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import {Apps, PermIdentity} from "@material-ui/icons";
import CustomDropdown from "components/CustomDropdown/CustomDropdown.js";
import Button from "components/CustomButtons/Button.js";
import styles from "assets/jss/material-kit-react/components/headerLinksStyle.js";
import MenuBookIcon from '@material-ui/icons/MenuBook';

import authenticationAPI from "api/authentication";
import authenticationUtil from "util/authentication";
import userUtil from "util/user";

const useStyles = makeStyles(styles);

export default function HeaderLinks(props) {
  const classes  = useStyles();
  const authToken = authenticationUtil.getAuthToken();

  const handleLogout = async() => {
    await authenticationAPI.logout();
    window.location.reload();
  }

  const getUserProfileLink = () => {
    const user = userUtil.getUser();
    return `/profile/${user.username}`
  }

  const redirectToProfilePage = () => {
    window.location.href = getUserProfileLink();
  }

  return (
    <List className={classes.list}>
      <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText="Probleme"
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={Apps}
          dropdownList={[
            <Link to="/problems" className={classes.dropdownLink}>
               Toate 
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
                Clasa a IX-a
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
                Clasa a X-a
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
                Clasa a XI-a
            </Link>,
             <Link to="/" className={classes.dropdownLink}>
               Concursuri
            </Link>,
          ]}
        />
      </ListItem>
      <ListItem className={classes.listItem}>
        <Link to="/submissions" className={classes.navLink}>
            Submissions
        </Link>
      </ListItem>
      <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText="Resurse"
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={MenuBookIcon}
          dropdownList={[
            // TODO
            <Link to="/" className={classes.dropdownLink}>
                Competitive Programming
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
                Web Development
            </Link>,
          ]}
        />
      </ListItem>
      {
        !authenticationUtil.isUserLoggedIn() &&
        <>
          <ListItem className={classes.listItem}>
            <Button style={{ paddingLeft: "8px", paddingRight: "8px"}} color="transparent" onClick={props.onSignup}>
              <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-user" aria-hidden="true"></i>
              Inregistrare
            </Button>
          </ListItem>
          <ListItem className={classes.listItem}>
            <Button style={{ paddingLeft: "8px", paddingRight: "8px" }} color="transparent" onClick={props.onLogin}>
              <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-sign-in" aria-hidden="true"></i>
              Autentificare
            </Button>
          </ListItem>
        </>
      }
      {
        authenticationUtil.isUserLoggedIn() &&
        <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText={authToken && authToken.user.username !== "" ? authToken.user.username : "User"}
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={PermIdentity}
          dropdownList={[
            <ListItem className={classes.listItem}>
              <Button
                  onClick={redirectToProfilePage} 
                  className={classes.dropdownLink} 
                  style={{ paddingLeft: "8px", paddingRight: "8px"}} 
                  color="transparent">
                <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-user" aria-hidden="true"></i>
                Profile
              </Button>
            </ListItem>,
            <ListItem className={classes.listItem}>
              <Button className={classes.dropdownLink} style={{ paddingLeft: "8px", paddingRight: "8px"}} color="transparent" onClick={() => handleLogout()}>
                <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-sign-out"></i>
                Log out
              </Button>
          </ListItem>,
          ]}
        />
        </ListItem> 
      }
    </List>
  );
}
