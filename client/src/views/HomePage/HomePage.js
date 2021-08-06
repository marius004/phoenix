import React, { useEffect } from "react";
import classNames from "classnames";
import { makeStyles } from "@material-ui/core/styles";
import Footer from "components/Footer/Footer.js";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem.js";
import Parallax from "components/Parallax/Parallax.js";

import styles from "assets/jss/material-kit-react/views/components.js";
import SectionFeatures from "./Sections/SectionCards.js";
import Navbar from "components/Navbar/Navbar.js";
import axios from "axios";
import config from "config";

const useStyles = makeStyles(styles);

export default function Components() {
  const classes = useStyles();


const updateProblem = () => {
  return axios.post(`${config.apiUrl}/submissions`, {
    "lang": "c",
    "problemId": 1,
    "sourceCode": `#include <stdio.h>\nint main() {\nint number1, number2, sum;\n scanf("%d %d", &number1, &number2);\n sum = number1 + number2; printf("%d", sum); return 0;\n}\n`
  }, config.cors)
}


  useEffect(async() => {
    await updateProblem();
  }, [])

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
                  style={{ fontWeight: 400 }} 
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
