import React, {useState, useEffect} from "react";

import Navbar from "components/Navbar/Navbar";
import Grid from '@material-ui/core/Grid';

import { makeStyles } from "@material-ui/core/styles";
import Alert from '@material-ui/lab/Alert';
import classNames from "classnames";
import styles from "assets/jss/material-kit-react/views/components.js";
import SubmissionFilter from "./Components/SubmissionFilter";
import { useHistory } from "react-router-dom";
import SubmissionTable from "./Components/SubmissionTable";
import Footer from "components/Footer/Footer";
import Loading from "views/Components/Loading";
import useQuery from "hooks/query";

import submissionAPI from "api/submission";

const useStyles = makeStyles(styles);

export default function SubmissionsPage() {
    const classes = useStyles();
  
    const query = useQuery();
    const history = useHistory();

    const [submissions, setSubmissions] = useState([]);
    const [loading, setLoading] = useState(true);

    const [username, setUsername] = useState("");
    const [problemName, setProblemName] = useState("");
    const [status, setStatus] = useState("-");
    const [language, setLanguage] = useState("-");
    const [score, setScore] = useState("-1");
    const [page, setPage] = useState("0");

    const isScoreFormattedOk = (score) => {
       const nr = parseInt(score);
       return nr >= -1 && nr <= 100;
    }

    const isPageFormattedOk = (page) => {
        const nr = parseInt(page);
        return nr >= 0;
    }

    useEffect(async() => {
        const usernameValue = query.get("username");
        const problemValue  = query.get("problem");
        const statusValue   = query.get("status");
        const languageValue = query.get("lang");
        const scoreValue    = query.get("score"); 
        const pageValue     = query.get("page");

        if (usernameValue !== null && usernameValue !== "")
            setUsername(usernameValue);
        if (problemValue !== null && problemValue !== "")
            setProblemName(problemValue);
        if (statusValue !== null && statusValue !== "")
            setStatus(status);
        if (languageValue !== null && languageValue !== "")
            setLanguage(languageValue);
        if (scoreValue !== null && scoreValue !== "" && isScoreFormattedOk(scoreValue))
            setScore(scoreValue);
        if (pageValue !== null && pageValue !== "" && isPageFormattedOk(pageValue)) 
            setPage(pageValue);

        const q = buildSubmissionQuery();
        await fetchSubmissions(q);
    }, []);

    const fetchSubmissions = async(query) => {
        try {
            const submissions = await submissionAPI.getByQuery(query, page);
    
            if (submissions !== null) {
                setSubmissions(submissions);
            } else {
                setSubmissions([]);
            }
        } catch(err) {
            console.log(err);
        } finally {
            setLoading(false);
        }
    }

    const handleSubmit = async(e) => {
        e.preventDefault();

        const query = buildSubmissionQuery()
        history.push(`/submissions?${query}`)

        try {
            await fetchSubmissions(query);
        } catch(err) {
            console.log(err);
        }
    }

    const buildSubmissionQuery = () => {
        const str = []

        if (username !== "") 
            str.push(`username=${username}`)
        if (problemName !== "")
            str.push(`problem=${problemName}`)
        if (status !== "-")
            str.push(`status=${status}`)
        if (language !== "-")
            str.push(`lang=${language}`)
        if (score >= 0)
            str.push(`score=${score}`)
        if (page > 0)
            str.push(`page=${page}`)

        return str.join("&")
    }

    if (loading) {
        return <Loading/>
    }

    return (
        <div>
            <Navbar color="white" fixed ={false}/> 
            <div style={{marginTop: "100px"}} className={classNames(classes.main, classes.mainRaised)}>
                <div style={{padding: "12px"}}>
                    <Alert severity="error">If your URL contains query search parameters, please click the search button to filter the submissions!! (bug fix coming soon)</Alert>
                </div>
                <Grid container spacing={3} style={{padding: "12px"}}>
                    <Grid item xl={4} md={3} xs={5}>
                       <SubmissionFilter
                            username={username}
                            problemName={problemName}
                            status={status}
                            language={language}
                            score={score}
                            page={page}

                            setUsername={setUsername}
                            setProblemName={setProblemName}
                            setStatus={setStatus}
                            setLanguage={setLanguage}
                            setScore={setScore}
                            setPage={setPage}

                            onSubmit={handleSubmit}
                       />
                    </Grid>
                    <Grid item xl={8} md={9} xs={9} style={{padding: "12px"}}>
                        <SubmissionTable submissions={submissions} />
                    </Grid>
                </Grid>
            </div>
            <Footer/>
        </div>
    );
}