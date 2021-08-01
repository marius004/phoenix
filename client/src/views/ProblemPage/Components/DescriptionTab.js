import React from "react";
import PropTypes from "prop-types";

const DescriptionTab = (props) => {
    const {problem} = props;
    
    return(<>
        <h1>{problem.name}</h1>    
    </>);
};

DescriptionTab.propTypes = {
    problem: PropTypes.object.isRequired
}

export default DescriptionTab;