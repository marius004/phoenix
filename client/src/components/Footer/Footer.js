/*eslint-disable*/
import React from "react";
// nodejs library to set properties for components
import PropTypes from "prop-types";
// nodejs library that concatenates classes
import classNames from "classnames";
// material-ui core components
import { List, ListItem } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";

import GitHubIcon from '@material-ui/icons/GitHub';
import InfoIcon from '@material-ui/icons/Info';
import EmailIcon from '@material-ui/icons/Email';

import styles from "assets/jss/material-kit-react/components/footerStyle.js";

const useStyles = makeStyles(styles);

export default function Footer(props) {
  const classes = useStyles();
  const { whiteFont } = props;
  const footerClasses = classNames({
    [classes.footer]: true,
    [classes.footerWhiteFont]: whiteFont,
  });
  const aClasses = classNames({
    [classes.a]: true,
    [classes.footerWhiteFont]: whiteFont,
  });
  return (
    <footer className={footerClasses}>
      <div className={classes.container}>
        <div className={classes.left}>
          <List className={classes.list}>
            <ListItem className={classes.inlineBlock}>
              <a href="https://github.com/marius004/phoenix" 
                className={classes.block}
                target="_blank">
                <GitHubIcon style={{ fontSize: 16, transform: "translate(0, 2px)" }} /> {"  "}
                Phoenix
              </a>
            </ListItem>
            <ListItem className={classes.inlineBlock}>
              <a href="/about" className={classes.block}>
                <InfoIcon style={{ fontSize: 18, transform: "translate(0, 4px)" }} /> {"  "}
                 About
              </a>
            </ListItem>
            <ListItem className={classes.inlineBlock}>
              <a
                href="mailto:scarlatmariusstefan2018@gmail.com"
                className={classes.block}
              >
              <EmailIcon style={{ fontSize: 18, transform: "translate(0, 4px)" }} /> {"  "}
                Contact
              </a>
            </ListItem>
          </List>
        </div>
      </div>
    </footer>
  );
}

Footer.propTypes = {
  whiteFont: PropTypes.bool,
};
