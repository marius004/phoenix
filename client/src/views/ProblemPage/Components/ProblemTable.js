import React, {useState, useEffect} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import PropTypes from "prop-types";
import userService from 'services/user.service';
import { Link } from 'react-router-dom';

const useStyles = makeStyles({
  gravatar: {
    borderRadius: "50%"
  }
});

export default function ProblemTable({ data }) {
  const classes = useStyles();
  const [gravatarData, setGravatarData] = useState({})

  const fetchGravatarData = async() => {
    try {
      const res = await userService.getGravatarEmailHash(data.authorId);
      setGravatarData(res.data);
    } catch(err) {
      console.error(err);
    }
  }

  const getGravatarURI = (imgSize) => {
    return `https://www.gravatar.com/avatar/${gravatarData.emailHash}?s=${imgSize}`;
  }

  useEffect(fetchGravatarData, [])

  return (
    <TableContainer component={Paper} style={{marginBottom: "20px"}}>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Posted by</TableCell>
            <TableCell align="right">Grade</TableCell>
            <TableCell align="right">Input/Output</TableCell>
            <TableCell align="right">Time Limit</TableCell>
            <TableCell align="right">Memory Limit</TableCell>
            <TableCell align="right">Stack Limit</TableCell>
            <TableCell align="right">Difficulty</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
        <TableRow key={data.id}>
            <TableCell component="th" scope="row">
              <Link to={() => `/profile/${gravatarData.username}`} style={{color: "blue"}}>
                <img className={classes.gravatar} src={getGravatarURI(22)}/>
                {"  "}{gravatarData.username} 
              </Link>
            </TableCell>
            <TableCell align="right">{data.grade}</TableCell>
            <TableCell align="right">{data.stream}</TableCell>
            <TableCell align="right">{data.timeLimit} s</TableCell>
            <TableCell align="right">{data.memoryLimit} KB</TableCell>
            <TableCell align="right">{data.stackLimit} KB</TableCell>
            <TableCell align="right">{data.difficulty}</TableCell>
        </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
}

ProblemTable.prototype = {
  data: PropTypes.object.isRequired
}