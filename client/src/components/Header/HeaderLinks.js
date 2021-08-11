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
import authenticationService from "services/authentication.service";
import userService from "services/user.service";

const useStyles = makeStyles(styles);

export default function HeaderLinks(props) {
  const classes  = useStyles();
  const authToken = authenticationService.getAuthToken();

  const handleLogout = async() => {
    await authenticationService.logout();
    window.location.reload();
  }

  const getUserProfileLink = () => {
    const user = userService.getUser();
    return `/profile/${user.username}`
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
        !authenticationService.isLoggedIn() &&
        <>
          <ListItem className={classes.dropdownLink}>
            <Button color="transparent" onClick={props.onSignup}>
              <i style={{ fontSize: "18px"}} className="fa fa-user" aria-hidden="true"/>
              Inregistrare
            </Button>
          </ListItem>
          <ListItem className={classes.dropdownLink}>
            <Button color="transparent" onClick={props.onLogin}>
              <i style={{ fontSize: "18px" }} className="fa fa-sign-in" aria-hidden="true"/>
              Autentificare
            </Button>
          </ListItem>
        </>
      }

      {
        authenticationService.isLoggedIn() &&
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
            <Link to={getUserProfileLink()} className={classes.dropdownLink}>
              <Button className={classes.dropdownLink} color="transparent">
                <i style={{ fontSize: "18px"}} className="fa fa-user" aria-hidden="true"></i>
                Profile
              </Button>
            </Link>,
            <Button className={classes.dropdownLink} color="transparent" onClick={() => handleLogout()}>
              <i style={{ fontSize: "18px"}} className="fa fa-sign-out"></i>
              Log out
            </Button>
          ]}
        />
        </ListItem> 
      }
    </List>
  );
}
