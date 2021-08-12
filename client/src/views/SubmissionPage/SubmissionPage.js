import React, {useState, useEffect} from "react";
import {useParams} from "react-router-dom";

import Navbar from "components/Navbar/Navbar";
import Grid from '@material-ui/core/Grid';

import { makeStyles } from "@material-ui/core/styles";
import styles from "assets/jss/material-kit-react/views/components.js";
import classNames from "classnames";

import submissionAPI from "api/submission";
import submissionTestAPI from "api/submissionTest";
import problemAPI from "api/problem";

const useStyles = makeStyles(styles);

export default function SubmissionPage(props) {

    const [problem, setProblem] = useState({});
    const [submission, setSubmission] = useState({});
    const [submissionTests, setSubmissionTests] = useState([]);

    const classes = useStyles();
    const { submissionId } = useParams();

    const fetchProblem = async (problemName) => {
        try {
            const problem = await problemAPI.getByName(problemName);
            setProblem(problem);
        } catch (err) {
          console.error(err);
        }
    }

    const fetchSubmission = async (submissionId) => {
        try {
           const submission = await submissionAPI.getById(submissionId);
           setSubmission(submission);
        } catch(err) {
            console.error(err);
        }
    }

    const fetchSubmissionTests = async (submissionId) => {
        try {
            const tests = await submissionTestAPI.getBySubmissionId(submissionId);
            setSubmissionTests(tests);
        } catch(err) {
            console.error(err);
        }
    }

    useEffect(async() => {

        await fetchSubmission(submissionId);
        await fetchSubmissionTests(submissionId);
        await fetchProblem(submission.problemName);

        console.log(problem, submission, submissionTests)
    }, []);

    return (
        <div>
            <Navbar color="white" fixed ={false}/> 
            <div style={{marginTop: "100px"}} className={classNames(classes.main, classes.mainRaised)}>
                <Grid container spacing={3} style={{padding: "12px"}}>
                    <Grid item xl={4} md={3} xs={5}>
                        <h3>Submission {submissionId}</h3>
                    </Grid>
                    <Grid item xl={8} md={9} xs={7}>
                        <h3>Submission {submissionId}</h3>
                    </Grid>
                    
                </Grid>
            </div>
        </div>
    );
}