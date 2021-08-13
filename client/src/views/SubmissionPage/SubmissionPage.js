import React, {useState, useEffect} from "react";
import {useParams} from "react-router-dom";

import Navbar from "components/Navbar/Navbar";
import Grid from '@material-ui/core/Grid';

import { makeStyles } from "@material-ui/core/styles";
import styles from "assets/jss/material-kit-react/views/components.js";
import classNames from "classnames";

import submissionAPI from "api/submission";
import submissionTestAPI from "api/submission-test";
import problemAPI from "api/problem";
import ProblemCard from "./Components/ProblemCard";
import SubmissionSourceCode from "./Components/SubmissionSourceCode";
import SubmissionStatus from "./Components/SubmissionStatus";

const useStyles = makeStyles(styles);

export default function SubmissionPage(props) {

    const [problem, setProblem] = useState({});
    const [submission, setSubmission] = useState({});
    const [submissionTests, setSubmissionTests] = useState([]);

    const classes = useStyles();
    const { submissionId } = useParams();  

    useEffect(async() => {

        const submission = await submissionAPI.getById(submissionId);
        const submissionTests = await submissionTestAPI.getBySubmissionId(submissionId);
        const problem = await problemAPI.getByName(submission.problemName);

        setProblem(problem);
        setSubmissionTests(submissionTests);
        setSubmission(submission);

        console.log(problem, submission, submissionTests)
    }, []);

    const authorProfile = (username) => {
        return `/profile/${username}`
    }

    return (
        <div style={{marginBottom: "40px"}}>
            <Navbar color="white" fixed ={false}/> 
            <div style={{marginTop: "100px"}} className={classNames(classes.main, classes.mainRaised)}>
                <Grid container spacing={1} style={{padding: "12px"}}>
                    <Grid item xl={2} md={3} xs={12} style={{textAlign: "center"}}>
                        <h3>Submission #{submission.id}</h3>
                        <h3>Author: {"  "}
                            <a style={{color: "inherit", textDecoration: "underline"}} href={authorProfile(submission.username)}>
                                {submission.username}
                            </a>    
                        </h3>
                        <h3>Score: {submission.score}</h3>
                    </Grid>
                    <Grid item xl={10} md={9} xs={12}>
                        <ProblemCard 
                            problem={problem} 
                            submission={submission}
                        />
                        <SubmissionStatus
                            submission={submission}
                            submissionTests={submissionTests}
                        />
                        <SubmissionSourceCode 
                            problem={problem} 
                            submission={submission}
                        />
                    </Grid>
                </Grid>
            </div>
        </div>
    );
}