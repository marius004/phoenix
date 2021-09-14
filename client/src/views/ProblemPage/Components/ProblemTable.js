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
import { Link } from 'react-router-dom';

import userAPI from 'api/user';
import avatarAPI from 'api/avatar';

const useStyles = makeStyles({
  gravatar: {
    borderRadius: "50%"
  }
});

export default function ProblemTable({ data }) {
  const classes = useStyles();
  const [avatar, setAvatar] = useState({image: "", username: ""})

  const fetchGravatarData = async() => {
    try {
      const avatar = await avatarAPI.get(data.authorId, 25);
      setAvatar(avatar);
    } catch(err) {
      console.error(err);
    }
  }

  useEffect(fetchGravatarData, [])

  return (
    <TableContainer component={Paper} style={{marginBottom: "20px"}}>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Posted by</TableCell>
            <TableCell align="right">Grade</TableCell>
            <TableCell align="right">ID</TableCell>
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
              <Link to={() => `/profile/${avatar.username}`} style={{color: "blue"}}>
              <img src={`data:image/png;base64,${avatar.image}`} alt="user icon"/> {"   "}
                {"  "}{avatar.username} 
              </Link>
            </TableCell>
            <TableCell align="right">{data.grade}</TableCell>
            <TableCell align="right">{data.id}</TableCell>
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