import React, {useState, useEffect} from "react";
import Box from '@material-ui/core/Box';

import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles({
    table: {
      minWidth: 650,
    },
    container: {
        overflow: "auto",
        maxHeight: "420px",
        padding: "20px",
        border: "1px solid #bdbdbd"
    }
});

export default function SubmissionStatus({ submission, submissionTests }) {
    const classes = useStyles();

    const rowStyle = (test) => {
        return test.score == 0 ? {backgroundColor: "#ff3d00"} : {backgroundColor: "#43a047"}
    } 

    return (
        <div>
            <h3>Submission Status</h3>
            {submission.hasCompileError && 
                <>
                    <h4 style={{color:"#bf360c"}}>Compilation Message</h4>
                    <Box style={{border: "2px solid #dd2c00", padding: "12px 16px", marginBottom: "20px"}}>
                        {submission.message}
                    </Box>
                </>
            }
            {submissionTests &&
                <> 
                    <TableContainer component={Paper}>
                        <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                    <TableCell>Test</TableCell>
                                    <TableCell>Time</TableCell>
                                    <TableCell>Memory</TableCell>
                                    <TableCell>Message</TableCell>
                                    <TableCell align="right">Score</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                            {submissionTests && submissionTests.map((test, indx) => (
                                <TableRow key={test.id} style={rowStyle(test)}>
                                    <TableCell component="th" scope="row">{indx + 1}</TableCell>
                                    <TableCell>{test.time.toFixed(3)}</TableCell>
                                    <TableCell>{test.memory} KB</TableCell>
                                    <TableCell>{test.message ? test.message : "-"}</TableCell>
                                    <TableCell align="right">{test.score}</TableCell>
                                </TableRow>
                            ))}
                            </TableBody>
                        </Table>
                    </TableContainer>

                    <h3>Final Score: {submission.score}</h3> 
                </>
            }
        </div>
    );
}