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
      <h2 style={{textAlign: "center", fontWeight: "bold"}}>Ce ofera Phoenix?</h2>
      <div className={classes.container}>
        <div className={classes.title}>
          <GridContainer justifyContent="center" style={{alignItems: "stretch"}}>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/graphs.gif").default} alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Peste 100 de probleme</h4>
                  <p>Oferim o categorie de probleme, de la structuri de date pana la programare dinamica si nu numai.</p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/rce.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Mediu de evaluare a submisiilor</h4>
                  <p>Sistemul de evaluare al submisiilor ofera posibilitatea de a rula codul sursa direct 
                    pe website, oferind feedback instant.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "226px" }} className={classes.imgCardTop} src={require("../../../assets/img/gopher.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Descrieri sugestive</h4>
                  <p>
                    Unele probleme nu sunt chiar asa de usor de rezolvat la inceput. De aceea oferim descrieri sugestive care sa va ajute atunci cand dati de greu.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/hacker.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Diversitate in limbaje de programare</h4>
                  <p>Nu toata lumea utilizeaza acelasi limbaj de programare. De aceea, oferim o gama larga din care ai ce sa alegi.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/coding.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Curs Design Patterns</h4>
                  <p>Design Patterns-urile sunt best practices care diferentiaza un programator bun de altul mai putin bun. 
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/concurrency.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Curs Concurrency</h4>
                  <p> Programarea concurenta este de multe ori o solutie inteligenta pentru a scala servere pentru un influx mai mare de requesturi.</p>
                </CardBody>
              </Card>
            </GridItem>
          </GridContainer>
        </div>
      </div>
    </div>
  );
}