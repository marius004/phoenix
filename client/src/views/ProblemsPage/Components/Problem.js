import React from "react";
import PropTypes from "prop-types";
import MDEditor from '@uiw/react-md-editor';
import NavPills from "components/NavPills/NavPills";
import Dashboard from "@material-ui/icons/Dashboard";
import EqualizerIcon from '@material-ui/icons/Equalizer';
import {Button} from "@material-ui/core";
import {Link} from "react-router-dom";

const Problem = ({problem}) => {

    const problemLink = () => {
        return `/problems/${problem.name}`
    }

    return(
        <div style={{backgroundColor: "#e0f2f1", border: "2px solid #00695c", marginBottom: "12px", borderRadius: "4px", padding: "6px"}}>
            <NavPills
                color="primary"
                horizontal={{
                    tabsGrid:{xs: 12, sm: 5, md: 3},
                    contentGrid:{xs: 12, sm: 7, md: 9},
                }}
                tabs={[
                    {
                        tabButton: problem.visible ? problem.name : problem.name + " (unpublished) " ,
                        tabIcon: Dashboard,
                        tabContent: (
                            <div style={{maxHeight: "240px", overflow: "auto", padding: "6px"}}>
                                <MDEditor.Markdown style={{marginBottom: "20px"}} source={problem.shortDescription}/>
                                <Link to={problemLink()}>
                                    <Button variant="contained" color="primary">
                                        Solve
                                    </Button>
                                </Link>
                            </div>
                        ),
                    },  
                    {
                        tabButton: "About",
                        tabIcon: EqualizerIcon,
                        tabContent: (
                            <div style={{maxHeight: "240px", overflow: "auto", padding: "6px"}}>
                                <span>Coming soon...</span>
                            </div>
                        ),
                    }, 
                ]}
            />
        </div>
    );
}

Problem.propTypes = {
    problem: PropTypes.object.isRequired
};

export default Problem;