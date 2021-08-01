import React from "react";
import PropTypes from "prop-types";

// TODO
const ProblemFilter = ({setProblems}) => {
    return (
        <div style={{backgroundColor: "#e0f2f1", border: "2px solid #00695c", marginBottom: "12px", borderRadius: "4px"}}>
            <h3 style={{fontWeight: "bold", textAlign: "center"}}>Filtre</h3>
            <h3>TODO</h3>
        </div>
    );
};

ProblemFilter.propTypes = {
    setProblems: PropTypes.func.isRequired,
};

export default ProblemFilter;