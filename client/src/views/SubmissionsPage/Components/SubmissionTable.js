import React, {useEffect} from 'react';
import { withStyles, makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import PropTypes from "prop-types";
import { Link } from 'react-router-dom';

const StyledTableCell = withStyles((theme) => ({
  head: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  body: {
    fontSize: 14,
  },
}))(TableCell);

const StyledTableRow = withStyles((theme) => ({
  root: {
    '&:nth-of-type(odd)': {
      backgroundColor: theme.palette.action.hover,
    },
  },
}))(TableRow);

export default function SubmissionTable({ submissions }) {
    const submissionDate = (time) => {
      const date = new Date(time);
      
      const year  = date.getFullYear();
      const month = date.getMonth() < 10 ? "0" + date.getMonth() : date.getMonth();
      const day   = date.getDay() < 10 ? "0" + date.getDay() : date.getDay();

      return day + "/" + month + "/" + year;
    }

    const submissionTime = (time) => {
      const date = new Date(time);
      
      const hour = date.getHours() < 10 ? "0" +  date.getHours() : date.getHours();
      const min  = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();

      return hour + ":" + min;
    }
  
    const submissionStatus = (submission) => {
      if (submission.status == "waiting")
          return "Waiting";
      if(submission.status == "working")
          return "Evaluating";
      if (submission.hasCompileError)
          return "Compilation Error";
      return `Evaluated: ${submission.score}`
    }

    if (submissions.length === 0) {
      return <h3 style={{textAlign: "center"}}>No submission</h3>
    }

   return (
    <div>
      <p style={{fontSize: "28px"}}> {submissions.length} records </p>
      <TableContainer component={Paper}>
        <Table aria-label="customized table">
          <TableHead>
            <TableRow>
              <StyledTableCell>ID</StyledTableCell>
              <StyledTableCell>User</StyledTableCell>
              <StyledTableCell>Date</StyledTableCell>
              <StyledTableCell>Problem</StyledTableCell>
              <StyledTableCell align="right">Status</StyledTableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {submissions.map((row) => (
              <StyledTableRow key={row.id}>
                {/* TODO add link to submission */}
                <StyledTableCell component="th" scope="row">
                  {row.id}
                </StyledTableCell>
                <StyledTableCell>
                    <img src={`https://www.gravatar.com/avatar/${row.emailHash}?s=25`} alt="user icon"/> {"   "}
                    {row.username}
                </StyledTableCell>
                <StyledTableCell>
                  {submissionDate(row.createdAt)} 
                  {"    "}
                  {submissionTime(row.createdAt)}
                </StyledTableCell>
                <StyledTableCell>
                  <Link to={`/problems/${row.problemName}`} style={{color: "black", textDecoration: "underline"}}>
                    {row.problemName}
                  </Link>
                </StyledTableCell>
                <StyledTableCell align="right">{submissionStatus(row)}</StyledTableCell>
              </StyledTableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
}

SubmissionTable.propTypes = {
    submissions: PropTypes.array.isRequired,
};