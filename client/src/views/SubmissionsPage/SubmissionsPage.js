import React, {useState, useEffect} from "react";

import Navbar from "components/Navbar/Navbar";
import Grid from '@material-ui/core/Grid';

import { makeStyles } from "@material-ui/core/styles";
import classNames from "classnames";
import styles from "assets/jss/material-kit-react/views/components.js";
import SubmissionFilter from "./Components/SubmissionFilter";
import submissionService from "services/submission.service";
import { useHistory } from "react-router-dom";
import SubmissionTable from "./Components/SubmissionTable";
import Footer from "components/Footer/Footer";
import useQuery from "hooks/query";

const useStyles = makeStyles(styles);

export default function SubmissionsPage() {
    const query = useQuery();
    const classes = useStyles();
    const history = useHistory();

    const [submissions, setSubmissions] = useState([]);

    const [username, setUsername] = useState("");
    const [problemName, setProblemName] = useState("");
    const [status, setStatus] = useState("-");
    const [language, setLanguage] = useState("-");
    const [score, setScore] = useState(-1);
    const [page, setPage] = useState(0);

    useEffect(() => {
        const usernameValue = query.get("username");
        const problemValue  = query.get("problem");
        const statusValue   = query.get("status");
        const languageValue = query.get("lang");
        const scoreValue    = query.get("score"); 
        const pageValue     = query.get("page");

        if (usernameValue !== null)
            setUsername(usernameValue);
        if (problemValue !== null)
            setProblemName(problemValue);
        if (statusValue !== null)
            setStatus(status);
        if (languageValue !== null)
            setLanguage(languageValue);
        if (scoreValue !== null) {
            const nr = parseInt(scoreValue);
            console.log(nr, typeof nr);
            setScore(nr);
        }
        if (pageValue !== null) {
            const nr = parseInt(pageValue);
            setPage(nr);
        }
    }, []);

    const fetchSubmissions = async() => {
        try {
            const query = buildSubmissionQuery();
            console.log("QUERY + PAGE: ", query, page);
            const res = await submissionService.getSubmissions(query, page);
    
            if (res.data != null) {
                setSubmissions(res.data);
            } else {
                setSubmissions([]);
            }

        } catch(err) {
            console.log(err);
        }
    }

    const handleSubmit = async(e) => {
        e.preventDefault();

        const query = buildSubmissionQuery()
        history.push(`/submissions?${query}`)

        try {
            await fetchSubmissions();
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

    return (
        <div>
            <Navbar color="white" fixed ={false}/> 
            <div style={{marginTop: "100px"}} className={classNames(classes.main, classes.mainRaised)}>
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