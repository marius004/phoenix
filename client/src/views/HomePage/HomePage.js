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


const updateProblem = (name) => {
  return axios.put(`${config.apiUrl}/problems/${name}`, {
    name: "marsx",
    description: "### ENUNT \n"
    +"Pe planeta Marte se desfășoară un intens război între trupele coloniștilor pământeni și triburile autohtone de reptilieni.\n"
    + "\n" 
    + "Pentru a dobândi controlul decisiv asupra planetei, pământenii au trimis trupe de commando în sistemul de galerii subterane de pe Marte, reprezentat sub forma unei configurații de `n` peșteri conectate prin `m` canale unidirecționale. Misiunea acestora este să ajungă din baza subterană a coloniștilor aflată în galeria `1` în orașul subteran al reptilienilor plasat în galeria `n` pentru a distruge rezistența extraterestră.\n" + 
    + "\n"
    + "Știind că fiecare dintre cele `m` canale dintre galerii este protejat de un anumit număr de soldați reptilieni, trupele de commando trebuie să aleagă rutele de deplasare cât mai vulnerabile în fața unui asalt (protejate de un număr minim de inamici).\n"
    + "\n"
    + "Se cere să aflați setul de galerii care vor fi obligatoriu traversate de pământeni pe parcursul misiunii."
    + "\n"
    + "### DATE DE INTRARE"
    + "\n"
    + "Se vor citi de la tastatură valorile `n` (numărul de galerii) și `m` (numărul de canale subterane), apoi `m` triplete de forma `u`, `v`, `w`, cu proprietatea că există canal unidirecțional între peștera `u` și peștera `v`, protejat de exact `w` soldați."
    + "\n"
    + "### DATE DE IEȘIRE"
    + "\n"
    + "Se va afișa pe prima linie valoarea `k` (numărul de galerii care vor fi obligatoriu traversate de coloniști), iar pe următoarea linie cele `k` galerii în ordine crescătoare." 
    + "\n"
    + "### RESTRICȚII ȘI PRECIZĂRI"
    + "\n"
    +"- Se garantează că există rută optimă între galeriile `1` și `n` pentru fiecare test."
    + "\n"
    + "- `1 <= n <= 100.000`"
    + "\n"
    + "- `1 <= m <= 200.000`" + "\n"
    + "- `1 <= u, v <= n`" + "\n"
    + "- `1 <= w <= 1.000.000.000`" + "\n"
    
    +"### EXEMPLU" + "\n"
    
    +"`Input`" + "\n"
    +"```" + "\n"
    +"5 6" + "\n"
    +"1 2 3" + "\n"
    +"1 3 4" + "\n"
    +"2 3 1" + "\n"
    +"2 4 5" + "\n"
    +"3 4 1" + "\n"
    +"4 5 8" + "\n"
    +"```" + "\n"
    
    +"`Output`" + "\n"
    +"```" + "\n"
    +"4" + "\n"
    +"1 3 4 5" + "\n"
    +"```" + "\n" + "\n" + "\n"
  }, config.cors)
}


  useEffect(async() => {
    // await updateProblem("sum01");
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
