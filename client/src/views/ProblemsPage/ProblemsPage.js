import Navbar from "components/Navbar/Navbar";
import React, { useEffect, useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import styles from "assets/jss/material-kit-react/views/components.js";
import classNames from "classnames";
import useQuery from "hooks/query";
import Problems from "views/ProblemsPage/Components/Problems";
import Footer from "components/Footer/Footer";
import Loading from "views/Components/Loading";

import problemAPI from "api/problem";

const useStyles = makeStyles(styles);

const ProblemsPage = () => {

    const [loading, setLoading] = useState(true);
    const [fetchingStatus, setFetchingStatus] = useState(200);

    const classes = useStyles();
    const [problems, setProblems] = React.useState([]);
    const query = useQuery();

    const fetchProblems = async () => {
        try {
            const problems = await problemAPI.getAll();
            setProblems(problems);
        }catch(err) {
            if (err.message == "Network Error") {
                setFetchingStatus(500);
            } else {
                const status = err.response.status;
                setFetchingStatus(status);
            }
        } finally {
            setLoading(false);
        }
    }

    useEffect(fetchProblems, []);

    if (loading) {
        return <Loading/>
    }

    // 4xx (bad request or not found)
    if (Math.round(fetchingStatus / 100) == 4) {
        return <NotFound message="Problema cautata nu a fost gasita"/>   
    }

    // 5xx (internal server error)
    if (Math.round(fetchingStatus / 100) === 5) {
        return <InternalServerError/>
    }

    return (<>
        <Navbar color="white" fixed={false}/> 
        <div style={{marginTop: "100px"}} className={classNames(classes.main, classes.mainRaised)}>
            <Problems problems={problems}/>
        </div>
        <Footer/>
    </>);
}

export default ProblemsPage;