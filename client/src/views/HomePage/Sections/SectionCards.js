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
      <h2 style={{textAlign: "center", fontWeight: "bold"}}>Ce este Phoenix?</h2>
      <div className={classes.container}>
        <div className={classes.title}>
          <GridContainer justifyContent="center" >
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/graphs.gif").default} alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Peste 100 de probleme</h4>
                  <p>Oferim o larga categorie de probleme, de la structuri de date avansate pana la programare dinamica,
                    dar si de la probleme de bac pana la cele de interviu.</p>
                {/* TODO add link to the problems page */}
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/rce.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Mediu de evaluare a submisiilor</h4>
                  <p>Sistemul de evaluare al submisiilor iti ofera posibilitatea de a rula codul sursa direct 
                    pe website, oferind in timp real feed-back.
                  </p>
                  {/* TODO about page */}
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/gopher.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Descrieri sugestive</h4>
                  <p>Algoritmica este dificil de aprofundat doar folosind pixul si hartia. Pe site-ul nostru, fiecare problema este
                    insotita de o descriere sugestiva, ideala pentru asimilarea intuitiva a notiunilor noi.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/hacker.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Diversitate in limbaje de programare</h4>
                  <p>Nu toata lumea utilizeaza acelasi limbaj de programare. De aceea, Phoenix vine in ajutorul vostru cu
                    o gama larga de limbaje de programare: de la clasicul C procedural la lambda calculus din Haskell si OOP-ul din Kotlin sau Java.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/coding.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Curs Design Patterns</h4>
                  <p>Design Patterns constituie solutii ideale pentru probleme comune din software design, fiind un instrument
                    esential si extrem de flexibil in crearea de aplicatii software scalabile.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
            <GridItem lg={4} md={6} sm={12}>
              <Card>
                <img style={{ height: "250px" }} className={classes.imgCardTop} src={require("../../../assets/img/concurrency.gif").default}  alt="Card-img-cap" />
                <CardBody>
                  <h4 className={classes.cardTitle}>Curs Concurrency</h4>
                  <p>Stapanirea abilitatilor de programare concurenta este o necesitate in procesul de creare a aplicatiilor
                    moderne, crescand scalabilitatea acestora prin posibilitatea de a rula doua procese in acelasi timp.
                  </p>
                </CardBody>
              </Card>
            </GridItem>
          </GridContainer>
        </div>
      </div>
    </div>
  );
}