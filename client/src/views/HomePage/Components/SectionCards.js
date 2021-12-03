import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import styles from "assets/jss/material-kit-react/views/componentsSections/basicsStyle.js";
import GridContainer from "components/Grid/GridContainer";
import GridItem from "components/Grid/GridItem";
import Card from "components/Card/Card.js";
import CardBody from "components/Card/CardBody.js";

const useStyles = makeStyles(styles);

export default function SectionCards() {
  const classes = useStyles();
  return (
    <div className={classes.sections}>
      {/* TODO animate */}
      <h2 style={{textAlign: "center", fontWeight: "bold"}}>What do we offer?</h2>
      <div className={classes.container}>
        <div className={classes.title}>
          <GridContainer justifyContent="center" style={{alignItems: "stretch"}}>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/graphs.gif").default} alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Over 100 problems</h4>
                  <p>We offer a large range of problems, from dynamic programming to binary trees.</p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/rce.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Reliable code evaluation engine</h4>
                  <p>The remote code evaluation engine offers the possibility to evaluate the source code directly on our website.</p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "226px" }} className={classes.imgCardTop} src={require("../../../assets/img/gopher.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Suggestive problem descriptions</h4>
                  <p>
                    Some problems require immagination and some experiece to solve. Thus, the description comes in handy to help you when you are stuck.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/hacker.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Various programming languages</h4>
                  <p> We offer a large range of programming languages, from C++ to Python.</p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/coding.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Free Design Patterns Course</h4>
                  <p> At phoenix, we strongly believe that design pattern make the difference between a good software engineer and a great one.</p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/concurrency.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Free Concurrency Course</h4>
                  <p>Concurrent Programming is a smart solution that can be used to improve the performance of your application.</p>
                </CardBody>
              </Card>
            </GridItem>
          </GridContainer>
        </div>
      </div>
    </div>
  );
}