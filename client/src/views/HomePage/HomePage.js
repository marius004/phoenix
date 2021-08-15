import React, { useEffect } from "react";
import classNames from "classnames";
import { makeStyles } from "@material-ui/core/styles";
import Footer from "components/Footer/Footer.js";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem.js";
import Parallax from "components/Parallax/Parallax.js";

import styles from "assets/jss/material-kit-react/views/components.js";
import SectionFeatures from "./Components/SectionCards.js";
import Navbar from "components/Navbar/Navbar.js";

const useStyles = makeStyles(styles);

export default function Components() {
  const classes = useStyles();

  return (
    <div>
      <Navbar/>
      <Parallax image={require("assets/img/bg4.jpg").default}>
        <div className={classes.container}>
          <GridContainer>
            <GridItem>
              <div className={classes.brand}>
                <h1 className={classes.title}>Phoenix</h1>
                <h3 
                  style={{ fontWeight: 500 }} 
                  className={classes.subtitle}
                  >
                  O platforma moderna pentru a invata programare
                </h3>
              </div>
            </GridItem>
          </GridContainer>
        </div>
      </Parallax>

      <div className={classNames(classes.main, classes.mainRaised)}>
        <SectionFeatures/>
      </div>
      <Footer />
    </div>
  );
}
