import React from  "react";
import PropTypes from "prop-types";
import DifficultyBadge from "../../Components/DifficultyBadge";
import MDEditor from '@uiw/react-md-editor';
import GradeBadge from "views/Components/GradeBadge";

export default function ProblemTab(props) {
    const {problem} = props;
    return(<>
        <h3 style={{fontWeight: "bold"}}>Problema <span style={{padding: "6px", borderRadius: "5px" ,backgroundColor: "#dce775"}}>{problem.name}</span></h3>
        <DifficultyBadge difficulty={problem.difficulty}/>
        <GradeBadge grade={problem.grade}/>
        <hr/>
        <MDEditor.Markdown style={{marginBottom: "20px"}} source={problem.description} />
    </>);
}

ProblemTab.propTypes = {
    problem: PropTypes.object.isRequired
}
