import React from "react";
import PropTypes from "prop-types";

const SubmissionTab = (props) => {
    const {problem} = props;
    
    return (<>
        <h2>No Submissions</h2>
    </>);
};

SubmissionTab.propTypes = {
    problem: PropTypes.object.isRequired
}

export default SubmissionTab;