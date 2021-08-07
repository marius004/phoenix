import React from "react";
import GridContainer from "components/Grid/GridContainer.js";
import GridItem from "components/Grid/GridItem";
import PropTypes from "prop-types";
import CustomDropdown from "components/CustomDropdown/CustomDropdown.js";
import styles from "assets/jss/material-kit-react/components/headerLinksStyle.js";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(styles);

export default function EditorSettings(props) {

    const classes = useStyles();

    const setFontSize = props.setFontSize;
    const setLang = props.setLang;
    const setEditorColor = props.setColor;

    return (
        <GridContainer style={{textAlign: "center", marginBottom: "12px"}}>
            <GridItem xl={4} lg={4} sm={4}>
                <CustomDropdown
                  buttonText="Font Size"
                  buttonProps={{color: "warning"}}
                  buttonStyle={{diplay: "inline-block"}}
                  dropdownList={[
                    <span className={classes.dropdownLink} onClick={() => setFontSize(18)}>
                        18
                    </span>, {divider: true},
                    <span className={classes.dropdownLink} onClick={() => setFontSize(19)}>
                        19
                    </span>, {divider: true},
                    <span className={classes.dropdownLink} onClick={() => setFontSize(20)}>
                        20
                    </span>, {divider: true},
                    <span className={classes.dropdownLink} onClick={() => setFontSize(21)}>
                        21
                    </span>, {divider: true},
                    <span className={classes.dropdownLink} onClick={() => setFontSize(22)}>
                        22
                    </span>,
                    ]}
                />
            </GridItem>
            <GridItem xl={4} lg={4} sm={4}>
                <CustomDropdown
                  buttonText="Programming Language"
                  buttonProps={{color: "info"}}
                  buttonStyle={{diplay: "inline-block"}}
                  dropdownList={[
                    <span className={classes.dropdownLink} onClick={() => setLang("c")}>
                        C 
                    </span>,
                  ]}
                />
            </GridItem>
            <GridItem xl={4} lg={4} sm={4}>
                {/* TODO ENGLEZIFICARE */}
                <CustomDropdown
                  buttonText="Background"
                  buttonProps={{color: "success"}}
                  buttonStyle={{diplay: "inline-block"}}
                  dropdownList={[
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#e0e0e0")}>
                        Gri 
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#ffe0b2")}>
                        Portocaliu
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#fff9c4")}>
                        Galben
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#e8f5e9")}>
                        Verde deschis
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#c8e6c9")}>
                        Verde inchis
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#e1f5fe")}>
                        Albastru deschis
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#b3e5fc")}>
                        Albastru inchis
                    </span>,{divider: true},
                    <span className={classes.dropdownLink} onClick={() => setEditorColor("#e6ee9c")}>
                        Lime
                    </span>,
                    ]}
                />
            </GridItem>
        </GridContainer>
    );
}

EditorSettings.propTypes = {
    setFontSize: PropTypes.func.isRequired,
    setLang: PropTypes.func.isRequired,
    setColor: PropTypes.func.isRequired,
}