import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import styles from "assets/jss/material-kit-react/views/componentsSections/basicsStyle.js";
import PropTypes from "prop-types";
import Problem from "./Problem";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem";

const useStyles = makeStyles(styles);

const Problems = ({problems}) => {
    const classes = useStyles();

    return(
        <div className={classes.sections} style={{padding: "22px 26px"}}>
            <h3 style={{fontWeight: "bold"}}>
                Problems {"    "}
                <span style={{padding: "6px", color: "white", borderRadius: "5px" ,backgroundColor: "#4db6ac"}}>
                    {problems && problems.length ? problems.length : 0}
                </span>
            </h3>
            <GridContainer>
                {problems && 
                    problems.map((problem, index) => {
                        return (
                            <GridItem key={index} xl={12} md={12} sm={12}>
                              <Problem problem={problem} />
                            </GridItem>
                        )
                    })
                }
            </GridContainer>
        </div>
    );
};

Problems.propTypes = {
    problems: PropTypes.array.isRequired
};

export default Problems;