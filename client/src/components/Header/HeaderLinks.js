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
import { Divider } from "@material-ui/core";

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

  const authedUserLinks = () => {
    const links = [
      <ListItem className={classes.listItem}>
        <Button
            onClick={redirectToProfilePage} 
            className={classes.dropdownLink} 
            style={{ paddingLeft: "8px", paddingRight: "8px"}} 
            color="transparent">
          <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-user" aria-hidden="true"></i>
          Profile
        </Button>
      </ListItem>
    ];

    if(authenticationUtil.isUserProposer()) {
      links.push(
        <Divider/>,
        <ListItem className={classes.listItem}>
          <Button className={classes.dropdownLink} style={{ paddingLeft: "8px", paddingRight: "8px"}} color="transparent" onClick={() => window.location.href = "/proposer"}>
          <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-pencil" aria-hidden="true"></i>
            Proposer Panel
          </Button>
        </ListItem>,
      );
    }

    if(authenticationUtil.isUserAdmin()) {
      links.push(
        <ListItem className={classes.listItem}>
          <Button className={classes.dropdownLink} style={{ paddingLeft: "8px", paddingRight: "8px"}} color="transparent" onClick={() => window.location.href = "/admin"}>
          <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-shield" aria-hidden="true"></i>
            Admin Panel
          </Button>
        </ListItem>,
      );
    }

    links.push(
      <Divider/>,
      <ListItem className={classes.listItem}>
        <Button className={classes.dropdownLink} style={{ paddingLeft: "8px", paddingRight: "8px"}} color="transparent" onClick={() => handleLogout()}>
          <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-sign-out"></i>
          Log out
        </Button>
      </ListItem>
    );

    return links;
  }

  return (
    <List className={classes.list}>
      <ListItem className={classes.listItem}>
        <Link to="/problems" className={classes.navLink}>
          Problems
        </Link>
      </ListItem>
      <ListItem className={classes.listItem}>
        <Link to="/submissions" className={classes.navLink}>
          Submissions
        </Link>
      </ListItem>
      <ListItem className={classes.listItem}>
        <CustomDropdown
          noLiPadding
          buttonText="Resourses"
          buttonProps={{
            className: classes.navLink,
            color: "transparent",
          }}
          buttonIcon={MenuBookIcon}
          dropdownList={[
            // TODO
            <Link to="/" className={classes.dropdownLink}>
              Design Patterns
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
              Competitive Programming
            </Link>,
            <Link to="/" className={classes.dropdownLink}>
              Concurrency
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
              Sign up
            </Button>
          </ListItem>
          <ListItem className={classes.listItem}>
            <Button style={{ paddingLeft: "8px", paddingRight: "8px" }} color="transparent" onClick={props.onLogin}>
              <i style={{ fontSize: "18px", marginRight: "4px" }} className="fa fa-sign-in" aria-hidden="true"></i>
              Log in
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
          dropdownList={authedUserLinks()}
        />
        </ListItem> 
      }
    </List>
  );
}
