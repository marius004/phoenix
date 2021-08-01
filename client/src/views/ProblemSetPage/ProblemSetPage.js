import Navbar from "components/Navbar/Navbar";
import React, { useEffect, useState } from "react";
import classNames from "classnames";
import useQuery from "hooks/query";
import { makeStyles } from "@material-ui/core/styles";
import styles from "assets/jss/material-kit-react/views/components.js";
import problemService from "services/problem.service";
import Problems from "views/Components/Problems";
import Footer from "components/Footer/Footer";
import Loading from "views/Loading";

const useStyles = makeStyles(styles);

const ProblemSetPage = () => {

    const [loading, setLoading] = useState(true);
    const [fetchingStatus, setFetchingStatus] = useState(200);

    const classes = useStyles();
    const [problems, setProblems] = React.useState([]);
    const query = useQuery();

    useEffect(() => {
        console.log(query.get("grade"))
    }, []);

    // TODO
    const getQueryFilters = () => {
        return {
            grade: query.get("grade"),
        };
    }

    const fetchProblems = async () => {
        try {
            const res = await problemService.getAll();
            setProblems(res.data);
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

    if (fetchingStatus === 404) {
        return <NotFound/>
    }

    if (fetchingStatus === 500) {
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

export default ProblemSetPage;